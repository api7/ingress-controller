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

// Code generated by informer-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	time "time"

	configv2alpha1 "github.com/apache/apisix-ingress-controller/pkg/kube/apisix/apis/config/v2alpha1"
	versioned "github.com/apache/apisix-ingress-controller/pkg/kube/apisix/client/clientset/versioned"
	internalinterfaces "github.com/apache/apisix-ingress-controller/pkg/kube/apisix/client/informers/externalversions/internalinterfaces"
	v2alpha1 "github.com/apache/apisix-ingress-controller/pkg/kube/apisix/client/listers/config/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ApisixClusterConfigInformer provides access to a shared informer and lister for
// ApisixClusterConfigs.
type ApisixClusterConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2alpha1.ApisixClusterConfigLister
}

type apisixClusterConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewApisixClusterConfigInformer constructs a new informer for ApisixClusterConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewApisixClusterConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredApisixClusterConfigInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredApisixClusterConfigInformer constructs a new informer for ApisixClusterConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredApisixClusterConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApisixV2alpha1().ApisixClusterConfigs().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApisixV2alpha1().ApisixClusterConfigs().Watch(context.TODO(), options)
			},
		},
		&configv2alpha1.ApisixClusterConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *apisixClusterConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredApisixClusterConfigInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *apisixClusterConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configv2alpha1.ApisixClusterConfig{}, f.defaultInformer)
}

func (f *apisixClusterConfigInformer) Lister() v2alpha1.ApisixClusterConfigLister {
	return v2alpha1.NewApisixClusterConfigLister(f.Informer().GetIndexer())
}
