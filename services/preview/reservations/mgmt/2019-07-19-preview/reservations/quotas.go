package reservations

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// QuotasClient is the client for the Quotas methods of the Reservations service.
type QuotasClient struct {
	BaseClient
}

// NewQuotasClient creates an instance of the QuotasClient client.
func NewQuotasClient() QuotasClient {
	return NewQuotasClientWithBaseURI(DefaultBaseURI)
}

// NewQuotasClientWithBaseURI creates an instance of the QuotasClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewQuotasClientWithBaseURI(baseURI string) QuotasClient {
	return QuotasClient{NewWithBaseURI(baseURI)}
}

// ListStatus this API gets the current quota limits and usages for the resource provider for the specified location.
// This response can be used to submit quotaRequests.
// Parameters:
// subscriptionID - azure subscription id.
// providerID - azure resource Provider id.
// location - azure region.
func (client QuotasClient) ListStatus(ctx context.Context, subscriptionID string, providerID string, location string) (result QuotaLimitsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotasClient.ListStatus")
		defer func() {
			sc := -1
			if result.ql.Response.Response != nil {
				sc = result.ql.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listStatusNextResults
	req, err := client.ListStatusPreparer(ctx, subscriptionID, providerID, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotasClient", "ListStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListStatusSender(req)
	if err != nil {
		result.ql.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "reservations.QuotasClient", "ListStatus", resp, "Failure sending request")
		return
	}

	result.ql, err = client.ListStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotasClient", "ListStatus", resp, "Failure responding to request")
	}

	return
}

// ListStatusPreparer prepares the ListStatus request.
func (client QuotasClient) ListStatusPreparer(ctx context.Context, subscriptionID string, providerID string, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"providerId":     autorest.Encode("path", providerID),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2019-07-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimits", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListStatusSender sends the ListStatus request. The method will close the
// http.Response Body if it receives an error.
func (client QuotasClient) ListStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListStatusResponder handles the response to the ListStatus request. The method always
// closes the http.Response Body.
func (client QuotasClient) ListStatusResponder(resp *http.Response) (result QuotaLimits, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listStatusNextResults retrieves the next set of results, if any.
func (client QuotasClient) listStatusNextResults(ctx context.Context, lastResults QuotaLimits) (result QuotaLimits, err error) {
	req, err := lastResults.quotaLimitsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "reservations.QuotasClient", "listStatusNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "reservations.QuotasClient", "listStatusNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotasClient", "listStatusNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListStatusComplete enumerates all values, automatically crossing page boundaries as required.
func (client QuotasClient) ListStatusComplete(ctx context.Context, subscriptionID string, providerID string, location string) (result QuotaLimitsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotasClient.ListStatus")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListStatus(ctx, subscriptionID, providerID, location)
	return
}