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

// FakeAwsInspectorAssessmentTemplates implements AwsInspectorAssessmentTemplateInterface
type FakeAwsInspectorAssessmentTemplates struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsinspectorassessmenttemplatesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsinspectorassessmenttemplates"}

var awsinspectorassessmenttemplatesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsInspectorAssessmentTemplate"}

// Get takes name of the awsInspectorAssessmentTemplate, and returns the corresponding awsInspectorAssessmentTemplate object, and an error if there is any.
func (c *FakeAwsInspectorAssessmentTemplates) Get(name string, options v1.GetOptions) (result *aws_v1.AwsInspectorAssessmentTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsinspectorassessmenttemplatesResource, c.ns, name), &aws_v1.AwsInspectorAssessmentTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsInspectorAssessmentTemplate), err
}

// List takes label and field selectors, and returns the list of AwsInspectorAssessmentTemplates that match those selectors.
func (c *FakeAwsInspectorAssessmentTemplates) List(opts v1.ListOptions) (result *aws_v1.AwsInspectorAssessmentTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsinspectorassessmenttemplatesResource, awsinspectorassessmenttemplatesKind, c.ns, opts), &aws_v1.AwsInspectorAssessmentTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsInspectorAssessmentTemplateList{}
	for _, item := range obj.(*aws_v1.AwsInspectorAssessmentTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsInspectorAssessmentTemplates.
func (c *FakeAwsInspectorAssessmentTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsinspectorassessmenttemplatesResource, c.ns, opts))

}

// Create takes the representation of a awsInspectorAssessmentTemplate and creates it.  Returns the server's representation of the awsInspectorAssessmentTemplate, and an error, if there is any.
func (c *FakeAwsInspectorAssessmentTemplates) Create(awsInspectorAssessmentTemplate *aws_v1.AwsInspectorAssessmentTemplate) (result *aws_v1.AwsInspectorAssessmentTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsinspectorassessmenttemplatesResource, c.ns, awsInspectorAssessmentTemplate), &aws_v1.AwsInspectorAssessmentTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsInspectorAssessmentTemplate), err
}

// Update takes the representation of a awsInspectorAssessmentTemplate and updates it. Returns the server's representation of the awsInspectorAssessmentTemplate, and an error, if there is any.
func (c *FakeAwsInspectorAssessmentTemplates) Update(awsInspectorAssessmentTemplate *aws_v1.AwsInspectorAssessmentTemplate) (result *aws_v1.AwsInspectorAssessmentTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsinspectorassessmenttemplatesResource, c.ns, awsInspectorAssessmentTemplate), &aws_v1.AwsInspectorAssessmentTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsInspectorAssessmentTemplate), err
}

// Delete takes name of the awsInspectorAssessmentTemplate and deletes it. Returns an error if one occurs.
func (c *FakeAwsInspectorAssessmentTemplates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsinspectorassessmenttemplatesResource, c.ns, name), &aws_v1.AwsInspectorAssessmentTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsInspectorAssessmentTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsinspectorassessmenttemplatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsInspectorAssessmentTemplateList{})
	return err
}

// Patch applies the patch and returns the patched awsInspectorAssessmentTemplate.
func (c *FakeAwsInspectorAssessmentTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsInspectorAssessmentTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsinspectorassessmenttemplatesResource, c.ns, name, data, subresources...), &aws_v1.AwsInspectorAssessmentTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsInspectorAssessmentTemplate), err
}
