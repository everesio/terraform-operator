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

package v1

import (
	v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AwsDynamodbGlobalTableLister helps list AwsDynamodbGlobalTables.
type AwsDynamodbGlobalTableLister interface {
	// List lists all AwsDynamodbGlobalTables in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsDynamodbGlobalTable, err error)
	// AwsDynamodbGlobalTables returns an object that can list and get AwsDynamodbGlobalTables.
	AwsDynamodbGlobalTables(namespace string) AwsDynamodbGlobalTableNamespaceLister
	AwsDynamodbGlobalTableListerExpansion
}

// awsDynamodbGlobalTableLister implements the AwsDynamodbGlobalTableLister interface.
type awsDynamodbGlobalTableLister struct {
	indexer cache.Indexer
}

// NewAwsDynamodbGlobalTableLister returns a new AwsDynamodbGlobalTableLister.
func NewAwsDynamodbGlobalTableLister(indexer cache.Indexer) AwsDynamodbGlobalTableLister {
	return &awsDynamodbGlobalTableLister{indexer: indexer}
}

// List lists all AwsDynamodbGlobalTables in the indexer.
func (s *awsDynamodbGlobalTableLister) List(selector labels.Selector) (ret []*v1.AwsDynamodbGlobalTable, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDynamodbGlobalTable))
	})
	return ret, err
}

// AwsDynamodbGlobalTables returns an object that can list and get AwsDynamodbGlobalTables.
func (s *awsDynamodbGlobalTableLister) AwsDynamodbGlobalTables(namespace string) AwsDynamodbGlobalTableNamespaceLister {
	return awsDynamodbGlobalTableNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsDynamodbGlobalTableNamespaceLister helps list and get AwsDynamodbGlobalTables.
type AwsDynamodbGlobalTableNamespaceLister interface {
	// List lists all AwsDynamodbGlobalTables in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsDynamodbGlobalTable, err error)
	// Get retrieves the AwsDynamodbGlobalTable from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsDynamodbGlobalTable, error)
	AwsDynamodbGlobalTableNamespaceListerExpansion
}

// awsDynamodbGlobalTableNamespaceLister implements the AwsDynamodbGlobalTableNamespaceLister
// interface.
type awsDynamodbGlobalTableNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsDynamodbGlobalTables in the indexer for a given namespace.
func (s awsDynamodbGlobalTableNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsDynamodbGlobalTable, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDynamodbGlobalTable))
	})
	return ret, err
}

// Get retrieves the AwsDynamodbGlobalTable from the indexer for a given namespace and name.
func (s awsDynamodbGlobalTableNamespaceLister) Get(name string) (*v1.AwsDynamodbGlobalTable, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsdynamodbglobaltable"), name)
	}
	return obj.(*v1.AwsDynamodbGlobalTable), nil
}
