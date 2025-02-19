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

// GET Metros retrieves all Equinix® Fabric™ metros, as well as latency data for each location.This performance data helps network planning engineers and administrators make strategic decisions about port locations and traffic routes.
type MetroResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// List of Fabric Metros.
	Data []Metro `json:"data,omitempty"`
}
