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

package v1alpha1

import (
	rest "k8s.io/client-go/rest"
	v1alpha1 "k8s.tars.io/api/crd/v1alpha1"
	"k8s.tars.io/client-go/clientset/versioned/scheme"
)

type CrdV1alpha1Interface interface {
	RESTClient() rest.Interface
	TAccountsGetter
	TConfigsGetter
	TDeploysGetter
	TEndpointsGetter
	TExitedRecordsGetter
	TImagesGetter
	TServersGetter
	TTemplatesGetter
	TTreesGetter
}

// CrdV1alpha1Client is used to interact with features provided by the crd group.
type CrdV1alpha1Client struct {
	restClient rest.Interface
}

func (c *CrdV1alpha1Client) TAccounts(namespace string) TAccountInterface {
	return newTAccounts(c, namespace)
}

func (c *CrdV1alpha1Client) TConfigs(namespace string) TConfigInterface {
	return newTConfigs(c, namespace)
}

func (c *CrdV1alpha1Client) TDeploys(namespace string) TDeployInterface {
	return newTDeploys(c, namespace)
}

func (c *CrdV1alpha1Client) TEndpoints(namespace string) TEndpointInterface {
	return newTEndpoints(c, namespace)
}

func (c *CrdV1alpha1Client) TExitedRecords(namespace string) TExitedRecordInterface {
	return newTExitedRecords(c, namespace)
}

func (c *CrdV1alpha1Client) TImages(namespace string) TImageInterface {
	return newTImages(c, namespace)
}

func (c *CrdV1alpha1Client) TServers(namespace string) TServerInterface {
	return newTServers(c, namespace)
}

func (c *CrdV1alpha1Client) TTemplates(namespace string) TTemplateInterface {
	return newTTemplates(c, namespace)
}

func (c *CrdV1alpha1Client) TTrees(namespace string) TTreeInterface {
	return newTTrees(c, namespace)
}

// NewForConfig creates a new CrdV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*CrdV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CrdV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new CrdV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CrdV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CrdV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *CrdV1alpha1Client {
	return &CrdV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CrdV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
