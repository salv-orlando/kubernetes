// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package authorization

import (
	api "k8s.io/kubernetes/pkg/api"
	conversion "k8s.io/kubernetes/pkg/conversion"
	reflect "reflect"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_LocalSubjectAccessReview, InType: reflect.TypeOf(func() *LocalSubjectAccessReview { var x *LocalSubjectAccessReview; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_NonResourceAttributes, InType: reflect.TypeOf(func() *NonResourceAttributes { var x *NonResourceAttributes; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_ResourceAttributes, InType: reflect.TypeOf(func() *ResourceAttributes { var x *ResourceAttributes; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_SelfSubjectAccessReview, InType: reflect.TypeOf(func() *SelfSubjectAccessReview { var x *SelfSubjectAccessReview; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_SelfSubjectAccessReviewSpec, InType: reflect.TypeOf(func() *SelfSubjectAccessReviewSpec { var x *SelfSubjectAccessReviewSpec; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_SubjectAccessReview, InType: reflect.TypeOf(func() *SubjectAccessReview { var x *SubjectAccessReview; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_SubjectAccessReviewSpec, InType: reflect.TypeOf(func() *SubjectAccessReviewSpec { var x *SubjectAccessReviewSpec; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_SubjectAccessReviewStatus, InType: reflect.TypeOf(func() *SubjectAccessReviewStatus { var x *SubjectAccessReviewStatus; return x }())},
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_authorization_LocalSubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LocalSubjectAccessReview)
		out := out.(*LocalSubjectAccessReview)
		out.TypeMeta = in.TypeMeta
		if err := DeepCopy_authorization_SubjectAccessReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_authorization_NonResourceAttributes(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NonResourceAttributes)
		out := out.(*NonResourceAttributes)
		out.Path = in.Path
		out.Verb = in.Verb
		return nil
	}
}

func DeepCopy_authorization_ResourceAttributes(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ResourceAttributes)
		out := out.(*ResourceAttributes)
		out.Namespace = in.Namespace
		out.Verb = in.Verb
		out.Group = in.Group
		out.Version = in.Version
		out.Resource = in.Resource
		out.Subresource = in.Subresource
		out.Name = in.Name
		return nil
	}
}

func DeepCopy_authorization_SelfSubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SelfSubjectAccessReview)
		out := out.(*SelfSubjectAccessReview)
		out.TypeMeta = in.TypeMeta
		if err := DeepCopy_authorization_SelfSubjectAccessReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_authorization_SelfSubjectAccessReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SelfSubjectAccessReviewSpec)
		out := out.(*SelfSubjectAccessReviewSpec)
		if in.ResourceAttributes != nil {
			in, out := &in.ResourceAttributes, &out.ResourceAttributes
			*out = new(ResourceAttributes)
			**out = **in
		} else {
			out.ResourceAttributes = nil
		}
		if in.NonResourceAttributes != nil {
			in, out := &in.NonResourceAttributes, &out.NonResourceAttributes
			*out = new(NonResourceAttributes)
			**out = **in
		} else {
			out.NonResourceAttributes = nil
		}
		return nil
	}
}

func DeepCopy_authorization_SubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReview)
		out := out.(*SubjectAccessReview)
		out.TypeMeta = in.TypeMeta
		if err := DeepCopy_authorization_SubjectAccessReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_authorization_SubjectAccessReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReviewSpec)
		out := out.(*SubjectAccessReviewSpec)
		if in.ResourceAttributes != nil {
			in, out := &in.ResourceAttributes, &out.ResourceAttributes
			*out = new(ResourceAttributes)
			**out = **in
		} else {
			out.ResourceAttributes = nil
		}
		if in.NonResourceAttributes != nil {
			in, out := &in.NonResourceAttributes, &out.NonResourceAttributes
			*out = new(NonResourceAttributes)
			**out = **in
		} else {
			out.NonResourceAttributes = nil
		}
		out.User = in.User
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Groups = nil
		}
		if in.Extra != nil {
			in, out := &in.Extra, &out.Extra
			*out = make(map[string][]string)
			for key, val := range *in {
				if newVal, err := c.DeepCopy(&val); err != nil {
					return err
				} else {
					(*out)[key] = *newVal.(*[]string)
				}
			}
		} else {
			out.Extra = nil
		}
		return nil
	}
}

func DeepCopy_authorization_SubjectAccessReviewStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReviewStatus)
		out := out.(*SubjectAccessReviewStatus)
		out.Allowed = in.Allowed
		out.Reason = in.Reason
		return nil
	}
}