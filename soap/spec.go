// Generated from Get_Candidates_Request.xml
package soap

import "encoding/xml"

type GetWorkersResponse struct {
    XMLName xml.Name `xml:"wd:Get_Workers_Response"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    RequestReferences *RequestReferences `xml:"wd:Request_References,omitempty"`
    RequestCriteria *RequestCriteria `xml:"wd:Request_Criteria,omitempty"`
    ResponseFilter *ResponseFilter `xml:"wd:Response_Filter,omitempty"`
    ResponseGroup *ResponseGroup `xml:"wd:Response_Group,omitempty"`
    ResponseResults *ResponseResults `xml:"wd:Response_Results,omitempty"`
    ResponseData *ResponseData `xml:"wd:Response_Data,omitempty"`
}

type RequestReferences struct {
    SkipNonExistingInstances *string `xml:"wd:Skip_Non_Existing_Instances,attr,omitempty"`
    IgnoreInvalidReferences *string `xml:"wd:Ignore_Invalid_References,attr,omitempty"`
    WorkerReference *WorkerReference `xml:"wd:Worker_Reference,omitempty"`
    ApplicantReference *ApplicantReference `xml:"wd:Applicant_Reference,omitempty"`
    CandidateReference *CandidateReference `xml:"wd:Candidate_Reference,omitempty"`
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
    SystemId *string `xml:"wd:System_ID,attr,omitempty"`
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
    WorkerReference *WorkerReference `xml:"wd:Worker_Reference,omitempty"`
    FormerWorkerReference *FormerWorkerReference `xml:"wd:Former_Worker_Reference,omitempty"`
    EmailAddress *EmailAddress `xml:"wd:Email_Address,omitempty"`
    FirstName *FirstName `xml:"wd:First_Name,omitempty"`
    LastName *LastName `xml:"wd:Last_Name,omitempty"`
    CandidateEmailAddress *CandidateEmailAddress `xml:"wd:Candidate_Email_Address,omitempty"`
    PreHireReference *PreHireReference `xml:"wd:Pre_Hire_Reference,omitempty"`
    JobRequisitionReference *JobRequisitionReference `xml:"wd:Job_Requisition_Reference,omitempty"`
    RecruitingStageReference *RecruitingStageReference `xml:"wd:Recruiting_Stage_Reference,omitempty"`
    ApplicantSourceReference *ApplicantSourceReference `xml:"wd:Applicant_Source_Reference,omitempty"`
    CandidateTagReference *CandidateTagReference `xml:"wd:Candidate_Tag_Reference,omitempty"`
    AppliedFrom *AppliedFrom `xml:"wd:Applied_From,omitempty"`
    AppliedThrough *AppliedThrough `xml:"wd:Applied_Through,omitempty"`
    InternalWorkersOnly *InternalWorkersOnly `xml:"wd:Internal_Workers_Only,omitempty"`
    CreatedFrom *CreatedFrom `xml:"wd:Created_From,omitempty"`
    CreatedThrough *CreatedThrough `xml:"wd:Created_Through,omitempty"`
    CriteriaMatchScoreCategoryReference *CriteriaMatchScoreCategoryReference `xml:"wd:Criteria_Match_Score_Category_Reference,omitempty"`
    StaticCandidatePoolReference *StaticCandidatePoolReference `xml:"wd:Static_Candidate_Pool_Reference,omitempty"`
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
    IncludeRecruitingInformation *IncludeRecruitingInformation `xml:"wd:Include_Recruiting_Information,omitempty"`
    IncludeQualificationProfile *IncludeQualificationProfile `xml:"wd:Include_Qualification_Profile,omitempty"`
    IncludeResume *IncludeResume `xml:"wd:Include_Resume,omitempty"`
    IncludeBackgroundCheck *IncludeBackgroundCheck `xml:"wd:Include_Background_Check,omitempty"`
    IncludeExternalIntegrationIdData *IncludeExternalIntegrationIdData `xml:"wd:Include_External_Integration_ID_Data,omitempty"`
    ExcludeAllAttachments *ExcludeAllAttachments `xml:"wd:Exclude_All_Attachments,omitempty"`
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
    TotalResults *TotalResults `xml:"wd:Total_Results,omitempty"`
    TotalPages *TotalPages `xml:"wd:Total_Pages,omitempty"`
    PageResults *PageResults `xml:"wd:Page_Results,omitempty"`
    Page *Page `xml:"wd:Page,omitempty"`
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
    Worker *Worker `xml:"wd:Worker,omitempty"`
    Candidate *Candidate `xml:"wd:Candidate,omitempty"`
    QuestionnaireReference *QuestionnaireReference `xml:"wd:Questionnaire_Reference,omitempty"`
    QuestionnaireAnswerData *QuestionnaireAnswerData `xml:"wd:Questionnaire_Answer_Data,omitempty"`
    Applicant *Applicant `xml:"wd:Applicant,omitempty"`
}

type Worker struct {
    WorkerReference *WorkerReference `xml:"wd:Worker_Reference,omitempty"`
    WorkerDescriptor *WorkerDescriptor `xml:"wd:Worker_Descriptor,omitempty"`
    UniversalIdentifierReference *UniversalIdentifierReference `xml:"wd:Universal_Identifier_Reference,omitempty"`
    WorkerData *WorkerData `xml:"wd:Worker_Data,omitempty"`
}

type WorkerDescriptor struct {
    Value *string `xml:",chardata"`
}

type UniversalIdentifierReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WorkerData struct {
    WorkerId *WorkerId `xml:"wd:Worker_ID,omitempty"`
    UserId *UserId `xml:"wd:User_ID,omitempty"`
    UniversalId *UniversalId `xml:"wd:Universal_ID,omitempty"`
    PersonalData *PersonalData `xml:"wd:Personal_Data,omitempty"`
    EmploymentData *EmploymentData `xml:"wd:Employment_Data,omitempty"`
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
    UniversalId *UniversalId `xml:"wd:Universal_ID,omitempty"`
    NameData *NameData `xml:"wd:Name_Data,omitempty"`
    PersonalInformationData *PersonalInformationData `xml:"wd:Personal_Information_Data,omitempty"`
    IdentificationData *IdentificationData `xml:"wd:Identification_Data,omitempty"`
    ContactData *ContactData `xml:"wd:Contact_Data,omitempty"`
    TobaccoUse *TobaccoUse `xml:"wd:Tobacco_Use,omitempty"`
}

type NameData struct {
    LegalNameData *LegalNameData `xml:"wd:Legal_Name_Data,omitempty"`
    PreferredNameData *PreferredNameData `xml:"wd:Preferred_Name_Data,omitempty"`
    AdditionalNameData *AdditionalNameData `xml:"wd:Additional_Name_Data,omitempty"`
    LegalName *LegalName `xml:"wd:Legal_Name,omitempty"`
    PreferredName *PreferredName `xml:"wd:Preferred_Name,omitempty"`
    FormattedName *string `xml:"wd:Formatted_Name,attr,omitempty"`
    ReportingName *string `xml:"wd:Reporting_Name,attr,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    PrefixData *PrefixData `xml:"wd:Prefix_Data,omitempty"`
    FirstName *FirstName `xml:"wd:First_Name,omitempty"`
    MiddleName *MiddleName `xml:"wd:Middle_Name,omitempty"`
    LastName *LastName `xml:"wd:Last_Name,omitempty"`
    SecondaryLastName *SecondaryLastName `xml:"wd:Secondary_Last_Name,omitempty"`
    TertiaryLastName *TertiaryLastName `xml:"wd:Tertiary_Last_Name,omitempty"`
    LocalNameDetailData *LocalNameDetailData `xml:"wd:Local_Name_Detail_Data,omitempty"`
    SuffixData *SuffixData `xml:"wd:Suffix_Data,omitempty"`
    FullNameForSingaporeAndMalaysia *FullNameForSingaporeAndMalaysia `xml:"wd:Full_Name_for_Singapore_and_Malaysia,omitempty"`
}

type LegalNameData struct {
    NameDetailData *NameDetailData `xml:"wd:Name_Detail_Data,omitempty"`
}

