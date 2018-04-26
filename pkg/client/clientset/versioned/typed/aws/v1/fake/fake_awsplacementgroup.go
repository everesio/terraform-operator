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

// FakeAwsPlacementGroups implements AwsPlacementGroupInterface
type FakeAwsPlacementGroups struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsplacementgroupsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsplacementgroups"}

var awsplacementgroupsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsPlacementGroup"}

// Get takes name of the awsPlacementGroup, and returns the corresponding awsPlacementGroup object, and an error if there is any.
func (c *FakeAwsPlacementGroups) Get(name string, options v1.GetOptions) (result *aws_v1.AwsPlacementGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsplacementgroupsResource, c.ns, name), &aws_v1.AwsPlacementGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsPlacementGroup), err
}

// List takes label and field selectors, and returns the list of AwsPlacementGroups that match those selectors.
func (c *FakeAwsPlacementGroups) List(opts v1.ListOptions) (result *aws_v1.AwsPlacementGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsplacementgroupsResource, awsplacementgroupsKind, c.ns, opts), &aws_v1.AwsPlacementGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsPlacementGroupList{}
	for _, item := range obj.(*aws_v1.AwsPlacementGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsPlacementGroups.
func (c *FakeAwsPlacementGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsplacementgroupsResource, c.ns, opts))

}

// Create takes the representation of a awsPlacementGroup and creates it.  Returns the server's representation of the awsPlacementGroup, and an error, if there is any.
func (c *FakeAwsPlacementGroups) Create(awsPlacementGroup *aws_v1.AwsPlacementGroup) (result *aws_v1.AwsPlacementGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsplacementgroupsResource, c.ns, awsPlacementGroup), &aws_v1.AwsPlacementGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsPlacementGroup), err
}

// Update takes the representation of a awsPlacementGroup and updates it. Returns the server's representation of the awsPlacementGroup, and an error, if there is any.
func (c *FakeAwsPlacementGroups) Update(awsPlacementGroup *aws_v1.AwsPlacementGroup) (result *aws_v1.AwsPlacementGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsplacementgroupsResource, c.ns, awsPlacementGroup), &aws_v1.AwsPlacementGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsPlacementGroup), err
}

// Delete takes name of the awsPlacementGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsPlacementGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsplacementgroupsResource, c.ns, name), &aws_v1.AwsPlacementGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsPlacementGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsplacementgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsPlacementGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsPlacementGroup.
func (c *FakeAwsPlacementGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsPlacementGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsplacementgroupsResource, c.ns, name, data, subresources...), &aws_v1.AwsPlacementGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsPlacementGroup), err
}