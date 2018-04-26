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

// AwsGlueJobsGetter has a method to return a AwsGlueJobInterface.
// A group's client should implement this interface.
type AwsGlueJobsGetter interface {
	AwsGlueJobs(namespace string) AwsGlueJobInterface
}

// AwsGlueJobInterface has methods to work with AwsGlueJob resources.
type AwsGlueJobInterface interface {
	Create(*v1.AwsGlueJob) (*v1.AwsGlueJob, error)
	Update(*v1.AwsGlueJob) (*v1.AwsGlueJob, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsGlueJob, error)
	List(opts meta_v1.ListOptions) (*v1.AwsGlueJobList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsGlueJob, err error)
	AwsGlueJobExpansion
}

// awsGlueJobs implements AwsGlueJobInterface
type awsGlueJobs struct {
	client rest.Interface
	ns     string
}

// newAwsGlueJobs returns a AwsGlueJobs
func newAwsGlueJobs(c *TrussleV1Client, namespace string) *awsGlueJobs {
	return &awsGlueJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsGlueJob, and returns the corresponding awsGlueJob object, and an error if there is any.
func (c *awsGlueJobs) Get(name string, options meta_v1.GetOptions) (result *v1.AwsGlueJob, err error) {
	result = &v1.AwsGlueJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsgluejobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsGlueJobs that match those selectors.
func (c *awsGlueJobs) List(opts meta_v1.ListOptions) (result *v1.AwsGlueJobList, err error) {
	result = &v1.AwsGlueJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsgluejobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsGlueJobs.
func (c *awsGlueJobs) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsgluejobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsGlueJob and creates it.  Returns the server's representation of the awsGlueJob, and an error, if there is any.
func (c *awsGlueJobs) Create(awsGlueJob *v1.AwsGlueJob) (result *v1.AwsGlueJob, err error) {
	result = &v1.AwsGlueJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsgluejobs").
		Body(awsGlueJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsGlueJob and updates it. Returns the server's representation of the awsGlueJob, and an error, if there is any.
func (c *awsGlueJobs) Update(awsGlueJob *v1.AwsGlueJob) (result *v1.AwsGlueJob, err error) {
	result = &v1.AwsGlueJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsgluejobs").
		Name(awsGlueJob.Name).
		Body(awsGlueJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsGlueJob and deletes it. Returns an error if one occurs.
func (c *awsGlueJobs) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsgluejobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsGlueJobs) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsgluejobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsGlueJob.
func (c *awsGlueJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsGlueJob, err error) {
	result = &v1.AwsGlueJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsgluejobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}