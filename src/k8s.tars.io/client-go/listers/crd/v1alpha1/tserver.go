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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "k8s.tars.io/api/crd/v1alpha1"
)

// TServerLister helps list TServers.
type TServerLister interface {
	// List lists all TServers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TServer, err error)
	// TServers returns an object that can list and get TServers.
	TServers(namespace string) TServerNamespaceLister
	TServerListerExpansion
}

// tServerLister implements the TServerLister interface.
type tServerLister struct {
	indexer cache.Indexer
}

// NewTServerLister returns a new TServerLister.
func NewTServerLister(indexer cache.Indexer) TServerLister {
	return &tServerLister{indexer: indexer}
}

// List lists all TServers in the indexer.
func (s *tServerLister) List(selector labels.Selector) (ret []*v1alpha1.TServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TServer))
	})
	return ret, err
}

// TServers returns an object that can list and get TServers.
func (s *tServerLister) TServers(namespace string) TServerNamespaceLister {
	return tServerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TServerNamespaceLister helps list and get TServers.
type TServerNamespaceLister interface {
	// List lists all TServers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TServer, err error)
	// Get retrieves the TServer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TServer, error)
	TServerNamespaceListerExpansion
}

// tServerNamespaceLister implements the TServerNamespaceLister
// interface.
type tServerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TServers in the indexer for a given namespace.
func (s tServerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TServer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TServer))
	})
	return ret, err
}

// Get retrieves the TServer from the indexer for a given namespace and name.
func (s tServerNamespaceLister) Get(name string) (*v1alpha1.TServer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tserver"), name)
	}
	return obj.(*v1alpha1.TServer), nil
}
