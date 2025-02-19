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

// LinkProtocolType : Type of Link Protocol
type LinkProtocolType string

// List of LinkProtocolType
const (
	UNTAGGED_LinkProtocolType   LinkProtocolType = "UNTAGGED"
	DOT1_Q_LinkProtocolType     LinkProtocolType = "DOT1Q"
	QINQ_LinkProtocolType       LinkProtocolType = "QINQ"
	EVPN_VXLAN_LinkProtocolType LinkProtocolType = "EVPN_VXLAN"
)
