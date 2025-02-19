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

// Connection specification
type Connection struct {
	Type_ *ConnectionType `json:"type"`
	// Connection URI
	Href string `json:"href,omitempty"`
	// Equinix-assigned connection identifier
	Uuid string `json:"uuid,omitempty"`
	// Customer-provided connection name
	Name string `json:"name"`
	// Customer-provided connection description
	Description string               `json:"description,omitempty"`
	State       *ConnectionState     `json:"state,omitempty"`
	Change      *Change              `json:"change,omitempty"`
	Operation   *ConnectionOperation `json:"operation,omitempty"`
	Order       *Order               `json:"order,omitempty"`
	// Preferences for notifications on connection configuration or status changes
	Notifications []SimplifiedNotification `json:"notifications,omitempty"`
	Account       *SimplifiedAccount       `json:"account,omitempty"`
	ChangeLog     *Changelog               `json:"changeLog,omitempty"`
	// Connection bandwidth in Mbps
	Bandwidth  int32                 `json:"bandwidth"`
	Redundancy *ConnectionRedundancy `json:"redundancy,omitempty"`
	// Connection property derived from access point locations
	IsRemote  bool                 `json:"isRemote,omitempty"`
	Direction *ConnectionDirection `json:"direction,omitempty"`
	ASide     *ConnectionSide      `json:"aSide"`
	ZSide     *ConnectionSide      `json:"zSide"`
	// Connection additional information
	AdditionalInfo []ConnectionSideAdditionalInfo `json:"additionalInfo,omitempty"`
	Project        *Project                       `json:"project,omitempty"`
}
