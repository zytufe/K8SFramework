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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha2 "k8s.tars.io/api/crd/v1alpha2"
)

// TDeployLister helps list TDeploys.
type TDeployLister interface {
	// List lists all TDeploys in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.TDeploy, err error)
	// TDeploys returns an object that can list and get TDeploys.
	TDeploys(namespace string) TDeployNamespaceLister
	TDeployListerExpansion
}

// tDeployLister implements the TDeployLister interface.
type tDeployLister struct {
	indexer cache.Indexer
}

// NewTDeployLister returns a new TDeployLister.
func NewTDeployLister(indexer cache.Indexer) TDeployLister {
	return &tDeployLister{indexer: indexer}
}

// List lists all TDeploys in the indexer.
func (s *tDeployLister) List(selector labels.Selector) (ret []*v1alpha2.TDeploy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.TDeploy))
	})
	return ret, err
}

// TDeploys returns an object that can list and get TDeploys.
func (s *tDeployLister) TDeploys(namespace string) TDeployNamespaceLister {
	return tDeployNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TDeployNamespaceLister helps list and get TDeploys.
type TDeployNamespaceLister interface {
	// List lists all TDeploys in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.TDeploy, err error)
	// Get retrieves the TDeploy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.TDeploy, error)
	TDeployNamespaceListerExpansion
}

// tDeployNamespaceLister implements the TDeployNamespaceLister
// interface.
type tDeployNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TDeploys in the indexer for a given namespace.
func (s tDeployNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.TDeploy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.TDeploy))
	})
	return ret, err
}

// Get retrieves the TDeploy from the indexer for a given namespace and name.
func (s tDeployNamespaceLister) Get(name string) (*v1alpha2.TDeploy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("tdeploy"), name)
	}
	return obj.(*v1alpha2.TDeploy), nil
}
