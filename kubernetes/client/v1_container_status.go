/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// ContainerStatus contains details for the current status of this container.
type V1ContainerStatus struct {

	// Container's ID in the format 'docker://<container_id>'.
	ContainerID string `json:"containerID,omitempty"`

	// The image the container is running. More info: https://kubernetes.io/docs/concepts/containers/images
	Image string `json:"image"`

	// ImageID of the container's image.
	ImageID string `json:"imageID"`

	// Details about the container's last termination condition.
	LastState *V1ContainerState `json:"lastState,omitempty"`

	// This must be a DNS_LABEL. Each container in a pod must have a unique name. Cannot be updated.
	Name string `json:"name"`

	// Specifies whether the container has passed its readiness probe.
	Ready bool `json:"ready"`

	// The number of times the container has been restarted, currently based on the number of dead containers that have not yet been removed. Note that this is calculated from dead containers. But those containers are subject to garbage collection. This value will get capped at 5 by GC.
	RestartCount int32 `json:"restartCount"`

	// Details about the container's current condition.
	State *V1ContainerState `json:"state,omitempty"`
}
