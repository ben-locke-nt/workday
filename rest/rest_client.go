// Copyright 2025 Nametag Inc.
//
// All information contained herein is the property of Nametag Inc.. The
// intellectual and technical concepts contained herein are proprietary, trade
// secrets, and/or confidential to Nametag, Inc. and may be covered by U.S.
// and Foreign Patents, patents in process, and are protected by trade secret or
// copyright law. Reproduction or distribution, in whole or in part, is
// forbidden except by express written permission of Nametag, Inc.

package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"workday/client"

	"github.com/gofrs/uuid"
	"github.com/samber/lo"
)

type RESTClient struct {
	oauthClient *http.Client
	apiBaseURL  string
	tenantName  string
}

func NewRESTClient() (*RESTClient, error) {
	tenantName := os.Getenv("WORKDAY_TENANT_NAME")
	hostName := os.Getenv("WORKDAY_REST_API_HOSTNAME")
	apiBaseURL := fmt.Sprintf("https://%s/api", hostName)

	oauth2Client, err := client.NewOAUTH2Client()
	if err != nil {
		return nil, err
	}

	return &RESTClient{
		oauthClient: oauth2Client,
		apiBaseURL:  apiBaseURL,
		tenantName:  tenantName,
	}, nil
}

// https://doc.workday.com/admin-guide/en-us/integrations/workday-rest-api/rest-api-fundamentals/dan1370797985682.html
// REST API format https://{tenantHostname}/api/{serviceName}/{serviceVersion}/{tenant}
func (c *RESTClient) buildURL(serviceName, serviceVersion string, path ...string) string {
	parts := []string{
		url.PathEscape(serviceName),
		url.PathEscape(serviceVersion),
		url.PathEscape(c.tenantName),
	}
	for _, p := range path {
		parts = append(parts, url.PathEscape(p))
	}
	joined, err := url.JoinPath(c.apiBaseURL, parts...)
	if err != nil {
		panic(err)
	}
	return joined
}

func (c *RESTClient) GetAWorker(workdayID string) {
	// https://community.workday.com/sites/default/files/file-hosting/restapi/index.html#staffing/v7/get-/workers
	// Secured by: FLW Service, Self-Service: Current Staffing Information, Worker Data: Public Worker Reports
	// Scope: Staffing
	// Note: If you get “Code S22” or “Permission Denied” error, please make sure the workday account has “Worker Data: Public Worker Reports” domain security policy and the API client has the Staffing functional area. <-- this is what fixed it
	// Note: Workday REST APIs require the Report/Task permissions in the security policy, whether the user is an ISU or an individual user. Here is the admin guide for REST API Security.
	path := c.buildURL("common", "v1", "workers", workdayID)
	res, err := c.oauthClient.Get(path)
	if err != nil {
		log.Printf("failed to get worker: %v", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to read worker response body: %v", err)
		return
	}

	report(res, body)
}

func (c *RESTClient) GetAProspect(workdayID string) {
	// Secured by: Prospects, Set Up: External Career Site Access
	// Scope: Recruiting, Talent Pipeline
	path := c.buildURL("recruiting", "v4", "prospects", workdayID)
	res, err := c.oauthClient.Get(path)
	if err != nil {
		log.Printf("failed to get prospect: %v", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to read prospect response body: %v", err)
		return
	}

	report(res, body)
}

func (c *RESTClient) CreateWorkerCustomObject(customObject *client.NametagWorkerIDVCustomObject) {
	newObjectNametagID, err := uuid.NewV4()
	if err != nil {
		log.Printf("failed to generate UUID for custom object: %v", err)
		return
	}

	newObjectWorkdayIDD, err := uuid.NewV4()
	if err != nil {
		log.Printf("failed to generate UUID for custom object: %v", err)
		return
	}

	customObject.NametagIDVCustomObject.WorkdayID = lo.ToPtr(newObjectWorkdayIDD.String()[:8])
	customObject.NametagIDVCustomObject.NametagID = lo.ToPtr(newObjectNametagID.String()[:8])

	bodyBytes, err := json.Marshal(customObject)
	if err != nil {
		log.Printf("failed to marshal custom object to JSON: %v", err)
		return
	}

	path := c.buildURL("common", "v1", "customObjects", "nametagIdvForWorker")
	res, err := c.oauthClient.Post(path, "application/json", strings.NewReader(string(bodyBytes)))
	if err != nil {
		log.Printf("failed to get prospect: %v", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to read prospect response body: %v", err)
		return
	}

	report(res, body)
}

func (c *RESTClient) GetWorkerCustomObject(workdayIDOfCustomObject string) {
	path := c.buildURL("common", "v1", "customObjects", "nametagIdvForWorker", workdayIDOfCustomObject)
	res, err := c.oauthClient.Get(path)
	if err != nil {
		log.Printf("failed to get prospect: %v", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to read prospect response body: %v", err)
		return
	}

	report(res, body)
}

func (c *RESTClient) GetJobApplicationCustomObject(workdayID string) {
	path := c.buildURL("common", "v1", "prospects", workdayID, "customObjects", "nametagIdvForJobApplication")
	res, err := c.oauthClient.Get(path)
	if err != nil {
		log.Printf("failed to get prospect: %v", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to read prospect response body: %v", err)
		return
	}

	report(res, body)
}

func (c *RESTClient) GetWorkerCustomObjects(workdayID string) {
	// OMG this works: [200](200 OK) GET
	// https://impl-services1.wd12.myworkday.com/ccx/api/common/v1/nametag_dpt1/workers/b3c518ac84321001c1c8ed9c98790000/customObjects/nametagIdvForWorker
	// {"total":1,"data":[{"id":"f5d0dde74ced30010bed10cec0a90000","assertedBirthDate":"1987-12-15","assertedAndVerifiedNamesMatch":"Yes","assertedName":"Ben Locke","verifiedBirthDate":"1986-12-15","assertedName2":"Judson Benton Locke","status":"Done","id":"jd1acc45","worker":{"id":"b3c518ac84321001c1c8ed9c98790000"}}]}
	path := c.buildURL("common", "v1", "workers", workdayID, "customObjects", "nametagIdvForWorker")
	res, err := c.oauthClient.Get(path)
	if err != nil {
		log.Printf("failed to get prospect: %v", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to read prospect response body: %v", err)
		return
	}

	report(res, body)
}

func report(res *http.Response, body []byte) {
	fmt.Printf("[%d](%s) %+v %s : %s\n", res.StatusCode, res.Status, res.Request.Method, res.Request.URL, string(body))
}
