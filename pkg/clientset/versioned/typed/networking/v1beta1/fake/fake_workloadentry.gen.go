// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWorkloadEntries implements WorkloadEntryInterface
type FakeWorkloadEntries struct {
	Fake *FakeNetworkingV1beta1
	ns   string
}

var workloadentriesResource = schema.GroupVersionResource{Group: "networking.istio.io", Version: "v1beta1", Resource: "workloadentries"}

var workloadentriesKind = schema.GroupVersionKind{Group: "networking.istio.io", Version: "v1beta1", Kind: "WorkloadEntry"}

// Get takes name of the workloadEntry, and returns the corresponding workloadEntry object, and an error if there is any.
func (c *FakeWorkloadEntries) Get(name string, options v1.GetOptions) (result *v1beta1.WorkloadEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(workloadentriesResource, c.ns, name), &v1beta1.WorkloadEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.WorkloadEntry), err
}

// List takes label and field selectors, and returns the list of WorkloadEntries that match those selectors.
func (c *FakeWorkloadEntries) List(opts v1.ListOptions) (result *v1beta1.WorkloadEntryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(workloadentriesResource, workloadentriesKind, c.ns, opts), &v1beta1.WorkloadEntryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.WorkloadEntryList{ListMeta: obj.(*v1beta1.WorkloadEntryList).ListMeta}
	for _, item := range obj.(*v1beta1.WorkloadEntryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workloadEntries.
func (c *FakeWorkloadEntries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(workloadentriesResource, c.ns, opts))

}

// Create takes the representation of a workloadEntry and creates it.  Returns the server's representation of the workloadEntry, and an error, if there is any.
func (c *FakeWorkloadEntries) Create(workloadEntry *v1beta1.WorkloadEntry) (result *v1beta1.WorkloadEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(workloadentriesResource, c.ns, workloadEntry), &v1beta1.WorkloadEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.WorkloadEntry), err
}

// Update takes the representation of a workloadEntry and updates it. Returns the server's representation of the workloadEntry, and an error, if there is any.
func (c *FakeWorkloadEntries) Update(workloadEntry *v1beta1.WorkloadEntry) (result *v1beta1.WorkloadEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(workloadentriesResource, c.ns, workloadEntry), &v1beta1.WorkloadEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.WorkloadEntry), err
}

// Delete takes name of the workloadEntry and deletes it. Returns an error if one occurs.
func (c *FakeWorkloadEntries) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(workloadentriesResource, c.ns, name), &v1beta1.WorkloadEntry{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkloadEntries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(workloadentriesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.WorkloadEntryList{})
	return err
}

// Patch applies the patch and returns the patched workloadEntry.
func (c *FakeWorkloadEntries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.WorkloadEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(workloadentriesResource, c.ns, name, pt, data, subresources...), &v1beta1.WorkloadEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.WorkloadEntry), err
}