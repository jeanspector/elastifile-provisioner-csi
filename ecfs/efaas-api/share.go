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

type Share struct {

	// [Output Only] Share name
	Name string `json:"name,omitempty"`

	// [Output Only] NFS mount point to be used in mount command.
	NfsMountPoint string `json:"nfsMountPoint,omitempty"`
}
