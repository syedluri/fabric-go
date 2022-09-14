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

// RouteTableEntrySortBy : Possible field names to use on sorting
type RouteTableEntrySortBy string

// List of RouteTableEntrySortBy
const (
	CHANGE_LOGCREATED_DATE_TIME_RouteTableEntrySortBy RouteTableEntrySortBy = "/changeLog/createdDateTime"
	CHANGE_LOGUPDATED_DATE_TIME_RouteTableEntrySortBy RouteTableEntrySortBy = "/changeLog/updatedDateTime"
	PREFIX_RouteTableEntrySortBy                      RouteTableEntrySortBy = "/prefix"
	NEXT_HOP_RouteTableEntrySortBy                    RouteTableEntrySortBy = "/nextHop"
	CONNECTIONNAME_RouteTableEntrySortBy              RouteTableEntrySortBy = "/connection/name"
	TYPE__RouteTableEntrySortBy                       RouteTableEntrySortBy = "/type"
)
