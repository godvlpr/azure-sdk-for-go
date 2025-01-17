//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorder_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListOperations.json
func ExampleManagementClient_ListOperations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListOperations(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListAddressesAtSubscriptionLevel.json
func ExampleManagementClient_ListAddressesAtSubscriptionLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListAddressesAtSubscriptionLevel(&armedgeorder.ManagementClientListAddressesAtSubscriptionLevelOptions{Filter: nil,
		SkipToken: nil,
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListProductFamilies.json
func ExampleManagementClient_ListProductFamilies() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListProductFamilies(armedgeorder.ProductFamiliesRequest{
		FilterableProperties: map[string][]*armedgeorder.FilterableProperty{
			"azurestackedge": {
				{
					Type: to.Ptr(armedgeorder.SupportedFilterTypesShipToCountries),
					SupportedValues: []*string{
						to.Ptr("US")},
				}},
		},
	},
		&armedgeorder.ManagementClientListProductFamiliesOptions{Expand: to.Ptr("<expand>"),
			SkipToken: nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListConfigurations.json
func ExampleManagementClient_ListConfigurations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListConfigurations(armedgeorder.ConfigurationsRequest{
		ConfigurationFilters: []*armedgeorder.ConfigurationFilters{
			{
				FilterableProperty: []*armedgeorder.FilterableProperty{
					{
						Type: to.Ptr(armedgeorder.SupportedFilterTypesShipToCountries),
						SupportedValues: []*string{
							to.Ptr("US")},
					}},
				HierarchyInformation: &armedgeorder.HierarchyInformation{
					ProductFamilyName: to.Ptr("<product-family-name>"),
					ProductLineName:   to.Ptr("<product-line-name>"),
					ProductName:       to.Ptr("<product-name>"),
				},
			}},
	},
		&armedgeorder.ManagementClientListConfigurationsOptions{SkipToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListProductFamiliesMetadata.json
func ExampleManagementClient_ListProductFamiliesMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListProductFamiliesMetadata(&armedgeorder.ManagementClientListProductFamiliesMetadataOptions{SkipToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListOrderAtSubscriptionLevel.json
func ExampleManagementClient_ListOrderAtSubscriptionLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListOrderAtSubscriptionLevel(&armedgeorder.ManagementClientListOrderAtSubscriptionLevelOptions{SkipToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListOrderItemsAtSubscriptionLevel.json
func ExampleManagementClient_ListOrderItemsAtSubscriptionLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListOrderItemsAtSubscriptionLevel(&armedgeorder.ManagementClientListOrderItemsAtSubscriptionLevelOptions{Filter: nil,
		Expand:    nil,
		SkipToken: nil,
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListAddressesAtResourceGroupLevel.json
func ExampleManagementClient_ListAddressesAtResourceGroupLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListAddressesAtResourceGroupLevel("<resource-group-name>",
		&armedgeorder.ManagementClientListAddressesAtResourceGroupLevelOptions{Filter: nil,
			SkipToken: nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/GetAddressByName.json
func ExampleManagementClient_GetAddressByName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetAddressByName(ctx,
		"<address-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/CreateAddress.json
func ExampleManagementClient_BeginCreateAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateAddress(ctx,
		"<address-name>",
		"<resource-group-name>",
		armedgeorder.AddressResource{
			Location: to.Ptr("<location>"),
			Properties: &armedgeorder.AddressProperties{
				ContactDetails: &armedgeorder.ContactDetails{
					ContactName: to.Ptr("<contact-name>"),
					EmailList: []*string{
						to.Ptr("xxxx@xxxx.xxx")},
					Phone:          to.Ptr("<phone>"),
					PhoneExtension: to.Ptr("<phone-extension>"),
				},
				ShippingAddress: &armedgeorder.ShippingAddress{
					AddressType:     to.Ptr(armedgeorder.AddressTypeNone),
					City:            to.Ptr("<city>"),
					CompanyName:     to.Ptr("<company-name>"),
					Country:         to.Ptr("<country>"),
					PostalCode:      to.Ptr("<postal-code>"),
					StateOrProvince: to.Ptr("<state-or-province>"),
					StreetAddress1:  to.Ptr("<street-address1>"),
					StreetAddress2:  to.Ptr("<street-address2>"),
				},
			},
		},
		&armedgeorder.ManagementClientBeginCreateAddressOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/DeleteAddressByName.json
func ExampleManagementClient_BeginDeleteAddressByName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDeleteAddressByName(ctx,
		"<address-name>",
		"<resource-group-name>",
		&armedgeorder.ManagementClientBeginDeleteAddressByNameOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/UpdateAddress.json
func ExampleManagementClient_BeginUpdateAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdateAddress(ctx,
		"<address-name>",
		"<resource-group-name>",
		armedgeorder.AddressUpdateParameter{
			Properties: &armedgeorder.AddressUpdateProperties{
				ContactDetails: &armedgeorder.ContactDetails{
					ContactName: to.Ptr("<contact-name>"),
					EmailList: []*string{
						to.Ptr("xxxx@xxxx.xxx")},
					Phone:          to.Ptr("<phone>"),
					PhoneExtension: to.Ptr("<phone-extension>"),
				},
				ShippingAddress: &armedgeorder.ShippingAddress{
					AddressType:     to.Ptr(armedgeorder.AddressTypeNone),
					City:            to.Ptr("<city>"),
					CompanyName:     to.Ptr("<company-name>"),
					Country:         to.Ptr("<country>"),
					PostalCode:      to.Ptr("<postal-code>"),
					StateOrProvince: to.Ptr("<state-or-province>"),
					StreetAddress1:  to.Ptr("<street-address1>"),
					StreetAddress2:  to.Ptr("<street-address2>"),
				},
			},
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
		},
		&armedgeorder.ManagementClientBeginUpdateAddressOptions{IfMatch: nil,
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListOrderAtResourceGroupLevel.json
func ExampleManagementClient_ListOrderAtResourceGroupLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListOrderAtResourceGroupLevel("<resource-group-name>",
		&armedgeorder.ManagementClientListOrderAtResourceGroupLevelOptions{SkipToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/GetOrderByName.json
func ExampleManagementClient_GetOrderByName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetOrderByName(ctx,
		"<order-name>",
		"<resource-group-name>",
		"<location>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListOrderItemsAtResourceGroupLevel.json
func ExampleManagementClient_ListOrderItemsAtResourceGroupLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListOrderItemsAtResourceGroupLevel("<resource-group-name>",
		&armedgeorder.ManagementClientListOrderItemsAtResourceGroupLevelOptions{Filter: nil,
			Expand:    nil,
			SkipToken: nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/GetOrderItemByName.json
func ExampleManagementClient_GetOrderItemByName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetOrderItemByName(ctx,
		"<order-item-name>",
		"<resource-group-name>",
		&armedgeorder.ManagementClientGetOrderItemByNameOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/CreateOrderItem.json
func ExampleManagementClient_BeginCreateOrderItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrderItem(ctx,
		"<order-item-name>",
		"<resource-group-name>",
		armedgeorder.OrderItemResource{
			Location: to.Ptr("<location>"),
			Properties: &armedgeorder.OrderItemProperties{
				AddressDetails: &armedgeorder.AddressDetails{
					ForwardAddress: &armedgeorder.AddressProperties{
						ContactDetails: &armedgeorder.ContactDetails{
							ContactName: to.Ptr("<contact-name>"),
							EmailList: []*string{
								to.Ptr("xxxx@xxxx.xxx")},
							Phone:          to.Ptr("<phone>"),
							PhoneExtension: to.Ptr("<phone-extension>"),
						},
						ShippingAddress: &armedgeorder.ShippingAddress{
							AddressType:     to.Ptr(armedgeorder.AddressTypeNone),
							City:            to.Ptr("<city>"),
							CompanyName:     to.Ptr("<company-name>"),
							Country:         to.Ptr("<country>"),
							PostalCode:      to.Ptr("<postal-code>"),
							StateOrProvince: to.Ptr("<state-or-province>"),
							StreetAddress1:  to.Ptr("<street-address1>"),
							StreetAddress2:  to.Ptr("<street-address2>"),
						},
					},
				},
				OrderID: to.Ptr("<order-id>"),
				OrderItemDetails: &armedgeorder.OrderItemDetails{
					OrderItemType: to.Ptr(armedgeorder.OrderItemTypePurchase),
					Preferences: &armedgeorder.Preferences{
						TransportPreferences: &armedgeorder.TransportPreferences{
							PreferredShipmentType: to.Ptr(armedgeorder.TransportShipmentTypesMicrosoftManaged),
						},
					},
					ProductDetails: &armedgeorder.ProductDetails{
						HierarchyInformation: &armedgeorder.HierarchyInformation{
							ConfigurationName: to.Ptr("<configuration-name>"),
							ProductFamilyName: to.Ptr("<product-family-name>"),
							ProductLineName:   to.Ptr("<product-line-name>"),
							ProductName:       to.Ptr("<product-name>"),
						},
					},
				},
			},
		},
		&armedgeorder.ManagementClientBeginCreateOrderItemOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/DeleteOrderItemByName.json
func ExampleManagementClient_BeginDeleteOrderItemByName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDeleteOrderItemByName(ctx,
		"<order-item-name>",
		"<resource-group-name>",
		&armedgeorder.ManagementClientBeginDeleteOrderItemByNameOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/UpdateOrderItem.json
func ExampleManagementClient_BeginUpdateOrderItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdateOrderItem(ctx,
		"<order-item-name>",
		"<resource-group-name>",
		armedgeorder.OrderItemUpdateParameter{
			Properties: &armedgeorder.OrderItemUpdateProperties{
				Preferences: &armedgeorder.Preferences{
					TransportPreferences: &armedgeorder.TransportPreferences{
						PreferredShipmentType: to.Ptr(armedgeorder.TransportShipmentTypesCustomerManaged),
					},
				},
			},
		},
		&armedgeorder.ManagementClientBeginUpdateOrderItemOptions{IfMatch: nil,
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/CancelOrderItem.json
func ExampleManagementClient_CancelOrderItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.CancelOrderItem(ctx,
		"<order-item-name>",
		"<resource-group-name>",
		armedgeorder.CancellationReason{
			Reason: to.Ptr("<reason>"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ReturnOrderItem.json
func ExampleManagementClient_BeginReturnOrderItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginReturnOrderItem(ctx,
		"<order-item-name>",
		"<resource-group-name>",
		armedgeorder.ReturnOrderItemDetails{
			ReturnReason: to.Ptr("<return-reason>"),
		},
		&armedgeorder.ManagementClientBeginReturnOrderItemOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
