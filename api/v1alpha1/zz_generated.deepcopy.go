// +build !ignore_autogenerated

/*


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
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomSubgraphImage) DeepCopyInto(out *CustomSubgraphImage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomSubgraphImage.
func (in *CustomSubgraphImage) DeepCopy() *CustomSubgraphImage {
	if in == nil {
		return nil
	}
	out := new(CustomSubgraphImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyRef) DeepCopyInto(out *SecretKeyRef) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyRef.
func (in *SecretKeyRef) DeepCopy() *SecretKeyRef {
	if in == nil {
		return nil
	}
	out := new(SecretKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringOrSecretParameter) DeepCopyInto(out *StringOrSecretParameter) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(SecretKeyRef)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringOrSecretParameter.
func (in *StringOrSecretParameter) DeepCopy() *StringOrSecretParameter {
	if in == nil {
		return nil
	}
	out := new(StringOrSecretParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinDataSource) DeepCopyInto(out *XJoinDataSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinDataSource.
func (in *XJoinDataSource) DeepCopy() *XJoinDataSource {
	if in == nil {
		return nil
	}
	out := new(XJoinDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XJoinDataSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinDataSourceList) DeepCopyInto(out *XJoinDataSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]XJoinDataSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinDataSourceList.
func (in *XJoinDataSourceList) DeepCopy() *XJoinDataSourceList {
	if in == nil {
		return nil
	}
	out := new(XJoinDataSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XJoinDataSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinDataSourcePipeline) DeepCopyInto(out *XJoinDataSourcePipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinDataSourcePipeline.
func (in *XJoinDataSourcePipeline) DeepCopy() *XJoinDataSourcePipeline {
	if in == nil {
		return nil
	}
	out := new(XJoinDataSourcePipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XJoinDataSourcePipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinDataSourcePipelineList) DeepCopyInto(out *XJoinDataSourcePipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]XJoinDataSourcePipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinDataSourcePipelineList.
func (in *XJoinDataSourcePipelineList) DeepCopy() *XJoinDataSourcePipelineList {
	if in == nil {
		return nil
	}
	out := new(XJoinDataSourcePipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XJoinDataSourcePipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinDataSourcePipelineSpec) DeepCopyInto(out *XJoinDataSourcePipelineSpec) {
	*out = *in
	if in.DatabaseHostname != nil {
		in, out := &in.DatabaseHostname, &out.DatabaseHostname
		*out = new(StringOrSecretParameter)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabasePort != nil {
		in, out := &in.DatabasePort, &out.DatabasePort
		*out = new(StringOrSecretParameter)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabaseUsername != nil {
		in, out := &in.DatabaseUsername, &out.DatabaseUsername
		*out = new(StringOrSecretParameter)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabasePassword != nil {
		in, out := &in.DatabasePassword, &out.DatabasePassword
		*out = new(StringOrSecretParameter)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(StringOrSecretParameter)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabaseTable != nil {
		in, out := &in.DatabaseTable, &out.DatabaseTable
		*out = new(StringOrSecretParameter)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinDataSourcePipelineSpec.
func (in *XJoinDataSourcePipelineSpec) DeepCopy() *XJoinDataSourcePipelineSpec {
	if in == nil {
		return nil
	}
	out := new(XJoinDataSourcePipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinDataSourcePipelineStatus) DeepCopyInto(out *XJoinDataSourcePipelineStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinDataSourcePipelineStatus.
func (in *XJoinDataSourcePipelineStatus) DeepCopy() *XJoinDataSourcePipelineStatus {
	if in == nil {
		return nil
	}
	out := new(XJoinDataSourcePipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinDataSourceSpec) DeepCopyInto(out *XJoinDataSourceSpec) {
	*out = *in
	if in.DatabaseHostname != nil {
		in, out := &in.DatabaseHostname, &out.DatabaseHostname
		*out = new(StringOrSecretParameter)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabasePort != nil {
		in, out := &in.DatabasePort, &out.DatabasePort
		*out = new(StringOrSecretParameter)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabaseUsername != nil {
		in, out := &in.DatabaseUsername, &out.DatabaseUsername
		*out = new(StringOrSecretParameter)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabasePassword != nil {
		in, out := &in.DatabasePassword, &out.DatabasePassword
		*out = new(StringOrSecretParameter)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(StringOrSecretParameter)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabaseTable != nil {
		in, out := &in.DatabaseTable, &out.DatabaseTable
		*out = new(StringOrSecretParameter)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinDataSourceSpec.
func (in *XJoinDataSourceSpec) DeepCopy() *XJoinDataSourceSpec {
	if in == nil {
		return nil
	}
	out := new(XJoinDataSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinDataSourceStatus) DeepCopyInto(out *XJoinDataSourceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinDataSourceStatus.
func (in *XJoinDataSourceStatus) DeepCopy() *XJoinDataSourceStatus {
	if in == nil {
		return nil
	}
	out := new(XJoinDataSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinIndex) DeepCopyInto(out *XJoinIndex) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinIndex.
func (in *XJoinIndex) DeepCopy() *XJoinIndex {
	if in == nil {
		return nil
	}
	out := new(XJoinIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XJoinIndex) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinIndexList) DeepCopyInto(out *XJoinIndexList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]XJoinIndex, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinIndexList.
func (in *XJoinIndexList) DeepCopy() *XJoinIndexList {
	if in == nil {
		return nil
	}
	out := new(XJoinIndexList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XJoinIndexList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinIndexPipeline) DeepCopyInto(out *XJoinIndexPipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinIndexPipeline.
func (in *XJoinIndexPipeline) DeepCopy() *XJoinIndexPipeline {
	if in == nil {
		return nil
	}
	out := new(XJoinIndexPipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XJoinIndexPipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinIndexPipelineList) DeepCopyInto(out *XJoinIndexPipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]XJoinIndexPipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinIndexPipelineList.
func (in *XJoinIndexPipelineList) DeepCopy() *XJoinIndexPipelineList {
	if in == nil {
		return nil
	}
	out := new(XJoinIndexPipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XJoinIndexPipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinIndexPipelineSpec) DeepCopyInto(out *XJoinIndexPipelineSpec) {
	*out = *in
	if in.CustomSubgraphImages != nil {
		in, out := &in.CustomSubgraphImages, &out.CustomSubgraphImages
		*out = make([]CustomSubgraphImage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinIndexPipelineSpec.
func (in *XJoinIndexPipelineSpec) DeepCopy() *XJoinIndexPipelineSpec {
	if in == nil {
		return nil
	}
	out := new(XJoinIndexPipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinIndexPipelineStatus) DeepCopyInto(out *XJoinIndexPipelineStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinIndexPipelineStatus.
func (in *XJoinIndexPipelineStatus) DeepCopy() *XJoinIndexPipelineStatus {
	if in == nil {
		return nil
	}
	out := new(XJoinIndexPipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinIndexSpec) DeepCopyInto(out *XJoinIndexSpec) {
	*out = *in
	if in.CustomSubgraphImages != nil {
		in, out := &in.CustomSubgraphImages, &out.CustomSubgraphImages
		*out = make([]CustomSubgraphImage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinIndexSpec.
func (in *XJoinIndexSpec) DeepCopy() *XJoinIndexSpec {
	if in == nil {
		return nil
	}
	out := new(XJoinIndexSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinIndexStatus) DeepCopyInto(out *XJoinIndexStatus) {
	*out = *in
	if in.DataSources != nil {
		in, out := &in.DataSources, &out.DataSources
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinIndexStatus.
func (in *XJoinIndexStatus) DeepCopy() *XJoinIndexStatus {
	if in == nil {
		return nil
	}
	out := new(XJoinIndexStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinIndexValidator) DeepCopyInto(out *XJoinIndexValidator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinIndexValidator.
func (in *XJoinIndexValidator) DeepCopy() *XJoinIndexValidator {
	if in == nil {
		return nil
	}
	out := new(XJoinIndexValidator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XJoinIndexValidator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinIndexValidatorList) DeepCopyInto(out *XJoinIndexValidatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]XJoinIndexValidator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinIndexValidatorList.
func (in *XJoinIndexValidatorList) DeepCopy() *XJoinIndexValidatorList {
	if in == nil {
		return nil
	}
	out := new(XJoinIndexValidatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XJoinIndexValidatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinIndexValidatorSpec) DeepCopyInto(out *XJoinIndexValidatorSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinIndexValidatorSpec.
func (in *XJoinIndexValidatorSpec) DeepCopy() *XJoinIndexValidatorSpec {
	if in == nil {
		return nil
	}
	out := new(XJoinIndexValidatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinIndexValidatorStatus) DeepCopyInto(out *XJoinIndexValidatorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinIndexValidatorStatus.
func (in *XJoinIndexValidatorStatus) DeepCopy() *XJoinIndexValidatorStatus {
	if in == nil {
		return nil
	}
	out := new(XJoinIndexValidatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinPipeline) DeepCopyInto(out *XJoinPipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinPipeline.
func (in *XJoinPipeline) DeepCopy() *XJoinPipeline {
	if in == nil {
		return nil
	}
	out := new(XJoinPipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XJoinPipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinPipelineList) DeepCopyInto(out *XJoinPipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]XJoinPipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinPipelineList.
func (in *XJoinPipelineList) DeepCopy() *XJoinPipelineList {
	if in == nil {
		return nil
	}
	out := new(XJoinPipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XJoinPipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinPipelineSpec) DeepCopyInto(out *XJoinPipelineSpec) {
	*out = *in
	if in.ResourceNamePrefix != nil {
		in, out := &in.ResourceNamePrefix, &out.ResourceNamePrefix
		*out = new(string)
		**out = **in
	}
	if in.KafkaCluster != nil {
		in, out := &in.KafkaCluster, &out.KafkaCluster
		*out = new(string)
		**out = **in
	}
	if in.KafkaClusterNamespace != nil {
		in, out := &in.KafkaClusterNamespace, &out.KafkaClusterNamespace
		*out = new(string)
		**out = **in
	}
	if in.ConnectCluster != nil {
		in, out := &in.ConnectCluster, &out.ConnectCluster
		*out = new(string)
		**out = **in
	}
	if in.ConnectClusterNamespace != nil {
		in, out := &in.ConnectClusterNamespace, &out.ConnectClusterNamespace
		*out = new(string)
		**out = **in
	}
	if in.HBIDBSecretName != nil {
		in, out := &in.HBIDBSecretName, &out.HBIDBSecretName
		*out = new(string)
		**out = **in
	}
	if in.ElasticSearchSecretName != nil {
		in, out := &in.ElasticSearchSecretName, &out.ElasticSearchSecretName
		*out = new(string)
		**out = **in
	}
	if in.ElasticSearchNamespace != nil {
		in, out := &in.ElasticSearchNamespace, &out.ElasticSearchNamespace
		*out = new(string)
		**out = **in
	}
	if in.ElasticSearchIndexTemplate != nil {
		in, out := &in.ElasticSearchIndexTemplate, &out.ElasticSearchIndexTemplate
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinPipelineSpec.
func (in *XJoinPipelineSpec) DeepCopy() *XJoinPipelineSpec {
	if in == nil {
		return nil
	}
	out := new(XJoinPipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XJoinPipelineStatus) DeepCopyInto(out *XJoinPipelineStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XJoinPipelineStatus.
func (in *XJoinPipelineStatus) DeepCopy() *XJoinPipelineStatus {
	if in == nil {
		return nil
	}
	out := new(XJoinPipelineStatus)
	in.DeepCopyInto(out)
	return out
}
