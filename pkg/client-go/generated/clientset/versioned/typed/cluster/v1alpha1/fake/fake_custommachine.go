/*
Copyright Kurator Authors.

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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kurator.dev/kurator/pkg/apis/cluster/v1alpha1"
)

// FakeCustomMachines implements CustomMachineInterface
type FakeCustomMachines struct {
	Fake *FakeClusterV1alpha1
	ns   string
}

var custommachinesResource = schema.GroupVersionResource{Group: "cluster.kurator.dev", Version: "v1alpha1", Resource: "custommachines"}

var custommachinesKind = schema.GroupVersionKind{Group: "cluster.kurator.dev", Version: "v1alpha1", Kind: "CustomMachine"}

// Get takes name of the customMachine, and returns the corresponding customMachine object, and an error if there is any.
func (c *FakeCustomMachines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CustomMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(custommachinesResource, c.ns, name), &v1alpha1.CustomMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomMachine), err
}

// List takes label and field selectors, and returns the list of CustomMachines that match those selectors.
func (c *FakeCustomMachines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CustomMachineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(custommachinesResource, custommachinesKind, c.ns, opts), &v1alpha1.CustomMachineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CustomMachineList{ListMeta: obj.(*v1alpha1.CustomMachineList).ListMeta}
	for _, item := range obj.(*v1alpha1.CustomMachineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested customMachines.
func (c *FakeCustomMachines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(custommachinesResource, c.ns, opts))

}

// Create takes the representation of a customMachine and creates it.  Returns the server's representation of the customMachine, and an error, if there is any.
func (c *FakeCustomMachines) Create(ctx context.Context, customMachine *v1alpha1.CustomMachine, opts v1.CreateOptions) (result *v1alpha1.CustomMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(custommachinesResource, c.ns, customMachine), &v1alpha1.CustomMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomMachine), err
}

// Update takes the representation of a customMachine and updates it. Returns the server's representation of the customMachine, and an error, if there is any.
func (c *FakeCustomMachines) Update(ctx context.Context, customMachine *v1alpha1.CustomMachine, opts v1.UpdateOptions) (result *v1alpha1.CustomMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(custommachinesResource, c.ns, customMachine), &v1alpha1.CustomMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomMachine), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCustomMachines) UpdateStatus(ctx context.Context, customMachine *v1alpha1.CustomMachine, opts v1.UpdateOptions) (*v1alpha1.CustomMachine, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(custommachinesResource, "status", c.ns, customMachine), &v1alpha1.CustomMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomMachine), err
}

// Delete takes name of the customMachine and deletes it. Returns an error if one occurs.
func (c *FakeCustomMachines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(custommachinesResource, c.ns, name, opts), &v1alpha1.CustomMachine{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCustomMachines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(custommachinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CustomMachineList{})
	return err
}

// Patch applies the patch and returns the patched customMachine.
func (c *FakeCustomMachines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CustomMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(custommachinesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CustomMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomMachine), err
}
