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
				// ContactData: &put_applicant_request.ContactData{
				// 	AddressData: &put_applicant_request.AddressData{
				// 		Add
				// 	},
				// },
				// PersonalInformationData: &put_applicant_request.PersonalInformationData{
				// 	BirthDate: lo.ToPtr(time.Now().AddDate(-5, 0, 0).Format("2006-01-02")),
				// },
			},
		},
	}

	request.ApplicantData.PersonalData.NameData.LegalNameData.NameDetailData.FirstName = lo.ToPtr(firstName)
	request.ApplicantData.PersonalData.NameData.LegalNameData.NameDetailData.MiddleName = lo.ToPtr(middleName)
	request.ApplicantData.PersonalData.NameData.LegalNameData.NameDetailData.LastName = lo.ToPtr(lastName)

	// request := &put_applicant_request.PutApplicantRequest{
	// 	XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
	// 	Version:      lo.ToPtr(WorkdayAPIVersion),
	// 	ApplicantReference: &put_applicant_request.ApplicantReference{
	// 		Id: &put_applicant_request.Id{
	// 			Type:  lo.ToPtr("WID"),
	// 			Value: lo.ToPtr(workdayID),
	// 		},
	// 	},
	// 	ApplicantData: &put_applicant_request.ApplicantData{
	// 		PersonalData: &put_applicant_request.PersonalData{
	// 			NameData: &put_applicant_request.NameData{
	// 				LegalNameData: &put_applicant_request.LegalNameData{
	// 					NameDetailData: &put_applicant_request.NameDetailData{
	// 						FirstName:  lo.ToPtr(firstName),
	// 						MiddleName: lo.ToPtr(middleName),
	// 						LastName:   lo.ToPtr(lastName),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	var response AnyXML
	if call := s.call("Recruiting/"+WorkdayAPIVersion, request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}
}
