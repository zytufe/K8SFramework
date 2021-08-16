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

package v1alpha2

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha2 "k8s.tars.io/api/crd/v1alpha2"
	scheme "k8s.tars.io/client-go/clientset/versioned/scheme"
)

// TDeploysGetter has a method to return a TDeployInterface.
// A group's client should implement this interface.
type TDeploysGetter interface {
	TDeploys(namespace string) TDeployInterface
}

// TDeployInterface has methods to work with TDeploy resources.
type TDeployInterface interface {
	Create(ctx context.Context, tDeploy *v1alpha2.TDeploy, opts v1.CreateOptions) (*v1alpha2.TDeploy, error)
	Update(ctx context.Context, tDeploy *v1alpha2.TDeploy, opts v1.UpdateOptions) (*v1alpha2.TDeploy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.TDeploy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.TDeployList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.TDeploy, err error)
	TDeployExpansion
}

// tDeploys implements TDeployInterface
type tDeploys struct {
	client rest.Interface
	ns     string
}

// newTDeploys returns a TDeploys
func newTDeploys(c *CrdV1alpha2Client, namespace string) *tDeploys {
	return &tDeploys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tDeploy, and returns the corresponding tDeploy object, and an error if there is any.
func (c *tDeploys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.TDeploy, err error) {
	result = &v1alpha2.TDeploy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tdeploys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TDeploys that match those selectors.
func (c *tDeploys) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.TDeployList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.TDeployList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tdeploys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tDeploys.
func (c *tDeploys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tdeploys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tDeploy and creates it.  Returns the server's representation of the tDeploy, and an error, if there is any.
func (c *tDeploys) Create(ctx context.Context, tDeploy *v1alpha2.TDeploy, opts v1.CreateOptions) (result *v1alpha2.TDeploy, err error) {
	result = &v1alpha2.TDeploy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tdeploys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tDeploy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tDeploy and updates it. Returns the server's representation of the tDeploy, and an error, if there is any.
func (c *tDeploys) Update(ctx context.Context, tDeploy *v1alpha2.TDeploy, opts v1.UpdateOptions) (result *v1alpha2.TDeploy, err error) {
	result = &v1alpha2.TDeploy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tdeploys").
		Name(tDeploy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tDeploy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tDeploy and deletes it. Returns an error if one occurs.
func (c *tDeploys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tdeploys").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tDeploys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tdeploys").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tDeploy.
func (c *tDeploys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.TDeploy, err error) {
	result = &v1alpha2.TDeploy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tdeploys").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
