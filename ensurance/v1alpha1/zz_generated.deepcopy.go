//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	analysisv1alpha1 "github.com/gocrane/api/analysis/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvoidanceAction) DeepCopyInto(out *AvoidanceAction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvoidanceAction.
func (in *AvoidanceAction) DeepCopy() *AvoidanceAction {
	if in == nil {
		return nil
	}
	out := new(AvoidanceAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AvoidanceAction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvoidanceActionList) DeepCopyInto(out *AvoidanceActionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AvoidanceAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvoidanceActionList.
func (in *AvoidanceActionList) DeepCopy() *AvoidanceActionList {
	if in == nil {
		return nil
	}
	out := new(AvoidanceActionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AvoidanceActionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvoidanceActionSpec) DeepCopyInto(out *AvoidanceActionSpec) {
	*out = *in
	if in.Throttle != nil {
		in, out := &in.Throttle, &out.Throttle
		*out = new(ThrottleAction)
		**out = **in
	}
	if in.Eviction != nil {
		in, out := &in.Eviction, &out.Eviction
		*out = new(EvictionAction)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvoidanceActionSpec.
func (in *AvoidanceActionSpec) DeepCopy() *AvoidanceActionSpec {
	if in == nil {
		return nil
	}
	out := new(AvoidanceActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvoidanceActionStatus) DeepCopyInto(out *AvoidanceActionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvoidanceActionStatus.
func (in *AvoidanceActionStatus) DeepCopy() *AvoidanceActionStatus {
	if in == nil {
		return nil
	}
	out := new(AvoidanceActionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvoidanceStrategy) DeepCopyInto(out *AvoidanceStrategy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvoidanceStrategy.
func (in *AvoidanceStrategy) DeepCopy() *AvoidanceStrategy {
	if in == nil {
		return nil
	}
	out := new(AvoidanceStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUBurst) DeepCopyInto(out *CPUBurst) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUBurst.
func (in *CPUBurst) DeepCopy() *CPUBurst {
	if in == nil {
		return nil
	}
	out := new(CPUBurst)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUQOS) DeepCopyInto(out *CPUQOS) {
	*out = *in
	if in.CPUPriority != nil {
		in, out := &in.CPUPriority, &out.CPUPriority
		*out = new(int32)
		**out = **in
	}
	if in.ContainerPriority != nil {
		in, out := &in.ContainerPriority, &out.ContainerPriority
		*out = make(map[string]*int32, len(*in))
		for key, val := range *in {
			var outVal *int32
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(int32)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	out.CPUBurst = in.CPUBurst
	out.HtIsolation = in.HtIsolation
	out.CPUSet = in.CPUSet
	in.RDT.DeepCopyInto(&out.RDT)
	in.ContainerRDT.DeepCopyInto(&out.ContainerRDT)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUQOS.
func (in *CPUQOS) DeepCopy() *CPUQOS {
	if in == nil {
		return nil
	}
	out := new(CPUQOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUSet) DeepCopyInto(out *CPUSet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUSet.
func (in *CPUSet) DeepCopy() *CPUSet {
	if in == nil {
		return nil
	}
	out := new(CPUSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUThrottle) DeepCopyInto(out *CPUThrottle) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUThrottle.
func (in *CPUThrottle) DeepCopy() *CPUThrottle {
	if in == nil {
		return nil
	}
	out := new(CPUThrottle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRDT) DeepCopyInto(out *ContainerRDT) {
	*out = *in
	if in.L3 != nil {
		in, out := &in.L3, &out.L3
		*out = make(map[string]RDTValue, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(RDTValue, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.MB != nil {
		in, out := &in.MB, &out.MB
		*out = make(map[string]RDTValue, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(RDTValue, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRDT.
func (in *ContainerRDT) DeepCopy() *ContainerRDT {
	if in == nil {
		return nil
	}
	out := new(ContainerRDT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskIOLimit) DeepCopyInto(out *DiskIOLimit) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskIOLimit.
func (in *DiskIOLimit) DeepCopy() *DiskIOLimit {
	if in == nil {
		return nil
	}
	out := new(DiskIOLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskIOQOS) DeepCopyInto(out *DiskIOQOS) {
	*out = *in
	out.DiskIOWeight = in.DiskIOWeight
	out.DiskIOLimit = in.DiskIOLimit
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskIOQOS.
func (in *DiskIOQOS) DeepCopy() *DiskIOQOS {
	if in == nil {
		return nil
	}
	out := new(DiskIOQOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskIOWeight) DeepCopyInto(out *DiskIOWeight) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskIOWeight.
func (in *DiskIOWeight) DeepCopy() *DiskIOWeight {
	if in == nil {
		return nil
	}
	out := new(DiskIOWeight)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvictionAction) DeepCopyInto(out *EvictionAction) {
	*out = *in
	if in.TerminationGracePeriodSeconds != nil {
		in, out := &in.TerminationGracePeriodSeconds, &out.TerminationGracePeriodSeconds
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvictionAction.
func (in *EvictionAction) DeepCopy() *EvictionAction {
	if in == nil {
		return nil
	}
	out := new(EvictionAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HtIsolation) DeepCopyInto(out *HtIsolation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HtIsolation.
func (in *HtIsolation) DeepCopy() *HtIsolation {
	if in == nil {
		return nil
	}
	out := new(HtIsolation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LowestPriorityCpuAvoidance) DeepCopyInto(out *LowestPriorityCpuAvoidance) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LowestPriorityCpuAvoidance.
func (in *LowestPriorityCpuAvoidance) DeepCopy() *LowestPriorityCpuAvoidance {
	if in == nil {
		return nil
	}
	out := new(LowestPriorityCpuAvoidance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LowestPriorityNodeCpuLimit) DeepCopyInto(out *LowestPriorityNodeCpuLimit) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LowestPriorityNodeCpuLimit.
func (in *LowestPriorityNodeCpuLimit) DeepCopy() *LowestPriorityNodeCpuLimit {
	if in == nil {
		return nil
	}
	out := new(LowestPriorityNodeCpuLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemAsyncReclaim) DeepCopyInto(out *MemAsyncReclaim) {
	*out = *in
	if in.AsyncRatio != nil {
		in, out := &in.AsyncRatio, &out.AsyncRatio
		*out = new(int64)
		**out = **in
	}
	if in.AsyncDistanceFactor != nil {
		in, out := &in.AsyncDistanceFactor, &out.AsyncDistanceFactor
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemAsyncReclaim.
func (in *MemAsyncReclaim) DeepCopy() *MemAsyncReclaim {
	if in == nil {
		return nil
	}
	out := new(MemAsyncReclaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemWatermark) DeepCopyInto(out *MemWatermark) {
	*out = *in
	if in.WatermarkRatio != nil {
		in, out := &in.WatermarkRatio, &out.WatermarkRatio
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemWatermark.
func (in *MemWatermark) DeepCopy() *MemWatermark {
	if in == nil {
		return nil
	}
	out := new(MemWatermark)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryQOS) DeepCopyInto(out *MemoryQOS) {
	*out = *in
	in.MemAsyncReclaim.DeepCopyInto(&out.MemAsyncReclaim)
	in.MemWatermark.DeepCopyInto(&out.MemWatermark)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryQOS.
func (in *MemoryQOS) DeepCopy() *MemoryQOS {
	if in == nil {
		return nil
	}
	out := new(MemoryQOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryThrottle) DeepCopyInto(out *MemoryThrottle) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryThrottle.
func (in *MemoryThrottle) DeepCopy() *MemoryThrottle {
	if in == nil {
		return nil
	}
	out := new(MemoryThrottle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricRule) DeepCopyInto(out *MetricRule) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	out.Value = in.Value.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricRule.
func (in *MetricRule) DeepCopy() *MetricRule {
	if in == nil {
		return nil
	}
	out := new(MetricRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetIOLimits) DeepCopyInto(out *NetIOLimits) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetIOLimits.
func (in *NetIOLimits) DeepCopy() *NetIOLimits {
	if in == nil {
		return nil
	}
	out := new(NetIOLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetIOQOS) DeepCopyInto(out *NetIOQOS) {
	*out = *in
	out.NetIOLimits = in.NetIOLimits
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetIOQOS.
func (in *NetIOQOS) DeepCopy() *NetIOQOS {
	if in == nil {
		return nil
	}
	out := new(NetIOQOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeCpuLimit) DeepCopyInto(out *NodeCpuLimit) {
	*out = *in
	out.LowestPriorityNodeCpuLimit = in.LowestPriorityNodeCpuLimit
	if in.LowestPriorityCoreCpuLimit != nil {
		in, out := &in.LowestPriorityCoreCpuLimit, &out.LowestPriorityCoreCpuLimit
		*out = make([]lowestPriorityCoreCpuLimit, len(*in))
		copy(*out, *in)
	}
	if in.LowestPriorityCoreCpuLimitPeriod != nil {
		in, out := &in.LowestPriorityCoreCpuLimitPeriod, &out.LowestPriorityCoreCpuLimitPeriod
		*out = make([]lowestPriorityCoreCpuLimitPeriod, len(*in))
		copy(*out, *in)
	}
	out.LowestPriorityCpuAvoidance = in.LowestPriorityCpuAvoidance
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeCpuLimit.
func (in *NodeCpuLimit) DeepCopy() *NodeCpuLimit {
	if in == nil {
		return nil
	}
	out := new(NodeCpuLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeLocalGet) DeepCopyInto(out *NodeLocalGet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeLocalGet.
func (in *NodeLocalGet) DeepCopy() *NodeLocalGet {
	if in == nil {
		return nil
	}
	out := new(NodeLocalGet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeQOS) DeepCopyInto(out *NodeQOS) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeQOS.
func (in *NodeQOS) DeepCopy() *NodeQOS {
	if in == nil {
		return nil
	}
	out := new(NodeQOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeQOS) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeQOSList) DeepCopyInto(out *NodeQOSList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeQOS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeQOSList.
func (in *NodeQOSList) DeepCopy() *NodeQOSList {
	if in == nil {
		return nil
	}
	out := new(NodeQOSList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeQOSList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeQOSSpec) DeepCopyInto(out *NodeQOSSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.NodeQualityProbe.DeepCopyInto(&out.NodeQualityProbe)
	if in.WaterLine != nil {
		in, out := &in.WaterLine, &out.WaterLine
		*out = make([]WaterLine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.NodeCpuLimit.DeepCopyInto(&out.NodeCpuLimit)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeQOSSpec.
func (in *NodeQOSSpec) DeepCopy() *NodeQOSSpec {
	if in == nil {
		return nil
	}
	out := new(NodeQOSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeQOSStatus) DeepCopyInto(out *NodeQOSStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeQOSStatus.
func (in *NodeQOSStatus) DeepCopy() *NodeQOSStatus {
	if in == nil {
		return nil
	}
	out := new(NodeQOSStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeQualityProbe) DeepCopyInto(out *NodeQualityProbe) {
	*out = *in
	if in.HTTPGet != nil {
		in, out := &in.HTTPGet, &out.HTTPGet
		*out = new(corev1.HTTPGetAction)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeLocalGet != nil {
		in, out := &in.NodeLocalGet, &out.NodeLocalGet
		*out = new(NodeLocalGet)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeQualityProbe.
func (in *NodeQualityProbe) DeepCopy() *NodeQualityProbe {
	if in == nil {
		return nil
	}
	out := new(NodeQualityProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodQOSEnsurancePolicy) DeepCopyInto(out *PodQOSEnsurancePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodQOSEnsurancePolicy.
func (in *PodQOSEnsurancePolicy) DeepCopy() *PodQOSEnsurancePolicy {
	if in == nil {
		return nil
	}
	out := new(PodQOSEnsurancePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodQOSEnsurancePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodQOSEnsurancePolicyList) DeepCopyInto(out *PodQOSEnsurancePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodQOSEnsurancePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodQOSEnsurancePolicyList.
func (in *PodQOSEnsurancePolicyList) DeepCopy() *PodQOSEnsurancePolicyList {
	if in == nil {
		return nil
	}
	out := new(PodQOSEnsurancePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodQOSEnsurancePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodQOSEnsurancePolicySpec) DeepCopyInto(out *PodQOSEnsurancePolicySpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.PodQualityProbe.DeepCopyInto(&out.PodQualityProbe)
	if in.ObjectiveEnsurances != nil {
		in, out := &in.ObjectiveEnsurances, &out.ObjectiveEnsurances
		*out = make([]QOSEnsurance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodQOSEnsurancePolicySpec.
func (in *PodQOSEnsurancePolicySpec) DeepCopy() *PodQOSEnsurancePolicySpec {
	if in == nil {
		return nil
	}
	out := new(PodQOSEnsurancePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodQOSEnsurancePolicyStatus) DeepCopyInto(out *PodQOSEnsurancePolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodQOSEnsurancePolicyStatus.
func (in *PodQOSEnsurancePolicyStatus) DeepCopy() *PodQOSEnsurancePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PodQOSEnsurancePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodQualityProbe) DeepCopyInto(out *PodQualityProbe) {
	*out = *in
	if in.HTTPGet != nil {
		in, out := &in.HTTPGet, &out.HTTPGet
		*out = new(corev1.HTTPGetAction)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodQualityProbe.
func (in *PodQualityProbe) DeepCopy() *PodQualityProbe {
	if in == nil {
		return nil
	}
	out := new(PodQualityProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QOSEnsurance) DeepCopyInto(out *QOSEnsurance) {
	*out = *in
	in.WaterLine.DeepCopyInto(&out.WaterLine)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QOSEnsurance.
func (in *QOSEnsurance) DeepCopy() *QOSEnsurance {
	if in == nil {
		return nil
	}
	out := new(QOSEnsurance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDT) DeepCopyInto(out *RDT) {
	*out = *in
	if in.L3 != nil {
		in, out := &in.L3, &out.L3
		*out = make(RDTValue, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MB != nil {
		in, out := &in.MB, &out.MB
		*out = make(RDTValue, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDT.
func (in *RDT) DeepCopy() *RDT {
	if in == nil {
		return nil
	}
	out := new(RDT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in RDTValue) DeepCopyInto(out *RDTValue) {
	{
		in := &in
		*out = make(RDTValue, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDTValue.
func (in RDTValue) DeepCopy() RDTValue {
	if in == nil {
		return nil
	}
	out := new(RDTValue)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMutation) DeepCopyInto(out *ResourceMutation) {
	*out = *in
	if in.RequestMutations != nil {
		in, out := &in.RequestMutations, &out.RequestMutations
		*out = make([]ResourcetMutation, len(*in))
		copy(*out, *in)
	}
	if in.LimitMutations != nil {
		in, out := &in.LimitMutations, &out.LimitMutations
		*out = make([]ResourcetMutation, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMutation.
func (in *ResourceMutation) DeepCopy() *ResourceMutation {
	if in == nil {
		return nil
	}
	out := new(ResourceMutation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePriority) DeepCopyInto(out *ResourcePriority) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePriority.
func (in *ResourcePriority) DeepCopy() *ResourcePriority {
	if in == nil {
		return nil
	}
	out := new(ResourcePriority)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQOS) DeepCopyInto(out *ResourceQOS) {
	*out = *in
	if in.CPUQOS != nil {
		in, out := &in.CPUQOS, &out.CPUQOS
		*out = new(CPUQOS)
		(*in).DeepCopyInto(*out)
	}
	if in.MemoryQOS != nil {
		in, out := &in.MemoryQOS, &out.MemoryQOS
		*out = new(MemoryQOS)
		(*in).DeepCopyInto(*out)
	}
	if in.NetIOQOS != nil {
		in, out := &in.NetIOQOS, &out.NetIOQOS
		*out = new(NetIOQOS)
		**out = **in
	}
	if in.DiskIOQOS != nil {
		in, out := &in.DiskIOQOS, &out.DiskIOQOS
		*out = new(DiskIOQOS)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQOS.
func (in *ResourceQOS) DeepCopy() *ResourceQOS {
	if in == nil {
		return nil
	}
	out := new(ResourceQOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcetMutation) DeepCopyInto(out *ResourcetMutation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcetMutation.
func (in *ResourcetMutation) DeepCopy() *ResourcetMutation {
	if in == nil {
		return nil
	}
	out := new(ResourcetMutation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicePolicy) DeepCopyInto(out *ServicePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePolicy.
func (in *ServicePolicy) DeepCopy() *ServicePolicy {
	if in == nil {
		return nil
	}
	out := new(ServicePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicePolicyList) DeepCopyInto(out *ServicePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServicePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePolicyList.
func (in *ServicePolicyList) DeepCopy() *ServicePolicyList {
	if in == nil {
		return nil
	}
	out := new(ServicePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicePolicySpec) DeepCopyInto(out *ServicePolicySpec) {
	*out = *in
	out.ResourcePriority = in.ResourcePriority
	out.AvoidanceStrategy = in.AvoidanceStrategy
	in.ResourceMutation.DeepCopyInto(&out.ResourceMutation)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePolicySpec.
func (in *ServicePolicySpec) DeepCopy() *ServicePolicySpec {
	if in == nil {
		return nil
	}
	out := new(ServicePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicePolicyStatus) DeepCopyInto(out *ServicePolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePolicyStatus.
func (in *ServicePolicyStatus) DeepCopy() *ServicePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ServicePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceQOS) DeepCopyInto(out *ServiceQOS) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceQOS.
func (in *ServiceQOS) DeepCopy() *ServiceQOS {
	if in == nil {
		return nil
	}
	out := new(ServiceQOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceQOS) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceQOSEnsurancePolicyStatus) DeepCopyInto(out *ServiceQOSEnsurancePolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceQOSEnsurancePolicyStatus.
func (in *ServiceQOSEnsurancePolicyStatus) DeepCopy() *ServiceQOSEnsurancePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceQOSEnsurancePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceQOSList) DeepCopyInto(out *ServiceQOSList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceQOS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceQOSList.
func (in *ServiceQOSList) DeepCopy() *ServiceQOSList {
	if in == nil {
		return nil
	}
	out := new(ServiceQOSList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceQOSList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceQOSSpec) DeepCopyInto(out *ServiceQOSSpec) {
	*out = *in
	if in.QOSClassSelectors != nil {
		in, out := &in.QOSClassSelectors, &out.QOSClassSelectors
		*out = make([]corev1.PodQOSClass, len(*in))
		copy(*out, *in)
	}
	if in.PrioritySelectors != nil {
		in, out := &in.PrioritySelectors, &out.PrioritySelectors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceSelectors != nil {
		in, out := &in.ResourceSelectors, &out.ResourceSelectors
		*out = make([]analysisv1alpha1.ResourceSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ResourceQOS.DeepCopyInto(&out.ResourceQOS)
	in.PodQualityProbe.DeepCopyInto(&out.PodQualityProbe)
	if in.WaterLine != nil {
		in, out := &in.WaterLine, &out.WaterLine
		*out = make([]WaterLine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowedActions != nil {
		in, out := &in.AllowedActions, &out.AllowedActions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceQOSSpec.
func (in *ServiceQOSSpec) DeepCopy() *ServiceQOSSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceQOSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThrottleAction) DeepCopyInto(out *ThrottleAction) {
	*out = *in
	out.CPUThrottle = in.CPUThrottle
	out.MemoryThrottle = in.MemoryThrottle
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThrottleAction.
func (in *ThrottleAction) DeepCopy() *ThrottleAction {
	if in == nil {
		return nil
	}
	out := new(ThrottleAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WaterLine) DeepCopyInto(out *WaterLine) {
	*out = *in
	if in.MetricRule != nil {
		in, out := &in.MetricRule, &out.MetricRule
		*out = new(MetricRule)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WaterLine.
func (in *WaterLine) DeepCopy() *WaterLine {
	if in == nil {
		return nil
	}
	out := new(WaterLine)
	in.DeepCopyInto(out)
	return out
}
