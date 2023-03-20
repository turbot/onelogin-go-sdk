/*
OneLogin API

OpenAPI Specification for OneLogin

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onelogin

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// BrandingServiceTemplatesApiService BrandingServiceTemplatesApi service
type BrandingServiceTemplatesApiService service

type ApiCreateMessageTemplateRequest struct {
	ctx context.Context
	ApiService *BrandingServiceTemplatesApiService
	brandId int32
	locale *string
	messageTemplate *MessageTemplate
}

// The 2 character language locale for the template. e.g. en &#x3D; English, es &#x3D; Spanish
func (r ApiCreateMessageTemplateRequest) Locale(locale string) ApiCreateMessageTemplateRequest {
	r.locale = &locale
	return r
}

func (r ApiCreateMessageTemplateRequest) MessageTemplate(messageTemplate MessageTemplate) ApiCreateMessageTemplateRequest {
	r.messageTemplate = &messageTemplate
	return r
}

func (r ApiCreateMessageTemplateRequest) Execute() (*MessageTemplate, *http.Response, error) {
	return r.ApiService.CreateMessageTemplateExecute(r)
}

/*
CreateMessageTemplate Create Message Template

Create Message Template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param brandId Unique identifier for the branding object.
 @return ApiCreateMessageTemplateRequest
*/
func (a *BrandingServiceTemplatesApiService) CreateMessageTemplate(ctx context.Context, brandId int32) ApiCreateMessageTemplateRequest {
	return ApiCreateMessageTemplateRequest{
		ApiService: a,
		ctx: ctx,
		brandId: brandId,
	}
}

