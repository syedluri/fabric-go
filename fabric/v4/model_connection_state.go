/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.3.52
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// ConnectionState : Connection status
type ConnectionState string

// List of ConnectionState
const (
	ACTIVE_ConnectionState         ConnectionState = "ACTIVE"
	CANCELLED_ConnectionState      ConnectionState = "CANCELLED"
	DEPROVISIONED_ConnectionState  ConnectionState = "DEPROVISIONED"
	DEPROVISIONING_ConnectionState ConnectionState = "DEPROVISIONING"
	DRAFT_ConnectionState          ConnectionState = "DRAFT"
	FAILED_ConnectionState         ConnectionState = "FAILED"
	PENDING_ConnectionState        ConnectionState = "PENDING"
	PROVISIONED_ConnectionState    ConnectionState = "PROVISIONED"
	PROVISIONING_ConnectionState   ConnectionState = "PROVISIONING"
	REPROVISIONING_ConnectionState ConnectionState = "REPROVISIONING"
	EMPTY_ConnectionState          ConnectionState = ""
)
