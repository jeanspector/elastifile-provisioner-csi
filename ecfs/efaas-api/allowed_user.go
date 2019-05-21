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

import (
	"time"
)

type AllowedUser struct {

	// User email address.
	User string `json:"user,omitempty"`

	// Email address of the user added this user to project.
	AddedBy string `json:"addedBy,omitempty"`

	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
}