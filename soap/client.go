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
	"time"
	"workday/client"
	"workday/soap/model"
	"workday/soap/model/change_home_contact_information_request"
	"workday/soap/model/change_legal_name_request"
	"workday/soap/model/get_applicants_request"
	"workday/soap/model/get_applicants_response"
	"workday/soap/model/get_candidates_request"
	"workday/soap/model/get_candidates_response"
	"workday/soap/model/get_workers_request"
	"workday/soap/model/get_workers_response"
	"workday/soap/model/put_applicant_request"
	"workday/soap/model/put_candidate_request"

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

func (s *Client) GetCandidate(permissionReport *model.PermissionCheck, workdayID string) {
	request := get_candidates_request.GetCandidatesRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		RequestReferences: &get_candidates_request.RequestReferences{
			CandidateReference: &get_candidates_request.CandidateReference{
				Id: &get_candidates_request.Id{
					Type:  lo.ToPtr("WID"),
					Value: lo.ToPtr(workdayID),
				},
			},
		},
		ResponseFilter: &get_candidates_request.ResponseFilter{
			Count: lo.ToPtr(1),
		},
	}

	var response get_candidates_response.GetCandidatesResponse
	if call := s.call("Recruiting/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}

	res, error := xml.MarshalIndent(response, "", "  ")
	if error != nil {
		fmt.Printf("SOAP response marshal error: %v\n", error)
		return
	}
	
	fmt.Printf("GetCandidate Response: %s\n", string(res))
}

func (s *Client) GetApplicant(permissionReport *model.PermissionCheck, workdayID string) *get_applicants_response.GetApplicantsResponse {
	request := get_applicants_request.GetApplicantsRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		RequestReferences: &get_applicants_request.RequestReferences{
			ApplicantReference: &get_applicants_request.ApplicantReference{
				Id: &get_applicants_request.Id{
					Type:  lo.ToPtr("WID"),
					Value: lo.ToPtr(workdayID),
				},
			},
		},
		ResponseFilter: &get_applicants_request.ResponseFilter{
			Count: lo.ToPtr(1),
		},
	}

	var response get_applicants_response.GetApplicantsResponse
	if call := s.call("Recruiting/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return nil
	}

	res, error := xml.MarshalIndent(response, "", "  ")
	if error != nil {
		fmt.Printf("SOAP response marshal error: %v\n", error)
		return nil
	}
	
	fmt.Printf("GetApplicant Response: %s\n", string(res))

	return &response
}

func (s *Client) GetWorker(permissionReport *model.PermissionCheck, workdayID string) {
	request := get_workers_request.GetWorkersRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		RequestReferences: &get_workers_request.RequestReferences{
			WorkerReference: &get_workers_request.WorkerReference{
				Id: &get_workers_request.Id{
					Type:  lo.ToPtr("WID"),
					Value: lo.ToPtr(workdayID),
				},
			},
		},
		ResponseFilter: &get_workers_request.ResponseFilter{
			Count: lo.ToPtr(1),
		},
		RequestCriteria: &get_workers_request.RequestCriteria{

		},
	}

	var response get_workers_response.GetWorkersResponse
	if call := s.call("Human_Resources/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}

	res, error := xml.MarshalIndent(response, "", "  ")
	if error != nil {
		fmt.Printf("SOAP response marshal error: %v\n", error)
		return
	}
	
	fmt.Printf("GetWorker Response: %s\n", string(res))

	// permissionReport.CanGetWorkerLegalName = response.ResponseData.Worker.WorkerData.PersonalData.NameData.LegalNameData.NameDetailData.FirstName != nil
	// 	&& response.ResponseData.Worker.WorkerData.PersonalData.NameData.LegalNameData.NameDetailData.LastName != nil
	// permissionReport.CanGetWorkerHomeAddress = response.ResponseData.Worker.WorkerData.PersonalData.ContactData.LocationData.AddressLine1 != nil
	// permissionReport.CanGetWorkerBirthDate = response.ResponseData.Worker.WorkerData.PersonalData.GlobalPersonalInformationData.DateOfBirth != nil
	// permissionReport.CanGetWorkerEmail = response.ResponseData.Worker.WorkerData.PersonalData.ContactData.EmailAddresses != nil
	// permissionReport.CanGetWorkerPhone = response.ResponseData.Worker.WorkerData.PersonalData.ContactData.PhoneNumbers != nil
}

