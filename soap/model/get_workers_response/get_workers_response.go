// Generated from ../model/get_workers_response/get_workers_response.go
package get_workers_response

import "encoding/xml"

type GetWorkersResponse struct {
    XMLName xml.Name `xml:"Get_Workers_Response"`
    Version *string `xml:"version,attr,omitempty"`
    RequestReferences *RequestReferences `xml:"Request_References,omitempty"`
    RequestCriteria *RequestCriteria `xml:"Request_Criteria,omitempty"`
    ResponseFilter *ResponseFilter `xml:"Response_Filter,omitempty"`
    ResponseGroup *ResponseGroup `xml:"Response_Group,omitempty"`
    ResponseResults *ResponseResults `xml:"Response_Results,omitempty"`
    ResponseData *ResponseData `xml:"Response_Data,omitempty"`
}

type RequestReferences struct {
    SkipNonExistingInstances *string `xml:"Skip_Non_Existing_Instances,attr,omitempty"`
    IgnoreInvalidReferences *string `xml:"Ignore_Invalid_References,attr,omitempty"`
    WorkerReference *WorkerReference `xml:"Worker_Reference,omitempty"`
}

type WorkerReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Id struct {
    Value *string `xml:",chardata"`
    Type *string `xml:"type,attr,omitempty"`
    ParentId *string `xml:"parent_id,attr,omitempty"`
    ParentType *string `xml:"parent_type,attr,omitempty"`
    SystemId *string `xml:"System_ID,attr,omitempty"`
}

type RequestCriteria struct {
    TransactionLogCriteriaData *TransactionLogCriteriaData `xml:"Transaction_Log_Criteria_Data,omitempty"`
    OrganizationReference *OrganizationReference `xml:"Organization_Reference,omitempty"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"`
    IncludeSubordinateOrganizations *IncludeSubordinateOrganizations `xml:"Include_Subordinate_Organizations,omitempty"`
    PositionReference *PositionReference `xml:"Position_Reference,omitempty"`
    EventReference *EventReference `xml:"Event_Reference,omitempty"`
    BenefitPlanReference *BenefitPlanReference `xml:"Benefit_Plan_Reference,omitempty"`
    FieldAndParameterCriteriaData *FieldAndParameterCriteriaData `xml:"Field_And_Parameter_Criteria_Data,omitempty"`
    NationalIdCriteriaData *NationalIdCriteriaData `xml:"National_ID_Criteria_Data,omitempty"`
    ExcludeInactiveWorkers *ExcludeInactiveWorkers `xml:"Exclude_Inactive_Workers,omitempty"`
    ExcludeEmployees *ExcludeEmployees `xml:"Exclude_Employees,omitempty"`
    ExcludeContingentWorkers *ExcludeContingentWorkers `xml:"Exclude_Contingent_Workers,omitempty"`
    EligibilityCriteriaData *EligibilityCriteriaData `xml:"Eligibility_Criteria_Data,omitempty"`
    UniversalIdReference *UniversalIdReference `xml:"Universal_ID_Reference,omitempty"`
}

type TransactionLogCriteriaData struct {
    TransactionDateRangeData *TransactionDateRangeData `xml:"Transaction_Date_Range_Data,omitempty"`
    TransactionTypeReferences *TransactionTypeReferences `xml:"Transaction_Type_References,omitempty"`
    SubscriberReference *SubscriberReference `xml:"Subscriber_Reference,omitempty"`
}

type TransactionDateRangeData struct {
    UpdatedFrom *UpdatedFrom `xml:"Updated_From,omitempty"`
    UpdatedThrough *UpdatedThrough `xml:"Updated_Through,omitempty"`
    EffectiveFrom *EffectiveFrom `xml:"Effective_From,omitempty"`
    EffectiveThrough *EffectiveThrough `xml:"Effective_Through,omitempty"`
}

type UpdatedFrom struct {
    Value *string `xml:",chardata"`
}

type UpdatedThrough struct {
    Value *string `xml:",chardata"`
}

type EffectiveFrom struct {
    Value *string `xml:",chardata"`
}

type EffectiveThrough struct {
    Value *string `xml:",chardata"`
}

type TransactionTypeReferences struct {
    TransactionTypeReference *TransactionTypeReference `xml:"Transaction_Type_Reference,omitempty"`
}

type TransactionTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SubscriberReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type OrganizationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CountryReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type IncludeSubordinateOrganizations struct {
    Value *string `xml:",chardata"`
}

type PositionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type EventReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type BenefitPlanReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type FieldAndParameterCriteriaData struct {
    ProviderReference *ProviderReference `xml:"Provider_Reference,omitempty"`
}

type ProviderReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type NationalIdCriteriaData struct {
    IdentifierId *IdentifierId `xml:"Identifier_ID,omitempty"`
    NationalIdTypeReference *NationalIdTypeReference `xml:"National_ID_Type_Reference,omitempty"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"`
}

type IdentifierId struct {
    Value *string `xml:",chardata"`
}

type NationalIdTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ExcludeInactiveWorkers struct {
    Value *string `xml:",chardata"`
}

type ExcludeEmployees struct {
    Value *string `xml:",chardata"`
}

type ExcludeContingentWorkers struct {
    Value *string `xml:",chardata"`
}

type EligibilityCriteriaData struct {
    FieldReference *FieldReference `xml:"Field_Reference,omitempty"`
    AsOfEffectiveDate *AsOfEffectiveDate `xml:"As_Of_Effective_Date,omitempty"`
    AsOfEntryDatetime *AsOfEntryDatetime `xml:"As_Of_Entry_DateTime,omitempty"`
}

type FieldReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type AsOfEffectiveDate struct {
    Value *string `xml:",chardata"`
}

type AsOfEntryDatetime struct {
    Value *string `xml:",chardata"`
}

type UniversalIdReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ResponseFilter struct {
    AsOfEffectiveDate *AsOfEffectiveDate `xml:"As_Of_Effective_Date,omitempty"`
    AsOfEntryDatetime *AsOfEntryDatetime `xml:"As_Of_Entry_DateTime,omitempty"`
    Page *Page `xml:"Page,omitempty"`
    Count *Count `xml:"Count,omitempty"`
}

type Page struct {
    Value *string `xml:",chardata"`
}

type Count struct {
    Value *string `xml:",chardata"`
}

type ResponseGroup struct {
    IncludeReference *IncludeReference `xml:"Include_Reference,omitempty"`
    IncludePersonalInformation *IncludePersonalInformation `xml:"Include_Personal_Information,omitempty"`
    ShowAllPersonalInformation *ShowAllPersonalInformation `xml:"Show_All_Personal_Information,omitempty"`
    IncludeAdditionalJobs *IncludeAdditionalJobs `xml:"Include_Additional_Jobs,omitempty"`
    IncludeEmploymentInformation *IncludeEmploymentInformation `xml:"Include_Employment_Information,omitempty"`
    IncludeCompensation *IncludeCompensation `xml:"Include_Compensation,omitempty"`
    IncludeOrganizations *IncludeOrganizations `xml:"Include_Organizations,omitempty"`
    ExcludeOrganizationSupportRoleData *ExcludeOrganizationSupportRoleData `xml:"Exclude_Organization_Support_Role_Data,omitempty"`
    ExcludeLocationHierarchies *ExcludeLocationHierarchies `xml:"Exclude_Location_Hierarchies,omitempty"`
    ExcludeCostCenters *ExcludeCostCenters `xml:"Exclude_Cost_Centers,omitempty"`
    ExcludeCostCenterHierarchies *ExcludeCostCenterHierarchies `xml:"Exclude_Cost_Center_Hierarchies,omitempty"`
    ExcludeCompanies *ExcludeCompanies `xml:"Exclude_Companies,omitempty"`
    ExcludeCompanyHierarchies *ExcludeCompanyHierarchies `xml:"Exclude_Company_Hierarchies,omitempty"`
    ExcludeMatrixOrganizations *ExcludeMatrixOrganizations `xml:"Exclude_Matrix_Organizations,omitempty"`
    ExcludePayGroups *ExcludePayGroups `xml:"Exclude_Pay_Groups,omitempty"`
    ExcludeRegions *ExcludeRegions `xml:"Exclude_Regions,omitempty"`
    ExcludeRegionHierarchies *ExcludeRegionHierarchies `xml:"Exclude_Region_Hierarchies,omitempty"`
    ExcludeSupervisoryOrganizations *ExcludeSupervisoryOrganizations `xml:"Exclude_Supervisory_Organizations,omitempty"`
    ExcludeTeams *ExcludeTeams `xml:"Exclude_Teams,omitempty"`
    ExcludeCustomOrganizations *ExcludeCustomOrganizations `xml:"Exclude_Custom_Organizations,omitempty"`
    IncludeRoles *IncludeRoles `xml:"Include_Roles,omitempty"`
    IncludeManagementChainData *IncludeManagementChainData `xml:"Include_Management_Chain_Data,omitempty"`
    IncludeMultipleManagersInManagementChainData *IncludeMultipleManagersInManagementChainData `xml:"Include_Multiple_Managers_in_Management_Chain_Data,omitempty"`
    IncludeBenefitEnrollments *IncludeBenefitEnrollments `xml:"Include_Benefit_Enrollments,omitempty"`
    IncludeBenefitEligibility *IncludeBenefitEligibility `xml:"Include_Benefit_Eligibility,omitempty"`
    IncludeRelatedPersons *IncludeRelatedPersons `xml:"Include_Related_Persons,omitempty"`
    IncludeQualifications *IncludeQualifications `xml:"Include_Qualifications,omitempty"`
    IncludeEmployeeReview *IncludeEmployeeReview `xml:"Include_Employee_Review,omitempty"`
    IncludeGoals *IncludeGoals `xml:"Include_Goals,omitempty"`
    IncludeDevelopmentItems *IncludeDevelopmentItems `xml:"Include_Development_Items,omitempty"`
    IncludeSkills *IncludeSkills `xml:"Include_Skills,omitempty"`
    IncludePhoto *IncludePhoto `xml:"Include_Photo,omitempty"`
    IncludeWorkerDocuments *IncludeWorkerDocuments `xml:"Include_Worker_Documents,omitempty"`
    IncludeTransactionLogData *IncludeTransactionLogData `xml:"Include_Transaction_Log_Data,omitempty"`
    IncludeSubeventsForCorrectedTransaction *IncludeSubeventsForCorrectedTransaction `xml:"Include_Subevents_for_Corrected_Transaction,omitempty"`
    IncludeSubeventsForRescindedTransaction *IncludeSubeventsForRescindedTransaction `xml:"Include_Subevents_for_Rescinded_Transaction,omitempty"`
    IncludeSuccessionProfile *IncludeSuccessionProfile `xml:"Include_Succession_Profile,omitempty"`
    IncludeTalentAssessment *IncludeTalentAssessment `xml:"Include_Talent_Assessment,omitempty"`
    IncludeEmployeeContractData *IncludeEmployeeContractData `xml:"Include_Employee_Contract_Data,omitempty"`
    IncludeContractsForTerminatedWorkers *IncludeContractsForTerminatedWorkers `xml:"Include_Contracts_for_Terminated_Workers,omitempty"`
    IncludeCollectiveAgreementData *IncludeCollectiveAgreementData `xml:"Include_Collective_Agreement_Data,omitempty"`
    IncludeProbationPeriodData *IncludeProbationPeriodData `xml:"Include_Probation_Period_Data,omitempty"`
    IncludeExtendedEmployeeContractDetails *IncludeExtendedEmployeeContractDetails `xml:"Include_Extended_Employee_Contract_Details,omitempty"`
    IncludeFeedbackReceived *IncludeFeedbackReceived `xml:"Include_Feedback_Received,omitempty"`
    IncludeUserAccount *IncludeUserAccount `xml:"Include_User_Account,omitempty"`
    IncludeCareer *IncludeCareer `xml:"Include_Career,omitempty"`
    IncludeAccountProvisioning *IncludeAccountProvisioning `xml:"Include_Account_Provisioning,omitempty"`
    IncludeBackgroundCheckData *IncludeBackgroundCheckData `xml:"Include_Background_Check_Data,omitempty"`
    IncludeContingentWorkerTaxAuthorityFormInformation *IncludeContingentWorkerTaxAuthorityFormInformation `xml:"Include_Contingent_Worker_Tax_Authority_Form_Information,omitempty"`
    ExcludeFunds *ExcludeFunds `xml:"Exclude_Funds,omitempty"`
    ExcludeFundHierarchies *ExcludeFundHierarchies `xml:"Exclude_Fund_Hierarchies,omitempty"`
    ExcludeGrants *ExcludeGrants `xml:"Exclude_Grants,omitempty"`
    ExcludeGrantHierarchies *ExcludeGrantHierarchies `xml:"Exclude_Grant_Hierarchies,omitempty"`
    ExcludeBusinessUnits *ExcludeBusinessUnits `xml:"Exclude_Business_Units,omitempty"`
    ExcludeBusinessUnitHierarchies *ExcludeBusinessUnitHierarchies `xml:"Exclude_Business_Unit_Hierarchies,omitempty"`
    ExcludePrograms *ExcludePrograms `xml:"Exclude_Programs,omitempty"`
    ExcludeProgramHierarchies *ExcludeProgramHierarchies `xml:"Exclude_Program_Hierarchies,omitempty"`
    ExcludeGifts *ExcludeGifts `xml:"Exclude_Gifts,omitempty"`
    ExcludeGiftHierarchies *ExcludeGiftHierarchies `xml:"Exclude_Gift_Hierarchies,omitempty"`
    ExcludeRetireeOrganizations *ExcludeRetireeOrganizations `xml:"Exclude_Retiree_Organizations,omitempty"`
}

