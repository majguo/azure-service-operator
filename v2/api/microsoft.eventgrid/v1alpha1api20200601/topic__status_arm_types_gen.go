// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

//Generated from:
type Topic_StatusARM struct {
	//Id: Fully qualified identifier of the resource.
	Id *string `json:"id,omitempty"`

	//Location: Location of the resource.
	Location *string `json:"location,omitempty"`

	//Name: Name of the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the topic.
	Properties *TopicProperties_StatusARM `json:"properties,omitempty"`

	//SystemData: The system metadata relating to Topic resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	//Tags: Tags of the resource.
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

//Generated from:
type TopicProperties_StatusARM struct {
	//Endpoint: Endpoint for the topic.
	Endpoint *string `json:"endpoint,omitempty"`

	//InboundIpRules: This can be used to restrict traffic from specific IPs instead
	//of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRule_StatusARM `json:"inboundIpRules,omitempty"`

	//InputSchema: This determines the format that Event Grid should expect for
	//incoming events published to the topic.
	InputSchema *TopicPropertiesStatusInputSchema `json:"inputSchema,omitempty"`

	//InputSchemaMapping: This enables publishing using custom event schemas. An
	//InputSchemaMapping can be specified to map various properties of a source schema
	//to various required properties of the EventGridEvent schema.
	InputSchemaMapping *InputSchemaMapping_StatusARM `json:"inputSchemaMapping,omitempty"`

	//MetricResourceId: Metric resource id for the topic.
	MetricResourceId           *string                                                         `json:"metricResourceId,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`

	//ProvisioningState: Provisioning state of the topic.
	ProvisioningState *TopicPropertiesStatusProvisioningState `json:"provisioningState,omitempty"`

	//PublicNetworkAccess: This determines if traffic is allowed over public network.
	//By default it is enabled.
	//You can further restrict to specific IPs by configuring <seealso
	//cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.TopicProperties.InboundIpRules"
	///>
	PublicNetworkAccess *TopicPropertiesStatusPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
}

//Generated from:
type PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM struct {
	//Id: Fully qualified identifier of the resource.
	Id *string `json:"id,omitempty"`
}
