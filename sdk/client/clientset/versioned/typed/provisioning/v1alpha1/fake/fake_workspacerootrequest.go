/*
Copyright The KCP Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kcp-dev/kcp/sdk/apis/provisioning/v1alpha1"
	provisioningv1alpha1 "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/provisioning/v1alpha1"
)

// FakeWorkspaceRootRequests implements WorkspaceRootRequestInterface
type FakeWorkspaceRootRequests struct {
	Fake *FakeProvisioningV1alpha1
}

var workspacerootrequestsResource = v1alpha1.SchemeGroupVersion.WithResource("workspacerootrequests")

var workspacerootrequestsKind = v1alpha1.SchemeGroupVersion.WithKind("WorkspaceRootRequest")

// Get takes name of the workspaceRootRequest, and returns the corresponding workspaceRootRequest object, and an error if there is any.
func (c *FakeWorkspaceRootRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WorkspaceRootRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(workspacerootrequestsResource, name), &v1alpha1.WorkspaceRootRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceRootRequest), err
}

// List takes label and field selectors, and returns the list of WorkspaceRootRequests that match those selectors.
func (c *FakeWorkspaceRootRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WorkspaceRootRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(workspacerootrequestsResource, workspacerootrequestsKind, opts), &v1alpha1.WorkspaceRootRequestList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WorkspaceRootRequestList{ListMeta: obj.(*v1alpha1.WorkspaceRootRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.WorkspaceRootRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workspaceRootRequests.
func (c *FakeWorkspaceRootRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(workspacerootrequestsResource, opts))
}

// Create takes the representation of a workspaceRootRequest and creates it.  Returns the server's representation of the workspaceRootRequest, and an error, if there is any.
func (c *FakeWorkspaceRootRequests) Create(ctx context.Context, workspaceRootRequest *v1alpha1.WorkspaceRootRequest, opts v1.CreateOptions) (result *v1alpha1.WorkspaceRootRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(workspacerootrequestsResource, workspaceRootRequest), &v1alpha1.WorkspaceRootRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceRootRequest), err
}

// Update takes the representation of a workspaceRootRequest and updates it. Returns the server's representation of the workspaceRootRequest, and an error, if there is any.
func (c *FakeWorkspaceRootRequests) Update(ctx context.Context, workspaceRootRequest *v1alpha1.WorkspaceRootRequest, opts v1.UpdateOptions) (result *v1alpha1.WorkspaceRootRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(workspacerootrequestsResource, workspaceRootRequest), &v1alpha1.WorkspaceRootRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceRootRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWorkspaceRootRequests) UpdateStatus(ctx context.Context, workspaceRootRequest *v1alpha1.WorkspaceRootRequest, opts v1.UpdateOptions) (*v1alpha1.WorkspaceRootRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(workspacerootrequestsResource, "status", workspaceRootRequest), &v1alpha1.WorkspaceRootRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceRootRequest), err
}

// Delete takes name of the workspaceRootRequest and deletes it. Returns an error if one occurs.
func (c *FakeWorkspaceRootRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(workspacerootrequestsResource, name, opts), &v1alpha1.WorkspaceRootRequest{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkspaceRootRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(workspacerootrequestsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.WorkspaceRootRequestList{})
	return err
}

// Patch applies the patch and returns the patched workspaceRootRequest.
func (c *FakeWorkspaceRootRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkspaceRootRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(workspacerootrequestsResource, name, pt, data, subresources...), &v1alpha1.WorkspaceRootRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceRootRequest), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied workspaceRootRequest.
func (c *FakeWorkspaceRootRequests) Apply(ctx context.Context, workspaceRootRequest *provisioningv1alpha1.WorkspaceRootRequestApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.WorkspaceRootRequest, err error) {
	if workspaceRootRequest == nil {
		return nil, fmt.Errorf("workspaceRootRequest provided to Apply must not be nil")
	}
	data, err := json.Marshal(workspaceRootRequest)
	if err != nil {
		return nil, err
	}
	name := workspaceRootRequest.Name
	if name == nil {
		return nil, fmt.Errorf("workspaceRootRequest.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(workspacerootrequestsResource, *name, types.ApplyPatchType, data), &v1alpha1.WorkspaceRootRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceRootRequest), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeWorkspaceRootRequests) ApplyStatus(ctx context.Context, workspaceRootRequest *provisioningv1alpha1.WorkspaceRootRequestApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.WorkspaceRootRequest, err error) {
	if workspaceRootRequest == nil {
		return nil, fmt.Errorf("workspaceRootRequest provided to Apply must not be nil")
	}
	data, err := json.Marshal(workspaceRootRequest)
	if err != nil {
		return nil, err
	}
	name := workspaceRootRequest.Name
	if name == nil {
		return nil, fmt.Errorf("workspaceRootRequest.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(workspacerootrequestsResource, *name, types.ApplyPatchType, data, "status"), &v1alpha1.WorkspaceRootRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceRootRequest), err
}
