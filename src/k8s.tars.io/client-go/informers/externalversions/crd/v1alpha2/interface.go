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

package v1alpha2

import (
	internalinterfaces "k8s.tars.io/client-go/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// TAccounts returns a TAccountInformer.
	TAccounts() TAccountInformer
	// TConfigs returns a TConfigInformer.
	TConfigs() TConfigInformer
	// TDeploys returns a TDeployInformer.
	TDeploys() TDeployInformer
	// TEndpoints returns a TEndpointInformer.
	TEndpoints() TEndpointInformer
	// TExitedRecords returns a TExitedRecordInformer.
	TExitedRecords() TExitedRecordInformer
	// TImages returns a TImageInformer.
	TImages() TImageInformer
	// TServers returns a TServerInformer.
	TServers() TServerInformer
	// TTemplates returns a TTemplateInformer.
	TTemplates() TTemplateInformer
	// TTrees returns a TTreeInformer.
	TTrees() TTreeInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// TAccounts returns a TAccountInformer.
func (v *version) TAccounts() TAccountInformer {
	return &tAccountInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TConfigs returns a TConfigInformer.
func (v *version) TConfigs() TConfigInformer {
	return &tConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TDeploys returns a TDeployInformer.
func (v *version) TDeploys() TDeployInformer {
	return &tDeployInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TEndpoints returns a TEndpointInformer.
func (v *version) TEndpoints() TEndpointInformer {
	return &tEndpointInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TExitedRecords returns a TExitedRecordInformer.
func (v *version) TExitedRecords() TExitedRecordInformer {
	return &tExitedRecordInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TImages returns a TImageInformer.
func (v *version) TImages() TImageInformer {
	return &tImageInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TServers returns a TServerInformer.
func (v *version) TServers() TServerInformer {
	return &tServerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TTemplates returns a TTemplateInformer.
func (v *version) TTemplates() TTemplateInformer {
	return &tTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TTrees returns a TTreeInformer.
func (v *version) TTrees() TTreeInformer {
	return &tTreeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
