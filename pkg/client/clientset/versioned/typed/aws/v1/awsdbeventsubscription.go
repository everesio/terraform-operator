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

package v1

import (
	v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	scheme "github.com/trussle/terraform-operator/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AwsDbEventSubscriptionsGetter has a method to return a AwsDbEventSubscriptionInterface.
// A group's client should implement this interface.
type AwsDbEventSubscriptionsGetter interface {
	AwsDbEventSubscriptions(namespace string) AwsDbEventSubscriptionInterface
}

// AwsDbEventSubscriptionInterface has methods to work with AwsDbEventSubscription resources.
type AwsDbEventSubscriptionInterface interface {
	Create(*v1.AwsDbEventSubscription) (*v1.AwsDbEventSubscription, error)
	Update(*v1.AwsDbEventSubscription) (*v1.AwsDbEventSubscription, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsDbEventSubscription, error)
	List(opts meta_v1.ListOptions) (*v1.AwsDbEventSubscriptionList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsDbEventSubscription, err error)
	AwsDbEventSubscriptionExpansion
}

// awsDbEventSubscriptions implements AwsDbEventSubscriptionInterface
type awsDbEventSubscriptions struct {
	client rest.Interface
	ns     string
}

// newAwsDbEventSubscriptions returns a AwsDbEventSubscriptions
func newAwsDbEventSubscriptions(c *TrussleV1Client, namespace string) *awsDbEventSubscriptions {
	return &awsDbEventSubscriptions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsDbEventSubscription, and returns the corresponding awsDbEventSubscription object, and an error if there is any.
func (c *awsDbEventSubscriptions) Get(name string, options meta_v1.GetOptions) (result *v1.AwsDbEventSubscription, err error) {
	result = &v1.AwsDbEventSubscription{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdbeventsubscriptions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDbEventSubscriptions that match those selectors.
func (c *awsDbEventSubscriptions) List(opts meta_v1.ListOptions) (result *v1.AwsDbEventSubscriptionList, err error) {
	result = &v1.AwsDbEventSubscriptionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdbeventsubscriptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDbEventSubscriptions.
func (c *awsDbEventSubscriptions) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsdbeventsubscriptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsDbEventSubscription and creates it.  Returns the server's representation of the awsDbEventSubscription, and an error, if there is any.
func (c *awsDbEventSubscriptions) Create(awsDbEventSubscription *v1.AwsDbEventSubscription) (result *v1.AwsDbEventSubscription, err error) {
	result = &v1.AwsDbEventSubscription{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsdbeventsubscriptions").
		Body(awsDbEventSubscription).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDbEventSubscription and updates it. Returns the server's representation of the awsDbEventSubscription, and an error, if there is any.
func (c *awsDbEventSubscriptions) Update(awsDbEventSubscription *v1.AwsDbEventSubscription) (result *v1.AwsDbEventSubscription, err error) {
	result = &v1.AwsDbEventSubscription{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsdbeventsubscriptions").
		Name(awsDbEventSubscription.Name).
		Body(awsDbEventSubscription).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDbEventSubscription and deletes it. Returns an error if one occurs.
func (c *awsDbEventSubscriptions) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdbeventsubscriptions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDbEventSubscriptions) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdbeventsubscriptions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDbEventSubscription.
func (c *awsDbEventSubscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsDbEventSubscription, err error) {
	result = &v1.AwsDbEventSubscription{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsdbeventsubscriptions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