type NameDetailData struct {
    FormattedName *string `xml:"wd:Formatted_Name,attr,omitempty"`
    ReportingName *string `xml:"wd:Reporting_Name,attr,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    PrefixData *PrefixData `xml:"wd:Prefix_Data,omitempty"`
    FirstName *FirstName `xml:"wd:First_Name,omitempty"`
    MiddleName *MiddleName `xml:"wd:Middle_Name,omitempty"`
    LastName *LastName `xml:"wd:Last_Name,omitempty"`
    SecondaryLastName *SecondaryLastName `xml:"wd:Secondary_Last_Name,omitempty"`
    TertiaryLastName *TertiaryLastName `xml:"wd:Tertiary_Last_Name,omitempty"`
    LocalNameDetailData *LocalNameDetailData `xml:"wd:Local_Name_Detail_Data,omitempty"`
    SuffixData *SuffixData `xml:"wd:Suffix_Data,omitempty"`
    FullNameForSingaporeAndMalaysia *FullNameForSingaporeAndMalaysia `xml:"wd:Full_Name_for_Singapore_and_Malaysia,omitempty"`
    TitleReference *TitleReference `xml:"wd:Title_Reference,omitempty"`
    SalutationSuffixReference *SalutationSuffixReference `xml:"wd:Salutation_Suffix_Reference,omitempty"`
    FullName *FullName `xml:"wd:Full_Name,omitempty"`
    HereditarySuffixReference *HereditarySuffixReference `xml:"wd:Hereditary_Suffix_Reference,omitempty"`
    SocialSuffixReference *SocialSuffixReference `xml:"wd:Social_Suffix_Reference,omitempty"`
    LocalPersonName *LocalPersonName `xml:"wd:Local_Person_Name,omitempty"`
}

type PrefixData struct {
    TitleReference *TitleReference `xml:"wd:Title_Reference,omitempty"`
    TitleDescriptor *TitleDescriptor `xml:"wd:Title_Descriptor,omitempty"`
    SalutationReference *SalutationReference `xml:"wd:Salutation_Reference,omitempty"`
}

type TitleReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type TitleDescriptor struct {
    Value *string `xml:",chardata"`
}

type SalutationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    LocalName *string `xml:"wd:Local_Name,attr,omitempty"`
    LocalScript *string `xml:"wd:Local_Script,attr,omitempty"`
    FirstName *FirstName `xml:"wd:First_Name,omitempty"`
    MiddleName *MiddleName `xml:"wd:Middle_Name,omitempty"`
    LastName *LastName `xml:"wd:Last_Name,omitempty"`
    SecondaryLastName *SecondaryLastName `xml:"wd:Secondary_Last_Name,omitempty"`
    FirstName2 *FirstName2 `xml:"wd:First_Name_2,omitempty"`
    MiddleName2 *MiddleName2 `xml:"wd:Middle_Name_2,omitempty"`
    LastName2 *LastName2 `xml:"wd:Last_Name_2,omitempty"`
    SecondaryLastName2 *SecondaryLastName2 `xml:"wd:Secondary_Last_Name_2,omitempty"`
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
    SocialSuffixReference *SocialSuffixReference `xml:"wd:Social_Suffix_Reference,omitempty"`
    SocialSuffixDescriptor *SocialSuffixDescriptor `xml:"wd:Social_Suffix_Descriptor,omitempty"`
    AcademicSuffixReference *AcademicSuffixReference `xml:"wd:Academic_Suffix_Reference,omitempty"`
    HereditarySuffixReference *HereditarySuffixReference `xml:"wd:Hereditary_Suffix_Reference,omitempty"`
    HonorarySuffixReference *HonorarySuffixReference `xml:"wd:Honorary_Suffix_Reference,omitempty"`
    ProfessionalSuffixReference *ProfessionalSuffixReference `xml:"wd:Professional_Suffix_Reference,omitempty"`
    ReligiousSuffixReference *ReligiousSuffixReference `xml:"wd:Religious_Suffix_Reference,omitempty"`
    RoyalSuffixReference *RoyalSuffixReference `xml:"wd:Royal_Suffix_Reference,omitempty"`
}

type SocialSuffixReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SocialSuffixDescriptor struct {
    Value *string `xml:",chardata"`
}

type AcademicSuffixReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type HereditarySuffixReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type HonorarySuffixReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ProfessionalSuffixReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ReligiousSuffixReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type RoyalSuffixReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type FullNameForSingaporeAndMalaysia struct {
    Value *string `xml:",chardata"`
}

type PreferredNameData struct {
    NameDetailData *NameDetailData `xml:"wd:Name_Detail_Data,omitempty"`
}

type AdditionalNameData struct {
    NameDetailData *NameDetailData `xml:"wd:Name_Detail_Data,omitempty"`
    NameTypeReference *NameTypeReference `xml:"wd:Name_Type_Reference,omitempty"`
}

type NameTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PersonalInformationData struct {
    PersonalInformationForCountryData *PersonalInformationForCountryData `xml:"wd:Personal_Information_For_Country_Data,omitempty"`
    BirthDate *BirthDate `xml:"wd:Birth_Date,omitempty"`
    DateOfDeath *DateOfDeath `xml:"wd:Date_of_Death,omitempty"`
    CountryOfBirthReference *CountryOfBirthReference `xml:"wd:Country_of_Birth_Reference,omitempty"`
    CountryRegionOfBirthReference *CountryRegionOfBirthReference `xml:"wd:Country_Region_of_Birth_Reference,omitempty"`
    RegionOfBirthDescriptor *RegionOfBirthDescriptor `xml:"wd:Region_of_Birth_Descriptor,omitempty"`
    CityOfBirth *CityOfBirth `xml:"wd:City_of_Birth,omitempty"`
    CityOfBirthReference *CityOfBirthReference `xml:"wd:City_of_Birth_Reference,omitempty"`
    CitizenshipStatusReference *CitizenshipStatusReference `xml:"wd:Citizenship_Status_Reference,omitempty"`
    PrimaryNationalityReference *PrimaryNationalityReference `xml:"wd:Primary_Nationality_Reference,omitempty"`
    AdditionalNationalitiesReference *AdditionalNationalitiesReference `xml:"wd:Additional_Nationalities_Reference,omitempty"`
    LastMedicalExamValidTo *LastMedicalExamValidTo `xml:"wd:Last_Medical_Exam_Valid_To,omitempty"`
    LastMedicalExamDate *LastMedicalExamDate `xml:"wd:Last_Medical_Exam_Date,omitempty"`
    MedicalExamNotes *MedicalExamNotes `xml:"wd:Medical_Exam_Notes,omitempty"`
    BloodTypeReference *BloodTypeReference `xml:"wd:Blood_Type_Reference,omitempty"`
    MilitaryServiceData *MilitaryServiceData `xml:"wd:Military_Service_Data,omitempty"`
    SexualOrientationReference *SexualOrientationReference `xml:"wd:Sexual_Orientation_Reference,omitempty"`
    GenderIdentityReference *GenderIdentityReference `xml:"wd:Gender_Identity_Reference,omitempty"`
    PronounReference *PronounReference `xml:"wd:Pronoun_Reference,omitempty"`
    NonCountrySpecificSectionData *NonCountrySpecificSectionData `xml:"wd:Non_Country_Specific_Section_Data,omitempty"`
    DateOfBirth *DateOfBirth `xml:"wd:Date_of_Birth,omitempty"`
    BirthCountryReference *BirthCountryReference `xml:"wd:Birth_Country_Reference,omitempty"`
    BirthRegionReference *BirthRegionReference `xml:"wd:Birth_Region_Reference,omitempty"`
    GenderReference *GenderReference `xml:"wd:Gender_Reference,omitempty"`
    ReportingGenderReference *ReportingGenderReference `xml:"wd:Reporting_Gender_Reference,omitempty"`
    DisabilityInformationData *DisabilityInformationData `xml:"wd:Disability_Information_Data,omitempty"`
    MaritalStatusReference *MaritalStatusReference `xml:"wd:Marital_Status_Reference,omitempty"`
    CitizenshipReference *CitizenshipReference `xml:"wd:Citizenship_Reference,omitempty"`
    AdditionalNationalityReference *AdditionalNationalityReference `xml:"wd:Additional_Nationality_Reference,omitempty"`
    EthnicityReference *EthnicityReference `xml:"wd:Ethnicity_Reference,omitempty"`
    RaceEthnicityDetailsReference *RaceEthnicityDetailsReference `xml:"wd:Race_Ethnicity_Details_Reference,omitempty"`
    HispanicOrLatino *HispanicOrLatino `xml:"wd:Hispanic_or_Latino,omitempty"`
    VisualSurveyEthnicityReference *VisualSurveyEthnicityReference `xml:"wd:Visual_Survey_Ethnicity_Reference,omitempty"`
    HispanicOrLatinoForVisualSurvey *HispanicOrLatinoForVisualSurvey `xml:"wd:Hispanic_or_Latino_for_Visual_Survey,omitempty"`
    AboriginalIndigenousIdentificationReference *AboriginalIndigenousIdentificationReference `xml:"wd:Aboriginal_Indigenous_Identification_Reference,omitempty"`
    AboriginalIndigenousIdentificationDetailsReference *AboriginalIndigenousIdentificationDetailsReference `xml:"wd:Aboriginal_Indigenous_Identification_Details_Reference,omitempty"`
    ReligionReference *ReligionReference `xml:"wd:Religion_Reference,omitempty"`
    HukouRegionReference *HukouRegionReference `xml:"wd:Hukou_Region_Reference,omitempty"`
    HukouSubregionReference *HukouSubregionReference `xml:"wd:Hukou_Subregion_Reference,omitempty"`
    HukouLocality *HukouLocality `xml:"wd:Hukou_Locality,omitempty"`
    HukouPostalCode *HukouPostalCode `xml:"wd:Hukou_Postal_Code,omitempty"`
    HukouTypeReference *HukouTypeReference `xml:"wd:Hukou_Type_Reference,omitempty"`
    NativeRegionReference *NativeRegionReference `xml:"wd:Native_Region_Reference,omitempty"`
    PersonnelFileAgency *PersonnelFileAgency `xml:"wd:Personnel_File_Agency,omitempty"`
    MilitaryInformationData *MilitaryInformationData `xml:"wd:Military_Information_Data,omitempty"`
    PoliticalAffiliationReference *PoliticalAffiliationReference `xml:"wd:Political_Affiliation_Reference,omitempty"`
    MaritalStatusDate *MaritalStatusDate `xml:"wd:Marital_Status_Date,omitempty"`
    UsesTobacco *UsesTobacco `xml:"wd:Uses_Tobacco,omitempty"`
    SocialBenefitsLocalityReference *SocialBenefitsLocalityReference `xml:"wd:Social_Benefits_Locality_Reference,omitempty"`
    LgbtIdentificationReference *LgbtIdentificationReference `xml:"wd:LGBT_Identification_Reference,omitempty"`
    RelativeNameInformationData *RelativeNameInformationData `xml:"wd:Relative_Name_Information_Data,omitempty"`
    VeteransPreferenceForRifReference *VeteransPreferenceForRifReference `xml:"wd:Veterans_Preference_For_RIF_Reference,omitempty"`
    SelectiveServiceRegistrationReference *SelectiveServiceRegistrationReference `xml:"wd:Selective_Service_Registration_Reference,omitempty"`
    VeteransPreferenceReference *VeteransPreferenceReference `xml:"wd:Veterans_Preference_Reference,omitempty"`
    ActiveMilitaryUniformedServiceReference *ActiveMilitaryUniformedServiceReference `xml:"wd:Active_Military_Uniformed_Service_Reference,omitempty"`
    DisabledVeteranLeaveDateReference *DisabledVeteranLeaveDateReference `xml:"wd:Disabled_Veteran_Leave_Date_Reference,omitempty"`
    UniformServiceReserveStatusReference *UniformServiceReserveStatusReference `xml:"wd:Uniform_Service_Reserve_Status_Reference,omitempty"`
    CountrySpecificSectionData *CountrySpecificSectionData `xml:"wd:Country_Specific_Section_Data,omitempty"`
    VeteransStatusReference *VeteransStatusReference `xml:"wd:Veterans_Status_Reference,omitempty"`
    DisabilityStatusReference *DisabilityStatusReference `xml:"wd:Disability_Status_Reference,omitempty"`
    DisabilityStatusLastUpdatedOn *DisabilityStatusLastUpdatedOn `xml:"wd:Disability_Status_Last_Updated_On,omitempty"`
    AdjudicationInfoData *AdjudicationInfoData `xml:"wd:Adjudication_Info_Data,omitempty"`
    VeteransPreferenceAttachmentData *VeteransPreferenceAttachmentData `xml:"wd:Veterans_Preference_Attachment_Data,omitempty"`
}

type PersonalInformationForCountryData struct {
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    CountryPersonalInformationData *CountryPersonalInformationData `xml:"wd:Country_Personal_Information_Data,omitempty"`
}

type CountryPersonalInformationData struct {
    MaritalStatusReference *MaritalStatusReference `xml:"wd:Marital_Status_Reference,omitempty"`
    MaritalStatusDate *MaritalStatusDate `xml:"wd:Marital_Status_Date,omitempty"`
    ReligionReference *ReligionReference `xml:"wd:Religion_Reference,omitempty"`
    DisabilityStatusData *DisabilityStatusData `xml:"wd:Disability_Status_Data,omitempty"`
    EthnicityReference *EthnicityReference `xml:"wd:Ethnicity_Reference,omitempty"`
    RaceEthnicityDetailsReference *RaceEthnicityDetailsReference `xml:"wd:Race_Ethnicity_Details_Reference,omitempty"`
    EthnicityVisualSurveyReference *EthnicityVisualSurveyReference `xml:"wd:Ethnicity_Visual_Survey_Reference,omitempty"`
    HispanicOrLatino *HispanicOrLatino `xml:"wd:Hispanic_or_Latino,omitempty"`
    HispanicOrLatinoVisualSurvey *HispanicOrLatinoVisualSurvey `xml:"wd:Hispanic_or_Latino_Visual_Survey,omitempty"`
    AboriginalIndigenousIdentificationReference *AboriginalIndigenousIdentificationReference `xml:"wd:Aboriginal_Indigenous_Identification_Reference,omitempty"`
    AboriginalIndigenousIdentificationDetailsReference *AboriginalIndigenousIdentificationDetailsReference `xml:"wd:Aboriginal_Indigenous_Identification_Details_Reference,omitempty"`
    HukouRegionReference *HukouRegionReference `xml:"wd:Hukou_Region_Reference,omitempty"`
    HukouSubregionReference *HukouSubregionReference `xml:"wd:Hukou_Subregion_Reference,omitempty"`
    HukouLocality *HukouLocality `xml:"wd:Hukou_Locality,omitempty"`
    HukouPostalCode *HukouPostalCode `xml:"wd:Hukou_Postal_Code,omitempty"`
    HukouTypeReference *HukouTypeReference `xml:"wd:Hukou_Type_Reference,omitempty"`
    LocalHukou *LocalHukou `xml:"wd:Local_Hukou,omitempty"`
    NativeRegionReference *NativeRegionReference `xml:"wd:Native_Region_Reference,omitempty"`
    NativeRegionDescriptor *NativeRegionDescriptor `xml:"wd:Native_Region_Descriptor,omitempty"`
    PersonnelFileAgencyForPerson *PersonnelFileAgencyForPerson `xml:"wd:Personnel_File_Agency_for_Person,omitempty"`
    PoliticalAffiliationReference *PoliticalAffiliationReference `xml:"wd:Political_Affiliation_Reference,omitempty"`
    SocialBenefitsLocalityReference *SocialBenefitsLocalityReference `xml:"wd:Social_Benefits_Locality_Reference,omitempty"`
    RelativeNameInformationData *RelativeNameInformationData `xml:"wd:Relative_Name_Information_Data,omitempty"`
    SexualOrientationAndGenderIdentityReference *SexualOrientationAndGenderIdentityReference `xml:"wd:Sexual_Orientation_and_Gender_Identity_Reference,omitempty"`
    GenderReference *GenderReference `xml:"wd:Gender_Reference,omitempty"`
    ReportingGenderReference *ReportingGenderReference `xml:"wd:Reporting_Gender_Reference,omitempty"`
    VeteransPreferenceReference *VeteransPreferenceReference `xml:"wd:Veterans_Preference_Reference,omitempty"`
    VeteransPreferenceForRifReference *VeteransPreferenceForRifReference `xml:"wd:Veterans_Preference_For_RIF_Reference,omitempty"`
    SelectiveServiceRegistrationReference *SelectiveServiceRegistrationReference `xml:"wd:Selective_Service_Registration_Reference,omitempty"`
    DisabledVeteranLeaveDateReference *DisabledVeteranLeaveDateReference `xml:"wd:Disabled_Veteran_Leave_Date_Reference,omitempty"`
    ActiveMilitaryUniformedServiceReference *ActiveMilitaryUniformedServiceReference `xml:"wd:Active_Military_Uniformed_Service_Reference,omitempty"`
    UniformServiceReserveStatusReference *UniformServiceReserveStatusReference `xml:"wd:Uniform_Service_Reserve_Status_Reference,omitempty"`
    CountrySpecificSectionData *CountrySpecificSectionData `xml:"wd:Country_Specific_Section_Data,omitempty"`
}

type MaritalStatusReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type MaritalStatusDate struct {
    Value *string `xml:",chardata"`
}

type ReligionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DisabilityStatusData struct {
    DisabilityReference *DisabilityReference `xml:"wd:Disability_Reference,omitempty"`
    DisabilityStatusDate *DisabilityStatusDate `xml:"wd:Disability_Status_Date,omitempty"`
    DisabilityDateKnown *DisabilityDateKnown `xml:"wd:Disability_Date_Known,omitempty"`
    DisabilityEndDate *DisabilityEndDate `xml:"wd:Disability_End_Date,omitempty"`
    DisabilityGradeReference *DisabilityGradeReference `xml:"wd:Disability_Grade_Reference,omitempty"`
    DisabilityDegree *DisabilityDegree `xml:"wd:Disability_Degree,omitempty"`
    DisabilityRemainingCapacity *DisabilityRemainingCapacity `xml:"wd:Disability_Remaining_Capacity,omitempty"`
    DisabilityCertificationAuthorityReference *DisabilityCertificationAuthorityReference `xml:"wd:Disability_Certification_Authority_Reference,omitempty"`
    DisabilityCertificationAuthority *DisabilityCertificationAuthority `xml:"wd:Disability_Certification_Authority,omitempty"`
    DisabilityCertifiedAt *DisabilityCertifiedAt `xml:"wd:Disability_Certified_At,omitempty"`
    DisabilityCertificationId *DisabilityCertificationId `xml:"wd:Disability_Certification_ID,omitempty"`
    DisabilityCertificationBasisReference *DisabilityCertificationBasisReference `xml:"wd:Disability_Certification_Basis_Reference,omitempty"`
    DisabilitySeverityRecognitionDate *DisabilitySeverityRecognitionDate `xml:"wd:Disability_Severity_Recognition_Date,omitempty"`
    DisabilityFteTowardQuota *DisabilityFteTowardQuota `xml:"wd:Disability_FTE_Toward_Quota,omitempty"`
    DisabilityWorkRestrictions *DisabilityWorkRestrictions `xml:"wd:Disability_Work_Restrictions,omitempty"`
    DisabilityAccommodationsRequested *DisabilityAccommodationsRequested `xml:"wd:Disability_Accommodations_Requested,omitempty"`
    DisabilityAccommodationsProvided *DisabilityAccommodationsProvided `xml:"wd:Disability_Accommodations_Provided,omitempty"`
    DisabilityRehabilitationRequested *DisabilityRehabilitationRequested `xml:"wd:Disability_Rehabilitation_Requested,omitempty"`
    DisabilityRehabilitationProvided *DisabilityRehabilitationProvided `xml:"wd:Disability_Rehabilitation_Provided,omitempty"`
    Note *Note `xml:"wd:Note,omitempty"`
    WorkerDocumentReference *WorkerDocumentReference `xml:"wd:Worker_Document_Reference,omitempty"`
    DisabilityStatusReference *DisabilityStatusReference `xml:"wd:Disability_Status_Reference,omitempty"`
    Delete *string `xml:"wd:Delete,attr,omitempty"`
}

type DisabilityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DisabilityDegree struct {
    Value *string `xml:",chardata"`
}

type DisabilityRemainingCapacity struct {
    Value *string `xml:",chardata"`
}

type DisabilityCertificationAuthorityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DisabilityStatusReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type EthnicityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type RaceEthnicityDetailsReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type EthnicityVisualSurveyReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type HispanicOrLatino struct {
    Value *string `xml:",chardata"`
}

type HispanicOrLatinoVisualSurvey struct {
    Value *string `xml:",chardata"`
}

type AboriginalIndigenousIdentificationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AboriginalIndigenousIdentificationDetailsReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type HukouRegionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type HukouSubregionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type HukouLocality struct {
    Value *string `xml:",chardata"`
}

type HukouPostalCode struct {
    Value *string `xml:",chardata"`
}

type HukouTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type LocalHukou struct {
    Value *string `xml:",chardata"`
}

type NativeRegionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type NativeRegionDescriptor struct {
    Value *string `xml:",chardata"`
}

type PersonnelFileAgencyForPerson struct {
    Value *string `xml:",chardata"`
}

type PoliticalAffiliationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SocialBenefitsLocalityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type RelativeNameInformationData struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    RelativeNameReferenceReference *RelativeNameReferenceReference `xml:"wd:Relative_Name_Reference_Reference,omitempty"`
    RelativeTypeReference *RelativeTypeReference `xml:"wd:Relative_Type_Reference,omitempty"`
    NameDetailData *NameDetailData `xml:"wd:Name_Detail_Data,omitempty"`
    ReplaceAll *string `xml:"wd:Replace_All,attr,omitempty"`
    RelativeNameData *RelativeNameData `xml:"wd:Relative_Name_Data,omitempty"`
}

type RelativeNameReferenceReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type RelativeTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SexualOrientationAndGenderIdentityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type GenderReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ReportingGenderReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type VeteransPreferenceReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type VeteransPreferenceForRifReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SelectiveServiceRegistrationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DisabledVeteranLeaveDateReference struct {
    Value *string `xml:",chardata"`
}

type ActiveMilitaryUniformedServiceReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type UniformServiceReserveStatusReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CountrySpecificSectionData struct {
    CountrySpecificSection1Data *CountrySpecificSection1Data `xml:"wd:Country_Specific_Section_1_Data,omitempty"`
    CountrySpecificSection2Data *CountrySpecificSection2Data `xml:"wd:Country_Specific_Section_2_Data,omitempty"`
    CountrySpecificSection3Data *CountrySpecificSection3Data `xml:"wd:Country_Specific_Section_3_Data,omitempty"`
}

type CountrySpecificSection1Data struct {
    Field1Reference *Field1Reference `xml:"wd:Field_1_Reference,omitempty"`
    Field2Reference *Field2Reference `xml:"wd:Field_2_Reference,omitempty"`
    Field3Reference *Field3Reference `xml:"wd:Field_3_Reference,omitempty"`
}

type Field1Reference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Field2Reference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Field3Reference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CountrySpecificSection2Data struct {
    Field1Reference *Field1Reference `xml:"wd:Field_1_Reference,omitempty"`
    Field2Reference *Field2Reference `xml:"wd:Field_2_Reference,omitempty"`
    Field3Reference *Field3Reference `xml:"wd:Field_3_Reference,omitempty"`
}

type CountrySpecificSection3Data struct {
    Field1Reference *Field1Reference `xml:"wd:Field_1_Reference,omitempty"`
    Field2Reference *Field2Reference `xml:"wd:Field_2_Reference,omitempty"`
    Field3Reference *Field3Reference `xml:"wd:Field_3_Reference,omitempty"`
}

type BirthDate struct {
    Value *string `xml:",chardata"`
}

type DateOfDeath struct {
    Value *string `xml:",chardata"`
}

type CountryOfBirthReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CountryRegionOfBirthReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type RegionOfBirthDescriptor struct {
    Value *string `xml:",chardata"`
}

type CityOfBirth struct {
    Value *string `xml:",chardata"`
}

type CityOfBirthReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CitizenshipStatusReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PrimaryNationalityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AdditionalNationalitiesReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type MilitaryServiceData struct {
    StatusReference *StatusReference `xml:"wd:Status_Reference,omitempty"`
    DischargeDate *DischargeDate `xml:"wd:Discharge_Date,omitempty"`
    StatusBeginDate *StatusBeginDate `xml:"wd:Status_Begin_Date,omitempty"`
    MilitaryServiceTypeReference *MilitaryServiceTypeReference `xml:"wd:Military_Service_Type_Reference,omitempty"`
    MilitaryRankReference *MilitaryRankReference `xml:"wd:Military_Rank_Reference,omitempty"`
    Notes *Notes `xml:"wd:Notes,omitempty"`
    MilitaryServiceReference *MilitaryServiceReference `xml:"wd:Military_Service_Reference,omitempty"`
    MilitaryDischargeTypeReference *MilitaryDischargeTypeReference `xml:"wd:Military_Discharge_Type_Reference,omitempty"`
    MilitaryStatusReference *MilitaryStatusReference `xml:"wd:Military_Status_Reference,omitempty"`
    MilitaryDischargeDate *MilitaryDischargeDate `xml:"wd:Military_Discharge_Date,omitempty"`
    MilitaryStatusBeginDate *MilitaryStatusBeginDate `xml:"wd:Military_Status_Begin_Date,omitempty"`
}

type StatusReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DischargeDate struct {
    Value *string `xml:",chardata"`
}

type StatusBeginDate struct {
    Value *string `xml:",chardata"`
}

type MilitaryServiceTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type MilitaryRankReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Notes struct {
    Value *string `xml:",chardata"`
}

type MilitaryServiceReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type MilitaryDischargeTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SexualOrientationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type GenderIdentityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PronounReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type NonCountrySpecificSectionData struct {
    NonCountrySpecificSection1Data *NonCountrySpecificSection1Data `xml:"wd:Non_Country_Specific_Section_1_Data,omitempty"`
    NonCountrySpecificSection2Data *NonCountrySpecificSection2Data `xml:"wd:Non_Country_Specific_Section_2_Data,omitempty"`
    NonCountrySpecificSection3Data *NonCountrySpecificSection3Data `xml:"wd:Non_Country_Specific_Section_3_Data,omitempty"`
}

type NonCountrySpecificSection1Data struct {
    Field1Reference *Field1Reference `xml:"wd:Field_1_Reference,omitempty"`
    Field2Reference *Field2Reference `xml:"wd:Field_2_Reference,omitempty"`
    Field3Reference *Field3Reference `xml:"wd:Field_3_Reference,omitempty"`
}

type NonCountrySpecificSection2Data struct {
    Field1Reference *Field1Reference `xml:"wd:Field_1_Reference,omitempty"`
    Field2Reference *Field2Reference `xml:"wd:Field_2_Reference,omitempty"`
    Field3Reference *Field3Reference `xml:"wd:Field_3_Reference,omitempty"`
}

type NonCountrySpecificSection3Data struct {
    Field1Reference *Field1Reference `xml:"wd:Field_1_Reference,omitempty"`
    Field2Reference *Field2Reference `xml:"wd:Field_2_Reference,omitempty"`
    Field3Reference *Field3Reference `xml:"wd:Field_3_Reference,omitempty"`
}

type IdentificationData struct {
    NationalId *NationalId `xml:"wd:National_ID,omitempty"`
    GovernmentId *GovernmentId `xml:"wd:Government_ID,omitempty"`
    VisaId *VisaId `xml:"wd:Visa_ID,omitempty"`
    PassportId *PassportId `xml:"wd:Passport_ID,omitempty"`
    LicenseId *LicenseId `xml:"wd:License_ID,omitempty"`
    CustomId *CustomId `xml:"wd:Custom_ID,omitempty"`
}

type NationalId struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    NationalIdReference *NationalIdReference `xml:"wd:National_ID_Reference,omitempty"`
    NationalIdData *NationalIdData `xml:"wd:National_ID_Data,omitempty"`
    NationalIdSharedReference *NationalIdSharedReference `xml:"wd:National_ID_Shared_Reference,omitempty"`
}

type NationalIdReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type NationalIdData struct {
    Id *Id `xml:"wd:ID,omitempty"`
    IdTypeReference *IdTypeReference `xml:"wd:ID_Type_Reference,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    IssuedDate *IssuedDate `xml:"wd:Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"wd:Expiration_Date,omitempty"`
    VerificationDate *VerificationDate `xml:"wd:Verification_Date,omitempty"`
    Series *Series `xml:"wd:Series,omitempty"`
    IssuingAgency *IssuingAgency `xml:"wd:Issuing_Agency,omitempty"`
    VerifiedByReference *VerifiedByReference `xml:"wd:Verified_By_Reference,omitempty"`
}

type IdTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type NationalIdSharedReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type GovernmentId struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    GovernmentIdReference *GovernmentIdReference `xml:"wd:Government_ID_Reference,omitempty"`
    GovernmentIdData *GovernmentIdData `xml:"wd:Government_ID_Data,omitempty"`
    GovernmentIdSharedReference *GovernmentIdSharedReference `xml:"wd:Government_ID_Shared_Reference,omitempty"`
}

type GovernmentIdReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type GovernmentIdData struct {
    Id *Id `xml:"wd:ID,omitempty"`
    IdTypeReference *IdTypeReference `xml:"wd:ID_Type_Reference,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    IssuedDate *IssuedDate `xml:"wd:Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"wd:Expiration_Date,omitempty"`
    VerificationDate *VerificationDate `xml:"wd:Verification_Date,omitempty"`
    VerifiedByReference *VerifiedByReference `xml:"wd:Verified_by_Reference,omitempty"`
}

type GovernmentIdSharedReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type VisaId struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    VisaIdReference *VisaIdReference `xml:"wd:Visa_ID_Reference,omitempty"`
    VisaIdData *VisaIdData `xml:"wd:Visa_ID_Data,omitempty"`
    VisaIdSharedReference *VisaIdSharedReference `xml:"wd:Visa_ID_Shared_Reference,omitempty"`
}

type VisaIdReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type VisaIdData struct {
    Id *Id `xml:"wd:ID,omitempty"`
    IdTypeReference *IdTypeReference `xml:"wd:ID_Type_Reference,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    IssuedDate *IssuedDate `xml:"wd:Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"wd:Expiration_Date,omitempty"`
    VerificationDate *VerificationDate `xml:"wd:Verification_Date,omitempty"`
    VerifiedByReference *VerifiedByReference `xml:"wd:Verified_By_Reference,omitempty"`
}

type VisaIdSharedReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PassportId struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    PassportIdReference *PassportIdReference `xml:"wd:Passport_ID_Reference,omitempty"`
    PassportIdData *PassportIdData `xml:"wd:Passport_ID_Data,omitempty"`
    PassportIdSharedReference *PassportIdSharedReference `xml:"wd:Passport_ID_Shared_Reference,omitempty"`
}

type PassportIdReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PassportIdData struct {
    Id *Id `xml:"wd:ID,omitempty"`
    IdTypeReference *IdTypeReference `xml:"wd:ID_Type_Reference,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    IssuedDate *IssuedDate `xml:"wd:Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"wd:Expiration_Date,omitempty"`
    VerificationDate *VerificationDate `xml:"wd:Verification_Date,omitempty"`
    VerifiedByReference *VerifiedByReference `xml:"wd:Verified_By_Reference,omitempty"`
}

type PassportIdSharedReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type LicenseId struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    LicenseIdReference *LicenseIdReference `xml:"wd:License_ID_Reference,omitempty"`
    LicenseIdData *LicenseIdData `xml:"wd:License_ID_Data,omitempty"`
    LicenseIdSharedReference *LicenseIdSharedReference `xml:"wd:License_ID_Shared_Reference,omitempty"`
}

type LicenseIdReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type LicenseIdData struct {
    Id *Id `xml:"wd:ID,omitempty"`
    IdTypeReference *IdTypeReference `xml:"wd:ID_Type_Reference,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    CountryRegionReference *CountryRegionReference `xml:"wd:Country_Region_Reference,omitempty"`
    CountryRegionDescriptor *CountryRegionDescriptor `xml:"wd:Country_Region_Descriptor,omitempty"`
    AuthorityReference *AuthorityReference `xml:"wd:Authority_Reference,omitempty"`
    LicenseClass *LicenseClass `xml:"wd:License_Class,omitempty"`
    IssuedDate *IssuedDate `xml:"wd:Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"wd:Expiration_Date,omitempty"`
    VerificationDate *VerificationDate `xml:"wd:Verification_Date,omitempty"`
    VerifiedByReference *VerifiedByReference `xml:"wd:Verified_By_Reference,omitempty"`
}

type CountryRegionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CountryRegionDescriptor struct {
    Value *string `xml:",chardata"`
}

type AuthorityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type LicenseClass struct {
    Value *string `xml:",chardata"`
}

type LicenseIdSharedReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CustomId struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    CustomIdReference *CustomIdReference `xml:"wd:Custom_ID_Reference,omitempty"`
    CustomIdData *CustomIdData `xml:"wd:Custom_ID_Data,omitempty"`
    CustomIdSharedReference *CustomIdSharedReference `xml:"wd:Custom_ID_Shared_Reference,omitempty"`
}

type CustomIdReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CustomIdData struct {
    Id *Id `xml:"wd:ID,omitempty"`
    IdTypeReference *IdTypeReference `xml:"wd:ID_Type_Reference,omitempty"`
    IssuedDate *IssuedDate `xml:"wd:Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"wd:Expiration_Date,omitempty"`
    OrganizationIdReference *OrganizationIdReference `xml:"wd:Organization_ID_Reference,omitempty"`
    CustomDescription *CustomDescription `xml:"wd:Custom_Description,omitempty"`
}

type OrganizationIdReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CustomDescription struct {
    Value *string `xml:",chardata"`
}

type CustomIdSharedReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ContactData struct {
    AddressData *AddressData `xml:"wd:Address_Data,omitempty"`
    PhoneData *PhoneData `xml:"wd:Phone_Data,omitempty"`
    EmailAddressData *EmailAddressData `xml:"wd:Email_Address_Data,omitempty"`
    InstantMessengerData *InstantMessengerData `xml:"wd:Instant_Messenger_Data,omitempty"`
    WebAddressData *WebAddressData `xml:"wd:Web_Address_Data,omitempty"`
    PhoneDeviceTypeReference *PhoneDeviceTypeReference `xml:"wd:Phone_Device_Type_Reference,omitempty"`
    CountryPhoneCodeReference *CountryPhoneCodeReference `xml:"wd:Country_Phone_Code_Reference,omitempty"`
    PhoneNumber *PhoneNumber `xml:"wd:Phone_Number,omitempty"`
    PhoneExtension *PhoneExtension `xml:"wd:Phone_Extension,omitempty"`
    EmailAddress *EmailAddress `xml:"wd:Email_Address,omitempty"`
    WebsiteData *WebsiteData `xml:"wd:Website_Data,omitempty"`
    LocationData *LocationData `xml:"wd:Location_Data,omitempty"`
}

type AddressData struct {
    FormattedAddress *string `xml:"wd:Formatted_Address,attr,omitempty"`
    AddressFormatType *string `xml:"wd:Address_Format_Type,attr,omitempty"`
    DefaultedBusinessSiteAddress *string `xml:"wd:Defaulted_Business_Site_Address,attr,omitempty"`
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    DoNotReplaceAll *string `xml:"wd:Do_Not_Replace_All,attr,omitempty"`
    EffectiveDate *string `xml:"wd:Effective_Date,attr,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    LastModified *LastModified `xml:"wd:Last_Modified,omitempty"`
    AddressLineData *AddressLineData `xml:"wd:Address_Line_Data,omitempty"`
    Municipality *Municipality `xml:"wd:Municipality,omitempty"`
    CountryCityReference *CountryCityReference `xml:"wd:Country_City_Reference,omitempty"`
    SubmunicipalityData *SubmunicipalityData `xml:"wd:Submunicipality_Data,omitempty"`
    CountryRegionReference *CountryRegionReference `xml:"wd:Country_Region_Reference,omitempty"`
    CountryRegionDescriptor *CountryRegionDescriptor `xml:"wd:Country_Region_Descriptor,omitempty"`
    SubregionData *SubregionData `xml:"wd:Subregion_Data,omitempty"`
    PostalCode *PostalCode `xml:"wd:Postal_Code,omitempty"`
    UsageData *UsageData `xml:"wd:Usage_Data,omitempty"`
    NumberOfDays *NumberOfDays `xml:"wd:Number_of_Days,omitempty"`
    MunicipalityLocal *MunicipalityLocal `xml:"wd:Municipality_Local,omitempty"`
    AddressReference *AddressReference `xml:"wd:Address_Reference,omitempty"`
    AddressId *AddressId `xml:"wd:Address_ID,omitempty"`
    DefaultAddress *string `xml:"wd:Default_Address,attr,omitempty"`
}

type LastModified struct {
    Value *string `xml:",chardata"`
}

type AddressLineData struct {
    Value *string `xml:",chardata"`
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Type *string `xml:"wd:Type,attr,omitempty"`
}

type Municipality struct {
    Value *string `xml:",chardata"`
}

type CountryCityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SubmunicipalityData struct {
    Value *string `xml:",chardata"`
    AddressComponentName *string `xml:"wd:Address_Component_Name,attr,omitempty"`
    Type *string `xml:"wd:Type,attr,omitempty"`
}

type SubregionData struct {
    Value *string `xml:",chardata"`
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Type *string `xml:"wd:Type,attr,omitempty"`
}

type PostalCode struct {
    Value *string `xml:",chardata"`
}

type UsageData struct {
    Public *string `xml:"wd:Public,attr,omitempty"`
    TypeData *TypeData `xml:"wd:Type_Data,omitempty"`
    UseForReference *UseForReference `xml:"wd:Use_For_Reference,omitempty"`
    UseForTenantedReference *UseForTenantedReference `xml:"wd:Use_For_Tenanted_Reference,omitempty"`
    Comments *Comments `xml:"wd:Comments,omitempty"`
}

type TypeData struct {
    Primary *string `xml:"wd:Primary,attr,omitempty"`
    TypeReference *TypeReference `xml:"wd:Type_Reference,omitempty"`
}

type TypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type UseForReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type UseForTenantedReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AddressId struct {
    Value *string `xml:",chardata"`
}

type PhoneData struct {
    AreaCode *string `xml:"wd:Area_Code,attr,omitempty"`
    TenantFormattedPhone *string `xml:"wd:Tenant_Formatted_Phone,attr,omitempty"`
    InternationalFormattedPhone *string `xml:"wd:International_Formatted_Phone,attr,omitempty"`
    PhoneNumberWithoutAreaCode *string `xml:"wd:Phone_Number_Without_Area_Code,attr,omitempty"`
    NationalFormattedPhone *string `xml:"wd:National_Formatted_Phone,attr,omitempty"`
    E164FormattedPhone *string `xml:"wd:E164_Formatted_Phone,attr,omitempty"`
    WorkdayTraditionalFormattedPhone *string `xml:"wd:Workday_Traditional_Formatted_Phone,attr,omitempty"`
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    DoNotReplaceAll *string `xml:"wd:Do_Not_Replace_All,attr,omitempty"`
    CountryIsoCode *CountryIsoCode `xml:"wd:Country_ISO_Code,omitempty"`
    InternationalPhoneCode *InternationalPhoneCode `xml:"wd:International_Phone_Code,omitempty"`
    PhoneNumber *PhoneNumber `xml:"wd:Phone_Number,omitempty"`
    PhoneExtension *PhoneExtension `xml:"wd:Phone_Extension,omitempty"`
    PhoneDeviceTypeReference *PhoneDeviceTypeReference `xml:"wd:Phone_Device_Type_Reference,omitempty"`
    UsageData *UsageData `xml:"wd:Usage_Data,omitempty"`
    PhoneReference *PhoneReference `xml:"wd:Phone_Reference,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
    FormattedPhone *string `xml:"wd:Formatted_Phone,attr,omitempty"`
    DeviceTypeReference *DeviceTypeReference `xml:"wd:Device_Type_Reference,omitempty"`
    CountryCodeReference *CountryCodeReference `xml:"wd:Country_Code_Reference,omitempty"`
    CompletePhoneNumber *CompletePhoneNumber `xml:"wd:Complete_Phone_Number,omitempty"`
    Extension *Extension `xml:"wd:Extension,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PhoneReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type EmailAddressData struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    DoNotReplaceAll *string `xml:"wd:Do_Not_Replace_All,attr,omitempty"`
    EmailAddress *EmailAddress `xml:"wd:Email_Address,omitempty"`
    EmailComment *EmailComment `xml:"wd:Email_Comment,omitempty"`
    UsageData *UsageData `xml:"wd:Usage_Data,omitempty"`
    EmailReference *EmailReference `xml:"wd:Email_Reference,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type EmailAddress struct {
    Value *string `xml:",chardata"`
}

type EmailComment struct {
    Value *string `xml:",chardata"`
}

type EmailReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type InstantMessengerData struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    DoNotReplaceAll *string `xml:"wd:Do_Not_Replace_All,attr,omitempty"`
    InstantMessengerAddress *InstantMessengerAddress `xml:"wd:Instant_Messenger_Address,omitempty"`
    InstantMessengerTypeReference *InstantMessengerTypeReference `xml:"wd:Instant_Messenger_Type_Reference,omitempty"`
    InstantMessengerComment *InstantMessengerComment `xml:"wd:Instant_Messenger_Comment,omitempty"`
    UsageData *UsageData `xml:"wd:Usage_Data,omitempty"`
    InstantMessengerReference *InstantMessengerReference `xml:"wd:Instant_Messenger_Reference,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type InstantMessengerAddress struct {
    Value *string `xml:",chardata"`
}

type InstantMessengerTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type InstantMessengerComment struct {
    Value *string `xml:",chardata"`
}

type InstantMessengerReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WebAddressData struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    DoNotReplaceAll *string `xml:"wd:Do_Not_Replace_All,attr,omitempty"`
    WebAddress *WebAddress `xml:"wd:Web_Address,omitempty"`
    WebAddressComment *WebAddressComment `xml:"wd:Web_Address_Comment,omitempty"`
    UsageData *UsageData `xml:"wd:Usage_Data,omitempty"`
    WebAddressReference *WebAddressReference `xml:"wd:Web_Address_Reference,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WebAddress struct {
    Value *string `xml:",chardata"`
}

type WebAddressComment struct {
    Value *string `xml:",chardata"`
}

type WebAddressReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type TobaccoUse struct {
    Value *string `xml:",chardata"`
}

type EmploymentData struct {
    WorkerJobData *WorkerJobData `xml:"wd:Worker_Job_Data,omitempty"`
    WorkerStatusData *WorkerStatusData `xml:"wd:Worker_Status_Data,omitempty"`
}

type WorkerJobData struct {
    PrimaryJob *string `xml:"wd:Primary_Job,attr,omitempty"`
    PositionData *PositionData `xml:"wd:Position_Data,omitempty"`
    PositionOrganizationsData *PositionOrganizationsData `xml:"wd:Position_Organizations_Data,omitempty"`
    PositionManagementChainsData *PositionManagementChainsData `xml:"wd:Position_Management_Chains_Data,omitempty"`
    PositionContractDetailData *PositionContractDetailData `xml:"wd:Position_Contract_Detail_Data,omitempty"`
}

type PositionData struct {
    EffectiveDate *string `xml:"wd:Effective_Date,attr,omitempty"`
    PositionReference *PositionReference `xml:"wd:Position_Reference,omitempty"`
    HeadcountReference *HeadcountReference `xml:"wd:Headcount_Reference,omitempty"`
    PositionId *PositionId `xml:"wd:Position_ID,omitempty"`
    PositionTitle *PositionTitle `xml:"wd:Position_Title,omitempty"`
    BusinessTitle *BusinessTitle `xml:"wd:Business_Title,omitempty"`
    StartDate *StartDate `xml:"wd:Start_Date,omitempty"`
    EndEmploymentDate *EndEmploymentDate `xml:"wd:End_Employment_Date,omitempty"`
    EndEmploymentReasonReference *EndEmploymentReasonReference `xml:"wd:End_Employment_Reason_Reference,omitempty"`
    WorkerTypeReference *WorkerTypeReference `xml:"wd:Worker_Type_Reference,omitempty"`
    PositionTimeTypeReference *PositionTimeTypeReference `xml:"wd:Position_Time_Type_Reference,omitempty"`
    JobExempt *JobExempt `xml:"wd:Job_Exempt,omitempty"`
    ScheduledWeeklyHours *ScheduledWeeklyHours `xml:"wd:Scheduled_Weekly_Hours,omitempty"`
    DefaultWeeklyHours *DefaultWeeklyHours `xml:"wd:Default_Weekly_Hours,omitempty"`
    WorkingTimeFrequencyReference *WorkingTimeFrequencyReference `xml:"wd:Working_Time_Frequency_Reference,omitempty"`
    WorkingTimeUnitReference *WorkingTimeUnitReference `xml:"wd:Working_Time_Unit_Reference,omitempty"`
    WorkingTimeValue *WorkingTimeValue `xml:"wd:Working_Time_Value,omitempty"`
    FullTimeEquivalentPercentage *FullTimeEquivalentPercentage `xml:"wd:Full_Time_Equivalent_Percentage,omitempty"`
    SpecifyPaidFte *SpecifyPaidFte `xml:"wd:Specify_Paid_FTE,omitempty"`
    PaidFte *PaidFte `xml:"wd:Paid_FTE,omitempty"`
    SpecifyWorkingFte *SpecifyWorkingFte `xml:"wd:Specify_Working_FTE,omitempty"`
    WorkingFte *WorkingFte `xml:"wd:Working_FTE,omitempty"`
    ExcludeFromHeadcount *ExcludeFromHeadcount `xml:"wd:Exclude_from_Headcount,omitempty"`
    PayRateTypeReference *PayRateTypeReference `xml:"wd:Pay_Rate_Type_Reference,omitempty"`
    JobClassificationSummaryData *JobClassificationSummaryData `xml:"wd:Job_Classification_Summary_Data,omitempty"`
    CompanyInsiderReference *CompanyInsiderReference `xml:"wd:Company_Insider_Reference,omitempty"`
    WorkShiftReference *WorkShiftReference `xml:"wd:Work_Shift_Reference,omitempty"`
    WorkHoursProfilesReference *WorkHoursProfilesReference `xml:"wd:Work_Hours_Profiles_Reference,omitempty"`
    FederalWithholdingFein *FederalWithholdingFein `xml:"wd:Federal_Withholding_FEIN,omitempty"`
    WorkerCompensationCodeData *WorkerCompensationCodeData `xml:"wd:Worker_Compensation_Code_Data,omitempty"`
    PositionPayrollReportingCodeData *PositionPayrollReportingCodeData `xml:"wd:Position_Payroll_Reporting_Code_Data,omitempty"`
    JobProfileSummaryData *JobProfileSummaryData `xml:"wd:Job_Profile_Summary_Data,omitempty"`
    BusinessSiteSummaryData *BusinessSiteSummaryData `xml:"wd:Business_Site_Summary_Data,omitempty"`
    PayrollInterfaceProcessingData *PayrollInterfaceProcessingData `xml:"wd:Payroll_Interface_Processing_Data,omitempty"`
    RegularPaidEquivalentHours *RegularPaidEquivalentHours `xml:"wd:Regular_Paid_Equivalent_Hours,omitempty"`
    WorkerHoursProfileClassification *WorkerHoursProfileClassification `xml:"wd:Worker_Hours_Profile_Classification,omitempty"`
    InternationalAssignmentData *InternationalAssignmentData `xml:"wd:International_Assignment_Data,omitempty"`
    WorkSpaceReference *WorkSpaceReference `xml:"wd:Work_Space__Reference,omitempty"`
    AcademicPaySetupData *AcademicPaySetupData `xml:"wd:Academic_Pay_Setup_Data,omitempty"`
    EndDate *EndDate `xml:"wd:End_Date,omitempty"`
    PayThroughDate *PayThroughDate `xml:"wd:Pay_Through_Date,omitempty"`
    CollectiveAgreementSummaryData *CollectiveAgreementSummaryData `xml:"wd:Collective_Agreement_Summary_Data,omitempty"`
    EmployeeProbationPeriodSummaryData *EmployeeProbationPeriodSummaryData `xml:"wd:Employee_Probation_Period_Summary_Data,omitempty"`
    ManagerAsOfLastDetectedManagerChangeReference *ManagerAsOfLastDetectedManagerChangeReference `xml:"wd:Manager_as_of_last_detected_manager_change_Reference,omitempty"`
}

type HeadcountReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WorkerTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PositionTimeTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WorkingTimeUnitReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type JobClassificationSummaryData struct {
    Additional *string `xml:"wd:Additional,attr,omitempty"`
    JobClassificationReference *JobClassificationReference `xml:"wd:Job_Classification_Reference,omitempty"`
    JobGroupReference *JobGroupReference `xml:"wd:Job_Group_Reference,omitempty"`
}

type JobClassificationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type JobGroupReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CompanyInsiderReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WorkShiftReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WorkHoursProfilesReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type FederalWithholdingFein struct {
    Value *string `xml:",chardata"`
}

type WorkerCompensationCodeData struct {
    WorkersCompensationCodeReference *WorkersCompensationCodeReference `xml:"wd:Workers_Compensation_Code_Reference,omitempty"`
    WorkersCompensationCode *WorkersCompensationCode `xml:"wd:Workers_Compensation_Code,omitempty"`
}

type WorkersCompensationCodeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WorkersCompensationCode struct {
    Value *string `xml:",chardata"`
}

type PositionPayrollReportingCodeData struct {
    PayrollReportingCodeReference *PayrollReportingCodeReference `xml:"wd:Payroll_Reporting_Code_Reference,omitempty"`
    Code *Code `xml:"wd:Code,omitempty"`
    FormattedCode *FormattedCode `xml:"wd:Formatted_Code,omitempty"`
    Name *Name `xml:"wd:Name,omitempty"`
    PayrollReportingTypeReference *PayrollReportingTypeReference `xml:"wd:Payroll_Reporting_Type_Reference,omitempty"`
}

type PayrollReportingCodeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type JobProfileSummaryData struct {
    JobProfileReference *JobProfileReference `xml:"wd:Job_Profile_Reference,omitempty"`
    JobExempt *JobExempt `xml:"wd:Job_Exempt,omitempty"`
    ManagementLevelReference *ManagementLevelReference `xml:"wd:Management_Level_Reference,omitempty"`
    JobCategoryReference *JobCategoryReference `xml:"wd:Job_Category_Reference,omitempty"`
    JobFamilyReference *JobFamilyReference `xml:"wd:Job_Family_Reference,omitempty"`
    JobProfileName *JobProfileName `xml:"wd:Job_Profile_Name,omitempty"`
    WorkShiftRequired *WorkShiftRequired `xml:"wd:Work_Shift_Required,omitempty"`
    CriticalJob *CriticalJob `xml:"wd:Critical_Job,omitempty"`
    DifficultyToFillReference *DifficultyToFillReference `xml:"wd:Difficulty_to_Fill_Reference,omitempty"`
}

type JobProfileReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ManagementLevelReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type JobCategoryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type JobFamilyReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type BusinessSiteSummaryData struct {
    LocationReference *LocationReference `xml:"wd:Location_Reference,omitempty"`
    Name *Name `xml:"wd:Name,omitempty"`
    LocationTypeReference *LocationTypeReference `xml:"wd:Location_Type_Reference,omitempty"`
    LocalReference *LocalReference `xml:"wd:Local_Reference,omitempty"`
    DisplayLanguageReference *DisplayLanguageReference `xml:"wd:Display_Language_Reference,omitempty"`
    TimeProfileReference *TimeProfileReference `xml:"wd:Time_Profile_Reference,omitempty"`
    ScheduledWeeklyHours *ScheduledWeeklyHours `xml:"wd:Scheduled_Weekly_Hours,omitempty"`
    AddressData *AddressData `xml:"wd:Address_Data,omitempty"`
}

type LocationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type LocationTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type LocalReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DisplayLanguageReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type TimeProfileReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PayrollInterfaceProcessingData struct {
    EffectiveDate *EffectiveDate `xml:"wd:Effective_Date,omitempty"`
    PayRateTypeReference *PayRateTypeReference `xml:"wd:Pay_Rate_Type_Reference,omitempty"`
    FrequencyReference *FrequencyReference `xml:"wd:Frequency_Reference,omitempty"`
    PayGroupReference *PayGroupReference `xml:"wd:Pay_Group_Reference,omitempty"`
    PayrollEntityReference *PayrollEntityReference `xml:"wd:Payroll_Entity_Reference,omitempty"`
    ExternalEmployeeTypeReference *ExternalEmployeeTypeReference `xml:"wd:External_Employee_Type_Reference,omitempty"`
    PayrollFileNumber *PayrollFileNumber `xml:"wd:Payroll_File_Number,omitempty"`
}

type EffectiveDate struct {
    Value *string `xml:",chardata"`
}

type FrequencyReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PayGroupReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PayrollEntityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ExternalEmployeeTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    InternationalAssignmentTypeReference *InternationalAssignmentTypeReference `xml:"wd:International_Assignment_Type_Reference,omitempty"`
    StartInternationalAssignmentReasonReference *StartInternationalAssignmentReasonReference `xml:"wd:Start_International_Assignment_Reason_Reference,omitempty"`
    ExpectedAssignmentEndDate *ExpectedAssignmentEndDate `xml:"wd:Expected_Assignment_End_Date,omitempty"`
    HostCountryReference *HostCountryReference `xml:"wd:Host_Country_Reference,omitempty"`
    HomeCountryReference *HomeCountryReference `xml:"wd:Home_Country_Reference,omitempty"`
}

type InternationalAssignmentTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type StartInternationalAssignmentReasonReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ExpectedAssignmentEndDate struct {
    Value *string `xml:",chardata"`
}

type HostCountryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type HomeCountryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WorkSpaceReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AcademicPaySetupData struct {
    AnnualWorkPeriodWorkPercentOfYear *AnnualWorkPeriodWorkPercentOfYear `xml:"wd:Annual_Work_Period_Work_Percent_of_Year,omitempty"`
    AnnualWorkPeriodStartDate *AnnualWorkPeriodStartDate `xml:"wd:Annual_Work_Period_Start_Date,omitempty"`
    AnnualWorkPeriodEndDate *AnnualWorkPeriodEndDate `xml:"wd:Annual_Work_Period_End_Date,omitempty"`
    DisbursementPlanPeriodStartDate *DisbursementPlanPeriodStartDate `xml:"wd:Disbursement_Plan_Period_Start_Date,omitempty"`
    DisbursementPlanPeriodEndDate *DisbursementPlanPeriodEndDate `xml:"wd:Disbursement_Plan_Period_End_Date,omitempty"`
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
    CollectiveAgreementData *CollectiveAgreementData `xml:"wd:Collective_Agreement_Data,omitempty"`
}

type CollectiveAgreementData struct {
    AssignEmployeeCollectiveAgreementEventReference *AssignEmployeeCollectiveAgreementEventReference `xml:"wd:Assign_Employee_Collective_Agreement_Event_Reference,omitempty"`
    EffectiveDate *EffectiveDate `xml:"wd:Effective_Date,omitempty"`
    EndAssignmentDate *EndAssignmentDate `xml:"wd:End_Assignment_Date,omitempty"`
    CollectiveAgreementData *CollectiveAgreementData `xml:"wd:Collective_Agreement_Data,omitempty"`
    CollectiveAgreementReference *CollectiveAgreementReference `xml:"wd:Collective_Agreement_Reference,omitempty"`
    CollectiveAgreementReviewDate *CollectiveAgreementReviewDate `xml:"wd:Collective_Agreement_Review_Date,omitempty"`
    CollectiveAgreementFactor *CollectiveAgreementFactor `xml:"wd:Collective_Agreement_Factor,omitempty"`
}

type AssignEmployeeCollectiveAgreementEventReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type EndAssignmentDate struct {
    Value *string `xml:",chardata"`
}

type CollectiveAgreementReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CollectiveAgreementReviewDate struct {
    Value *string `xml:",chardata"`
}

type CollectiveAgreementFactor struct {
    Order *Order `xml:"wd:Order,omitempty"`
    CollectiveAgreementFactorReference *CollectiveAgreementFactorReference `xml:"wd:Collective_Agreement_Factor_Reference,omitempty"`
    CollectiveAgreementFactorOptionReference *CollectiveAgreementFactorOptionReference `xml:"wd:Collective_Agreement_Factor_Option_Reference,omitempty"`
}

type Order struct {
    Value *string `xml:",chardata"`
}

type CollectiveAgreementFactorReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CollectiveAgreementFactorOptionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type EmployeeProbationPeriodSummaryData struct {
    EmployeeProbationPeriodDetailData *EmployeeProbationPeriodDetailData `xml:"wd:Employee_Probation_Period_Detail_Data,omitempty"`
}

type EmployeeProbationPeriodDetailData struct {
    EmployeeProbationPeriodEventReference *EmployeeProbationPeriodEventReference `xml:"wd:Employee_Probation_Period_Event_Reference,omitempty"`
    EmployeeProbationPeriodReference *EmployeeProbationPeriodReference `xml:"wd:Employee_Probation_Period_Reference,omitempty"`
    EffectiveDate *EffectiveDate `xml:"wd:Effective_Date,omitempty"`
    StartDate *StartDate `xml:"wd:Start_Date,omitempty"`
    EndDate *EndDate `xml:"wd:End_Date,omitempty"`
    ProbationTypeReference *ProbationTypeReference `xml:"wd:Probation_Type_Reference,omitempty"`
    ProbationReasonReference *ProbationReasonReference `xml:"wd:Probation_Reason_Reference,omitempty"`
    ExtendedEndDate *ExtendedEndDate `xml:"wd:Extended_End_Date,omitempty"`
    Note *Note `xml:"wd:Note,omitempty"`
}

type EmployeeProbationPeriodEventReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type EmployeeProbationPeriodReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ProbationTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ProbationReasonReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ExtendedEndDate struct {
    Value *string `xml:",chardata"`
}

type ManagerAsOfLastDetectedManagerChangeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PositionOrganizationsData struct {
    PositionOrganizationData *PositionOrganizationData `xml:"wd:Position_Organization_Data,omitempty"`
    PayGroupAssignmentCorrectOrRescindData *PayGroupAssignmentCorrectOrRescindData `xml:"wd:Pay_Group_Assignment_Correct_or_Rescind_Data,omitempty"`
}

type PositionOrganizationData struct {
    OrganizationReference *OrganizationReference `xml:"wd:Organization_Reference,omitempty"`
    OrganizationData *OrganizationData `xml:"wd:Organization_Data,omitempty"`
}

type OrganizationData struct {
    OrganizationReferenceId *OrganizationReferenceId `xml:"wd:Organization_Reference_ID,omitempty"`
    OrganizationCode *OrganizationCode `xml:"wd:Organization_Code,omitempty"`
    IntegrationIdData *IntegrationIdData `xml:"wd:Integration_ID_Data,omitempty"`
    OrganizationName *OrganizationName `xml:"wd:Organization_Name,omitempty"`
    OrganizationTypeReference *OrganizationTypeReference `xml:"wd:Organization_Type_Reference,omitempty"`
    OrganizationSubtypeReference *OrganizationSubtypeReference `xml:"wd:Organization_Subtype_Reference,omitempty"`
    PrimaryBusinessSiteReference *PrimaryBusinessSiteReference `xml:"wd:Primary_Business_Site_Reference,omitempty"`
    OrganizationSupportRoleData *OrganizationSupportRoleData `xml:"wd:Organization_Support_Role_Data,omitempty"`
    DateOfPayGroupAssignment *DateOfPayGroupAssignment `xml:"wd:Date_of_Pay_Group_Assignment,omitempty"`
    UsedInChangeOrganizationAssignments *UsedInChangeOrganizationAssignments `xml:"wd:Used_in_Change_Organization_Assignments,omitempty"`
}

type OrganizationReferenceId struct {
    Value *string `xml:",chardata"`
}

type OrganizationCode struct {
    Value *string `xml:",chardata"`
}

type IntegrationIdData struct {
    Id *Id `xml:"wd:ID,omitempty"`
}

type OrganizationName struct {
    Value *string `xml:",chardata"`
}

type OrganizationTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type OrganizationSubtypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PrimaryBusinessSiteReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type OrganizationSupportRoleData struct {
    OrganizationSupportRole *OrganizationSupportRole `xml:"wd:Organization_Support_Role,omitempty"`
}

type OrganizationSupportRole struct {
    OrganizationRoleReference *OrganizationRoleReference `xml:"wd:Organization_Role_Reference,omitempty"`
    OrganizationRoleData *OrganizationRoleData `xml:"wd:Organization_Role_Data,omitempty"`
}

type OrganizationRoleReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type OrganizationRoleData struct {
    WorkerReference *WorkerReference `xml:"wd:Worker_Reference,omitempty"`
    AssignmentFrom *AssignmentFrom `xml:"wd:Assignment_From,omitempty"`
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
    EffectiveDate *EffectiveDate `xml:"wd:Effective_Date,omitempty"`
    CompletedOn *CompletedOn `xml:"wd:Completed_On,omitempty"`
    EventCorrected *EventCorrected `xml:"wd:Event_Corrected,omitempty"`
    EventRescinded *EventRescinded `xml:"wd:Event_Rescinded,omitempty"`
    OriginalPayGroupReference *OriginalPayGroupReference `xml:"wd:Original_Pay_Group_Reference,omitempty"`
    ProposedPayGroupReference *ProposedPayGroupReference `xml:"wd:Proposed_Pay_Group_Reference,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ProposedPayGroupReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PositionManagementChainsData struct {
    PositionSupervisoryManagementChainData *PositionSupervisoryManagementChainData `xml:"wd:Position_Supervisory_Management_Chain_Data,omitempty"`
    PositionMatrixManagementChainData *PositionMatrixManagementChainData `xml:"wd:Position_Matrix_Management_Chain_Data,omitempty"`
}

type PositionSupervisoryManagementChainData struct {
    ManagementChainData *ManagementChainData `xml:"wd:Management_Chain_Data,omitempty"`
}

type ManagementChainData struct {
    OrganizationReference *OrganizationReference `xml:"wd:Organization_Reference,omitempty"`
    ManagerReference *ManagerReference `xml:"wd:Manager_Reference,omitempty"`
    Manager *Manager `xml:"wd:Manager,omitempty"`
}

type ManagerReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Manager struct {
    WorkerReference *WorkerReference `xml:"wd:Worker_Reference,omitempty"`
    WorkerDescriptor *WorkerDescriptor `xml:"wd:Worker_Descriptor,omitempty"`
}

type PositionMatrixManagementChainData struct {
    ManagementChainData *ManagementChainData `xml:"wd:Management_Chain_Data,omitempty"`
}

type PositionContractDetailData struct {
    ContractPayRate *ContractPayRate `xml:"wd:Contract_Pay_Rate,omitempty"`
    CurrencyReference *CurrencyReference `xml:"wd:Currency_Reference,omitempty"`
    FrequencyReference *FrequencyReference `xml:"wd:Frequency_Reference,omitempty"`
    ContractAssignmentDetails *ContractAssignmentDetails `xml:"wd:Contract_Assignment_Details,omitempty"`
    ContractEndDate *ContractEndDate `xml:"wd:Contract_End_Date,omitempty"`
    SupplierReference *SupplierReference `xml:"wd:Supplier_Reference,omitempty"`
}

type ContractPayRate struct {
    Value *string `xml:",chardata"`
}

type CurrencyReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ContractAssignmentDetails struct {
    Value *string `xml:",chardata"`
}

type ContractEndDate struct {
    Value *string `xml:",chardata"`
}

type SupplierReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WorkerStatusData struct {
    Active *Active `xml:"wd:Active,omitempty"`
    ActiveStatusDate *ActiveStatusDate `xml:"wd:Active_Status_Date,omitempty"`
    HireDate *HireDate `xml:"wd:Hire_Date,omitempty"`
    OriginalHireDate *OriginalHireDate `xml:"wd:Original_Hire_Date,omitempty"`
    HireReasonReference *HireReasonReference `xml:"wd:Hire_Reason_Reference,omitempty"`
    EndEmploymentDate *EndEmploymentDate `xml:"wd:End_Employment_Date,omitempty"`
    ContinuousServiceDate *ContinuousServiceDate `xml:"wd:Continuous_Service_Date,omitempty"`
    FirstDayOfWork *FirstDayOfWork `xml:"wd:First_Day_of_Work,omitempty"`
    ExpectedRetirementDate *ExpectedRetirementDate `xml:"wd:Expected_Retirement_Date,omitempty"`
    RetirementEligibilityDate *RetirementEligibilityDate `xml:"wd:Retirement_Eligibility_Date,omitempty"`
    Retired *Retired `xml:"wd:Retired,omitempty"`
    RetirementDate *RetirementDate `xml:"wd:Retirement_Date,omitempty"`
    SeniorityDate *SeniorityDate `xml:"wd:Seniority_Date,omitempty"`
    SeveranceDate *SeveranceDate `xml:"wd:Severance_Date,omitempty"`
    BenefitsServiceDate *BenefitsServiceDate `xml:"wd:Benefits_Service_Date,omitempty"`
    CompanyServiceDate *CompanyServiceDate `xml:"wd:Company_Service_Date,omitempty"`
    TimeOffServiceDate *TimeOffServiceDate `xml:"wd:Time_Off_Service_Date,omitempty"`
    VestingDate *VestingDate `xml:"wd:Vesting_Date,omitempty"`
    DateEnteredWorkforce *DateEnteredWorkforce `xml:"wd:Date_Entered_Workforce,omitempty"`
    DaysUnemployed *DaysUnemployed `xml:"wd:Days_Unemployed,omitempty"`
    MonthsContinuousPriorEmployment *MonthsContinuousPriorEmployment `xml:"wd:Months_Continuous_Prior_Employment,omitempty"`
    Terminated *Terminated `xml:"wd:Terminated,omitempty"`
    TerminationDate *TerminationDate `xml:"wd:Termination_Date,omitempty"`
    PayThroughDate *PayThroughDate `xml:"wd:Pay_Through_Date,omitempty"`
    AgreementSignatureDate *AgreementSignatureDate `xml:"wd:Agreement_Signature_Date,omitempty"`
    DismissalProcessStartDate *DismissalProcessStartDate `xml:"wd:Dismissal_Process_Start_Date,omitempty"`
    NoticePeriodStartDate *NoticePeriodStartDate `xml:"wd:Notice_Period_Start_Date,omitempty"`
    PrimaryTerminationReasonReference *PrimaryTerminationReasonReference `xml:"wd:Primary_Termination_Reason_Reference,omitempty"`
    PrimaryTerminationCategoryReference *PrimaryTerminationCategoryReference `xml:"wd:Primary_Termination_Category_Reference,omitempty"`
    TerminationInvoluntary *TerminationInvoluntary `xml:"wd:Termination_Involuntary,omitempty"`
    SecondaryTerminationReasonsData *SecondaryTerminationReasonsData `xml:"wd:Secondary_Termination_Reasons_Data,omitempty"`
    LocalTerminationReasonReference *LocalTerminationReasonReference `xml:"wd:Local_Termination_Reason_Reference,omitempty"`
    EligibleForHireReference *EligibleForHireReference `xml:"wd:Eligible_for_Hire_Reference,omitempty"`
    RegrettableTermination *RegrettableTermination `xml:"wd:Regrettable_Termination,omitempty"`
    EligibleForRehireOnLatestTerminationReference *EligibleForRehireOnLatestTerminationReference `xml:"wd:Eligible_for_Rehire_on_Latest_Termination_Reference,omitempty"`
    HireRescinded *HireRescinded `xml:"wd:Hire_Rescinded,omitempty"`
    TerminationLastDayOfWork *TerminationLastDayOfWork `xml:"wd:Termination_Last_Day_of_Work,omitempty"`
    ResignationDate *ResignationDate `xml:"wd:Resignation_Date,omitempty"`
    LastDateForWhichPaid *LastDateForWhichPaid `xml:"wd:Last_Date_for_Which_Paid,omitempty"`
    ExpectedDateOfReturn *ExpectedDateOfReturn `xml:"wd:Expected_Date_of_Return,omitempty"`
    NotReturning *NotReturning `xml:"wd:Not_Returning,omitempty"`
    ReturnUnknown *ReturnUnknown `xml:"wd:Return_Unknown,omitempty"`
    ProbationStartDate *ProbationStartDate `xml:"wd:Probation_Start_Date,omitempty"`
    ProbationEndDate *ProbationEndDate `xml:"wd:Probation_End_Date,omitempty"`
    LeaveStatusData *LeaveStatusData `xml:"wd:Leave_Status_Data,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PrimaryTerminationCategoryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type TerminationInvoluntary struct {
    Value *string `xml:",chardata"`
}

