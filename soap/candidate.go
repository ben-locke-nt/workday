package soap

import (
	"fmt"
	"time"
	"workday/soap/model"
	"workday/soap/model/get_candidates_request"
	"workday/soap/model/get_candidates_response"
	"workday/soap/model/put_candidate_request"

	"github.com/samber/lo"
)

func (s *Client) GetCandidate(permissionReport *model.PermissionCheck, workdayID string) (*get_candidates_response.CandidateData, error) {
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
			AsOfEffectiveDate: lo.ToPtr(time.Now().Format("2006-01-02")),
			Count:             lo.ToPtr(1),
		},
		// TODO - why does this break validation?
		// RequestCriteria: &get_candidates_request.RequestCriteria{
		// 	InternalWorkersOnly: lo.ToPtr(false),
		// },
		ResponseGroup: &get_candidates_request.ResponseGroup{
			IncludeReference:      lo.ToPtr(false),
			ExcludeAllAttachments: lo.ToPtr(true),
		},
	}

	var response get_candidates_response.GetCandidatesResponse
	if call := s.call("Recruiting/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return nil, fmt.Errorf("call SOAP service: %w", call.Error)
	}

	return response.ResponseData.Candidate.CandidateData, nil
}

func (s *Client) UpdateCandidateName(permissionReport *model.PermissionCheck, workdayID, firstName, middleName, lastName string) error {
	candidate, err := s.GetCandidate(permissionReport, workdayID)
	if err != nil {
		fmt.Printf("Could not find candidate with Workday ID: %s\n", workdayID)
		return err
	}

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
			ContactData: &put_candidate_request.ContactData{
				LocationData: &put_candidate_request.LocationData{
					CountryReference: &put_candidate_request.CountryReference{
						Id: &put_candidate_request.Id{
							Type:  candidate.ContactData.LocationData.CountryReference.Id.Type,
							Value: candidate.ContactData.LocationData.CountryReference.Id.Value,
						},
					},
					AddressLine1: lo.ToPtr("hello street"),
				},
			},
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
		return fmt.Errorf("call SOAP service: %w", call.Error)
	}

	return nil
}
