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

// VirtualDevice Access Point Type
type ServiceProfileAccessPointTypeVd struct {
	Type_               *ServiceProfileAccessPointTypeEnum `json:"type"`
	Uuid                string                             `json:"uuid,omitempty"`
	SupportedBandwidths *[]int32                           `json:"supportedBandwidths,omitempty"`
	// Allow remote connections to Service Profile
	AllowRemoteConnections bool `json:"allowRemoteConnections,omitempty"`
	AllowCustomBandwidth   bool `json:"allowCustomBandwidth,omitempty"`
}
