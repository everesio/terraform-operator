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

// FakeAwsSnsTopics implements AwsSnsTopicInterface
type FakeAwsSnsTopics struct {
	Fake *FakeTrussleV1
	ns   string
}

var awssnstopicsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awssnstopics"}

var awssnstopicsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsSnsTopic"}

// Get takes name of the awsSnsTopic, and returns the corresponding awsSnsTopic object, and an error if there is any.
func (c *FakeAwsSnsTopics) Get(name string, options v1.GetOptions) (result *aws_v1.AwsSnsTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssnstopicsResource, c.ns, name), &aws_v1.AwsSnsTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSnsTopic), err
}

// List takes label and field selectors, and returns the list of AwsSnsTopics that match those selectors.
func (c *FakeAwsSnsTopics) List(opts v1.ListOptions) (result *aws_v1.AwsSnsTopicList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssnstopicsResource, awssnstopicsKind, c.ns, opts), &aws_v1.AwsSnsTopicList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsSnsTopicList{}
	for _, item := range obj.(*aws_v1.AwsSnsTopicList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSnsTopics.
func (c *FakeAwsSnsTopics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssnstopicsResource, c.ns, opts))

}

// Create takes the representation of a awsSnsTopic and creates it.  Returns the server's representation of the awsSnsTopic, and an error, if there is any.
func (c *FakeAwsSnsTopics) Create(awsSnsTopic *aws_v1.AwsSnsTopic) (result *aws_v1.AwsSnsTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssnstopicsResource, c.ns, awsSnsTopic), &aws_v1.AwsSnsTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSnsTopic), err
}

// Update takes the representation of a awsSnsTopic and updates it. Returns the server's representation of the awsSnsTopic, and an error, if there is any.
func (c *FakeAwsSnsTopics) Update(awsSnsTopic *aws_v1.AwsSnsTopic) (result *aws_v1.AwsSnsTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssnstopicsResource, c.ns, awsSnsTopic), &aws_v1.AwsSnsTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSnsTopic), err
}

// Delete takes name of the awsSnsTopic and deletes it. Returns an error if one occurs.
func (c *FakeAwsSnsTopics) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssnstopicsResource, c.ns, name), &aws_v1.AwsSnsTopic{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSnsTopics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssnstopicsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsSnsTopicList{})
	return err
}

// Patch applies the patch and returns the patched awsSnsTopic.
func (c *FakeAwsSnsTopics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsSnsTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssnstopicsResource, c.ns, name, data, subresources...), &aws_v1.AwsSnsTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSnsTopic), err
}