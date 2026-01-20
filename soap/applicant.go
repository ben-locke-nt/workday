package soap

import (
	"fmt"
	"workday/soap/model"
	"workday/soap/model/get_applicants_request"
	"workday/soap/model/get_applicants_response"
	"workday/soap/model/put_applicant_request"

	"github.com/samber/lo"
)

func (s *Client) GetApplicant(permissionReport *model.PermissionCheck, workdayID string) *get_applicants_response.ApplicantData {
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
		ResponseGroup: &get_applicants_request.ResponseGroup{
			IncludeReference:                 lo.ToPtr(false),
			IncludePersonalInformation:       lo.ToPtr(true),
			ShowAllPersonalInformation:       lo.ToPtr(true),
			IncludeRecruitingInformation:     lo.ToPtr(false),
			IncludeQualificationProfile:      lo.ToPtr(false),
			IncludeResume:                    lo.ToPtr(false),
			IncludeBackgroundCheck:           lo.ToPtr(false),
			IncludeExternalIntegrationIdData: lo.ToPtr(true),
		},
	}

	var response get_applicants_response.GetApplicantsResponse
	if call := s.call("Recruiting/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return nil
	}

	return response.ResponseData.Applicant.ApplicantData
}

func (s *Client) UpdateApplicant(permissionReport *model.PermissionCheck, workdayID, firstName, middleName, lastName string) {
	applicant := s.GetApplicant(permissionReport, workdayID)
	if applicant == nil {
		fmt.Printf("Could not find applicant with Workday ID: %s\n", workdayID)
		return
	}

	fmt.Println("got applicant: " +
		lo.FromPtr(applicant.PersonalData.NameData.LegalNameData.NameDetailData.FirstName))

	request := &put_applicant_request.PutApplicantRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		ApplicantReference: &put_applicant_request.ApplicantReference{
			Id: &put_applicant_request.Id{
				Type:  lo.ToPtr("WID"),
				Value: lo.ToPtr(workdayID),
			},
		},
		ApplicantData: &put_applicant_request.ApplicantData{
			PersonalData: &put_applicant_request.PersonalData{
				NameData: &put_applicant_request.NameData{
					LegalNameData: &put_applicant_request.LegalNameData{
						NameDetailData: &put_applicant_request.NameDetailData{
							FirstName:  lo.ToPtr(firstName),
							MiddleName: lo.ToPtr(middleName),
							LastName:   lo.ToPtr(lastName),
							CountryReference: &put_applicant_request.CountryReference{
								Id: &put_applicant_request.Id{
									Type:  applicant.PersonalData.NameData.LegalNameData.NameDetailData.CountryReference.Id.Type,
									Value: applicant.PersonalData.NameData.LegalNameData.NameDetailData.CountryReference.Id.Value,
								},
							},
						},
					},
				},
				// PersonalInformationData: &put_applicant_request.PersonalInformationData{
				// 	// Error received:
				// 	// The field Date of Birth is not tracked for the specified Location Context.
				// 	// The Location Context is derived from the Country of the Location for the Position.
				// 	//BirthDate: lo.ToPtr("2000-04-01"),
				// },
				ContactData: &put_applicant_request.ContactData{
					AddressData: &put_applicant_request.AddressData{
						DoNotReplaceAll: lo.ToPtr("true"),
						Delete: lo.ToPtr("false"),
						CountryReference: &put_applicant_request.CountryReference{
							Id: &put_applicant_request.Id{
								Type: applicant.PersonalData.ContactData.AddressData.CountryReference.Id.Type,
								Value: applicant.PersonalData.ContactData.AddressData.CountryReference.Id.Value,
							},
						},
						Municipality: applicant.PersonalData.ContactData.AddressData.Municipality,
						AddressLineData: &put_applicant_request.AddressLineData{
							Descriptor: applicant.PersonalData.ContactData.AddressData.AddressLineData.Descriptor,
							Type: applicant.PersonalData.ContactData.AddressData.AddressLineData.Type,
							Value: lo.ToPtr("1234 Saint James Place"),
						},
						PostalCode: applicant.PersonalData.ContactData.AddressData.PostalCode,
						UsageData: &put_applicant_request.UsageData{
							Public: lo.ToPtr("0"),
							TypeData: &put_applicant_request.TypeData{
								Primary: lo.ToPtr("1"),
								TypeReference: &put_applicant_request.TypeReference{
									Id: &put_applicant_request.Id{
										Type: lo.ToPtr("Communication_Usage_Type_ID"),
										Value: lo.ToPtr("HOME"),
									},
								},
							},
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
