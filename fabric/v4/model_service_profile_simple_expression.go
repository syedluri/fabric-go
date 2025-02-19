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

type ServiceProfileSimpleExpression struct {
	// Possible field names to use on filters:  * `/name` - Service Profile name  * `/uuid` - Service Profile uuid  * `/state` - Service Profile status  * `/metros/code` - Service Profile metro code  * `/visibility` - Service Profile package  * `/type` - Service Profile package
	Property string `json:"property,omitempty"`
	// Possible operators to use on filters:  * `=` - equal
	Operator string   `json:"operator,omitempty"`
	Values   []string `json:"values,omitempty"`
}
