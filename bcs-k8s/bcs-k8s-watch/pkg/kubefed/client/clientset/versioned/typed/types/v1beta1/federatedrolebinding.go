/*
Copyright 2020 The Kubernetes Authors.

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

package v1beta1

import (
	v1beta1 "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"
	scheme "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedRoleBindingsGetter has a method to return a FederatedRoleBindingInterface.
// A group's client should implement this interface.
type FederatedRoleBindingsGetter interface {
	FederatedRoleBindings(namespace string) FederatedRoleBindingInterface
}

// FederatedRoleBindingInterface has methods to work with FederatedRoleBinding resources.
type FederatedRoleBindingInterface interface {
	Create(*v1beta1.FederatedRoleBinding) (*v1beta1.FederatedRoleBinding, error)
	Update(*v1beta1.FederatedRoleBinding) (*v1beta1.FederatedRoleBinding, error)
	UpdateStatus(*v1beta1.FederatedRoleBinding) (*v1beta1.FederatedRoleBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.FederatedRoleBinding, error)
	List(opts v1.ListOptions) (*v1beta1.FederatedRoleBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedRoleBinding, err error)
	FederatedRoleBindingExpansion
}

// federatedRoleBindings implements FederatedRoleBindingInterface
type federatedRoleBindings struct {
	client rest.Interface
	ns     string
}

// newFederatedRoleBindings returns a FederatedRoleBindings
func newFederatedRoleBindings(c *TypesV1beta1Client, namespace string) *federatedRoleBindings {
	return &federatedRoleBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedRoleBinding, and returns the corresponding federatedRoleBinding object, and an error if there is any.
func (c *federatedRoleBindings) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedRoleBinding, err error) {
	result = &v1beta1.FederatedRoleBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedrolebindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedRoleBindings that match those selectors.
func (c *federatedRoleBindings) List(opts v1.ListOptions) (result *v1beta1.FederatedRoleBindingList, err error) {
	result = &v1beta1.FederatedRoleBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedRoleBindings.
func (c *federatedRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedRoleBinding and creates it.  Returns the server's representation of the federatedRoleBinding, and an error, if there is any.
func (c *federatedRoleBindings) Create(federatedRoleBinding *v1beta1.FederatedRoleBinding) (result *v1beta1.FederatedRoleBinding, err error) {
	result = &v1beta1.FederatedRoleBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedrolebindings").
		Body(federatedRoleBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedRoleBinding and updates it. Returns the server's representation of the federatedRoleBinding, and an error, if there is any.
func (c *federatedRoleBindings) Update(federatedRoleBinding *v1beta1.FederatedRoleBinding) (result *v1beta1.FederatedRoleBinding, err error) {
	result = &v1beta1.FederatedRoleBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedrolebindings").
		Name(federatedRoleBinding.Name).
		Body(federatedRoleBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedRoleBindings) UpdateStatus(federatedRoleBinding *v1beta1.FederatedRoleBinding) (result *v1beta1.FederatedRoleBinding, err error) {
	result = &v1beta1.FederatedRoleBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedrolebindings").
		Name(federatedRoleBinding.Name).
		SubResource("status").
		Body(federatedRoleBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedRoleBinding and deletes it. Returns an error if one occurs.
func (c *federatedRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedrolebindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedrolebindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedRoleBinding.
func (c *federatedRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedRoleBinding, err error) {
	result = &v1beta1.FederatedRoleBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedrolebindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}