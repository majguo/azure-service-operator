// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210701storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=machinelearningservices.azure.com,resources=workspaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=machinelearningservices.azure.com,resources={workspaces/status,workspaces/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210701.Workspace
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/resourceDefinitions/workspaces
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Workspaces_Spec  `json:"spec,omitempty"`
	Status            Workspace_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Workspace{}

// GetConditions returns the conditions of the resource
func (workspace *Workspace) GetConditions() conditions.Conditions {
	return workspace.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (workspace *Workspace) SetConditions(conditions conditions.Conditions) {
	workspace.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Workspace{}

// AzureName returns the Azure name of the resource
func (workspace *Workspace) AzureName() string {
	return workspace.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-07-01"
func (workspace Workspace) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (workspace *Workspace) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (workspace *Workspace) GetSpec() genruntime.ConvertibleSpec {
	return &workspace.Spec
}

// GetStatus returns the status of this resource
func (workspace *Workspace) GetStatus() genruntime.ConvertibleStatus {
	return &workspace.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.MachineLearningServices/workspaces"
func (workspace *Workspace) GetType() string {
	return "Microsoft.MachineLearningServices/workspaces"
}

// NewEmptyStatus returns a new empty (blank) status
func (workspace *Workspace) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Workspace_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (workspace *Workspace) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(workspace.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  workspace.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (workspace *Workspace) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Workspace_Status); ok {
		workspace.Status = *st
		return nil
	}

	// Convert status to required version
	var st Workspace_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	workspace.Status = st
	return nil
}

// Hub marks that this Workspace is the hub type for conversion
func (workspace *Workspace) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (workspace *Workspace) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: workspace.Spec.OriginalVersion,
		Kind:    "Workspace",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210701.Workspace
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/resourceDefinitions/workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

// Storage version of v1beta20210701.APIVersion
// +kubebuilder:validation:Enum={"2021-07-01"}
type APIVersion string

const APIVersionValue = APIVersion("2021-07-01")

// Storage version of v1beta20210701.Workspace_Status
type Workspace_Status struct {
	AllowPublicAccessWhenBehindVnet *bool                                                  `json:"allowPublicAccessWhenBehindVnet,omitempty"`
	ApplicationInsights             *string                                                `json:"applicationInsights,omitempty"`
	Conditions                      []conditions.Condition                                 `json:"conditions,omitempty"`
	ContainerRegistry               *string                                                `json:"containerRegistry,omitempty"`
	Description                     *string                                                `json:"description,omitempty"`
	DiscoveryUrl                    *string                                                `json:"discoveryUrl,omitempty"`
	Encryption                      *EncryptionProperty_Status                             `json:"encryption,omitempty"`
	FriendlyName                    *string                                                `json:"friendlyName,omitempty"`
	HbiWorkspace                    *bool                                                  `json:"hbiWorkspace,omitempty"`
	Id                              *string                                                `json:"id,omitempty"`
	Identity                        *Identity_Status                                       `json:"identity,omitempty"`
	ImageBuildCompute               *string                                                `json:"imageBuildCompute,omitempty"`
	KeyVault                        *string                                                `json:"keyVault,omitempty"`
	Location                        *string                                                `json:"location,omitempty"`
	MlFlowTrackingUri               *string                                                `json:"mlFlowTrackingUri,omitempty"`
	Name                            *string                                                `json:"name,omitempty"`
	NotebookInfo                    *NotebookResourceInfo_Status                           `json:"notebookInfo,omitempty"`
	PrimaryUserAssignedIdentity     *string                                                `json:"primaryUserAssignedIdentity,omitempty"`
	PrivateEndpointConnections      []PrivateEndpointConnection_Status_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PrivateLinkCount                *int                                                   `json:"privateLinkCount,omitempty"`
	PropertyBag                     genruntime.PropertyBag                                 `json:"$propertyBag,omitempty"`
	ProvisioningState               *string                                                `json:"provisioningState,omitempty"`
	PublicNetworkAccess             *string                                                `json:"publicNetworkAccess,omitempty"`
	ServiceManagedResourcesSettings *ServiceManagedResourcesSettings_Status                `json:"serviceManagedResourcesSettings,omitempty"`
	ServiceProvisionedResourceGroup *string                                                `json:"serviceProvisionedResourceGroup,omitempty"`
	SharedPrivateLinkResources      []SharedPrivateLinkResource_Status                     `json:"sharedPrivateLinkResources,omitempty"`
	Sku                             *Sku_Status                                            `json:"sku,omitempty"`
	StorageAccount                  *string                                                `json:"storageAccount,omitempty"`
	StorageHnsEnabled               *bool                                                  `json:"storageHnsEnabled,omitempty"`
	SystemData                      *SystemData_Status                                     `json:"systemData,omitempty"`
	Tags                            map[string]string                                      `json:"tags,omitempty"`
	TenantId                        *string                                                `json:"tenantId,omitempty"`
	Type                            *string                                                `json:"type,omitempty"`
	WorkspaceId                     *string                                                `json:"workspaceId,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Workspace_Status{}

// ConvertStatusFrom populates our Workspace_Status from the provided source
func (workspace *Workspace_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(workspace)
}

// ConvertStatusTo populates the provided destination from our Workspace_Status
func (workspace *Workspace_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(workspace)
}

// Storage version of v1beta20210701.Workspaces_Spec
type Workspaces_Spec struct {
	AllowPublicAccessWhenBehindVnet *bool `json:"allowPublicAccessWhenBehindVnet,omitempty"`

	// ApplicationInsightsReference: ARM id of the application insights associated with this workspace. This cannot be changed
	// once the workspace has been created
	ApplicationInsightsReference *genruntime.ResourceReference `armReference:"ApplicationInsights" json:"applicationInsightsReference,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// ContainerRegistryReference: ARM id of the container registry associated with this workspace. This cannot be changed once
	// the workspace has been created
	ContainerRegistryReference *genruntime.ResourceReference `armReference:"ContainerRegistry" json:"containerRegistryReference,omitempty"`
	Description                *string                       `json:"description,omitempty"`
	DiscoveryUrl               *string                       `json:"discoveryUrl,omitempty"`
	Encryption                 *EncryptionProperty           `json:"encryption,omitempty"`
	FriendlyName               *string                       `json:"friendlyName,omitempty"`
	HbiWorkspace               *bool                         `json:"hbiWorkspace,omitempty"`
	Identity                   *Identity                     `json:"identity,omitempty"`
	ImageBuildCompute          *string                       `json:"imageBuildCompute,omitempty"`

	// KeyVaultReference: ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has
	// been created
	KeyVaultReference *genruntime.ResourceReference `armReference:"KeyVault" json:"keyVaultReference,omitempty"`
	Location          *string                       `json:"location,omitempty"`
	OperatorSpec      *WorkspaceOperatorSpec        `json:"operatorSpec,omitempty"`
	OriginalVersion   string                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`

	// PrimaryUserAssignedIdentityReference: The user assigned identity resource id that represents the workspace identity.
	PrimaryUserAssignedIdentityReference *genruntime.ResourceReference                           `armReference:"PrimaryUserAssignedIdentity" json:"primaryUserAssignedIdentityReference,omitempty"`
	PropertyBag                          genruntime.PropertyBag                                  `json:"$propertyBag,omitempty"`
	PublicNetworkAccess                  *string                                                 `json:"publicNetworkAccess,omitempty"`
	ServiceManagedResourcesSettings      *ServiceManagedResourcesSettings                        `json:"serviceManagedResourcesSettings,omitempty"`
	SharedPrivateLinkResources           []Workspaces_Spec_Properties_SharedPrivateLinkResources `json:"sharedPrivateLinkResources,omitempty"`
	Sku                                  *Sku                                                    `json:"sku,omitempty"`

	// StorageAccountReference: ARM id of the storage account associated with this workspace. This cannot be changed once the
	// workspace has been created
	StorageAccountReference *genruntime.ResourceReference `armReference:"StorageAccount" json:"storageAccountReference,omitempty"`
	SystemData              *SystemData                   `json:"systemData,omitempty"`
	Tags                    map[string]string             `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Workspaces_Spec{}

// ConvertSpecFrom populates our Workspaces_Spec from the provided source
func (workspaces *Workspaces_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == workspaces {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(workspaces)
}

// ConvertSpecTo populates the provided destination from our Workspaces_Spec
func (workspaces *Workspaces_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == workspaces {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(workspaces)
}

// Storage version of v1beta20210701.EncryptionProperty
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/EncryptionProperty
type EncryptionProperty struct {
	Identity           *IdentityForCmk        `json:"identity,omitempty"`
	KeyVaultProperties *KeyVaultProperties    `json:"keyVaultProperties,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status             *string                `json:"status,omitempty"`
}

// Storage version of v1beta20210701.EncryptionProperty_Status
type EncryptionProperty_Status struct {
	Identity           *IdentityForCmk_Status     `json:"identity,omitempty"`
	KeyVaultProperties *KeyVaultProperties_Status `json:"keyVaultProperties,omitempty"`
	PropertyBag        genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Status             *string                    `json:"status,omitempty"`
}

// Storage version of v1beta20210701.Identity
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/Identity
type Identity struct {
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type                   *string                `json:"type,omitempty"`
	UserAssignedIdentities map[string]v1.JSON     `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1beta20210701.Identity_Status
type Identity_Status struct {
	PrincipalId            *string                                `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	TenantId               *string                                `json:"tenantId,omitempty"`
	Type                   *string                                `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentity_Status `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1beta20210701.NotebookResourceInfo_Status
type NotebookResourceInfo_Status struct {
	Fqdn                     *string                          `json:"fqdn,omitempty"`
	NotebookPreparationError *NotebookPreparationError_Status `json:"notebookPreparationError,omitempty"`
	PropertyBag              genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	ResourceId               *string                          `json:"resourceId,omitempty"`
}

// Storage version of v1beta20210701.PrivateEndpointConnection_Status_SubResourceEmbedded
type PrivateEndpointConnection_Status_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	Identity    *Identity_Status       `json:"identity,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Sku         *Sku_Status            `json:"sku,omitempty"`
	SystemData  *SystemData_Status     `json:"systemData,omitempty"`
}

// Storage version of v1beta20210701.ServiceManagedResourcesSettings
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/ServiceManagedResourcesSettings
type ServiceManagedResourcesSettings struct {
	CosmosDb    *CosmosDbSettings      `json:"cosmosDb,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.ServiceManagedResourcesSettings_Status
type ServiceManagedResourcesSettings_Status struct {
	CosmosDb    *CosmosDbSettings_Status `json:"cosmosDb,omitempty"`
	PropertyBag genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.SharedPrivateLinkResource_Status
type SharedPrivateLinkResource_Status struct {
	GroupId               *string                `json:"groupId,omitempty"`
	Name                  *string                `json:"name,omitempty"`
	PrivateLinkResourceId *string                `json:"privateLinkResourceId,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequestMessage        *string                `json:"requestMessage,omitempty"`
	Status                *string                `json:"status,omitempty"`
}

// Storage version of v1beta20210701.Sku
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/Sku
type Sku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20210701.Sku_Status
type Sku_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20210701.SystemData
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/SystemData
type SystemData struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.SystemData_Status
type SystemData_Status struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.WorkspaceOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type WorkspaceOperatorSpec struct {
	PropertyBag genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	Secrets     *WorkspaceOperatorSecrets `json:"secrets,omitempty"`
}

// Storage version of v1beta20210701.Workspaces_Spec_Properties_SharedPrivateLinkResources
type Workspaces_Spec_Properties_SharedPrivateLinkResources struct {
	GroupId *string `json:"groupId,omitempty"`
	Name    *string `json:"name,omitempty"`

	// PrivateLinkResourceReference: The resource id that private link links to.
	PrivateLinkResourceReference *genruntime.ResourceReference `armReference:"PrivateLinkResourceId" json:"privateLinkResourceReference,omitempty"`
	PropertyBag                  genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	RequestMessage               *string                       `json:"requestMessage,omitempty"`
	Status                       *string                       `json:"status,omitempty"`
}

// Storage version of v1beta20210701.CosmosDbSettings
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/CosmosDbSettings
type CosmosDbSettings struct {
	CollectionsThroughput *int                   `json:"collectionsThroughput,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.CosmosDbSettings_Status
type CosmosDbSettings_Status struct {
	CollectionsThroughput *int                   `json:"collectionsThroughput,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.IdentityForCmk
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/IdentityForCmk
type IdentityForCmk struct {
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UserAssignedIdentity *string                `json:"userAssignedIdentity,omitempty"`
}

// Storage version of v1beta20210701.IdentityForCmk_Status
type IdentityForCmk_Status struct {
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UserAssignedIdentity *string                `json:"userAssignedIdentity,omitempty"`
}

// Storage version of v1beta20210701.KeyVaultProperties
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/KeyVaultProperties
type KeyVaultProperties struct {
	IdentityClientId *string                `json:"identityClientId,omitempty"`
	KeyIdentifier    *string                `json:"keyIdentifier,omitempty"`
	KeyVaultArmId    *string                `json:"keyVaultArmId,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.KeyVaultProperties_Status
type KeyVaultProperties_Status struct {
	IdentityClientId *string                `json:"identityClientId,omitempty"`
	KeyIdentifier    *string                `json:"keyIdentifier,omitempty"`
	KeyVaultArmId    *string                `json:"keyVaultArmId,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.NotebookPreparationError_Status
type NotebookPreparationError_Status struct {
	ErrorMessage *string                `json:"errorMessage,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StatusCode   *int                   `json:"statusCode,omitempty"`
}

// Storage version of v1beta20210701.UserAssignedIdentity_Status
type UserAssignedIdentity_Status struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId    *string                `json:"tenantId,omitempty"`
}

// Storage version of v1beta20210701.WorkspaceOperatorSecrets
type WorkspaceOperatorSecrets struct {
	AppInsightsInstrumentationKey *genruntime.SecretDestination `json:"appInsightsInstrumentationKey,omitempty"`
	ContainerRegistryPassword     *genruntime.SecretDestination `json:"containerRegistryPassword,omitempty"`
	ContainerRegistryPassword2    *genruntime.SecretDestination `json:"containerRegistryPassword2,omitempty"`
	ContainerRegistryUserName     *genruntime.SecretDestination `json:"containerRegistryUserName,omitempty"`
	PrimaryNotebookAccessKey      *genruntime.SecretDestination `json:"primaryNotebookAccessKey,omitempty"`
	PropertyBag                   genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecondaryNotebookAccessKey    *genruntime.SecretDestination `json:"secondaryNotebookAccessKey,omitempty"`
	UserStorageKey                *genruntime.SecretDestination `json:"userStorageKey,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
