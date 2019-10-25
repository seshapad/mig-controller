package costmanagement

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ReportConfigClient is the client for the ReportConfig methods of the Costmanagement service.
type ReportConfigClient struct {
	BaseClient
}

// NewReportConfigClient creates an instance of the ReportConfigClient client.
func NewReportConfigClient(subscriptionID string) ReportConfigClient {
	return NewReportConfigClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewReportConfigClientWithBaseURI creates an instance of the ReportConfigClient client.
func NewReportConfigClientWithBaseURI(baseURI string, subscriptionID string) ReportConfigClient {
	return ReportConfigClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the operation to create or update a report config. Update operation requires latest eTag to be set in
// the request mandatorily. You may obtain the latest eTag by performing a get operation. Create operation does not
// require eTag.
// Parameters:
// reportConfigName - report Config Name.
// parameters - parameters supplied to the CreateOrUpdate Report Config operation.
func (client ReportConfigClient) CreateOrUpdate(ctx context.Context, reportConfigName string, parameters ReportConfig) (result ReportConfig, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReportConfigClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ReportConfigProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Schedule", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Schedule.RecurrencePeriod", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Schedule.RecurrencePeriod.From", Name: validation.Null, Rule: true, Chain: nil}}},
					}},
					{Target: "parameters.ReportConfigProperties.DeliveryInfo", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.DeliveryInfo.Destination", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.DeliveryInfo.Destination.ResourceID", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.ReportConfigProperties.DeliveryInfo.Destination.Container", Name: validation.Null, Rule: true, Chain: nil},
							}},
						}},
					{Target: "parameters.ReportConfigProperties.Definition", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Type", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ReportConfigProperties.Definition.TimePeriod", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.TimePeriod.From", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "parameters.ReportConfigProperties.Definition.TimePeriod.To", Name: validation.Null, Rule: true, Chain: nil},
								}},
							{Target: "parameters.ReportConfigProperties.Definition.Dataset", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Grouping", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Grouping", Name: validation.MaxItems, Rule: 2, Chain: nil}}},
									{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter", Name: validation.Null, Rule: false,
										Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.And", Name: validation.Null, Rule: false,
											Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.And", Name: validation.MinItems, Rule: 2, Chain: nil}}},
											{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Or", Name: validation.Null, Rule: false,
												Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Or", Name: validation.MinItems, Rule: 2, Chain: nil}}},
											{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Not", Name: validation.Null, Rule: false, Chain: nil},
											{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Dimension", Name: validation.Null, Rule: false,
												Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Dimension.Name", Name: validation.Null, Rule: true, Chain: nil},
													{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Dimension.Operator", Name: validation.Null, Rule: true, Chain: nil},
													{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Dimension.Values", Name: validation.Null, Rule: true,
														Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Dimension.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
												}},
											{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Tag", Name: validation.Null, Rule: false,
												Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Tag.Name", Name: validation.Null, Rule: true, Chain: nil},
													{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Tag.Operator", Name: validation.Null, Rule: true, Chain: nil},
													{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Tag.Values", Name: validation.Null, Rule: true,
														Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Tag.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
												}},
										}},
								}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("costmanagement.ReportConfigClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, reportConfigName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ReportConfigClient) CreateOrUpdatePreparer(ctx context.Context, reportConfigName string, parameters ReportConfig) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportConfigName": autorest.Encode("path", reportConfigName),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/reportconfigs/{reportConfigName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ReportConfigClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ReportConfigClient) CreateOrUpdateResponder(resp *http.Response) (result ReportConfig, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateByResourceGroupName the operation to create or update a report config. Update operation requires
// latest eTag to be set in the request mandatorily. You may obtain the latest eTag by performing a get operation.
// Create operation does not require eTag.
// Parameters:
// resourceGroupName - azure Resource Group Name.
// reportConfigName - report Config Name.
// parameters - parameters supplied to the CreateOrUpdate Report Config operation.
func (client ReportConfigClient) CreateOrUpdateByResourceGroupName(ctx context.Context, resourceGroupName string, reportConfigName string, parameters ReportConfig) (result ReportConfig, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReportConfigClient.CreateOrUpdateByResourceGroupName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ReportConfigProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Schedule", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Schedule.RecurrencePeriod", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Schedule.RecurrencePeriod.From", Name: validation.Null, Rule: true, Chain: nil}}},
					}},
					{Target: "parameters.ReportConfigProperties.DeliveryInfo", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.DeliveryInfo.Destination", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.DeliveryInfo.Destination.ResourceID", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.ReportConfigProperties.DeliveryInfo.Destination.Container", Name: validation.Null, Rule: true, Chain: nil},
							}},
						}},
					{Target: "parameters.ReportConfigProperties.Definition", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Type", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ReportConfigProperties.Definition.TimePeriod", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.TimePeriod.From", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "parameters.ReportConfigProperties.Definition.TimePeriod.To", Name: validation.Null, Rule: true, Chain: nil},
								}},
							{Target: "parameters.ReportConfigProperties.Definition.Dataset", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Grouping", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Grouping", Name: validation.MaxItems, Rule: 2, Chain: nil}}},
									{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter", Name: validation.Null, Rule: false,
										Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.And", Name: validation.Null, Rule: false,
											Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.And", Name: validation.MinItems, Rule: 2, Chain: nil}}},
											{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Or", Name: validation.Null, Rule: false,
												Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Or", Name: validation.MinItems, Rule: 2, Chain: nil}}},
											{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Not", Name: validation.Null, Rule: false, Chain: nil},
											{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Dimension", Name: validation.Null, Rule: false,
												Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Dimension.Name", Name: validation.Null, Rule: true, Chain: nil},
													{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Dimension.Operator", Name: validation.Null, Rule: true, Chain: nil},
													{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Dimension.Values", Name: validation.Null, Rule: true,
														Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Dimension.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
												}},
											{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Tag", Name: validation.Null, Rule: false,
												Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Tag.Name", Name: validation.Null, Rule: true, Chain: nil},
													{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Tag.Operator", Name: validation.Null, Rule: true, Chain: nil},
													{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Tag.Values", Name: validation.Null, Rule: true,
														Chain: []validation.Constraint{{Target: "parameters.ReportConfigProperties.Definition.Dataset.Filter.Tag.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
												}},
										}},
								}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("costmanagement.ReportConfigClient", "CreateOrUpdateByResourceGroupName", err.Error())
	}

	req, err := client.CreateOrUpdateByResourceGroupNamePreparer(ctx, resourceGroupName, reportConfigName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "CreateOrUpdateByResourceGroupName", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateByResourceGroupNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "CreateOrUpdateByResourceGroupName", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "CreateOrUpdateByResourceGroupName", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateByResourceGroupNamePreparer prepares the CreateOrUpdateByResourceGroupName request.
func (client ReportConfigClient) CreateOrUpdateByResourceGroupNamePreparer(ctx context.Context, resourceGroupName string, reportConfigName string, parameters ReportConfig) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportConfigName":  autorest.Encode("path", reportConfigName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CostManagement/reportconfigs/{reportConfigName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateByResourceGroupNameSender sends the CreateOrUpdateByResourceGroupName request. The method will close the
// http.Response Body if it receives an error.
func (client ReportConfigClient) CreateOrUpdateByResourceGroupNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateByResourceGroupNameResponder handles the response to the CreateOrUpdateByResourceGroupName request. The method always
// closes the http.Response Body.
func (client ReportConfigClient) CreateOrUpdateByResourceGroupNameResponder(resp *http.Response) (result ReportConfig, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete a report.
// Parameters:
// reportConfigName - report Config Name.
func (client ReportConfigClient) Delete(ctx context.Context, reportConfigName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReportConfigClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, reportConfigName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ReportConfigClient) DeletePreparer(ctx context.Context, reportConfigName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportConfigName": autorest.Encode("path", reportConfigName),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/reportconfigs/{reportConfigName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReportConfigClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ReportConfigClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteByResourceGroupName the operation to delete a report config.
// Parameters:
// resourceGroupName - azure Resource Group Name.
// reportConfigName - report Config Name.
func (client ReportConfigClient) DeleteByResourceGroupName(ctx context.Context, resourceGroupName string, reportConfigName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReportConfigClient.DeleteByResourceGroupName")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteByResourceGroupNamePreparer(ctx, resourceGroupName, reportConfigName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "DeleteByResourceGroupName", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteByResourceGroupNameSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "DeleteByResourceGroupName", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "DeleteByResourceGroupName", resp, "Failure responding to request")
	}

	return
}

// DeleteByResourceGroupNamePreparer prepares the DeleteByResourceGroupName request.
func (client ReportConfigClient) DeleteByResourceGroupNamePreparer(ctx context.Context, resourceGroupName string, reportConfigName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportConfigName":  autorest.Encode("path", reportConfigName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CostManagement/reportconfigs/{reportConfigName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteByResourceGroupNameSender sends the DeleteByResourceGroupName request. The method will close the
// http.Response Body if it receives an error.
func (client ReportConfigClient) DeleteByResourceGroupNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteByResourceGroupNameResponder handles the response to the DeleteByResourceGroupName request. The method always
// closes the http.Response Body.
func (client ReportConfigClient) DeleteByResourceGroupNameResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the report config for a subscription by report config name.
// Parameters:
// reportConfigName - report Config Name.
func (client ReportConfigClient) Get(ctx context.Context, reportConfigName string) (result ReportConfig, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReportConfigClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, reportConfigName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReportConfigClient) GetPreparer(ctx context.Context, reportConfigName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportConfigName": autorest.Encode("path", reportConfigName),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/reportconfigs/{reportConfigName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReportConfigClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReportConfigClient) GetResponder(resp *http.Response) (result ReportConfig, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByResourceGroupName gets the report config for a resource group under a subscription by report config name.
// Parameters:
// resourceGroupName - azure Resource Group Name.
// reportConfigName - report Config Name.
func (client ReportConfigClient) GetByResourceGroupName(ctx context.Context, resourceGroupName string, reportConfigName string) (result ReportConfig, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReportConfigClient.GetByResourceGroupName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByResourceGroupNamePreparer(ctx, resourceGroupName, reportConfigName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "GetByResourceGroupName", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByResourceGroupNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "GetByResourceGroupName", resp, "Failure sending request")
		return
	}

	result, err = client.GetByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "GetByResourceGroupName", resp, "Failure responding to request")
	}

	return
}

// GetByResourceGroupNamePreparer prepares the GetByResourceGroupName request.
func (client ReportConfigClient) GetByResourceGroupNamePreparer(ctx context.Context, resourceGroupName string, reportConfigName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportConfigName":  autorest.Encode("path", reportConfigName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CostManagement/reportconfigs/{reportConfigName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByResourceGroupNameSender sends the GetByResourceGroupName request. The method will close the
// http.Response Body if it receives an error.
func (client ReportConfigClient) GetByResourceGroupNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetByResourceGroupNameResponder handles the response to the GetByResourceGroupName request. The method always
// closes the http.Response Body.
func (client ReportConfigClient) GetByResourceGroupNameResponder(resp *http.Response) (result ReportConfig, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all report configs for a subscription.
func (client ReportConfigClient) List(ctx context.Context) (result ReportConfigListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReportConfigClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReportConfigClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/reportconfigs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReportConfigClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReportConfigClient) ListResponder(resp *http.Response) (result ReportConfigListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupName lists all report configs for a resource group under a subscription.
// Parameters:
// resourceGroupName - azure Resource Group Name.
func (client ReportConfigClient) ListByResourceGroupName(ctx context.Context, resourceGroupName string) (result ReportConfigListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReportConfigClient.ListByResourceGroupName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByResourceGroupNamePreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "ListByResourceGroupName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "ListByResourceGroupName", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportConfigClient", "ListByResourceGroupName", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupNamePreparer prepares the ListByResourceGroupName request.
func (client ReportConfigClient) ListByResourceGroupNamePreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CostManagement/reportconfigs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupNameSender sends the ListByResourceGroupName request. The method will close the
// http.Response Body if it receives an error.
func (client ReportConfigClient) ListByResourceGroupNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByResourceGroupNameResponder handles the response to the ListByResourceGroupName request. The method always
// closes the http.Response Body.
func (client ReportConfigClient) ListByResourceGroupNameResponder(resp *http.Response) (result ReportConfigListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}