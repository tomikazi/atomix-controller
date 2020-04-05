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

package v1beta1

import (
	"time"

	scheme "github.com/atomix/kubernetes-controller/pkg/client/v1beta1/clientset/versioned/scheme"
	v1beta1 "github.com/atomix/kubernetes-controller/pkg/apis/cloud/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProtocolsGetter has a method to return a ProtocolInterface.
// A group's client should implement this interface.
type ProtocolsGetter interface {
	Protocols(namespace string) ProtocolInterface
}

// ProtocolInterface has methods to work with Protocol resources.
type ProtocolInterface interface {
	Create(*v1beta1.Protocol) (*v1beta1.Protocol, error)
	Update(*v1beta1.Protocol) (*v1beta1.Protocol, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.Protocol, error)
	List(opts v1.ListOptions) (*v1beta1.ProtocolList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Protocol, err error)
	ProtocolExpansion
}

// protocols implements ProtocolInterface
type protocols struct {
	client rest.Interface
	ns     string
}

// newProtocols returns a Protocols
func newProtocols(c *CloudV1beta1Client, namespace string) *protocols {
	return &protocols{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the protocol, and returns the corresponding protocol object, and an error if there is any.
func (c *protocols) Get(name string, options v1.GetOptions) (result *v1beta1.Protocol, err error) {
	result = &v1beta1.Protocol{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("protocols").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Protocols that match those selectors.
func (c *protocols) List(opts v1.ListOptions) (result *v1beta1.ProtocolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ProtocolList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("protocols").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested protocols.
func (c *protocols) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("protocols").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a protocol and creates it.  Returns the server's representation of the protocol, and an error, if there is any.
func (c *protocols) Create(protocol *v1beta1.Protocol) (result *v1beta1.Protocol, err error) {
	result = &v1beta1.Protocol{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("protocols").
		Body(protocol).
		Do().
		Into(result)
	return
}

// Update takes the representation of a protocol and updates it. Returns the server's representation of the protocol, and an error, if there is any.
func (c *protocols) Update(protocol *v1beta1.Protocol) (result *v1beta1.Protocol, err error) {
	result = &v1beta1.Protocol{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("protocols").
		Name(protocol.Name).
		Body(protocol).
		Do().
		Into(result)
	return
}

// Delete takes name of the protocol and deletes it. Returns an error if one occurs.
func (c *protocols) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("protocols").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *protocols) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("protocols").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched protocol.
func (c *protocols) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Protocol, err error) {
	result = &v1beta1.Protocol{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("protocols").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
