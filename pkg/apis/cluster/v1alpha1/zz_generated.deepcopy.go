//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright Kurator Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	v1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in CIDRBlocks) DeepCopyInto(out *CIDRBlocks) {
	{
		in := &in
		*out = make(CIDRBlocks, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CIDRBlocks.
func (in CIDRBlocks) DeepCopy() CIDRBlocks {
	if in == nil {
		return nil
	}
	out := new(CIDRBlocks)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNIConfig) DeepCopyInto(out *CNIConfig) {
	*out = *in
	in.ExtraArgs.DeepCopyInto(&out.ExtraArgs)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNIConfig.
func (in *CNIConfig) DeepCopy() *CNIConfig {
	if in == nil {
		return nil
	}
	out := new(CNIConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (c *Cluster) DeepCopyInto(out *Cluster) {
	*out = *c
	out.TypeMeta = c.TypeMeta
	c.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	c.Spec.DeepCopyInto(&out.Spec)
	c.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (c *Cluster) DeepCopy() *Cluster {
	if c == nil {
		return nil
	}
	out := new(Cluster)
	c.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (c *Cluster) DeepCopyObject() runtime.Object {
	if c := c.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.Credential != nil {
		in, out := &in.Credential, &out.Credential
		*out = new(CredentialConfig)
		**out = **in
	}
	in.Network.DeepCopyInto(&out.Network)
	in.Master.DeepCopyInto(&out.Master)
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = make([]WorkerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.PodIdentity = in.PodIdentity
	if in.AdditionalResources != nil {
		in, out := &in.AdditionalResources, &out.AdditionalResources
		*out = make([]ResourceRef, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialConfig) DeepCopyInto(out *CredentialConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialConfig.
func (in *CredentialConfig) DeepCopy() *CredentialConfig {
	if in == nil {
		return nil
	}
	out := new(CredentialConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfig) DeepCopyInto(out *MachineConfig) {
	*out = *in
	if in.RootVolume != nil {
		in, out := &in.RootVolume, &out.RootVolume
		*out = new(Volume)
		**out = **in
	}
	if in.NonRootVolumes != nil {
		in, out := &in.NonRootVolumes, &out.NonRootVolumes
		*out = make([]Volume, len(*in))
		copy(*out, *in)
	}
	in.ExtraArgs.DeepCopyInto(&out.ExtraArgs)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfig.
func (in *MachineConfig) DeepCopy() *MachineConfig {
	if in == nil {
		return nil
	}
	out := new(MachineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterConfig) DeepCopyInto(out *MasterConfig) {
	*out = *in
	in.NodeConfig.DeepCopyInto(&out.NodeConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterConfig.
func (in *MasterConfig) DeepCopy() *MasterConfig {
	if in == nil {
		return nil
	}
	out := new(MasterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfig) DeepCopyInto(out *NetworkConfig) {
	*out = *in
	out.VPC = in.VPC
	if in.PodCIDRs != nil {
		in, out := &in.PodCIDRs, &out.PodCIDRs
		*out = make(CIDRBlocks, len(*in))
		copy(*out, *in)
	}
	if in.ServiceCIDRs != nil {
		in, out := &in.ServiceCIDRs, &out.ServiceCIDRs
		*out = make(CIDRBlocks, len(*in))
		copy(*out, *in)
	}
	in.CNI.DeepCopyInto(&out.CNI)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfig.
func (in *NetworkConfig) DeepCopy() *NetworkConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfig) DeepCopyInto(out *NodeConfig) {
	*out = *in
	in.MachineConfig.DeepCopyInto(&out.MachineConfig)
	in.NodeRegistrationConfig.DeepCopyInto(&out.NodeRegistrationConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfig.
func (in *NodeConfig) DeepCopy() *NodeConfig {
	if in == nil {
		return nil
	}
	out := new(NodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRegistrationConfig) DeepCopyInto(out *NodeRegistrationConfig) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]v1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRegistrationConfig.
func (in *NodeRegistrationConfig) DeepCopy() *NodeRegistrationConfig {
	if in == nil {
		return nil
	}
	out := new(NodeRegistrationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeUpgradeStrategy) DeepCopyInto(out *NodeUpgradeStrategy) {
	*out = *in
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = new(RollingUpdateNodeUpgradeStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeUpgradeStrategy.
func (in *NodeUpgradeStrategy) DeepCopy() *NodeUpgradeStrategy {
	if in == nil {
		return nil
	}
	out := new(NodeUpgradeStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodIdentityConfig) DeepCopyInto(out *PodIdentityConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodIdentityConfig.
func (in *PodIdentityConfig) DeepCopy() *PodIdentityConfig {
	if in == nil {
		return nil
	}
	out := new(PodIdentityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRef) DeepCopyInto(out *ResourceRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRef.
func (in *ResourceRef) DeepCopy() *ResourceRef {
	if in == nil {
		return nil
	}
	out := new(ResourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateNodeUpgradeStrategy) DeepCopyInto(out *RollingUpdateNodeUpgradeStrategy) {
	*out = *in
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxSurge != nil {
		in, out := &in.MaxSurge, &out.MaxSurge
		*out = new(intstr.IntOrString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateNodeUpgradeStrategy.
func (in *RollingUpdateNodeUpgradeStrategy) DeepCopy() *RollingUpdateNodeUpgradeStrategy {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateNodeUpgradeStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCConfig) DeepCopyInto(out *VPCConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCConfig.
func (in *VPCConfig) DeepCopy() *VPCConfig {
	if in == nil {
		return nil
	}
	out := new(VPCConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerConfig) DeepCopyInto(out *WorkerConfig) {
	*out = *in
	in.NodeConfig.DeepCopyInto(&out.NodeConfig)
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(NodeUpgradeStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerConfig.
func (in *WorkerConfig) DeepCopy() *WorkerConfig {
	if in == nil {
		return nil
	}
	out := new(WorkerConfig)
	in.DeepCopyInto(out)
	return out
}
