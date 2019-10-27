/*
 * Elastifile FaaS API
 *
 * Elastifile Filesystem as a Service API
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package efaasapi

type Snapshots struct {
	// [Output Only] The unique identifier for the snapshot. This identifier is defined by the server.
	Id string `json:"id,omitempty"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035
	Name string `json:"name"`
	// Snapshot retention policy. The number of days to hold the snapshot till automatic deletion. Default 0, meaning no retention policy defined.
	Retention float32 `json:"retention"`
	// [Output Only] The filesystem instance id that this snapshot was taken for.
	InstanceId string `json:"instanceId,omitempty"`
	// [Output Only] The filesystem instance name that this snapshot was taken for.
	InstanceName string `json:"instanceName,omitempty"`
	// [Output Only] The filesystem id that this snapshot was taken for.
	FilesystemId string `json:"filesystemId,omitempty"`
	// [Output Only] The filesystem name that this snapshot was taken for.
	FilesystemName string `json:"filesystemName,omitempty"`
	// Snapshot region location.
	Region string `json:"region,omitempty"`
	// [Output Only] Snapshot size in bytes.
	Size int32 `json:"size,omitempty"`
	// Snapshot scheduling Daily, Weekly, Monthly or Manual.
	Schedule string `json:"schedule,omitempty"`
	// [Output Only] If exists, this object includes the snapshot share parameters.
	Share *Share `json:"share,omitempty"`
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
	// 
	DeletionTime string `json:"deletionTime,omitempty"`
	// [Output Only] The status of the snapshot. A snapshot can be used to mount a previous copy of the filesystem, only after the snapshot has been successfully created and the status is set to READY. Possible values are PENDING, READY.
	Status string `json:"status,omitempty"`
}