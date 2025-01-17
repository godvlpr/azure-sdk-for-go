//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VPNSiteLinksClient contains the methods for the VPNSiteLinks group.
// Don't use this type directly, use NewVPNSiteLinksClient() instead.
type VPNSiteLinksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVPNSiteLinksClient creates a new instance of VPNSiteLinksClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVPNSiteLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VPNSiteLinksClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &VPNSiteLinksClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Retrieves the details of a VPN site link.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnSite.
// vpnSiteName - The name of the VpnSite.
// vpnSiteLinkName - The name of the VpnSiteLink being retrieved.
// options - VPNSiteLinksClientGetOptions contains the optional parameters for the VPNSiteLinksClient.Get method.
func (client *VPNSiteLinksClient) Get(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteLinkName string, options *VPNSiteLinksClientGetOptions) (VPNSiteLinksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vpnSiteName, vpnSiteLinkName, options)
	if err != nil {
		return VPNSiteLinksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VPNSiteLinksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VPNSiteLinksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNSiteLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteLinkName string, options *VPNSiteLinksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks/{vpnSiteLinkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnSiteName == "" {
		return nil, errors.New("parameter vpnSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	if vpnSiteLinkName == "" {
		return nil, errors.New("parameter vpnSiteLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteLinkName}", url.PathEscape(vpnSiteLinkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VPNSiteLinksClient) getHandleResponse(resp *http.Response) (VPNSiteLinksClientGetResponse, error) {
	result := VPNSiteLinksClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNSiteLink); err != nil {
		return VPNSiteLinksClientGetResponse{}, err
	}
	return result, nil
}

// ListByVPNSite - Lists all the vpnSiteLinks in a resource group for a vpn site.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnSite.
// vpnSiteName - The name of the VpnSite.
// options - VPNSiteLinksClientListByVPNSiteOptions contains the optional parameters for the VPNSiteLinksClient.ListByVPNSite
// method.
func (client *VPNSiteLinksClient) ListByVPNSite(resourceGroupName string, vpnSiteName string, options *VPNSiteLinksClientListByVPNSiteOptions) *VPNSiteLinksClientListByVPNSitePager {
	return &VPNSiteLinksClientListByVPNSitePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByVPNSiteCreateRequest(ctx, resourceGroupName, vpnSiteName, options)
		},
		advancer: func(ctx context.Context, resp VPNSiteLinksClientListByVPNSiteResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVPNSiteLinksResult.NextLink)
		},
	}
}

// listByVPNSiteCreateRequest creates the ListByVPNSite request.
func (client *VPNSiteLinksClient) listByVPNSiteCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VPNSiteLinksClientListByVPNSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnSiteName == "" {
		return nil, errors.New("parameter vpnSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByVPNSiteHandleResponse handles the ListByVPNSite response.
func (client *VPNSiteLinksClient) listByVPNSiteHandleResponse(resp *http.Response) (VPNSiteLinksClientListByVPNSiteResponse, error) {
	result := VPNSiteLinksClientListByVPNSiteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNSiteLinksResult); err != nil {
		return VPNSiteLinksClientListByVPNSiteResponse{}, err
	}
	return result, nil
}
