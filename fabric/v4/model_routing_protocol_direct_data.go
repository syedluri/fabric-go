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

type RoutingProtocolDirectData struct {
	// Routing protocol type
	Type_      string                `json:"type,omitempty"`
	Name       string                `json:"name,omitempty"`
	DirectIpv4 *DirectConnectionIpv4 `json:"directIpv4,omitempty"`
	DirectIpv6 *DirectConnectionIpv6 `json:"directIpv6,omitempty"`
	// Routing Protocol URI
	Href string `json:"href,omitempty"`
	// Routing protocol identifier
	Uuid      string                    `json:"uuid,omitempty"`
	State     string                    `json:"state,omitempty"`
	Operation *RoutingProtocolOperation `json:"operation,omitempty"`
	Change    *RoutingProtocolChange    `json:"change,omitempty"`
	Changelog *Changelog                `json:"changelog,omitempty"`
}
