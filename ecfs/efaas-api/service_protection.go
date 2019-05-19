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

type ServiceProtection struct {

	// Defines the Availability Zones support
	ProtectionMode string `json:"protectionMode,omitempty"`

	// Number of data copies
	ReplicationLevel int32 `json:"replicationLevel,omitempty"`
}
