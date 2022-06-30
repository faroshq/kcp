//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	conditionsv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/third_party/conditions/apis/conditions/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspace) DeepCopyInto(out *ClusterWorkspace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspace.
func (in *ClusterWorkspace) DeepCopy() *ClusterWorkspace {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterWorkspace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceList) DeepCopyInto(out *ClusterWorkspaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterWorkspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceList.
func (in *ClusterWorkspaceList) DeepCopy() *ClusterWorkspaceList {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterWorkspaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceLocation) DeepCopyInto(out *ClusterWorkspaceLocation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceLocation.
func (in *ClusterWorkspaceLocation) DeepCopy() *ClusterWorkspaceLocation {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceShard) DeepCopyInto(out *ClusterWorkspaceShard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceShard.
func (in *ClusterWorkspaceShard) DeepCopy() *ClusterWorkspaceShard {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceShard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterWorkspaceShard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceShardList) DeepCopyInto(out *ClusterWorkspaceShardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterWorkspaceShard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceShardList.
func (in *ClusterWorkspaceShardList) DeepCopy() *ClusterWorkspaceShardList {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceShardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterWorkspaceShardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceShardSpec) DeepCopyInto(out *ClusterWorkspaceShardSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceShardSpec.
func (in *ClusterWorkspaceShardSpec) DeepCopy() *ClusterWorkspaceShardSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceShardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceShardStatus) DeepCopyInto(out *ClusterWorkspaceShardStatus) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceShardStatus.
func (in *ClusterWorkspaceShardStatus) DeepCopy() *ClusterWorkspaceShardStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceShardStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceSpec) DeepCopyInto(out *ClusterWorkspaceSpec) {
	*out = *in
	out.Type = in.Type
	if in.Shard != nil {
		in, out := &in.Shard, &out.Shard
		*out = new(ShardConstraints)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceSpec.
func (in *ClusterWorkspaceSpec) DeepCopy() *ClusterWorkspaceSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceStatus) DeepCopyInto(out *ClusterWorkspaceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Location = in.Location
	if in.Initializers != nil {
		in, out := &in.Initializers, &out.Initializers
		*out = make([]ClusterWorkspaceInitializer, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceStatus.
func (in *ClusterWorkspaceStatus) DeepCopy() *ClusterWorkspaceStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceType) DeepCopyInto(out *ClusterWorkspaceType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceType.
func (in *ClusterWorkspaceType) DeepCopy() *ClusterWorkspaceType {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterWorkspaceType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceTypeExtension) DeepCopyInto(out *ClusterWorkspaceTypeExtension) {
	*out = *in
	if in.With != nil {
		in, out := &in.With, &out.With
		*out = make([]ClusterWorkspaceTypeReference, len(*in))
		copy(*out, *in)
	}
	if in.Without != nil {
		in, out := &in.Without, &out.Without
		*out = make([]ClusterWorkspaceTypeReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceTypeExtension.
func (in *ClusterWorkspaceTypeExtension) DeepCopy() *ClusterWorkspaceTypeExtension {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceTypeExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceTypeList) DeepCopyInto(out *ClusterWorkspaceTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterWorkspaceType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceTypeList.
func (in *ClusterWorkspaceTypeList) DeepCopy() *ClusterWorkspaceTypeList {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterWorkspaceTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceTypeReference) DeepCopyInto(out *ClusterWorkspaceTypeReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceTypeReference.
func (in *ClusterWorkspaceTypeReference) DeepCopy() *ClusterWorkspaceTypeReference {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceTypeReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceTypeSpec) DeepCopyInto(out *ClusterWorkspaceTypeSpec) {
	*out = *in
	in.Extend.DeepCopyInto(&out.Extend)
	if in.AdditionalWorkspaceLabels != nil {
		in, out := &in.AdditionalWorkspaceLabels, &out.AdditionalWorkspaceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AllowedChildWorkspaceTypes != nil {
		in, out := &in.AllowedChildWorkspaceTypes, &out.AllowedChildWorkspaceTypes
		*out = make([]ClusterWorkspaceTypeName, len(*in))
		copy(*out, *in)
	}
	if in.AllowedParentWorkspaceTypes != nil {
		in, out := &in.AllowedParentWorkspaceTypes, &out.AllowedParentWorkspaceTypes
		*out = make([]ClusterWorkspaceTypeName, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceTypeSpec.
func (in *ClusterWorkspaceTypeSpec) DeepCopy() *ClusterWorkspaceTypeSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkspaceTypeStatus) DeepCopyInto(out *ClusterWorkspaceTypeStatus) {
	*out = *in
	if in.Initializers != nil {
		in, out := &in.Initializers, &out.Initializers
		*out = make([]ClusterWorkspaceInitializer, len(*in))
		copy(*out, *in)
	}
	if in.TypeAliases != nil {
		in, out := &in.TypeAliases, &out.TypeAliases
		*out = make([]ClusterWorkspaceTypeReference, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VirtualWorkspaces != nil {
		in, out := &in.VirtualWorkspaces, &out.VirtualWorkspaces
		*out = make([]VirtualWorkspace, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkspaceTypeStatus.
func (in *ClusterWorkspaceTypeStatus) DeepCopy() *ClusterWorkspaceTypeStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkspaceTypeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardConstraints) DeepCopyInto(out *ShardConstraints) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardConstraints.
func (in *ShardConstraints) DeepCopy() *ShardConstraints {
	if in == nil {
		return nil
	}
	out := new(ShardConstraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualWorkspace) DeepCopyInto(out *VirtualWorkspace) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualWorkspace.
func (in *VirtualWorkspace) DeepCopy() *VirtualWorkspace {
	if in == nil {
		return nil
	}
	out := new(VirtualWorkspace)
	in.DeepCopyInto(out)
	return out
}
