package securityinsight

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// EntityQueryTemplatesClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource
// provider
type EntityQueryTemplatesClient struct {
	BaseClient
}

// NewEntityQueryTemplatesClient creates an instance of the EntityQueryTemplatesClient client.
func NewEntityQueryTemplatesClient(subscriptionID string) EntityQueryTemplatesClient {
	return NewEntityQueryTemplatesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEntityQueryTemplatesClientWithBaseURI creates an instance of the EntityQueryTemplatesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewEntityQueryTemplatesClientWithBaseURI(baseURI string, subscriptionID string) EntityQueryTemplatesClient {
	return EntityQueryTemplatesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets an entity query.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// entityQueryTemplateID - entity query template ID
func (client EntityQueryTemplatesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryTemplateID string) (result EntityQueryTemplateModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EntityQueryTemplatesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.EntityQueryTemplatesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, entityQueryTemplateID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.EntityQueryTemplatesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.EntityQueryTemplatesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.EntityQueryTemplatesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client EntityQueryTemplatesClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryTemplateID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"entityQueryTemplateId": autorest.Encode("path", entityQueryTemplateID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"workspaceName":         autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2022-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entityQueryTemplates/{entityQueryTemplateId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EntityQueryTemplatesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EntityQueryTemplatesClient) GetResponder(resp *http.Response) (result EntityQueryTemplateModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all entity query templates.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// kind - the entity template query kind we want to fetch
func (client EntityQueryTemplatesClient) List(ctx context.Context, resourceGroupName string, workspaceName string, kind string) (result EntityQueryTemplateListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EntityQueryTemplatesClient.List")
		defer func() {
			sc := -1
			if result.eqtl.Response.Response != nil {
				sc = result.eqtl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.EntityQueryTemplatesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName, kind)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.EntityQueryTemplatesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.eqtl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.EntityQueryTemplatesClient", "List", resp, "Failure sending request")
		return
	}

	result.eqtl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.EntityQueryTemplatesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.eqtl.hasNextLink() && result.eqtl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client EntityQueryTemplatesClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string, kind string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2022-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(kind)) > 0 {
		queryParameters["kind"] = autorest.Encode("query", kind)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entityQueryTemplates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client EntityQueryTemplatesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client EntityQueryTemplatesClient) ListResponder(resp *http.Response) (result EntityQueryTemplateList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client EntityQueryTemplatesClient) listNextResults(ctx context.Context, lastResults EntityQueryTemplateList) (result EntityQueryTemplateList, err error) {
	req, err := lastResults.entityQueryTemplateListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityinsight.EntityQueryTemplatesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityinsight.EntityQueryTemplatesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.EntityQueryTemplatesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client EntityQueryTemplatesClient) ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, kind string) (result EntityQueryTemplateListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EntityQueryTemplatesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, workspaceName, kind)
	return
}