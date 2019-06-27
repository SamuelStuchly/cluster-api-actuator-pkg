// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAutoscaler) DeepCopyInto(out *ClusterAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAutoscaler.
func (in *ClusterAutoscaler) DeepCopy() *ClusterAutoscaler {
	if in == nil {
		return nil
	}
	out := new(ClusterAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAutoscalerList) DeepCopyInto(out *ClusterAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAutoscalerList.
func (in *ClusterAutoscalerList) DeepCopy() *ClusterAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(ClusterAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAutoscalerSpec) DeepCopyInto(out *ClusterAutoscalerSpec) {
	*out = *in
	if in.ResourceLimits != nil {
		in, out := &in.ResourceLimits, &out.ResourceLimits
		*out = new(ResourceLimits)
		(*in).DeepCopyInto(*out)
	}
	if in.ScaleDown != nil {
		in, out := &in.ScaleDown, &out.ScaleDown
		*out = new(ScaleDownConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxPodGracePeriod != nil {
		in, out := &in.MaxPodGracePeriod, &out.MaxPodGracePeriod
		*out = new(int32)
		**out = **in
	}
	if in.PodPriorityThreshold != nil {
		in, out := &in.PodPriorityThreshold, &out.PodPriorityThreshold
		*out = new(int32)
		**out = **in
	}
	if in.BalanceSimilarNodeGroups != nil {
		in, out := &in.BalanceSimilarNodeGroups, &out.BalanceSimilarNodeGroups
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAutoscalerSpec.
func (in *ClusterAutoscalerSpec) DeepCopy() *ClusterAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAutoscalerStatus) DeepCopyInto(out *ClusterAutoscalerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAutoscalerStatus.
func (in *ClusterAutoscalerStatus) DeepCopy() *ClusterAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GPULimit) DeepCopyInto(out *GPULimit) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GPULimit.
func (in *GPULimit) DeepCopy() *GPULimit {
	if in == nil {
		return nil
	}
	out := new(GPULimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceLimits) DeepCopyInto(out *ResourceLimits) {
	*out = *in
	if in.MaxNodesTotal != nil {
		in, out := &in.MaxNodesTotal, &out.MaxNodesTotal
		*out = new(int32)
		**out = **in
	}
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		*out = new(ResourceRange)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(ResourceRange)
		**out = **in
	}
	if in.GPUS != nil {
		in, out := &in.GPUS, &out.GPUS
		*out = make([]GPULimit, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceLimits.
func (in *ResourceLimits) DeepCopy() *ResourceLimits {
	if in == nil {
		return nil
	}
	out := new(ResourceLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRange) DeepCopyInto(out *ResourceRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRange.
func (in *ResourceRange) DeepCopy() *ResourceRange {
	if in == nil {
		return nil
	}
	out := new(ResourceRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleDownConfig) DeepCopyInto(out *ScaleDownConfig) {
	*out = *in
	if in.DelayAfterAdd != nil {
		in, out := &in.DelayAfterAdd, &out.DelayAfterAdd
		*out = new(string)
		**out = **in
	}
	if in.DelayAfterDelete != nil {
		in, out := &in.DelayAfterDelete, &out.DelayAfterDelete
		*out = new(string)
		**out = **in
	}
	if in.DelayAfterFailure != nil {
		in, out := &in.DelayAfterFailure, &out.DelayAfterFailure
		*out = new(string)
		**out = **in
	}
	if in.UnneededTime != nil {
		in, out := &in.UnneededTime, &out.UnneededTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleDownConfig.
func (in *ScaleDownConfig) DeepCopy() *ScaleDownConfig {
	if in == nil {
		return nil
	}
	out := new(ScaleDownConfig)
	in.DeepCopyInto(out)
	return out
}
