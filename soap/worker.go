package soap

import (
	"fmt"
	"time"
	"workday/soap/model"
	"workday/soap/model/change_home_contact_information_request"
	"workday/soap/model/change_legal_name_request"
	"workday/soap/model/change_personal_information_request"
	"workday/soap/model/change_personal_information_response"
	"workday/soap/model/employee_personal_info"
	"workday/soap/model/employee_personal_info_get"
	"workday/soap/model/get_workers_request"
	"workday/soap/model/get_workers_response"

	"github.com/samber/lo"
)

func (s *Client) GetEmployee(permissionReport *model.PermissionCheck, workdayID string) {
	request := employee_personal_info_get.EmployeePersonalInfoGet{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		EmployeeReference: &employee_personal_info_get.EmployeeReference{
			IntegrationIdReference: &employee_personal_info_get.IntegrationIdReference{
				Id: &employee_personal_info_get.Id{
					SystemId: lo.ToPtr("WID"),
					Value:    lo.ToPtr(workdayID),
				},
			},
		},
	}

	var response employee_personal_info.EmployeePersonalInfo
	if call := s.call("Human_Resources/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}
}

func (s *Client) GetWorker(permissionReport *model.PermissionCheck, workdayID string) (*get_workers_response.WorkerData, error) {
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
			ExcludeInactiveWorkers:          lo.ToPtr(false),
			ExcludeEmployees:                lo.ToPtr(false),
			ExcludeContingentWorkers:        lo.ToPtr(false),
			IncludeSubordinateOrganizations: lo.ToPtr(true),
		},
		ResponseGroup: &get_workers_request.ResponseGroup{
			IncludeReference:                                   lo.ToPtr(false),
			IncludePersonalInformation:                         lo.ToPtr(true),
			ShowAllPersonalInformation:                         lo.ToPtr(true),
			IncludeAdditionalJobs:                              lo.ToPtr(false),
			IncludeEmploymentInformation:                       lo.ToPtr(false),
			IncludeCompensation:                                lo.ToPtr(false),
			IncludeOrganizations:                               lo.ToPtr(false),
			ExcludeOrganizationSupportRoleData:                 lo.ToPtr(true),
			ExcludeLocationHierarchies:                         lo.ToPtr(true),
			ExcludeCostCenters:                                 lo.ToPtr(true),
			ExcludeCostCenterHierarchies:                       lo.ToPtr(true),
			ExcludeCompanies:                                   lo.ToPtr(true),
			ExcludeCompanyHierarchies:                          lo.ToPtr(true),
			ExcludeMatrixOrganizations:                         lo.ToPtr(true),
			ExcludePayGroups:                                   lo.ToPtr(true),
			ExcludeRegions:                                     lo.ToPtr(true),
			ExcludeRegionHierarchies:                           lo.ToPtr(true),
			ExcludeSupervisoryOrganizations:                    lo.ToPtr(true),
			ExcludeTeams:                                       lo.ToPtr(true),
			ExcludeCustomOrganizations:                         lo.ToPtr(true),
			IncludeRoles:                                       lo.ToPtr(false),
			IncludeManagementChainData:                         lo.ToPtr(false),
			IncludeMultipleManagersInManagementChainData:       lo.ToPtr(false),
			IncludeBenefitEnrollments:                          lo.ToPtr(false),
			IncludeBenefitEligibility:                          lo.ToPtr(false),
			IncludeRelatedPersons:                              lo.ToPtr(false),
			IncludeQualifications:                              lo.ToPtr(false),
			IncludeEmployeeReview:                              lo.ToPtr(false),
			IncludeGoals:                                       lo.ToPtr(false),
			IncludeDevelopmentItems:                            lo.ToPtr(false),
			IncludeSkills:                                      lo.ToPtr(false),
			IncludePhoto:                                       lo.ToPtr(false),
			IncludeWorkerDocuments:                             lo.ToPtr(false),
			IncludeTransactionLogData:                          lo.ToPtr(false),
			IncludeSubeventsForCorrectedTransaction:            lo.ToPtr(false),
			IncludeSubeventsForRescindedTransaction:            lo.ToPtr(false),
			IncludeSuccessionProfile:                           lo.ToPtr(false),
			IncludeTalentAssessment:                            lo.ToPtr(false),
			IncludeEmployeeContractData:                        lo.ToPtr(false),
			IncludeContractsForTerminatedWorkers:               lo.ToPtr(false),
			IncludeCollectiveAgreementData:                     lo.ToPtr(false),
			IncludeProbationPeriodData:                         lo.ToPtr(false),
			IncludeExtendedEmployeeContractDetails:             lo.ToPtr(false),
			IncludeFeedbackReceived:                            lo.ToPtr(false),
			IncludeUserAccount:                                 lo.ToPtr(false),
			IncludeCareer:                                      lo.ToPtr(false),
			IncludeAccountProvisioning:                         lo.ToPtr(false),
			IncludeBackgroundCheckData:                         lo.ToPtr(false),
			IncludeContingentWorkerTaxAuthorityFormInformation: lo.ToPtr(false),
			ExcludeFunds:                                       lo.ToPtr(true),
			ExcludeFundHierarchies:                             lo.ToPtr(true),
			ExcludeGrants:                                      lo.ToPtr(true),
			ExcludeGrantHierarchies:                            lo.ToPtr(true),
			ExcludeBusinessUnits:                               lo.ToPtr(true),
			ExcludeBusinessUnitHierarchies:                     lo.ToPtr(true),
			ExcludePrograms:                                    lo.ToPtr(true),
			ExcludeProgramHierarchies:                          lo.ToPtr(true),
			ExcludeGifts:                                       lo.ToPtr(true),
			ExcludeGiftHierarchies:                             lo.ToPtr(true),
			ExcludeRetireeOrganizations:                        lo.ToPtr(true),
		},
	}

	var response get_workers_response.GetWorkersResponse
	if call := s.call("Human_Resources/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return nil, fmt.Errorf("call SOAP service: %w", call.Error)
	}

	//permissionReport.CanGetWorkerLegalName = response.ResponseData.Worker.WorkerData.PersonalData.NameData.LegalNameData != nil
	//permissionReport.CanGetWorkerBirthDate = response.ResponseData.Worker.WorkerData.PersonalData.PersonalInformationData.BirthDate != nil
	//permissionReport.CanGetWorkerAddress = response.ResponseData.Worker.WorkerData.PersonalData.ContactData.

	return response.ResponseData.Worker.WorkerData, nil
}

