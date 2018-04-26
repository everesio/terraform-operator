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

// AwsAutoscalingAttachmentLister helps list AwsAutoscalingAttachments.
type AwsAutoscalingAttachmentLister interface {
	// List lists all AwsAutoscalingAttachments in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsAutoscalingAttachment, err error)
	// AwsAutoscalingAttachments returns an object that can list and get AwsAutoscalingAttachments.
	AwsAutoscalingAttachments(namespace string) AwsAutoscalingAttachmentNamespaceLister
	AwsAutoscalingAttachmentListerExpansion
}

// awsAutoscalingAttachmentLister implements the AwsAutoscalingAttachmentLister interface.
type awsAutoscalingAttachmentLister struct {
	indexer cache.Indexer
}

// NewAwsAutoscalingAttachmentLister returns a new AwsAutoscalingAttachmentLister.
func NewAwsAutoscalingAttachmentLister(indexer cache.Indexer) AwsAutoscalingAttachmentLister {
	return &awsAutoscalingAttachmentLister{indexer: indexer}
}

// List lists all AwsAutoscalingAttachments in the indexer.
func (s *awsAutoscalingAttachmentLister) List(selector labels.Selector) (ret []*v1.AwsAutoscalingAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsAutoscalingAttachment))
	})
	return ret, err
}

// AwsAutoscalingAttachments returns an object that can list and get AwsAutoscalingAttachments.
func (s *awsAutoscalingAttachmentLister) AwsAutoscalingAttachments(namespace string) AwsAutoscalingAttachmentNamespaceLister {
	return awsAutoscalingAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsAutoscalingAttachmentNamespaceLister helps list and get AwsAutoscalingAttachments.
type AwsAutoscalingAttachmentNamespaceLister interface {
	// List lists all AwsAutoscalingAttachments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsAutoscalingAttachment, err error)
	// Get retrieves the AwsAutoscalingAttachment from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsAutoscalingAttachment, error)
	AwsAutoscalingAttachmentNamespaceListerExpansion
}

// awsAutoscalingAttachmentNamespaceLister implements the AwsAutoscalingAttachmentNamespaceLister
// interface.
type awsAutoscalingAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsAutoscalingAttachments in the indexer for a given namespace.
func (s awsAutoscalingAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsAutoscalingAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsAutoscalingAttachment))
	})
	return ret, err
}

// Get retrieves the AwsAutoscalingAttachment from the indexer for a given namespace and name.
func (s awsAutoscalingAttachmentNamespaceLister) Get(name string) (*v1.AwsAutoscalingAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsautoscalingattachment"), name)
	}
	return obj.(*v1.AwsAutoscalingAttachment), nil
}