type IncludeReference struct {
    Value *string `xml:",chardata"`
}

type IncludePersonalInformation struct {
    Value *string `xml:",chardata"`
}

type ShowAllPersonalInformation struct {
    Value *string `xml:",chardata"`
}

type IncludeAdditionalJobs struct {
    Value *string `xml:",chardata"`
}

type IncludeEmploymentInformation struct {
    Value *string `xml:",chardata"`
}

type IncludeCompensation struct {
    Value *string `xml:",chardata"`
}

type IncludeOrganizations struct {
    Value *string `xml:",chardata"`
}

type ExcludeOrganizationSupportRoleData struct {
    Value *string `xml:",chardata"`
}

type ExcludeLocationHierarchies struct {
    Value *string `xml:",chardata"`
}

type ExcludeCostCenters struct {
    Value *string `xml:",chardata"`
}

type ExcludeCostCenterHierarchies struct {
    Value *string `xml:",chardata"`
}

type ExcludeCompanies struct {
    Value *string `xml:",chardata"`
}

type ExcludeCompanyHierarchies struct {
    Value *string `xml:",chardata"`
}

type ExcludeMatrixOrganizations struct {
    Value *string `xml:",chardata"`
}

type ExcludePayGroups struct {
    Value *string `xml:",chardata"`
}

type ExcludeRegions struct {
    Value *string `xml:",chardata"`
}

type ExcludeRegionHierarchies struct {
    Value *string `xml:",chardata"`
}

type ExcludeSupervisoryOrganizations struct {
    Value *string `xml:",chardata"`
}

type ExcludeTeams struct {
    Value *string `xml:",chardata"`
}

type ExcludeCustomOrganizations struct {
    Value *string `xml:",chardata"`
}

type IncludeRoles struct {
    Value *string `xml:",chardata"`
}

type IncludeManagementChainData struct {
    Value *string `xml:",chardata"`
}

type IncludeMultipleManagersInManagementChainData struct {
    Value *string `xml:",chardata"`
}

type IncludeBenefitEnrollments struct {
    Value *string `xml:",chardata"`
}

type IncludeBenefitEligibility struct {
    Value *string `xml:",chardata"`
}

type IncludeRelatedPersons struct {
    Value *string `xml:",chardata"`
}

type IncludeQualifications struct {
    Value *string `xml:",chardata"`
}

type IncludeEmployeeReview struct {
    Value *string `xml:",chardata"`
}

type IncludeGoals struct {
    Value *string `xml:",chardata"`
}

type IncludeDevelopmentItems struct {
    Value *string `xml:",chardata"`
}

type IncludeSkills struct {
    Value *string `xml:",chardata"`
}

type IncludePhoto struct {
    Value *string `xml:",chardata"`
}

type IncludeWorkerDocuments struct {
    Value *string `xml:",chardata"`
}

type IncludeTransactionLogData struct {
    Value *string `xml:",chardata"`
}

type IncludeSubeventsForCorrectedTransaction struct {
    Value *string `xml:",chardata"`
}

type IncludeSubeventsForRescindedTransaction struct {
    Value *string `xml:",chardata"`
}

type IncludeSuccessionProfile struct {
    Value *string `xml:",chardata"`
}

type IncludeTalentAssessment struct {
    Value *string `xml:",chardata"`
}

type IncludeEmployeeContractData struct {
    Value *string `xml:",chardata"`
}

type IncludeContractsForTerminatedWorkers struct {
    Value *string `xml:",chardata"`
}

type IncludeCollectiveAgreementData struct {
    Value *string `xml:",chardata"`
}

type IncludeProbationPeriodData struct {
    Value *string `xml:",chardata"`
}

type IncludeExtendedEmployeeContractDetails struct {
    Value *string `xml:",chardata"`
}

type IncludeFeedbackReceived struct {
    Value *string `xml:",chardata"`
}

type IncludeUserAccount struct {
    Value *string `xml:",chardata"`
}

type IncludeCareer struct {
    Value *string `xml:",chardata"`
}

type IncludeAccountProvisioning struct {
    Value *string `xml:",chardata"`
}

type IncludeBackgroundCheckData struct {
    Value *string `xml:",chardata"`
}

type IncludeContingentWorkerTaxAuthorityFormInformation struct {
    Value *string `xml:",chardata"`
}

type ExcludeFunds struct {
    Value *string `xml:",chardata"`
}

type ExcludeFundHierarchies struct {
    Value *string `xml:",chardata"`
}

type ExcludeGrants struct {
    Value *string `xml:",chardata"`
}

type ExcludeGrantHierarchies struct {
    Value *string `xml:",chardata"`
}

type ExcludeBusinessUnits struct {
    Value *string `xml:",chardata"`
}

type ExcludeBusinessUnitHierarchies struct {
    Value *string `xml:",chardata"`
}

type ExcludePrograms struct {
    Value *string `xml:",chardata"`
}

type ExcludeProgramHierarchies struct {
    Value *string `xml:",chardata"`
}

type ExcludeGifts struct {
    Value *string `xml:",chardata"`
}

type ExcludeGiftHierarchies struct {
    Value *string `xml:",chardata"`
}

type ExcludeRetireeOrganizations struct {
    Value *string `xml:",chardata"`
}

type ResponseResults struct {
    TotalResults *TotalResults `xml:"Total_Results,omitempty"`
    TotalPages *TotalPages `xml:"Total_Pages,omitempty"`
    PageResults *PageResults `xml:"Page_Results,omitempty"`
    Page *Page `xml:"Page,omitempty"`
}

type TotalResults struct {
    Value *string `xml:",chardata"`
}

type TotalPages struct {
    Value *string `xml:",chardata"`
}

type PageResults struct {
    Value *string `xml:",chardata"`
}

type ResponseData struct {
    Worker *Worker `xml:"Worker,omitempty"`
}

type Worker struct {
    WorkerReference *WorkerReference `xml:"Worker_Reference,omitempty"`
    WorkerDescriptor *WorkerDescriptor `xml:"Worker_Descriptor,omitempty"`
    UniversalIdentifierReference *UniversalIdentifierReference `xml:"Universal_Identifier_Reference,omitempty"`
    WorkerData *WorkerData `xml:"Worker_Data,omitempty"`
}

type WorkerDescriptor struct {
    Value *string `xml:",chardata"`
}

type UniversalIdentifierReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WorkerData struct {
    WorkerId *WorkerId `xml:"Worker_ID,omitempty"`
    UserId *UserId `xml:"User_ID,omitempty"`
    UniversalId *UniversalId `xml:"Universal_ID,omitempty"`
    PersonalData *PersonalData `xml:"Personal_Data,omitempty"`
    EmploymentData *EmploymentData `xml:"Employment_Data,omitempty"`
}

type WorkerId struct {
    Value *string `xml:",chardata"`
}

type UserId struct {
    Value *string `xml:",chardata"`
}

type UniversalId struct {
    Value *string `xml:",chardata"`
}

type PersonalData struct {
    UniversalId *UniversalId `xml:"Universal_ID,omitempty"`
    NameData *NameData `xml:"Name_Data,omitempty"`
    PersonalInformationData *PersonalInformationData `xml:"Personal_Information_Data,omitempty"`
    IdentificationData *IdentificationData `xml:"Identification_Data,omitempty"`
    ContactData *ContactData `xml:"Contact_Data,omitempty"`
    TobaccoUse *TobaccoUse `xml:"Tobacco_Use,omitempty"`
}

type NameData struct {
    LegalNameData *LegalNameData `xml:"Legal_Name_Data,omitempty"`
    PreferredNameData *PreferredNameData `xml:"Preferred_Name_Data,omitempty"`
    AdditionalNameData *AdditionalNameData `xml:"Additional_Name_Data,omitempty"`
}

type LegalNameData struct {
    NameDetailData *NameDetailData `xml:"Name_Detail_Data,omitempty"`
}

type NameDetailData struct {
    FormattedName *string `xml:"Formatted_Name,attr,omitempty"`
    ReportingName *string `xml:"Reporting_Name,attr,omitempty"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"`
    PrefixData *PrefixData `xml:"Prefix_Data,omitempty"`
    FirstName *FirstName `xml:"First_Name,omitempty"`
    MiddleName *MiddleName `xml:"Middle_Name,omitempty"`
    LastName *LastName `xml:"Last_Name,omitempty"`
    SecondaryLastName *SecondaryLastName `xml:"Secondary_Last_Name,omitempty"`
    TertiaryLastName *TertiaryLastName `xml:"Tertiary_Last_Name,omitempty"`
    LocalNameDetailData *LocalNameDetailData `xml:"Local_Name_Detail_Data,omitempty"`
    SuffixData *SuffixData `xml:"Suffix_Data,omitempty"`
    FullNameForSingaporeAndMalaysia *FullNameForSingaporeAndMalaysia `xml:"Full_Name_for_Singapore_and_Malaysia,omitempty"`
}

type PrefixData struct {
    TitleReference *TitleReference `xml:"Title_Reference,omitempty"`
    TitleDescriptor *TitleDescriptor `xml:"Title_Descriptor,omitempty"`
    SalutationReference *SalutationReference `xml:"Salutation_Reference,omitempty"`
}

type TitleReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type TitleDescriptor struct {
    Value *string `xml:",chardata"`
}

type SalutationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type FirstName struct {
    Value *string `xml:",chardata"`
}

type MiddleName struct {
    Value *string `xml:",chardata"`
}

type LastName struct {
    Value *string `xml:",chardata"`
}

type SecondaryLastName struct {
    Value *string `xml:",chardata"`
}

type TertiaryLastName struct {
    Value *string `xml:",chardata"`
}

