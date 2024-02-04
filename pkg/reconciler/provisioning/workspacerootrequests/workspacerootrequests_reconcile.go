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

	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilserrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/kcp-dev/kcp/pkg/admission/workspacetypeexists"
	"github.com/kcp-dev/kcp/pkg/indexers"
	corev1alpha1 "github.com/kcp-dev/kcp/sdk/apis/core/v1alpha1"
	provisioningv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/provisioning/v1alpha1"
	tenancyv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/tenancy/v1alpha1"
)

type reconcileStatus int

const (
	reconcileStatusStopAndRequeue reconcileStatus = iota
	reconcileStatusContinue
)

type reconciler interface {
	reconcile(ctx context.Context, request *provisioningv1alpha1.WorkspaceRootRequest) (reconcileStatus, error)
}

func (c *Controller) getWorkspaceType(path logicalcluster.Path, name string) (*tenancyv1alpha1.WorkspaceType, error) {
	return indexers.ByPathAndName[*tenancyv1alpha1.WorkspaceType](tenancyv1alpha1.Resource("workspacetypes"), c.workspaceTypeIndexer, path, name)
}

func (c *Controller) reconcile(ctx context.Context, request *provisioningv1alpha1.WorkspaceRootRequest) (bool, error) {
	reconcilers := []reconciler{
		&logicalClusterCreator{
			getLogicalCluster: func(ctx context.Context, cluster logicalcluster.Name, name string) (*corev1alpha1.LogicalCluster, error) {
				return c.kcpClusterClient.Cluster(cluster.Path()).CoreV1alpha1().LogicalClusters().Get(ctx, name, metav1.GetOptions{})
			},
			createLogicalCluster: func(ctx context.Context, cluster logicalcluster.Name, logicalCluster *corev1alpha1.LogicalCluster) (*corev1alpha1.LogicalCluster, error) {
				return c.kcpClusterClient.Cluster(cluster.Path()).CoreV1alpha1().LogicalClusters().Create(ctx, logicalCluster, metav1.CreateOptions{})
			},
			transitiveTypeResolver: workspacetypeexists.NewTransitiveTypeResolver(c.getWorkspaceType),
			getWorkspaceType:       c.getWorkspaceType,
		},
	}

	var errs []error

	requeue := false
	for _, r := range reconcilers {
		var err error
		var status reconcileStatus
		status, err = r.reconcile(ctx, request)
		if err != nil {
			errs = append(errs, err)
		}
		if status == reconcileStatusStopAndRequeue {
			requeue = true
			break
		}
	}

	return requeue, utilserrors.NewAggregate(errs)
}
