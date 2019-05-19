/* 
 * Elastifile FaaS API
 *
 * Elastifile Filesystem as a Service API
 *
 * OpenAPI spec version: 2.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package EfaasApi

type ServiceClass struct {

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id string `json:"id,omitempty"`

	// [Output Only] Name of the resource.
	Name string `json:"name"`

	// [Output Only] A textual description of the resource.
	Description string `json:"description"`

	// 
	LongDescription string `json:"longDescription,omitempty"`

	// ServiceClass supported regions
	Regions []Region `json:"regions,omitempty"`

	// Storage backend device type
	StorageBackend string `json:"storageBackend"`

	CapacityUnits CapacityUnits `json:"capacityUnits,omitempty"`

	ServiceProtection ServiceProtection `json:"serviceProtection,omitempty"`

	// 
	NodeType string `json:"nodeType"`

	ClearTier ClearTier `json:"clearTier,omitempty"`

	StoragePrice StoragePrice `json:"storagePrice,omitempty"`

	// 
	DisplayOrder int32 `json:"displayOrder,omitempty"`

	Deprecated Deprecated `json:"deprecated,omitempty"`
}
