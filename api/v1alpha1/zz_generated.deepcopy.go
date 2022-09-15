//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Actions) DeepCopyInto(out *Actions) {
	*out = *in
	in.BeforeEnable.DeepCopyInto(&out.BeforeEnable)
	in.AfterCompletion.DeepCopyInto(&out.AfterCompletion)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Actions.
func (in *Actions) DeepCopy() *Actions {
	if in == nil {
		return nil
	}
	out := new(Actions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AfterCompletion) DeepCopyInto(out *AfterCompletion) {
	*out = *in
	if in.AddClusterLabels != nil {
		in, out := &in.AddClusterLabels, &out.AddClusterLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeleteClusterLabels != nil {
		in, out := &in.DeleteClusterLabels, &out.DeleteClusterLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeleteObjects != nil {
		in, out := &in.DeleteObjects, &out.DeleteObjects
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AfterCompletion.
func (in *AfterCompletion) DeepCopy() *AfterCompletion {
	if in == nil {
		return nil
	}
	out := new(AfterCompletion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatus) DeepCopyInto(out *BackupStatus) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatus.
func (in *BackupStatus) DeepCopy() *BackupStatus {
	if in == nil {
		return nil
	}
	out := new(BackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeforeEnable) DeepCopyInto(out *BeforeEnable) {
	*out = *in
	if in.AddClusterLabels != nil {
		in, out := &in.AddClusterLabels, &out.AddClusterLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeleteClusterLabels != nil {
		in, out := &in.DeleteClusterLabels, &out.DeleteClusterLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeforeEnable.
func (in *BeforeEnable) DeepCopy() *BeforeEnable {
	if in == nil {
		return nil
	}
	out := new(BeforeEnable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockingCR) DeepCopyInto(out *BlockingCR) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockingCR.
func (in *BlockingCR) DeepCopy() *BlockingCR {
	if in == nil {
		return nil
	}
	out := new(BlockingCR)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGroupUpgrade) DeepCopyInto(out *ClusterGroupUpgrade) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGroupUpgrade.
func (in *ClusterGroupUpgrade) DeepCopy() *ClusterGroupUpgrade {
	if in == nil {
		return nil
	}
	out := new(ClusterGroupUpgrade)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterGroupUpgrade) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGroupUpgradeList) DeepCopyInto(out *ClusterGroupUpgradeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterGroupUpgrade, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGroupUpgradeList.
func (in *ClusterGroupUpgradeList) DeepCopy() *ClusterGroupUpgradeList {
	if in == nil {
		return nil
	}
	out := new(ClusterGroupUpgradeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterGroupUpgradeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGroupUpgradeSpec) DeepCopyInto(out *ClusterGroupUpgradeSpec) {
	*out = *in
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterSelector != nil {
		in, out := &in.ClusterSelector, &out.ClusterSelector
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterLabelSelectors != nil {
		in, out := &in.ClusterLabelSelectors, &out.ClusterLabelSelectors
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RemediationStrategy != nil {
		in, out := &in.RemediationStrategy, &out.RemediationStrategy
		*out = new(RemediationStrategySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ManagedPolicies != nil {
		in, out := &in.ManagedPolicies, &out.ManagedPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BlockingCRs != nil {
		in, out := &in.BlockingCRs, &out.BlockingCRs
		*out = make([]BlockingCR, len(*in))
		copy(*out, *in)
	}
	in.Actions.DeepCopyInto(&out.Actions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGroupUpgradeSpec.
func (in *ClusterGroupUpgradeSpec) DeepCopy() *ClusterGroupUpgradeSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterGroupUpgradeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGroupUpgradeStatus) DeepCopyInto(out *ClusterGroupUpgradeStatus) {
	*out = *in
	if in.PlacementBindings != nil {
		in, out := &in.PlacementBindings, &out.PlacementBindings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PlacementRules != nil {
		in, out := &in.PlacementRules, &out.PlacementRules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CopiedPolicies != nil {
		in, out := &in.CopiedPolicies, &out.CopiedPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RemediationPlan != nil {
		in, out := &in.RemediationPlan, &out.RemediationPlan
		*out = make([][]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
		}
	}
	if in.ManagedPoliciesNs != nil {
		in, out := &in.ManagedPoliciesNs, &out.ManagedPoliciesNs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SafeResourceNames != nil {
		in, out := &in.SafeResourceNames, &out.SafeResourceNames
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ManagedPoliciesForUpgrade != nil {
		in, out := &in.ManagedPoliciesForUpgrade, &out.ManagedPoliciesForUpgrade
		*out = make([]ManagedPolicyForUpgrade, len(*in))
		copy(*out, *in)
	}
	if in.ManagedPoliciesCompliantBeforeUpgrade != nil {
		in, out := &in.ManagedPoliciesCompliantBeforeUpgrade, &out.ManagedPoliciesCompliantBeforeUpgrade
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ManagedPoliciesContent != nil {
		in, out := &in.ManagedPoliciesContent, &out.ManagedPoliciesContent
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]ClusterState, len(*in))
		copy(*out, *in)
	}
	in.Status.DeepCopyInto(&out.Status)
	if in.Precaching != nil {
		in, out := &in.Precaching, &out.Precaching
		*out = new(PrecachingStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(BackupStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGroupUpgradeStatus.
func (in *ClusterGroupUpgradeStatus) DeepCopy() *ClusterGroupUpgradeStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterGroupUpgradeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRemediationProgress) DeepCopyInto(out *ClusterRemediationProgress) {
	*out = *in
	if in.PolicyIndex != nil {
		in, out := &in.PolicyIndex, &out.PolicyIndex
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRemediationProgress.
func (in *ClusterRemediationProgress) DeepCopy() *ClusterRemediationProgress {
	if in == nil {
		return nil
	}
	out := new(ClusterRemediationProgress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterState) DeepCopyInto(out *ClusterState) {
	*out = *in
	out.CurrentPolicy = in.CurrentPolicy
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterState.
func (in *ClusterState) DeepCopy() *ClusterState {
	if in == nil {
		return nil
	}
	out := new(ClusterState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPolicyForUpgrade) DeepCopyInto(out *ManagedPolicyForUpgrade) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPolicyForUpgrade.
func (in *ManagedPolicyForUpgrade) DeepCopy() *ManagedPolicyForUpgrade {
	if in == nil {
		return nil
	}
	out := new(ManagedPolicyForUpgrade)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorUpgradeSpec) DeepCopyInto(out *OperatorUpgradeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorUpgradeSpec.
func (in *OperatorUpgradeSpec) DeepCopy() *OperatorUpgradeSpec {
	if in == nil {
		return nil
	}
	out := new(OperatorUpgradeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyContent) DeepCopyInto(out *PolicyContent) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyContent.
func (in *PolicyContent) DeepCopy() *PolicyContent {
	if in == nil {
		return nil
	}
	out := new(PolicyContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrecachingSpec) DeepCopyInto(out *PrecachingSpec) {
	*out = *in
	if in.OperatorsIndexes != nil {
		in, out := &in.OperatorsIndexes, &out.OperatorsIndexes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OperatorsPackagesAndChannels != nil {
		in, out := &in.OperatorsPackagesAndChannels, &out.OperatorsPackagesAndChannels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrecachingSpec.
func (in *PrecachingSpec) DeepCopy() *PrecachingSpec {
	if in == nil {
		return nil
	}
	out := new(PrecachingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrecachingStatus) DeepCopyInto(out *PrecachingStatus) {
	*out = *in
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(PrecachingSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrecachingStatus.
func (in *PrecachingStatus) DeepCopy() *PrecachingStatus {
	if in == nil {
		return nil
	}
	out := new(PrecachingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemediationStrategySpec) DeepCopyInto(out *RemediationStrategySpec) {
	*out = *in
	if in.Canaries != nil {
		in, out := &in.Canaries, &out.Canaries
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemediationStrategySpec.
func (in *RemediationStrategySpec) DeepCopy() *RemediationStrategySpec {
	if in == nil {
		return nil
	}
	out := new(RemediationStrategySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeStatus) DeepCopyInto(out *UpgradeStatus) {
	*out = *in
	in.StartedAt.DeepCopyInto(&out.StartedAt)
	in.CompletedAt.DeepCopyInto(&out.CompletedAt)
	in.CurrentBatchStartedAt.DeepCopyInto(&out.CurrentBatchStartedAt)
	if in.CurrentBatchRemediationProgress != nil {
		in, out := &in.CurrentBatchRemediationProgress, &out.CurrentBatchRemediationProgress
		*out = make(map[string]*ClusterRemediationProgress, len(*in))
		for key, val := range *in {
			var outVal *ClusterRemediationProgress
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ClusterRemediationProgress)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeStatus.
func (in *UpgradeStatus) DeepCopy() *UpgradeStatus {
	if in == nil {
		return nil
	}
	out := new(UpgradeStatus)
	in.DeepCopyInto(out)
	return out
}
