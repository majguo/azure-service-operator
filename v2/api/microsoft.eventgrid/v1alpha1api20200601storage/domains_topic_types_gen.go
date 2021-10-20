// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20200601.DomainsTopic
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/resourceDefinitions/domains_topics
type DomainsTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainsTopics_Spec `json:"spec,omitempty"`
	Status            DomainTopic_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DomainsTopic{}

// GetConditions returns the conditions of the resource
func (domainsTopic *DomainsTopic) GetConditions() conditions.Conditions {
	return domainsTopic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (domainsTopic *DomainsTopic) SetConditions(conditions conditions.Conditions) {
	domainsTopic.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &DomainsTopic{}

// AzureName returns the Azure name of the resource
func (domainsTopic *DomainsTopic) AzureName() string {
	return domainsTopic.Spec.AzureName
}

// GetResourceKind returns the kind of the resource
func (domainsTopic *DomainsTopic) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (domainsTopic *DomainsTopic) GetSpec() genruntime.ConvertibleSpec {
	return &domainsTopic.Spec
}

// GetStatus returns the status of this resource
func (domainsTopic *DomainsTopic) GetStatus() genruntime.ConvertibleStatus {
	return &domainsTopic.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains/topics"
func (domainsTopic *DomainsTopic) GetType() string {
	return "Microsoft.EventGrid/domains/topics"
}

// NewEmptyStatus returns a new empty (blank) status
func (domainsTopic *DomainsTopic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DomainTopic_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (domainsTopic *DomainsTopic) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(domainsTopic.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: domainsTopic.Namespace,
		Name:      domainsTopic.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (domainsTopic *DomainsTopic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DomainTopic_Status); ok {
		domainsTopic.Status = *st
		return nil
	}

	// Convert status to required version
	var st DomainTopic_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	domainsTopic.Status = st
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (domainsTopic *DomainsTopic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: domainsTopic.Spec.OriginalVersion,
		Kind:    "DomainsTopic",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20200601.DomainsTopic
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/resourceDefinitions/domains_topics
type DomainsTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainsTopic `json:"items"`
}

//Storage version of v1alpha1api20200601.DomainTopic_Status
//Generated from:
type DomainTopic_Status struct {
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	SystemData        *SystemData_Status     `json:"systemData,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DomainTopic_Status{}

// ConvertStatusFrom populates our DomainTopic_Status from the provided source
func (domainTopicStatus *DomainTopic_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == domainTopicStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(domainTopicStatus)
}

// ConvertStatusTo populates the provided destination from our DomainTopic_Status
func (domainTopicStatus *DomainTopic_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == domainTopicStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(domainTopicStatus)
}

//Storage version of v1alpha1api20200601.DomainsTopics_Spec
type DomainsTopics_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string  `json:"azureName"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"microsoft.eventgrid.azure.com" json:"owner" kind:"Domain"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DomainsTopics_Spec{}

// ConvertSpecFrom populates our DomainsTopics_Spec from the provided source
func (domainsTopicsSpec *DomainsTopics_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == domainsTopicsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(domainsTopicsSpec)
}

// ConvertSpecTo populates the provided destination from our DomainsTopics_Spec
func (domainsTopicsSpec *DomainsTopics_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == domainsTopicsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(domainsTopicsSpec)
}

func init() {
	SchemeBuilder.Register(&DomainsTopic{}, &DomainsTopicList{})
}