func (s *Client) UpdateWorkerName(permissionReport *model.PermissionCheck, workdayID, firstName, middleName, lastName string) {
	request := change_legal_name_request.ChangeLegalNameRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		ChangeLegalNameData: &change_legal_name_request.ChangeLegalNameData{
			PersonReference: &change_legal_name_request.PersonReference{
				Id: &change_legal_name_request.Id{
					Type: lo.ToPtr("WID"),
					Value: lo.ToPtr(workdayID),
				},
			},
			NameData: &change_legal_name_request.NameData{
				FirstName: lo.ToPtr(firstName),
				MiddleName: lo.ToPtr(middleName),
				LastName: lo.ToPtr(lastName),
			},
			EffectiveDate: lo.ToPtr(time.Now().Format("2006-01-02")),
		},
	}

	var response AnyXML
	if call := s.call("Human_Resources/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}
}

func (s *Client) UpdateWorkerAddress(permissionReport *model.PermissionCheck, workdayID string) {
	request := &change_home_contact_information_request.ChangeHomeContactInformationRequest{

	}

	var response AnyXML
	if call := s.call("Human_Resources/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}
}

func (s *Client) UpdateCandidateName(permissionReport *model.PermissionCheck, workdayID, firstName, middleName, lastName string) {
	request := put_candidate_request.PutCandidateRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		CandidateReference: &put_candidate_request.CandidateReference{
			Id: &put_candidate_request.Id{
				Type:  lo.ToPtr("WID"),
				Value: lo.ToPtr(workdayID),
			},
		},
		CandidateData: &put_candidate_request.CandidateData{
			NameData: &put_candidate_request.NameData{
				LegalName: &put_candidate_request.LegalName{
					NameDetailData: &put_candidate_request.NameDetailData{
						FirstName: lo.ToPtr(firstName),
						// They don;t accept middle names :shrug:....if configured
						LastName: lo.ToPtr(lastName),
					},
				},
			},
			// Can't do these...at least not without a ridiculous amount of tomfoolery.
			// ContactData: &ContactData{
			// 	LocationData: &LocationData{
			// 		AddressLine1: &AddressLine1{
			// 			Value: lo.ToPtr("hello street"),
			// 		},
			// 	},
			// },
			// JobApplicationData: &JobApplicationData{
			// 	JobAppliedToData: &JobAppliedToData{
			// 		GlobalPersonalInformationData: &GlobalPersonalInformationData{
			// 			DateOfBirth: &DateOfBirth{
			// 				Value: lo.ToPtr("1990-02-12"),
			// 			},
			// 		},
			// 	},
			// },
		},
	}

	var response AnyXML
	if call := s.call("Recruiting/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}
}

func (s *Client) UpdateApplicant(permissionReport *model.PermissionCheck, workdayID, firstName, middleName, lastName string) {
	applicant := s.GetApplicant(permissionReport, workdayID)
	if applicant == nil {
		fmt.Printf("Could not find applicant with Workday ID: %s\n", workdayID)
		return
	}
	
	fmt.Println("got applicant: " +
		lo.FromPtr(applicant.ResponseData.Applicant.ApplicantData.PersonalData.NameData.LegalNameData.NameDetailData.FirstName))

	request := &put_applicant_request.PutApplicantRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		ApplicantReference: &put_applicant_request.ApplicantReference{
			Id: &put_applicant_request.Id{
				Type: lo.ToPtr("WID"),
				Value: lo.ToPtr(workdayID),
			},
		},
		ApplicantData: &put_applicant_request.ApplicantData{
			PersonalData: &put_applicant_request.PersonalData{
				NameData: &put_applicant_request.NameData{
					LegalNameData: &put_applicant_request.LegalNameData{
						NameDetailData: &put_applicant_request.NameDetailData{
							FirstName: lo.ToPtr(firstName),
							MiddleName: lo.ToPtr(middleName),
							LastName: lo.ToPtr(lastName),
						},
					},
				},
			},
		},
	}

	var response AnyXML
	if call := s.call("Recruiting/"+WorkdayAPIVersion, request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}
}
