package meta

import (
	k8sMetaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sWatchV1 "k8s.io/apimachinery/pkg/watch"
	k8sInformers "k8s.io/client-go/informers"
	k8sInformersAppsV1 "k8s.io/client-go/informers/apps/v1"
	k8sInformersCoreV1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/metadata/metadatainformer"
	"k8s.io/client-go/tools/cache"
	crdInformers "k8s.tars.io/client-go/informers/externalversions"
	crdInformersV1alpha2 "k8s.tars.io/client-go/informers/externalversions/crd/v1alpha2"
)

type EventsReceiver interface {
	EnqueueObj(resourceName string, resourceEvent k8sWatchV1.EventType, resourceObj interface{})
}

type Informers struct {
	K8sInformerFactory        k8sInformers.SharedInformerFactory
	K8sMetadataInformerFactor metadatainformer.SharedInformerFactory
	CrdInformerFactory        crdInformers.SharedInformerFactory

	synced  bool
	synceds []cache.InformerSynced

	ServiceInformer k8sInformersCoreV1.ServiceInformer
	PodInformer     k8sInformersCoreV1.PodInformer

	DaemonSetInformer   k8sInformersAppsV1.DaemonSetInformer
	StatefulSetInformer k8sInformersAppsV1.StatefulSetInformer

	TServerInformer       crdInformersV1alpha2.TServerInformer
	TEndpointInformer     crdInformersV1alpha2.TEndpointInformer
	TTemplateInformer     crdInformersV1alpha2.TTemplateInformer
	TImageInformer        crdInformersV1alpha2.TImageInformer
	TTreeInformer         crdInformersV1alpha2.TTreeInformer
	TExitedRecordInformer crdInformersV1alpha2.TExitedRecordInformer
	TDeployInformer       crdInformersV1alpha2.TDeployInformer
	TAccountInformer      crdInformersV1alpha2.TAccountInformer

	TConfigInformer k8sInformers.GenericInformer

	receivers []EventsReceiver
}

func (i *Informers) send(name string, event k8sWatchV1.EventType, obj interface{}) {
	for _, r := range i.receivers {
		r.EnqueueObj(name, event, obj)
	}
}

func (i *Informers) Register(r EventsReceiver) {
	i.receivers = append(i.receivers, r)
}

func setEventHandler(resourceKind string, resourceInformer cache.SharedInformer, informers *Informers) {
	eventHandler := cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			informers.send(resourceKind, k8sWatchV1.Added, obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldMeta := oldObj.(k8sMetaV1.Object)
			newMeta := newObj.(k8sMetaV1.Object)
			if newMeta.GetResourceVersion() != oldMeta.GetResourceVersion() {
				informers.send(resourceKind, k8sWatchV1.Modified, newObj)
			}
		},
		DeleteFunc: func(obj interface{}) {
			informers.send(resourceKind, k8sWatchV1.Deleted, obj)
		},
	}
	resourceInformer.AddEventHandler(eventHandler)
}

func (i *Informers) Start(stop chan struct{}) {
	i.K8sInformerFactory.Start(stop)
	i.K8sMetadataInformerFactor.Start(stop)
	i.CrdInformerFactory.Start(stop)
}

func (i *Informers) Synced() bool {
	return i.synced
}

func (i *Informers) WaitForCacheSync(stopCh chan struct{}) bool {
	i.synced = cache.WaitForCacheSync(stopCh, i.synceds...)
	return i.synced
}
