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

// FakeAwsSfnActivities implements AwsSfnActivityInterface
type FakeAwsSfnActivities struct {
	Fake *FakeTrussleV1
	ns   string
}

var awssfnactivitiesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awssfnactivities"}

var awssfnactivitiesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsSfnActivity"}

// Get takes name of the awsSfnActivity, and returns the corresponding awsSfnActivity object, and an error if there is any.
func (c *FakeAwsSfnActivities) Get(name string, options v1.GetOptions) (result *aws_v1.AwsSfnActivity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssfnactivitiesResource, c.ns, name), &aws_v1.AwsSfnActivity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSfnActivity), err
}

// List takes label and field selectors, and returns the list of AwsSfnActivities that match those selectors.
func (c *FakeAwsSfnActivities) List(opts v1.ListOptions) (result *aws_v1.AwsSfnActivityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssfnactivitiesResource, awssfnactivitiesKind, c.ns, opts), &aws_v1.AwsSfnActivityList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsSfnActivityList{}
	for _, item := range obj.(*aws_v1.AwsSfnActivityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSfnActivities.
func (c *FakeAwsSfnActivities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssfnactivitiesResource, c.ns, opts))

}

// Create takes the representation of a awsSfnActivity and creates it.  Returns the server's representation of the awsSfnActivity, and an error, if there is any.
func (c *FakeAwsSfnActivities) Create(awsSfnActivity *aws_v1.AwsSfnActivity) (result *aws_v1.AwsSfnActivity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssfnactivitiesResource, c.ns, awsSfnActivity), &aws_v1.AwsSfnActivity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSfnActivity), err
}

// Update takes the representation of a awsSfnActivity and updates it. Returns the server's representation of the awsSfnActivity, and an error, if there is any.
func (c *FakeAwsSfnActivities) Update(awsSfnActivity *aws_v1.AwsSfnActivity) (result *aws_v1.AwsSfnActivity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssfnactivitiesResource, c.ns, awsSfnActivity), &aws_v1.AwsSfnActivity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSfnActivity), err
}

// Delete takes name of the awsSfnActivity and deletes it. Returns an error if one occurs.
func (c *FakeAwsSfnActivities) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssfnactivitiesResource, c.ns, name), &aws_v1.AwsSfnActivity{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSfnActivities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssfnactivitiesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsSfnActivityList{})
	return err
}

// Patch applies the patch and returns the patched awsSfnActivity.
func (c *FakeAwsSfnActivities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsSfnActivity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssfnactivitiesResource, c.ns, name, data, subresources...), &aws_v1.AwsSfnActivity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSfnActivity), err
}