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

// FakeAwsIotPolicies implements AwsIotPolicyInterface
type FakeAwsIotPolicies struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsiotpoliciesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsiotpolicies"}

var awsiotpoliciesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsIotPolicy"}

// Get takes name of the awsIotPolicy, and returns the corresponding awsIotPolicy object, and an error if there is any.
func (c *FakeAwsIotPolicies) Get(name string, options v1.GetOptions) (result *aws_v1.AwsIotPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsiotpoliciesResource, c.ns, name), &aws_v1.AwsIotPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotPolicy), err
}

// List takes label and field selectors, and returns the list of AwsIotPolicies that match those selectors.
func (c *FakeAwsIotPolicies) List(opts v1.ListOptions) (result *aws_v1.AwsIotPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsiotpoliciesResource, awsiotpoliciesKind, c.ns, opts), &aws_v1.AwsIotPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsIotPolicyList{}
	for _, item := range obj.(*aws_v1.AwsIotPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIotPolicies.
func (c *FakeAwsIotPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsiotpoliciesResource, c.ns, opts))

}

// Create takes the representation of a awsIotPolicy and creates it.  Returns the server's representation of the awsIotPolicy, and an error, if there is any.
func (c *FakeAwsIotPolicies) Create(awsIotPolicy *aws_v1.AwsIotPolicy) (result *aws_v1.AwsIotPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsiotpoliciesResource, c.ns, awsIotPolicy), &aws_v1.AwsIotPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotPolicy), err
}

// Update takes the representation of a awsIotPolicy and updates it. Returns the server's representation of the awsIotPolicy, and an error, if there is any.
func (c *FakeAwsIotPolicies) Update(awsIotPolicy *aws_v1.AwsIotPolicy) (result *aws_v1.AwsIotPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsiotpoliciesResource, c.ns, awsIotPolicy), &aws_v1.AwsIotPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotPolicy), err
}

// Delete takes name of the awsIotPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAwsIotPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsiotpoliciesResource, c.ns, name), &aws_v1.AwsIotPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIotPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsiotpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsIotPolicyList{})
	return err
}

// Patch applies the patch and returns the patched awsIotPolicy.
func (c *FakeAwsIotPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsIotPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsiotpoliciesResource, c.ns, name, data, subresources...), &aws_v1.AwsIotPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotPolicy), err
}