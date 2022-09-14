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

// EquinixStatus : Connection status
type EquinixStatus string

// List of EquinixStatus
const (
	REJECTED_ACK_EquinixStatus                    EquinixStatus = "REJECTED_ACK"
	REJECTED_EquinixStatus                        EquinixStatus = "REJECTED"
	PENDING_DELETE_EquinixStatus                  EquinixStatus = "PENDING_DELETE"
	PROVISIONED_EquinixStatus                     EquinixStatus = "PROVISIONED"
	BEING_REPROVISIONED_EquinixStatus             EquinixStatus = "BEING_REPROVISIONED"
	BEING_DEPROVISIONED_EquinixStatus             EquinixStatus = "BEING_DEPROVISIONED"
	BEING_PROVISIONED_EquinixStatus               EquinixStatus = "BEING_PROVISIONED"
	CREATED_EquinixStatus                         EquinixStatus = "CREATED"
	ERRORED_EquinixStatus                         EquinixStatus = "ERRORED"
	PENDING_DEPROVISIONING_EquinixStatus          EquinixStatus = "PENDING_DEPROVISIONING"
	APPROVED_EquinixStatus                        EquinixStatus = "APPROVED"
	ORDERING_EquinixStatus                        EquinixStatus = "ORDERING"
	PENDING_APPROVAL_EquinixStatus                EquinixStatus = "PENDING_APPROVAL"
	NOT_PROVISIONED_EquinixStatus                 EquinixStatus = "NOT_PROVISIONED"
	DEPROVISIONING_EquinixStatus                  EquinixStatus = "DEPROVISIONING"
	NOT_DEPROVISIONED_EquinixStatus               EquinixStatus = "NOT_DEPROVISIONED"
	PENDING_AUTO_APPROVAL_EquinixStatus           EquinixStatus = "PENDING_AUTO_APPROVAL"
	PROVISIONING_EquinixStatus                    EquinixStatus = "PROVISIONING"
	PENDING_BGP_PEERING_EquinixStatus             EquinixStatus = "PENDING_BGP_PEERING"
	PENDING_PROVIDER_VLAN_EquinixStatus           EquinixStatus = "PENDING_PROVIDER_VLAN"
	DEPROVISIONED_EquinixStatus                   EquinixStatus = "DEPROVISIONED"
	DELETED_EquinixStatus                         EquinixStatus = "DELETED"
	PENDING_BANDWIDTH_APPROVAL_EquinixStatus      EquinixStatus = "PENDING_BANDWIDTH_APPROVAL"
	AUTO_APPROVAL_FAILED_EquinixStatus            EquinixStatus = "AUTO_APPROVAL_FAILED"
	UPDATE_PENDING_EquinixStatus                  EquinixStatus = "UPDATE_PENDING"
	DELETED_API_EquinixStatus                     EquinixStatus = "DELETED_API"
	MODIFIED_EquinixStatus                        EquinixStatus = "MODIFIED"
	PENDING_PROVIDER_VLAN_ERROR_EquinixStatus     EquinixStatus = "PENDING_PROVIDER_VLAN_ERROR"
	DRAFT_EquinixStatus                           EquinixStatus = "DRAFT"
	CANCELLED_EquinixStatus                       EquinixStatus = "CANCELLED"
	PENDING_INTERFACE_CONFIGURATION_EquinixStatus EquinixStatus = "PENDING_INTERFACE_CONFIGURATION"
)