// Execute executes the request
//  @return MessageTemplate
func (a *BrandingServiceTemplatesApiService) CreateMessageTemplateExecute(r ApiCreateMessageTemplateRequest) (*MessageTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessageTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingServiceTemplatesApiService.CreateMessageTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2/branding/brands/{brand_id}/templates"
	localVarPath = strings.Replace(localVarPath, "{"+"brand_id"+"}", url.PathEscape(parameterValueToString(r.brandId, "brandId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.locale != nil {
		parameterAddToQuery(localVarQueryParams, "locale", r.locale, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.messageTemplate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteMessageTemplateRequest struct {
	ctx context.Context
	ApiService *BrandingServiceTemplatesApiService
	brandId int32
	templateId int32
}

func (r ApiDeleteMessageTemplateRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteMessageTemplateExecute(r)
}

/*
DeleteMessageTemplate Delete Message Template

Delete Message Template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param brandId Unique identifier for the branding object.
 @param templateId Unique identifier for the template to return.
 @return ApiDeleteMessageTemplateRequest
*/
func (a *BrandingServiceTemplatesApiService) DeleteMessageTemplate(ctx context.Context, brandId int32, templateId int32) ApiDeleteMessageTemplateRequest {
	return ApiDeleteMessageTemplateRequest{
		ApiService: a,
		ctx: ctx,
		brandId: brandId,
		templateId: templateId,
	}
}

// Execute executes the request
func (a *BrandingServiceTemplatesApiService) DeleteMessageTemplateExecute(r ApiDeleteMessageTemplateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingServiceTemplatesApiService.DeleteMessageTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2/branding/brands/{brand_id}/templates/{template_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"brand_id"+"}", url.PathEscape(parameterValueToString(r.brandId, "brandId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"template_id"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetMasterByTypeRequest struct {
	ctx context.Context
	ApiService *BrandingServiceTemplatesApiService
	templateType string
}

func (r ApiGetMasterByTypeRequest) Execute() (*MessageTemplate, *http.Response, error) {
	return r.ApiService.GetMasterByTypeExecute(r)
}

/*
GetMasterByType Get Master Template by Type

Get Master Template by Type

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param templateType The message template type to return.
 @return ApiGetMasterByTypeRequest
*/
func (a *BrandingServiceTemplatesApiService) GetMasterByType(ctx context.Context, templateType string) ApiGetMasterByTypeRequest {
	return ApiGetMasterByTypeRequest{
		ApiService: a,
		ctx: ctx,
		templateType: templateType,
	}
}

// Execute executes the request
//  @return MessageTemplate
func (a *BrandingServiceTemplatesApiService) GetMasterByTypeExecute(r ApiGetMasterByTypeRequest) (*MessageTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessageTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingServiceTemplatesApiService.GetMasterByType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2/branding/brands/master/templates/{template_type}"
	localVarPath = strings.Replace(localVarPath, "{"+"template_type"+"}", url.PathEscape(parameterValueToString(r.templateType, "templateType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetMessageTemplateByIdRequest struct {
	ctx context.Context
	ApiService *BrandingServiceTemplatesApiService
	brandId int32
	templateId int32
}

func (r ApiGetMessageTemplateByIdRequest) Execute() (*MessageTemplate, *http.Response, error) {
	return r.ApiService.GetMessageTemplateByIdExecute(r)
}

/*
GetMessageTemplateById Get Message Template

Get Message Template by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param brandId Unique identifier for the branding object.
 @param templateId Unique identifier for the template to return.
 @return ApiGetMessageTemplateByIdRequest
*/
func (a *BrandingServiceTemplatesApiService) GetMessageTemplateById(ctx context.Context, brandId int32, templateId int32) ApiGetMessageTemplateByIdRequest {
	return ApiGetMessageTemplateByIdRequest{
		ApiService: a,
		ctx: ctx,
		brandId: brandId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return MessageTemplate
func (a *BrandingServiceTemplatesApiService) GetMessageTemplateByIdExecute(r ApiGetMessageTemplateByIdRequest) (*MessageTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessageTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingServiceTemplatesApiService.GetMessageTemplateById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2/branding/brands/{brand_id}/templates/{template_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"brand_id"+"}", url.PathEscape(parameterValueToString(r.brandId, "brandId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"template_id"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTemplateByLocaleRequest struct {
	ctx context.Context
	ApiService *BrandingServiceTemplatesApiService
	brandId int32
	templateType string
	locale string
}

func (r ApiGetTemplateByLocaleRequest) Execute() (*MessageTemplate, *http.Response, error) {
	return r.ApiService.GetTemplateByLocaleExecute(r)
}

/*
GetTemplateByLocale Get Template by Type & Locale

Get Template by Type and Locale

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param brandId Unique identifier for the branding object.
 @param templateType The message template type to return.
 @param locale The 2 character language locale for the template. e.g. en = English, es = Spanish
 @return ApiGetTemplateByLocaleRequest
*/
func (a *BrandingServiceTemplatesApiService) GetTemplateByLocale(ctx context.Context, brandId int32, templateType string, locale string) ApiGetTemplateByLocaleRequest {
	return ApiGetTemplateByLocaleRequest{
		ApiService: a,
		ctx: ctx,
		brandId: brandId,
		templateType: templateType,
		locale: locale,
	}
}

// Execute executes the request
//  @return MessageTemplate
func (a *BrandingServiceTemplatesApiService) GetTemplateByLocaleExecute(r ApiGetTemplateByLocaleRequest) (*MessageTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessageTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingServiceTemplatesApiService.GetTemplateByLocale")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2/branding/brands/{brand_id}/templates/{template_type}/{locale}"
	localVarPath = strings.Replace(localVarPath, "{"+"brand_id"+"}", url.PathEscape(parameterValueToString(r.brandId, "brandId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"template_type"+"}", url.PathEscape(parameterValueToString(r.templateType, "templateType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"locale"+"}", url.PathEscape(parameterValueToString(r.locale, "locale")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListMessageTemplatesRequest struct {
	ctx context.Context
	ApiService *BrandingServiceTemplatesApiService
	brandId int32
	locale *string
}

// The 2 character language locale for the template. e.g. en &#x3D; English, es &#x3D; Spanish
func (r ApiListMessageTemplatesRequest) Locale(locale string) ApiListMessageTemplatesRequest {
	r.locale = &locale
	return r
}

func (r ApiListMessageTemplatesRequest) Execute() ([]ListMessageTemplates200ResponseInner, *http.Response, error) {
	return r.ApiService.ListMessageTemplatesExecute(r)
}

/*
ListMessageTemplates List Message Templates

List Message Templates

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param brandId Unique identifier for the branding object.
 @return ApiListMessageTemplatesRequest
*/
func (a *BrandingServiceTemplatesApiService) ListMessageTemplates(ctx context.Context, brandId int32) ApiListMessageTemplatesRequest {
	return ApiListMessageTemplatesRequest{
		ApiService: a,
		ctx: ctx,
		brandId: brandId,
	}
}

// Execute executes the request
//  @return []ListMessageTemplates200ResponseInner
func (a *BrandingServiceTemplatesApiService) ListMessageTemplatesExecute(r ApiListMessageTemplatesRequest) ([]ListMessageTemplates200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListMessageTemplates200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingServiceTemplatesApiService.ListMessageTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2/branding/brands/{brand_id}/templates"
	localVarPath = strings.Replace(localVarPath, "{"+"brand_id"+"}", url.PathEscape(parameterValueToString(r.brandId, "brandId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.locale != nil {
		parameterAddToQuery(localVarQueryParams, "locale", r.locale, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateMessageTemplateByIdRequest struct {
	ctx context.Context
	ApiService *BrandingServiceTemplatesApiService
	brandId int32
	templateId int32
}

func (r ApiUpdateMessageTemplateByIdRequest) Execute() (*MessageTemplate, *http.Response, error) {
	return r.ApiService.UpdateMessageTemplateByIdExecute(r)
}

/*
UpdateMessageTemplateById Update Message Template

Update Message Template by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param brandId Unique identifier for the branding object.
 @param templateId Unique identifier for the template to return.
 @return ApiUpdateMessageTemplateByIdRequest
*/
func (a *BrandingServiceTemplatesApiService) UpdateMessageTemplateById(ctx context.Context, brandId int32, templateId int32) ApiUpdateMessageTemplateByIdRequest {
	return ApiUpdateMessageTemplateByIdRequest{
		ApiService: a,
		ctx: ctx,
		brandId: brandId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return MessageTemplate
func (a *BrandingServiceTemplatesApiService) UpdateMessageTemplateByIdExecute(r ApiUpdateMessageTemplateByIdRequest) (*MessageTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessageTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingServiceTemplatesApiService.UpdateMessageTemplateById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2/branding/brands/{brand_id}/templates/{template_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"brand_id"+"}", url.PathEscape(parameterValueToString(r.brandId, "brandId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"template_id"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateTemplateByLocaleRequest struct {
	ctx context.Context
	ApiService *BrandingServiceTemplatesApiService
	brandId int32
	templateType string
	locale string
}

func (r ApiUpdateTemplateByLocaleRequest) Execute() (*MessageTemplate, *http.Response, error) {
	return r.ApiService.UpdateTemplateByLocaleExecute(r)
}

/*
UpdateTemplateByLocale Update Template by Type & Locale

Update Template by Type and Locale

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param brandId Unique identifier for the branding object.
 @param templateType The message template type to return.
 @param locale The 2 character language locale for the template. e.g. en = English, es = Spanish
 @return ApiUpdateTemplateByLocaleRequest
*/
func (a *BrandingServiceTemplatesApiService) UpdateTemplateByLocale(ctx context.Context, brandId int32, templateType string, locale string) ApiUpdateTemplateByLocaleRequest {
	return ApiUpdateTemplateByLocaleRequest{
		ApiService: a,
		ctx: ctx,
		brandId: brandId,
		templateType: templateType,
		locale: locale,
	}
}

// Execute executes the request
//  @return MessageTemplate
func (a *BrandingServiceTemplatesApiService) UpdateTemplateByLocaleExecute(r ApiUpdateTemplateByLocaleRequest) (*MessageTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessageTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingServiceTemplatesApiService.UpdateTemplateByLocale")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2/branding/brands/{brand_id}/templates/{template_type}/{locale}"
	localVarPath = strings.Replace(localVarPath, "{"+"brand_id"+"}", url.PathEscape(parameterValueToString(r.brandId, "brandId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"template_type"+"}", url.PathEscape(parameterValueToString(r.templateType, "templateType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"locale"+"}", url.PathEscape(parameterValueToString(r.locale, "locale")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v AltErr
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
