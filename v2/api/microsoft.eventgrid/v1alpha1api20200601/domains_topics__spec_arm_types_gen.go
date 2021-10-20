// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DomainsTopics_SpecARM struct {
	//APIVersion: API Version of the resource type, optional when apiProfile is used
	//on the template
	APIVersion DomainsTopicsSpecAPIVersion `json:"apiVersion"`

	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	//Name: Name of the domain topic.
	Name string `json:"name"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Resource type
	Type DomainsTopicsSpecType `json:"type"`
}

var _ genruntime.ARMResourceSpec = &DomainsTopics_SpecARM{}

// GetAPIVersion returns the APIVersion of the resource
func (domainsTopicsSpecARM DomainsTopics_SpecARM) GetAPIVersion() string {
	return string(domainsTopicsSpecARM.APIVersion)
}

// GetName returns the Name of the resource
func (domainsTopicsSpecARM DomainsTopics_SpecARM) GetName() string {
	return domainsTopicsSpecARM.Name
}

// GetType returns the Type of the resource
func (domainsTopicsSpecARM DomainsTopics_SpecARM) GetType() string {
	return string(domainsTopicsSpecARM.Type)
}

// +kubebuilder:validation:Enum={"2020-06-01"}
type DomainsTopicsSpecAPIVersion string

const DomainsTopicsSpecAPIVersion20200601 = DomainsTopicsSpecAPIVersion("2020-06-01")

// +kubebuilder:validation:Enum={"Microsoft.EventGrid/domains/topics"}
type DomainsTopicsSpecType string

const DomainsTopicsSpecTypeMicrosoftEventGridDomainsTopics = DomainsTopicsSpecType("Microsoft.EventGrid/domains/topics")
