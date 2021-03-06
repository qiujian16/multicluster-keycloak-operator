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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationDomain) DeepCopyInto(out *AuthorizationDomain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationDomain.
func (in *AuthorizationDomain) DeepCopy() *AuthorizationDomain {
	if in == nil {
		return nil
	}
	out := new(AuthorizationDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthorizationDomain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationDomainList) DeepCopyInto(out *AuthorizationDomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthorizationDomain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationDomainList.
func (in *AuthorizationDomainList) DeepCopy() *AuthorizationDomainList {
	if in == nil {
		return nil
	}
	out := new(AuthorizationDomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthorizationDomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationDomainSpec) DeepCopyInto(out *AuthorizationDomainSpec) {
	*out = *in
	if in.IdentityProviders != nil {
		in, out := &in.IdentityProviders, &out.IdentityProviders
		*out = make([]IdentityProvider, len(*in))
		copy(*out, *in)
	}
	out.IssuerCertificate = in.IssuerCertificate
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationDomainSpec.
func (in *AuthorizationDomainSpec) DeepCopy() *AuthorizationDomainSpec {
	if in == nil {
		return nil
	}
	out := new(AuthorizationDomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationDomainStatus) DeepCopyInto(out *AuthorizationDomainStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationDomainStatus.
func (in *AuthorizationDomainStatus) DeepCopy() *AuthorizationDomainStatus {
	if in == nil {
		return nil
	}
	out := new(AuthorizationDomainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProvider) DeepCopyInto(out *IdentityProvider) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProvider.
func (in *IdentityProvider) DeepCopy() *IdentityProvider {
	if in == nil {
		return nil
	}
	out := new(IdentityProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerCertificate) DeepCopyInto(out *IssuerCertificate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerCertificate.
func (in *IssuerCertificate) DeepCopy() *IssuerCertificate {
	if in == nil {
		return nil
	}
	out := new(IssuerCertificate)
	in.DeepCopyInto(out)
	return out
}
