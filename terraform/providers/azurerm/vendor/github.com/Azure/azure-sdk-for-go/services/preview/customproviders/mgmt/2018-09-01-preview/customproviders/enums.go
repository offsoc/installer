package customproviders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ActionRouting enumerates the values for action routing.
type ActionRouting string

const (
	// Proxy ...
	Proxy ActionRouting = "Proxy"
)

// PossibleActionRoutingValues returns an array of possible values for the ActionRouting const type.
func PossibleActionRoutingValues() []ActionRouting {
	return []ActionRouting{Proxy}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Accepted ...
	Accepted ProvisioningState = "Accepted"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Running ...
	Running ProvisioningState = "Running"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Accepted, Deleting, Failed, Running, Succeeded}
}

// ResourceTypeRouting enumerates the values for resource type routing.
type ResourceTypeRouting string

const (
	// ResourceTypeRoutingProxy ...
	ResourceTypeRoutingProxy ResourceTypeRouting = "Proxy"
	// ResourceTypeRoutingProxyCache ...
	ResourceTypeRoutingProxyCache ResourceTypeRouting = "Proxy,Cache"
)

// PossibleResourceTypeRoutingValues returns an array of possible values for the ResourceTypeRouting const type.
func PossibleResourceTypeRoutingValues() []ResourceTypeRouting {
	return []ResourceTypeRouting{ResourceTypeRoutingProxy, ResourceTypeRoutingProxyCache}
}

// ValidationType enumerates the values for validation type.
type ValidationType string

const (
	// Swagger ...
	Swagger ValidationType = "Swagger"
)

// PossibleValidationTypeValues returns an array of possible values for the ValidationType const type.
func PossibleValidationTypeValues() []ValidationType {
	return []ValidationType{Swagger}
}