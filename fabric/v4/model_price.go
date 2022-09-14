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

type Price struct {
	// An absolute URL that returns specified pricing data
	Href  string       `json:"href,omitempty"`
	Type_ *ProductType `json:"type,omitempty"`
	// Equinix-assigned product code
	Code string `json:"code,omitempty"`
	// Full product name
	Name string `json:"name,omitempty"`
	// Product description
	Description string             `json:"description,omitempty"`
	Account     *SimplifiedAccount `json:"account,omitempty"`
	Charges     []PriceCharge      `json:"charges,omitempty"`
	// Product offering price currency
	Currency string `json:"currency,omitempty"`
	// In months. No value means unlimited
	TermLength int32                   `json:"termLength,omitempty"`
	Catgory    *PriceCategory          `json:"catgory,omitempty"`
	Connection *VirtualConnectionPrice `json:"connection,omitempty"`
	IpBlock    *IpBlockPrice           `json:"ipBlock,omitempty"`
	Gateway    *FabricGatewayPrice     `json:"gateway,omitempty"`
	Port       *VirtualPortPrice       `json:"port,omitempty"`
}