type SecondaryTerminationReasonsData struct {
    SecondaryTerminationReasonReference *SecondaryTerminationReasonReference `xml:"wd:Secondary_Termination_Reason_Reference,omitempty"`
    SecondaryTerminationReasonCategoryReference *SecondaryTerminationReasonCategoryReference `xml:"wd:Secondary_Termination_Reason_Category_Reference,omitempty"`
}

type SecondaryTerminationReasonReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SecondaryTerminationReasonCategoryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type LocalTerminationReasonReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type EligibleForHireReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type RegrettableTermination struct {
    Value *string `xml:",chardata"`
}

type EligibleForRehireOnLatestTerminationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    LeaveRequestEventReference *LeaveRequestEventReference `xml:"wd:Leave_Request_Event_Reference,omitempty"`
    LeaveRequestDescription *LeaveRequestDescription `xml:"wd:Leave_Request_Description,omitempty"`
    LeaveReturnEventReference *LeaveReturnEventReference `xml:"wd:Leave_Return_Event_Reference,omitempty"`
    OnLeave *OnLeave `xml:"wd:On_Leave,omitempty"`
    LeaveStartDate *LeaveStartDate `xml:"wd:Leave_Start_Date,omitempty"`
    EstimatedLeaveEndDate *EstimatedLeaveEndDate `xml:"wd:Estimated_Leave_End_Date,omitempty"`
    LeaveEndDate *LeaveEndDate `xml:"wd:Leave_End_Date,omitempty"`
    FirstDayOfWork *FirstDayOfWork `xml:"wd:First_Day_Of_Work,omitempty"`
    LeaveLastDayOfWork *LeaveLastDayOfWork `xml:"wd:Leave_Last_Day_of_Work,omitempty"`
    LeaveOfAbsenceTypeReference *LeaveOfAbsenceTypeReference `xml:"wd:Leave_of_Absence_Type_Reference,omitempty"`
    BenefitsEffect *BenefitsEffect `xml:"wd:Benefits_Effect,omitempty"`
    PayrollEffect *PayrollEffect `xml:"wd:Payroll_Effect,omitempty"`
    PaidTimeOffAccrualEffect *PaidTimeOffAccrualEffect `xml:"wd:Paid_Time_Off_Accrual_Effect,omitempty"`
    ContinuousServiceAccrualEffect *ContinuousServiceAccrualEffect `xml:"wd:Continuous_Service_Accrual_Effect,omitempty"`
    StockVestingEffect *StockVestingEffect `xml:"wd:Stock_Vesting_Effect,omitempty"`
    LeaveTypeReasonReference *LeaveTypeReasonReference `xml:"wd:Leave_Type_Reason_Reference,omitempty"`
    LeaveRequestAdditionalFields *LeaveRequestAdditionalFields `xml:"wd:Leave_Request_Additional_Fields,omitempty"`
}

type LeaveRequestEventReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type LeaveRequestDescription struct {
    Value *string `xml:",chardata"`
}

type LeaveReturnEventReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type LeaveRequestAdditionalFields struct {
    LastDateForWhichPaid *LastDateForWhichPaid `xml:"wd:Last_Date_for_Which_Paid,omitempty"`
    ExpectedDueDate *ExpectedDueDate `xml:"wd:Expected_Due_Date,omitempty"`
    ChildSBirthDate *ChildSBirthDate `xml:"wd:Child_s_Birth_Date,omitempty"`
    StillbirthBabyDeceased *StillbirthBabyDeceased `xml:"wd:Stillbirth_Baby_Deceased,omitempty"`
    DateBabyArrivedHomeFromHospital *DateBabyArrivedHomeFromHospital `xml:"wd:Date_Baby_Arrived_Home_From_Hospital,omitempty"`
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

type GetApplicantsRequest struct {
    XMLName xml.Name `xml:"wd:Get_Applicants_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    RequestReferences *RequestReferences `xml:"wd:Request_References,omitempty"`
    RequestCriteria *RequestCriteria `xml:"wd:Request_Criteria,omitempty"`
    ResponseFilter *ResponseFilter `xml:"wd:Response_Filter,omitempty"`
    ResponseGroup *ResponseGroup `xml:"wd:Response_Group,omitempty"`
}

type ApplicantReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type FormerWorkerReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type IncludeRecruitingInformation struct {
    Value *string `xml:",chardata"`
}

type IncludeQualificationProfile struct {
    Value *string `xml:",chardata"`
}

type IncludeResume struct {
    Value *string `xml:",chardata"`
}

type IncludeBackgroundCheck struct {
    Value *string `xml:",chardata"`
}

type IncludeExternalIntegrationIdData struct {
    Value *string `xml:",chardata"`
}

type ChangeHomeContactInformationRequest struct {
    XMLName xml.Name `xml:"wd:Change_Home_Contact_Information_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    BusinessProcessParameters *BusinessProcessParameters `xml:"wd:Business_Process_Parameters,omitempty"`
    ChangeHomeContactInformationData *ChangeHomeContactInformationData `xml:"wd:Change_Home_Contact_Information_Data,omitempty"`
}

type BusinessProcessParameters struct {
    AutoComplete *AutoComplete `xml:"wd:Auto_Complete,omitempty"`
    RunNow *RunNow `xml:"wd:Run_Now,omitempty"`
    DiscardOnExitValidationError *DiscardOnExitValidationError `xml:"wd:Discard_On_Exit_Validation_Error,omitempty"`
    CommentData *CommentData `xml:"wd:Comment_Data,omitempty"`
    BusinessProcessAttachmentData *BusinessProcessAttachmentData `xml:"wd:Business_Process_Attachment_Data,omitempty"`
}

type AutoComplete struct {
    Value *string `xml:",chardata"`
}

type RunNow struct {
    Value *string `xml:",chardata"`
}

type DiscardOnExitValidationError struct {
    Value *string `xml:",chardata"`
}

type CommentData struct {
    Comment *Comment `xml:"wd:Comment,omitempty"`
    WorkerReference *WorkerReference `xml:"wd:Worker_Reference,omitempty"`
}

type Comment struct {
    Value *string `xml:",chardata"`
}

type BusinessProcessAttachmentData struct {
    FileName *FileName `xml:"wd:File_Name,omitempty"`
    EventAttachmentDescription *EventAttachmentDescription `xml:"wd:Event_Attachment_Description,omitempty"`
    EventAttachmentCategoryReference *EventAttachmentCategoryReference `xml:"wd:Event_Attachment_Category_Reference,omitempty"`
    File *File `xml:"wd:File,omitempty"`
    ContentType *ContentType `xml:"wd:Content_Type,omitempty"`
}

type FileName struct {
    Value *string `xml:",chardata"`
}

type EventAttachmentDescription struct {
    Value *string `xml:",chardata"`
}

type EventAttachmentCategoryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type File struct {
    Value *string `xml:",chardata"`
}

type ContentType struct {
    Value *string `xml:",chardata"`
}

type ChangeHomeContactInformationData struct {
    PersonReference *PersonReference `xml:"wd:Person_Reference,omitempty"`
    UniversalIdReference *UniversalIdReference `xml:"wd:Universal_ID_Reference,omitempty"`
    EventEffectiveDate *EventEffectiveDate `xml:"wd:Event_Effective_Date,omitempty"`
    PersonContactInformationData *PersonContactInformationData `xml:"wd:Person_Contact_Information_Data,omitempty"`
}

type PersonReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type EventEffectiveDate struct {
    Value *string `xml:",chardata"`
}

type PersonContactInformationData struct {
    PersonAddressInformationData *PersonAddressInformationData `xml:"wd:Person_Address_Information_Data,omitempty"`
    PersonPhoneInformationData *PersonPhoneInformationData `xml:"wd:Person_Phone_Information_Data,omitempty"`
    PersonEmailInformationData *PersonEmailInformationData `xml:"wd:Person_Email_Information_Data,omitempty"`
    PersonInstantMessengerInformationData *PersonInstantMessengerInformationData `xml:"wd:Person_Instant_Messenger_Information_Data,omitempty"`
    PersonWebAddressInformationData *PersonWebAddressInformationData `xml:"wd:Person_Web_Address_Information_Data,omitempty"`
}

type PersonAddressInformationData struct {
    ReplaceAll *string `xml:"wd:Replace_All,attr,omitempty"`
    AddressInformationData *AddressInformationData `xml:"wd:Address_Information_Data,omitempty"`
}

type AddressInformationData struct {
    AddressFormatType *string `xml:"wd:Address_Format_Type,attr,omitempty"`
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    EffectiveDate *string `xml:"wd:Effective_Date,attr,omitempty"`
    AddressData *AddressData `xml:"wd:Address_Data,omitempty"`
    UsageData *UsageData `xml:"wd:Usage_Data,omitempty"`
    NumberOfDays *NumberOfDays `xml:"wd:Number_of_Days,omitempty"`
    DaysOfTheWeekReference *DaysOfTheWeekReference `xml:"wd:Days_of_the_Week_Reference,omitempty"`
    AddressReference *AddressReference `xml:"wd:Address_Reference,omitempty"`
    AddressId *AddressId `xml:"wd:Address_ID,omitempty"`
}

type DaysOfTheWeekReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PersonPhoneInformationData struct {
    ReplaceAll *string `xml:"wd:Replace_All,attr,omitempty"`
    PhoneInformationData *PhoneInformationData `xml:"wd:Phone_Information_Data,omitempty"`
}

type PhoneInformationData struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    PhoneData *PhoneData `xml:"wd:Phone_Data,omitempty"`
    UsageData *UsageData `xml:"wd:Usage_Data,omitempty"`
    PhoneReference *PhoneReference `xml:"wd:Phone_Reference,omitempty"`
    PhoneId *PhoneId `xml:"wd:Phone_ID,omitempty"`
}

type DeviceTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CountryCodeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CompletePhoneNumber struct {
    Value *string `xml:",chardata"`
}

type Extension struct {
    Value *string `xml:",chardata"`
}

type PhoneId struct {
    Value *string `xml:",chardata"`
}

type PersonEmailInformationData struct {
    ReplaceAll *string `xml:"wd:Replace_All,attr,omitempty"`
    EmailInformationData *EmailInformationData `xml:"wd:Email_Information_Data,omitempty"`
}

type EmailInformationData struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    EmailData *EmailData `xml:"wd:Email_Data,omitempty"`
    UsageData *UsageData `xml:"wd:Usage_Data,omitempty"`
    EmailReference *EmailReference `xml:"wd:Email_Reference,omitempty"`
    EmailId *EmailId `xml:"wd:Email_ID,omitempty"`
}

type EmailData struct {
    EmailAddress *EmailAddress `xml:"wd:Email_Address,omitempty"`
    EmailComment *EmailComment `xml:"wd:Email_Comment,omitempty"`
}

type EmailId struct {
    Value *string `xml:",chardata"`
}

type PersonInstantMessengerInformationData struct {
    ReplaceAll *string `xml:"wd:Replace_All,attr,omitempty"`
    PersonInstantMessengerData *PersonInstantMessengerData `xml:"wd:Person_Instant_Messenger_Data,omitempty"`
}

type PersonInstantMessengerData struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    InstantMessengerData *InstantMessengerData `xml:"wd:Instant_Messenger_Data,omitempty"`
    UsageData *UsageData `xml:"wd:Usage_Data,omitempty"`
    InstantMessengerReference *InstantMessengerReference `xml:"wd:Instant_Messenger_Reference,omitempty"`
    InstantMessengerId *InstantMessengerId `xml:"wd:Instant_Messenger_ID,omitempty"`
}

type InstantMessengerId struct {
    Value *string `xml:",chardata"`
}

type PersonWebAddressInformationData struct {
    ReplaceAll *string `xml:"wd:Replace_All,attr,omitempty"`
    PersonWebAddressData *PersonWebAddressData `xml:"wd:Person_Web_Address_Data,omitempty"`
}

type PersonWebAddressData struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    WebAddressData *WebAddressData `xml:"wd:Web_Address_Data,omitempty"`
    UsageData *UsageData `xml:"wd:Usage_Data,omitempty"`
    WebAddressReference *WebAddressReference `xml:"wd:Web_Address_Reference,omitempty"`
    WebAddressId *WebAddressId `xml:"wd:Web_Address_ID,omitempty"`
}

type WebAddressId struct {
    Value *string `xml:",chardata"`
}

type GetWorkersRequest struct {
    XMLName xml.Name `xml:"wd:Get_Workers_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    RequestReferences *RequestReferences `xml:"wd:Request_References,omitempty"`
    RequestCriteria *RequestCriteria `xml:"wd:Request_Criteria,omitempty"`
    ResponseFilter *ResponseFilter `xml:"wd:Response_Filter,omitempty"`
    ResponseGroup *ResponseGroup `xml:"wd:Response_Group,omitempty"`
}

type PutApplicantResponse struct {
    XMLName xml.Name `xml:"wd:Put_Applicant_Response"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    ApplicantReference *ApplicantReference `xml:"wd:Applicant_Reference,omitempty"`
    ExceptionsResponseData *ExceptionsResponseData `xml:"wd:Exceptions_Response_Data,omitempty"`
}

type ExceptionsResponseData struct {
    ExceptionsData *ExceptionsData `xml:"wd:Exceptions_Data,omitempty"`
}

type ExceptionsData struct {
    ExceptionData *ExceptionData `xml:"wd:Exception_Data,omitempty"`
}

type ExceptionData struct {
    Classification *Classification `xml:"wd:Classification,omitempty"`
    Message *Message `xml:"wd:Message,omitempty"`
}

type Classification struct {
    Value *string `xml:",chardata"`
}

type Message struct {
    Value *string `xml:",chardata"`
}

type ChangeHomeContactInformationResponse struct {
    XMLName xml.Name `xml:"wd:Change_Home_Contact_Information_Response"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    ChangeHomeContactInformationEventReference *ChangeHomeContactInformationEventReference `xml:"wd:Change_Home_Contact_Information_Event_Reference,omitempty"`
}

type ChangeHomeContactInformationEventReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PutApplicantRequest struct {
    XMLName xml.Name `xml:"wd:Put_Applicant_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    AddOnly *string `xml:"wd:Add_Only,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    ApplicantReference *ApplicantReference `xml:"wd:Applicant_Reference,omitempty"`
    ApplicantData *ApplicantData `xml:"wd:Applicant_Data,omitempty"`
}

type ApplicantData struct {
    ApplicantId *ApplicantId `xml:"wd:Applicant_ID,omitempty"`
    PersonalData *PersonalData `xml:"wd:Personal_Data,omitempty"`
    QualificationData *QualificationData `xml:"wd:Qualification_Data,omitempty"`
    RecruitingData *RecruitingData `xml:"wd:Recruiting_Data,omitempty"`
    ResumeData *ResumeData `xml:"wd:Resume_Data,omitempty"`
    BackgroundCheckData *BackgroundCheckData `xml:"wd:Background_Check_Data,omitempty"`
    ExternalIntegrationIdData *ExternalIntegrationIdData `xml:"wd:External_Integration_ID_Data,omitempty"`
    DocumentFieldResultData *DocumentFieldResultData `xml:"wd:Document_Field_Result_Data,omitempty"`
}

type ApplicantId struct {
    Value *string `xml:",chardata"`
}

type QualificationData struct {
    ExternalJobHistory *ExternalJobHistory `xml:"wd:External_Job_History,omitempty"`
    Competency *Competency `xml:"wd:Competency,omitempty"`
    Certification *Certification `xml:"wd:Certification,omitempty"`
    Training *Training `xml:"wd:Training,omitempty"`
    AwardAndActivity *AwardAndActivity `xml:"wd:Award_and_Activity,omitempty"`
    OrganizationMembership *OrganizationMembership `xml:"wd:Organization_Membership,omitempty"`
    Education *Education `xml:"wd:Education,omitempty"`
    WorkExperience *WorkExperience `xml:"wd:Work_Experience,omitempty"`
    Language *Language `xml:"wd:Language,omitempty"`
    InternalProjectExperience *InternalProjectExperience `xml:"wd:Internal_Project_Experience,omitempty"`
}

type ExternalJobHistory struct {
    JobHistoryReference *JobHistoryReference `xml:"wd:Job_History_Reference,omitempty"`
    JobHistoryData *JobHistoryData `xml:"wd:Job_History_Data,omitempty"`
}

type JobHistoryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type JobHistoryData struct {
    JobHistoryId *JobHistoryId `xml:"wd:Job_History_ID,omitempty"`
    RemoveJobHistory *RemoveJobHistory `xml:"wd:Remove_Job_History,omitempty"`
    JobTitle *JobTitle `xml:"wd:Job_Title,omitempty"`
    Company *Company `xml:"wd:Company,omitempty"`
    JobHistoryCompanyReference *JobHistoryCompanyReference `xml:"wd:Job_History_Company_Reference,omitempty"`
    StartDate *StartDate `xml:"wd:Start_Date,omitempty"`
    EndDate *EndDate `xml:"wd:End_Date,omitempty"`
    ResponsibilitiesAndAchievements *ResponsibilitiesAndAchievements `xml:"wd:Responsibilities_and_Achievements,omitempty"`
    Location *Location `xml:"wd:Location,omitempty"`
    JobReference *JobReference `xml:"wd:Job_Reference,omitempty"`
    Contact *Contact `xml:"wd:Contact,omitempty"`
}

