// Copyright 2025 Nametag Inc.
//
// All information contained herein is the property of Nametag Inc.. The
// intellectual and technical concepts contained herein are proprietary, trade
// secrets, and/or confidential to Nametag, Inc. and may be covered by U.S.
// and Foreign Patents, patents in process, and are protected by trade secret or
// copyright law. Reproduction or distribution, in whole or in part, is
// forbidden except by express written permission of Nametag, Inc.

package soap

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/google/uuid"
	"github.com/hooklift/gowsdl/soap"
	"github.com/samber/lo"

	workday "github.com/nametaginc/nt/server/workday/client/.dev/client"
	"github.com/nametaginc/nt/server/workday/client/.dev/soap/model/human_resources"
	"github.com/nametaginc/nt/server/workday/client/.dev/soap/model/human_resources/gowsdl"
	"github.com/nametaginc/nt/server/workday/client/.dev/soap/model/human_resources/xgen"
)

type Client struct {
	credentialedClient *http.Client
	apiURL             string
}

func New() (*Client, error) {
	apiURL := os.Getenv("WORKDAY_SOAP_API_URL")

	oauth2Client, err := workday.NewOAUTH2Client()
	if err != nil {
		return nil, err
	}

	return &Client{
		credentialedClient: oauth2Client,
		apiURL:             apiURL,
	}, nil
}

func (s *Client) call(servicePath string, request any, response any) *callResponse {
	serviceURL, err := url.JoinPath(s.apiURL, servicePath)
	if err != nil {
		return (&callResponse{}).WithError(fmt.Errorf("build service URL: %w", err))
	}

	// https://community-content.workday.com/content/workday-community/en-us/kits-and-tools/products/platform-and-product-extensions/integrations/workday-api-headers.html?lang=en-us#accordion-1ac3d0cbad-item-a028c583f1
	headers := map[string]string{
		"wd-external-request-id":     uuid.New().String(),
		"wd-external-application-id": "Nametag",
		"wd-external-originator-id":  "Nametag",
	}

	soapClientImpl := soap.NewClient(serviceURL, soap.WithHTTPClient(s.credentialedClient), soap.WithHTTPHeaders(headers))
	if err := soapClientImpl.Call("''", request, response); err != nil {
		return (&callResponse{}).WithError(fmt.Errorf("call SOAP service: %w", err))
	}

	return &callResponse{
		Error: nil,
	}
}

type callResponse struct {
	Error             error  `json:"error,omitempty"`
	HTTPStatusCode    int    `json:"http_status_code,omitempty"`
	HTTPStatusMessage string `json:"http_status_message,omitempty"`
}

func (s *callResponse) WithError(err error) *callResponse {
	s.Error = err
	return s
}

func (s *callResponse) Success() bool {
	return s.Error == nil && s.HTTPStatusCode == 200
}

func (s *Client) GetProspect(workdayID string) {
	// TODO
}

func (s *Client) GetJobApplication(workdayID string) {
	// TODO
}

func (s *Client) GetPreHire(workdayID string) {
	// TODO
}

func (s *Client) GetWorker(workdayID string) {
	s.GetWorkerCoded(workdayID)
}

func (s *Client) GetWorkerCoded(workdayID string) {
	request := human_resources.NewGetWorkerByWorkdayIDRequest(workdayID)

	req, err := xml.Marshal(request)
	if err != nil {
		fmt.Printf("Error marshalling request: %v\n", err)
		return
	}

	fmt.Printf("Request: %s\n", string(req))

	var response human_resources.GetWorkersResponse
	if call := s.call("Human_Resources/v46.0", request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}

	res, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling response: %v\n", err)
		return
	}

	fmt.Printf("Response: %s\n", string(res))
}

// Problems with gowsdl:
// IDType is a raw string, but it actually has a required attribute "type". (manually fixable):
//
//	type IDType struct {
//		Type  string `xml:"type,attr,omitempty" json:"type,omitempty"`
//		Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
//	}
//
// SOAP request error: call SOAP service: xml: unsupported type: soap.XSDDate
func (s *Client) GetWorkerPersonalInfoGowsdl(workdayID string) {
	request := gowsdl.Employee_Personal_Info_Get{
		Employee_Reference: &gowsdl.Employee_ReferenceType{
			Integration_ID_Reference: &gowsdl.External_Integration_ID_Reference_DataType{
				ID: &gowsdl.IDType{
					Type:  "WID",
					Value: workdayID,
				},
			},
		},
	}

	var response gowsdl.Employee_Personal_Info
	call := s.call("Human_Resources/v46.0", &request, &response)
	if !call.Success() {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}

	fmt.Printf("Received %s personal information help texts\n", response.Employee_Reference.Integration_ID_Reference.ID.Value)
}

// func (s *SOAPClient) GetWorkerGoxmlstruct(workdayID string) {
// 	request := &human_resources.GetWorkersRequest{
// 		XMLNamespace: model.WorkdayXMLNamespace,
// 		RequestReferences: human_resources.RequestReferences{
// 			WorkerReference: human_resources.WorkerReference{
// 				ID: workdayID,
// 				Descriptor: "WID",
// 			},
// 		}
// 	}

// 	req, err := xml.Marshal(request)
// 	if err != nil {
// 		fmt.Printf("Error marshalling request: %v\n", err)
// 		return
// 	}
// 	fmt.Printf("Request: %s\n", string(req))

// 	// var response human_resources.GetWorkersResponse
// 	// call := s.call("Human_Resources/v46.0", &request, &response)
// 	// if !call.Success() {
// 	// 	fmt.Printf("SOAP request error: %v\n", call.Error)
// 	// 	return
// 	// }

// 	// fmt.Printf("Received %s personal information help texts\n", response.Employee_Reference.Integration_ID_Reference.ID.Value)
// }

// Problems with xgen:
// Output code invalid; requires manual fix for a go variable starting with a number (process error on ./Human_Resources.xsd: 13618:2: expected '}', found 24)
func (s *Client) GetWorkerXgen(workdayID string) {
	request := xgen.EmployeePersonalInfoGetType{
		EmployeeReference: &xgen.EmployeeReferenceType{
			IntegrationIDReference: &xgen.ExternalIntegrationIDReferenceDataType{
				ID: &xgen.IDType{
					SystemIDAttr: lo.ToPtr("WID"),
					Value:        workdayID,
				},
			},
		},
	}

	rep, err := xml.Marshal(request)
	if err != nil {
		fmt.Printf("Error marshalling request: %v\n", err)
		return
	}
	fmt.Printf("Request: %s\n", string(rep))

	var response xgen.EmployeePersonalInfoType
	call := s.call("Human_Resources/v46.0", &request, &response)
	if !call.Success() {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}

	fmt.Printf("Received %s personal information help texts\n", response.EmployeeReference.IntegrationIDReference.ID.Value)
}

// // Problems with wsdl2go:
// // Output code invalid; requires manual fix for several fields _{1,2,3,...}
// // A bunch of undefined types, including some we need like change legal name
// func (s *SOAPClient) GetWorkerWsdl2go(workdayID string) {
// 	request := wsdl2go.Employee_Personal_Info_GetType{
// 		Employee_Reference: &wsdl2go.Employee_Personal_Info_GetType{
// 			Integration_ID_Reference: &wsdl2go.External_Integration_ID_Reference_DataType{
// 				ID: &wsdl2go.IDType{
// 					System_ID: lo.ToPtr("WID"),
// 					Content: workdayID,
// 				},
// 			},
// 		},
// 	}
// }