func (s *Client) UpdateWorkerName(permissionReport *model.PermissionCheck, workdayID, firstName, middleName, lastName string) {
	request := change_legal_name_request.ChangeLegalNameRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		ChangeLegalNameData: &change_legal_name_request.ChangeLegalNameData{
			PersonReference: &change_legal_name_request.PersonReference{
				Id: &change_legal_name_request.Id{
					Type:  lo.ToPtr("WID"),
					Value: lo.ToPtr(workdayID),
				},
			},
			NameData: &change_legal_name_request.NameData{
				FirstName:  lo.ToPtr(firstName),
				MiddleName: lo.ToPtr(middleName),
				LastName:   lo.ToPtr(lastName),
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
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		ChangeHomeContactInformationData: &change_home_contact_information_request.ChangeHomeContactInformationData{
			PersonReference: &change_home_contact_information_request.PersonReference{
				Id: &change_home_contact_information_request.Id{
					Type:  lo.ToPtr("WID"),
					Value: lo.ToPtr(workdayID),
				},
			},
			PersonContactInformationData: &change_home_contact_information_request.PersonContactInformationData{
				PersonAddressInformationData: &change_home_contact_information_request.PersonAddressInformationData{
					ReplaceAll: lo.ToPtr("false"),
					AddressInformationData: &change_home_contact_information_request.AddressInformationData{
						Delete:        lo.ToPtr("false"),
						AddressData:   &change_home_contact_information_request.AddressData{},
						EffectiveDate: lo.ToPtr(time.Now().Format("2006-01-02")),
					},
				},
			},
		},
	}

	var response AnyXML
	if call := s.call("Human_Resources/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}
}

func (s *Client) UpdateWorkerBirthDate(permissionReport *model.PermissionCheck, workdayID string) {
	request := &change_personal_information_request.ChangePersonalInformationRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		ChangePersonalInformationData: &change_personal_information_request.ChangePersonalInformationData{
			PersonReference: &change_personal_information_request.PersonReference{
				Id: &change_personal_information_request.Id{
					Type:  lo.ToPtr("WID"),
					Value: lo.ToPtr(workdayID),
				},
			},
			PersonalInformationData: &change_personal_information_request.PersonalInformationData{
				DateOfBirth: lo.ToPtr("1990-03-03"),
			},
		},
	}

	var response change_personal_information_response.ChangePersonalInformationResponse
	if call := s.call("Human_Resources/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return
	}
}
