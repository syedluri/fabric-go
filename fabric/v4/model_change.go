/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.4
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

import (
	"time"
)

// Current state of latest connection change
type Change struct {
	// Uniquely identifies a change
	Uuid string `json:"uuid,omitempty"`
	// Type of change
	Type_ string `json:"type"`
	// Current outcome of the change flow
	Status string `json:"status,omitempty"`
	// Set when change flow starts
	CreatedDateTime time.Time `json:"createdDateTime"`
	// Set when change object is updated
	UpdatedDateTime time.Time `json:"updatedDateTime,omitempty"`
	// Additional information
	Information string                     `json:"information,omitempty"`
	Data        *ConnectionChangeOperation `json:"data,omitempty"`
}
