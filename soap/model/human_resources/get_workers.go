// Copyright 2026 Nametag Inc.
//
// All information contained herein is the property of Nametag Inc.. The
// intellectual and technical concepts contained herein are proprietary, trade
// secrets, and/or confidential to Nametag, Inc. and may be covered by U.S.
// and Foreign Patents, patents in process, and are protected by trade secret or
// copyright law. Reproduction or distribution, in whole or in part, is
// forbidden except by express written permission of Nametag, Inc.

package human_resources

import (
	"encoding/xml"

	"github.com/nametaginc/nt/server/workday/client/.dev/soap/model"
)

type GetWorkersRequest struct {
	XMLName xml.Name `xml:"wd:Get_Workers_Request,omitempty"`
	*model.XMLTopLevel
	References *model.WorkdayRequestReferencesList
	Pagination *model.PaginationFilter
}

type WorkerReference struct {
	XMLName   xml.Name `xml:"wd:Worker_Reference,omitempty"`
	Reference *model.WorkdayObjectID
}

func NewGetWorkersRequest() *GetWorkersRequest {
	return &GetWorkersRequest{
		XMLTopLevel: model.NewXMLTopLevel(HumanResourcesServiceVersion),
	}
}

func NewGetWorkerByWorkdayIDRequest(workdayID string) *GetWorkersRequest {
	request := NewGetWorkersRequest()
	request.References = &model.WorkdayRequestReferencesList{
		References: []any{
			&WorkerReference{
				Reference: model.NewWorkdayObjectID(workdayID),
			},
		},
	}
	request.Pagination = &model.PaginationFilter{
		Count: 1,
	}
	return request
}


