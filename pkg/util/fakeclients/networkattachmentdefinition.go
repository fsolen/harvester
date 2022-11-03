/*
Copyright 2022 Rancher Labs, Inc.

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

package fakeclients

import (
	"context"

	"k8s.io/apimachinery/pkg/labels"

	cnitype "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/k8s.cni.cncf.io/v1"
	ctlcniv1 "github.com/harvester/harvester/pkg/generated/controllers/k8s.cni.cncf.io/v1"
	cniv1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NetworkAttachmentDefinitionCache func(namespace string) cnitype.NetworkAttachmentDefinitionInterface

func (c NetworkAttachmentDefinitionCache) Get(namespace, name string) (*cniv1.NetworkAttachmentDefinition, error) {
	return c(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

func (c NetworkAttachmentDefinitionCache) List(nemspace string, selector labels.Selector) ([]*cniv1.NetworkAttachmentDefinition, error) {
	panic("implement me")
}

func (c NetworkAttachmentDefinitionCache) AddIndexer(indexName string, indexer ctlcniv1.NetworkAttachmentDefinitionIndexer) {
	panic("implement me")
}
func (c NetworkAttachmentDefinitionCache) GetByIndex(indexName, key string) ([]*cniv1.NetworkAttachmentDefinition, error) {
	panic("implement me")
}