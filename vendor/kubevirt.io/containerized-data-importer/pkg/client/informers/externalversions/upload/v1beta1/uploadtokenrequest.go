/*
Copyright 2018 The CDI Authors.

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

package v1beta1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	uploadv1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/upload/v1beta1"
	versioned "kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned"
	internalinterfaces "kubevirt.io/containerized-data-importer/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "kubevirt.io/containerized-data-importer/pkg/client/listers/upload/v1beta1"
)

// UploadTokenRequestInformer provides access to a shared informer and lister for
// UploadTokenRequests.
type UploadTokenRequestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.UploadTokenRequestLister
}

type uploadTokenRequestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewUploadTokenRequestInformer constructs a new informer for UploadTokenRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewUploadTokenRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredUploadTokenRequestInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredUploadTokenRequestInformer constructs a new informer for UploadTokenRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredUploadTokenRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.UploadV1beta1().UploadTokenRequests(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.UploadV1beta1().UploadTokenRequests(namespace).Watch(context.TODO(), options)
			},
		},
		&uploadv1beta1.UploadTokenRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *uploadTokenRequestInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredUploadTokenRequestInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *uploadTokenRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&uploadv1beta1.UploadTokenRequest{}, f.defaultInformer)
}

func (f *uploadTokenRequestInformer) Lister() v1beta1.UploadTokenRequestLister {
	return v1beta1.NewUploadTokenRequestLister(f.Informer().GetIndexer())
}
