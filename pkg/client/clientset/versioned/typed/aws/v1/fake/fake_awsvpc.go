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

// FakeAwsVpcs implements AwsVpcInterface
type FakeAwsVpcs struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsvpcsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsvpcs"}

var awsvpcsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsVpc"}

// Get takes name of the awsVpc, and returns the corresponding awsVpc object, and an error if there is any.
func (c *FakeAwsVpcs) Get(name string, options v1.GetOptions) (result *aws_v1.AwsVpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsvpcsResource, c.ns, name), &aws_v1.AwsVpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpc), err
}

// List takes label and field selectors, and returns the list of AwsVpcs that match those selectors.
func (c *FakeAwsVpcs) List(opts v1.ListOptions) (result *aws_v1.AwsVpcList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsvpcsResource, awsvpcsKind, c.ns, opts), &aws_v1.AwsVpcList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsVpcList{}
	for _, item := range obj.(*aws_v1.AwsVpcList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsVpcs.
func (c *FakeAwsVpcs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsvpcsResource, c.ns, opts))

}

// Create takes the representation of a awsVpc and creates it.  Returns the server's representation of the awsVpc, and an error, if there is any.
func (c *FakeAwsVpcs) Create(awsVpc *aws_v1.AwsVpc) (result *aws_v1.AwsVpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsvpcsResource, c.ns, awsVpc), &aws_v1.AwsVpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpc), err
}

// Update takes the representation of a awsVpc and updates it. Returns the server's representation of the awsVpc, and an error, if there is any.
func (c *FakeAwsVpcs) Update(awsVpc *aws_v1.AwsVpc) (result *aws_v1.AwsVpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsvpcsResource, c.ns, awsVpc), &aws_v1.AwsVpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpc), err
}

// Delete takes name of the awsVpc and deletes it. Returns an error if one occurs.
func (c *FakeAwsVpcs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsvpcsResource, c.ns, name), &aws_v1.AwsVpc{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsVpcs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsvpcsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsVpcList{})
	return err
}

// Patch applies the patch and returns the patched awsVpc.
func (c *FakeAwsVpcs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsVpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsvpcsResource, c.ns, name, data, subresources...), &aws_v1.AwsVpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpc), err
}
