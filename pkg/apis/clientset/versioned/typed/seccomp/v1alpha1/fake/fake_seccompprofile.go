/*
Copyright 2020 The Knative Authors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/imjasonh/seccomp-profile/pkg/apis/seccomp/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSeccompProfiles implements SeccompProfileInterface
type FakeSeccompProfiles struct {
	Fake *FakeSeccompV1alpha1
}

var seccompprofilesResource = schema.GroupVersionResource{Group: "seccomp.imjasonh.dev", Version: "v1alpha1", Resource: "seccompprofiles"}

var seccompprofilesKind = schema.GroupVersionKind{Group: "seccomp.imjasonh.dev", Version: "v1alpha1", Kind: "SeccompProfile"}

// Get takes name of the seccompProfile, and returns the corresponding seccompProfile object, and an error if there is any.
func (c *FakeSeccompProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SeccompProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(seccompprofilesResource, name), &v1alpha1.SeccompProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeccompProfile), err
}

// List takes label and field selectors, and returns the list of SeccompProfiles that match those selectors.
func (c *FakeSeccompProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SeccompProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(seccompprofilesResource, seccompprofilesKind, opts), &v1alpha1.SeccompProfileList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SeccompProfileList{ListMeta: obj.(*v1alpha1.SeccompProfileList).ListMeta}
	for _, item := range obj.(*v1alpha1.SeccompProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested seccompProfiles.
func (c *FakeSeccompProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(seccompprofilesResource, opts))
}

// Create takes the representation of a seccompProfile and creates it.  Returns the server's representation of the seccompProfile, and an error, if there is any.
func (c *FakeSeccompProfiles) Create(ctx context.Context, seccompProfile *v1alpha1.SeccompProfile, opts v1.CreateOptions) (result *v1alpha1.SeccompProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(seccompprofilesResource, seccompProfile), &v1alpha1.SeccompProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeccompProfile), err
}

// Update takes the representation of a seccompProfile and updates it. Returns the server's representation of the seccompProfile, and an error, if there is any.
func (c *FakeSeccompProfiles) Update(ctx context.Context, seccompProfile *v1alpha1.SeccompProfile, opts v1.UpdateOptions) (result *v1alpha1.SeccompProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(seccompprofilesResource, seccompProfile), &v1alpha1.SeccompProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeccompProfile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSeccompProfiles) UpdateStatus(ctx context.Context, seccompProfile *v1alpha1.SeccompProfile, opts v1.UpdateOptions) (*v1alpha1.SeccompProfile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(seccompprofilesResource, "status", seccompProfile), &v1alpha1.SeccompProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeccompProfile), err
}

// Delete takes name of the seccompProfile and deletes it. Returns an error if one occurs.
func (c *FakeSeccompProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(seccompprofilesResource, name, opts), &v1alpha1.SeccompProfile{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSeccompProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(seccompprofilesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SeccompProfileList{})
	return err
}

// Patch applies the patch and returns the patched seccompProfile.
func (c *FakeSeccompProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SeccompProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(seccompprofilesResource, name, pt, data, subresources...), &v1alpha1.SeccompProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeccompProfile), err
}
