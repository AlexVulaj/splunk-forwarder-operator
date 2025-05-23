//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplunkFilter) DeepCopyInto(out *SplunkFilter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplunkFilter.
func (in *SplunkFilter) DeepCopy() *SplunkFilter {
	if in == nil {
		return nil
	}
	out := new(SplunkFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplunkForwarder) DeepCopyInto(out *SplunkForwarder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplunkForwarder.
func (in *SplunkForwarder) DeepCopy() *SplunkForwarder {
	if in == nil {
		return nil
	}
	out := new(SplunkForwarder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SplunkForwarder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplunkForwarderInputs) DeepCopyInto(out *SplunkForwarderInputs) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplunkForwarderInputs.
func (in *SplunkForwarderInputs) DeepCopy() *SplunkForwarderInputs {
	if in == nil {
		return nil
	}
	out := new(SplunkForwarderInputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplunkForwarderList) DeepCopyInto(out *SplunkForwarderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SplunkForwarder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplunkForwarderList.
func (in *SplunkForwarderList) DeepCopy() *SplunkForwarderList {
	if in == nil {
		return nil
	}
	out := new(SplunkForwarderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SplunkForwarderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplunkForwarderSpec) DeepCopyInto(out *SplunkForwarderSpec) {
	*out = *in
	if in.SplunkInputs != nil {
		in, out := &in.SplunkInputs, &out.SplunkInputs
		*out = make([]SplunkForwarderInputs, len(*in))
		copy(*out, *in)
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]SplunkFilter, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplunkForwarderSpec.
func (in *SplunkForwarderSpec) DeepCopy() *SplunkForwarderSpec {
	if in == nil {
		return nil
	}
	out := new(SplunkForwarderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplunkForwarderStatus) DeepCopyInto(out *SplunkForwarderStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplunkForwarderStatus.
func (in *SplunkForwarderStatus) DeepCopy() *SplunkForwarderStatus {
	if in == nil {
		return nil
	}
	out := new(SplunkForwarderStatus)
	in.DeepCopyInto(out)
	return out
}