type JobHistoryId struct {
    Value *string `xml:",chardata"`
}

type RemoveJobHistory struct {
    Value *string `xml:",chardata"`
}

type JobTitle struct {
    Value *string `xml:",chardata"`
}

type Company struct {
    Value *string `xml:",chardata"`
}

type JobHistoryCompanyReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ResponsibilitiesAndAchievements struct {
    Value *string `xml:",chardata"`
}

type Location struct {
    Value *string `xml:",chardata"`
}

type JobReference struct {
    Value *string `xml:",chardata"`
}

type Contact struct {
    Value *string `xml:",chardata"`
}

type Competency struct {
    CompetencyLevelBehaviorReference *CompetencyLevelBehaviorReference `xml:"wd:Competency_Level_Behavior_Reference,omitempty"`
    CompetencyLevelBehaviorDescriptor *CompetencyLevelBehaviorDescriptor `xml:"wd:Competency_Level_Behavior_Descriptor,omitempty"`
    AssessmentComment *AssessmentComment `xml:"wd:Assessment_Comment,omitempty"`
    AssessedOn *AssessedOn `xml:"wd:Assessed_On,omitempty"`
    AssessedByPersonReference *AssessedByPersonReference `xml:"wd:Assessed_By_Person_Reference,omitempty"`
    AssessedByWorkerDescriptor *AssessedByWorkerDescriptor `xml:"wd:Assessed_By_Worker_Descriptor,omitempty"`
    CompetencyReference *CompetencyReference `xml:"wd:Competency_Reference,omitempty"`
    CompetencyDescriptor *CompetencyDescriptor `xml:"wd:Competency_Descriptor,omitempty"`
}

type CompetencyLevelBehaviorReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CompetencyLevelBehaviorDescriptor struct {
    Value *string `xml:",chardata"`
}

type AssessmentComment struct {
    Value *string `xml:",chardata"`
}

type AssessedOn struct {
    Value *string `xml:",chardata"`
}

type AssessedByPersonReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AssessedByWorkerDescriptor struct {
    Value *string `xml:",chardata"`
}

type CompetencyReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CompetencyDescriptor struct {
    Value *string `xml:",chardata"`
}

type Certification struct {
    CertificationReference *CertificationReference `xml:"wd:Certification_Reference,omitempty"`
    CertificationData *CertificationData `xml:"wd:Certification_Data,omitempty"`
}

type CertificationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CertificationData struct {
    CertificationId *CertificationId `xml:"wd:Certification_ID,omitempty"`
    RemoveCertification *RemoveCertification `xml:"wd:Remove_Certification,omitempty"`
    CertificationReference *CertificationReference `xml:"wd:Certification_Reference,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    CertificationName *CertificationName `xml:"wd:Certification_Name,omitempty"`
    Issuer *Issuer `xml:"wd:Issuer,omitempty"`
    CertificationNumber *CertificationNumber `xml:"wd:Certification_Number,omitempty"`
    IssuedDate *IssuedDate `xml:"wd:Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"wd:Expiration_Date,omitempty"`
    ExaminationScore *ExaminationScore `xml:"wd:Examination_Score,omitempty"`
    ExaminationDate *ExaminationDate `xml:"wd:Examination_Date,omitempty"`
    SpecialtyAchievementData *SpecialtyAchievementData `xml:"wd:Specialty_Achievement_Data,omitempty"`
    WorkerDocumentData *WorkerDocumentData `xml:"wd:Worker_Document_Data,omitempty"`
    CertificationSkillReference *CertificationSkillReference `xml:"wd:Certification_Skill_Reference,omitempty"`
    CertificationDataDetails *CertificationDataDetails `xml:"wd:Certification_Data_Details,omitempty"`
}

type CertificationId struct {
    Value *string `xml:",chardata"`
}

type RemoveCertification struct {
    Value *string `xml:",chardata"`
}

type CertificationName struct {
    Value *string `xml:",chardata"`
}

type Issuer struct {
    Value *string `xml:",chardata"`
}

type CertificationNumber struct {
    Value *string `xml:",chardata"`
}

type ExaminationScore struct {
    Value *string `xml:",chardata"`
}

type ExaminationDate struct {
    Value *string `xml:",chardata"`
}

type SpecialtyAchievementData struct {
    SpecialtyReference *SpecialtyReference `xml:"wd:Specialty_Reference,omitempty"`
    StartDate *StartDate `xml:"wd:Start_Date,omitempty"`
    EndDate *EndDate `xml:"wd:End_Date,omitempty"`
    SubspecialtyReference *SubspecialtyReference `xml:"wd:Subspecialty_Reference,omitempty"`
}

type SpecialtyReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SubspecialtyReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WorkerDocumentData struct {
    FileName *FileName `xml:"wd:File_Name,omitempty"`
    Comment *Comment `xml:"wd:Comment,omitempty"`
    File *File `xml:"wd:File,omitempty"`
    DocumentCategoryReference *DocumentCategoryReference `xml:"wd:Document_Category_Reference,omitempty"`
    ContentType *ContentType `xml:"wd:Content_Type,omitempty"`
}

type DocumentCategoryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Training struct {
    TrainingReference *TrainingReference `xml:"wd:Training_Reference,omitempty"`
    TrainingData *TrainingData `xml:"wd:Training_Data,omitempty"`
    Value *string `xml:",chardata"`
}

type TrainingReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type TrainingData struct {
    TrainingId *TrainingId `xml:"wd:Training_ID,omitempty"`
    RemoveTraining *RemoveTraining `xml:"wd:Remove_Training,omitempty"`
    Training *Training `xml:"wd:Training,omitempty"`
    Description *Description `xml:"wd:Description,omitempty"`
    TrainingTypeReference *TrainingTypeReference `xml:"wd:Training_Type_Reference,omitempty"`
    CompletionDate *CompletionDate `xml:"wd:Completion_Date,omitempty"`
    TrainingDuration *TrainingDuration `xml:"wd:Training_Duration,omitempty"`
}

type TrainingId struct {
    Value *string `xml:",chardata"`
}

type RemoveTraining struct {
    Value *string `xml:",chardata"`
}

type Description struct {
    Value *string `xml:",chardata"`
}

type TrainingTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CompletionDate struct {
    Value *string `xml:",chardata"`
}

type TrainingDuration struct {
    Value *string `xml:",chardata"`
}

type AwardAndActivity struct {
    AwardAndActivityReference *AwardAndActivityReference `xml:"wd:Award_and_Activity_Reference,omitempty"`
    AwardAndActivityData *AwardAndActivityData `xml:"wd:Award_and_Activity_Data,omitempty"`
}

type AwardAndActivityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AwardAndActivityData struct {
    AwardAndActivityId *AwardAndActivityId `xml:"wd:Award_and_Activity_ID,omitempty"`
    RemoveAwardAndActivity *RemoveAwardAndActivity `xml:"wd:Remove_Award_and_Activity,omitempty"`
    AwardAndActivityTypeReference *AwardAndActivityTypeReference `xml:"wd:Award_and_Activity_Type_Reference,omitempty"`
    Title *Title `xml:"wd:Title,omitempty"`
    SponsorIssuer *SponsorIssuer `xml:"wd:Sponsor_Issuer,omitempty"`
    StartDate *StartDate `xml:"wd:Start_Date,omitempty"`
    EndDate *EndDate `xml:"wd:End_Date,omitempty"`
    Description *Description `xml:"wd:Description,omitempty"`
    Url *Url `xml:"wd:URL,omitempty"`
    RelatedPositionReference *RelatedPositionReference `xml:"wd:Related_Position_Reference,omitempty"`
}

type AwardAndActivityId struct {
    Value *string `xml:",chardata"`
}

type RemoveAwardAndActivity struct {
    Value *string `xml:",chardata"`
}

type AwardAndActivityTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Title struct {
    Value *string `xml:",chardata"`
}

type SponsorIssuer struct {
    Value *string `xml:",chardata"`
}

type Url struct {
    Value *string `xml:",chardata"`
}

type RelatedPositionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type OrganizationMembership struct {
    OrganizationProfessionalAffiliationReference *OrganizationProfessionalAffiliationReference `xml:"wd:Organization_Professional_Affiliation_Reference,omitempty"`
    OrganizationProfessionalAffiliationData *OrganizationProfessionalAffiliationData `xml:"wd:Organization_Professional_Affiliation_Data,omitempty"`
}

type OrganizationProfessionalAffiliationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type OrganizationProfessionalAffiliationData struct {
    ProfessionalAffiliationId *ProfessionalAffiliationId `xml:"wd:Professional_Affiliation_ID,omitempty"`
    RemoveProfessionalAffiliation *RemoveProfessionalAffiliation `xml:"wd:Remove_Professional_Affiliation,omitempty"`
    ProfessionalAffiliationReference *ProfessionalAffiliationReference `xml:"wd:Professional_Affiliation_Reference,omitempty"`
    ProfessionalAffiliation *ProfessionalAffiliation `xml:"wd:Professional_Affiliation,omitempty"`
    ProfessionalAffiliationTypeReference *ProfessionalAffiliationTypeReference `xml:"wd:Professional_Affiliation_Type_Reference,omitempty"`
    Affiliation *Affiliation `xml:"wd:Affiliation,omitempty"`
    ProfessionalAffiliationRelationshipTypeReference *ProfessionalAffiliationRelationshipTypeReference `xml:"wd:Professional_Affiliation_Relationship_Type_Reference,omitempty"`
    BeginDate *BeginDate `xml:"wd:Begin_Date,omitempty"`
    EndDate *EndDate `xml:"wd:End_Date,omitempty"`
    ContactInformationData *ContactInformationData `xml:"wd:Contact_Information_Data,omitempty"`
}

type ProfessionalAffiliationId struct {
    Value *string `xml:",chardata"`
}

type RemoveProfessionalAffiliation struct {
    Value *string `xml:",chardata"`
}

type ProfessionalAffiliationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ProfessionalAffiliation struct {
    Value *string `xml:",chardata"`
}

type ProfessionalAffiliationTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Affiliation struct {
    Value *string `xml:",chardata"`
}

type ProfessionalAffiliationRelationshipTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type BeginDate struct {
    Value *string `xml:",chardata"`
}

type ContactInformationData struct {
    AddressData *AddressData `xml:"wd:Address_Data,omitempty"`
    PhoneData *PhoneData `xml:"wd:Phone_Data,omitempty"`
    EmailAddressData *EmailAddressData `xml:"wd:Email_Address_Data,omitempty"`
    InstantMessengerData *InstantMessengerData `xml:"wd:Instant_Messenger_Data,omitempty"`
    WebAddressData *WebAddressData `xml:"wd:Web_Address_Data,omitempty"`
}

type Education struct {
    EducationReference *EducationReference `xml:"wd:Education_Reference,omitempty"`
    EducationData *EducationData `xml:"wd:Education_Data,omitempty"`
}

type EducationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type EducationData struct {
    EducationId *EducationId `xml:"wd:Education_ID,omitempty"`
    RemoveEducation *RemoveEducation `xml:"wd:Remove_Education,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    SchoolReference *SchoolReference `xml:"wd:School_Reference,omitempty"`
    SchoolName *SchoolName `xml:"wd:School_Name,omitempty"`
    SchoolTypeReference *SchoolTypeReference `xml:"wd:School_Type_Reference,omitempty"`
    Location *Location `xml:"wd:Location,omitempty"`
    DegreeReference *DegreeReference `xml:"wd:Degree_Reference,omitempty"`
    DegreeCompletedReference *DegreeCompletedReference `xml:"wd:Degree_Completed_Reference,omitempty"`
    DateDegreeReceived *DateDegreeReceived `xml:"wd:Date_Degree_Received,omitempty"`
    FieldOfStudyReference *FieldOfStudyReference `xml:"wd:Field_of_Study_Reference,omitempty"`
    GradeAverage *GradeAverage `xml:"wd:Grade_Average,omitempty"`
    FirstYearAttended *FirstYearAttended `xml:"wd:First_Year_Attended,omitempty"`
    LastYearAttended *LastYearAttended `xml:"wd:Last_Year_Attended,omitempty"`
    IsHighestLevelOfEducation *IsHighestLevelOfEducation `xml:"wd:Is_Highest_Level_of_Education,omitempty"`
    FirstDayAttended *FirstDayAttended `xml:"wd:First_Day_Attended,omitempty"`
    LastDayAttended *LastDayAttended `xml:"wd:Last_Day_Attended,omitempty"`
    EducationAttachmentData *EducationAttachmentData `xml:"wd:Education_Attachment_Data,omitempty"`
}

type EducationId struct {
    Value *string `xml:",chardata"`
}

type RemoveEducation struct {
    Value *string `xml:",chardata"`
}

type SchoolReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SchoolName struct {
    Value *string `xml:",chardata"`
}

type SchoolTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DegreeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DegreeCompletedReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DateDegreeReceived struct {
    Value *string `xml:",chardata"`
}

type FieldOfStudyReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type GradeAverage struct {
    Value *string `xml:",chardata"`
}

type FirstYearAttended struct {
    Value *string `xml:",chardata"`
}

type LastYearAttended struct {
    Value *string `xml:",chardata"`
}

type IsHighestLevelOfEducation struct {
    Value *string `xml:",chardata"`
}

type FirstDayAttended struct {
    Value *string `xml:",chardata"`
}

type LastDayAttended struct {
    Value *string `xml:",chardata"`
}

type EducationAttachmentData struct {
    FileName *FileName `xml:"wd:File_Name,omitempty"`
    Comment *Comment `xml:"wd:Comment,omitempty"`
    File *File `xml:"wd:File,omitempty"`
    DocumentCategoryReference *DocumentCategoryReference `xml:"wd:Document_Category_Reference,omitempty"`
}

type WorkExperience struct {
    ExperienceReference *ExperienceReference `xml:"wd:Experience_Reference,omitempty"`
    RemoveExperience *RemoveExperience `xml:"wd:Remove_Experience,omitempty"`
    ExperienceRatingReference *ExperienceRatingReference `xml:"wd:Experience_Rating_Reference,omitempty"`
    ExperienceComment *ExperienceComment `xml:"wd:Experience_Comment,omitempty"`
}

type ExperienceReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type RemoveExperience struct {
    Value *string `xml:",chardata"`
}

type ExperienceRatingReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ExperienceComment struct {
    Value *string `xml:",chardata"`
}

type Language struct {
    LanguageReference *LanguageReference `xml:"wd:Language_Reference,omitempty"`
    RemoveLanguage *RemoveLanguage `xml:"wd:Remove_Language,omitempty"`
    NativeLanguage *NativeLanguage `xml:"wd:Native_Language,omitempty"`
    LanguageAbility *LanguageAbility `xml:"wd:Language_Ability,omitempty"`
    AssessedOn *AssessedOn `xml:"wd:Assessed_On,omitempty"`
    Note *Note `xml:"wd:Note,omitempty"`
    AssessedByWorkerReference *AssessedByWorkerReference `xml:"wd:Assessed_by_Worker_Reference,omitempty"`
    Native *Native `xml:"wd:Native,omitempty"`
}

type LanguageReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type RemoveLanguage struct {
    Value *string `xml:",chardata"`
}

type NativeLanguage struct {
    Value *string `xml:",chardata"`
}

type LanguageAbility struct {
    LanguageAbilityData *LanguageAbilityData `xml:"wd:Language_Ability_Data,omitempty"`
}

type LanguageAbilityData struct {
    LanguageProficiencyReference *LanguageProficiencyReference `xml:"wd:Language_Proficiency_Reference,omitempty"`
    LanguageAbilityTypeReference *LanguageAbilityTypeReference `xml:"wd:Language_Ability_Type_Reference,omitempty"`
}

type LanguageProficiencyReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type LanguageAbilityTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AssessedByWorkerReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type InternalProjectExperience struct {
    InternalProjectExperienceReference *InternalProjectExperienceReference `xml:"wd:Internal_Project_Experience_Reference,omitempty"`
    InternalProjectExperienceData *InternalProjectExperienceData `xml:"wd:Internal_Project_Experience_Data,omitempty"`
}

type InternalProjectExperienceReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type InternalProjectExperienceData struct {
    InternalProjectExperienceId *InternalProjectExperienceId `xml:"wd:Internal_Project_Experience_ID,omitempty"`
    RemoveInternalProjectExperience *RemoveInternalProjectExperience `xml:"wd:Remove_Internal_Project_Experience,omitempty"`
    InternalProject *InternalProject `xml:"wd:Internal_Project,omitempty"`
    Description *Description `xml:"wd:Description,omitempty"`
    StartDate *StartDate `xml:"wd:Start_Date,omitempty"`
    EndDate *EndDate `xml:"wd:End_Date,omitempty"`
    ProjectLeader *ProjectLeader `xml:"wd:Project_Leader,omitempty"`
}

type InternalProjectExperienceId struct {
    Value *string `xml:",chardata"`
}

type RemoveInternalProjectExperience struct {
    Value *string `xml:",chardata"`
}

type InternalProject struct {
    Value *string `xml:",chardata"`
}

type ProjectLeader struct {
    Value *string `xml:",chardata"`
}

type RecruitingData struct {
    ApplicantEnteredDate *ApplicantEnteredDate `xml:"wd:Applicant_Entered_Date,omitempty"`
    ApplicantComments *ApplicantComments `xml:"wd:Applicant_Comments,omitempty"`
    EligibleForHireReference *EligibleForHireReference `xml:"wd:Eligible_For_Hire_Reference,omitempty"`
    EligibleForRehireComments *EligibleForRehireComments `xml:"wd:Eligible_for_Rehire_Comments,omitempty"`
    ApplicantHasMarkedAsNoShowReference *ApplicantHasMarkedAsNoShowReference `xml:"wd:Applicant_Has_Marked_as_No_Show_Reference,omitempty"`
    ApplicantSourceReference *ApplicantSourceReference `xml:"wd:Applicant_Source_Reference,omitempty"`
    ReferredByWorkerReference *ReferredByWorkerReference `xml:"wd:Referred_by_Worker_Reference,omitempty"`
    PositionsConsideredForReference *PositionsConsideredForReference `xml:"wd:Positions_Considered_for_Reference,omitempty"`
}

type ApplicantEnteredDate struct {
    Value *string `xml:",chardata"`
}

type ApplicantComments struct {
    Value *string `xml:",chardata"`
}

type EligibleForRehireComments struct {
    Value *string `xml:",chardata"`
}

type ApplicantHasMarkedAsNoShowReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ApplicantSourceReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ReferredByWorkerReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PositionsConsideredForReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ResumeData struct {
    Resume *Resume `xml:"wd:Resume,omitempty"`
    FileId *FileId `xml:"wd:File_ID,omitempty"`
    File *File `xml:"wd:File,omitempty"`
    Filename *Filename `xml:"wd:FileName,omitempty"`
    Comments *Comments `xml:"wd:Comments,omitempty"`
    Summary *Summary `xml:"wd:Summary,omitempty"`
    SkillData *SkillData `xml:"wd:Skill_Data,omitempty"`
    ExperienceData *ExperienceData `xml:"wd:Experience_Data,omitempty"`
    EducationData *EducationData `xml:"wd:Education_Data,omitempty"`
    CertificationData *CertificationData `xml:"wd:Certification_Data,omitempty"`
    LanguageData *LanguageData `xml:"wd:Language_Data,omitempty"`
    QuestionnaireResponseData *QuestionnaireResponseData `xml:"wd:Questionnaire_Response_Data,omitempty"`
}

type Resume struct {
    ResumeReference *ResumeReference `xml:"wd:Resume_Reference,omitempty"`
    ResumeData *ResumeData `xml:"wd:Resume_Data,omitempty"`
}

type ResumeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type FileId struct {
    Value *string `xml:",chardata"`
}

type Filename struct {
    Value *string `xml:",chardata"`
}

type BackgroundCheckData struct {
    StatusDate *StatusDate `xml:"wd:Status_Date,omitempty"`
    StatusReference *StatusReference `xml:"wd:Status_Reference,omitempty"`
    StatusComment *StatusComment `xml:"wd:Status_Comment,omitempty"`
}

type StatusDate struct {
    Value *string `xml:",chardata"`
}

type StatusComment struct {
    Value *string `xml:",chardata"`
}

type ExternalIntegrationIdData struct {
    Id *Id `xml:"wd:ID,omitempty"`
}

type DocumentFieldResultData struct {
    FieldReference *FieldReference `xml:"wd:Field_Reference,omitempty"`
    Value *Value `xml:"wd:Value,omitempty"`
}

type Value struct {
    Value *string `xml:",chardata"`
}

type ChangePersonalInformationResponse struct {
    XMLName xml.Name `xml:"wd:Change_Personal_Information_Response"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    PersonalInformationChangeEventReference *PersonalInformationChangeEventReference `xml:"wd:Personal_Information_Change_Event_Reference,omitempty"`
    PersonalInformationData *PersonalInformationData `xml:"wd:Personal_Information_Data,omitempty"`
}

type PersonalInformationChangeEventReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DateOfBirth struct {
    Value *string `xml:",chardata"`
}

type BirthCountryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type BirthRegionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DisabilityInformationData struct {
    ReplaceAll *string `xml:"wd:Replace_All,attr,omitempty"`
    DisabilityStatusInformationData *DisabilityStatusInformationData `xml:"wd:Disability_Status_Information_Data,omitempty"`
}

type DisabilityStatusInformationData struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    DisabilityStatusReference *DisabilityStatusReference `xml:"wd:Disability_Status_Reference,omitempty"`
    DisabilityStatusData *DisabilityStatusData `xml:"wd:Disability_Status_Data,omitempty"`
}

type CitizenshipReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AdditionalNationalityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type VisualSurveyEthnicityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type HispanicOrLatinoForVisualSurvey struct {
    Value *string `xml:",chardata"`
}

type PersonnelFileAgency struct {
    Value *string `xml:",chardata"`
}

type MilitaryInformationData struct {
    ReplaceAll *string `xml:"wd:Replace_All,attr,omitempty"`
    MilitaryServiceInformationData *MilitaryServiceInformationData `xml:"wd:Military_Service_Information_Data,omitempty"`
}

type MilitaryServiceInformationData struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    MilitaryServiceReference *MilitaryServiceReference `xml:"wd:Military_Service_Reference,omitempty"`
    MilitaryServiceData *MilitaryServiceData `xml:"wd:Military_Service_Data,omitempty"`
}

type MilitaryStatusReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type MilitaryDischargeDate struct {
    Value *string `xml:",chardata"`
}

type MilitaryStatusBeginDate struct {
    Value *string `xml:",chardata"`
}

type UsesTobacco struct {
    Value *string `xml:",chardata"`
}

type LgbtIdentificationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type RelativeNameData struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    RelativeNameReferenceReference *RelativeNameReferenceReference `xml:"wd:Relative_Name_Reference_Reference,omitempty"`
    RelativeTypeReference *RelativeTypeReference `xml:"wd:Relative_Type_Reference,omitempty"`
    NameDetailData *NameDetailData `xml:"wd:Name_Detail_Data,omitempty"`
}

type ChangeLegalNameResponse struct {
    XMLName xml.Name `xml:"wd:Change_Legal_Name_Response"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    LegalNameChangeEventReference *LegalNameChangeEventReference `xml:"wd:Legal_Name_Change_Event_Reference,omitempty"`
}

type LegalNameChangeEventReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PutCandidateResponse struct {
    XMLName xml.Name `xml:"wd:Put_Candidate_Response"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    CandidateReference *CandidateReference `xml:"wd:Candidate_Reference,omitempty"`
    CandidateJobApplicationData *CandidateJobApplicationData `xml:"wd:Candidate_Job_Application_Data,omitempty"`
}

type CandidateReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CandidateJobApplicationData struct {
    JobApplicationReference *JobApplicationReference `xml:"wd:Job_Application_Reference,omitempty"`
    JobRequisitionReference *JobRequisitionReference `xml:"wd:Job_Requisition_Reference,omitempty"`
}

type JobApplicationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type JobRequisitionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type GetCandidatesResponse struct {
    XMLName xml.Name `xml:"wd:Get_Candidates_Response"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    RequestReferences *RequestReferences `xml:"wd:Request_References,omitempty"`
    RequestCriteria *RequestCriteria `xml:"wd:Request_Criteria,omitempty"`
    ResponseFilter *ResponseFilter `xml:"wd:Response_Filter,omitempty"`
    ResponseGroup *ResponseGroup `xml:"wd:Response_Group,omitempty"`
    ResponseResults *ResponseResults `xml:"wd:Response_Results,omitempty"`
    ResponseData *ResponseData `xml:"wd:Response_Data,omitempty"`
}

type CandidateEmailAddress struct {
    Value *string `xml:",chardata"`
}

type PreHireReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type RecruitingStageReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CandidateTagReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AppliedFrom struct {
    Value *string `xml:",chardata"`
}

type AppliedThrough struct {
    Value *string `xml:",chardata"`
}

type InternalWorkersOnly struct {
    Value *string `xml:",chardata"`
}

type CreatedFrom struct {
    Value *string `xml:",chardata"`
}

type CreatedThrough struct {
    Value *string `xml:",chardata"`
}

type CriteriaMatchScoreCategoryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type StaticCandidatePoolReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ExcludeAllAttachments struct {
    Value *string `xml:",chardata"`
}

type Candidate struct {
    CandidateReference *CandidateReference `xml:"wd:Candidate_Reference,omitempty"`
    CandidateData *CandidateData `xml:"wd:Candidate_Data,omitempty"`
}

type CandidateData struct {
    CandidateId *CandidateId `xml:"wd:Candidate_ID,omitempty"`
    PreHireReference *PreHireReference `xml:"wd:Pre_Hire_Reference,omitempty"`
    WorkerReference *WorkerReference `xml:"wd:Worker_Reference,omitempty"`
    NameData *NameData `xml:"wd:Name_Data,omitempty"`
    ContactData *ContactData `xml:"wd:Contact_Data,omitempty"`
    SocialMediaAccountData *SocialMediaAccountData `xml:"wd:Social_Media_Account_Data,omitempty"`
    StatusData *StatusData `xml:"wd:Status_Data,omitempty"`
    JobApplicationData *JobApplicationData `xml:"wd:Job_Application_Data,omitempty"`
    ProspectData *ProspectData `xml:"wd:Prospect_Data,omitempty"`
    ReferralData *ReferralData `xml:"wd:Referral_Data,omitempty"`
    CandidateIdentificationData *CandidateIdentificationData `xml:"wd:Candidate_Identification_Data,omitempty"`
    LanguageReference *LanguageReference `xml:"wd:Language_Reference,omitempty"`
    CandidateTagReference *CandidateTagReference `xml:"wd:Candidate_Tag_Reference,omitempty"`
    CandidatePoolData *CandidatePoolData `xml:"wd:Candidate_Pool_Data,omitempty"`
    CandidateAttachmentData *CandidateAttachmentData `xml:"wd:Candidate_Attachment_Data,omitempty"`
}

type CandidateId struct {
    Value *string `xml:",chardata"`
}

type LegalName struct {
    NameDetailData *NameDetailData `xml:"wd:Name_Detail_Data,omitempty"`
}

type SalutationSuffixReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type FullName struct {
    Value *string `xml:",chardata"`
}

type LocalPersonName struct {
    FirstNameLocal *FirstNameLocal `xml:"wd:First_Name___Local,omitempty"`
    MiddleNameLocal *MiddleNameLocal `xml:"wd:Middle_Name___Local,omitempty"`
    LastNameLocal *LastNameLocal `xml:"wd:Last_Name___Local,omitempty"`
    SecondaryLastNameLocal *SecondaryLastNameLocal `xml:"wd:Secondary_Last_Name___Local,omitempty"`
    FirstName2Local *FirstName2Local `xml:"wd:First_Name_2___Local,omitempty"`
    LastName2Local *LastName2Local `xml:"wd:Last_Name_2___Local,omitempty"`
}

type FirstNameLocal struct {
    Value *string `xml:",chardata"`
}

type MiddleNameLocal struct {
    Value *string `xml:",chardata"`
}

type LastNameLocal struct {
    Value *string `xml:",chardata"`
}

type SecondaryLastNameLocal struct {
    Value *string `xml:",chardata"`
}

type FirstName2Local struct {
    Value *string `xml:",chardata"`
}

type LastName2Local struct {
    Value *string `xml:",chardata"`
}

type PreferredName struct {
    DeletePreferredNameData *DeletePreferredNameData `xml:"wd:Delete_Preferred_Name_Data,omitempty"`
    NameDetailData *NameDetailData `xml:"wd:Name_Detail_Data,omitempty"`
}

type DeletePreferredNameData struct {
    Value *string `xml:",chardata"`
}

type CountryPhoneCodeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WebsiteData struct {
    UrlAddress *UrlAddress `xml:"wd:URL_Address,omitempty"`
}

type UrlAddress struct {
    Value *string `xml:",chardata"`
}

type LocationData struct {
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    AddressLine1 *AddressLine1 `xml:"wd:Address_Line_1,omitempty"`
    AddressLine2 *AddressLine2 `xml:"wd:Address_Line_2,omitempty"`
    AddressLine3 *AddressLine3 `xml:"wd:Address_Line_3,omitempty"`
    AddressLine4 *AddressLine4 `xml:"wd:Address_Line_4,omitempty"`
    AddressLine5 *AddressLine5 `xml:"wd:Address_Line_5,omitempty"`
    AddressLine6 *AddressLine6 `xml:"wd:Address_Line_6,omitempty"`
    AddressLine7 *AddressLine7 `xml:"wd:Address_Line_7,omitempty"`
    AddressLine8 *AddressLine8 `xml:"wd:Address_Line_8,omitempty"`
    AddressLine9 *AddressLine9 `xml:"wd:Address_Line_9,omitempty"`
    AddressLine1Local *AddressLine1Local `xml:"wd:Address_Line_1___Local,omitempty"`
    AddressLine2Local *AddressLine2Local `xml:"wd:Address_Line_2___Local,omitempty"`
    AddressLine3Local *AddressLine3Local `xml:"wd:Address_Line_3___Local,omitempty"`
    AddressLine4Local *AddressLine4Local `xml:"wd:Address_Line_4___Local,omitempty"`
    AddressLine5Local *AddressLine5Local `xml:"wd:Address_Line_5___Local,omitempty"`
    AddressLine6Local *AddressLine6Local `xml:"wd:Address_Line_6___Local,omitempty"`
    AddressLine7Local *AddressLine7Local `xml:"wd:Address_Line_7___Local,omitempty"`
    AddressLine8Local *AddressLine8Local `xml:"wd:Address_Line_8___Local,omitempty"`
    AddressLine9Local *AddressLine9Local `xml:"wd:Address_Line_9___Local,omitempty"`
    City *City `xml:"wd:City,omitempty"`
    CityLocal *CityLocal `xml:"wd:City___Local,omitempty"`
    CitySubdivision1 *CitySubdivision1 `xml:"wd:City_Subdivision_1,omitempty"`
    CitySubdivision2 *CitySubdivision2 `xml:"wd:City_Subdivision_2,omitempty"`
    CitySubdivision1Local *CitySubdivision1Local `xml:"wd:City_Subdivision_1___Local,omitempty"`
    CitySubdivision2Local *CitySubdivision2Local `xml:"wd:City_Subdivision_2___Local,omitempty"`
    CountryCityReference *CountryCityReference `xml:"wd:Country_City_Reference,omitempty"`
    CountryRegionReference *CountryRegionReference `xml:"wd:Country_Region_Reference,omitempty"`
    RegionSubdivision1 *RegionSubdivision1 `xml:"wd:Region_Subdivision_1,omitempty"`
    RegionSubdivision2 *RegionSubdivision2 `xml:"wd:Region_Subdivision_2,omitempty"`
    RegionSubdivision1Local *RegionSubdivision1Local `xml:"wd:Region_Subdivision_1___Local,omitempty"`
    RegionSubdivision2Local *RegionSubdivision2Local `xml:"wd:Region_Subdivision_2___Local,omitempty"`
    PostalCode *PostalCode `xml:"wd:Postal_Code,omitempty"`
}

type AddressLine1 struct {
    Value *string `xml:",chardata"`
}

type AddressLine2 struct {
    Value *string `xml:",chardata"`
}

type AddressLine3 struct {
    Value *string `xml:",chardata"`
}

type AddressLine4 struct {
    Value *string `xml:",chardata"`
}

type AddressLine5 struct {
    Value *string `xml:",chardata"`
}

type AddressLine6 struct {
    Value *string `xml:",chardata"`
}

type AddressLine7 struct {
    Value *string `xml:",chardata"`
}

type AddressLine8 struct {
    Value *string `xml:",chardata"`
}

type AddressLine9 struct {
    Value *string `xml:",chardata"`
}

type AddressLine1Local struct {
    Value *string `xml:",chardata"`
}

type AddressLine2Local struct {
    Value *string `xml:",chardata"`
}

type AddressLine3Local struct {
    Value *string `xml:",chardata"`
}

type AddressLine4Local struct {
    Value *string `xml:",chardata"`
}

type AddressLine5Local struct {
    Value *string `xml:",chardata"`
}

type AddressLine6Local struct {
    Value *string `xml:",chardata"`
}

type AddressLine7Local struct {
    Value *string `xml:",chardata"`
}

type AddressLine8Local struct {
    Value *string `xml:",chardata"`
}

type AddressLine9Local struct {
    Value *string `xml:",chardata"`
}

type City struct {
    Value *string `xml:",chardata"`
}

type CityLocal struct {
    Value *string `xml:",chardata"`
}

type CitySubdivision1 struct {
    Value *string `xml:",chardata"`
}

type CitySubdivision2 struct {
    Value *string `xml:",chardata"`
}

type CitySubdivision1Local struct {
    Value *string `xml:",chardata"`
}

type CitySubdivision2Local struct {
    Value *string `xml:",chardata"`
}

type RegionSubdivision1 struct {
    Value *string `xml:",chardata"`
}

type RegionSubdivision2 struct {
    Value *string `xml:",chardata"`
}

type RegionSubdivision1Local struct {
    Value *string `xml:",chardata"`
}

type RegionSubdivision2Local struct {
    Value *string `xml:",chardata"`
}

type SocialMediaAccountData struct {
    SocialNetworkTypeReference *SocialNetworkTypeReference `xml:"wd:Social_Network_Type_Reference,omitempty"`
    SocialNetworkAccountUrl *SocialNetworkAccountUrl `xml:"wd:Social_Network_Account_URL,omitempty"`
    SocialNetworkAccountUserName *SocialNetworkAccountUserName `xml:"wd:Social_Network_Account_User_Name,omitempty"`
}

type SocialNetworkTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SocialNetworkAccountUrl struct {
    Value *string `xml:",chardata"`
}

type SocialNetworkAccountUserName struct {
    Value *string `xml:",chardata"`
}

type StatusData struct {
    DoNotHire *DoNotHire `xml:"wd:Do_Not_Hire,omitempty"`
    Withdrawn *Withdrawn `xml:"wd:Withdrawn,omitempty"`
}

type DoNotHire struct {
    Value *string `xml:",chardata"`
}

type Withdrawn struct {
    Value *string `xml:",chardata"`
}

type JobApplicationData struct {
    JobApplicationReference *JobApplicationReference `xml:"wd:Job_Application_Reference,omitempty"`
    JobAppliedToData *JobAppliedToData `xml:"wd:Job_Applied_To_Data,omitempty"`
    ResumeAttachmentData *ResumeAttachmentData `xml:"wd:Resume_Attachment_Data,omitempty"`
    ResumeData *ResumeData `xml:"wd:Resume_Data,omitempty"`
}

type JobAppliedToData struct {
    JobApplicationId *JobApplicationId `xml:"wd:Job_Application_ID,omitempty"`
    JobRequisitionReference *JobRequisitionReference `xml:"wd:Job_Requisition_Reference,omitempty"`
    JobApplicationDate *JobApplicationDate `xml:"wd:Job_Application_Date,omitempty"`
    StageReference *StageReference `xml:"wd:Stage_Reference,omitempty"`
    WorkflowStepReference *WorkflowStepReference `xml:"wd:Workflow_Step_Reference,omitempty"`
    DispositionReference *DispositionReference `xml:"wd:Disposition_Reference,omitempty"`
    StatusTimestamp *StatusTimestamp `xml:"wd:Status_Timestamp,omitempty"`
    SourceReference *SourceReference `xml:"wd:Source_Reference,omitempty"`
    ReferredByWorkerReference *ReferredByWorkerReference `xml:"wd:Referred_By_Worker_Reference,omitempty"`
    AddedByWorkerReference *AddedByWorkerReference `xml:"wd:Added_By_Worker_Reference,omitempty"`
    CriteriaMatchScoreReference *CriteriaMatchScoreReference `xml:"wd:Criteria_Match_Score_Reference,omitempty"`
    JobApplicationAvailabilityData *JobApplicationAvailabilityData `xml:"wd:Job_Application_Availability_Data,omitempty"`
    PersonalInformationData *PersonalInformationData `xml:"wd:Personal_Information_Data,omitempty"`
    GlobalPersonalInformationData *GlobalPersonalInformationData `xml:"wd:Global_Personal_Information_Data,omitempty"`
}

type JobApplicationId struct {
    Value *string `xml:",chardata"`
}

type JobApplicationDate struct {
    Value *string `xml:",chardata"`
}

type StageReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WorkflowStepReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DispositionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type StatusTimestamp struct {
    Value *string `xml:",chardata"`
}

type SourceReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AddedByWorkerReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CriteriaMatchScoreReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type JobApplicationAvailabilityData struct {
    CandidateAvailabilityTemplateVersionReference *CandidateAvailabilityTemplateVersionReference `xml:"wd:Candidate_Availability_Template_Version_Reference,omitempty"`
    Flexible *Flexible `xml:"wd:Flexible,omitempty"`
    NoneWorkForMe *NoneWorkForMe `xml:"wd:None_Work_for_Me,omitempty"`
    AvailabilityNotes *AvailabilityNotes `xml:"wd:Availability_Notes,omitempty"`
    MondayAvailabilityReference *MondayAvailabilityReference `xml:"wd:Monday_Availability_Reference,omitempty"`
    TuesdayAvailabilityReference *TuesdayAvailabilityReference `xml:"wd:Tuesday_Availability_Reference,omitempty"`
    WednesdayAvailabilityReference *WednesdayAvailabilityReference `xml:"wd:Wednesday_Availability_Reference,omitempty"`
    ThursdayAvailabilityReference *ThursdayAvailabilityReference `xml:"wd:Thursday_Availability_Reference,omitempty"`
    FridayAvailabilityReference *FridayAvailabilityReference `xml:"wd:Friday_Availability_Reference,omitempty"`
    SaturdayAvailabilityReference *SaturdayAvailabilityReference `xml:"wd:Saturday_Availability_Reference,omitempty"`
    SundayAvailabilityReference *SundayAvailabilityReference `xml:"wd:Sunday_Availability_Reference,omitempty"`
}

type CandidateAvailabilityTemplateVersionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Flexible struct {
    Value *string `xml:",chardata"`
}

type NoneWorkForMe struct {
    Value *string `xml:",chardata"`
}

type AvailabilityNotes struct {
    Value *string `xml:",chardata"`
}

type MondayAvailabilityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type TuesdayAvailabilityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WednesdayAvailabilityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ThursdayAvailabilityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type FridayAvailabilityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SaturdayAvailabilityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SundayAvailabilityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type VeteransStatusReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DisabilityStatusLastUpdatedOn struct {
    Value *string `xml:",chardata"`
}

type AdjudicationInfoData struct {
    AdjudicationStatusReference *AdjudicationStatusReference `xml:"wd:Adjudication_Status_Reference,omitempty"`
}

type AdjudicationStatusReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type VeteransPreferenceAttachmentData struct {
    Id *Id `xml:"wd:ID,omitempty"`
    Filename *Filename `xml:"wd:Filename,omitempty"`
    FileContent *FileContent `xml:"wd:File_Content,omitempty"`
    MimeTypeReference *MimeTypeReference `xml:"wd:Mime_Type_Reference,omitempty"`
    Comment *Comment `xml:"wd:Comment,omitempty"`
}

type FileContent struct {
    Value *string `xml:",chardata"`
}

type MimeTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type GlobalPersonalInformationData struct {
    GenderReference *GenderReference `xml:"wd:Gender_Reference,omitempty"`
    DateOfBirth *DateOfBirth `xml:"wd:Date_of_Birth,omitempty"`
    BirthCountryReference *BirthCountryReference `xml:"wd:Birth_Country_Reference,omitempty"`
    BirthRegionReference *BirthRegionReference `xml:"wd:Birth_Region_Reference,omitempty"`
    CityOfBirth *CityOfBirth `xml:"wd:City_of_Birth,omitempty"`
    MaritalStatusReference *MaritalStatusReference `xml:"wd:Marital_Status_Reference,omitempty"`
    HispanicOrLatino *HispanicOrLatino `xml:"wd:Hispanic_or_Latino,omitempty"`
    AboriginalIndigenousIdentificationReference *AboriginalIndigenousIdentificationReference `xml:"wd:Aboriginal_Indigenous_Identification_Reference,omitempty"`
    AboriginalIndigenousIdentificationDetailsReference *AboriginalIndigenousIdentificationDetailsReference `xml:"wd:Aboriginal_Indigenous_Identification_Details_Reference,omitempty"`
    EthnicityReference *EthnicityReference `xml:"wd:Ethnicity_Reference,omitempty"`
    ReligionReference *ReligionReference `xml:"wd:Religion_Reference,omitempty"`
    CitizenshipReference *CitizenshipReference `xml:"wd:Citizenship_Reference,omitempty"`
    NationalityReference *NationalityReference `xml:"wd:Nationality_Reference,omitempty"`
    AdditionalNationalityReference *AdditionalNationalityReference `xml:"wd:Additional_Nationality_Reference,omitempty"`
    SocialBenefitsLocalityReference *SocialBenefitsLocalityReference `xml:"wd:Social_Benefits_Locality_Reference,omitempty"`
    LgbtIdentificationReference *LgbtIdentificationReference `xml:"wd:LGBT_Identification_Reference,omitempty"`
    SexualOrientationReference *SexualOrientationReference `xml:"wd:Sexual_Orientation_Reference,omitempty"`
    GenderIdentityReference *GenderIdentityReference `xml:"wd:Gender_Identity_Reference,omitempty"`
    PronounReference *PronounReference `xml:"wd:Pronoun_Reference,omitempty"`
    DisabilityStatusData *DisabilityStatusData `xml:"wd:Disability_Status_Data,omitempty"`
    MilitaryServiceInformationData *MilitaryServiceInformationData `xml:"wd:Military_Service_Information_Data,omitempty"`
}

type NationalityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ResumeAttachmentData struct {
    Id *Id `xml:"wd:ID,omitempty"`
    Filename *Filename `xml:"wd:Filename,omitempty"`
    FileContent *FileContent `xml:"wd:File_Content,omitempty"`
    MimeTypeReference *MimeTypeReference `xml:"wd:Mime_Type_Reference,omitempty"`
    Comment *Comment `xml:"wd:Comment,omitempty"`
}

type Summary struct {
    Value *string `xml:",chardata"`
}

type SkillData struct {
    SkillReference *SkillReference `xml:"wd:Skill_Reference,omitempty"`
    SkillName *SkillName `xml:"wd:Skill_Name,omitempty"`
}

type SkillReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SkillName struct {
    Value *string `xml:",chardata"`
}

type ExperienceData struct {
    CompanyReference *CompanyReference `xml:"wd:Company_Reference,omitempty"`
    CompanyName *CompanyName `xml:"wd:Company_Name,omitempty"`
    Title *Title `xml:"wd:Title,omitempty"`
    Location *Location `xml:"wd:Location,omitempty"`
    StartMonth *StartMonth `xml:"wd:Start_Month,omitempty"`
    StartYear *StartYear `xml:"wd:Start_Year,omitempty"`
    EndMonth *EndMonth `xml:"wd:End_Month,omitempty"`
    EndYear *EndYear `xml:"wd:End_Year,omitempty"`
    CurrentlyWorkHere *CurrentlyWorkHere `xml:"wd:Currently_Work_Here,omitempty"`
    Description *Description `xml:"wd:Description,omitempty"`
}

type CompanyReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CompanyName struct {
    Value *string `xml:",chardata"`
}

type StartMonth struct {
    Value *string `xml:",chardata"`
}

type StartYear struct {
    Value *string `xml:",chardata"`
}

type EndMonth struct {
    Value *string `xml:",chardata"`
}

type EndYear struct {
    Value *string `xml:",chardata"`
}

type CurrentlyWorkHere struct {
    Value *string `xml:",chardata"`
}

type CertificationSkillReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CertificationDataDetails struct {
    CertificationId *CertificationId `xml:"wd:Certification_ID,omitempty"`
    RemoveCertification *RemoveCertification `xml:"wd:Remove_Certification,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    CertificationReference *CertificationReference `xml:"wd:Certification_Reference,omitempty"`
    CertificationNumber *CertificationNumber `xml:"wd:Certification_Number,omitempty"`
    ExamDate *ExamDate `xml:"wd:Exam_Date,omitempty"`
    ExamScore *ExamScore `xml:"wd:Exam_Score,omitempty"`
    IssuedDate *IssuedDate `xml:"wd:Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"wd:Expiration_Date,omitempty"`
    SpecialtyDataDetails *SpecialtyDataDetails `xml:"wd:Specialty_Data_Details,omitempty"`
    AttachmentData *AttachmentData `xml:"wd:Attachment_Data,omitempty"`
}

