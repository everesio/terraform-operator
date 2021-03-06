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

package fake

import (
	aws_v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAwsElasticBeanstalkApplications implements AwsElasticBeanstalkApplicationInterface
type FakeAwsElasticBeanstalkApplications struct {
	Fake *FakeTrussleV1
	ns   string
}

var awselasticbeanstalkapplicationsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awselasticbeanstalkapplications"}

var awselasticbeanstalkapplicationsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsElasticBeanstalkApplication"}

// Get takes name of the awsElasticBeanstalkApplication, and returns the corresponding awsElasticBeanstalkApplication object, and an error if there is any.
func (c *FakeAwsElasticBeanstalkApplications) Get(name string, options v1.GetOptions) (result *aws_v1.AwsElasticBeanstalkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awselasticbeanstalkapplicationsResource, c.ns, name), &aws_v1.AwsElasticBeanstalkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsElasticBeanstalkApplication), err
}

// List takes label and field selectors, and returns the list of AwsElasticBeanstalkApplications that match those selectors.
func (c *FakeAwsElasticBeanstalkApplications) List(opts v1.ListOptions) (result *aws_v1.AwsElasticBeanstalkApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awselasticbeanstalkapplicationsResource, awselasticbeanstalkapplicationsKind, c.ns, opts), &aws_v1.AwsElasticBeanstalkApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsElasticBeanstalkApplicationList{}
	for _, item := range obj.(*aws_v1.AwsElasticBeanstalkApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsElasticBeanstalkApplications.
func (c *FakeAwsElasticBeanstalkApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awselasticbeanstalkapplicationsResource, c.ns, opts))

}

// Create takes the representation of a awsElasticBeanstalkApplication and creates it.  Returns the server's representation of the awsElasticBeanstalkApplication, and an error, if there is any.
func (c *FakeAwsElasticBeanstalkApplications) Create(awsElasticBeanstalkApplication *aws_v1.AwsElasticBeanstalkApplication) (result *aws_v1.AwsElasticBeanstalkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awselasticbeanstalkapplicationsResource, c.ns, awsElasticBeanstalkApplication), &aws_v1.AwsElasticBeanstalkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsElasticBeanstalkApplication), err
}

// Update takes the representation of a awsElasticBeanstalkApplication and updates it. Returns the server's representation of the awsElasticBeanstalkApplication, and an error, if there is any.
func (c *FakeAwsElasticBeanstalkApplications) Update(awsElasticBeanstalkApplication *aws_v1.AwsElasticBeanstalkApplication) (result *aws_v1.AwsElasticBeanstalkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awselasticbeanstalkapplicationsResource, c.ns, awsElasticBeanstalkApplication), &aws_v1.AwsElasticBeanstalkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsElasticBeanstalkApplication), err
}

// Delete takes name of the awsElasticBeanstalkApplication and deletes it. Returns an error if one occurs.
func (c *FakeAwsElasticBeanstalkApplications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awselasticbeanstalkapplicationsResource, c.ns, name), &aws_v1.AwsElasticBeanstalkApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsElasticBeanstalkApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awselasticbeanstalkapplicationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsElasticBeanstalkApplicationList{})
	return err
}

// Patch applies the patch and returns the patched awsElasticBeanstalkApplication.
func (c *FakeAwsElasticBeanstalkApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsElasticBeanstalkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awselasticbeanstalkapplicationsResource, c.ns, name, data, subresources...), &aws_v1.AwsElasticBeanstalkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsElasticBeanstalkApplication), err
}
