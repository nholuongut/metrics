/*
Copyright Nho Luong DevOps.

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
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

// FakePodMetricses implements PodMetricsInterface
type FakePodMetricses struct {
	Fake *FakeMetricsV1beta1
	ns   string
}

var podmetricsesResource = v1beta1.SchemeGroupVersion.WithResource("pods")

var podmetricsesKind = v1beta1.SchemeGroupVersion.WithKind("PodMetrics")

// Get takes name of the podMetrics, and returns the corresponding podMetrics object, and an error if there is any.
func (c *FakePodMetricses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.PodMetrics, err error) {
	emptyResult := &v1beta1.PodMetrics{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(podmetricsesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.PodMetrics), err
}

// List takes label and field selectors, and returns the list of PodMetricses that match those selectors.
func (c *FakePodMetricses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PodMetricsList, err error) {
	emptyResult := &v1beta1.PodMetricsList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(podmetricsesResource, podmetricsesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PodMetricsList{ListMeta: obj.(*v1beta1.PodMetricsList).ListMeta}
	for _, item := range obj.(*v1beta1.PodMetricsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podMetricses.
func (c *FakePodMetricses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(podmetricsesResource, c.ns, opts))

}