type LocalNameDetailData struct {
    LocalName *string `xml:"Local_Name,attr,omitempty"`
    LocalScript *string `xml:"Local_Script,attr,omitempty"`
    FirstName *FirstName `xml:"First_Name,omitempty"`
    MiddleName *MiddleName `xml:"Middle_Name,omitempty"`
    LastName *LastName `xml:"Last_Name,omitempty"`
    SecondaryLastName *SecondaryLastName `xml:"Secondary_Last_Name,omitempty"`
    FirstName2 *FirstName2 `xml:"First_Name_2,omitempty"`
    MiddleName2 *MiddleName2 `xml:"Middle_Name_2,omitempty"`
    LastName2 *LastName2 `xml:"Last_Name_2,omitempty"`
    SecondaryLastName2 *SecondaryLastName2 `xml:"Secondary_Last_Name_2,omitempty"`
}

type FirstName2 struct {
    Value *string `xml:",chardata"`
}

type MiddleName2 struct {
    Value *string `xml:",chardata"`
}

type LastName2 struct {
    Value *string `xml:",chardata"`
}

type SecondaryLastName2 struct {
    Value *string `xml:",chardata"`
}

type SuffixData struct {
    SocialSuffixReference *SocialSuffixReference `xml:"Social_Suffix_Reference,omitempty"`
    SocialSuffixDescriptor *SocialSuffixDescriptor `xml:"Social_Suffix_Descriptor,omitempty"`
    AcademicSuffixReference *AcademicSuffixReference `xml:"Academic_Suffix_Reference,omitempty"`
    HereditarySuffixReference *HereditarySuffixReference `xml:"Hereditary_Suffix_Reference,omitempty"`
    HonorarySuffixReference *HonorarySuffixReference `xml:"Honorary_Suffix_Reference,omitempty"`
    ProfessionalSuffixReference *ProfessionalSuffixReference `xml:"Professional_Suffix_Reference,omitempty"`
    ReligiousSuffixReference *ReligiousSuffixReference `xml:"Religious_Suffix_Reference,omitempty"`
    RoyalSuffixReference *RoyalSuffixReference `xml:"Royal_Suffix_Reference,omitempty"`
}

type SocialSuffixReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SocialSuffixDescriptor struct {
    Value *string `xml:",chardata"`
}

type AcademicSuffixReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type HereditarySuffixReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type HonorarySuffixReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ProfessionalSuffixReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ReligiousSuffixReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type RoyalSuffixReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type FullNameForSingaporeAndMalaysia struct {
    Value *string `xml:",chardata"`
}

type PreferredNameData struct {
    NameDetailData *NameDetailData `xml:"Name_Detail_Data,omitempty"`
}

type AdditionalNameData struct {
    NameDetailData *NameDetailData `xml:"Name_Detail_Data,omitempty"`
    NameTypeReference *NameTypeReference `xml:"Name_Type_Reference,omitempty"`
}

type NameTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PersonalInformationData struct {
    PersonalInformationForCountryData *PersonalInformationForCountryData `xml:"Personal_Information_For_Country_Data,omitempty"`
    BirthDate *BirthDate `xml:"Birth_Date,omitempty"`
    DateOfDeath *DateOfDeath `xml:"Date_of_Death,omitempty"`
    CountryOfBirthReference *CountryOfBirthReference `xml:"Country_of_Birth_Reference,omitempty"`
    CountryRegionOfBirthReference *CountryRegionOfBirthReference `xml:"Country_Region_of_Birth_Reference,omitempty"`
    RegionOfBirthDescriptor *RegionOfBirthDescriptor `xml:"Region_of_Birth_Descriptor,omitempty"`
    CityOfBirth *CityOfBirth `xml:"City_of_Birth,omitempty"`
    CityOfBirthReference *CityOfBirthReference `xml:"City_of_Birth_Reference,omitempty"`
    CitizenshipStatusReference *CitizenshipStatusReference `xml:"Citizenship_Status_Reference,omitempty"`
    PrimaryNationalityReference *PrimaryNationalityReference `xml:"Primary_Nationality_Reference,omitempty"`
    AdditionalNationalitiesReference *AdditionalNationalitiesReference `xml:"Additional_Nationalities_Reference,omitempty"`
    LastMedicalExamValidTo *LastMedicalExamValidTo `xml:"Last_Medical_Exam_Valid_To,omitempty"`
    LastMedicalExamDate *LastMedicalExamDate `xml:"Last_Medical_Exam_Date,omitempty"`
    MedicalExamNotes *MedicalExamNotes `xml:"Medical_Exam_Notes,omitempty"`
    BloodTypeReference *BloodTypeReference `xml:"Blood_Type_Reference,omitempty"`
    MilitaryServiceData *MilitaryServiceData `xml:"Military_Service_Data,omitempty"`
    SexualOrientationReference *SexualOrientationReference `xml:"Sexual_Orientation_Reference,omitempty"`
    GenderIdentityReference *GenderIdentityReference `xml:"Gender_Identity_Reference,omitempty"`
    PronounReference *PronounReference `xml:"Pronoun_Reference,omitempty"`
    NonCountrySpecificSectionData *NonCountrySpecificSectionData `xml:"Non_Country_Specific_Section_Data,omitempty"`
}

type PersonalInformationForCountryData struct {
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"`
    CountryPersonalInformationData *CountryPersonalInformationData `xml:"Country_Personal_Information_Data,omitempty"`
}

type CountryPersonalInformationData struct {
    MaritalStatusReference *MaritalStatusReference `xml:"Marital_Status_Reference,omitempty"`
    MaritalStatusDate *MaritalStatusDate `xml:"Marital_Status_Date,omitempty"`
    ReligionReference *ReligionReference `xml:"Religion_Reference,omitempty"`
    DisabilityStatusData *DisabilityStatusData `xml:"Disability_Status_Data,omitempty"`
    EthnicityReference *EthnicityReference `xml:"Ethnicity_Reference,omitempty"`
    RaceEthnicityDetailsReference *RaceEthnicityDetailsReference `xml:"Race_Ethnicity_Details_Reference,omitempty"`
    EthnicityVisualSurveyReference *EthnicityVisualSurveyReference `xml:"Ethnicity_Visual_Survey_Reference,omitempty"`
    HispanicOrLatino *HispanicOrLatino `xml:"Hispanic_or_Latino,omitempty"`
    HispanicOrLatinoVisualSurvey *HispanicOrLatinoVisualSurvey `xml:"Hispanic_or_Latino_Visual_Survey,omitempty"`
    AboriginalIndigenousIdentificationReference *AboriginalIndigenousIdentificationReference `xml:"Aboriginal_Indigenous_Identification_Reference,omitempty"`
    AboriginalIndigenousIdentificationDetailsReference *AboriginalIndigenousIdentificationDetailsReference `xml:"Aboriginal_Indigenous_Identification_Details_Reference,omitempty"`
    HukouRegionReference *HukouRegionReference `xml:"Hukou_Region_Reference,omitempty"`
    HukouSubregionReference *HukouSubregionReference `xml:"Hukou_Subregion_Reference,omitempty"`
    HukouLocality *HukouLocality `xml:"Hukou_Locality,omitempty"`
    HukouPostalCode *HukouPostalCode `xml:"Hukou_Postal_Code,omitempty"`
    HukouTypeReference *HukouTypeReference `xml:"Hukou_Type_Reference,omitempty"`
    LocalHukou *LocalHukou `xml:"Local_Hukou,omitempty"`
    NativeRegionReference *NativeRegionReference `xml:"Native_Region_Reference,omitempty"`
    NativeRegionDescriptor *NativeRegionDescriptor `xml:"Native_Region_Descriptor,omitempty"`
    PersonnelFileAgencyForPerson *PersonnelFileAgencyForPerson `xml:"Personnel_File_Agency_for_Person,omitempty"`
    PoliticalAffiliationReference *PoliticalAffiliationReference `xml:"Political_Affiliation_Reference,omitempty"`
    SocialBenefitsLocalityReference *SocialBenefitsLocalityReference `xml:"Social_Benefits_Locality_Reference,omitempty"`
    RelativeNameInformationData *RelativeNameInformationData `xml:"Relative_Name_Information_Data,omitempty"`
    SexualOrientationAndGenderIdentityReference *SexualOrientationAndGenderIdentityReference `xml:"Sexual_Orientation_and_Gender_Identity_Reference,omitempty"`
    GenderReference *GenderReference `xml:"Gender_Reference,omitempty"`
    ReportingGenderReference *ReportingGenderReference `xml:"Reporting_Gender_Reference,omitempty"`
    VeteransPreferenceReference *VeteransPreferenceReference `xml:"Veterans_Preference_Reference,omitempty"`
    VeteransPreferenceForRifReference *VeteransPreferenceForRifReference `xml:"Veterans_Preference_For_RIF_Reference,omitempty"`
    SelectiveServiceRegistrationReference *SelectiveServiceRegistrationReference `xml:"Selective_Service_Registration_Reference,omitempty"`
    DisabledVeteranLeaveDateReference *DisabledVeteranLeaveDateReference `xml:"Disabled_Veteran_Leave_Date_Reference,omitempty"`
    ActiveMilitaryUniformedServiceReference *ActiveMilitaryUniformedServiceReference `xml:"Active_Military_Uniformed_Service_Reference,omitempty"`
    UniformServiceReserveStatusReference *UniformServiceReserveStatusReference `xml:"Uniform_Service_Reserve_Status_Reference,omitempty"`
    CountrySpecificSectionData *CountrySpecificSectionData `xml:"Country_Specific_Section_Data,omitempty"`
}

type MaritalStatusReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type MaritalStatusDate struct {
    Value *string `xml:",chardata"`
}

type ReligionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DisabilityStatusData struct {
    DisabilityReference *DisabilityReference `xml:"Disability_Reference,omitempty"`
    DisabilityStatusDate *DisabilityStatusDate `xml:"Disability_Status_Date,omitempty"`
    DisabilityDateKnown *DisabilityDateKnown `xml:"Disability_Date_Known,omitempty"`
    DisabilityEndDate *DisabilityEndDate `xml:"Disability_End_Date,omitempty"`
    DisabilityGradeReference *DisabilityGradeReference `xml:"Disability_Grade_Reference,omitempty"`
    DisabilityDegree *DisabilityDegree `xml:"Disability_Degree,omitempty"`
    DisabilityRemainingCapacity *DisabilityRemainingCapacity `xml:"Disability_Remaining_Capacity,omitempty"`
    DisabilityCertificationAuthorityReference *DisabilityCertificationAuthorityReference `xml:"Disability_Certification_Authority_Reference,omitempty"`
    DisabilityCertificationAuthority *DisabilityCertificationAuthority `xml:"Disability_Certification_Authority,omitempty"`
    DisabilityCertifiedAt *DisabilityCertifiedAt `xml:"Disability_Certified_At,omitempty"`
    DisabilityCertificationId *DisabilityCertificationId `xml:"Disability_Certification_ID,omitempty"`
    DisabilityCertificationBasisReference *DisabilityCertificationBasisReference `xml:"Disability_Certification_Basis_Reference,omitempty"`
    DisabilitySeverityRecognitionDate *DisabilitySeverityRecognitionDate `xml:"Disability_Severity_Recognition_Date,omitempty"`
    DisabilityFteTowardQuota *DisabilityFteTowardQuota `xml:"Disability_FTE_Toward_Quota,omitempty"`
    DisabilityWorkRestrictions *DisabilityWorkRestrictions `xml:"Disability_Work_Restrictions,omitempty"`
    DisabilityAccommodationsRequested *DisabilityAccommodationsRequested `xml:"Disability_Accommodations_Requested,omitempty"`
    DisabilityAccommodationsProvided *DisabilityAccommodationsProvided `xml:"Disability_Accommodations_Provided,omitempty"`
    DisabilityRehabilitationRequested *DisabilityRehabilitationRequested `xml:"Disability_Rehabilitation_Requested,omitempty"`
    DisabilityRehabilitationProvided *DisabilityRehabilitationProvided `xml:"Disability_Rehabilitation_Provided,omitempty"`
    Note *Note `xml:"Note,omitempty"`
    WorkerDocumentReference *WorkerDocumentReference `xml:"Worker_Document_Reference,omitempty"`
    DisabilityStatusReference *DisabilityStatusReference `xml:"Disability_Status_Reference,omitempty"`
}

type DisabilityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DisabilityStatusDate struct {
    Value *string `xml:",chardata"`
}

type DisabilityDateKnown struct {
    Value *string `xml:",chardata"`
}

type DisabilityEndDate struct {
    Value *string `xml:",chardata"`
}

type DisabilityGradeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DisabilityDegree struct {
    Value *string `xml:",chardata"`
}

type DisabilityRemainingCapacity struct {
    Value *string `xml:",chardata"`
}

type DisabilityCertificationAuthorityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DisabilityCertificationAuthority struct {
    Value *string `xml:",chardata"`
}

type DisabilityCertifiedAt struct {
    Value *string `xml:",chardata"`
}

type DisabilityCertificationId struct {
    Value *string `xml:",chardata"`
}

type DisabilityCertificationBasisReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DisabilitySeverityRecognitionDate struct {
    Value *string `xml:",chardata"`
}

type DisabilityFteTowardQuota struct {
    Value *string `xml:",chardata"`
}

type DisabilityWorkRestrictions struct {
    Value *string `xml:",chardata"`
}

type DisabilityAccommodationsRequested struct {
    Value *string `xml:",chardata"`
}

type DisabilityAccommodationsProvided struct {
    Value *string `xml:",chardata"`
}

type DisabilityRehabilitationRequested struct {
    Value *string `xml:",chardata"`
}

type DisabilityRehabilitationProvided struct {
    Value *string `xml:",chardata"`
}

type Note struct {
    Value *string `xml:",chardata"`
}

type WorkerDocumentReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DisabilityStatusReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type EthnicityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type RaceEthnicityDetailsReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type EthnicityVisualSurveyReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type HispanicOrLatino struct {
    Value *string `xml:",chardata"`
}

type HispanicOrLatinoVisualSurvey struct {
    Value *string `xml:",chardata"`
}

type AboriginalIndigenousIdentificationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type AboriginalIndigenousIdentificationDetailsReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type HukouRegionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type HukouSubregionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type HukouLocality struct {
    Value *string `xml:",chardata"`
}

type HukouPostalCode struct {
    Value *string `xml:",chardata"`
}

type HukouTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type LocalHukou struct {
    Value *string `xml:",chardata"`
}

type NativeRegionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type NativeRegionDescriptor struct {
    Value *string `xml:",chardata"`
}

type PersonnelFileAgencyForPerson struct {
    Value *string `xml:",chardata"`
}

type PoliticalAffiliationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SocialBenefitsLocalityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type RelativeNameInformationData struct {
    Delete *string `xml:"Delete,attr,omitempty"`
    RelativeNameReferenceReference *RelativeNameReferenceReference `xml:"Relative_Name_Reference_Reference,omitempty"`
    RelativeTypeReference *RelativeTypeReference `xml:"Relative_Type_Reference,omitempty"`
    NameDetailData *NameDetailData `xml:"Name_Detail_Data,omitempty"`
}

type RelativeNameReferenceReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type RelativeTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SexualOrientationAndGenderIdentityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type GenderReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ReportingGenderReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type VeteransPreferenceReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type VeteransPreferenceForRifReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SelectiveServiceRegistrationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DisabledVeteranLeaveDateReference struct {
    Value *string `xml:",chardata"`
}

type ActiveMilitaryUniformedServiceReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type UniformServiceReserveStatusReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CountrySpecificSectionData struct {
    CountrySpecificSection1Data *CountrySpecificSection1Data `xml:"Country_Specific_Section_1_Data,omitempty"`
    CountrySpecificSection2Data *CountrySpecificSection2Data `xml:"Country_Specific_Section_2_Data,omitempty"`
    CountrySpecificSection3Data *CountrySpecificSection3Data `xml:"Country_Specific_Section_3_Data,omitempty"`
}

type CountrySpecificSection1Data struct {
    Field1Reference *Field1Reference `xml:"Field_1_Reference,omitempty"`
    Field2Reference *Field2Reference `xml:"Field_2_Reference,omitempty"`
    Field3Reference *Field3Reference `xml:"Field_3_Reference,omitempty"`
}

type Field1Reference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Field2Reference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Field3Reference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CountrySpecificSection2Data struct {
    Field1Reference *Field1Reference `xml:"Field_1_Reference,omitempty"`
    Field2Reference *Field2Reference `xml:"Field_2_Reference,omitempty"`
    Field3Reference *Field3Reference `xml:"Field_3_Reference,omitempty"`
}

type CountrySpecificSection3Data struct {
    Field1Reference *Field1Reference `xml:"Field_1_Reference,omitempty"`
    Field2Reference *Field2Reference `xml:"Field_2_Reference,omitempty"`
    Field3Reference *Field3Reference `xml:"Field_3_Reference,omitempty"`
}

type BirthDate struct {
    Value *string `xml:",chardata"`
}

type DateOfDeath struct {
    Value *string `xml:",chardata"`
}

type CountryOfBirthReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CountryRegionOfBirthReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type RegionOfBirthDescriptor struct {
    Value *string `xml:",chardata"`
}

type CityOfBirth struct {
    Value *string `xml:",chardata"`
}

type CityOfBirthReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CitizenshipStatusReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PrimaryNationalityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type AdditionalNationalitiesReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type LastMedicalExamValidTo struct {
    Value *string `xml:",chardata"`
}

type LastMedicalExamDate struct {
    Value *string `xml:",chardata"`
}

type MedicalExamNotes struct {
    Value *string `xml:",chardata"`
}

type BloodTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type MilitaryServiceData struct {
    StatusReference *StatusReference `xml:"Status_Reference,omitempty"`
    DischargeDate *DischargeDate `xml:"Discharge_Date,omitempty"`
    StatusBeginDate *StatusBeginDate `xml:"Status_Begin_Date,omitempty"`
    MilitaryServiceTypeReference *MilitaryServiceTypeReference `xml:"Military_Service_Type_Reference,omitempty"`
    MilitaryRankReference *MilitaryRankReference `xml:"Military_Rank_Reference,omitempty"`
    Notes *Notes `xml:"Notes,omitempty"`
    MilitaryServiceReference *MilitaryServiceReference `xml:"Military_Service_Reference,omitempty"`
    MilitaryDischargeTypeReference *MilitaryDischargeTypeReference `xml:"Military_Discharge_Type_Reference,omitempty"`
}

type StatusReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DischargeDate struct {
    Value *string `xml:",chardata"`
}

type StatusBeginDate struct {
    Value *string `xml:",chardata"`
}

type MilitaryServiceTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type MilitaryRankReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Notes struct {
    Value *string `xml:",chardata"`
}

type MilitaryServiceReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type MilitaryDischargeTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SexualOrientationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type GenderIdentityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PronounReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type NonCountrySpecificSectionData struct {
    NonCountrySpecificSection1Data *NonCountrySpecificSection1Data `xml:"Non_Country_Specific_Section_1_Data,omitempty"`
    NonCountrySpecificSection2Data *NonCountrySpecificSection2Data `xml:"Non_Country_Specific_Section_2_Data,omitempty"`
    NonCountrySpecificSection3Data *NonCountrySpecificSection3Data `xml:"Non_Country_Specific_Section_3_Data,omitempty"`
}

type NonCountrySpecificSection1Data struct {
    Field1Reference *Field1Reference `xml:"Field_1_Reference,omitempty"`
    Field2Reference *Field2Reference `xml:"Field_2_Reference,omitempty"`
    Field3Reference *Field3Reference `xml:"Field_3_Reference,omitempty"`
}

type NonCountrySpecificSection2Data struct {
    Field1Reference *Field1Reference `xml:"Field_1_Reference,omitempty"`
    Field2Reference *Field2Reference `xml:"Field_2_Reference,omitempty"`
    Field3Reference *Field3Reference `xml:"Field_3_Reference,omitempty"`
}

type NonCountrySpecificSection3Data struct {
    Field1Reference *Field1Reference `xml:"Field_1_Reference,omitempty"`
    Field2Reference *Field2Reference `xml:"Field_2_Reference,omitempty"`
    Field3Reference *Field3Reference `xml:"Field_3_Reference,omitempty"`
}

type IdentificationData struct {
    NationalId *NationalId `xml:"National_ID,omitempty"`
    GovernmentId *GovernmentId `xml:"Government_ID,omitempty"`
    VisaId *VisaId `xml:"Visa_ID,omitempty"`
    PassportId *PassportId `xml:"Passport_ID,omitempty"`
    LicenseId *LicenseId `xml:"License_ID,omitempty"`
    CustomId *CustomId `xml:"Custom_ID,omitempty"`
}

type NationalId struct {
    Delete *string `xml:"Delete,attr,omitempty"`
    NationalIdReference *NationalIdReference `xml:"National_ID_Reference,omitempty"`
    NationalIdData *NationalIdData `xml:"National_ID_Data,omitempty"`
    NationalIdSharedReference *NationalIdSharedReference `xml:"National_ID_Shared_Reference,omitempty"`
}

type NationalIdReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type NationalIdData struct {
    Id *Id `xml:"ID,omitempty"`
    IdTypeReference *IdTypeReference `xml:"ID_Type_Reference,omitempty"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"`
    IssuedDate *IssuedDate `xml:"Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"Expiration_Date,omitempty"`
    VerificationDate *VerificationDate `xml:"Verification_Date,omitempty"`
    Series *Series `xml:"Series,omitempty"`
    IssuingAgency *IssuingAgency `xml:"Issuing_Agency,omitempty"`
    VerifiedByReference *VerifiedByReference `xml:"Verified_By_Reference,omitempty"`
}

type IdTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type IssuedDate struct {
    Value *string `xml:",chardata"`
}

type ExpirationDate struct {
    Value *string `xml:",chardata"`
}

type VerificationDate struct {
    Value *string `xml:",chardata"`
}

type Series struct {
    Value *string `xml:",chardata"`
}

type IssuingAgency struct {
    Value *string `xml:",chardata"`
}

type VerifiedByReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type NationalIdSharedReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type GovernmentId struct {
    Delete *string `xml:"Delete,attr,omitempty"`
    GovernmentIdReference *GovernmentIdReference `xml:"Government_ID_Reference,omitempty"`
    GovernmentIdData *GovernmentIdData `xml:"Government_ID_Data,omitempty"`
    GovernmentIdSharedReference *GovernmentIdSharedReference `xml:"Government_ID_Shared_Reference,omitempty"`
}

type GovernmentIdReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type GovernmentIdData struct {
    Id *Id `xml:"ID,omitempty"`
    IdTypeReference *IdTypeReference `xml:"ID_Type_Reference,omitempty"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"`
    IssuedDate *IssuedDate `xml:"Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"Expiration_Date,omitempty"`
    VerificationDate *VerificationDate `xml:"Verification_Date,omitempty"`
    VerifiedByReference *VerifiedByReference `xml:"Verified_by_Reference,omitempty"`
}

type GovernmentIdSharedReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type VisaId struct {
    Delete *string `xml:"Delete,attr,omitempty"`
    VisaIdReference *VisaIdReference `xml:"Visa_ID_Reference,omitempty"`
    VisaIdData *VisaIdData `xml:"Visa_ID_Data,omitempty"`
    VisaIdSharedReference *VisaIdSharedReference `xml:"Visa_ID_Shared_Reference,omitempty"`
}

type VisaIdReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type VisaIdData struct {
    Id *Id `xml:"ID,omitempty"`
    IdTypeReference *IdTypeReference `xml:"ID_Type_Reference,omitempty"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"`
    IssuedDate *IssuedDate `xml:"Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"Expiration_Date,omitempty"`
    VerificationDate *VerificationDate `xml:"Verification_Date,omitempty"`
    VerifiedByReference *VerifiedByReference `xml:"Verified_By_Reference,omitempty"`
}

type VisaIdSharedReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PassportId struct {
    Delete *string `xml:"Delete,attr,omitempty"`
    PassportIdReference *PassportIdReference `xml:"Passport_ID_Reference,omitempty"`
    PassportIdData *PassportIdData `xml:"Passport_ID_Data,omitempty"`
    PassportIdSharedReference *PassportIdSharedReference `xml:"Passport_ID_Shared_Reference,omitempty"`
}

type PassportIdReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PassportIdData struct {
    Id *Id `xml:"ID,omitempty"`
    IdTypeReference *IdTypeReference `xml:"ID_Type_Reference,omitempty"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"`
    IssuedDate *IssuedDate `xml:"Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"Expiration_Date,omitempty"`
    VerificationDate *VerificationDate `xml:"Verification_Date,omitempty"`
    VerifiedByReference *VerifiedByReference `xml:"Verified_By_Reference,omitempty"`
}

type PassportIdSharedReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type LicenseId struct {
    Delete *string `xml:"Delete,attr,omitempty"`
    LicenseIdReference *LicenseIdReference `xml:"License_ID_Reference,omitempty"`
    LicenseIdData *LicenseIdData `xml:"License_ID_Data,omitempty"`
    LicenseIdSharedReference *LicenseIdSharedReference `xml:"License_ID_Shared_Reference,omitempty"`
}

type LicenseIdReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type LicenseIdData struct {
    Id *Id `xml:"ID,omitempty"`
    IdTypeReference *IdTypeReference `xml:"ID_Type_Reference,omitempty"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"`
    CountryRegionReference *CountryRegionReference `xml:"Country_Region_Reference,omitempty"`
    CountryRegionDescriptor *CountryRegionDescriptor `xml:"Country_Region_Descriptor,omitempty"`
    AuthorityReference *AuthorityReference `xml:"Authority_Reference,omitempty"`
    LicenseClass *LicenseClass `xml:"License_Class,omitempty"`
    IssuedDate *IssuedDate `xml:"Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"Expiration_Date,omitempty"`
    VerificationDate *VerificationDate `xml:"Verification_Date,omitempty"`
    VerifiedByReference *VerifiedByReference `xml:"Verified_By_Reference,omitempty"`
}

type CountryRegionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CountryRegionDescriptor struct {
    Value *string `xml:",chardata"`
}

type AuthorityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type LicenseClass struct {
    Value *string `xml:",chardata"`
}

type LicenseIdSharedReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CustomId struct {
    Delete *string `xml:"Delete,attr,omitempty"`
    CustomIdReference *CustomIdReference `xml:"Custom_ID_Reference,omitempty"`
    CustomIdData *CustomIdData `xml:"Custom_ID_Data,omitempty"`
    CustomIdSharedReference *CustomIdSharedReference `xml:"Custom_ID_Shared_Reference,omitempty"`
}

type CustomIdReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CustomIdData struct {
    Id *Id `xml:"ID,omitempty"`
    IdTypeReference *IdTypeReference `xml:"ID_Type_Reference,omitempty"`
    IssuedDate *IssuedDate `xml:"Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"Expiration_Date,omitempty"`
    OrganizationIdReference *OrganizationIdReference `xml:"Organization_ID_Reference,omitempty"`
    CustomDescription *CustomDescription `xml:"Custom_Description,omitempty"`
}

type OrganizationIdReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CustomDescription struct {
    Value *string `xml:",chardata"`
}

type CustomIdSharedReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ContactData struct {
    AddressData *AddressData `xml:"Address_Data,omitempty"`
    PhoneData *PhoneData `xml:"Phone_Data,omitempty"`
    EmailAddressData *EmailAddressData `xml:"Email_Address_Data,omitempty"`
    InstantMessengerData *InstantMessengerData `xml:"Instant_Messenger_Data,omitempty"`
    WebAddressData *WebAddressData `xml:"Web_Address_Data,omitempty"`
}

type AddressData struct {
    FormattedAddress *string `xml:"Formatted_Address,attr,omitempty"`
    AddressFormatType *string `xml:"Address_Format_Type,attr,omitempty"`
    DefaultedBusinessSiteAddress *string `xml:"Defaulted_Business_Site_Address,attr,omitempty"`
    Delete *string `xml:"Delete,attr,omitempty"`
    DoNotReplaceAll *string `xml:"Do_Not_Replace_All,attr,omitempty"`
    EffectiveDate *string `xml:"Effective_Date,attr,omitempty"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"`
    LastModified *LastModified `xml:"Last_Modified,omitempty"`
    AddressLineData *AddressLineData `xml:"Address_Line_Data,omitempty"`
    Municipality *Municipality `xml:"Municipality,omitempty"`
    CountryCityReference *CountryCityReference `xml:"Country_City_Reference,omitempty"`
    SubmunicipalityData *SubmunicipalityData `xml:"Submunicipality_Data,omitempty"`
    CountryRegionReference *CountryRegionReference `xml:"Country_Region_Reference,omitempty"`
    CountryRegionDescriptor *CountryRegionDescriptor `xml:"Country_Region_Descriptor,omitempty"`
    SubregionData *SubregionData `xml:"Subregion_Data,omitempty"`
    PostalCode *PostalCode `xml:"Postal_Code,omitempty"`
    UsageData *UsageData `xml:"Usage_Data,omitempty"`
    NumberOfDays *NumberOfDays `xml:"Number_of_Days,omitempty"`
    MunicipalityLocal *MunicipalityLocal `xml:"Municipality_Local,omitempty"`
    AddressReference *AddressReference `xml:"Address_Reference,omitempty"`
    AddressId *AddressId `xml:"Address_ID,omitempty"`
}

type LastModified struct {
    Value *string `xml:",chardata"`
}

type AddressLineData struct {
    Value *string `xml:",chardata"`
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Type *string `xml:"Type,attr,omitempty"`
}

type Municipality struct {
    Value *string `xml:",chardata"`
}

type CountryCityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SubmunicipalityData struct {
    Value *string `xml:",chardata"`
    AddressComponentName *string `xml:"Address_Component_Name,attr,omitempty"`
    Type *string `xml:"Type,attr,omitempty"`
}

type SubregionData struct {
    Value *string `xml:",chardata"`
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Type *string `xml:"Type,attr,omitempty"`
}

type PostalCode struct {
    Value *string `xml:",chardata"`
}

type UsageData struct {
    Public *string `xml:"Public,attr,omitempty"`
    TypeData *TypeData `xml:"Type_Data,omitempty"`
    UseForReference *UseForReference `xml:"Use_For_Reference,omitempty"`
    UseForTenantedReference *UseForTenantedReference `xml:"Use_For_Tenanted_Reference,omitempty"`
    Comments *Comments `xml:"Comments,omitempty"`
}

type TypeData struct {
    Primary *string `xml:"Primary,attr,omitempty"`
    TypeReference *TypeReference `xml:"Type_Reference,omitempty"`
}

type TypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type UseForReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type UseForTenantedReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Comments struct {
    Value *string `xml:",chardata"`
}

type NumberOfDays struct {
    Value *string `xml:",chardata"`
}

type MunicipalityLocal struct {
    Value *string `xml:",chardata"`
}

type AddressReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type AddressId struct {
    Value *string `xml:",chardata"`
}

type PhoneData struct {
    AreaCode *string `xml:"Area_Code,attr,omitempty"`
    TenantFormattedPhone *string `xml:"Tenant_Formatted_Phone,attr,omitempty"`
    InternationalFormattedPhone *string `xml:"International_Formatted_Phone,attr,omitempty"`
    PhoneNumberWithoutAreaCode *string `xml:"Phone_Number_Without_Area_Code,attr,omitempty"`
    NationalFormattedPhone *string `xml:"National_Formatted_Phone,attr,omitempty"`
    E164FormattedPhone *string `xml:"E164_Formatted_Phone,attr,omitempty"`
    WorkdayTraditionalFormattedPhone *string `xml:"Workday_Traditional_Formatted_Phone,attr,omitempty"`
    Delete *string `xml:"Delete,attr,omitempty"`
    DoNotReplaceAll *string `xml:"Do_Not_Replace_All,attr,omitempty"`
    CountryIsoCode *CountryIsoCode `xml:"Country_ISO_Code,omitempty"`
    InternationalPhoneCode *InternationalPhoneCode `xml:"International_Phone_Code,omitempty"`
    PhoneNumber *PhoneNumber `xml:"Phone_Number,omitempty"`
    PhoneExtension *PhoneExtension `xml:"Phone_Extension,omitempty"`
    PhoneDeviceTypeReference *PhoneDeviceTypeReference `xml:"Phone_Device_Type_Reference,omitempty"`
    UsageData *UsageData `xml:"Usage_Data,omitempty"`
    PhoneReference *PhoneReference `xml:"Phone_Reference,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CountryIsoCode struct {
    Value *string `xml:",chardata"`
}

type InternationalPhoneCode struct {
    Value *string `xml:",chardata"`
}

type PhoneNumber struct {
    Value *string `xml:",chardata"`
}

type PhoneExtension struct {
    Value *string `xml:",chardata"`
}

type PhoneDeviceTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PhoneReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type EmailAddressData struct {
    Delete *string `xml:"Delete,attr,omitempty"`
    DoNotReplaceAll *string `xml:"Do_Not_Replace_All,attr,omitempty"`
    EmailAddress *EmailAddress `xml:"Email_Address,omitempty"`
    EmailComment *EmailComment `xml:"Email_Comment,omitempty"`
    UsageData *UsageData `xml:"Usage_Data,omitempty"`
    EmailReference *EmailReference `xml:"Email_Reference,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type EmailAddress struct {
    Value *string `xml:",chardata"`
}

type EmailComment struct {
    Value *string `xml:",chardata"`
}

type EmailReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type InstantMessengerData struct {
    Delete *string `xml:"Delete,attr,omitempty"`
    DoNotReplaceAll *string `xml:"Do_Not_Replace_All,attr,omitempty"`
    InstantMessengerAddress *InstantMessengerAddress `xml:"Instant_Messenger_Address,omitempty"`
    InstantMessengerTypeReference *InstantMessengerTypeReference `xml:"Instant_Messenger_Type_Reference,omitempty"`
    InstantMessengerComment *InstantMessengerComment `xml:"Instant_Messenger_Comment,omitempty"`
    UsageData *UsageData `xml:"Usage_Data,omitempty"`
    InstantMessengerReference *InstantMessengerReference `xml:"Instant_Messenger_Reference,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type InstantMessengerAddress struct {
    Value *string `xml:",chardata"`
}

type InstantMessengerTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type InstantMessengerComment struct {
    Value *string `xml:",chardata"`
}

type InstantMessengerReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WebAddressData struct {
    Delete *string `xml:"Delete,attr,omitempty"`
    DoNotReplaceAll *string `xml:"Do_Not_Replace_All,attr,omitempty"`
    WebAddress *WebAddress `xml:"Web_Address,omitempty"`
    WebAddressComment *WebAddressComment `xml:"Web_Address_Comment,omitempty"`
    UsageData *UsageData `xml:"Usage_Data,omitempty"`
    WebAddressReference *WebAddressReference `xml:"Web_Address_Reference,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WebAddress struct {
    Value *string `xml:",chardata"`
}

type WebAddressComment struct {
    Value *string `xml:",chardata"`
}

type WebAddressReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type TobaccoUse struct {
    Value *string `xml:",chardata"`
}

type EmploymentData struct {
    WorkerJobData *WorkerJobData `xml:"Worker_Job_Data,omitempty"`
    WorkerStatusData *WorkerStatusData `xml:"Worker_Status_Data,omitempty"`
}

