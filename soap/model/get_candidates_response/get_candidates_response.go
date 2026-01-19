// Generated from ../model/get_candidates_response/get_candidates_response.go
package get_candidates_response

import "encoding/xml"

type GetCandidatesResponse struct {
    XMLName xml.Name `xml:"Get_Candidates_Response"`
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
    CandidateReference *CandidateReference `xml:"Candidate_Reference,omitempty"`
}

type CandidateReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Id struct {
    Value *string `xml:",chardata"`
    Type *string `xml:"type,attr,omitempty"`
    ParentId *string `xml:"parent_id,attr,omitempty"`
    ParentType *string `xml:"parent_type,attr,omitempty"`
}

type RequestCriteria struct {
    TransactionLogCriteriaData *TransactionLogCriteriaData `xml:"Transaction_Log_Criteria_Data,omitempty"`
    FirstName *FirstName `xml:"First_Name,omitempty"`
    LastName *LastName `xml:"Last_Name,omitempty"`
    CandidateEmailAddress *CandidateEmailAddress `xml:"Candidate_Email_Address,omitempty"`
    PreHireReference *PreHireReference `xml:"Pre_Hire_Reference,omitempty"`
    WorkerReference *WorkerReference `xml:"Worker_Reference,omitempty"`
    JobRequisitionReference *JobRequisitionReference `xml:"Job_Requisition_Reference,omitempty"`
    RecruitingStageReference *RecruitingStageReference `xml:"Recruiting_Stage_Reference,omitempty"`
    ApplicantSourceReference *ApplicantSourceReference `xml:"Applicant_Source_Reference,omitempty"`
    CandidateTagReference *CandidateTagReference `xml:"Candidate_Tag_Reference,omitempty"`
    AppliedFrom *AppliedFrom `xml:"Applied_From,omitempty"`
    AppliedThrough *AppliedThrough `xml:"Applied_Through,omitempty"`
    InternalWorkersOnly *InternalWorkersOnly `xml:"Internal_Workers_Only,omitempty"`
    CreatedFrom *CreatedFrom `xml:"Created_From,omitempty"`
    CreatedThrough *CreatedThrough `xml:"Created_Through,omitempty"`
    CriteriaMatchScoreCategoryReference *CriteriaMatchScoreCategoryReference `xml:"Criteria_Match_Score_Category_Reference,omitempty"`
    StaticCandidatePoolReference *StaticCandidatePoolReference `xml:"Static_Candidate_Pool_Reference,omitempty"`
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

type FirstName struct {
    Value *string `xml:",chardata"`
}

type LastName struct {
    Value *string `xml:",chardata"`
}

type CandidateEmailAddress struct {
    Value *string `xml:",chardata"`
}

type PreHireReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WorkerReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type JobRequisitionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type RecruitingStageReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ApplicantSourceReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CandidateTagReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
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
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type StaticCandidatePoolReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ResponseFilter struct {
    AsOfEffectiveDate *AsOfEffectiveDate `xml:"As_Of_Effective_Date,omitempty"`
    AsOfEntryDatetime *AsOfEntryDatetime `xml:"As_Of_Entry_DateTime,omitempty"`
    Page *Page `xml:"Page,omitempty"`
    Count *Count `xml:"Count,omitempty"`
}

type AsOfEffectiveDate struct {
    Value *string `xml:",chardata"`
}

type AsOfEntryDatetime struct {
    Value *string `xml:",chardata"`
}

type Page struct {
    Value *string `xml:",chardata"`
}

type Count struct {
    Value *string `xml:",chardata"`
}

type ResponseGroup struct {
    IncludeReference *IncludeReference `xml:"Include_Reference,omitempty"`
    ExcludeAllAttachments *ExcludeAllAttachments `xml:"Exclude_All_Attachments,omitempty"`
}

type IncludeReference struct {
    Value *string `xml:",chardata"`
}

type ExcludeAllAttachments struct {
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
    Candidate *Candidate `xml:"Candidate,omitempty"`
    QuestionnaireReference *QuestionnaireReference `xml:"Questionnaire_Reference,omitempty"`
    QuestionnaireAnswerData *QuestionnaireAnswerData `xml:"Questionnaire_Answer_Data,omitempty"`
}

type Candidate struct {
    CandidateReference *CandidateReference `xml:"Candidate_Reference,omitempty"`
    CandidateData *CandidateData `xml:"Candidate_Data,omitempty"`
}

type CandidateData struct {
    CandidateId *CandidateId `xml:"Candidate_ID,omitempty"`
    PreHireReference *PreHireReference `xml:"Pre_Hire_Reference,omitempty"`
    WorkerReference *WorkerReference `xml:"Worker_Reference,omitempty"`
    NameData *NameData `xml:"Name_Data,omitempty"`
    ContactData *ContactData `xml:"Contact_Data,omitempty"`
    SocialMediaAccountData *SocialMediaAccountData `xml:"Social_Media_Account_Data,omitempty"`
    StatusData *StatusData `xml:"Status_Data,omitempty"`
    JobApplicationData *JobApplicationData `xml:"Job_Application_Data,omitempty"`
    ProspectData *ProspectData `xml:"Prospect_Data,omitempty"`
    ReferralData *ReferralData `xml:"Referral_Data,omitempty"`
    CandidateIdentificationData *CandidateIdentificationData `xml:"Candidate_Identification_Data,omitempty"`
    LanguageReference *LanguageReference `xml:"Language_Reference,omitempty"`
    CandidateTagReference *CandidateTagReference `xml:"Candidate_Tag_Reference,omitempty"`
    CandidatePoolData *CandidatePoolData `xml:"Candidate_Pool_Data,omitempty"`
    CandidateAttachmentData *CandidateAttachmentData `xml:"Candidate_Attachment_Data,omitempty"`
}

type CandidateId struct {
    Value *string `xml:",chardata"`
}

type NameData struct {
    LegalName *LegalName `xml:"Legal_Name,omitempty"`
    PreferredName *PreferredName `xml:"Preferred_Name,omitempty"`
}

type LegalName struct {
    NameDetailData *NameDetailData `xml:"Name_Detail_Data,omitempty"`
}

type NameDetailData struct {
    TitleReference *TitleReference `xml:"Title_Reference,omitempty"`
    SalutationSuffixReference *SalutationSuffixReference `xml:"Salutation_Suffix_Reference,omitempty"`
    FullName *FullName `xml:"Full_Name,omitempty"`
    FirstName *FirstName `xml:"First_Name,omitempty"`
    MiddleName *MiddleName `xml:"Middle_Name,omitempty"`
    HereditarySuffixReference *HereditarySuffixReference `xml:"Hereditary_Suffix_Reference,omitempty"`
    LastName *LastName `xml:"Last_Name,omitempty"`
    SecondaryLastName *SecondaryLastName `xml:"Secondary_Last_Name,omitempty"`
    TertiaryLastName *TertiaryLastName `xml:"Tertiary_Last_Name,omitempty"`
    SocialSuffixReference *SocialSuffixReference `xml:"Social_Suffix_Reference,omitempty"`
    LocalPersonName *LocalPersonName `xml:"Local_Person_Name,omitempty"`
}

type TitleReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SalutationSuffixReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type FullName struct {
    Value *string `xml:",chardata"`
}

type MiddleName struct {
    Value *string `xml:",chardata"`
}

type HereditarySuffixReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SecondaryLastName struct {
    Value *string `xml:",chardata"`
}

type TertiaryLastName struct {
    Value *string `xml:",chardata"`
}

type SocialSuffixReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type LocalPersonName struct {
    FirstNameLocal *FirstNameLocal `xml:"First_Name___Local,omitempty"`
    MiddleNameLocal *MiddleNameLocal `xml:"Middle_Name___Local,omitempty"`
    LastNameLocal *LastNameLocal `xml:"Last_Name___Local,omitempty"`
    SecondaryLastNameLocal *SecondaryLastNameLocal `xml:"Secondary_Last_Name___Local,omitempty"`
    FirstName2Local *FirstName2Local `xml:"First_Name_2___Local,omitempty"`
    LastName2Local *LastName2Local `xml:"Last_Name_2___Local,omitempty"`
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
    DeletePreferredNameData *DeletePreferredNameData `xml:"Delete_Preferred_Name_Data,omitempty"`
    NameDetailData *NameDetailData `xml:"Name_Detail_Data,omitempty"`
}

type DeletePreferredNameData struct {
    Value *string `xml:",chardata"`
}

type ContactData struct {
    PhoneDeviceTypeReference *PhoneDeviceTypeReference `xml:"Phone_Device_Type_Reference,omitempty"`
    CountryPhoneCodeReference *CountryPhoneCodeReference `xml:"Country_Phone_Code_Reference,omitempty"`
    PhoneNumber *PhoneNumber `xml:"Phone_Number,omitempty"`
    PhoneExtension *PhoneExtension `xml:"Phone_Extension,omitempty"`
    EmailAddress *EmailAddress `xml:"Email_Address,omitempty"`
    WebsiteData *WebsiteData `xml:"Website_Data,omitempty"`
    LocationData *LocationData `xml:"Location_Data,omitempty"`
}

type PhoneDeviceTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CountryPhoneCodeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PhoneNumber struct {
    Value *string `xml:",chardata"`
}

type PhoneExtension struct {
    Value *string `xml:",chardata"`
}

type EmailAddress struct {
    Value *string `xml:",chardata"`
}

type WebsiteData struct {
    UrlAddress *UrlAddress `xml:"URL_Address,omitempty"`
}

type UrlAddress struct {
    Value *string `xml:",chardata"`
}

type LocationData struct {
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"`
    AddressLine1 *AddressLine1 `xml:"Address_Line_1,omitempty"`
    AddressLine2 *AddressLine2 `xml:"Address_Line_2,omitempty"`
    AddressLine3 *AddressLine3 `xml:"Address_Line_3,omitempty"`
    AddressLine4 *AddressLine4 `xml:"Address_Line_4,omitempty"`
    AddressLine5 *AddressLine5 `xml:"Address_Line_5,omitempty"`
    AddressLine6 *AddressLine6 `xml:"Address_Line_6,omitempty"`
    AddressLine7 *AddressLine7 `xml:"Address_Line_7,omitempty"`
    AddressLine8 *AddressLine8 `xml:"Address_Line_8,omitempty"`
    AddressLine9 *AddressLine9 `xml:"Address_Line_9,omitempty"`
    AddressLine1Local *AddressLine1Local `xml:"Address_Line_1___Local,omitempty"`
    AddressLine2Local *AddressLine2Local `xml:"Address_Line_2___Local,omitempty"`
    AddressLine3Local *AddressLine3Local `xml:"Address_Line_3___Local,omitempty"`
    AddressLine4Local *AddressLine4Local `xml:"Address_Line_4___Local,omitempty"`
    AddressLine5Local *AddressLine5Local `xml:"Address_Line_5___Local,omitempty"`
    AddressLine6Local *AddressLine6Local `xml:"Address_Line_6___Local,omitempty"`
    AddressLine7Local *AddressLine7Local `xml:"Address_Line_7___Local,omitempty"`
    AddressLine8Local *AddressLine8Local `xml:"Address_Line_8___Local,omitempty"`
    AddressLine9Local *AddressLine9Local `xml:"Address_Line_9___Local,omitempty"`
    City *City `xml:"City,omitempty"`
    CityLocal *CityLocal `xml:"City___Local,omitempty"`
    CitySubdivision1 *CitySubdivision1 `xml:"City_Subdivision_1,omitempty"`
    CitySubdivision2 *CitySubdivision2 `xml:"City_Subdivision_2,omitempty"`
    CitySubdivision1Local *CitySubdivision1Local `xml:"City_Subdivision_1___Local,omitempty"`
    CitySubdivision2Local *CitySubdivision2Local `xml:"City_Subdivision_2___Local,omitempty"`
    CountryCityReference *CountryCityReference `xml:"Country_City_Reference,omitempty"`
    CountryRegionReference *CountryRegionReference `xml:"Country_Region_Reference,omitempty"`
    RegionSubdivision1 *RegionSubdivision1 `xml:"Region_Subdivision_1,omitempty"`
    RegionSubdivision2 *RegionSubdivision2 `xml:"Region_Subdivision_2,omitempty"`
    RegionSubdivision1Local *RegionSubdivision1Local `xml:"Region_Subdivision_1___Local,omitempty"`
    RegionSubdivision2Local *RegionSubdivision2Local `xml:"Region_Subdivision_2___Local,omitempty"`
    PostalCode *PostalCode `xml:"Postal_Code,omitempty"`
}

type CountryReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
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

type CountryCityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CountryRegionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
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

type PostalCode struct {
    Value *string `xml:",chardata"`
}

type SocialMediaAccountData struct {
    SocialNetworkTypeReference *SocialNetworkTypeReference `xml:"Social_Network_Type_Reference,omitempty"`
    SocialNetworkAccountUrl *SocialNetworkAccountUrl `xml:"Social_Network_Account_URL,omitempty"`
    SocialNetworkAccountUserName *SocialNetworkAccountUserName `xml:"Social_Network_Account_User_Name,omitempty"`
}

type SocialNetworkTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SocialNetworkAccountUrl struct {
    Value *string `xml:",chardata"`
}

type SocialNetworkAccountUserName struct {
    Value *string `xml:",chardata"`
}

type StatusData struct {
    DoNotHire *DoNotHire `xml:"Do_Not_Hire,omitempty"`
    Withdrawn *Withdrawn `xml:"Withdrawn,omitempty"`
}

type DoNotHire struct {
    Value *string `xml:",chardata"`
}

type Withdrawn struct {
    Value *string `xml:",chardata"`
}

type JobApplicationData struct {
    JobApplicationReference *JobApplicationReference `xml:"Job_Application_Reference,omitempty"`
    JobAppliedToData *JobAppliedToData `xml:"Job_Applied_To_Data,omitempty"`
    ResumeAttachmentData *ResumeAttachmentData `xml:"Resume_Attachment_Data,omitempty"`
    ResumeData *ResumeData `xml:"Resume_Data,omitempty"`
}

type JobApplicationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type JobAppliedToData struct {
    JobApplicationId *JobApplicationId `xml:"Job_Application_ID,omitempty"`
    JobRequisitionReference *JobRequisitionReference `xml:"Job_Requisition_Reference,omitempty"`
    JobApplicationDate *JobApplicationDate `xml:"Job_Application_Date,omitempty"`
    StageReference *StageReference `xml:"Stage_Reference,omitempty"`
    WorkflowStepReference *WorkflowStepReference `xml:"Workflow_Step_Reference,omitempty"`
    DispositionReference *DispositionReference `xml:"Disposition_Reference,omitempty"`
    StatusTimestamp *StatusTimestamp `xml:"Status_Timestamp,omitempty"`
    SourceReference *SourceReference `xml:"Source_Reference,omitempty"`
    ReferredByWorkerReference *ReferredByWorkerReference `xml:"Referred_By_Worker_Reference,omitempty"`
    AddedByWorkerReference *AddedByWorkerReference `xml:"Added_By_Worker_Reference,omitempty"`
    CriteriaMatchScoreReference *CriteriaMatchScoreReference `xml:"Criteria_Match_Score_Reference,omitempty"`
    JobApplicationAvailabilityData *JobApplicationAvailabilityData `xml:"Job_Application_Availability_Data,omitempty"`
    PersonalInformationData *PersonalInformationData `xml:"Personal_Information_Data,omitempty"`
    GlobalPersonalInformationData *GlobalPersonalInformationData `xml:"Global_Personal_Information_Data,omitempty"`
}

type JobApplicationId struct {
    Value *string `xml:",chardata"`
}

type JobApplicationDate struct {
    Value *string `xml:",chardata"`
}

type StageReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WorkflowStepReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DispositionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type StatusTimestamp struct {
    Value *string `xml:",chardata"`
}

type SourceReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ReferredByWorkerReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type AddedByWorkerReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CriteriaMatchScoreReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type JobApplicationAvailabilityData struct {
    CandidateAvailabilityTemplateVersionReference *CandidateAvailabilityTemplateVersionReference `xml:"Candidate_Availability_Template_Version_Reference,omitempty"`
    Flexible *Flexible `xml:"Flexible,omitempty"`
    NoneWorkForMe *NoneWorkForMe `xml:"None_Work_for_Me,omitempty"`
    AvailabilityNotes *AvailabilityNotes `xml:"Availability_Notes,omitempty"`
    MondayAvailabilityReference *MondayAvailabilityReference `xml:"Monday_Availability_Reference,omitempty"`
    TuesdayAvailabilityReference *TuesdayAvailabilityReference `xml:"Tuesday_Availability_Reference,omitempty"`
    WednesdayAvailabilityReference *WednesdayAvailabilityReference `xml:"Wednesday_Availability_Reference,omitempty"`
    ThursdayAvailabilityReference *ThursdayAvailabilityReference `xml:"Thursday_Availability_Reference,omitempty"`
    FridayAvailabilityReference *FridayAvailabilityReference `xml:"Friday_Availability_Reference,omitempty"`
    SaturdayAvailabilityReference *SaturdayAvailabilityReference `xml:"Saturday_Availability_Reference,omitempty"`
    SundayAvailabilityReference *SundayAvailabilityReference `xml:"Sunday_Availability_Reference,omitempty"`
}

type CandidateAvailabilityTemplateVersionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
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
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type TuesdayAvailabilityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WednesdayAvailabilityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ThursdayAvailabilityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type FridayAvailabilityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SaturdayAvailabilityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SundayAvailabilityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PersonalInformationData struct {
    GenderReference *GenderReference `xml:"Gender_Reference,omitempty"`
    EthnicityReference *EthnicityReference `xml:"Ethnicity_Reference,omitempty"`
    VeteransStatusReference *VeteransStatusReference `xml:"Veterans_Status_Reference,omitempty"`
    HispanicOrLatino *HispanicOrLatino `xml:"Hispanic_or_Latino,omitempty"`
    DisabilityStatusReference *DisabilityStatusReference `xml:"Disability_Status_Reference,omitempty"`
    DisabilityStatusLastUpdatedOn *DisabilityStatusLastUpdatedOn `xml:"Disability_Status_Last_Updated_On,omitempty"`
    VeteransPreferenceReference *VeteransPreferenceReference `xml:"Veterans_Preference_Reference,omitempty"`
    AdjudicationInfoData *AdjudicationInfoData `xml:"Adjudication_Info_Data,omitempty"`
    VeteransPreferenceAttachmentData *VeteransPreferenceAttachmentData `xml:"Veterans_Preference_Attachment_Data,omitempty"`
}

type GenderReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type EthnicityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type VeteransStatusReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type HispanicOrLatino struct {
    Value *string `xml:",chardata"`
}

type DisabilityStatusReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DisabilityStatusLastUpdatedOn struct {
    Value *string `xml:",chardata"`
}

type VeteransPreferenceReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type AdjudicationInfoData struct {
    AdjudicationStatusReference *AdjudicationStatusReference `xml:"Adjudication_Status_Reference,omitempty"`
}

type AdjudicationStatusReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type VeteransPreferenceAttachmentData struct {
    Id *Id `xml:"ID,omitempty"`
    Filename *Filename `xml:"Filename,omitempty"`
    FileContent *FileContent `xml:"File_Content,omitempty"`
    MimeTypeReference *MimeTypeReference `xml:"Mime_Type_Reference,omitempty"`
    Comment *Comment `xml:"Comment,omitempty"`
}

type Filename struct {
    Value *string `xml:",chardata"`
}

type FileContent struct {
    Value *string `xml:",chardata"`
}

type MimeTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Comment struct {
    Value *string `xml:",chardata"`
}

type GlobalPersonalInformationData struct {
    GenderReference *GenderReference `xml:"Gender_Reference,omitempty"`
    DateOfBirth *DateOfBirth `xml:"Date_of_Birth,omitempty"`
    BirthCountryReference *BirthCountryReference `xml:"Birth_Country_Reference,omitempty"`
    BirthRegionReference *BirthRegionReference `xml:"Birth_Region_Reference,omitempty"`
    CityOfBirth *CityOfBirth `xml:"City_of_Birth,omitempty"`
    MaritalStatusReference *MaritalStatusReference `xml:"Marital_Status_Reference,omitempty"`
    HispanicOrLatino *HispanicOrLatino `xml:"Hispanic_or_Latino,omitempty"`
    AboriginalIndigenousIdentificationReference *AboriginalIndigenousIdentificationReference `xml:"Aboriginal_Indigenous_Identification_Reference,omitempty"`
    AboriginalIndigenousIdentificationDetailsReference *AboriginalIndigenousIdentificationDetailsReference `xml:"Aboriginal_Indigenous_Identification_Details_Reference,omitempty"`
    EthnicityReference *EthnicityReference `xml:"Ethnicity_Reference,omitempty"`
    ReligionReference *ReligionReference `xml:"Religion_Reference,omitempty"`
    CitizenshipReference *CitizenshipReference `xml:"Citizenship_Reference,omitempty"`
    NationalityReference *NationalityReference `xml:"Nationality_Reference,omitempty"`
    AdditionalNationalityReference *AdditionalNationalityReference `xml:"Additional_Nationality_Reference,omitempty"`
    SocialBenefitsLocalityReference *SocialBenefitsLocalityReference `xml:"Social_Benefits_Locality_Reference,omitempty"`
    LgbtIdentificationReference *LgbtIdentificationReference `xml:"LGBT_Identification_Reference,omitempty"`
    SexualOrientationReference *SexualOrientationReference `xml:"Sexual_Orientation_Reference,omitempty"`
    GenderIdentityReference *GenderIdentityReference `xml:"Gender_Identity_Reference,omitempty"`
    PronounReference *PronounReference `xml:"Pronoun_Reference,omitempty"`
    DisabilityStatusData *DisabilityStatusData `xml:"Disability_Status_Data,omitempty"`
    MilitaryServiceInformationData *MilitaryServiceInformationData `xml:"Military_Service_Information_Data,omitempty"`
}

type DateOfBirth struct {
    Value *string `xml:",chardata"`
}

type BirthCountryReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type BirthRegionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CityOfBirth struct {
    Value *string `xml:",chardata"`
}

type MaritalStatusReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type AboriginalIndigenousIdentificationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type AboriginalIndigenousIdentificationDetailsReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ReligionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CitizenshipReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type NationalityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type AdditionalNationalityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SocialBenefitsLocalityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type LgbtIdentificationReference struct {
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

type DisabilityStatusData struct {
    Delete *string `xml:"Delete,attr,omitempty"`
    DisabilityStatusReference *DisabilityStatusReference `xml:"Disability_Status_Reference,omitempty"`
    DisabilityReference *DisabilityReference `xml:"Disability_Reference,omitempty"`
    DisabilityStatusDate *DisabilityStatusDate `xml:"Disability_Status_Date,omitempty"`
}

type DisabilityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DisabilityStatusDate struct {
    Value *string `xml:",chardata"`
}

type MilitaryServiceInformationData struct {
    Delete *string `xml:"Delete,attr,omitempty"`
    MilitaryServiceReference *MilitaryServiceReference `xml:"Military_Service_Reference,omitempty"`
    MilitaryServiceData *MilitaryServiceData `xml:"Military_Service_Data,omitempty"`
}

type MilitaryServiceReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type MilitaryServiceData struct {
    MilitaryStatusReference *MilitaryStatusReference `xml:"Military_Status_Reference,omitempty"`
    MilitaryDischargeDate *MilitaryDischargeDate `xml:"Military_Discharge_Date,omitempty"`
    MilitaryStatusBeginDate *MilitaryStatusBeginDate `xml:"Military_Status_Begin_Date,omitempty"`
    MilitaryServiceTypeReference *MilitaryServiceTypeReference `xml:"Military_Service_Type_Reference,omitempty"`
    MilitaryRankReference *MilitaryRankReference `xml:"Military_Rank_Reference,omitempty"`
    Notes *Notes `xml:"Notes,omitempty"`
    MilitaryDischargeTypeReference *MilitaryDischargeTypeReference `xml:"Military_Discharge_Type_Reference,omitempty"`
}

type MilitaryStatusReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type MilitaryDischargeDate struct {
    Value *string `xml:",chardata"`
}

type MilitaryStatusBeginDate struct {
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

type MilitaryDischargeTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ResumeAttachmentData struct {
    Id *Id `xml:"ID,omitempty"`
    Filename *Filename `xml:"Filename,omitempty"`
    FileContent *FileContent `xml:"File_Content,omitempty"`
    MimeTypeReference *MimeTypeReference `xml:"Mime_Type_Reference,omitempty"`
    Comment *Comment `xml:"Comment,omitempty"`
}

type ResumeData struct {
    Summary *Summary `xml:"Summary,omitempty"`
    SkillData *SkillData `xml:"Skill_Data,omitempty"`
    ExperienceData *ExperienceData `xml:"Experience_Data,omitempty"`
    EducationData *EducationData `xml:"Education_Data,omitempty"`
    CertificationData *CertificationData `xml:"Certification_Data,omitempty"`
    LanguageData *LanguageData `xml:"Language_Data,omitempty"`
    QuestionnaireResponseData *QuestionnaireResponseData `xml:"Questionnaire_Response_Data,omitempty"`
}

type Summary struct {
    Value *string `xml:",chardata"`
}

type SkillData struct {
    SkillReference *SkillReference `xml:"Skill_Reference,omitempty"`
    SkillName *SkillName `xml:"Skill_Name,omitempty"`
}

type SkillReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SkillName struct {
    Value *string `xml:",chardata"`
}

type ExperienceData struct {
    CompanyReference *CompanyReference `xml:"Company_Reference,omitempty"`
    CompanyName *CompanyName `xml:"Company_Name,omitempty"`
    Title *Title `xml:"Title,omitempty"`
    Location *Location `xml:"Location,omitempty"`
    StartMonth *StartMonth `xml:"Start_Month,omitempty"`
    StartYear *StartYear `xml:"Start_Year,omitempty"`
    EndMonth *EndMonth `xml:"End_Month,omitempty"`
    EndYear *EndYear `xml:"End_Year,omitempty"`
    CurrentlyWorkHere *CurrentlyWorkHere `xml:"Currently_Work_Here,omitempty"`
    Description *Description `xml:"Description,omitempty"`
}

type CompanyReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CompanyName struct {
    Value *string `xml:",chardata"`
}

type Title struct {
    Value *string `xml:",chardata"`
}

type Location struct {
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

type Description struct {
    Value *string `xml:",chardata"`
}

type EducationData struct {
    SchoolReference *SchoolReference `xml:"School_Reference,omitempty"`
    SchoolName *SchoolName `xml:"School_Name,omitempty"`
    FirstYearAttended *FirstYearAttended `xml:"First_Year_Attended,omitempty"`
    LastYearAttended *LastYearAttended `xml:"Last_Year_Attended,omitempty"`
    DegreeReference *DegreeReference `xml:"Degree_Reference,omitempty"`
    FieldOfStudyReference *FieldOfStudyReference `xml:"Field_of_Study_Reference,omitempty"`
    GradeAverage *GradeAverage `xml:"Grade_Average,omitempty"`
}

type SchoolReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SchoolName struct {
    Value *string `xml:",chardata"`
}

type FirstYearAttended struct {
    Value *string `xml:",chardata"`
}

type LastYearAttended struct {
    Value *string `xml:",chardata"`
}

type DegreeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type FieldOfStudyReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type GradeAverage struct {
    Value *string `xml:",chardata"`
}

type CertificationData struct {
    CertificationSkillReference *CertificationSkillReference `xml:"Certification_Skill_Reference,omitempty"`
    CertificationDataDetails *CertificationDataDetails `xml:"Certification_Data_Details,omitempty"`
}

type CertificationSkillReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CertificationDataDetails struct {
    CertificationId *CertificationId `xml:"Certification_ID,omitempty"`
    RemoveCertification *RemoveCertification `xml:"Remove_Certification,omitempty"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"`
    CertificationReference *CertificationReference `xml:"Certification_Reference,omitempty"`
    CertificationNumber *CertificationNumber `xml:"Certification_Number,omitempty"`
    ExamDate *ExamDate `xml:"Exam_Date,omitempty"`
    ExamScore *ExamScore `xml:"Exam_Score,omitempty"`
    IssuedDate *IssuedDate `xml:"Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"Expiration_Date,omitempty"`
    SpecialtyDataDetails *SpecialtyDataDetails `xml:"Specialty_Data_Details,omitempty"`
    AttachmentData *AttachmentData `xml:"Attachment_Data,omitempty"`
}

type CertificationId struct {
    Value *string `xml:",chardata"`
}

type RemoveCertification struct {
    Value *string `xml:",chardata"`
}

type CertificationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CertificationNumber struct {
    Value *string `xml:",chardata"`
}

type ExamDate struct {
    Value *string `xml:",chardata"`
}

type ExamScore struct {
    Value *string `xml:",chardata"`
}

type IssuedDate struct {
    Value *string `xml:",chardata"`
}

type ExpirationDate struct {
    Value *string `xml:",chardata"`
}

type SpecialtyDataDetails struct {
    SpecialtyReference *SpecialtyReference `xml:"Specialty_Reference,omitempty"`
    SpecialtyStartDate *SpecialtyStartDate `xml:"Specialty_Start_Date,omitempty"`
    SpecialtyEndDate *SpecialtyEndDate `xml:"Specialty_End_Date,omitempty"`
    SubspecialtyReference *SubspecialtyReference `xml:"Subspecialty_Reference,omitempty"`
}

type SpecialtyReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SpecialtyStartDate struct {
    Value *string `xml:",chardata"`
}

type SpecialtyEndDate struct {
    Value *string `xml:",chardata"`
}

type SubspecialtyReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type AttachmentData struct {
    Id *Id `xml:"ID,omitempty"`
    Filename *Filename `xml:"Filename,omitempty"`
    FileContent *FileContent `xml:"File_Content,omitempty"`
    MimeTypeReference *MimeTypeReference `xml:"Mime_Type_Reference,omitempty"`
    Comment *Comment `xml:"Comment,omitempty"`
}

type LanguageData struct {
    LanguageReference *LanguageReference `xml:"Language_Reference,omitempty"`
    Language *Language `xml:"Language,omitempty"`
}

type LanguageReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Language struct {
    Native *Native `xml:"Native,omitempty"`
    LanguageAbility *LanguageAbility `xml:"Language_Ability,omitempty"`
}

type Native struct {
    Value *string `xml:",chardata"`
}

type LanguageAbility struct {
    LanguageAbilityData *LanguageAbilityData `xml:"Language_Ability_Data,omitempty"`
}

type LanguageAbilityData struct {
    LanguageProficiencyReference *LanguageProficiencyReference `xml:"Language_Proficiency_Reference,omitempty"`
    LanguageAbilityTypeReference *LanguageAbilityTypeReference `xml:"Language_Ability_Type_Reference,omitempty"`
}

type LanguageProficiencyReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type LanguageAbilityTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type QuestionnaireResponseData struct {
    ResponseData *ResponseData `xml:"Response_Data,omitempty"`
}

type QuestionnaireReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type QuestionnaireAnswerData struct {
    QuestionOrder *QuestionOrder `xml:"Question_Order,omitempty"`
    MultipleChoiceAnswerData *MultipleChoiceAnswerData `xml:"Multiple_Choice_Answer_Data,omitempty"`
    AnswerDate *AnswerDate `xml:"Answer_Date,omitempty"`
    AnswerNumeric *AnswerNumeric `xml:"Answer_Numeric,omitempty"`
    AnswerText *AnswerText `xml:"Answer_Text,omitempty"`
    AttachmentAnswerData *AttachmentAnswerData `xml:"Attachment_Answer_Data,omitempty"`
}

type QuestionOrder struct {
    Value *string `xml:",chardata"`
}

type MultipleChoiceAnswerData struct {
    MultipleChoiceAnswerReference *MultipleChoiceAnswerReference `xml:"Multiple_Choice_Answer_Reference,omitempty"`
    MultipleChoiceAnswerText *MultipleChoiceAnswerText `xml:"Multiple_Choice_Answer_Text,omitempty"`
}

type MultipleChoiceAnswerReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
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
    AnswerAttachment *AnswerAttachment `xml:"Answer_Attachment,omitempty"`
}

type AnswerAttachment struct {
    Id *Id `xml:"ID,omitempty"`
    Filename *Filename `xml:"Filename,omitempty"`
    FileContent *FileContent `xml:"File_Content,omitempty"`
    MimeTypeReference *MimeTypeReference `xml:"Mime_Type_Reference,omitempty"`
    Comment *Comment `xml:"Comment,omitempty"`
}

type ProspectData struct {
    Prospect *Prospect `xml:"Prospect,omitempty"`
    Confidential *Confidential `xml:"Confidential,omitempty"`
    ReferralConsentGiven *ReferralConsentGiven `xml:"Referral_Consent_Given,omitempty"`
    LevelReference *LevelReference `xml:"Level_Reference,omitempty"`
    ProspectStatusReference *ProspectStatusReference `xml:"Prospect_Status_Reference,omitempty"`
    ProspectTypeReference *ProspectTypeReference `xml:"Prospect_Type_Reference,omitempty"`
    SourceReference *SourceReference `xml:"Source_Reference,omitempty"`
    ReferredByWorkerReference *ReferredByWorkerReference `xml:"Referred_By_Worker_Reference,omitempty"`
    ReferralJobReference *ReferralJobReference `xml:"Referral_Job_Reference,omitempty"`
    ReferralJobAreaReference *ReferralJobAreaReference `xml:"Referral_Job_Area_Reference,omitempty"`
    ReferralRelationshipReference *ReferralRelationshipReference `xml:"Referral_Relationship_Reference,omitempty"`
    ReferralComment *ReferralComment `xml:"Referral_Comment,omitempty"`
    AddedByWorkerReference *AddedByWorkerReference `xml:"Added_By_Worker_Reference,omitempty"`
    ProspectAttachmentData *ProspectAttachmentData `xml:"Prospect_Attachment_Data,omitempty"`
    ResumeData *ResumeData `xml:"Resume_Data,omitempty"`
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
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ProspectStatusReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ProspectTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ReferralJobReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ReferralJobAreaReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ReferralRelationshipReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ReferralComment struct {
    Value *string `xml:",chardata"`
}

type ProspectAttachmentData struct {
    ResumeAttachments *ResumeAttachments `xml:"Resume_Attachments,omitempty"`
}

type ResumeAttachments struct {
    Id *Id `xml:"ID,omitempty"`
    Filename *Filename `xml:"Filename,omitempty"`
    FileContent *FileContent `xml:"File_Content,omitempty"`
    MimeTypeReference *MimeTypeReference `xml:"Mime_Type_Reference,omitempty"`
    Comment *Comment `xml:"Comment,omitempty"`
}

type ReferralData struct {
    ReferralOwnershipDetailsData *ReferralOwnershipDetailsData `xml:"Referral_Ownership_Details_Data,omitempty"`
}

type ReferralOwnershipDetailsData struct {
    CandidateOwnerWorkerReference *CandidateOwnerWorkerReference `xml:"Candidate_Owner_Worker_Reference,omitempty"`
    OwnershipStartDate *OwnershipStartDate `xml:"Ownership_Start_Date,omitempty"`
    OwnershipEndDate *OwnershipEndDate `xml:"Ownership_End_Date,omitempty"`
}

type CandidateOwnerWorkerReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type OwnershipStartDate struct {
    Value *string `xml:",chardata"`
}

type OwnershipEndDate struct {
    Value *string `xml:",chardata"`
}

type CandidateIdentificationData struct {
    NationalId *NationalId `xml:"National_ID,omitempty"`
    GovernmentId *GovernmentId `xml:"Government_ID,omitempty"`
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

type CandidatePoolData struct {
    CandidatePoolMembershipData *CandidatePoolMembershipData `xml:"Candidate_Pool_Membership_Data,omitempty"`
}

type CandidatePoolMembershipData struct {
    StaticCandidatePoolReference *StaticCandidatePoolReference `xml:"Static_Candidate_Pool_Reference,omitempty"`
    ReadinessStatusReference *ReadinessStatusReference `xml:"Readiness_Status_Reference,omitempty"`
}

type ReadinessStatusReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CandidateAttachmentData struct {
    CandidateAttachmentReference *CandidateAttachmentReference `xml:"Candidate_Attachment_Reference,omitempty"`
    AttachmentData *AttachmentData `xml:"Attachment_Data,omitempty"`
    DocumentCategoryReference *DocumentCategoryReference `xml:"Document_Category_Reference,omitempty"`
}

type CandidateAttachmentReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DocumentCategoryReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

