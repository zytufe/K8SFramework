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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha2 "k8s.tars.io/api/crd/v1alpha2"
)

// FakeTImages implements TImageInterface
type FakeTImages struct {
	Fake *FakeCrdV1alpha2
	ns   string
}

var timagesResource = schema.GroupVersionResource{Group: "crd", Version: "v1alpha2", Resource: "timages"}

var timagesKind = schema.GroupVersionKind{Group: "crd", Version: "v1alpha2", Kind: "TImage"}

// Get takes name of the tImage, and returns the corresponding tImage object, and an error if there is any.
func (c *FakeTImages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.TImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(timagesResource, c.ns, name), &v1alpha2.TImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TImage), err
}

// List takes label and field selectors, and returns the list of TImages that match those selectors.
func (c *FakeTImages) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.TImageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(timagesResource, timagesKind, c.ns, opts), &v1alpha2.TImageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.TImageList{ListMeta: obj.(*v1alpha2.TImageList).ListMeta}
	for _, item := range obj.(*v1alpha2.TImageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tImages.
func (c *FakeTImages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(timagesResource, c.ns, opts))

}

// Create takes the representation of a tImage and creates it.  Returns the server's representation of the tImage, and an error, if there is any.
func (c *FakeTImages) Create(ctx context.Context, tImage *v1alpha2.TImage, opts v1.CreateOptions) (result *v1alpha2.TImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(timagesResource, c.ns, tImage), &v1alpha2.TImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TImage), err
}

// Update takes the representation of a tImage and updates it. Returns the server's representation of the tImage, and an error, if there is any.
func (c *FakeTImages) Update(ctx context.Context, tImage *v1alpha2.TImage, opts v1.UpdateOptions) (result *v1alpha2.TImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(timagesResource, c.ns, tImage), &v1alpha2.TImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TImage), err
}

// Delete takes name of the tImage and deletes it. Returns an error if one occurs.
func (c *FakeTImages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(timagesResource, c.ns, name), &v1alpha2.TImage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTImages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(timagesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.TImageList{})
	return err
}

// Patch applies the patch and returns the patched tImage.
func (c *FakeTImages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.TImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(timagesResource, c.ns, name, pt, data, subresources...), &v1alpha2.TImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TImage), err
}
