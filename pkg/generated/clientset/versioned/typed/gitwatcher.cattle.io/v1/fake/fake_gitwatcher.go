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

// Code generated by main. DO NOT EDIT.

package fake

import (
	gitwatchercattleiov1 "github.com/rancher/gitwatcher/pkg/apis/gitwatcher.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGitWatchers implements GitWatcherInterface
type FakeGitWatchers struct {
	Fake *FakeGitwatcherV1
	ns   string
}

var gitwatchersResource = schema.GroupVersionResource{Group: "gitwatcher.cattle.io", Version: "v1", Resource: "gitwatchers"}

var gitwatchersKind = schema.GroupVersionKind{Group: "gitwatcher.cattle.io", Version: "v1", Kind: "GitWatcher"}

// Get takes name of the gitWatcher, and returns the corresponding gitWatcher object, and an error if there is any.
func (c *FakeGitWatchers) Get(name string, options v1.GetOptions) (result *gitwatchercattleiov1.GitWatcher, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gitwatchersResource, c.ns, name), &gitwatchercattleiov1.GitWatcher{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gitwatchercattleiov1.GitWatcher), err
}

// List takes label and field selectors, and returns the list of GitWatchers that match those selectors.
func (c *FakeGitWatchers) List(opts v1.ListOptions) (result *gitwatchercattleiov1.GitWatcherList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gitwatchersResource, gitwatchersKind, c.ns, opts), &gitwatchercattleiov1.GitWatcherList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &gitwatchercattleiov1.GitWatcherList{ListMeta: obj.(*gitwatchercattleiov1.GitWatcherList).ListMeta}
	for _, item := range obj.(*gitwatchercattleiov1.GitWatcherList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gitWatchers.
func (c *FakeGitWatchers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gitwatchersResource, c.ns, opts))

}

// Create takes the representation of a gitWatcher and creates it.  Returns the server's representation of the gitWatcher, and an error, if there is any.
func (c *FakeGitWatchers) Create(gitWatcher *gitwatchercattleiov1.GitWatcher) (result *gitwatchercattleiov1.GitWatcher, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gitwatchersResource, c.ns, gitWatcher), &gitwatchercattleiov1.GitWatcher{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gitwatchercattleiov1.GitWatcher), err
}

// Update takes the representation of a gitWatcher and updates it. Returns the server's representation of the gitWatcher, and an error, if there is any.
func (c *FakeGitWatchers) Update(gitWatcher *gitwatchercattleiov1.GitWatcher) (result *gitwatchercattleiov1.GitWatcher, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gitwatchersResource, c.ns, gitWatcher), &gitwatchercattleiov1.GitWatcher{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gitwatchercattleiov1.GitWatcher), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGitWatchers) UpdateStatus(gitWatcher *gitwatchercattleiov1.GitWatcher) (*gitwatchercattleiov1.GitWatcher, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gitwatchersResource, "status", c.ns, gitWatcher), &gitwatchercattleiov1.GitWatcher{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gitwatchercattleiov1.GitWatcher), err
}

// Delete takes name of the gitWatcher and deletes it. Returns an error if one occurs.
func (c *FakeGitWatchers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gitwatchersResource, c.ns, name), &gitwatchercattleiov1.GitWatcher{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGitWatchers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gitwatchersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &gitwatchercattleiov1.GitWatcherList{})
	return err
}

// Patch applies the patch and returns the patched gitWatcher.
func (c *FakeGitWatchers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *gitwatchercattleiov1.GitWatcher, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gitwatchersResource, c.ns, name, pt, data, subresources...), &gitwatchercattleiov1.GitWatcher{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gitwatchercattleiov1.GitWatcher), err
}