type WorkerJobData struct {
    PrimaryJob *string `xml:"Primary_Job,attr,omitempty"`
    PositionData *PositionData `xml:"Position_Data,omitempty"`
    PositionOrganizationsData *PositionOrganizationsData `xml:"Position_Organizations_Data,omitempty"`
    PositionManagementChainsData *PositionManagementChainsData `xml:"Position_Management_Chains_Data,omitempty"`
    PositionContractDetailData *PositionContractDetailData `xml:"Position_Contract_Detail_Data,omitempty"`
}

type PositionData struct {
    EffectiveDate *string `xml:"Effective_Date,attr,omitempty"`
    PositionReference *PositionReference `xml:"Position_Reference,omitempty"`
    HeadcountReference *HeadcountReference `xml:"Headcount_Reference,omitempty"`
    PositionId *PositionId `xml:"Position_ID,omitempty"`
    PositionTitle *PositionTitle `xml:"Position_Title,omitempty"`
    BusinessTitle *BusinessTitle `xml:"Business_Title,omitempty"`
    StartDate *StartDate `xml:"Start_Date,omitempty"`
    EndEmploymentDate *EndEmploymentDate `xml:"End_Employment_Date,omitempty"`
    EndEmploymentReasonReference *EndEmploymentReasonReference `xml:"End_Employment_Reason_Reference,omitempty"`
    WorkerTypeReference *WorkerTypeReference `xml:"Worker_Type_Reference,omitempty"`
    PositionTimeTypeReference *PositionTimeTypeReference `xml:"Position_Time_Type_Reference,omitempty"`
    JobExempt *JobExempt `xml:"Job_Exempt,omitempty"`
    ScheduledWeeklyHours *ScheduledWeeklyHours `xml:"Scheduled_Weekly_Hours,omitempty"`
    DefaultWeeklyHours *DefaultWeeklyHours `xml:"Default_Weekly_Hours,omitempty"`
    WorkingTimeFrequencyReference *WorkingTimeFrequencyReference `xml:"Working_Time_Frequency_Reference,omitempty"`
    WorkingTimeUnitReference *WorkingTimeUnitReference `xml:"Working_Time_Unit_Reference,omitempty"`
    WorkingTimeValue *WorkingTimeValue `xml:"Working_Time_Value,omitempty"`
    FullTimeEquivalentPercentage *FullTimeEquivalentPercentage `xml:"Full_Time_Equivalent_Percentage,omitempty"`
    SpecifyPaidFte *SpecifyPaidFte `xml:"Specify_Paid_FTE,omitempty"`
    PaidFte *PaidFte `xml:"Paid_FTE,omitempty"`
    SpecifyWorkingFte *SpecifyWorkingFte `xml:"Specify_Working_FTE,omitempty"`
    WorkingFte *WorkingFte `xml:"Working_FTE,omitempty"`
    ExcludeFromHeadcount *ExcludeFromHeadcount `xml:"Exclude_from_Headcount,omitempty"`
    PayRateTypeReference *PayRateTypeReference `xml:"Pay_Rate_Type_Reference,omitempty"`
    JobClassificationSummaryData *JobClassificationSummaryData `xml:"Job_Classification_Summary_Data,omitempty"`
    CompanyInsiderReference *CompanyInsiderReference `xml:"Company_Insider_Reference,omitempty"`
    WorkShiftReference *WorkShiftReference `xml:"Work_Shift_Reference,omitempty"`
    WorkHoursProfilesReference *WorkHoursProfilesReference `xml:"Work_Hours_Profiles_Reference,omitempty"`
    FederalWithholdingFein *FederalWithholdingFein `xml:"Federal_Withholding_FEIN,omitempty"`
    WorkerCompensationCodeData *WorkerCompensationCodeData `xml:"Worker_Compensation_Code_Data,omitempty"`
    PositionPayrollReportingCodeData *PositionPayrollReportingCodeData `xml:"Position_Payroll_Reporting_Code_Data,omitempty"`
    JobProfileSummaryData *JobProfileSummaryData `xml:"Job_Profile_Summary_Data,omitempty"`
    BusinessSiteSummaryData *BusinessSiteSummaryData `xml:"Business_Site_Summary_Data,omitempty"`
    PayrollInterfaceProcessingData *PayrollInterfaceProcessingData `xml:"Payroll_Interface_Processing_Data,omitempty"`
    RegularPaidEquivalentHours *RegularPaidEquivalentHours `xml:"Regular_Paid_Equivalent_Hours,omitempty"`
    WorkerHoursProfileClassification *WorkerHoursProfileClassification `xml:"Worker_Hours_Profile_Classification,omitempty"`
    InternationalAssignmentData *InternationalAssignmentData `xml:"International_Assignment_Data,omitempty"`
    WorkSpaceReference *WorkSpaceReference `xml:"Work_Space__Reference,omitempty"`
    AcademicPaySetupData *AcademicPaySetupData `xml:"Academic_Pay_Setup_Data,omitempty"`
    EndDate *EndDate `xml:"End_Date,omitempty"`
    PayThroughDate *PayThroughDate `xml:"Pay_Through_Date,omitempty"`
    CollectiveAgreementSummaryData *CollectiveAgreementSummaryData `xml:"Collective_Agreement_Summary_Data,omitempty"`
    EmployeeProbationPeriodSummaryData *EmployeeProbationPeriodSummaryData `xml:"Employee_Probation_Period_Summary_Data,omitempty"`
    ManagerAsOfLastDetectedManagerChangeReference *ManagerAsOfLastDetectedManagerChangeReference `xml:"Manager_as_of_last_detected_manager_change_Reference,omitempty"`
}

type HeadcountReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PositionId struct {
    Value *string `xml:",chardata"`
}

type PositionTitle struct {
    Value *string `xml:",chardata"`
}

type BusinessTitle struct {
    Value *string `xml:",chardata"`
}

type StartDate struct {
    Value *string `xml:",chardata"`
}

type EndEmploymentDate struct {
    Value *string `xml:",chardata"`
}

type EndEmploymentReasonReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WorkerTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PositionTimeTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type JobExempt struct {
    Value *string `xml:",chardata"`
}

type ScheduledWeeklyHours struct {
    Value *string `xml:",chardata"`
}

type DefaultWeeklyHours struct {
    Value *string `xml:",chardata"`
}

type WorkingTimeFrequencyReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WorkingTimeUnitReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WorkingTimeValue struct {
    Value *string `xml:",chardata"`
}

type FullTimeEquivalentPercentage struct {
    Value *string `xml:",chardata"`
}

type SpecifyPaidFte struct {
    Value *string `xml:",chardata"`
}

type PaidFte struct {
    Value *string `xml:",chardata"`
}

type SpecifyWorkingFte struct {
    Value *string `xml:",chardata"`
}

type WorkingFte struct {
    Value *string `xml:",chardata"`
}

type ExcludeFromHeadcount struct {
    Value *string `xml:",chardata"`
}

type PayRateTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type JobClassificationSummaryData struct {
    Additional *string `xml:"Additional,attr,omitempty"`
    JobClassificationReference *JobClassificationReference `xml:"Job_Classification_Reference,omitempty"`
    JobGroupReference *JobGroupReference `xml:"Job_Group_Reference,omitempty"`
}

type JobClassificationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type JobGroupReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CompanyInsiderReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WorkShiftReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WorkHoursProfilesReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type FederalWithholdingFein struct {
    Value *string `xml:",chardata"`
}

type WorkerCompensationCodeData struct {
    WorkersCompensationCodeReference *WorkersCompensationCodeReference `xml:"Workers_Compensation_Code_Reference,omitempty"`
    WorkersCompensationCode *WorkersCompensationCode `xml:"Workers_Compensation_Code,omitempty"`
}

type WorkersCompensationCodeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WorkersCompensationCode struct {
    Value *string `xml:",chardata"`
}

type PositionPayrollReportingCodeData struct {
    PayrollReportingCodeReference *PayrollReportingCodeReference `xml:"Payroll_Reporting_Code_Reference,omitempty"`
    Code *Code `xml:"Code,omitempty"`
    FormattedCode *FormattedCode `xml:"Formatted_Code,omitempty"`
    Name *Name `xml:"Name,omitempty"`
    PayrollReportingTypeReference *PayrollReportingTypeReference `xml:"Payroll_Reporting_Type_Reference,omitempty"`
}

type PayrollReportingCodeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Code struct {
    Value *string `xml:",chardata"`
}

type FormattedCode struct {
    Value *string `xml:",chardata"`
}

type Name struct {
    Value *string `xml:",chardata"`
}

type PayrollReportingTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type JobProfileSummaryData struct {
    JobProfileReference *JobProfileReference `xml:"Job_Profile_Reference,omitempty"`
    JobExempt *JobExempt `xml:"Job_Exempt,omitempty"`
    ManagementLevelReference *ManagementLevelReference `xml:"Management_Level_Reference,omitempty"`
    JobCategoryReference *JobCategoryReference `xml:"Job_Category_Reference,omitempty"`
    JobFamilyReference *JobFamilyReference `xml:"Job_Family_Reference,omitempty"`
    JobProfileName *JobProfileName `xml:"Job_Profile_Name,omitempty"`
    WorkShiftRequired *WorkShiftRequired `xml:"Work_Shift_Required,omitempty"`
    CriticalJob *CriticalJob `xml:"Critical_Job,omitempty"`
    DifficultyToFillReference *DifficultyToFillReference `xml:"Difficulty_to_Fill_Reference,omitempty"`
}

type JobProfileReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ManagementLevelReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type JobCategoryReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type JobFamilyReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type JobProfileName struct {
    Value *string `xml:",chardata"`
}

type WorkShiftRequired struct {
    Value *string `xml:",chardata"`
}

type CriticalJob struct {
    Value *string `xml:",chardata"`
}

type DifficultyToFillReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type BusinessSiteSummaryData struct {
    LocationReference *LocationReference `xml:"Location_Reference,omitempty"`
    Name *Name `xml:"Name,omitempty"`
    LocationTypeReference *LocationTypeReference `xml:"Location_Type_Reference,omitempty"`
    LocalReference *LocalReference `xml:"Local_Reference,omitempty"`
    DisplayLanguageReference *DisplayLanguageReference `xml:"Display_Language_Reference,omitempty"`
    TimeProfileReference *TimeProfileReference `xml:"Time_Profile_Reference,omitempty"`
    ScheduledWeeklyHours *ScheduledWeeklyHours `xml:"Scheduled_Weekly_Hours,omitempty"`
    AddressData *AddressData `xml:"Address_Data,omitempty"`
}

type LocationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type LocationTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type LocalReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DisplayLanguageReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type TimeProfileReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PayrollInterfaceProcessingData struct {
    EffectiveDate *EffectiveDate `xml:"Effective_Date,omitempty"`
    PayRateTypeReference *PayRateTypeReference `xml:"Pay_Rate_Type_Reference,omitempty"`
    FrequencyReference *FrequencyReference `xml:"Frequency_Reference,omitempty"`
    PayGroupReference *PayGroupReference `xml:"Pay_Group_Reference,omitempty"`
    PayrollEntityReference *PayrollEntityReference `xml:"Payroll_Entity_Reference,omitempty"`
    ExternalEmployeeTypeReference *ExternalEmployeeTypeReference `xml:"External_Employee_Type_Reference,omitempty"`
    PayrollFileNumber *PayrollFileNumber `xml:"Payroll_File_Number,omitempty"`
}

type EffectiveDate struct {
    Value *string `xml:",chardata"`
}

type FrequencyReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PayGroupReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PayrollEntityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ExternalEmployeeTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PayrollFileNumber struct {
    Value *string `xml:",chardata"`
}

type RegularPaidEquivalentHours struct {
    Value *string `xml:",chardata"`
}

type WorkerHoursProfileClassification struct {
    Value *string `xml:",chardata"`
}

