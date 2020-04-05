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

// PartitionsGetter has a method to return a PartitionInterface.
// A group's client should implement this interface.
type PartitionsGetter interface {
	Partitions(namespace string) PartitionInterface
}

// PartitionInterface has methods to work with Partition resources.
type PartitionInterface interface {
	Create(*v1beta1.Partition) (*v1beta1.Partition, error)
	Update(*v1beta1.Partition) (*v1beta1.Partition, error)
	UpdateStatus(*v1beta1.Partition) (*v1beta1.Partition, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.Partition, error)
	List(opts v1.ListOptions) (*v1beta1.PartitionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Partition, err error)
	PartitionExpansion
}

// partitions implements PartitionInterface
type partitions struct {
	client rest.Interface
	ns     string
}

// newPartitions returns a Partitions
func newPartitions(c *CloudV1beta1Client, namespace string) *partitions {
	return &partitions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the partition, and returns the corresponding partition object, and an error if there is any.
func (c *partitions) Get(name string, options v1.GetOptions) (result *v1beta1.Partition, err error) {
	result = &v1beta1.Partition{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("partitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Partitions that match those selectors.
func (c *partitions) List(opts v1.ListOptions) (result *v1beta1.PartitionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.PartitionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("partitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested partitions.
func (c *partitions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("partitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a partition and creates it.  Returns the server's representation of the partition, and an error, if there is any.
func (c *partitions) Create(partition *v1beta1.Partition) (result *v1beta1.Partition, err error) {
	result = &v1beta1.Partition{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("partitions").
		Body(partition).
		Do().
		Into(result)
	return
}

// Update takes the representation of a partition and updates it. Returns the server's representation of the partition, and an error, if there is any.
func (c *partitions) Update(partition *v1beta1.Partition) (result *v1beta1.Partition, err error) {
	result = &v1beta1.Partition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("partitions").
		Name(partition.Name).
		Body(partition).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *partitions) UpdateStatus(partition *v1beta1.Partition) (result *v1beta1.Partition, err error) {
	result = &v1beta1.Partition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("partitions").
		Name(partition.Name).
		SubResource("status").
		Body(partition).
		Do().
		Into(result)
	return
}

// Delete takes name of the partition and deletes it. Returns an error if one occurs.
func (c *partitions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("partitions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *partitions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("partitions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched partition.
func (c *partitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Partition, err error) {
	result = &v1beta1.Partition{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("partitions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
