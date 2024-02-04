/*
Copyright 2022 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package workspacerootrequests

import (
	"context"
	"fmt"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"

	"github.com/kcp-dev/kcp/pkg/logging"
	"github.com/kcp-dev/kcp/pkg/reconciler/committer"
	provisioningv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/provisioning/v1alpha1"
	kcpclientset "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/cluster"
	provisioningv1alpha1client "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/provisioning/v1alpha1"
	provisioningv1alpha1informers "github.com/kcp-dev/kcp/sdk/client/informers/externalversions/provisioning/v1alpha1"
	provisioningv1alpha1listers "github.com/kcp-dev/kcp/sdk/client/listers/provisioning/v1alpha1"
)

const (
	ControllerName = "kcp-workspacerootrequests"
)

func NewController(
	kcpClusterClient kcpclientset.ClusterInterface,
	workspaceRootRequestInformer provisioningv1alpha1informers.WorkspaceRootRequestClusterInformer,
) (*Controller, error) {
	queue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), ControllerName)

	c := &Controller{
		queue:                       queue,
		kcpClusterClient:            kcpClusterClient,
		workspaceRootRequestIndexer: workspaceRootRequestInformer.Informer().GetIndexer(),
		workspaceRootRequestLister:  workspaceRootRequestInformer.Lister(),
		commit:                      committer.NewCommitter[*provisioningv1alpha1.WorkspaceRootRequest, provisioningv1alpha1client.WorkspaceRootRequestInterface, *provisioningv1alpha1.WorkspaceRootRequestSpec, *provisioningv1alpha1.WorkspaceRootRequestStatus](kcpClusterClient.ProvisioningV1alpha1().WorkspaceRootRequests()),
	}
	_, _ = workspaceRootRequestInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    func(obj interface{}) { c.enqueue(obj) },
		UpdateFunc: func(obj, _ interface{}) { c.enqueue(obj) },
		DeleteFunc: func(obj interface{}) { c.enqueue(obj) },
	})

	return c, nil
}

type workspaceRootRequestResource = committer.Resource[*provisioningv1alpha1.WorkspaceRootRequestSpec, *provisioningv1alpha1.WorkspaceRootRequestStatus]

// Controller watches Workspaces and WorkspaceShards in order to make sure every workspace
// is scheduled to a valid Shard.
type Controller struct {
	queue workqueue.RateLimitingInterface

	kcpClusterClient kcpclientset.ClusterInterface

	workspaceRootRequestIndexer cache.Indexer
	workspaceRootRequestLister  provisioningv1alpha1listers.WorkspaceRootRequestClusterLister

	// commit creates a patch and submits it, if needed.
	commit func(ctx context.Context, old, new *workspaceRootRequestResource) error
}

func (c *Controller) enqueue(obj interface{}) {
	key, err := kcpcache.DeletionHandlingMetaClusterNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
		return
	}
	logger := logging.WithQueueKey(logging.WithReconciler(klog.Background(), ControllerName), key)
	logger.V(2).Info("queueing WorkspaceRootRequest")
	c.queue.Add(key)
}

func (c *Controller) Start(ctx context.Context, numThreads int) {
	defer runtime.HandleCrash()
	defer c.queue.ShutDown()

	logger := logging.WithReconciler(klog.FromContext(ctx), ControllerName)
	ctx = klog.NewContext(ctx, logger)
	logger.Info("Starting controller")
	defer logger.Info("Shutting down controller")

	for i := 0; i < numThreads; i++ {
		go wait.Until(func() { c.startWorker(ctx) }, time.Second, ctx.Done())
	}

	<-ctx.Done()
}

func (c *Controller) startWorker(ctx context.Context) {
	for c.processNextWorkItem(ctx) {
	}
}

func (c *Controller) processNextWorkItem(ctx context.Context) bool {
	// Wait until there is a new item in the working queue
	k, quit := c.queue.Get()
	if quit {
		return false
	}
	key := k.(string)

	logger := logging.WithQueueKey(klog.FromContext(ctx), key)
	ctx = klog.NewContext(ctx, logger)
	logger.V(1).Info("processing key")

	// No matter what, tell the queue we're done with this key, to unblock
	// other workers.
	defer c.queue.Done(key)

	if requeue, err := c.process(ctx, key); err != nil {
		runtime.HandleError(fmt.Errorf("%q controller failed to sync %q, err: %w", ControllerName, key, err))
		c.queue.AddRateLimited(key)
		return true
	} else if requeue {
		// only requeue if we didn't error, but we still want to requeue
		c.queue.Add(key)
		return true
	}
	c.queue.Forget(key)
	return true
}

func (c *Controller) process(ctx context.Context, key string) (bool, error) {
	logger := klog.FromContext(ctx)

	clusterName, _, name, err := kcpcache.SplitMetaClusterNamespaceKey(key)
	if err != nil {
		logger.Error(err, "unable to decode key")
		return false, nil
	}

	workspaceRootRequest, err := c.workspaceRootRequestLister.Cluster(clusterName).Get(name)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			logger.Error(err, "failed to get WorkspaceRootRequest from lister", "cluster", clusterName)
		}

		return false, nil // nothing we can do here
	}

	old := workspaceRootRequest
	workspaceRootRequest = workspaceRootRequest.DeepCopy()

	logger = logging.WithObject(logger, workspaceRootRequest)
	ctx = klog.NewContext(ctx, logger)

	var errs []error
	requeue, err := c.reconcile(ctx, workspaceRootRequest)
	if err != nil {
		errs = append(errs, err)
	}

	// If the object being reconciled changed as a result, update it.
	oldResource := &workspaceRootRequestResource{ObjectMeta: old.ObjectMeta, Spec: &old.Spec, Status: &old.Status}
	newResource := &workspaceRootRequestResource{ObjectMeta: workspaceRootRequest.ObjectMeta, Spec: &workspaceRootRequest.Spec, Status: &workspaceRootRequest.Status}
	if err := c.commit(ctx, oldResource, newResource); err != nil {
		errs = append(errs, err)
	}

	return requeue, utilerrors.NewAggregate(errs)
}
