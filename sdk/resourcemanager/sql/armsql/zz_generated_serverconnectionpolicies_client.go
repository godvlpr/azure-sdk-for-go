//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ServerConnectionPoliciesClient contains the methods for the ServerConnectionPolicies group.
// Don't use this type directly, use NewServerConnectionPoliciesClient() instead.
type ServerConnectionPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServerConnectionPoliciesClient creates a new instance of ServerConnectionPoliciesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServerConnectionPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServerConnectionPoliciesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ServerConnectionPoliciesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Updates a server connection policy
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// connectionPolicyName - The name of the connection policy.
// parameters - The required parameters for updating a server connection policy.
// options - ServerConnectionPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerConnectionPoliciesClient.BeginCreateOrUpdate
// method.
func (client *ServerConnectionPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, connectionPolicyName ConnectionPolicyName, parameters ServerConnectionPolicy, options *ServerConnectionPoliciesClientBeginCreateOrUpdateOptions) (ServerConnectionPoliciesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, connectionPolicyName, parameters, options)
	if err != nil {
		return ServerConnectionPoliciesClientCreateOrUpdatePollerResponse{}, err
	}
	result := ServerConnectionPoliciesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerConnectionPoliciesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return ServerConnectionPoliciesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ServerConnectionPoliciesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Updates a server connection policy
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServerConnectionPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, connectionPolicyName ConnectionPolicyName, parameters ServerConnectionPolicy, options *ServerConnectionPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, connectionPolicyName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerConnectionPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, connectionPolicyName ConnectionPolicyName, parameters ServerConnectionPolicy, options *ServerConnectionPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies/{connectionPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if connectionPolicyName == "" {
		return nil, errors.New("parameter connectionPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionPolicyName}", url.PathEscape(string(connectionPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Gets a server connection policy
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// connectionPolicyName - The name of the connection policy.
// options - ServerConnectionPoliciesClientGetOptions contains the optional parameters for the ServerConnectionPoliciesClient.Get
// method.
func (client *ServerConnectionPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string, connectionPolicyName ConnectionPolicyName, options *ServerConnectionPoliciesClientGetOptions) (ServerConnectionPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, connectionPolicyName, options)
	if err != nil {
		return ServerConnectionPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerConnectionPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerConnectionPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerConnectionPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, connectionPolicyName ConnectionPolicyName, options *ServerConnectionPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies/{connectionPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if connectionPolicyName == "" {
		return nil, errors.New("parameter connectionPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionPolicyName}", url.PathEscape(string(connectionPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerConnectionPoliciesClient) getHandleResponse(resp *http.Response) (ServerConnectionPoliciesClientGetResponse, error) {
	result := ServerConnectionPoliciesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerConnectionPolicy); err != nil {
		return ServerConnectionPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// ListByServer - Lists connection policy
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// options - ServerConnectionPoliciesClientListByServerOptions contains the optional parameters for the ServerConnectionPoliciesClient.ListByServer
// method.
func (client *ServerConnectionPoliciesClient) ListByServer(resourceGroupName string, serverName string, options *ServerConnectionPoliciesClientListByServerOptions) *ServerConnectionPoliciesClientListByServerPager {
	return &ServerConnectionPoliciesClientListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp ServerConnectionPoliciesClientListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServerConnectionPolicyListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerConnectionPoliciesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerConnectionPoliciesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerConnectionPoliciesClient) listByServerHandleResponse(resp *http.Response) (ServerConnectionPoliciesClientListByServerResponse, error) {
	result := ServerConnectionPoliciesClientListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerConnectionPolicyListResult); err != nil {
		return ServerConnectionPoliciesClientListByServerResponse{}, err
	}
	return result, nil
}
