# VirtualGateway

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [optional] [default to null]
**Name** | **string** | Customer-provided Fabric Gateway name | [optional] [default to null]
**Location** | [***SimplifiedLocationWithoutIbx**](SimplifiedLocationWithoutIBX.md) |  | [optional] [default to null]
**Package_** | [***VirtualGatewayPackageType**](VirtualGatewayPackageType.md) |  | [optional] [default to null]
**Order** | [***Order**](Order.md) |  | [optional] [default to null]
**Project** | [***Project**](Project.md) |  | [optional] [default to null]
**Account** | [***SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] [default to null]
**Notifications** | [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on connection configuration or status changes | [optional] [default to null]
**Href** | **string** | Fabric Gateway URI | [optional] [default to null]
**Uuid** | **string** | Equinix-assigned access point identifier | [optional] [default to null]
**State** | [***VirtualGatewayAccessPointState**](VirtualGatewayAccessPointState.md) |  | [optional] [default to null]
**BgpIpv4RoutesCount** | **int32** | Access point used and maximum number of IPv4 BGP routes | [optional] [default to null]
**BgpIpv6RoutesCount** | **int32** | Access point used and maximum number of IPv6 BGP routes | [optional] [default to null]
**ConnectionsCount** | **int32** | Number of connections associated with this Access point | [optional] [default to null]
**ChangeLog** | [***Changelog**](Changelog.md) |  | [optional] [default to null]
**Change** | [***GatewayChange**](GatewayChange.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

