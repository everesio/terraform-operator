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

// AwsS3BucketPoliciesGetter has a method to return a AwsS3BucketPolicyInterface.
// A group's client should implement this interface.
type AwsS3BucketPoliciesGetter interface {
	AwsS3BucketPolicies(namespace string) AwsS3BucketPolicyInterface
}

// AwsS3BucketPolicyInterface has methods to work with AwsS3BucketPolicy resources.
type AwsS3BucketPolicyInterface interface {
	Create(*v1.AwsS3BucketPolicy) (*v1.AwsS3BucketPolicy, error)
	Update(*v1.AwsS3BucketPolicy) (*v1.AwsS3BucketPolicy, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsS3BucketPolicy, error)
	List(opts meta_v1.ListOptions) (*v1.AwsS3BucketPolicyList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsS3BucketPolicy, err error)
	AwsS3BucketPolicyExpansion
}

// awsS3BucketPolicies implements AwsS3BucketPolicyInterface
type awsS3BucketPolicies struct {
	client rest.Interface
	ns     string
}

// newAwsS3BucketPolicies returns a AwsS3BucketPolicies
func newAwsS3BucketPolicies(c *TrussleV1Client, namespace string) *awsS3BucketPolicies {
	return &awsS3BucketPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsS3BucketPolicy, and returns the corresponding awsS3BucketPolicy object, and an error if there is any.
func (c *awsS3BucketPolicies) Get(name string, options meta_v1.GetOptions) (result *v1.AwsS3BucketPolicy, err error) {
	result = &v1.AwsS3BucketPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awss3bucketpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsS3BucketPolicies that match those selectors.
func (c *awsS3BucketPolicies) List(opts meta_v1.ListOptions) (result *v1.AwsS3BucketPolicyList, err error) {
	result = &v1.AwsS3BucketPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awss3bucketpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsS3BucketPolicies.
func (c *awsS3BucketPolicies) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awss3bucketpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsS3BucketPolicy and creates it.  Returns the server's representation of the awsS3BucketPolicy, and an error, if there is any.
func (c *awsS3BucketPolicies) Create(awsS3BucketPolicy *v1.AwsS3BucketPolicy) (result *v1.AwsS3BucketPolicy, err error) {
	result = &v1.AwsS3BucketPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awss3bucketpolicies").
		Body(awsS3BucketPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsS3BucketPolicy and updates it. Returns the server's representation of the awsS3BucketPolicy, and an error, if there is any.
func (c *awsS3BucketPolicies) Update(awsS3BucketPolicy *v1.AwsS3BucketPolicy) (result *v1.AwsS3BucketPolicy, err error) {
	result = &v1.AwsS3BucketPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awss3bucketpolicies").
		Name(awsS3BucketPolicy.Name).
		Body(awsS3BucketPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsS3BucketPolicy and deletes it. Returns an error if one occurs.
func (c *awsS3BucketPolicies) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awss3bucketpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsS3BucketPolicies) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awss3bucketpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsS3BucketPolicy.
func (c *awsS3BucketPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsS3BucketPolicy, err error) {
	result = &v1.AwsS3BucketPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awss3bucketpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}