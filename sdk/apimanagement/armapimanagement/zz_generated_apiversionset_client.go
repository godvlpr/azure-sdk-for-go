// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// APIVersionSetClient contains the methods for the APIVersionSet group.
// Don't use this type directly, use NewAPIVersionSetClient() instead.
type APIVersionSetClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewAPIVersionSetClient creates a new instance of APIVersionSetClient with the specified values.
func NewAPIVersionSetClient(con *armcore.Connection, subscriptionID string) *APIVersionSetClient {
	return &APIVersionSetClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or Updates a Api Version Set.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APIVersionSetClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, parameters APIVersionSetContract, options *APIVersionSetCreateOrUpdateOptions) (APIVersionSetCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, versionSetID, parameters, options)
	if err != nil {
		return APIVersionSetCreateOrUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return APIVersionSetCreateOrUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return APIVersionSetCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *APIVersionSetClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, parameters APIVersionSetContract, options *APIVersionSetCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *APIVersionSetClient) createOrUpdateHandleResponse(resp *azcore.Response) (APIVersionSetCreateOrUpdateResponse, error) {
	result := APIVersionSetCreateOrUpdateResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.APIVersionSetContract); err != nil {
		return APIVersionSetCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *APIVersionSetClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Deletes specific Api Version Set.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APIVersionSetClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, ifMatch string, options *APIVersionSetDeleteOptions) (APIVersionSetDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, versionSetID, ifMatch, options)
	if err != nil {
		return APIVersionSetDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return APIVersionSetDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return APIVersionSetDeleteResponse{}, client.deleteHandleError(resp)
	}
	return APIVersionSetDeleteResponse{RawResponse: resp.Response}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *APIVersionSetClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, ifMatch string, options *APIVersionSetDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("If-Match", ifMatch)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *APIVersionSetClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the details of the Api Version Set specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APIVersionSetClient) Get(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, options *APIVersionSetGetOptions) (APIVersionSetGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, versionSetID, options)
	if err != nil {
		return APIVersionSetGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return APIVersionSetGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return APIVersionSetGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *APIVersionSetClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, options *APIVersionSetGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APIVersionSetClient) getHandleResponse(resp *azcore.Response) (APIVersionSetGetResponse, error) {
	result := APIVersionSetGetResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.APIVersionSetContract); err != nil {
		return APIVersionSetGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *APIVersionSetClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetEntityTag - Gets the entity state (Etag) version of the Api Version Set specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APIVersionSetClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, options *APIVersionSetGetEntityTagOptions) (APIVersionSetGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, versionSetID, options)
	if err != nil {
		return APIVersionSetGetEntityTagResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return APIVersionSetGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *APIVersionSetClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, options *APIVersionSetGetEntityTagOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *APIVersionSetClient) getEntityTagHandleResponse(resp *azcore.Response) (APIVersionSetGetEntityTagResponse, error) {
	result := APIVersionSetGetEntityTagResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Lists a collection of API Version Sets in the specified service instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APIVersionSetClient) ListByService(resourceGroupName string, serviceName string, options *APIVersionSetListByServiceOptions) APIVersionSetListByServicePager {
	return &apiVersionSetListByServicePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
		},
		advancer: func(ctx context.Context, resp APIVersionSetListByServiceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.APIVersionSetCollection.NextLink)
		},
	}
}

// listByServiceCreateRequest creates the ListByService request.
func (client *APIVersionSetClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *APIVersionSetListByServiceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *APIVersionSetClient) listByServiceHandleResponse(resp *azcore.Response) (APIVersionSetListByServiceResponse, error) {
	result := APIVersionSetListByServiceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.APIVersionSetCollection); err != nil {
		return APIVersionSetListByServiceResponse{}, err
	}
	return result, nil
}

// listByServiceHandleError handles the ListByService error response.
func (client *APIVersionSetClient) listByServiceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Update - Updates the details of the Api VersionSet specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APIVersionSetClient) Update(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, ifMatch string, parameters APIVersionSetUpdateParameters, options *APIVersionSetUpdateOptions) (APIVersionSetUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, versionSetID, ifMatch, parameters, options)
	if err != nil {
		return APIVersionSetUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return APIVersionSetUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return APIVersionSetUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *APIVersionSetClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, ifMatch string, parameters APIVersionSetUpdateParameters, options *APIVersionSetUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("If-Match", ifMatch)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleResponse handles the Update response.
func (client *APIVersionSetClient) updateHandleResponse(resp *azcore.Response) (APIVersionSetUpdateResponse, error) {
	result := APIVersionSetUpdateResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.APIVersionSetContract); err != nil {
		return APIVersionSetUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *APIVersionSetClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}