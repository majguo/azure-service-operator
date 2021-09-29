// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=microsoft.network.azure.com,resources=networksecuritygroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.network.azure.com,resources={networksecuritygroups/status,networksecuritygroups/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201101.NetworkSecurityGroup
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/networkSecurityGroups
type NetworkSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSecurityGroups_Spec                                           `json:"spec,omitempty"`
	Status            NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NetworkSecurityGroup{}

// GetConditions returns the conditions of the resource
func (networkSecurityGroup *NetworkSecurityGroup) GetConditions() conditions.Conditions {
	return networkSecurityGroup.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (networkSecurityGroup *NetworkSecurityGroup) SetConditions(conditions conditions.Conditions) {
	networkSecurityGroup.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &NetworkSecurityGroup{}

// AzureName returns the Azure name of the resource
func (networkSecurityGroup *NetworkSecurityGroup) AzureName() string {
	return networkSecurityGroup.Spec.AzureName
}

// GetResourceKind returns the kind of the resource
func (networkSecurityGroup *NetworkSecurityGroup) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (networkSecurityGroup *NetworkSecurityGroup) GetSpec() genruntime.ConvertibleSpec {
	return &networkSecurityGroup.Spec
}

// GetStatus returns the status of this resource
func (networkSecurityGroup *NetworkSecurityGroup) GetStatus() genruntime.ConvertibleStatus {
	return &networkSecurityGroup.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/networkSecurityGroups"
func (networkSecurityGroup *NetworkSecurityGroup) GetType() string {
	return "Microsoft.Network/networkSecurityGroups"
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (networkSecurityGroup *NetworkSecurityGroup) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(networkSecurityGroup.Spec)
	return &genruntime.ResourceReference{Group: group, Kind: kind, Namespace: networkSecurityGroup.Namespace, Name: networkSecurityGroup.Spec.Owner.Name}
}

// SetStatus sets the status of this resource
func (networkSecurityGroup *NetworkSecurityGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded); ok {
		networkSecurityGroup.Status = *st
		return nil
	}

	// Convert status to required version
	var st NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	networkSecurityGroup.Status = st
	return nil
}

// Hub marks that this NetworkSecurityGroup is the hub type for conversion
func (networkSecurityGroup *NetworkSecurityGroup) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (networkSecurityGroup *NetworkSecurityGroup) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: networkSecurityGroup.Spec.OriginalVersion,
		Kind:    "NetworkSecurityGroup",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201101.NetworkSecurityGroup
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/networkSecurityGroups
type NetworkSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityGroup `json:"items"`
}

//Storage version of v1alpha1api20201101.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
//Generated from:
type NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded struct {
	Conditions           []conditions.Condition                                             `json:"conditions,omitempty"`
	DefaultSecurityRules []SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded     `json:"defaultSecurityRules,omitempty"`
	Etag                 *string                                                            `json:"etag,omitempty"`
	FlowLogs             []FlowLog_Status_SubResourceEmbedded                               `json:"flowLogs,omitempty"`
	Id                   *string                                                            `json:"id,omitempty"`
	Location             *string                                                            `json:"location,omitempty"`
	Name                 *string                                                            `json:"name,omitempty"`
	NetworkInterfaces    []NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded `json:"networkInterfaces,omitempty"`
	PropertyBag          genruntime.PropertyBag                                             `json:"$propertyBag,omitempty"`
	ProvisioningState    *string                                                            `json:"provisioningState,omitempty"`
	ResourceGuid         *string                                                            `json:"resourceGuid,omitempty"`
	SecurityRules        []SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded     `json:"securityRules,omitempty"`
	Subnets              []Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded           `json:"subnets,omitempty"`
	Tags                 map[string]string                                                  `json:"tags,omitempty"`
	Type                 *string                                                            `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded{}

// ConvertStatusFrom populates our NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded from the provided source
func (networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded *NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded)
}

// ConvertStatusTo populates the provided destination from our NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
func (networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded *NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded)
}

//Storage version of v1alpha1api20201101.NetworkSecurityGroups_Spec
type NetworkSecurityGroups_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string  `json:"azureName"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"microsoft.resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NetworkSecurityGroups_Spec{}

// ConvertSpecFrom populates our NetworkSecurityGroups_Spec from the provided source
func (networkSecurityGroupsSpec *NetworkSecurityGroups_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == networkSecurityGroupsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(networkSecurityGroupsSpec)
}

// ConvertSpecTo populates the provided destination from our NetworkSecurityGroups_Spec
func (networkSecurityGroupsSpec *NetworkSecurityGroups_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == networkSecurityGroupsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(networkSecurityGroupsSpec)
}

//Storage version of v1alpha1api20201101.FlowLog_Status_SubResourceEmbedded
//Generated from:
type FlowLog_Status_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded
//Generated from:
type NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded struct {
	ExtendedLocation *ExtendedLocation_Status `json:"extendedLocation,omitempty"`
	Id               *string                  `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded
//Generated from:
type SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded
//Generated from:
type Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityGroup{}, &NetworkSecurityGroupList{})
}