type GetWorkersResponse struct {
	XMLName           xml.Name `xml:"Get_Workers_Response,omitempty"`
	Text              string   `xml:",chardata"`
	Version           string   `xml:"version,attr,omitempty"`
	Bsvc              string   `xml:"bsvc,attr,omitempty"`
	RequestReferences struct {
		Text                     string `xml:",chardata"`
		SkipNonExistingInstances string `xml:"Skip_Non_Existing_Instances,attr,omitempty"`
		IgnoreInvalidReferences  string `xml:"Ignore_Invalid_References,attr,omitempty"`
		WorkerReference          struct {
			Text       string `xml:",chardata"`
			Descriptor string `xml:"Descriptor,attr,omitempty"`
			ID         struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr,omitempty"`
			} `xml:"ID,omitempty"`
		} `xml:"Worker_Reference,omitempty"`
	} `xml:"Request_References,omitempty"`
	RequestCriteria struct {
		Text                       string `xml:",chardata"`
		TransactionLogCriteriaData struct {
			Text                     string `xml:",chardata"`
			TransactionDateRangeData struct {
				Text             string `xml:",chardata"`
				UpdatedFrom      string `xml:"Updated_From,omitempty"`
				UpdatedThrough   string `xml:"Updated_Through,omitempty"`
				EffectiveFrom    string `xml:"Effective_From,omitempty"`
				EffectiveThrough string `xml:"Effective_Through,omitempty"`
			} `xml:"Transaction_Date_Range_Data,omitempty"`
			TransactionTypeReferences struct {
				Text                     string `xml:",chardata"`
				TransactionTypeReference struct {
					Text       string `xml:",chardata"`
					Descriptor string `xml:"Descriptor,attr,omitempty"`
					ID         struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr,omitempty"`
					} `xml:"ID,omitempty"`
				} `xml:"Transaction_Type_Reference,omitempty"`
			} `xml:"Transaction_Type_References,omitempty"`
			SubscriberReference struct {
				Text       string `xml:",chardata"`
				Descriptor string `xml:"Descriptor,attr,omitempty"`
				ID         struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr,omitempty"`
				} `xml:"ID,omitempty"`
			} `xml:"Subscriber_Reference,omitempty"`
		} `xml:"Transaction_Log_Criteria_Data,omitempty"`
		OrganizationReference struct {
			Text       string `xml:",chardata"`
			Descriptor string `xml:"Descriptor,attr,omitempty"`
			ID         struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr,omitempty"`
			} `xml:"ID,omitempty"`
		} `xml:"Organization_Reference,omitempty"`
		CountryReference struct {
			Text       string `xml:",chardata"`
			Descriptor string `xml:"Descriptor,attr,omitempty"`
			ID         struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr,omitempty"`
			} `xml:"ID,omitempty"`
		} `xml:"Country_Reference,omitempty"`
		IncludeSubordinateOrganizations string `xml:"Include_Subordinate_Organizations,omitempty"`
		PositionReference               struct {
			Text       string `xml:",chardata"`
			Descriptor string `xml:"Descriptor,attr,omitempty"`
			ID         struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr,omitempty"`
			} `xml:"ID,omitempty"`
		} `xml:"Position_Reference,omitempty"`
		EventReference struct {
			Text       string `xml:",chardata"`
			Descriptor string `xml:"Descriptor,attr,omitempty"`
			ID         struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr,omitempty"`
			} `xml:"ID,omitempty"`
		} `xml:"Event_Reference,omitempty"`
		BenefitPlanReference struct {
			Text       string `xml:",chardata"`
			Descriptor string `xml:"Descriptor,attr,omitempty"`
			ID         struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr,omitempty"`
			} `xml:"ID,omitempty"`
		} `xml:"Benefit_Plan_Reference,omitempty"`
		FieldAndParameterCriteriaData struct {
			Text              string `xml:",chardata"`
			ProviderReference struct {
				Text       string `xml:",chardata"`
				Descriptor string `xml:"Descriptor,attr,omitempty"`
				ID         struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr,omitempty"`
				} `xml:"ID,omitempty"`
			} `xml:"Provider_Reference,omitempty"`
		} `xml:"Field_And_Parameter_Criteria_Data,omitempty"`
		NationalIDCriteriaData struct {
			Text                    string `xml:",chardata"`
			IdentifierID            string `xml:"Identifier_ID,omitempty"`
			NationalIDTypeReference struct {
				Text       string `xml:",chardata"`
				Descriptor string `xml:"Descriptor,attr,omitempty"`
				ID         struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr,omitempty"`
				} `xml:"ID,omitempty"`
			} `xml:"National_ID_Type_Reference,omitempty"`
			CountryReference struct {
				Text       string `xml:",chardata"`
				Descriptor string `xml:"Descriptor,attr,omitempty"`
				ID         struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr,omitempty"`
				} `xml:"ID,omitempty"`
			} `xml:"Country_Reference,omitempty"`
		} `xml:"National_ID_Criteria_Data,omitempty"`
		ExcludeInactiveWorkers   string `xml:"Exclude_Inactive_Workers,omitempty"`
		ExcludeEmployees         string `xml:"Exclude_Employees,omitempty"`
		ExcludeContingentWorkers string `xml:"Exclude_Contingent_Workers,omitempty"`
		EligibilityCriteriaData  struct {
			Text           string `xml:",chardata"`
			FieldReference struct {
				Text       string `xml:",chardata"`
				Descriptor string `xml:"Descriptor,attr,omitempty"`
				ID         struct {
					Text       string `xml:",chardata"`
					Type       string `xml:"type,attr,omitempty"`
					ParentID   string `xml:"parent_id,attr,omitempty"`
					ParentType string `xml:"parent_type,attr,omitempty"`
				} `xml:"ID,omitempty"`
			} `xml:"Field_Reference,omitempty"`
			AsOfEffectiveDate string `xml:"As_Of_Effective_Date,omitempty"`
			AsOfEntryDateTime string `xml:"As_Of_Entry_DateTime,omitempty"`
		} `xml:"Eligibility_Criteria_Data,omitempty"`
		UniversalIDReference struct {
			Text       string `xml:",chardata"`
			Descriptor string `xml:"Descriptor,attr,omitempty"`
			ID         struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr,omitempty"`
			} `xml:"ID,omitempty"`
		} `xml:"Universal_ID_Reference,omitempty"`
	} `xml:"Request_Criteria,omitempty"`
	ResponseFilter struct {
		Text              string `xml:",chardata"`
		AsOfEffectiveDate string `xml:"As_Of_Effective_Date,omitempty"`
		AsOfEntryDateTime string `xml:"As_Of_Entry_DateTime,omitempty"`
		Page              string `xml:"Page,omitempty"`
		Count             string `xml:"Count,omitempty"`
	} `xml:"Response_Filter,omitempty"`
	ResponseGroup struct {
		Text                                               string `xml:",chardata"`
		IncludeReference                                   string `xml:"Include_Reference,omitempty"`
		IncludePersonalInformation                         string `xml:"Include_Personal_Information,omitempty"`
		ShowAllPersonalInformation                         string `xml:"Show_All_Personal_Information,omitempty"`
		IncludeAdditionalJobs                              string `xml:"Include_Additional_Jobs,omitempty"`
		IncludeEmploymentInformation                       string `xml:"Include_Employment_Information,omitempty"`
		IncludeCompensation                                string `xml:"Include_Compensation,omitempty"`
		IncludeOrganizations                               string `xml:"Include_Organizations,omitempty"`
		ExcludeOrganizationSupportRoleData                 string `xml:"Exclude_Organization_Support_Role_Data,omitempty"`
		ExcludeLocationHierarchies                         string `xml:"Exclude_Location_Hierarchies,omitempty"`
		ExcludeCostCenters                                 string `xml:"Exclude_Cost_Centers,omitempty"`
		ExcludeCostCenterHierarchies                       string `xml:"Exclude_Cost_Center_Hierarchies,omitempty"`
		ExcludeCompanies                                   string `xml:"Exclude_Companies,omitempty"`
		ExcludeCompanyHierarchies                          string `xml:"Exclude_Company_Hierarchies,omitempty"`
		ExcludeMatrixOrganizations                         string `xml:"Exclude_Matrix_Organizations,omitempty"`
		ExcludePayGroups                                   string `xml:"Exclude_Pay_Groups,omitempty"`
		ExcludeRegions                                     string `xml:"Exclude_Regions,omitempty"`
		ExcludeRegionHierarchies                           string `xml:"Exclude_Region_Hierarchies,omitempty"`
		ExcludeSupervisoryOrganizations                    string `xml:"Exclude_Supervisory_Organizations,omitempty"`
		ExcludeTeams                                       string `xml:"Exclude_Teams,omitempty"`
		ExcludeCustomOrganizations                         string `xml:"Exclude_Custom_Organizations,omitempty"`
		IncludeRoles                                       string `xml:"Include_Roles,omitempty"`
		IncludeManagementChainData                         string `xml:"Include_Management_Chain_Data,omitempty"`
		IncludeMultipleManagersInManagementChainData       string `xml:"Include_Multiple_Managers_in_Management_Chain_Data,omitempty"`
		IncludeBenefitEnrollments                          string `xml:"Include_Benefit_Enrollments,omitempty"`
		IncludeBenefitEligibility                          string `xml:"Include_Benefit_Eligibility,omitempty"`
		IncludeRelatedPersons                              string `xml:"Include_Related_Persons,omitempty"`
		IncludeQualifications                              string `xml:"Include_Qualifications,omitempty"`
		IncludeEmployeeReview                              string `xml:"Include_Employee_Review,omitempty"`
		IncludeGoals                                       string `xml:"Include_Goals,omitempty"`
		IncludeDevelopmentItems                            string `xml:"Include_Development_Items,omitempty"`
		IncludeSkills                                      string `xml:"Include_Skills,omitempty"`
		IncludePhoto                                       string `xml:"Include_Photo,omitempty"`
		IncludeWorkerDocuments                             string `xml:"Include_Worker_Documents,omitempty"`
		IncludeTransactionLogData                          string `xml:"Include_Transaction_Log_Data,omitempty"`
		IncludeSubeventsForCorrectedTransaction            string `xml:"Include_Subevents_for_Corrected_Transaction,omitempty"`
		IncludeSubeventsForRescindedTransaction            string `xml:"Include_Subevents_for_Rescinded_Transaction,omitempty"`
		IncludeSuccessionProfile                           string `xml:"Include_Succession_Profile,omitempty"`
		IncludeTalentAssessment                            string `xml:"Include_Talent_Assessment,omitempty"`
		IncludeEmployeeContractData                        string `xml:"Include_Employee_Contract_Data,omitempty"`
		IncludeContractsForTerminatedWorkers               string `xml:"Include_Contracts_for_Terminated_Workers,omitempty"`
		IncludeCollectiveAgreementData                     string `xml:"Include_Collective_Agreement_Data,omitempty"`
		IncludeProbationPeriodData                         string `xml:"Include_Probation_Period_Data,omitempty"`
		IncludeExtendedEmployeeContractDetails             string `xml:"Include_Extended_Employee_Contract_Details,omitempty"`
		IncludeFeedbackReceived                            string `xml:"Include_Feedback_Received,omitempty"`
		IncludeUserAccount                                 string `xml:"Include_User_Account,omitempty"`
		IncludeCareer                                      string `xml:"Include_Career,omitempty"`
		IncludeAccountProvisioning                         string `xml:"Include_Account_Provisioning,omitempty"`
		IncludeBackgroundCheckData                         string `xml:"Include_Background_Check_Data,omitempty"`
		IncludeContingentWorkerTaxAuthorityFormInformation string `xml:"Include_Contingent_Worker_Tax_Authority_Form_Information,omitempty"`
		ExcludeFunds                                       string `xml:"Exclude_Funds,omitempty"`
		ExcludeFundHierarchies                             string `xml:"Exclude_Fund_Hierarchies,omitempty"`
		ExcludeGrants                                      string `xml:"Exclude_Grants,omitempty"`
		ExcludeGrantHierarchies                            string `xml:"Exclude_Grant_Hierarchies,omitempty"`
		ExcludeBusinessUnits                               string `xml:"Exclude_Business_Units,omitempty"`
		ExcludeBusinessUnitHierarchies                     string `xml:"Exclude_Business_Unit_Hierarchies,omitempty"`
		ExcludePrograms                                    string `xml:"Exclude_Programs,omitempty"`
		ExcludeProgramHierarchies                          string `xml:"Exclude_Program_Hierarchies,omitempty"`
		ExcludeGifts                                       string `xml:"Exclude_Gifts,omitempty"`
		ExcludeGiftHierarchies                             string `xml:"Exclude_Gift_Hierarchies,omitempty"`
		ExcludeRetireeOrganizations                        string `xml:"Exclude_Retiree_Organizations,omitempty"`
	} `xml:"Response_Group,omitempty"`
	ResponseResults struct {
		Text         string `xml:",chardata"`
		TotalResults string `xml:"Total_Results,omitempty"`
		TotalPages   string `xml:"Total_Pages,omitempty"`
		PageResults  string `xml:"Page_Results,omitempty"`
		Page         string `xml:"Page,omitempty"`
	} `xml:"Response_Results,omitempty"`
	ResponseData struct {
		Text   string `xml:",chardata"`
		Worker struct {
			Text            string `xml:",chardata"`
			WorkerReference struct {
				Text       string `xml:",chardata"`
				Descriptor string `xml:"Descriptor,attr,omitempty"`
				ID         struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr,omitempty"`
				} `xml:"ID,omitempty"`
			} `xml:"Worker_Reference,omitempty"`
			WorkerDescriptor             string `xml:"Worker_Descriptor,omitempty"`
			UniversalIdentifierReference struct {
				Text       string `xml:",chardata"`
				Descriptor string `xml:"Descriptor,attr,omitempty"`
				ID         struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr,omitempty"`
				} `xml:"ID,omitempty"`
			} `xml:"Universal_Identifier_Reference,omitempty"`
			WorkerData struct {
				Text         string `xml:",chardata"`
				WorkerID     string `xml:"Worker_ID,omitempty"`
				UserID       string `xml:"User_ID,omitempty"`
				UniversalID  string `xml:"Universal_ID,omitempty"`
				PersonalData struct {
					Text        string `xml:",chardata"`
					UniversalID string `xml:"Universal_ID,omitempty"`
					NameData    struct {
						Text          string `xml:",chardata"`
						LegalNameData struct {
							Text           string `xml:",chardata"`
							NameDetailData struct {
								Text             string `xml:",chardata"`
								FormattedName    string `xml:"Formatted_Name,attr,omitempty"`
								ReportingName    string `xml:"Reporting_Name,attr,omitempty"`
								CountryReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Country_Reference,omitempty"`
								PrefixData struct {
									Text           string `xml:",chardata"`
									TitleReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Title_Reference,omitempty"`
									TitleDescriptor     string `xml:"Title_Descriptor,omitempty"`
									SalutationReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Salutation_Reference,omitempty"`
								} `xml:"Prefix_Data,omitempty"`
								FirstName           string `xml:"First_Name,omitempty"`
								MiddleName          string `xml:"Middle_Name,omitempty"`
								LastName            string `xml:"Last_Name,omitempty"`
								SecondaryLastName   string `xml:"Secondary_Last_Name,omitempty"`
								TertiaryLastName    string `xml:"Tertiary_Last_Name,omitempty"`
								LocalNameDetailData struct {
									Text               string `xml:",chardata"`
									LocalName          string `xml:"Local_Name,attr,omitempty"`
									LocalScript        string `xml:"Local_Script,attr,omitempty"`
									FirstName          string `xml:"First_Name,omitempty"`
									MiddleName         string `xml:"Middle_Name,omitempty"`
									LastName           string `xml:"Last_Name,omitempty"`
									SecondaryLastName  string `xml:"Secondary_Last_Name,omitempty"`
									FirstName2         string `xml:"First_Name_2,omitempty"`
									MiddleName2        string `xml:"Middle_Name_2,omitempty"`
									LastName2          string `xml:"Last_Name_2,omitempty"`
									SecondaryLastName2 string `xml:"Secondary_Last_Name_2,omitempty"`
								} `xml:"Local_Name_Detail_Data,omitempty"`
								SuffixData struct {
									Text                  string `xml:",chardata"`
									SocialSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Social_Suffix_Reference,omitempty"`
									SocialSuffixDescriptor  string `xml:"Social_Suffix_Descriptor,omitempty"`
									AcademicSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Academic_Suffix_Reference,omitempty"`
									HereditarySuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Hereditary_Suffix_Reference,omitempty"`
									HonorarySuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Honorary_Suffix_Reference,omitempty"`
									ProfessionalSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Professional_Suffix_Reference,omitempty"`
									ReligiousSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Religious_Suffix_Reference,omitempty"`
									RoyalSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Royal_Suffix_Reference,omitempty"`
								} `xml:"Suffix_Data,omitempty"`
								FullNameForSingaporeAndMalaysia string `xml:"Full_Name_for_Singapore_and_Malaysia,omitempty"`
							} `xml:"Name_Detail_Data,omitempty"`
						} `xml:"Legal_Name_Data,omitempty"`
						PreferredNameData struct {
							Text           string `xml:",chardata"`
							NameDetailData struct {
								Text             string `xml:",chardata"`
								FormattedName    string `xml:"Formatted_Name,attr,omitempty"`
								ReportingName    string `xml:"Reporting_Name,attr,omitempty"`
								CountryReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Country_Reference,omitempty"`
								PrefixData struct {
									Text           string `xml:",chardata"`
									TitleReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Title_Reference,omitempty"`
									TitleDescriptor     string `xml:"Title_Descriptor,omitempty"`
									SalutationReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Salutation_Reference,omitempty"`
								} `xml:"Prefix_Data,omitempty"`
								FirstName           string `xml:"First_Name,omitempty"`
								MiddleName          string `xml:"Middle_Name,omitempty"`
								LastName            string `xml:"Last_Name,omitempty"`
								SecondaryLastName   string `xml:"Secondary_Last_Name,omitempty"`
								TertiaryLastName    string `xml:"Tertiary_Last_Name,omitempty"`
								LocalNameDetailData struct {
									Text               string `xml:",chardata"`
									LocalName          string `xml:"Local_Name,attr,omitempty"`
									LocalScript        string `xml:"Local_Script,attr,omitempty"`
									FirstName          string `xml:"First_Name,omitempty"`
									MiddleName         string `xml:"Middle_Name,omitempty"`
									LastName           string `xml:"Last_Name,omitempty"`
									SecondaryLastName  string `xml:"Secondary_Last_Name,omitempty"`
									FirstName2         string `xml:"First_Name_2,omitempty"`
									MiddleName2        string `xml:"Middle_Name_2,omitempty"`
									LastName2          string `xml:"Last_Name_2,omitempty"`
									SecondaryLastName2 string `xml:"Secondary_Last_Name_2,omitempty"`
								} `xml:"Local_Name_Detail_Data,omitempty"`
								SuffixData struct {
									Text                  string `xml:",chardata"`
									SocialSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Social_Suffix_Reference,omitempty"`
									SocialSuffixDescriptor  string `xml:"Social_Suffix_Descriptor,omitempty"`
									AcademicSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Academic_Suffix_Reference,omitempty"`
									HereditarySuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Hereditary_Suffix_Reference,omitempty"`
									HonorarySuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Honorary_Suffix_Reference,omitempty"`
									ProfessionalSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Professional_Suffix_Reference,omitempty"`
									ReligiousSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Religious_Suffix_Reference,omitempty"`
									RoyalSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Royal_Suffix_Reference,omitempty"`
								} `xml:"Suffix_Data,omitempty"`
								FullNameForSingaporeAndMalaysia string `xml:"Full_Name_for_Singapore_and_Malaysia,omitempty"`
							} `xml:"Name_Detail_Data,omitempty"`
						} `xml:"Preferred_Name_Data,omitempty"`
						AdditionalNameData struct {
							Text           string `xml:",chardata"`
							NameDetailData struct {
								Text             string `xml:",chardata"`
								FormattedName    string `xml:"Formatted_Name,attr,omitempty"`
								ReportingName    string `xml:"Reporting_Name,attr,omitempty"`
								CountryReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Country_Reference,omitempty"`
								PrefixData struct {
									Text           string `xml:",chardata"`
									TitleReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Title_Reference,omitempty"`
									TitleDescriptor     string `xml:"Title_Descriptor,omitempty"`
									SalutationReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Salutation_Reference,omitempty"`
								} `xml:"Prefix_Data,omitempty"`
								FirstName           string `xml:"First_Name,omitempty"`
								MiddleName          string `xml:"Middle_Name,omitempty"`
								LastName            string `xml:"Last_Name,omitempty"`
								SecondaryLastName   string `xml:"Secondary_Last_Name,omitempty"`
								TertiaryLastName    string `xml:"Tertiary_Last_Name,omitempty"`
								LocalNameDetailData struct {
									Text               string `xml:",chardata"`
									LocalName          string `xml:"Local_Name,attr,omitempty"`
									LocalScript        string `xml:"Local_Script,attr,omitempty"`
									FirstName          string `xml:"First_Name,omitempty"`
									MiddleName         string `xml:"Middle_Name,omitempty"`
									LastName           string `xml:"Last_Name,omitempty"`
									SecondaryLastName  string `xml:"Secondary_Last_Name,omitempty"`
									FirstName2         string `xml:"First_Name_2,omitempty"`
									MiddleName2        string `xml:"Middle_Name_2,omitempty"`
									LastName2          string `xml:"Last_Name_2,omitempty"`
									SecondaryLastName2 string `xml:"Secondary_Last_Name_2,omitempty"`
								} `xml:"Local_Name_Detail_Data,omitempty"`
								SuffixData struct {
									Text                  string `xml:",chardata"`
									SocialSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Social_Suffix_Reference,omitempty"`
									SocialSuffixDescriptor  string `xml:"Social_Suffix_Descriptor,omitempty"`
									AcademicSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Academic_Suffix_Reference,omitempty"`
									HereditarySuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Hereditary_Suffix_Reference,omitempty"`
									HonorarySuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Honorary_Suffix_Reference,omitempty"`
									ProfessionalSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Professional_Suffix_Reference,omitempty"`
									ReligiousSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Religious_Suffix_Reference,omitempty"`
									RoyalSuffixReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Royal_Suffix_Reference,omitempty"`
								} `xml:"Suffix_Data,omitempty"`
								FullNameForSingaporeAndMalaysia string `xml:"Full_Name_for_Singapore_and_Malaysia,omitempty"`
							} `xml:"Name_Detail_Data,omitempty"`
							NameTypeReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Name_Type_Reference,omitempty"`
						} `xml:"Additional_Name_Data,omitempty"`
					} `xml:"Name_Data,omitempty"`
					PersonalInformationData struct {
						Text                              string `xml:",chardata"`
						PersonalInformationForCountryData struct {
							Text             string `xml:",chardata"`
							CountryReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Country_Reference,omitempty"`
							CountryPersonalInformationData struct {
								Text                   string `xml:",chardata"`
								MaritalStatusReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text       string `xml:",chardata"`
										Type       string `xml:"type,attr,omitempty"`
										ParentID   string `xml:"parent_id,attr,omitempty"`
										ParentType string `xml:"parent_type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Marital_Status_Reference,omitempty"`
								MaritalStatusDate string `xml:"Marital_Status_Date,omitempty"`
								ReligionReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Religion_Reference,omitempty"`
								DisabilityStatusData struct {
									Text                string `xml:",chardata"`
									DisabilityReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Disability_Reference,omitempty"`
									DisabilityStatusDate     string `xml:"Disability_Status_Date,omitempty"`
									DisabilityDateKnown      string `xml:"Disability_Date_Known,omitempty"`
									DisabilityEndDate        string `xml:"Disability_End_Date,omitempty"`
									DisabilityGradeReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Disability_Grade_Reference,omitempty"`
									DisabilityDegree                          string `xml:"Disability_Degree,omitempty"`
									DisabilityRemainingCapacity               string `xml:"Disability_Remaining_Capacity,omitempty"`
									DisabilityCertificationAuthorityReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Disability_Certification_Authority_Reference,omitempty"`
									DisabilityCertificationAuthority      string `xml:"Disability_Certification_Authority,omitempty"`
									DisabilityCertifiedAt                 string `xml:"Disability_Certified_At,omitempty"`
									DisabilityCertificationID             string `xml:"Disability_Certification_ID,omitempty"`
									DisabilityCertificationBasisReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Disability_Certification_Basis_Reference,omitempty"`
									DisabilitySeverityRecognitionDate string `xml:"Disability_Severity_Recognition_Date,omitempty"`
									DisabilityFTETowardQuota          string `xml:"Disability_FTE_Toward_Quota,omitempty"`
									DisabilityWorkRestrictions        string `xml:"Disability_Work_Restrictions,omitempty"`
									DisabilityAccommodationsRequested string `xml:"Disability_Accommodations_Requested,omitempty"`
									DisabilityAccommodationsProvided  string `xml:"Disability_Accommodations_Provided,omitempty"`
									DisabilityRehabilitationRequested string `xml:"Disability_Rehabilitation_Requested,omitempty"`
									DisabilityRehabilitationProvided  string `xml:"Disability_Rehabilitation_Provided,omitempty"`
									Note                              string `xml:"Note,omitempty"`
									WorkerDocumentReference           struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Worker_Document_Reference,omitempty"`
									DisabilityStatusReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Disability_Status_Reference,omitempty"`
								} `xml:"Disability_Status_Data,omitempty"`
								EthnicityReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Ethnicity_Reference,omitempty"`
								RaceEthnicityDetailsReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Race_Ethnicity_Details_Reference,omitempty"`
								EthnicityVisualSurveyReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Ethnicity_Visual_Survey_Reference,omitempty"`
								HispanicOrLatino                            string `xml:"Hispanic_or_Latino,omitempty"`
								HispanicOrLatinoVisualSurvey                string `xml:"Hispanic_or_Latino_Visual_Survey,omitempty"`
								AboriginalIndigenousIdentificationReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Aboriginal_Indigenous_Identification_Reference,omitempty"`
								AboriginalIndigenousIdentificationDetailsReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Aboriginal_Indigenous_Identification_Details_Reference,omitempty"`
								HukouRegionReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Hukou_Region_Reference,omitempty"`
								HukouSubregionReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Hukou_Subregion_Reference,omitempty"`
								HukouLocality      string `xml:"Hukou_Locality,omitempty"`
								HukouPostalCode    string `xml:"Hukou_Postal_Code,omitempty"`
								HukouTypeReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Hukou_Type_Reference,omitempty"`
								LocalHukou            string `xml:"Local_Hukou,omitempty"`
								NativeRegionReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Native_Region_Reference,omitempty"`
								NativeRegionDescriptor        string `xml:"Native_Region_Descriptor,omitempty"`
								PersonnelFileAgencyForPerson  string `xml:"Personnel_File_Agency_for_Person,omitempty"`
								PoliticalAffiliationReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Political_Affiliation_Reference,omitempty"`
								SocialBenefitsLocalityReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Social_Benefits_Locality_Reference,omitempty"`
								RelativeNameInformationData struct {
									Text                           string `xml:",chardata"`
									Delete                         string `xml:"Delete,attr,omitempty"`
									RelativeNameReferenceReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Relative_Name_Reference_Reference,omitempty"`
									RelativeTypeReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Relative_Type_Reference,omitempty"`
									NameDetailData struct {
										Text             string `xml:",chardata"`
										FormattedName    string `xml:"Formatted_Name,attr,omitempty"`
										ReportingName    string `xml:"Reporting_Name,attr,omitempty"`
										CountryReference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Country_Reference,omitempty"`
										PrefixData struct {
											Text           string `xml:",chardata"`
											TitleReference struct {
												Text       string `xml:",chardata"`
												Descriptor string `xml:"Descriptor,attr,omitempty"`
												ID         struct {
													Text string `xml:",chardata"`
													Type string `xml:"type,attr,omitempty"`
												} `xml:"ID,omitempty"`
											} `xml:"Title_Reference,omitempty"`
											TitleDescriptor     string `xml:"Title_Descriptor,omitempty"`
											SalutationReference struct {
												Text       string `xml:",chardata"`
												Descriptor string `xml:"Descriptor,attr,omitempty"`
												ID         struct {
													Text string `xml:",chardata"`
													Type string `xml:"type,attr,omitempty"`
												} `xml:"ID,omitempty"`
											} `xml:"Salutation_Reference,omitempty"`
										} `xml:"Prefix_Data,omitempty"`
										FirstName           string `xml:"First_Name,omitempty"`
										MiddleName          string `xml:"Middle_Name,omitempty"`
										LastName            string `xml:"Last_Name,omitempty"`
										SecondaryLastName   string `xml:"Secondary_Last_Name,omitempty"`
										TertiaryLastName    string `xml:"Tertiary_Last_Name,omitempty"`
										LocalNameDetailData struct {
											Text               string `xml:",chardata"`
											LocalName          string `xml:"Local_Name,attr,omitempty"`
											LocalScript        string `xml:"Local_Script,attr,omitempty"`
											FirstName          string `xml:"First_Name,omitempty"`
											MiddleName         string `xml:"Middle_Name,omitempty"`
											LastName           string `xml:"Last_Name,omitempty"`
											SecondaryLastName  string `xml:"Secondary_Last_Name,omitempty"`
											FirstName2         string `xml:"First_Name_2,omitempty"`
											MiddleName2        string `xml:"Middle_Name_2,omitempty"`
											LastName2          string `xml:"Last_Name_2,omitempty"`
											SecondaryLastName2 string `xml:"Secondary_Last_Name_2,omitempty"`
										} `xml:"Local_Name_Detail_Data,omitempty"`
										SuffixData struct {
											Text                  string `xml:",chardata"`
											SocialSuffixReference struct {
												Text       string `xml:",chardata"`
												Descriptor string `xml:"Descriptor,attr,omitempty"`
												ID         struct {
													Text string `xml:",chardata"`
													Type string `xml:"type,attr,omitempty"`
												} `xml:"ID,omitempty"`
											} `xml:"Social_Suffix_Reference,omitempty"`
											SocialSuffixDescriptor  string `xml:"Social_Suffix_Descriptor,omitempty"`
											AcademicSuffixReference struct {
												Text       string `xml:",chardata"`
												Descriptor string `xml:"Descriptor,attr,omitempty"`
												ID         struct {
													Text string `xml:",chardata"`
													Type string `xml:"type,attr,omitempty"`
												} `xml:"ID,omitempty"`
											} `xml:"Academic_Suffix_Reference,omitempty"`
											HereditarySuffixReference struct {
												Text       string `xml:",chardata"`
												Descriptor string `xml:"Descriptor,attr,omitempty"`
												ID         struct {
													Text string `xml:",chardata"`
													Type string `xml:"type,attr,omitempty"`
												} `xml:"ID,omitempty"`
											} `xml:"Hereditary_Suffix_Reference,omitempty"`
											HonorarySuffixReference struct {
												Text       string `xml:",chardata"`
												Descriptor string `xml:"Descriptor,attr,omitempty"`
												ID         struct {
													Text string `xml:",chardata"`
													Type string `xml:"type,attr,omitempty"`
												} `xml:"ID,omitempty"`
											} `xml:"Honorary_Suffix_Reference,omitempty"`
											ProfessionalSuffixReference struct {
												Text       string `xml:",chardata"`
												Descriptor string `xml:"Descriptor,attr,omitempty"`
												ID         struct {
													Text string `xml:",chardata"`
													Type string `xml:"type,attr,omitempty"`
												} `xml:"ID,omitempty"`
											} `xml:"Professional_Suffix_Reference,omitempty"`
											ReligiousSuffixReference struct {
												Text       string `xml:",chardata"`
												Descriptor string `xml:"Descriptor,attr,omitempty"`
												ID         struct {
													Text string `xml:",chardata"`
													Type string `xml:"type,attr,omitempty"`
												} `xml:"ID,omitempty"`
											} `xml:"Religious_Suffix_Reference,omitempty"`
											RoyalSuffixReference struct {
												Text       string `xml:",chardata"`
												Descriptor string `xml:"Descriptor,attr,omitempty"`
												ID         struct {
													Text string `xml:",chardata"`
													Type string `xml:"type,attr,omitempty"`
												} `xml:"ID,omitempty"`
											} `xml:"Royal_Suffix_Reference,omitempty"`
										} `xml:"Suffix_Data,omitempty"`
										FullNameForSingaporeAndMalaysia string `xml:"Full_Name_for_Singapore_and_Malaysia,omitempty"`
									} `xml:"Name_Detail_Data,omitempty"`
								} `xml:"Relative_Name_Information_Data,omitempty"`
								SexualOrientationAndGenderIdentityReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Sexual_Orientation_and_Gender_Identity_Reference,omitempty"`
								GenderReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Gender_Reference,omitempty"`
								ReportingGenderReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Reporting_Gender_Reference,omitempty"`
								VeteransPreferenceReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Veterans_Preference_Reference,omitempty"`
								VeteransPreferenceForRIFReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Veterans_Preference_For_RIF_Reference,omitempty"`
								SelectiveServiceRegistrationReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Selective_Service_Registration_Reference,omitempty"`
								DisabledVeteranLeaveDateReference       string `xml:"Disabled_Veteran_Leave_Date_Reference,omitempty"`
								ActiveMilitaryUniformedServiceReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Active_Military_Uniformed_Service_Reference,omitempty"`
								UniformServiceReserveStatusReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Uniform_Service_Reserve_Status_Reference,omitempty"`
								CountrySpecificSectionData struct {
									Text                        string `xml:",chardata"`
									CountrySpecificSection1Data struct {
										Text            string `xml:",chardata"`
										Field1Reference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Field_1_Reference,omitempty"`
										Field2Reference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Field_2_Reference,omitempty"`
										Field3Reference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Field_3_Reference,omitempty"`
									} `xml:"Country_Specific_Section_1_Data,omitempty"`
									CountrySpecificSection2Data struct {
										Text            string `xml:",chardata"`
										Field1Reference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Field_1_Reference,omitempty"`
										Field2Reference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Field_2_Reference,omitempty"`
										Field3Reference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Field_3_Reference,omitempty"`
									} `xml:"Country_Specific_Section_2_Data,omitempty"`
									CountrySpecificSection3Data struct {
										Text            string `xml:",chardata"`
										Field1Reference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Field_1_Reference,omitempty"`
										Field2Reference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Field_2_Reference,omitempty"`
										Field3Reference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Field_3_Reference,omitempty"`
									} `xml:"Country_Specific_Section_3_Data,omitempty"`
								} `xml:"Country_Specific_Section_Data,omitempty"`
							} `xml:"Country_Personal_Information_Data,omitempty"`
						} `xml:"Personal_Information_For_Country_Data,omitempty"`
						BirthDate               string `xml:"Birth_Date,omitempty"`
						DateOfDeath             string `xml:"Date_of_Death,omitempty"`
						CountryOfBirthReference struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Country_of_Birth_Reference,omitempty"`
						CountryRegionOfBirthReference struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Country_Region_of_Birth_Reference,omitempty"`
						RegionOfBirthDescriptor string `xml:"Region_of_Birth_Descriptor,omitempty"`
						CityOfBirth             string `xml:"City_of_Birth,omitempty"`
						CityOfBirthReference    struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text       string `xml:",chardata"`
								Type       string `xml:"type,attr,omitempty"`
								ParentID   string `xml:"parent_id,attr,omitempty"`
								ParentType string `xml:"parent_type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"City_of_Birth_Reference,omitempty"`
						CitizenshipStatusReference struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Citizenship_Status_Reference,omitempty"`
						PrimaryNationalityReference struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Primary_Nationality_Reference,omitempty"`
						AdditionalNationalitiesReference struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Additional_Nationalities_Reference,omitempty"`
						LastMedicalExamValidTo string `xml:"Last_Medical_Exam_Valid_To,omitempty"`
						LastMedicalExamDate    string `xml:"Last_Medical_Exam_Date,omitempty"`
						MedicalExamNotes       string `xml:"Medical_Exam_Notes,omitempty"`
						BloodTypeReference     struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Blood_Type_Reference,omitempty"`
						MilitaryServiceData struct {
							Text            string `xml:",chardata"`
							StatusReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Status_Reference,omitempty"`
							DischargeDate                string `xml:"Discharge_Date,omitempty"`
							StatusBeginDate              string `xml:"Status_Begin_Date,omitempty"`
							MilitaryServiceTypeReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Military_Service_Type_Reference,omitempty"`
							MilitaryRankReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Military_Rank_Reference,omitempty"`
							Notes                    string `xml:"Notes,omitempty"`
							MilitaryServiceReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Military_Service_Reference,omitempty"`
							MilitaryDischargeTypeReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Military_Discharge_Type_Reference,omitempty"`
						} `xml:"Military_Service_Data,omitempty"`
						SexualOrientationReference struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Sexual_Orientation_Reference,omitempty"`
						GenderIdentityReference struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Gender_Identity_Reference,omitempty"`
						PronounReference struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Pronoun_Reference,omitempty"`
						NonCountrySpecificSectionData struct {
							Text                           string `xml:",chardata"`
							NonCountrySpecificSection1Data struct {
								Text            string `xml:",chardata"`
								Field1Reference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Field_1_Reference,omitempty"`
								Field2Reference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Field_2_Reference,omitempty"`
								Field3Reference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Field_3_Reference,omitempty"`
							} `xml:"Non_Country_Specific_Section_1_Data,omitempty"`
							NonCountrySpecificSection2Data struct {
								Text            string `xml:",chardata"`
								Field1Reference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Field_1_Reference,omitempty"`
								Field2Reference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Field_2_Reference,omitempty"`
								Field3Reference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Field_3_Reference,omitempty"`
							} `xml:"Non_Country_Specific_Section_2_Data,omitempty"`
							NonCountrySpecificSection3Data struct {
								Text            string `xml:",chardata"`
								Field1Reference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Field_1_Reference,omitempty"`
								Field2Reference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Field_2_Reference,omitempty"`
								Field3Reference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Field_3_Reference,omitempty"`
							} `xml:"Non_Country_Specific_Section_3_Data,omitempty"`
						} `xml:"Non_Country_Specific_Section_Data,omitempty"`
					} `xml:"Personal_Information_Data,omitempty"`
					IdentificationData struct {
						Text       string `xml:",chardata"`
						NationalID struct {
							Text                string `xml:",chardata"`
							Delete              string `xml:"Delete,attr,omitempty"`
							NationalIDReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"National_ID_Reference,omitempty"`
							NationalIDData struct {
								Text            string `xml:",chardata"`
								ID              string `xml:"ID,omitempty"`
								IDTypeReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"ID_Type_Reference,omitempty"`
								CountryReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Country_Reference,omitempty"`
								IssuedDate          string `xml:"Issued_Date,omitempty"`
								ExpirationDate      string `xml:"Expiration_Date,omitempty"`
								VerificationDate    string `xml:"Verification_Date,omitempty"`
								Series              string `xml:"Series,omitempty"`
								IssuingAgency       string `xml:"Issuing_Agency,omitempty"`
								VerifiedByReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Verified_By_Reference,omitempty"`
							} `xml:"National_ID_Data,omitempty"`
							NationalIDSharedReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"National_ID_Shared_Reference,omitempty"`
						} `xml:"National_ID,omitempty"`
						GovernmentID struct {
							Text                  string `xml:",chardata"`
							Delete                string `xml:"Delete,attr,omitempty"`
							GovernmentIDReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Government_ID_Reference,omitempty"`
							GovernmentIDData struct {
								Text            string `xml:",chardata"`
								ID              string `xml:"ID,omitempty"`
								IDTypeReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"ID_Type_Reference,omitempty"`
								CountryReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Country_Reference,omitempty"`
								IssuedDate          string `xml:"Issued_Date,omitempty"`
								ExpirationDate      string `xml:"Expiration_Date,omitempty"`
								VerificationDate    string `xml:"Verification_Date,omitempty"`
								VerifiedByReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Verified_by_Reference,omitempty"`
							} `xml:"Government_ID_Data,omitempty"`
							GovernmentIDSharedReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Government_ID_Shared_Reference,omitempty"`
						} `xml:"Government_ID,omitempty"`
						VisaID struct {
							Text            string `xml:",chardata"`
							Delete          string `xml:"Delete,attr,omitempty"`
							VisaIDReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Visa_ID_Reference,omitempty"`
							VisaIDData struct {
								Text            string `xml:",chardata"`
								ID              string `xml:"ID,omitempty"`
								IDTypeReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"ID_Type_Reference,omitempty"`
								CountryReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Country_Reference,omitempty"`
								IssuedDate          string `xml:"Issued_Date,omitempty"`
								ExpirationDate      string `xml:"Expiration_Date,omitempty"`
								VerificationDate    string `xml:"Verification_Date,omitempty"`
								VerifiedByReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Verified_By_Reference,omitempty"`
							} `xml:"Visa_ID_Data,omitempty"`
							VisaIDSharedReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Visa_ID_Shared_Reference,omitempty"`
						} `xml:"Visa_ID,omitempty"`
						PassportID struct {
							Text                string `xml:",chardata"`
							Delete              string `xml:"Delete,attr,omitempty"`
							PassportIDReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Passport_ID_Reference,omitempty"`
							PassportIDData struct {
								Text            string `xml:",chardata"`
								ID              string `xml:"ID,omitempty"`
								IDTypeReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"ID_Type_Reference,omitempty"`
								CountryReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Country_Reference,omitempty"`
								IssuedDate          string `xml:"Issued_Date,omitempty"`
								ExpirationDate      string `xml:"Expiration_Date,omitempty"`
								VerificationDate    string `xml:"Verification_Date,omitempty"`
								VerifiedByReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Verified_By_Reference,omitempty"`
							} `xml:"Passport_ID_Data,omitempty"`
							PassportIDSharedReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Passport_ID_Shared_Reference,omitempty"`
						} `xml:"Passport_ID,omitempty"`
						LicenseID struct {
							Text               string `xml:",chardata"`
							Delete             string `xml:"Delete,attr,omitempty"`
							LicenseIDReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"License_ID_Reference,omitempty"`
							LicenseIDData struct {
								Text            string `xml:",chardata"`
								ID              string `xml:"ID,omitempty"`
								IDTypeReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"ID_Type_Reference,omitempty"`
								CountryReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Country_Reference,omitempty"`
								CountryRegionReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Country_Region_Reference,omitempty"`
								CountryRegionDescriptor string `xml:"Country_Region_Descriptor,omitempty"`
								AuthorityReference      struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Authority_Reference,omitempty"`
								LicenseClass        string `xml:"License_Class,omitempty"`
								IssuedDate          string `xml:"Issued_Date,omitempty"`
								ExpirationDate      string `xml:"Expiration_Date,omitempty"`
								VerificationDate    string `xml:"Verification_Date,omitempty"`
								VerifiedByReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Verified_By_Reference,omitempty"`
							} `xml:"License_ID_Data,omitempty"`
							LicenseIDSharedReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"License_ID_Shared_Reference,omitempty"`
						} `xml:"License_ID,omitempty"`
						CustomID struct {
							Text              string `xml:",chardata"`
							Delete            string `xml:"Delete,attr,omitempty"`
							CustomIDReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Custom_ID_Reference,omitempty"`
							CustomIDData struct {
								Text            string `xml:",chardata"`
								ID              string `xml:"ID,omitempty"`
								IDTypeReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"ID_Type_Reference,omitempty"`
								IssuedDate              string `xml:"Issued_Date,omitempty"`
								ExpirationDate          string `xml:"Expiration_Date,omitempty"`
								OrganizationIDReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Organization_ID_Reference,omitempty"`
								CustomDescription string `xml:"Custom_Description,omitempty"`
							} `xml:"Custom_ID_Data,omitempty"`
							CustomIDSharedReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Custom_ID_Shared_Reference,omitempty"`
						} `xml:"Custom_ID,omitempty"`
					} `xml:"Identification_Data,omitempty"`
					ContactData struct {
						Text        string `xml:",chardata"`
						AddressData struct {
							Text                         string `xml:",chardata"`
							FormattedAddress             string `xml:"Formatted_Address,attr,omitempty"`
							AddressFormatType            string `xml:"Address_Format_Type,attr,omitempty"`
							DefaultedBusinessSiteAddress string `xml:"Defaulted_Business_Site_Address,attr,omitempty"`
							Delete                       string `xml:"Delete,attr,omitempty"`
							DoNotReplaceAll              string `xml:"Do_Not_Replace_All,attr,omitempty"`
							EffectiveDate                string `xml:"Effective_Date,attr,omitempty"`
							CountryReference             struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Country_Reference,omitempty"`
							LastModified    string `xml:"Last_Modified,omitempty"`
							AddressLineData struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								Type       string `xml:"Type,attr,omitempty"`
							} `xml:"Address_Line_Data,omitempty"`
							Municipality         string `xml:"Municipality,omitempty"`
							CountryCityReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text       string `xml:",chardata"`
									Type       string `xml:"type,attr,omitempty"`
									ParentID   string `xml:"parent_id,attr,omitempty"`
									ParentType string `xml:"parent_type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Country_City_Reference,omitempty"`
							SubmunicipalityData struct {
								Text                 string `xml:",chardata"`
								AddressComponentName string `xml:"Address_Component_Name,attr,omitempty"`
								Type                 string `xml:"Type,attr,omitempty"`
							} `xml:"Submunicipality_Data,omitempty"`
							CountryRegionReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Country_Region_Reference,omitempty"`
							CountryRegionDescriptor string `xml:"Country_Region_Descriptor,omitempty"`
							SubregionData           struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								Type       string `xml:"Type,attr,omitempty"`
							} `xml:"Subregion_Data,omitempty"`
							PostalCode string `xml:"Postal_Code,omitempty"`
							UsageData  struct {
								Text     string `xml:",chardata"`
								Public   string `xml:"Public,attr,omitempty"`
								TypeData struct {
									Text          string `xml:",chardata"`
									Primary       string `xml:"Primary,attr,omitempty"`
									TypeReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Type_Reference,omitempty"`
								} `xml:"Type_Data,omitempty"`
								UseForReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Use_For_Reference,omitempty"`
								UseForTenantedReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Use_For_Tenanted_Reference,omitempty"`
								Comments string `xml:"Comments,omitempty"`
							} `xml:"Usage_Data,omitempty"`
							NumberOfDays      string `xml:"Number_of_Days,omitempty"`
							MunicipalityLocal string `xml:"Municipality_Local,omitempty"`
							AddressReference  struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Address_Reference,omitempty"`
							AddressID string `xml:"Address_ID,omitempty"`
						} `xml:"Address_Data,omitempty"`
						PhoneData struct {
							Text                             string `xml:",chardata"`
							AreaCode                         string `xml:"Area_Code,attr,omitempty"`
							TenantFormattedPhone             string `xml:"Tenant_Formatted_Phone,attr,omitempty"`
							InternationalFormattedPhone      string `xml:"International_Formatted_Phone,attr,omitempty"`
							PhoneNumberWithoutAreaCode       string `xml:"Phone_Number_Without_Area_Code,attr,omitempty"`
							NationalFormattedPhone           string `xml:"National_Formatted_Phone,attr,omitempty"`
							E164FormattedPhone               string `xml:"E164_Formatted_Phone,attr,omitempty"`
							WorkdayTraditionalFormattedPhone string `xml:"Workday_Traditional_Formatted_Phone,attr,omitempty"`
							Delete                           string `xml:"Delete,attr,omitempty"`
							DoNotReplaceAll                  string `xml:"Do_Not_Replace_All,attr,omitempty"`
							CountryISOCode                   string `xml:"Country_ISO_Code,omitempty"`
							InternationalPhoneCode           string `xml:"International_Phone_Code,omitempty"`
							PhoneNumber                      string `xml:"Phone_Number,omitempty"`
							PhoneExtension                   string `xml:"Phone_Extension,omitempty"`
							PhoneDeviceTypeReference         struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Phone_Device_Type_Reference,omitempty"`
							UsageData struct {
								Text     string `xml:",chardata"`
								Public   string `xml:"Public,attr,omitempty"`
								TypeData struct {
									Text          string `xml:",chardata"`
									Primary       string `xml:"Primary,attr,omitempty"`
									TypeReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Type_Reference,omitempty"`
								} `xml:"Type_Data,omitempty"`
								UseForReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Use_For_Reference,omitempty"`
								UseForTenantedReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Use_For_Tenanted_Reference,omitempty"`
								Comments string `xml:"Comments,omitempty"`
							} `xml:"Usage_Data,omitempty"`
							PhoneReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Phone_Reference,omitempty"`
							ID string `xml:"ID,omitempty"`
						} `xml:"Phone_Data,omitempty"`
						EmailAddressData struct {
							Text            string `xml:",chardata"`
							Delete          string `xml:"Delete,attr,omitempty"`
							DoNotReplaceAll string `xml:"Do_Not_Replace_All,attr,omitempty"`
							EmailAddress    string `xml:"Email_Address,omitempty"`
							EmailComment    string `xml:"Email_Comment,omitempty"`
							UsageData       struct {
								Text     string `xml:",chardata"`
								Public   string `xml:"Public,attr,omitempty"`
								TypeData struct {
									Text          string `xml:",chardata"`
									Primary       string `xml:"Primary,attr,omitempty"`
									TypeReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Type_Reference,omitempty"`
								} `xml:"Type_Data,omitempty"`
								UseForReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Use_For_Reference,omitempty"`
								UseForTenantedReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Use_For_Tenanted_Reference,omitempty"`
								Comments string `xml:"Comments,omitempty"`
							} `xml:"Usage_Data,omitempty"`
							EmailReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Email_Reference,omitempty"`
							ID string `xml:"ID,omitempty"`
						} `xml:"Email_Address_Data,omitempty"`
						InstantMessengerData struct {
							Text                          string `xml:",chardata"`
							Delete                        string `xml:"Delete,attr,omitempty"`
							DoNotReplaceAll               string `xml:"Do_Not_Replace_All,attr,omitempty"`
							InstantMessengerAddress       string `xml:"Instant_Messenger_Address,omitempty"`
							InstantMessengerTypeReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Instant_Messenger_Type_Reference,omitempty"`
							InstantMessengerComment string `xml:"Instant_Messenger_Comment,omitempty"`
							UsageData               struct {
								Text     string `xml:",chardata"`
								Public   string `xml:"Public,attr,omitempty"`
								TypeData struct {
									Text          string `xml:",chardata"`
									Primary       string `xml:"Primary,attr,omitempty"`
									TypeReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Type_Reference,omitempty"`
								} `xml:"Type_Data,omitempty"`
								UseForReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Use_For_Reference,omitempty"`
								UseForTenantedReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Use_For_Tenanted_Reference,omitempty"`
								Comments string `xml:"Comments,omitempty"`
							} `xml:"Usage_Data,omitempty"`
							InstantMessengerReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Instant_Messenger_Reference,omitempty"`
							ID string `xml:"ID,omitempty"`
						} `xml:"Instant_Messenger_Data,omitempty"`
						WebAddressData struct {
							Text              string `xml:",chardata"`
							Delete            string `xml:"Delete,attr,omitempty"`
							DoNotReplaceAll   string `xml:"Do_Not_Replace_All,attr,omitempty"`
							WebAddress        string `xml:"Web_Address,omitempty"`
							WebAddressComment string `xml:"Web_Address_Comment,omitempty"`
							UsageData         struct {
								Text     string `xml:",chardata"`
								Public   string `xml:"Public,attr,omitempty"`
								TypeData struct {
									Text          string `xml:",chardata"`
									Primary       string `xml:"Primary,attr,omitempty"`
									TypeReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Type_Reference,omitempty"`
								} `xml:"Type_Data,omitempty"`
								UseForReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Use_For_Reference,omitempty"`
								UseForTenantedReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Use_For_Tenanted_Reference,omitempty"`
								Comments string `xml:"Comments,omitempty"`
							} `xml:"Usage_Data,omitempty"`
							WebAddressReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Web_Address_Reference,omitempty"`
							ID string `xml:"ID,omitempty"`
						} `xml:"Web_Address_Data,omitempty"`
					} `xml:"Contact_Data,omitempty"`
					TobaccoUse string `xml:"Tobacco_Use,omitempty"`
				} `xml:"Personal_Data,omitempty"`
				EmploymentData struct {
					Text          string `xml:",chardata"`
					WorkerJobData struct {
						Text         string `xml:",chardata"`
						PrimaryJob   string `xml:"Primary_Job,attr,omitempty"`
						PositionData struct {
							Text              string `xml:",chardata"`
							EffectiveDate     string `xml:"Effective_Date,attr,omitempty"`
							PositionReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Position_Reference,omitempty"`
							HeadcountReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Headcount_Reference,omitempty"`
							PositionID                   string `xml:"Position_ID,omitempty"`
							PositionTitle                string `xml:"Position_Title,omitempty"`
							BusinessTitle                string `xml:"Business_Title,omitempty"`
							StartDate                    string `xml:"Start_Date,omitempty"`
							EndEmploymentDate            string `xml:"End_Employment_Date,omitempty"`
							EndEmploymentReasonReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"End_Employment_Reason_Reference,omitempty"`
							WorkerTypeReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Worker_Type_Reference,omitempty"`
							PositionTimeTypeReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Position_Time_Type_Reference,omitempty"`
							JobExempt                     string `xml:"Job_Exempt,omitempty"`
							ScheduledWeeklyHours          string `xml:"Scheduled_Weekly_Hours,omitempty"`
							DefaultWeeklyHours            string `xml:"Default_Weekly_Hours,omitempty"`
							WorkingTimeFrequencyReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Working_Time_Frequency_Reference,omitempty"`
							WorkingTimeUnitReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Working_Time_Unit_Reference,omitempty"`
							WorkingTimeValue             string `xml:"Working_Time_Value,omitempty"`
							FullTimeEquivalentPercentage string `xml:"Full_Time_Equivalent_Percentage,omitempty"`
							SpecifyPaidFTE               string `xml:"Specify_Paid_FTE,omitempty"`
							PaidFTE                      string `xml:"Paid_FTE,omitempty"`
							SpecifyWorkingFTE            string `xml:"Specify_Working_FTE,omitempty"`
							WorkingFTE                   string `xml:"Working_FTE,omitempty"`
							ExcludeFromHeadcount         string `xml:"Exclude_from_Headcount,omitempty"`
							PayRateTypeReference         struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Pay_Rate_Type_Reference,omitempty"`
							JobClassificationSummaryData struct {
								Text                       string `xml:",chardata"`
								Additional                 string `xml:"Additional,attr,omitempty"`
								JobClassificationReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Job_Classification_Reference,omitempty"`
								JobGroupReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Job_Group_Reference,omitempty"`
							} `xml:"Job_Classification_Summary_Data,omitempty"`
							CompanyInsiderReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Company_Insider_Reference,omitempty"`
							WorkShiftReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Work_Shift_Reference,omitempty"`
							WorkHoursProfilesReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Work_Hours_Profiles_Reference,omitempty"`
							FederalWithholdingFEIN     string `xml:"Federal_Withholding_FEIN,omitempty"`
							WorkerCompensationCodeData struct {
								Text                             string `xml:",chardata"`
								WorkersCompensationCodeReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Workers_Compensation_Code_Reference,omitempty"`
								WorkersCompensationCode string `xml:"Workers_Compensation_Code,omitempty"`
							} `xml:"Worker_Compensation_Code_Data,omitempty"`
							PositionPayrollReportingCodeData struct {
								Text                          string `xml:",chardata"`
								PayrollReportingCodeReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text       string `xml:",chardata"`
										Type       string `xml:"type,attr,omitempty"`
										ParentID   string `xml:"parent_id,attr,omitempty"`
										ParentType string `xml:"parent_type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Payroll_Reporting_Code_Reference,omitempty"`
								Code                          string `xml:"Code,omitempty"`
								FormattedCode                 string `xml:"Formatted_Code,omitempty"`
								Name                          string `xml:"Name,omitempty"`
								PayrollReportingTypeReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Payroll_Reporting_Type_Reference,omitempty"`
							} `xml:"Position_Payroll_Reporting_Code_Data,omitempty"`
							JobProfileSummaryData struct {
								Text                string `xml:",chardata"`
								JobProfileReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Job_Profile_Reference,omitempty"`
								JobExempt                string `xml:"Job_Exempt,omitempty"`
								ManagementLevelReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Management_Level_Reference,omitempty"`
								JobCategoryReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Job_Category_Reference,omitempty"`
								JobFamilyReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Job_Family_Reference,omitempty"`
								JobProfileName            string `xml:"Job_Profile_Name,omitempty"`
								WorkShiftRequired         string `xml:"Work_Shift_Required,omitempty"`
								CriticalJob               string `xml:"Critical_Job,omitempty"`
								DifficultyToFillReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Difficulty_to_Fill_Reference,omitempty"`
							} `xml:"Job_Profile_Summary_Data,omitempty"`
							BusinessSiteSummaryData struct {
								Text              string `xml:",chardata"`
								LocationReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Location_Reference,omitempty"`
								Name                  string `xml:"Name,omitempty"`
								LocationTypeReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Location_Type_Reference,omitempty"`
								LocalReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Local_Reference,omitempty"`
								DisplayLanguageReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Display_Language_Reference,omitempty"`
								TimeProfileReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Time_Profile_Reference,omitempty"`
								ScheduledWeeklyHours string `xml:"Scheduled_Weekly_Hours,omitempty"`
								AddressData          struct {
									Text                         string `xml:",chardata"`
									FormattedAddress             string `xml:"Formatted_Address,attr,omitempty"`
									AddressFormatType            string `xml:"Address_Format_Type,attr,omitempty"`
									DefaultedBusinessSiteAddress string `xml:"Defaulted_Business_Site_Address,attr,omitempty"`
									Delete                       string `xml:"Delete,attr,omitempty"`
									DoNotReplaceAll              string `xml:"Do_Not_Replace_All,attr,omitempty"`
									EffectiveDate                string `xml:"Effective_Date,attr,omitempty"`
									CountryReference             struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Country_Reference,omitempty"`
									LastModified    string `xml:"Last_Modified,omitempty"`
									AddressLineData struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										Type       string `xml:"Type,attr,omitempty"`
									} `xml:"Address_Line_Data,omitempty"`
									Municipality         string `xml:"Municipality,omitempty"`
									CountryCityReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text       string `xml:",chardata"`
											Type       string `xml:"type,attr,omitempty"`
											ParentID   string `xml:"parent_id,attr,omitempty"`
											ParentType string `xml:"parent_type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Country_City_Reference,omitempty"`
									SubmunicipalityData struct {
										Text                 string `xml:",chardata"`
										AddressComponentName string `xml:"Address_Component_Name,attr,omitempty"`
										Type                 string `xml:"Type,attr,omitempty"`
									} `xml:"Submunicipality_Data,omitempty"`
									CountryRegionReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Country_Region_Reference,omitempty"`
									CountryRegionDescriptor string `xml:"Country_Region_Descriptor,omitempty"`
									SubregionData           struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										Type       string `xml:"Type,attr,omitempty"`
									} `xml:"Subregion_Data,omitempty"`
									PostalCode string `xml:"Postal_Code,omitempty"`
									UsageData  struct {
										Text     string `xml:",chardata"`
										Public   string `xml:"Public,attr,omitempty"`
										TypeData struct {
											Text          string `xml:",chardata"`
											Primary       string `xml:"Primary,attr,omitempty"`
											TypeReference struct {
												Text       string `xml:",chardata"`
												Descriptor string `xml:"Descriptor,attr,omitempty"`
												ID         struct {
													Text string `xml:",chardata"`
													Type string `xml:"type,attr,omitempty"`
												} `xml:"ID,omitempty"`
											} `xml:"Type_Reference,omitempty"`
										} `xml:"Type_Data,omitempty"`
										UseForReference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Use_For_Reference,omitempty"`
										UseForTenantedReference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Use_For_Tenanted_Reference,omitempty"`
										Comments string `xml:"Comments,omitempty"`
									} `xml:"Usage_Data,omitempty"`
									NumberOfDays      string `xml:"Number_of_Days,omitempty"`
									MunicipalityLocal string `xml:"Municipality_Local,omitempty"`
									AddressReference  struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Address_Reference,omitempty"`
									AddressID string `xml:"Address_ID,omitempty"`
								} `xml:"Address_Data,omitempty"`
							} `xml:"Business_Site_Summary_Data,omitempty"`
							PayrollInterfaceProcessingData struct {
								Text                 string `xml:",chardata"`
								EffectiveDate        string `xml:"Effective_Date,omitempty"`
								PayRateTypeReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Pay_Rate_Type_Reference,omitempty"`
								FrequencyReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Frequency_Reference,omitempty"`
								PayGroupReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text       string `xml:",chardata"`
										Type       string `xml:"type,attr,omitempty"`
										ParentID   string `xml:"parent_id,attr,omitempty"`
										ParentType string `xml:"parent_type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Pay_Group_Reference,omitempty"`
								PayrollEntityReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Payroll_Entity_Reference,omitempty"`
								ExternalEmployeeTypeReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"External_Employee_Type_Reference,omitempty"`
								PayrollFileNumber string `xml:"Payroll_File_Number,omitempty"`
							} `xml:"Payroll_Interface_Processing_Data,omitempty"`
							RegularPaidEquivalentHours       string `xml:"Regular_Paid_Equivalent_Hours,omitempty"`
							WorkerHoursProfileClassification string `xml:"Worker_Hours_Profile_Classification,omitempty"`
							InternationalAssignmentData      struct {
								Text                                 string `xml:",chardata"`
								InternationalAssignmentTypeReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"International_Assignment_Type_Reference,omitempty"`
								StartInternationalAssignmentReasonReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Start_International_Assignment_Reason_Reference,omitempty"`
								ExpectedAssignmentEndDate string `xml:"Expected_Assignment_End_Date,omitempty"`
								HostCountryReference      struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Host_Country_Reference,omitempty"`
								HomeCountryReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Home_Country_Reference,omitempty"`
							} `xml:"International_Assignment_Data,omitempty"`
							WorkSpaceReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Work_Space__Reference,omitempty"`
							AcademicPaySetupData struct {
								Text                              string `xml:",chardata"`
								AnnualWorkPeriodWorkPercentOfYear string `xml:"Annual_Work_Period_Work_Percent_of_Year,omitempty"`
								AnnualWorkPeriodStartDate         string `xml:"Annual_Work_Period_Start_Date,omitempty"`
								AnnualWorkPeriodEndDate           string `xml:"Annual_Work_Period_End_Date,omitempty"`
								DisbursementPlanPeriodStartDate   string `xml:"Disbursement_Plan_Period_Start_Date,omitempty"`
								DisbursementPlanPeriodEndDate     string `xml:"Disbursement_Plan_Period_End_Date,omitempty"`
							} `xml:"Academic_Pay_Setup_Data,omitempty"`
							EndDate                        string `xml:"End_Date,omitempty"`
							PayThroughDate                 string `xml:"Pay_Through_Date,omitempty"`
							CollectiveAgreementSummaryData struct {
								Text                    string `xml:",chardata"`
								CollectiveAgreementData struct {
									Text                                            string `xml:",chardata"`
									AssignEmployeeCollectiveAgreementEventReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Assign_Employee_Collective_Agreement_Event_Reference,omitempty"`
									EffectiveDate           string `xml:"Effective_Date,omitempty"`
									EndAssignmentDate       string `xml:"End_Assignment_Date,omitempty"`
									CollectiveAgreementData struct {
										Text                         string `xml:",chardata"`
										CollectiveAgreementReference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Collective_Agreement_Reference,omitempty"`
										CollectiveAgreementReviewDate string `xml:"Collective_Agreement_Review_Date,omitempty"`
										CollectiveAgreementFactor     struct {
											Text                               string `xml:",chardata"`
											Order                              string `xml:"Order,omitempty"`
											CollectiveAgreementFactorReference struct {
												Text       string `xml:",chardata"`
												Descriptor string `xml:"Descriptor,attr,omitempty"`
												ID         struct {
													Text string `xml:",chardata"`
													Type string `xml:"type,attr,omitempty"`
												} `xml:"ID,omitempty"`
											} `xml:"Collective_Agreement_Factor_Reference,omitempty"`
											CollectiveAgreementFactorOptionReference struct {
												Text       string `xml:",chardata"`
												Descriptor string `xml:"Descriptor,attr,omitempty"`
												ID         struct {
													Text string `xml:",chardata"`
													Type string `xml:"type,attr,omitempty"`
												} `xml:"ID,omitempty"`
											} `xml:"Collective_Agreement_Factor_Option_Reference,omitempty"`
										} `xml:"Collective_Agreement_Factor,omitempty"`
									} `xml:"Collective_Agreement_Data,omitempty"`
								} `xml:"Collective_Agreement_Data,omitempty"`
							} `xml:"Collective_Agreement_Summary_Data,omitempty"`
							EmployeeProbationPeriodSummaryData struct {
								Text                              string `xml:",chardata"`
								EmployeeProbationPeriodDetailData struct {
									Text                                  string `xml:",chardata"`
									EmployeeProbationPeriodEventReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Employee_Probation_Period_Event_Reference,omitempty"`
									EmployeeProbationPeriodReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Employee_Probation_Period_Reference,omitempty"`
									EffectiveDate          string `xml:"Effective_Date,omitempty"`
									StartDate              string `xml:"Start_Date,omitempty"`
									EndDate                string `xml:"End_Date,omitempty"`
									ProbationTypeReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Probation_Type_Reference,omitempty"`
									ProbationReasonReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Probation_Reason_Reference,omitempty"`
									ExtendedEndDate string `xml:"Extended_End_Date,omitempty"`
									Note            string `xml:"Note,omitempty"`
								} `xml:"Employee_Probation_Period_Detail_Data,omitempty"`
							} `xml:"Employee_Probation_Period_Summary_Data,omitempty"`
							ManagerAsOfLastDetectedManagerChangeReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Manager_as_of_last_detected_manager_change_Reference,omitempty"`
						} `xml:"Position_Data,omitempty"`
						PositionOrganizationsData struct {
							Text                     string `xml:",chardata"`
							PositionOrganizationData struct {
								Text                  string `xml:",chardata"`
								OrganizationReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Organization_Reference,omitempty"`
								OrganizationData struct {
									Text                    string `xml:",chardata"`
									OrganizationReferenceID string `xml:"Organization_Reference_ID,omitempty"`
									OrganizationCode        string `xml:"Organization_Code,omitempty"`
									IntegrationIDData       struct {
										Text string `xml:",chardata"`
										ID   struct {
											Text     string `xml:",chardata"`
											SystemID string `xml:"System_ID,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Integration_ID_Data,omitempty"`
									OrganizationName          string `xml:"Organization_Name,omitempty"`
									OrganizationTypeReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Organization_Type_Reference,omitempty"`
									OrganizationSubtypeReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Organization_Subtype_Reference,omitempty"`
									PrimaryBusinessSiteReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Primary_Business_Site_Reference,omitempty"`
									OrganizationSupportRoleData struct {
										Text                    string `xml:",chardata"`
										OrganizationSupportRole struct {
											Text                      string `xml:",chardata"`
											OrganizationRoleReference struct {
												Text       string `xml:",chardata"`
												Descriptor string `xml:"Descriptor,attr,omitempty"`
												ID         struct {
													Text string `xml:",chardata"`
													Type string `xml:"type,attr,omitempty"`
												} `xml:"ID,omitempty"`
											} `xml:"Organization_Role_Reference,omitempty"`
											OrganizationRoleData struct {
												Text            string `xml:",chardata"`
												WorkerReference struct {
													Text       string `xml:",chardata"`
													Descriptor string `xml:"Descriptor,attr,omitempty"`
													ID         struct {
														Text string `xml:",chardata"`
														Type string `xml:"type,attr,omitempty"`
													} `xml:"ID,omitempty"`
												} `xml:"Worker_Reference,omitempty"`
												AssignmentFrom string `xml:"Assignment_From,omitempty"`
											} `xml:"Organization_Role_Data,omitempty"`
										} `xml:"Organization_Support_Role,omitempty"`
									} `xml:"Organization_Support_Role_Data,omitempty"`
									DateOfPayGroupAssignment            string `xml:"Date_of_Pay_Group_Assignment,omitempty"`
									UsedInChangeOrganizationAssignments string `xml:"Used_in_Change_Organization_Assignments,omitempty"`
								} `xml:"Organization_Data,omitempty"`
							} `xml:"Position_Organization_Data,omitempty"`
							PayGroupAssignmentCorrectOrRescindData struct {
								Text                      string `xml:",chardata"`
								EffectiveDate             string `xml:"Effective_Date,omitempty"`
								CompletedOn               string `xml:"Completed_On,omitempty"`
								EventCorrected            string `xml:"Event_Corrected,omitempty"`
								EventRescinded            string `xml:"Event_Rescinded,omitempty"`
								OriginalPayGroupReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Original_Pay_Group_Reference,omitempty"`
								ProposedPayGroupReference struct {
									Text       string `xml:",chardata"`
									Descriptor string `xml:"Descriptor,attr,omitempty"`
									ID         struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr,omitempty"`
									} `xml:"ID,omitempty"`
								} `xml:"Proposed_Pay_Group_Reference,omitempty"`
							} `xml:"Pay_Group_Assignment_Correct_or_Rescind_Data,omitempty"`
						} `xml:"Position_Organizations_Data,omitempty"`
						PositionManagementChainsData struct {
							Text                                   string `xml:",chardata"`
							PositionSupervisoryManagementChainData struct {
								Text                string `xml:",chardata"`
								ManagementChainData struct {
									Text                  string `xml:",chardata"`
									OrganizationReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Organization_Reference,omitempty"`
									ManagerReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Manager_Reference,omitempty"`
									Manager struct {
										Text            string `xml:",chardata"`
										WorkerReference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Worker_Reference,omitempty"`
										WorkerDescriptor string `xml:"Worker_Descriptor,omitempty"`
									} `xml:"Manager,omitempty"`
								} `xml:"Management_Chain_Data,omitempty"`
							} `xml:"Position_Supervisory_Management_Chain_Data,omitempty"`
							PositionMatrixManagementChainData struct {
								Text                string `xml:",chardata"`
								ManagementChainData struct {
									Text                  string `xml:",chardata"`
									OrganizationReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Organization_Reference,omitempty"`
									ManagerReference struct {
										Text       string `xml:",chardata"`
										Descriptor string `xml:"Descriptor,attr,omitempty"`
										ID         struct {
											Text string `xml:",chardata"`
											Type string `xml:"type,attr,omitempty"`
										} `xml:"ID,omitempty"`
									} `xml:"Manager_Reference,omitempty"`
									Manager struct {
										Text            string `xml:",chardata"`
										WorkerReference struct {
											Text       string `xml:",chardata"`
											Descriptor string `xml:"Descriptor,attr,omitempty"`
											ID         struct {
												Text string `xml:",chardata"`
												Type string `xml:"type,attr,omitempty"`
											} `xml:"ID,omitempty"`
										} `xml:"Worker_Reference,omitempty"`
										WorkerDescriptor string `xml:"Worker_Descriptor,omitempty"`
									} `xml:"Manager,omitempty"`
								} `xml:"Management_Chain_Data,omitempty"`
							} `xml:"Position_Matrix_Management_Chain_Data,omitempty"`
						} `xml:"Position_Management_Chains_Data,omitempty"`
						PositionContractDetailData struct {
							Text              string `xml:",chardata"`
							ContractPayRate   string `xml:"Contract_Pay_Rate,omitempty"`
							CurrencyReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Currency_Reference,omitempty"`
							FrequencyReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Frequency_Reference,omitempty"`
							ContractAssignmentDetails string `xml:"Contract_Assignment_Details,omitempty"`
							ContractEndDate           string `xml:"Contract_End_Date,omitempty"`
							SupplierReference         struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Supplier_Reference,omitempty"`
						} `xml:"Position_Contract_Detail_Data,omitempty"`
					} `xml:"Worker_Job_Data,omitempty"`
					WorkerStatusData struct {
						Text                string `xml:",chardata"`
						Active              string `xml:"Active,omitempty"`
						ActiveStatusDate    string `xml:"Active_Status_Date,omitempty"`
						HireDate            string `xml:"Hire_Date,omitempty"`
						OriginalHireDate    string `xml:"Original_Hire_Date,omitempty"`
						HireReasonReference struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Hire_Reason_Reference,omitempty"`
						EndEmploymentDate                 string `xml:"End_Employment_Date,omitempty"`
						ContinuousServiceDate             string `xml:"Continuous_Service_Date,omitempty"`
						FirstDayOfWork                    string `xml:"First_Day_of_Work,omitempty"`
						ExpectedRetirementDate            string `xml:"Expected_Retirement_Date,omitempty"`
						RetirementEligibilityDate         string `xml:"Retirement_Eligibility_Date,omitempty"`
						Retired                           string `xml:"Retired,omitempty"`
						RetirementDate                    string `xml:"Retirement_Date,omitempty"`
						SeniorityDate                     string `xml:"Seniority_Date,omitempty"`
						SeveranceDate                     string `xml:"Severance_Date,omitempty"`
						BenefitsServiceDate               string `xml:"Benefits_Service_Date,omitempty"`
						CompanyServiceDate                string `xml:"Company_Service_Date,omitempty"`
						TimeOffServiceDate                string `xml:"Time_Off_Service_Date,omitempty"`
						VestingDate                       string `xml:"Vesting_Date,omitempty"`
						DateEnteredWorkforce              string `xml:"Date_Entered_Workforce,omitempty"`
						DaysUnemployed                    string `xml:"Days_Unemployed,omitempty"`
						MonthsContinuousPriorEmployment   string `xml:"Months_Continuous_Prior_Employment,omitempty"`
						Terminated                        string `xml:"Terminated,omitempty"`
						TerminationDate                   string `xml:"Termination_Date,omitempty"`
						PayThroughDate                    string `xml:"Pay_Through_Date,omitempty"`
						AgreementSignatureDate            string `xml:"Agreement_Signature_Date,omitempty"`
						DismissalProcessStartDate         string `xml:"Dismissal_Process_Start_Date,omitempty"`
						NoticePeriodStartDate             string `xml:"Notice_Period_Start_Date,omitempty"`
						PrimaryTerminationReasonReference struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Primary_Termination_Reason_Reference,omitempty"`
						PrimaryTerminationCategoryReference struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Primary_Termination_Category_Reference,omitempty"`
						TerminationInvoluntary          string `xml:"Termination_Involuntary,omitempty"`
						SecondaryTerminationReasonsData struct {
							Text                                string `xml:",chardata"`
							SecondaryTerminationReasonReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Secondary_Termination_Reason_Reference,omitempty"`
							SecondaryTerminationReasonCategoryReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Secondary_Termination_Reason_Category_Reference,omitempty"`
						} `xml:"Secondary_Termination_Reasons_Data,omitempty"`
						LocalTerminationReasonReference struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Local_Termination_Reason_Reference,omitempty"`
						EligibleForHireReference struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Eligible_for_Hire_Reference,omitempty"`
						RegrettableTermination                        string `xml:"Regrettable_Termination,omitempty"`
						EligibleForRehireOnLatestTerminationReference struct {
							Text       string `xml:",chardata"`
							Descriptor string `xml:"Descriptor,attr,omitempty"`
							ID         struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr,omitempty"`
							} `xml:"ID,omitempty"`
						} `xml:"Eligible_for_Rehire_on_Latest_Termination_Reference,omitempty"`
						HireRescinded            string `xml:"Hire_Rescinded,omitempty"`
						TerminationLastDayOfWork string `xml:"Termination_Last_Day_of_Work,omitempty"`
						ResignationDate          string `xml:"Resignation_Date,omitempty"`
						LastDateForWhichPaid     string `xml:"Last_Date_for_Which_Paid,omitempty"`
						ExpectedDateOfReturn     string `xml:"Expected_Date_of_Return,omitempty"`
						NotReturning             string `xml:"Not_Returning,omitempty"`
						ReturnUnknown            string `xml:"Return_Unknown,omitempty"`
						ProbationStartDate       string `xml:"Probation_Start_Date,omitempty"`
						ProbationEndDate         string `xml:"Probation_End_Date,omitempty"`
						LeaveStatusData          struct {
							Text                       string `xml:",chardata"`
							LeaveRequestEventReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Leave_Request_Event_Reference,omitempty"`
							LeaveRequestDescription   string `xml:"Leave_Request_Description,omitempty"`
							LeaveReturnEventReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Leave_Return_Event_Reference,omitempty"`
							OnLeave                     string `xml:"On_Leave,omitempty"`
							LeaveStartDate              string `xml:"Leave_Start_Date,omitempty"`
							EstimatedLeaveEndDate       string `xml:"Estimated_Leave_End_Date,omitempty"`
							LeaveEndDate                string `xml:"Leave_End_Date,omitempty"`
							FirstDayOfWork              string `xml:"First_Day_Of_Work,omitempty"`
							LeaveLastDayOfWork          string `xml:"Leave_Last_Day_of_Work,omitempty"`
							LeaveOfAbsenceTypeReference struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Leave_of_Absence_Type_Reference,omitempty"`
							BenefitsEffect                 string `xml:"Benefits_Effect,omitempty"`
							PayrollEffect                  string `xml:"Payroll_Effect,omitempty"`
							PaidTimeOffAccrualEffect       string `xml:"Paid_Time_Off_Accrual_Effect,omitempty"`
							ContinuousServiceAccrualEffect string `xml:"Continuous_Service_Accrual_Effect,omitempty"`
							StockVestingEffect             string `xml:"Stock_Vesting_Effect,omitempty"`
							LeaveTypeReasonReference       struct {
								Text       string `xml:",chardata"`
								Descriptor string `xml:"Descriptor,attr,omitempty"`
								ID         struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr,omitempty"`
								} `xml:"ID,omitempty"`
							} `xml:"Leave_Type_Reason_Reference,omitempty"`
							LeaveRequestAdditionalFields struct {
								Text                            string `xml:",chardata"`
								LastDateForWhichPaid            string `xml:"Last_Date_for_Which_Paid,omitempty"`
								ExpectedDueDate                 string `xml:"Expected_Due_Date,omitempty"`
								ChildSBirthDate                 string `xml:"Child_s_Birth_Date,omitempty"`
								StillbirthBabyDeceased          string `xml:"Stillbirth_Baby_Deceased,omitempty"`
								DateBabyArrivedHomeFromHospital string `xml:"Date_Baby_Arrived_Home_From_Hospital,omitempty"`
							} `xml:"Leave_Request_Additional_Fields,omitempty"`
						} `xml:"Leave_Status_Data,omitempty"`
					} `xml:"Worker_Status_Data,omitempty"`
				} `xml:"Employment_Data,omitempty"`
			} `xml:"Worker_Data,omitempty"`
		} `xml:"Worker,omitempty"`
	} `xml:"Response_Data,omitempty"`
} 