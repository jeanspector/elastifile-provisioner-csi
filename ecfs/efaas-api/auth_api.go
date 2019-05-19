/* 
 * Elastifile FaaS API
 *
 * Elastifile Filesystem as a Service API
 *
 * OpenAPI spec version: 2.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package EfaasApi

import (
	"net/url"
	"strings"
	"encoding/json"
)

type AuthApi struct {
	Configuration *Configuration
}

func NewAuthApi() *AuthApi {
	configuration := NewConfiguration()
	return &AuthApi{
		Configuration: configuration,
	}
}

func NewAuthApiWithBasePath(basePath string) *AuthApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &AuthApi{
		Configuration: configuration,
	}
}

/**
 * Authenticates using JWT from Google OAuth
 * Validates the Google Oauth token.
 *
 * @param payload 
 * @return *Login
 */
func (a AuthApi) Login(payload Login) (*Login, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/auth/login"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(cloud-console-partner)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// authentication '(cloud-console-partner-autopush)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// authentication '(google-id)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// authentication '(cloud-console-partner-local)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &payload
	var successPayload = new(Login)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "Login", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

