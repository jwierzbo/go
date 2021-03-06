/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
type V1HostAlias struct {

	// Hostnames for the above IP address.
	Hostnames []string `json:"hostnames,omitempty"`

	// IP address of the host file entry.
	Ip string `json:"ip,omitempty"`
}
