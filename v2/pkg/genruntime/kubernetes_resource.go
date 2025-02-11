/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ARMOwned interface {
	// Owner returns the ResourceReference of the owner, or nil if there is no owner
	Owner() *ResourceReference
}

// KubernetesResource is an Azure resource. This interface contains the common set of
// methods that apply to all ASO ARM resources.
type KubernetesResource interface {
	ARMOwned

	// TODO: I think we need this?
	// KnownOwner() *KnownResourceReference

	// AzureName returns the Azure name of the resource
	AzureName() string

	// GetType returns the type of the resource according to Azure. For example Microsoft.Resources/resourceGroups or
	// Microsoft.Network/networkSecurityGroups/securityRules
	GetType() string

	// GetResourceKind returns the ResourceKind of the resource.
	GetResourceKind() ResourceKind

	// Some types, but not all, have a corresponding:
	// 	SetAzureName(name string)
	// They do not if the name must be a fixed value (like 'default').

	// GetAPIVersion returns the API Version of the resource
	GetAPIVersion() string

	// GetSpec returns the specification of the resource
	GetSpec() ConvertibleSpec

	// GetStatus returns the current status of the resource
	GetStatus() ConvertibleStatus

	// NewEmptyStatus returns a blank status ready for population
	NewEmptyStatus() ConvertibleStatus

	// SetStatus updates the status of the resource
	SetStatus(status ConvertibleStatus) error
}

// NewEmptyVersionedResource returns a new blank resource based on the passed metaObject; the original API version used
// (if available) from when the resource was first created is used to identify the version to return.
// Returns an empty resource.
func NewEmptyVersionedResource(metaObject ARMMetaObject, scheme *runtime.Scheme) (ARMMetaObject, error) {
	return NewEmptyVersionedResourceFromGVK(scheme, GetOriginalGVK(metaObject))
}

// NewEmptyVersionedResourceFromGVK creates a new empty versioned resource from the specified GVK
func NewEmptyVersionedResourceFromGVK(scheme *runtime.Scheme, gvk schema.GroupVersionKind) (ARMMetaObject, error) {
	// Create an empty resource at the desired version
	rsrc, err := scheme.New(gvk)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create new %s", gvk)
	}

	// Convert it to our interface
	mo, ok := rsrc.(ARMMetaObject)
	if !ok {
		return nil, errors.Errorf("expected resource %s to implement genruntime.ARMMetaObject", gvk)
	}

	// Ensure GVK is populated
	mo.GetObjectKind().SetGroupVersionKind(gvk)

	// Return the empty resource
	return mo, nil
}
