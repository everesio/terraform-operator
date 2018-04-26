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

// AwsDmsCertificateLister helps list AwsDmsCertificates.
type AwsDmsCertificateLister interface {
	// List lists all AwsDmsCertificates in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsDmsCertificate, err error)
	// AwsDmsCertificates returns an object that can list and get AwsDmsCertificates.
	AwsDmsCertificates(namespace string) AwsDmsCertificateNamespaceLister
	AwsDmsCertificateListerExpansion
}

// awsDmsCertificateLister implements the AwsDmsCertificateLister interface.
type awsDmsCertificateLister struct {
	indexer cache.Indexer
}

// NewAwsDmsCertificateLister returns a new AwsDmsCertificateLister.
func NewAwsDmsCertificateLister(indexer cache.Indexer) AwsDmsCertificateLister {
	return &awsDmsCertificateLister{indexer: indexer}
}

// List lists all AwsDmsCertificates in the indexer.
func (s *awsDmsCertificateLister) List(selector labels.Selector) (ret []*v1.AwsDmsCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDmsCertificate))
	})
	return ret, err
}

// AwsDmsCertificates returns an object that can list and get AwsDmsCertificates.
func (s *awsDmsCertificateLister) AwsDmsCertificates(namespace string) AwsDmsCertificateNamespaceLister {
	return awsDmsCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsDmsCertificateNamespaceLister helps list and get AwsDmsCertificates.
type AwsDmsCertificateNamespaceLister interface {
	// List lists all AwsDmsCertificates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsDmsCertificate, err error)
	// Get retrieves the AwsDmsCertificate from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsDmsCertificate, error)
	AwsDmsCertificateNamespaceListerExpansion
}

// awsDmsCertificateNamespaceLister implements the AwsDmsCertificateNamespaceLister
// interface.
type awsDmsCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsDmsCertificates in the indexer for a given namespace.
func (s awsDmsCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsDmsCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDmsCertificate))
	})
	return ret, err
}

// Get retrieves the AwsDmsCertificate from the indexer for a given namespace and name.
func (s awsDmsCertificateNamespaceLister) Get(name string) (*v1.AwsDmsCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsdmscertificate"), name)
	}
	return obj.(*v1.AwsDmsCertificate), nil
}