type InternationalAssignmentData struct {
    InternationalAssignmentTypeReference *InternationalAssignmentTypeReference `xml:"International_Assignment_Type_Reference,omitempty"`
    StartInternationalAssignmentReasonReference *StartInternationalAssignmentReasonReference `xml:"Start_International_Assignment_Reason_Reference,omitempty"`
    ExpectedAssignmentEndDate *ExpectedAssignmentEndDate `xml:"Expected_Assignment_End_Date,omitempty"`
    HostCountryReference *HostCountryReference `xml:"Host_Country_Reference,omitempty"`
    HomeCountryReference *HomeCountryReference `xml:"Home_Country_Reference,omitempty"`
}

type InternationalAssignmentTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type StartInternationalAssignmentReasonReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ExpectedAssignmentEndDate struct {
    Value *string `xml:",chardata"`
}

type HostCountryReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type HomeCountryReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WorkSpaceReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type AcademicPaySetupData struct {
    AnnualWorkPeriodWorkPercentOfYear *AnnualWorkPeriodWorkPercentOfYear `xml:"Annual_Work_Period_Work_Percent_of_Year,omitempty"`
    AnnualWorkPeriodStartDate *AnnualWorkPeriodStartDate `xml:"Annual_Work_Period_Start_Date,omitempty"`
    AnnualWorkPeriodEndDate *AnnualWorkPeriodEndDate `xml:"Annual_Work_Period_End_Date,omitempty"`
    DisbursementPlanPeriodStartDate *DisbursementPlanPeriodStartDate `xml:"Disbursement_Plan_Period_Start_Date,omitempty"`
    DisbursementPlanPeriodEndDate *DisbursementPlanPeriodEndDate `xml:"Disbursement_Plan_Period_End_Date,omitempty"`
}

type AnnualWorkPeriodWorkPercentOfYear struct {
    Value *string `xml:",chardata"`
}

type AnnualWorkPeriodStartDate struct {
    Value *string `xml:",chardata"`
}

type AnnualWorkPeriodEndDate struct {
    Value *string `xml:",chardata"`
}

type DisbursementPlanPeriodStartDate struct {
    Value *string `xml:",chardata"`
}

type DisbursementPlanPeriodEndDate struct {
    Value *string `xml:",chardata"`
}

type EndDate struct {
    Value *string `xml:",chardata"`
}

type PayThroughDate struct {
    Value *string `xml:",chardata"`
}

type CollectiveAgreementSummaryData struct {
    CollectiveAgreementData *CollectiveAgreementData `xml:"Collective_Agreement_Data,omitempty"`
}

type CollectiveAgreementData struct {
    AssignEmployeeCollectiveAgreementEventReference *AssignEmployeeCollectiveAgreementEventReference `xml:"Assign_Employee_Collective_Agreement_Event_Reference,omitempty"`
    EffectiveDate *EffectiveDate `xml:"Effective_Date,omitempty"`
    EndAssignmentDate *EndAssignmentDate `xml:"End_Assignment_Date,omitempty"`
    CollectiveAgreementData *CollectiveAgreementData `xml:"Collective_Agreement_Data,omitempty"`
    CollectiveAgreementReference *CollectiveAgreementReference `xml:"Collective_Agreement_Reference,omitempty"`
    CollectiveAgreementReviewDate *CollectiveAgreementReviewDate `xml:"Collective_Agreement_Review_Date,omitempty"`
    CollectiveAgreementFactor *CollectiveAgreementFactor `xml:"Collective_Agreement_Factor,omitempty"`
}

type AssignEmployeeCollectiveAgreementEventReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type EndAssignmentDate struct {
    Value *string `xml:",chardata"`
}

type CollectiveAgreementReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CollectiveAgreementReviewDate struct {
    Value *string `xml:",chardata"`
}

type CollectiveAgreementFactor struct {
    Order *Order `xml:"Order,omitempty"`
    CollectiveAgreementFactorReference *CollectiveAgreementFactorReference `xml:"Collective_Agreement_Factor_Reference,omitempty"`
    CollectiveAgreementFactorOptionReference *CollectiveAgreementFactorOptionReference `xml:"Collective_Agreement_Factor_Option_Reference,omitempty"`
}

type Order struct {
    Value *string `xml:",chardata"`
}

type CollectiveAgreementFactorReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CollectiveAgreementFactorOptionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type EmployeeProbationPeriodSummaryData struct {
    EmployeeProbationPeriodDetailData *EmployeeProbationPeriodDetailData `xml:"Employee_Probation_Period_Detail_Data,omitempty"`
}

type EmployeeProbationPeriodDetailData struct {
    EmployeeProbationPeriodEventReference *EmployeeProbationPeriodEventReference `xml:"Employee_Probation_Period_Event_Reference,omitempty"`
    EmployeeProbationPeriodReference *EmployeeProbationPeriodReference `xml:"Employee_Probation_Period_Reference,omitempty"`
    EffectiveDate *EffectiveDate `xml:"Effective_Date,omitempty"`
    StartDate *StartDate `xml:"Start_Date,omitempty"`
    EndDate *EndDate `xml:"End_Date,omitempty"`
    ProbationTypeReference *ProbationTypeReference `xml:"Probation_Type_Reference,omitempty"`
    ProbationReasonReference *ProbationReasonReference `xml:"Probation_Reason_Reference,omitempty"`
    ExtendedEndDate *ExtendedEndDate `xml:"Extended_End_Date,omitempty"`
    Note *Note `xml:"Note,omitempty"`
}

type EmployeeProbationPeriodEventReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type EmployeeProbationPeriodReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ProbationTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ProbationReasonReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ExtendedEndDate struct {
    Value *string `xml:",chardata"`
}

type ManagerAsOfLastDetectedManagerChangeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PositionOrganizationsData struct {
    PositionOrganizationData *PositionOrganizationData `xml:"Position_Organization_Data,omitempty"`
    PayGroupAssignmentCorrectOrRescindData *PayGroupAssignmentCorrectOrRescindData `xml:"Pay_Group_Assignment_Correct_or_Rescind_Data,omitempty"`
}

type PositionOrganizationData struct {
    OrganizationReference *OrganizationReference `xml:"Organization_Reference,omitempty"`
    OrganizationData *OrganizationData `xml:"Organization_Data,omitempty"`
}

type OrganizationData struct {
    OrganizationReferenceId *OrganizationReferenceId `xml:"Organization_Reference_ID,omitempty"`
    OrganizationCode *OrganizationCode `xml:"Organization_Code,omitempty"`
    IntegrationIdData *IntegrationIdData `xml:"Integration_ID_Data,omitempty"`
    OrganizationName *OrganizationName `xml:"Organization_Name,omitempty"`
    OrganizationTypeReference *OrganizationTypeReference `xml:"Organization_Type_Reference,omitempty"`
    OrganizationSubtypeReference *OrganizationSubtypeReference `xml:"Organization_Subtype_Reference,omitempty"`
    PrimaryBusinessSiteReference *PrimaryBusinessSiteReference `xml:"Primary_Business_Site_Reference,omitempty"`
    OrganizationSupportRoleData *OrganizationSupportRoleData `xml:"Organization_Support_Role_Data,omitempty"`
    DateOfPayGroupAssignment *DateOfPayGroupAssignment `xml:"Date_of_Pay_Group_Assignment,omitempty"`
    UsedInChangeOrganizationAssignments *UsedInChangeOrganizationAssignments `xml:"Used_in_Change_Organization_Assignments,omitempty"`
}

type OrganizationReferenceId struct {
    Value *string `xml:",chardata"`
}

type OrganizationCode struct {
    Value *string `xml:",chardata"`
}

type IntegrationIdData struct {
    Id *Id `xml:"ID,omitempty"`
}

type OrganizationName struct {
    Value *string `xml:",chardata"`
}

type OrganizationTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type OrganizationSubtypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PrimaryBusinessSiteReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type OrganizationSupportRoleData struct {
    OrganizationSupportRole *OrganizationSupportRole `xml:"Organization_Support_Role,omitempty"`
}

type OrganizationSupportRole struct {
    OrganizationRoleReference *OrganizationRoleReference `xml:"Organization_Role_Reference,omitempty"`
    OrganizationRoleData *OrganizationRoleData `xml:"Organization_Role_Data,omitempty"`
}

type OrganizationRoleReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type OrganizationRoleData struct {
    WorkerReference *WorkerReference `xml:"Worker_Reference,omitempty"`
    AssignmentFrom *AssignmentFrom `xml:"Assignment_From,omitempty"`
}

type AssignmentFrom struct {
    Value *string `xml:",chardata"`
}

type DateOfPayGroupAssignment struct {
    Value *string `xml:",chardata"`
}

type UsedInChangeOrganizationAssignments struct {
    Value *string `xml:",chardata"`
}

type PayGroupAssignmentCorrectOrRescindData struct {
    EffectiveDate *EffectiveDate `xml:"Effective_Date,omitempty"`
    CompletedOn *CompletedOn `xml:"Completed_On,omitempty"`
    EventCorrected *EventCorrected `xml:"Event_Corrected,omitempty"`
    EventRescinded *EventRescinded `xml:"Event_Rescinded,omitempty"`
    OriginalPayGroupReference *OriginalPayGroupReference `xml:"Original_Pay_Group_Reference,omitempty"`
    ProposedPayGroupReference *ProposedPayGroupReference `xml:"Proposed_Pay_Group_Reference,omitempty"`
}

type CompletedOn struct {
    Value *string `xml:",chardata"`
}

type EventCorrected struct {
    Value *string `xml:",chardata"`
}

type EventRescinded struct {
    Value *string `xml:",chardata"`
}

type OriginalPayGroupReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ProposedPayGroupReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PositionManagementChainsData struct {
    PositionSupervisoryManagementChainData *PositionSupervisoryManagementChainData `xml:"Position_Supervisory_Management_Chain_Data,omitempty"`
    PositionMatrixManagementChainData *PositionMatrixManagementChainData `xml:"Position_Matrix_Management_Chain_Data,omitempty"`
}

type PositionSupervisoryManagementChainData struct {
    ManagementChainData *ManagementChainData `xml:"Management_Chain_Data,omitempty"`
}

type ManagementChainData struct {
    OrganizationReference *OrganizationReference `xml:"Organization_Reference,omitempty"`
    ManagerReference *ManagerReference `xml:"Manager_Reference,omitempty"`
    Manager *Manager `xml:"Manager,omitempty"`
}

type ManagerReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Manager struct {
    WorkerReference *WorkerReference `xml:"Worker_Reference,omitempty"`
    WorkerDescriptor *WorkerDescriptor `xml:"Worker_Descriptor,omitempty"`
}

type PositionMatrixManagementChainData struct {
    ManagementChainData *ManagementChainData `xml:"Management_Chain_Data,omitempty"`
}

type PositionContractDetailData struct {
    ContractPayRate *ContractPayRate `xml:"Contract_Pay_Rate,omitempty"`
    CurrencyReference *CurrencyReference `xml:"Currency_Reference,omitempty"`
    FrequencyReference *FrequencyReference `xml:"Frequency_Reference,omitempty"`
    ContractAssignmentDetails *ContractAssignmentDetails `xml:"Contract_Assignment_Details,omitempty"`
    ContractEndDate *ContractEndDate `xml:"Contract_End_Date,omitempty"`
    SupplierReference *SupplierReference `xml:"Supplier_Reference,omitempty"`
}

type ContractPayRate struct {
    Value *string `xml:",chardata"`
}

type CurrencyReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ContractAssignmentDetails struct {
    Value *string `xml:",chardata"`
}

type ContractEndDate struct {
    Value *string `xml:",chardata"`
}

type SupplierReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WorkerStatusData struct {
    Active *Active `xml:"Active,omitempty"`
    ActiveStatusDate *ActiveStatusDate `xml:"Active_Status_Date,omitempty"`
    HireDate *HireDate `xml:"Hire_Date,omitempty"`
    OriginalHireDate *OriginalHireDate `xml:"Original_Hire_Date,omitempty"`
    HireReasonReference *HireReasonReference `xml:"Hire_Reason_Reference,omitempty"`
    EndEmploymentDate *EndEmploymentDate `xml:"End_Employment_Date,omitempty"`
    ContinuousServiceDate *ContinuousServiceDate `xml:"Continuous_Service_Date,omitempty"`
    FirstDayOfWork *FirstDayOfWork `xml:"First_Day_of_Work,omitempty"`
    ExpectedRetirementDate *ExpectedRetirementDate `xml:"Expected_Retirement_Date,omitempty"`
    RetirementEligibilityDate *RetirementEligibilityDate `xml:"Retirement_Eligibility_Date,omitempty"`
    Retired *Retired `xml:"Retired,omitempty"`
    RetirementDate *RetirementDate `xml:"Retirement_Date,omitempty"`
    SeniorityDate *SeniorityDate `xml:"Seniority_Date,omitempty"`
    SeveranceDate *SeveranceDate `xml:"Severance_Date,omitempty"`
    BenefitsServiceDate *BenefitsServiceDate `xml:"Benefits_Service_Date,omitempty"`
    CompanyServiceDate *CompanyServiceDate `xml:"Company_Service_Date,omitempty"`
    TimeOffServiceDate *TimeOffServiceDate `xml:"Time_Off_Service_Date,omitempty"`
    VestingDate *VestingDate `xml:"Vesting_Date,omitempty"`
    DateEnteredWorkforce *DateEnteredWorkforce `xml:"Date_Entered_Workforce,omitempty"`
    DaysUnemployed *DaysUnemployed `xml:"Days_Unemployed,omitempty"`
    MonthsContinuousPriorEmployment *MonthsContinuousPriorEmployment `xml:"Months_Continuous_Prior_Employment,omitempty"`
    Terminated *Terminated `xml:"Terminated,omitempty"`
    TerminationDate *TerminationDate `xml:"Termination_Date,omitempty"`
    PayThroughDate *PayThroughDate `xml:"Pay_Through_Date,omitempty"`
    AgreementSignatureDate *AgreementSignatureDate `xml:"Agreement_Signature_Date,omitempty"`
    DismissalProcessStartDate *DismissalProcessStartDate `xml:"Dismissal_Process_Start_Date,omitempty"`
    NoticePeriodStartDate *NoticePeriodStartDate `xml:"Notice_Period_Start_Date,omitempty"`
    PrimaryTerminationReasonReference *PrimaryTerminationReasonReference `xml:"Primary_Termination_Reason_Reference,omitempty"`
    PrimaryTerminationCategoryReference *PrimaryTerminationCategoryReference `xml:"Primary_Termination_Category_Reference,omitempty"`
    TerminationInvoluntary *TerminationInvoluntary `xml:"Termination_Involuntary,omitempty"`
    SecondaryTerminationReasonsData *SecondaryTerminationReasonsData `xml:"Secondary_Termination_Reasons_Data,omitempty"`
    LocalTerminationReasonReference *LocalTerminationReasonReference `xml:"Local_Termination_Reason_Reference,omitempty"`
    EligibleForHireReference *EligibleForHireReference `xml:"Eligible_for_Hire_Reference,omitempty"`
    RegrettableTermination *RegrettableTermination `xml:"Regrettable_Termination,omitempty"`
    EligibleForRehireOnLatestTerminationReference *EligibleForRehireOnLatestTerminationReference `xml:"Eligible_for_Rehire_on_Latest_Termination_Reference,omitempty"`
    HireRescinded *HireRescinded `xml:"Hire_Rescinded,omitempty"`
    TerminationLastDayOfWork *TerminationLastDayOfWork `xml:"Termination_Last_Day_of_Work,omitempty"`
    ResignationDate *ResignationDate `xml:"Resignation_Date,omitempty"`
    LastDateForWhichPaid *LastDateForWhichPaid `xml:"Last_Date_for_Which_Paid,omitempty"`
    ExpectedDateOfReturn *ExpectedDateOfReturn `xml:"Expected_Date_of_Return,omitempty"`
    NotReturning *NotReturning `xml:"Not_Returning,omitempty"`
    ReturnUnknown *ReturnUnknown `xml:"Return_Unknown,omitempty"`
    ProbationStartDate *ProbationStartDate `xml:"Probation_Start_Date,omitempty"`
    ProbationEndDate *ProbationEndDate `xml:"Probation_End_Date,omitempty"`
    LeaveStatusData *LeaveStatusData `xml:"Leave_Status_Data,omitempty"`
}

type Active struct {
    Value *string `xml:",chardata"`
}

type ActiveStatusDate struct {
    Value *string `xml:",chardata"`
}

type HireDate struct {
    Value *string `xml:",chardata"`
}

type OriginalHireDate struct {
    Value *string `xml:",chardata"`
}

type HireReasonReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ContinuousServiceDate struct {
    Value *string `xml:",chardata"`
}

type FirstDayOfWork struct {
    Value *string `xml:",chardata"`
}

type ExpectedRetirementDate struct {
    Value *string `xml:",chardata"`
}

type RetirementEligibilityDate struct {
    Value *string `xml:",chardata"`
}

type Retired struct {
    Value *string `xml:",chardata"`
}

type RetirementDate struct {
    Value *string `xml:",chardata"`
}

type SeniorityDate struct {
    Value *string `xml:",chardata"`
}

type SeveranceDate struct {
    Value *string `xml:",chardata"`
}

type BenefitsServiceDate struct {
    Value *string `xml:",chardata"`
}

type CompanyServiceDate struct {
    Value *string `xml:",chardata"`
}

type TimeOffServiceDate struct {
    Value *string `xml:",chardata"`
}

type VestingDate struct {
    Value *string `xml:",chardata"`
}

type DateEnteredWorkforce struct {
    Value *string `xml:",chardata"`
}

type DaysUnemployed struct {
    Value *string `xml:",chardata"`
}

type MonthsContinuousPriorEmployment struct {
    Value *string `xml:",chardata"`
}

type Terminated struct {
    Value *string `xml:",chardata"`
}

type TerminationDate struct {
    Value *string `xml:",chardata"`
}

type AgreementSignatureDate struct {
    Value *string `xml:",chardata"`
}

type DismissalProcessStartDate struct {
    Value *string `xml:",chardata"`
}

type NoticePeriodStartDate struct {
    Value *string `xml:",chardata"`
}

type PrimaryTerminationReasonReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PrimaryTerminationCategoryReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type TerminationInvoluntary struct {
    Value *string `xml:",chardata"`
}

type SecondaryTerminationReasonsData struct {
    SecondaryTerminationReasonReference *SecondaryTerminationReasonReference `xml:"Secondary_Termination_Reason_Reference,omitempty"`
    SecondaryTerminationReasonCategoryReference *SecondaryTerminationReasonCategoryReference `xml:"Secondary_Termination_Reason_Category_Reference,omitempty"`
}

type SecondaryTerminationReasonReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SecondaryTerminationReasonCategoryReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type LocalTerminationReasonReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type EligibleForHireReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type RegrettableTermination struct {
    Value *string `xml:",chardata"`
}

type EligibleForRehireOnLatestTerminationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type HireRescinded struct {
    Value *string `xml:",chardata"`
}

type TerminationLastDayOfWork struct {
    Value *string `xml:",chardata"`
}

type ResignationDate struct {
    Value *string `xml:",chardata"`
}

type LastDateForWhichPaid struct {
    Value *string `xml:",chardata"`
}

type ExpectedDateOfReturn struct {
    Value *string `xml:",chardata"`
}

type NotReturning struct {
    Value *string `xml:",chardata"`
}

type ReturnUnknown struct {
    Value *string `xml:",chardata"`
}

type ProbationStartDate struct {
    Value *string `xml:",chardata"`
}

type ProbationEndDate struct {
    Value *string `xml:",chardata"`
}

type LeaveStatusData struct {
    LeaveRequestEventReference *LeaveRequestEventReference `xml:"Leave_Request_Event_Reference,omitempty"`
    LeaveRequestDescription *LeaveRequestDescription `xml:"Leave_Request_Description,omitempty"`
    LeaveReturnEventReference *LeaveReturnEventReference `xml:"Leave_Return_Event_Reference,omitempty"`
    OnLeave *OnLeave `xml:"On_Leave,omitempty"`
    LeaveStartDate *LeaveStartDate `xml:"Leave_Start_Date,omitempty"`
    EstimatedLeaveEndDate *EstimatedLeaveEndDate `xml:"Estimated_Leave_End_Date,omitempty"`
    LeaveEndDate *LeaveEndDate `xml:"Leave_End_Date,omitempty"`
    FirstDayOfWork *FirstDayOfWork `xml:"First_Day_Of_Work,omitempty"`
    LeaveLastDayOfWork *LeaveLastDayOfWork `xml:"Leave_Last_Day_of_Work,omitempty"`
    LeaveOfAbsenceTypeReference *LeaveOfAbsenceTypeReference `xml:"Leave_of_Absence_Type_Reference,omitempty"`
    BenefitsEffect *BenefitsEffect `xml:"Benefits_Effect,omitempty"`
    PayrollEffect *PayrollEffect `xml:"Payroll_Effect,omitempty"`
    PaidTimeOffAccrualEffect *PaidTimeOffAccrualEffect `xml:"Paid_Time_Off_Accrual_Effect,omitempty"`
    ContinuousServiceAccrualEffect *ContinuousServiceAccrualEffect `xml:"Continuous_Service_Accrual_Effect,omitempty"`
    StockVestingEffect *StockVestingEffect `xml:"Stock_Vesting_Effect,omitempty"`
    LeaveTypeReasonReference *LeaveTypeReasonReference `xml:"Leave_Type_Reason_Reference,omitempty"`
    LeaveRequestAdditionalFields *LeaveRequestAdditionalFields `xml:"Leave_Request_Additional_Fields,omitempty"`
}

type LeaveRequestEventReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type LeaveRequestDescription struct {
    Value *string `xml:",chardata"`
}

type LeaveReturnEventReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type OnLeave struct {
    Value *string `xml:",chardata"`
}

type LeaveStartDate struct {
    Value *string `xml:",chardata"`
}

type EstimatedLeaveEndDate struct {
    Value *string `xml:",chardata"`
}

type LeaveEndDate struct {
    Value *string `xml:",chardata"`
}

type LeaveLastDayOfWork struct {
    Value *string `xml:",chardata"`
}

type LeaveOfAbsenceTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type BenefitsEffect struct {
    Value *string `xml:",chardata"`
}

type PayrollEffect struct {
    Value *string `xml:",chardata"`
}

type PaidTimeOffAccrualEffect struct {
    Value *string `xml:",chardata"`
}

type ContinuousServiceAccrualEffect struct {
    Value *string `xml:",chardata"`
}

type StockVestingEffect struct {
    Value *string `xml:",chardata"`
}

type LeaveTypeReasonReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type LeaveRequestAdditionalFields struct {
    LastDateForWhichPaid *LastDateForWhichPaid `xml:"Last_Date_for_Which_Paid,omitempty"`
    ExpectedDueDate *ExpectedDueDate `xml:"Expected_Due_Date,omitempty"`
    ChildSBirthDate *ChildSBirthDate `xml:"Child_s_Birth_Date,omitempty"`
    StillbirthBabyDeceased *StillbirthBabyDeceased `xml:"Stillbirth_Baby_Deceased,omitempty"`
    DateBabyArrivedHomeFromHospital *DateBabyArrivedHomeFromHospital `xml:"Date_Baby_Arrived_Home_From_Hospital,omitempty"`
}

type ExpectedDueDate struct {
    Value *string `xml:",chardata"`
}

type ChildSBirthDate struct {
    Value *string `xml:",chardata"`
}

type StillbirthBabyDeceased struct {
    Value *string `xml:",chardata"`
}

type DateBabyArrivedHomeFromHospital struct {
    Value *string `xml:",chardata"`
}

