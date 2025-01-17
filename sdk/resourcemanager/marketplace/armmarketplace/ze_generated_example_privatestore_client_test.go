//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/GetPrivateStores.json
func ExamplePrivateStoreClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.List(&armmarketplace.PrivateStoreClientListOptions{UseCache: nil})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/GetPrivateStore.json
func ExamplePrivateStoreClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/PrivateStores_update.json
func ExamplePrivateStoreClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.CreateOrUpdate(ctx,
		"<private-store-id>",
		&armmarketplace.PrivateStoreClientCreateOrUpdateOptions{Payload: &armmarketplace.PrivateStore{
			Properties: &armmarketplace.PrivateStoreProperties{
				Availability: to.Ptr(armmarketplace.AvailabilityDisabled),
				ETag:         to.Ptr("<etag>"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/DeletePrivateStore.json
func ExamplePrivateStoreClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Delete(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/QueryOffers.json
func ExamplePrivateStoreClient_QueryOffers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.QueryOffers(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/BillingAccounts.json
func ExamplePrivateStoreClient_BillingAccounts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.BillingAccounts(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/CollectionsToSubscriptionsMapping.json
func ExamplePrivateStoreClient_CollectionsToSubscriptionsMapping() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CollectionsToSubscriptionsMapping(ctx,
		"<private-store-id>",
		&armmarketplace.PrivateStoreClientCollectionsToSubscriptionsMappingOptions{Payload: &armmarketplace.CollectionsToSubscriptionsMappingPayload{
			Properties: &armmarketplace.CollectionsToSubscriptionsMappingProperties{
				SubscriptionIDs: []*string{
					to.Ptr("b340914e-353d-453a-85fb-8f9b65b51f91"),
					to.Ptr("f2baa04d-5bfc-461b-b6d8-61b403c9ec48")},
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/QueryApprovedPlans.json
func ExamplePrivateStoreClient_QueryApprovedPlans() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.QueryApprovedPlans(ctx,
		"<private-store-id>",
		&armmarketplace.PrivateStoreClientQueryApprovedPlansOptions{Payload: &armmarketplace.QueryApprovedPlansPayload{
			Properties: &armmarketplace.QueryApprovedPlans{
				OfferID: to.Ptr("<offer-id>"),
				PlanIDs: []*string{
					to.Ptr("testPlanA"),
					to.Ptr("testPlanB"),
					to.Ptr("testPlanC")},
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/BulkCollectionsAction.json
func ExamplePrivateStoreClient_BulkCollectionsAction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.BulkCollectionsAction(ctx,
		"<private-store-id>",
		&armmarketplace.PrivateStoreClientBulkCollectionsActionOptions{Payload: &armmarketplace.BulkCollectionsPayload{
			Properties: &armmarketplace.BulkCollectionsDetails{
				Action: to.Ptr("<action>"),
				CollectionIDs: []*string{
					to.Ptr("c752f021-1c37-4af5-b82f-74c51c27b44a"),
					to.Ptr("f47ef1c7-e908-4f39-ae29-db181634ad8d")},
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/GetApprovalRequestsList.json
func ExamplePrivateStoreClient_GetApprovalRequestsList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetApprovalRequestsList(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/GetRequestApproval.json
func ExamplePrivateStoreClient_GetRequestApproval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetRequestApproval(ctx,
		"<private-store-id>",
		"<request-approval-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/CreateApprovalRequest.json
func ExamplePrivateStoreClient_CreateApprovalRequest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateApprovalRequest(ctx,
		"<private-store-id>",
		"<request-approval-id>",
		&armmarketplace.PrivateStoreClientCreateApprovalRequestOptions{Payload: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/QueryRequestApproval.json
func ExamplePrivateStoreClient_QueryRequestApproval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.QueryRequestApproval(ctx,
		"<private-store-id>",
		"<request-approval-id>",
		&armmarketplace.PrivateStoreClientQueryRequestApprovalOptions{Payload: &armmarketplace.QueryRequestApprovalProperties{
			Properties: &armmarketplace.RequestDetails{
				PlanIDs: []*string{
					to.Ptr("testPlanA"),
					to.Ptr("testPlanB"),
					to.Ptr("*")},
				PublisherID: to.Ptr("<publisher-id>"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/AdminRequestApprovalsList.json
func ExamplePrivateStoreClient_AdminRequestApprovalsList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.AdminRequestApprovalsList(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/GetAdminRequestApproval.json
func ExamplePrivateStoreClient_GetAdminRequestApproval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetAdminRequestApproval(ctx,
		"<private-store-id>",
		"<admin-request-approval-id>",
		"<publisher-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/UpdateAdminRequestApproval.json
func ExamplePrivateStoreClient_UpdateAdminRequestApproval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.UpdateAdminRequestApproval(ctx,
		"<private-store-id>",
		"<admin-request-approval-id>",
		&armmarketplace.PrivateStoreClientUpdateAdminRequestApprovalOptions{Payload: &armmarketplace.AdminRequestApprovalsResource{
			Properties: &armmarketplace.AdminRequestApprovalProperties{
				AdminAction: to.Ptr(armmarketplace.AdminActionApproved),
				ApprovedPlans: []*string{
					to.Ptr("testPlan")},
				CollectionIDs: []*string{
					to.Ptr("f8ee227e-85d7-477d-abbf-854d6decaf70"),
					to.Ptr("39246ad6-c521-4fed-8de7-77dede2e873f")},
				Comment:     to.Ptr("<comment>"),
				PublisherID: to.Ptr("<publisher-id>"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/NotificationsState.json
func ExamplePrivateStoreClient_QueryNotificationsState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.QueryNotificationsState(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/AcknowledgeNotification.json
func ExamplePrivateStoreClient_AcknowledgeOfferNotification() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.AcknowledgeOfferNotification(ctx,
		"<private-store-id>",
		"<offer-id>",
		&armmarketplace.PrivateStoreClientAcknowledgeOfferNotificationOptions{Payload: &armmarketplace.AcknowledgeOfferNotificationProperties{
			Properties: &armmarketplace.AcknowledgeOfferNotificationDetails{
				Acknowledge: to.Ptr(false),
				Dismiss:     to.Ptr(false),
				RemoveOffer: to.Ptr(false),
				RemovePlans: []*string{
					to.Ptr("testPlanA")},
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/WithdrawPlan.json
func ExamplePrivateStoreClient_WithdrawPlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.WithdrawPlan(ctx,
		"<private-store-id>",
		"<request-approval-id>",
		&armmarketplace.PrivateStoreClientWithdrawPlanOptions{Payload: &armmarketplace.WithdrawProperties{
			Properties: &armmarketplace.WithdrawDetails{
				PlanID:      to.Ptr("<plan-id>"),
				PublisherID: to.Ptr("<publisher-id>"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/FetchAllSubscriptionsInTenant.json
func ExamplePrivateStoreClient_FetchAllSubscriptionsInTenant() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.FetchAllSubscriptionsInTenant(ctx,
		"<private-store-id>",
		&armmarketplace.PrivateStoreClientFetchAllSubscriptionsInTenantOptions{NextPageToken: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/ListNewPlansNotifications.json
func ExamplePrivateStoreClient_ListNewPlansNotifications() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.ListNewPlansNotifications(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/ListStopSellOffersPlansNotifications.json
func ExamplePrivateStoreClient_ListStopSellOffersPlansNotifications() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.ListStopSellOffersPlansNotifications(ctx,
		"<private-store-id>",
		&armmarketplace.PrivateStoreClientListStopSellOffersPlansNotificationsOptions{StopSellSubscriptions: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/ListSubscriptionsContext.json
func ExamplePrivateStoreClient_ListSubscriptionsContext() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.ListSubscriptionsContext(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
