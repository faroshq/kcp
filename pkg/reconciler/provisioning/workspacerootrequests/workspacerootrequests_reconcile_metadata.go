/*
Copyright 2021 The KCP Authors.

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

	"github.com/davecgh/go-spew/spew"
	"github.com/kcp-dev/logicalcluster/v3"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kcp-dev/kcp/pkg/admission/workspacetypeexists"
	indexrewriters "github.com/kcp-dev/kcp/pkg/index/rewriters"
	reconcilerworkspace "github.com/kcp-dev/kcp/pkg/reconciler/tenancy/workspace"
	"github.com/kcp-dev/kcp/sdk/apis/core"
	corev1alpha1 "github.com/kcp-dev/kcp/sdk/apis/core/v1alpha1"
	provisioningv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/provisioning/v1alpha1"
	tenancyv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/tenancy/v1alpha1"
)

type logicalClusterCreator struct {
	getLogicalCluster    func(ctx context.Context, cluster logicalcluster.Name, name string) (*corev1alpha1.LogicalCluster, error)
	createLogicalCluster func(ctx context.Context, cluster logicalcluster.Name, logicalCluster *corev1alpha1.LogicalCluster) (*corev1alpha1.LogicalCluster, error)

	getWorkspaceType       func(clusterName logicalcluster.Path, name string) (*tenancyv1alpha1.WorkspaceType, error)
	transitiveTypeResolver workspacetypeexists.TransitiveTypeResolver
}

func (r *logicalClusterCreator) reconcile(ctx context.Context, request *provisioningv1alpha1.WorkspaceRootRequest) (reconcileStatus, error) {

	request.Status.Phase = corev1alpha1.LogicalClusterPhaseScheduling
	homeClusterName := indexrewriters.HomeClusterName(request.Name)

	current, err := r.getLogicalCluster(ctx, homeClusterName, corev1alpha1.LogicalClusterName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			logicalCluster := &corev1alpha1.LogicalCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1alpha1.LogicalClusterName,
					Annotations: map[string]string{
						//tenancyv1alpha1.ExperimentalWorkspaceOwnerAnnotationKey: userInfo,
						tenancyv1alpha1.LogicalClusterTypeAnnotationKey: "root:faros",
						core.LogicalClusterPathAnnotationKey:            fmt.Sprintf("user:%s", request.Name),
					},
				},
			}

			logicalCluster.Spec.Initializers, err = reconcilerworkspace.LogicalClustersInitializers(r.transitiveTypeResolver, r.getWorkspaceType, core.RootCluster.Path(), "home")
			if err != nil {
				return reconcileStatusStopAndRequeue, err
			}

			current, err = r.createLogicalCluster(ctx, homeClusterName, logicalCluster)
			if err != nil {
				return reconcileStatusStopAndRequeue, err
			}
		}
	}

	spew.Dump(current)

	request.Status.URL = current.Status.URL
	request.Status.Phase = current.Status.Phase
	request.Status.Conditions = current.Status.Conditions

	return reconcileStatusContinue, nil
}
