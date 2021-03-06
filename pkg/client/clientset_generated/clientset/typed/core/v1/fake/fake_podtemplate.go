/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
	api "k8s.io/kubernetes/pkg/api"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

// FakePodTemplates implements PodTemplateInterface
type FakePodTemplates struct {
	Fake *FakeCoreV1
	ns   string
}

var podtemplatesResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "podtemplates"}

func (c *FakePodTemplates) Create(podTemplate *v1.PodTemplate) (result *v1.PodTemplate, err error) {
	obj, err := c.Fake.
		Invokes(core.NewCreateAction(podtemplatesResource, c.ns, podTemplate), &v1.PodTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PodTemplate), err
}

func (c *FakePodTemplates) Update(podTemplate *v1.PodTemplate) (result *v1.PodTemplate, err error) {
	obj, err := c.Fake.
		Invokes(core.NewUpdateAction(podtemplatesResource, c.ns, podTemplate), &v1.PodTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PodTemplate), err
}

func (c *FakePodTemplates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(core.NewDeleteAction(podtemplatesResource, c.ns, name), &v1.PodTemplate{})

	return err
}

func (c *FakePodTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := core.NewDeleteCollectionAction(podtemplatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1.PodTemplateList{})
	return err
}

func (c *FakePodTemplates) Get(name string, options meta_v1.GetOptions) (result *v1.PodTemplate, err error) {
	obj, err := c.Fake.
		Invokes(core.NewGetAction(podtemplatesResource, c.ns, name), &v1.PodTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PodTemplate), err
}

func (c *FakePodTemplates) List(opts v1.ListOptions) (result *v1.PodTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(core.NewListAction(podtemplatesResource, c.ns, opts), &v1.PodTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := core.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.PodTemplateList{}
	for _, item := range obj.(*v1.PodTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podTemplates.
func (c *FakePodTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(core.NewWatchAction(podtemplatesResource, c.ns, opts))

}

// Patch applies the patch and returns the patched podTemplate.
func (c *FakePodTemplates) Patch(name string, pt api.PatchType, data []byte, subresources ...string) (result *v1.PodTemplate, err error) {
	obj, err := c.Fake.
		Invokes(core.NewPatchSubresourceAction(podtemplatesResource, c.ns, name, data, subresources...), &v1.PodTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PodTemplate), err
}
