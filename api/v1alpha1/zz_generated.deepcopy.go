//go:build !ignore_autogenerated

/*
2024 NVIDIA CORPORATION & AFFILIATES
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
func (in *ConfigurationTemplateSpec) DeepCopyInto(out *ConfigurationTemplateSpec) {
	*out = *in
	if in.PciPerformanceOptimized != nil {
		in, out := &in.PciPerformanceOptimized, &out.PciPerformanceOptimized
		*out = new(PciPerformanceOptimizedSpec)
		**out = **in
	}
	if in.RoceOptimized != nil {
		in, out := &in.RoceOptimized, &out.RoceOptimized
		*out = new(RoceOptimizedSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.GpuDirectOptimized != nil {
		in, out := &in.GpuDirectOptimized, &out.GpuDirectOptimized
		*out = new(GpuDirectOptimizedSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationTemplateSpec.
func (in *ConfigurationTemplateSpec) DeepCopy() *ConfigurationTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GpuDirectOptimizedSpec) DeepCopyInto(out *GpuDirectOptimizedSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GpuDirectOptimizedSpec.
func (in *GpuDirectOptimizedSpec) DeepCopy() *GpuDirectOptimizedSpec {
	if in == nil {
		return nil
	}
	out := new(GpuDirectOptimizedSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicConfigurationTemplate) DeepCopyInto(out *NicConfigurationTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicConfigurationTemplate.
func (in *NicConfigurationTemplate) DeepCopy() *NicConfigurationTemplate {
	if in == nil {
		return nil
	}
	out := new(NicConfigurationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NicConfigurationTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicConfigurationTemplateList) DeepCopyInto(out *NicConfigurationTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NicConfigurationTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicConfigurationTemplateList.
func (in *NicConfigurationTemplateList) DeepCopy() *NicConfigurationTemplateList {
	if in == nil {
		return nil
	}
	out := new(NicConfigurationTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NicConfigurationTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicConfigurationTemplateSpec) DeepCopyInto(out *NicConfigurationTemplateSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NicSelector != nil {
		in, out := &in.NicSelector, &out.NicSelector
		*out = new(NicSelectorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(ConfigurationTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicConfigurationTemplateSpec.
func (in *NicConfigurationTemplateSpec) DeepCopy() *NicConfigurationTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(NicConfigurationTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicConfigurationTemplateStatus) DeepCopyInto(out *NicConfigurationTemplateStatus) {
	*out = *in
	if in.NicDevices != nil {
		in, out := &in.NicDevices, &out.NicDevices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicConfigurationTemplateStatus.
func (in *NicConfigurationTemplateStatus) DeepCopy() *NicConfigurationTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(NicConfigurationTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicDevice) DeepCopyInto(out *NicDevice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicDevice.
func (in *NicDevice) DeepCopy() *NicDevice {
	if in == nil {
		return nil
	}
	out := new(NicDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NicDevice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicDeviceConfigurationSpec) DeepCopyInto(out *NicDeviceConfigurationSpec) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(ConfigurationTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicDeviceConfigurationSpec.
func (in *NicDeviceConfigurationSpec) DeepCopy() *NicDeviceConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(NicDeviceConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicDeviceList) DeepCopyInto(out *NicDeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NicDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicDeviceList.
func (in *NicDeviceList) DeepCopy() *NicDeviceList {
	if in == nil {
		return nil
	}
	out := new(NicDeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NicDeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicDevicePortSpec) DeepCopyInto(out *NicDevicePortSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicDevicePortSpec.
func (in *NicDevicePortSpec) DeepCopy() *NicDevicePortSpec {
	if in == nil {
		return nil
	}
	out := new(NicDevicePortSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicDeviceSpec) DeepCopyInto(out *NicDeviceSpec) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(NicDeviceConfigurationSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicDeviceSpec.
func (in *NicDeviceSpec) DeepCopy() *NicDeviceSpec {
	if in == nil {
		return nil
	}
	out := new(NicDeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicDeviceStatus) DeepCopyInto(out *NicDeviceStatus) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]NicDevicePortSpec, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicDeviceStatus.
func (in *NicDeviceStatus) DeepCopy() *NicDeviceStatus {
	if in == nil {
		return nil
	}
	out := new(NicDeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicSelectorSpec) DeepCopyInto(out *NicSelectorSpec) {
	*out = *in
	if in.PciAddresses != nil {
		in, out := &in.PciAddresses, &out.PciAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SerialNumbers != nil {
		in, out := &in.SerialNumbers, &out.SerialNumbers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicSelectorSpec.
func (in *NicSelectorSpec) DeepCopy() *NicSelectorSpec {
	if in == nil {
		return nil
	}
	out := new(NicSelectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PciPerformanceOptimizedSpec) DeepCopyInto(out *PciPerformanceOptimizedSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PciPerformanceOptimizedSpec.
func (in *PciPerformanceOptimizedSpec) DeepCopy() *PciPerformanceOptimizedSpec {
	if in == nil {
		return nil
	}
	out := new(PciPerformanceOptimizedSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QosSpec) DeepCopyInto(out *QosSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QosSpec.
func (in *QosSpec) DeepCopy() *QosSpec {
	if in == nil {
		return nil
	}
	out := new(QosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoceOptimizedSpec) DeepCopyInto(out *RoceOptimizedSpec) {
	*out = *in
	if in.Qos != nil {
		in, out := &in.Qos, &out.Qos
		*out = new(QosSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoceOptimizedSpec.
func (in *RoceOptimizedSpec) DeepCopy() *RoceOptimizedSpec {
	if in == nil {
		return nil
	}
	out := new(RoceOptimizedSpec)
	in.DeepCopyInto(out)
	return out
}
