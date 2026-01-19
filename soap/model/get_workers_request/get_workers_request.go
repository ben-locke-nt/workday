// Generated from ../model/get_workers_request/get_workers_request.go
package get_workers_request

import "encoding/xml"

type GetWorkersRequest struct {
    XMLName xml.Name `xml:"wd:Get_Workers_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    RequestReferences *RequestReferences `xml:"wd:Request_References,omitempty"`
    RequestCriteria *RequestCriteria `xml:"wd:Request_Criteria,omitempty"`
    ResponseFilter *ResponseFilter `xml:"wd:Response_Filter,omitempty"`
    ResponseGroup *ResponseGroup `xml:"wd:Response_Group,omitempty"`
}

type RequestReferences struct {
    SkipNonExistingInstances *string `xml:"wd:Skip_Non_Existing_Instances,attr,omitempty"`
    IgnoreInvalidReferences *string `xml:"wd:Ignore_Invalid_References,attr,omitempty"`
    WorkerReference *WorkerReference `xml:"wd:Worker_Reference,omitempty"`
}

type WorkerReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Id struct {
    Value *string `xml:",chardata"`
    Type *string `xml:"wd:type,attr,omitempty"`
    ParentId *string `xml:"wd:parent_id,attr,omitempty"`
    ParentType *string `xml:"wd:parent_type,attr,omitempty"`
}

type RequestCriteria struct {
    TransactionLogCriteriaData *TransactionLogCriteriaData `xml:"wd:Transaction_Log_Criteria_Data,omitempty"`
    OrganizationReference *OrganizationReference `xml:"wd:Organization_Reference,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    IncludeSubordinateOrganizations *IncludeSubordinateOrganizations `xml:"wd:Include_Subordinate_Organizations,omitempty"`
    PositionReference *PositionReference `xml:"wd:Position_Reference,omitempty"`
    EventReference *EventReference `xml:"wd:Event_Reference,omitempty"`
    BenefitPlanReference *BenefitPlanReference `xml:"wd:Benefit_Plan_Reference,omitempty"`
    FieldAndParameterCriteriaData *FieldAndParameterCriteriaData `xml:"wd:Field_And_Parameter_Criteria_Data,omitempty"`
    NationalIdCriteriaData *NationalIdCriteriaData `xml:"wd:National_ID_Criteria_Data,omitempty"`
    ExcludeInactiveWorkers *ExcludeInactiveWorkers `xml:"wd:Exclude_Inactive_Workers,omitempty"`
    ExcludeEmployees *ExcludeEmployees `xml:"wd:Exclude_Employees,omitempty"`
    ExcludeContingentWorkers *ExcludeContingentWorkers `xml:"wd:Exclude_Contingent_Workers,omitempty"`
    EligibilityCriteriaData *EligibilityCriteriaData `xml:"wd:Eligibility_Criteria_Data,omitempty"`
    UniversalIdReference *UniversalIdReference `xml:"wd:Universal_ID_Reference,omitempty"`
}

type TransactionLogCriteriaData struct {
    TransactionDateRangeData *TransactionDateRangeData `xml:"wd:Transaction_Date_Range_Data,omitempty"`
    TransactionTypeReferences *TransactionTypeReferences `xml:"wd:Transaction_Type_References,omitempty"`
    SubscriberReference *SubscriberReference `xml:"wd:Subscriber_Reference,omitempty"`
}

type TransactionDateRangeData struct {
    UpdatedFrom *UpdatedFrom `xml:"wd:Updated_From,omitempty"`
    UpdatedThrough *UpdatedThrough `xml:"wd:Updated_Through,omitempty"`
    EffectiveFrom *EffectiveFrom `xml:"wd:Effective_From,omitempty"`
    EffectiveThrough *EffectiveThrough `xml:"wd:Effective_Through,omitempty"`
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
    TransactionTypeReference *TransactionTypeReference `xml:"wd:Transaction_Type_Reference,omitempty"`
}

type TransactionTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SubscriberReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type OrganizationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CountryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type IncludeSubordinateOrganizations struct {
    Value *string `xml:",chardata"`
}

type PositionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type EventReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type BenefitPlanReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type FieldAndParameterCriteriaData struct {
    ProviderReference *ProviderReference `xml:"wd:Provider_Reference,omitempty"`
}

type ProviderReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type NationalIdCriteriaData struct {
    IdentifierId *IdentifierId `xml:"wd:Identifier_ID,omitempty"`
    NationalIdTypeReference *NationalIdTypeReference `xml:"wd:National_ID_Type_Reference,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
}

type IdentifierId struct {
    Value *string `xml:",chardata"`
}

type NationalIdTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    FieldReference *FieldReference `xml:"wd:Field_Reference,omitempty"`
    AsOfEffectiveDate *AsOfEffectiveDate `xml:"wd:As_Of_Effective_Date,omitempty"`
    AsOfEntryDatetime *AsOfEntryDatetime `xml:"wd:As_Of_Entry_DateTime,omitempty"`
}

type FieldReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AsOfEffectiveDate struct {
    Value *string `xml:",chardata"`
}

type AsOfEntryDatetime struct {
    Value *string `xml:",chardata"`
}

type UniversalIdReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ResponseFilter struct {
    AsOfEffectiveDate *AsOfEffectiveDate `xml:"wd:As_Of_Effective_Date,omitempty"`
    AsOfEntryDatetime *AsOfEntryDatetime `xml:"wd:As_Of_Entry_DateTime,omitempty"`
    Page *Page `xml:"wd:Page,omitempty"`
    Count *Count `xml:"wd:Count,omitempty"`
}

type Page struct {
    Value *string `xml:",chardata"`
}

type Count struct {
    Value *string `xml:",chardata"`
}

type ResponseGroup struct {
    IncludeReference *IncludeReference `xml:"wd:Include_Reference,omitempty"`
    IncludePersonalInformation *IncludePersonalInformation `xml:"wd:Include_Personal_Information,omitempty"`
    ShowAllPersonalInformation *ShowAllPersonalInformation `xml:"wd:Show_All_Personal_Information,omitempty"`
    IncludeAdditionalJobs *IncludeAdditionalJobs `xml:"wd:Include_Additional_Jobs,omitempty"`
    IncludeEmploymentInformation *IncludeEmploymentInformation `xml:"wd:Include_Employment_Information,omitempty"`
    IncludeCompensation *IncludeCompensation `xml:"wd:Include_Compensation,omitempty"`
    IncludeOrganizations *IncludeOrganizations `xml:"wd:Include_Organizations,omitempty"`
    ExcludeOrganizationSupportRoleData *ExcludeOrganizationSupportRoleData `xml:"wd:Exclude_Organization_Support_Role_Data,omitempty"`
    ExcludeLocationHierarchies *ExcludeLocationHierarchies `xml:"wd:Exclude_Location_Hierarchies,omitempty"`
    ExcludeCostCenters *ExcludeCostCenters `xml:"wd:Exclude_Cost_Centers,omitempty"`
    ExcludeCostCenterHierarchies *ExcludeCostCenterHierarchies `xml:"wd:Exclude_Cost_Center_Hierarchies,omitempty"`
    ExcludeCompanies *ExcludeCompanies `xml:"wd:Exclude_Companies,omitempty"`
    ExcludeCompanyHierarchies *ExcludeCompanyHierarchies `xml:"wd:Exclude_Company_Hierarchies,omitempty"`
    ExcludeMatrixOrganizations *ExcludeMatrixOrganizations `xml:"wd:Exclude_Matrix_Organizations,omitempty"`
    ExcludePayGroups *ExcludePayGroups `xml:"wd:Exclude_Pay_Groups,omitempty"`
    ExcludeRegions *ExcludeRegions `xml:"wd:Exclude_Regions,omitempty"`
    ExcludeRegionHierarchies *ExcludeRegionHierarchies `xml:"wd:Exclude_Region_Hierarchies,omitempty"`
    ExcludeSupervisoryOrganizations *ExcludeSupervisoryOrganizations `xml:"wd:Exclude_Supervisory_Organizations,omitempty"`
    ExcludeTeams *ExcludeTeams `xml:"wd:Exclude_Teams,omitempty"`
    ExcludeCustomOrganizations *ExcludeCustomOrganizations `xml:"wd:Exclude_Custom_Organizations,omitempty"`
    IncludeRoles *IncludeRoles `xml:"wd:Include_Roles,omitempty"`
    IncludeManagementChainData *IncludeManagementChainData `xml:"wd:Include_Management_Chain_Data,omitempty"`
    IncludeMultipleManagersInManagementChainData *IncludeMultipleManagersInManagementChainData `xml:"wd:Include_Multiple_Managers_in_Management_Chain_Data,omitempty"`
    IncludeBenefitEnrollments *IncludeBenefitEnrollments `xml:"wd:Include_Benefit_Enrollments,omitempty"`
    IncludeBenefitEligibility *IncludeBenefitEligibility `xml:"wd:Include_Benefit_Eligibility,omitempty"`
    IncludeRelatedPersons *IncludeRelatedPersons `xml:"wd:Include_Related_Persons,omitempty"`
    IncludeQualifications *IncludeQualifications `xml:"wd:Include_Qualifications,omitempty"`
    IncludeEmployeeReview *IncludeEmployeeReview `xml:"wd:Include_Employee_Review,omitempty"`
    IncludeGoals *IncludeGoals `xml:"wd:Include_Goals,omitempty"`
    IncludeDevelopmentItems *IncludeDevelopmentItems `xml:"wd:Include_Development_Items,omitempty"`
    IncludeSkills *IncludeSkills `xml:"wd:Include_Skills,omitempty"`
    IncludePhoto *IncludePhoto `xml:"wd:Include_Photo,omitempty"`
    IncludeWorkerDocuments *IncludeWorkerDocuments `xml:"wd:Include_Worker_Documents,omitempty"`
    IncludeTransactionLogData *IncludeTransactionLogData `xml:"wd:Include_Transaction_Log_Data,omitempty"`
    IncludeSubeventsForCorrectedTransaction *IncludeSubeventsForCorrectedTransaction `xml:"wd:Include_Subevents_for_Corrected_Transaction,omitempty"`
    IncludeSubeventsForRescindedTransaction *IncludeSubeventsForRescindedTransaction `xml:"wd:Include_Subevents_for_Rescinded_Transaction,omitempty"`
    IncludeSuccessionProfile *IncludeSuccessionProfile `xml:"wd:Include_Succession_Profile,omitempty"`
    IncludeTalentAssessment *IncludeTalentAssessment `xml:"wd:Include_Talent_Assessment,omitempty"`
    IncludeEmployeeContractData *IncludeEmployeeContractData `xml:"wd:Include_Employee_Contract_Data,omitempty"`
    IncludeContractsForTerminatedWorkers *IncludeContractsForTerminatedWorkers `xml:"wd:Include_Contracts_for_Terminated_Workers,omitempty"`
    IncludeCollectiveAgreementData *IncludeCollectiveAgreementData `xml:"wd:Include_Collective_Agreement_Data,omitempty"`
    IncludeProbationPeriodData *IncludeProbationPeriodData `xml:"wd:Include_Probation_Period_Data,omitempty"`
    IncludeExtendedEmployeeContractDetails *IncludeExtendedEmployeeContractDetails `xml:"wd:Include_Extended_Employee_Contract_Details,omitempty"`
    IncludeFeedbackReceived *IncludeFeedbackReceived `xml:"wd:Include_Feedback_Received,omitempty"`
    IncludeUserAccount *IncludeUserAccount `xml:"wd:Include_User_Account,omitempty"`
    IncludeCareer *IncludeCareer `xml:"wd:Include_Career,omitempty"`
    IncludeAccountProvisioning *IncludeAccountProvisioning `xml:"wd:Include_Account_Provisioning,omitempty"`
    IncludeBackgroundCheckData *IncludeBackgroundCheckData `xml:"wd:Include_Background_Check_Data,omitempty"`
    IncludeContingentWorkerTaxAuthorityFormInformation *IncludeContingentWorkerTaxAuthorityFormInformation `xml:"wd:Include_Contingent_Worker_Tax_Authority_Form_Information,omitempty"`
    ExcludeFunds *ExcludeFunds `xml:"wd:Exclude_Funds,omitempty"`
    ExcludeFundHierarchies *ExcludeFundHierarchies `xml:"wd:Exclude_Fund_Hierarchies,omitempty"`
    ExcludeGrants *ExcludeGrants `xml:"wd:Exclude_Grants,omitempty"`
    ExcludeGrantHierarchies *ExcludeGrantHierarchies `xml:"wd:Exclude_Grant_Hierarchies,omitempty"`
    ExcludeBusinessUnits *ExcludeBusinessUnits `xml:"wd:Exclude_Business_Units,omitempty"`
    ExcludeBusinessUnitHierarchies *ExcludeBusinessUnitHierarchies `xml:"wd:Exclude_Business_Unit_Hierarchies,omitempty"`
    ExcludePrograms *ExcludePrograms `xml:"wd:Exclude_Programs,omitempty"`
    ExcludeProgramHierarchies *ExcludeProgramHierarchies `xml:"wd:Exclude_Program_Hierarchies,omitempty"`
    ExcludeGifts *ExcludeGifts `xml:"wd:Exclude_Gifts,omitempty"`
    ExcludeGiftHierarchies *ExcludeGiftHierarchies `xml:"wd:Exclude_Gift_Hierarchies,omitempty"`
    ExcludeRetireeOrganizations *ExcludeRetireeOrganizations `xml:"wd:Exclude_Retiree_Organizations,omitempty"`
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

