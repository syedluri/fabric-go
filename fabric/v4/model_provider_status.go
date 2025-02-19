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

// ProviderStatus : Connection provider readiness status
type ProviderStatus string

// List of ProviderStatus
const (
	AVAILABLE_ProviderStatus        ProviderStatus = "AVAILABLE"
	DEPROVISIONED_ProviderStatus    ProviderStatus = "DEPROVISIONED"
	DEPROVISIONING_ProviderStatus   ProviderStatus = "DEPROVISIONING"
	FAILED_ProviderStatus           ProviderStatus = "FAILED"
	NOT_AVAILABLE_ProviderStatus    ProviderStatus = "NOT_AVAILABLE"
	PENDING_APPROVAL_ProviderStatus ProviderStatus = "PENDING_APPROVAL"
	PROVISIONED_ProviderStatus      ProviderStatus = "PROVISIONED"
	PROVISIONING_ProviderStatus     ProviderStatus = "PROVISIONING"
	REJECTED_ProviderStatus         ProviderStatus = "REJECTED"
	PENDING_BGP_ProviderStatus      ProviderStatus = "PENDING_BGP"
	OUT_OF_BANDWIDTH_ProviderStatus ProviderStatus = "OUT_OF_BANDWIDTH"
	DELETED_ProviderStatus          ProviderStatus = "DELETED"
	ERROR__ProviderStatus           ProviderStatus = "ERROR"
	ERRORED_ProviderStatus          ProviderStatus = "ERRORED"
	NOTPROVISIONED_ProviderStatus   ProviderStatus = "NOTPROVISIONED"
	NOT_PROVISIONED_ProviderStatus  ProviderStatus = "NOT_PROVISIONED"
	ORDERING_ProviderStatus         ProviderStatus = "ORDERING"
	DELETING_ProviderStatus         ProviderStatus = "DELETING"
	PENDING_DELETE_ProviderStatus   ProviderStatus = "PENDING DELETE"
	NA_ProviderStatus               ProviderStatus = "N/A"
)
