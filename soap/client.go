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
	"workday/client"

	"github.com/google/uuid"
	"github.com/hooklift/gowsdl/soap"
	"github.com/samber/lo"
)

type Client struct {
	credentialedClient *http.Client
	apiURL             string
}

func New() (*Client, error) {
	apiURL := os.Getenv("WORKDAY_SOAP_API_URL")

	oauth2Client, err := client.NewOAUTH2Client()
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

	req, err := xml.MarshalIndent(request, "", "  ")
	if err != nil {
		return (&callResponse{}).WithError(fmt.Errorf("marshal SOAP request: %w", err))
	}

	fmt.Println("=========================================================================")

	fmt.Printf("SOAP Request: %s\n", string(req))

	// https://community-content.workday.com/content/workday-community/en-us/kits-and-tools/products/platform-and-product-extensions/integrations/workday-api-headers.html?lang=en-us#accordion-1ac3d0cbad-item-a028c583f1
	headers := map[string]string{
		"wd-external-request-id":     uuid.New().String(),
		"wd-external-application-id": "Nametag",
		"wd-external-originator-id":  "Nametag-dev",
	}

	soapClientImpl := soap.NewClient(serviceURL, soap.WithHTTPClient(s.credentialedClient), soap.WithHTTPHeaders(headers))
	if err := soapClientImpl.Call("''", request, response); err != nil {
		return (&callResponse{}).WithError(fmt.Errorf("call SOAP service: %w", err))
	}

	res, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		return (&callResponse{}).WithError(fmt.Errorf("marshal SOAP response: %w", err))
	}

	fmt.Println("-------------------------------------------------------------------------")

	fmt.Printf("SOAP Response: %s\n", string(res))

	fmt.Println("=========================================================================")

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

type AnyXML struct {
	Content string `xml:",innerxml"`
}

func (s *Client) GetCandidate(workdayID string) {
	request := GetCandidatesRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version: lo.ToPtr(WorkdayAPIVersion),
		RequestReferences: &RequestReferences{
			CandidateReference: &CandidateReference{
				Id: &Id{
					Type:  lo.ToPtr("WID"),
					Value: lo.ToPtr(workdayID),
				},
			},
		},
		ResponseFilter: &ResponseFilter{
			Count: &Count{
				Value: lo.ToPtr("1"),
			},
		},
	}

	var response AnyXML
	if call := s.call("Recruiting/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}
}

func (s *Client) GetApplicant(workdayID string) {
	request := GetApplicantsRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version: lo.ToPtr(WorkdayAPIVersion),
		RequestReferences: &RequestReferences{
			ApplicantReference: &ApplicantReference{
				Id: &Id{
					Type:  lo.ToPtr("WID"),
					Value: lo.ToPtr(workdayID),
				},
			},
		},
		ResponseFilter: &ResponseFilter{
			Count: &Count{
				Value: lo.ToPtr("1"),
			},
		},
	}

	var response AnyXML
	if call := s.call("Recruiting/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}
}

func (s *Client) GetWorker(workdayID string) {
	request := GetWorkersRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version: lo.ToPtr(WorkdayAPIVersion),
		RequestReferences: &RequestReferences{
			WorkerReference: &WorkerReference{
				Id: &Id{
					Type:  lo.ToPtr("WID"),
					Value: lo.ToPtr(workdayID),
				},
			},
		},
		ResponseFilter: &ResponseFilter{
			Count: &Count{
				Value: lo.ToPtr("1"),
			},
		},
	}

	var response AnyXML
	if call := s.call("Human_Resources/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}
}

func (s *Client) UpdateWorker(workdayID string) {
	// TODO
}

func (s *Client) UpdateCandidate(workdayID string) {
	// TODO
}

func (s *Client) UpdateApplicant(workdayID string) {
	// TODO
}