type ExamDate struct {
    Value *string `xml:",chardata"`
}

type ExamScore struct {
    Value *string `xml:",chardata"`
}

type SpecialtyDataDetails struct {
    SpecialtyReference *SpecialtyReference `xml:"wd:Specialty_Reference,omitempty"`
    SpecialtyStartDate *SpecialtyStartDate `xml:"wd:Specialty_Start_Date,omitempty"`
    SpecialtyEndDate *SpecialtyEndDate `xml:"wd:Specialty_End_Date,omitempty"`
    SubspecialtyReference *SubspecialtyReference `xml:"wd:Subspecialty_Reference,omitempty"`
}

type SpecialtyStartDate struct {
    Value *string `xml:",chardata"`
}

type SpecialtyEndDate struct {
    Value *string `xml:",chardata"`
}

type AttachmentData struct {
    Id *Id `xml:"wd:ID,omitempty"`
    Filename *Filename `xml:"wd:Filename,omitempty"`
    FileContent *FileContent `xml:"wd:File_Content,omitempty"`
    MimeTypeReference *MimeTypeReference `xml:"wd:Mime_Type_Reference,omitempty"`
    Comment *Comment `xml:"wd:Comment,omitempty"`
}

type LanguageData struct {
    LanguageReference *LanguageReference `xml:"wd:Language_Reference,omitempty"`
    Language *Language `xml:"wd:Language,omitempty"`
}

type Native struct {
    Value *string `xml:",chardata"`
}

type QuestionnaireResponseData struct {
    ResponseData *ResponseData `xml:"wd:Response_Data,omitempty"`
}

type QuestionnaireReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type QuestionnaireAnswerData struct {
    QuestionOrder *QuestionOrder `xml:"wd:Question_Order,omitempty"`
    MultipleChoiceAnswerData *MultipleChoiceAnswerData `xml:"wd:Multiple_Choice_Answer_Data,omitempty"`
    AnswerDate *AnswerDate `xml:"wd:Answer_Date,omitempty"`
    AnswerNumeric *AnswerNumeric `xml:"wd:Answer_Numeric,omitempty"`
    AnswerText *AnswerText `xml:"wd:Answer_Text,omitempty"`
    AttachmentAnswerData *AttachmentAnswerData `xml:"wd:Attachment_Answer_Data,omitempty"`
}

type QuestionOrder struct {
    Value *string `xml:",chardata"`
}

type MultipleChoiceAnswerData struct {
    MultipleChoiceAnswerReference *MultipleChoiceAnswerReference `xml:"wd:Multiple_Choice_Answer_Reference,omitempty"`
    MultipleChoiceAnswerText *MultipleChoiceAnswerText `xml:"wd:Multiple_Choice_Answer_Text,omitempty"`
}

type MultipleChoiceAnswerReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type MultipleChoiceAnswerText struct {
    Value *string `xml:",chardata"`
}

type AnswerDate struct {
    Value *string `xml:",chardata"`
}

type AnswerNumeric struct {
    Value *string `xml:",chardata"`
}

type AnswerText struct {
    Value *string `xml:",chardata"`
}

type AttachmentAnswerData struct {
    AnswerAttachment *AnswerAttachment `xml:"wd:Answer_Attachment,omitempty"`
}

type AnswerAttachment struct {
    Id *Id `xml:"wd:ID,omitempty"`
    Filename *Filename `xml:"wd:Filename,omitempty"`
    FileContent *FileContent `xml:"wd:File_Content,omitempty"`
    MimeTypeReference *MimeTypeReference `xml:"wd:Mime_Type_Reference,omitempty"`
    Comment *Comment `xml:"wd:Comment,omitempty"`
}

type ProspectData struct {
    Prospect *Prospect `xml:"wd:Prospect,omitempty"`
    Confidential *Confidential `xml:"wd:Confidential,omitempty"`
    ReferralConsentGiven *ReferralConsentGiven `xml:"wd:Referral_Consent_Given,omitempty"`
    LevelReference *LevelReference `xml:"wd:Level_Reference,omitempty"`
    ProspectStatusReference *ProspectStatusReference `xml:"wd:Prospect_Status_Reference,omitempty"`
    ProspectTypeReference *ProspectTypeReference `xml:"wd:Prospect_Type_Reference,omitempty"`
    SourceReference *SourceReference `xml:"wd:Source_Reference,omitempty"`
    ReferredByWorkerReference *ReferredByWorkerReference `xml:"wd:Referred_By_Worker_Reference,omitempty"`
    ReferralJobReference *ReferralJobReference `xml:"wd:Referral_Job_Reference,omitempty"`
    ReferralJobAreaReference *ReferralJobAreaReference `xml:"wd:Referral_Job_Area_Reference,omitempty"`
    ReferralRelationshipReference *ReferralRelationshipReference `xml:"wd:Referral_Relationship_Reference,omitempty"`
    ReferralComment *ReferralComment `xml:"wd:Referral_Comment,omitempty"`
    AddedByWorkerReference *AddedByWorkerReference `xml:"wd:Added_By_Worker_Reference,omitempty"`
    ProspectAttachmentData *ProspectAttachmentData `xml:"wd:Prospect_Attachment_Data,omitempty"`
    ResumeData *ResumeData `xml:"wd:Resume_Data,omitempty"`
}

type Prospect struct {
    Value *string `xml:",chardata"`
}

type Confidential struct {
    Value *string `xml:",chardata"`
}

type ReferralConsentGiven struct {
    Value *string `xml:",chardata"`
}

type LevelReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ProspectStatusReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ProspectTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ReferralJobReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ReferralJobAreaReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ReferralRelationshipReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ReferralComment struct {
    Value *string `xml:",chardata"`
}

type ProspectAttachmentData struct {
    ResumeAttachments *ResumeAttachments `xml:"wd:Resume_Attachments,omitempty"`
}

type ResumeAttachments struct {
    Id *Id `xml:"wd:ID,omitempty"`
    Filename *Filename `xml:"wd:Filename,omitempty"`
    FileContent *FileContent `xml:"wd:File_Content,omitempty"`
    MimeTypeReference *MimeTypeReference `xml:"wd:Mime_Type_Reference,omitempty"`
    Comment *Comment `xml:"wd:Comment,omitempty"`
}

type ReferralData struct {
    ReferralOwnershipDetailsData *ReferralOwnershipDetailsData `xml:"wd:Referral_Ownership_Details_Data,omitempty"`
}

type ReferralOwnershipDetailsData struct {
    CandidateOwnerWorkerReference *CandidateOwnerWorkerReference `xml:"wd:Candidate_Owner_Worker_Reference,omitempty"`
    OwnershipStartDate *OwnershipStartDate `xml:"wd:Ownership_Start_Date,omitempty"`
    OwnershipEndDate *OwnershipEndDate `xml:"wd:Ownership_End_Date,omitempty"`
}

type CandidateOwnerWorkerReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type OwnershipStartDate struct {
    Value *string `xml:",chardata"`
}

type OwnershipEndDate struct {
    Value *string `xml:",chardata"`
}

type CandidateIdentificationData struct {
    NationalId *NationalId `xml:"wd:National_ID,omitempty"`
    GovernmentId *GovernmentId `xml:"wd:Government_ID,omitempty"`
}

type CandidatePoolData struct {
    CandidatePoolMembershipData *CandidatePoolMembershipData `xml:"wd:Candidate_Pool_Membership_Data,omitempty"`
}

type CandidatePoolMembershipData struct {
    StaticCandidatePoolReference *StaticCandidatePoolReference `xml:"wd:Static_Candidate_Pool_Reference,omitempty"`
    ReadinessStatusReference *ReadinessStatusReference `xml:"wd:Readiness_Status_Reference,omitempty"`
}

type ReadinessStatusReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CandidateAttachmentData struct {
    CandidateAttachmentReference *CandidateAttachmentReference `xml:"wd:Candidate_Attachment_Reference,omitempty"`
    AttachmentData *AttachmentData `xml:"wd:Attachment_Data,omitempty"`
    DocumentCategoryReference *DocumentCategoryReference `xml:"wd:Document_Category_Reference,omitempty"`
}

type CandidateAttachmentReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ChangePersonalInformationRequest struct {
    XMLName xml.Name `xml:"wd:Change_Personal_Information_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    BusinessProcessParameters *BusinessProcessParameters `xml:"wd:Business_Process_Parameters,omitempty"`
    ChangePersonalInformationData *ChangePersonalInformationData `xml:"wd:Change_Personal_Information_Data,omitempty"`
}

type ChangePersonalInformationData struct {
    PersonReference *PersonReference `xml:"wd:Person_Reference,omitempty"`
    UniversalIdReference *UniversalIdReference `xml:"wd:Universal_ID_Reference,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    PersonalInformationData *PersonalInformationData `xml:"wd:Personal_Information_Data,omitempty"`
}

type GetApplicantsResponse struct {
    XMLName xml.Name `xml:"wd:Get_Applicants_Response"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    RequestReferences *RequestReferences `xml:"wd:Request_References,omitempty"`
    RequestCriteria *RequestCriteria `xml:"wd:Request_Criteria,omitempty"`
    ResponseFilter *ResponseFilter `xml:"wd:Response_Filter,omitempty"`
    ResponseGroup *ResponseGroup `xml:"wd:Response_Group,omitempty"`
    ResponseResults *ResponseResults `xml:"wd:Response_Results,omitempty"`
    ResponseData *ResponseData `xml:"wd:Response_Data,omitempty"`
}

type Applicant struct {
    ApplicantReference *ApplicantReference `xml:"wd:Applicant_Reference,omitempty"`
    ApplicantData *ApplicantData `xml:"wd:Applicant_Data,omitempty"`
}

type PutCandidateRequest struct {
    XMLName xml.Name `xml:"wd:Put_Candidate_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    AddOnly *string `xml:"wd:Add_Only,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    CandidateReference *CandidateReference `xml:"wd:Candidate_Reference,omitempty"`
    CandidateData *CandidateData `xml:"wd:Candidate_Data,omitempty"`
}

type ChangeLegalNameRequest struct {
    XMLName xml.Name `xml:"wd:Change_Legal_Name_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    BusinessProcessParameters *BusinessProcessParameters `xml:"wd:Business_Process_Parameters,omitempty"`
    ChangeLegalNameData *ChangeLegalNameData `xml:"wd:Change_Legal_Name_Data,omitempty"`
}

type ChangeLegalNameData struct {
    PersonReference *PersonReference `xml:"wd:Person_Reference,omitempty"`
    UniversalIdReference *UniversalIdReference `xml:"wd:Universal_ID_Reference,omitempty"`
    EffectiveDate *EffectiveDate `xml:"wd:Effective_Date,omitempty"`
    NameData *NameData `xml:"wd:Name_Data,omitempty"`
}

type GetCandidatesRequest struct {
    XMLName xml.Name `xml:"wd:Get_Candidates_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    RequestReferences *RequestReferences `xml:"wd:Request_References,omitempty"`
    RequestCriteria *RequestCriteria `xml:"wd:Request_Criteria,omitempty"`
    ResponseFilter *ResponseFilter `xml:"wd:Response_Filter,omitempty"`
    ResponseGroup *ResponseGroup `xml:"wd:Response_Group,omitempty"`
}

