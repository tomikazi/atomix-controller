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
	"github.com/atomix/kubernetes-controller/pkg/client/v1beta1/clientset/versioned/scheme"
	v1beta1 "github.com/atomix/kubernetes-controller/pkg/apis/cloud/v1beta1"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type CloudV1beta1Interface interface {
	RESTClient() rest.Interface
	ClustersGetter
	DatabasesGetter
	PartitionsGetter
	ProtocolsGetter
}

// CloudV1beta1Client is used to interact with features provided by the cloud.atomix.io group.
type CloudV1beta1Client struct {
	restClient rest.Interface
}

func (c *CloudV1beta1Client) Clusters(namespace string) ClusterInterface {
	return newClusters(c, namespace)
}

func (c *CloudV1beta1Client) Databases(namespace string) DatabaseInterface {
	return newDatabases(c, namespace)
}

func (c *CloudV1beta1Client) Partitions(namespace string) PartitionInterface {
	return newPartitions(c, namespace)
}

func (c *CloudV1beta1Client) Protocols(namespace string) ProtocolInterface {
	return newProtocols(c, namespace)
}

// NewForConfig creates a new CloudV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*CloudV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CloudV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new CloudV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CloudV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CloudV1beta1Client for the given RESTClient.
func New(c rest.Interface) *CloudV1beta1Client {
	return &CloudV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CloudV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
