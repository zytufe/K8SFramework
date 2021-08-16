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

// TTreeLister helps list TTrees.
type TTreeLister interface {
	// List lists all TTrees in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.TTree, err error)
	// TTrees returns an object that can list and get TTrees.
	TTrees(namespace string) TTreeNamespaceLister
	TTreeListerExpansion
}

// tTreeLister implements the TTreeLister interface.
type tTreeLister struct {
	indexer cache.Indexer
}

// NewTTreeLister returns a new TTreeLister.
func NewTTreeLister(indexer cache.Indexer) TTreeLister {
	return &tTreeLister{indexer: indexer}
}

// List lists all TTrees in the indexer.
func (s *tTreeLister) List(selector labels.Selector) (ret []*v1alpha2.TTree, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.TTree))
	})
	return ret, err
}

// TTrees returns an object that can list and get TTrees.
func (s *tTreeLister) TTrees(namespace string) TTreeNamespaceLister {
	return tTreeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TTreeNamespaceLister helps list and get TTrees.
type TTreeNamespaceLister interface {
	// List lists all TTrees in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.TTree, err error)
	// Get retrieves the TTree from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.TTree, error)
	TTreeNamespaceListerExpansion
}

// tTreeNamespaceLister implements the TTreeNamespaceLister
// interface.
type tTreeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TTrees in the indexer for a given namespace.
func (s tTreeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.TTree, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.TTree))
	})
	return ret, err
}

// Get retrieves the TTree from the indexer for a given namespace and name.
func (s tTreeNamespaceLister) Get(name string) (*v1alpha2.TTree, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("ttree"), name)
	}
	return obj.(*v1alpha2.TTree), nil
}
