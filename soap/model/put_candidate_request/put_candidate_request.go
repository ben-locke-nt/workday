// Generated from ../model/put_candidate_request/put_candidate_request.go
package put_candidate_request

import "encoding/xml"

type PutCandidateRequest struct {
    XMLName xml.Name `xml:"wd:Put_Candidate_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    AddOnly *string `xml:"wd:Add_Only,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    CandidateReference *CandidateReference `xml:"wd:Candidate_Reference,omitempty"`
    CandidateData *CandidateData `xml:"wd:Candidate_Data,omitempty"`
}

type CandidateReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Id struct {
    Value *string `xml:",chardata"`
    Type *string `xml:"wd:type,attr,omitempty"`
    ParentId *string `xml:"wd:parent_id,attr,omitempty"`
    ParentType *string `xml:"wd:parent_type,attr,omitempty"`
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

type PreHireReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type WorkerReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type NameData struct {
    LegalName *LegalName `xml:"wd:Legal_Name,omitempty"`
    PreferredName *PreferredName `xml:"wd:Preferred_Name,omitempty"`
}

type LegalName struct {
    NameDetailData *NameDetailData `xml:"wd:Name_Detail_Data,omitempty"`
}

type NameDetailData struct {
    TitleReference *TitleReference `xml:"wd:Title_Reference,omitempty"`
    SalutationSuffixReference *SalutationSuffixReference `xml:"wd:Salutation_Suffix_Reference,omitempty"`
    FullName *FullName `xml:"wd:Full_Name,omitempty"`
    FirstName *FirstName `xml:"wd:First_Name,omitempty"`
    MiddleName *MiddleName `xml:"wd:Middle_Name,omitempty"`
    HereditarySuffixReference *HereditarySuffixReference `xml:"wd:Hereditary_Suffix_Reference,omitempty"`
    LastName *LastName `xml:"wd:Last_Name,omitempty"`
    SecondaryLastName *SecondaryLastName `xml:"wd:Secondary_Last_Name,omitempty"`
    TertiaryLastName *TertiaryLastName `xml:"wd:Tertiary_Last_Name,omitempty"`
    SocialSuffixReference *SocialSuffixReference `xml:"wd:Social_Suffix_Reference,omitempty"`
    LocalPersonName *LocalPersonName `xml:"wd:Local_Person_Name,omitempty"`
}

type TitleReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SalutationSuffixReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type FullName struct {
    Value *string `xml:",chardata"`
}

type FirstName struct {
    Value *string `xml:",chardata"`
}

type MiddleName struct {
    Value *string `xml:",chardata"`
}

type HereditarySuffixReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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

type SocialSuffixReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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

type ContactData struct {
    PhoneDeviceTypeReference *PhoneDeviceTypeReference `xml:"wd:Phone_Device_Type_Reference,omitempty"`
    CountryPhoneCodeReference *CountryPhoneCodeReference `xml:"wd:Country_Phone_Code_Reference,omitempty"`
    PhoneNumber *PhoneNumber `xml:"wd:Phone_Number,omitempty"`
    PhoneExtension *PhoneExtension `xml:"wd:Phone_Extension,omitempty"`
    EmailAddress *EmailAddress `xml:"wd:Email_Address,omitempty"`
    WebsiteData *WebsiteData `xml:"wd:Website_Data,omitempty"`
    LocationData *LocationData `xml:"wd:Location_Data,omitempty"`
}

type PhoneDeviceTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CountryPhoneCodeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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

type CountryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CountryRegionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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

type JobApplicationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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

type JobRequisitionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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

type ReferredByWorkerReference struct {
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

type PersonalInformationData struct {
    GenderReference *GenderReference `xml:"wd:Gender_Reference,omitempty"`
    EthnicityReference *EthnicityReference `xml:"wd:Ethnicity_Reference,omitempty"`
    VeteransStatusReference *VeteransStatusReference `xml:"wd:Veterans_Status_Reference,omitempty"`
    HispanicOrLatino *HispanicOrLatino `xml:"wd:Hispanic_or_Latino,omitempty"`
    DisabilityStatusReference *DisabilityStatusReference `xml:"wd:Disability_Status_Reference,omitempty"`
    DisabilityStatusLastUpdatedOn *DisabilityStatusLastUpdatedOn `xml:"wd:Disability_Status_Last_Updated_On,omitempty"`
    VeteransPreferenceReference *VeteransPreferenceReference `xml:"wd:Veterans_Preference_Reference,omitempty"`
    AdjudicationInfoData *AdjudicationInfoData `xml:"wd:Adjudication_Info_Data,omitempty"`
    VeteransPreferenceAttachmentData *VeteransPreferenceAttachmentData `xml:"wd:Veterans_Preference_Attachment_Data,omitempty"`
}

type GenderReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type EthnicityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type VeteransStatusReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type HispanicOrLatino struct {
    Value *string `xml:",chardata"`
}

type DisabilityStatusReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DisabilityStatusLastUpdatedOn struct {
    Value *string `xml:",chardata"`
}

type VeteransPreferenceReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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

type Filename struct {
    Value *string `xml:",chardata"`
}

type FileContent struct {
    Value *string `xml:",chardata"`
}

type MimeTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Comment struct {
    Value *string `xml:",chardata"`
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

type CityOfBirth struct {
    Value *string `xml:",chardata"`
}

type MaritalStatusReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AboriginalIndigenousIdentificationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AboriginalIndigenousIdentificationDetailsReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ReligionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CitizenshipReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type NationalityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AdditionalNationalityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SocialBenefitsLocalityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type LgbtIdentificationReference struct {
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

type DisabilityStatusData struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    DisabilityStatusReference *DisabilityStatusReference `xml:"wd:Disability_Status_Reference,omitempty"`
    DisabilityReference *DisabilityReference `xml:"wd:Disability_Reference,omitempty"`
    DisabilityStatusDate *DisabilityStatusDate `xml:"wd:Disability_Status_Date,omitempty"`
}

type DisabilityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DisabilityStatusDate struct {
    Value *string `xml:",chardata"`
}

type MilitaryServiceInformationData struct {
    Delete *string `xml:"wd:Delete,attr,omitempty"`
    MilitaryServiceReference *MilitaryServiceReference `xml:"wd:Military_Service_Reference,omitempty"`
    MilitaryServiceData *MilitaryServiceData `xml:"wd:Military_Service_Data,omitempty"`
}

type MilitaryServiceReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type MilitaryServiceData struct {
    MilitaryStatusReference *MilitaryStatusReference `xml:"wd:Military_Status_Reference,omitempty"`
    MilitaryDischargeDate *MilitaryDischargeDate `xml:"wd:Military_Discharge_Date,omitempty"`
    MilitaryStatusBeginDate *MilitaryStatusBeginDate `xml:"wd:Military_Status_Begin_Date,omitempty"`
    MilitaryServiceTypeReference *MilitaryServiceTypeReference `xml:"wd:Military_Service_Type_Reference,omitempty"`
    MilitaryRankReference *MilitaryRankReference `xml:"wd:Military_Rank_Reference,omitempty"`
    Notes *Notes `xml:"wd:Notes,omitempty"`
    MilitaryDischargeTypeReference *MilitaryDischargeTypeReference `xml:"wd:Military_Discharge_Type_Reference,omitempty"`
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

type MilitaryDischargeTypeReference struct {
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

type ResumeData struct {
    Summary *Summary `xml:"wd:Summary,omitempty"`
    SkillData *SkillData `xml:"wd:Skill_Data,omitempty"`
    ExperienceData *ExperienceData `xml:"wd:Experience_Data,omitempty"`
    EducationData *EducationData `xml:"wd:Education_Data,omitempty"`
    CertificationData *CertificationData `xml:"wd:Certification_Data,omitempty"`
    LanguageData *LanguageData `xml:"wd:Language_Data,omitempty"`
    QuestionnaireResponseData *QuestionnaireResponseData `xml:"wd:Questionnaire_Response_Data,omitempty"`
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
    SchoolReference *SchoolReference `xml:"wd:School_Reference,omitempty"`
    SchoolName *SchoolName `xml:"wd:School_Name,omitempty"`
    FirstYearAttended *FirstYearAttended `xml:"wd:First_Year_Attended,omitempty"`
    LastYearAttended *LastYearAttended `xml:"wd:Last_Year_Attended,omitempty"`
    DegreeReference *DegreeReference `xml:"wd:Degree_Reference,omitempty"`
    FieldOfStudyReference *FieldOfStudyReference `xml:"wd:Field_of_Study_Reference,omitempty"`
    GradeAverage *GradeAverage `xml:"wd:Grade_Average,omitempty"`
}

type SchoolReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type FieldOfStudyReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type GradeAverage struct {
    Value *string `xml:",chardata"`
}

type CertificationData struct {
    CertificationSkillReference *CertificationSkillReference `xml:"wd:Certification_Skill_Reference,omitempty"`
    CertificationDataDetails *CertificationDataDetails `xml:"wd:Certification_Data_Details,omitempty"`
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

type CertificationId struct {
    Value *string `xml:",chardata"`
}

type RemoveCertification struct {
    Value *string `xml:",chardata"`
}

type CertificationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    SpecialtyReference *SpecialtyReference `xml:"wd:Specialty_Reference,omitempty"`
    SpecialtyStartDate *SpecialtyStartDate `xml:"wd:Specialty_Start_Date,omitempty"`
    SpecialtyEndDate *SpecialtyEndDate `xml:"wd:Specialty_End_Date,omitempty"`
    SubspecialtyReference *SubspecialtyReference `xml:"wd:Subspecialty_Reference,omitempty"`
}

type SpecialtyReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SpecialtyStartDate struct {
    Value *string `xml:",chardata"`
}

type SpecialtyEndDate struct {
    Value *string `xml:",chardata"`
}

type SubspecialtyReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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

type LanguageReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Language struct {
    Native *Native `xml:"wd:Native,omitempty"`
    LanguageAbility *LanguageAbility `xml:"wd:Language_Ability,omitempty"`
}

type Native struct {
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

type QuestionnaireResponseData struct {
    ResponseData *ResponseData `xml:"wd:Response_Data,omitempty"`
}

type ResponseData struct {
    QuestionnaireReference *QuestionnaireReference `xml:"wd:Questionnaire_Reference,omitempty"`
    QuestionnaireAnswerData *QuestionnaireAnswerData `xml:"wd:Questionnaire_Answer_Data,omitempty"`
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

type CandidateTagReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CandidatePoolData struct {
    CandidatePoolMembershipData *CandidatePoolMembershipData `xml:"wd:Candidate_Pool_Membership_Data,omitempty"`
}

type CandidatePoolMembershipData struct {
    StaticCandidatePoolReference *StaticCandidatePoolReference `xml:"wd:Static_Candidate_Pool_Reference,omitempty"`
    ReadinessStatusReference *ReadinessStatusReference `xml:"wd:Readiness_Status_Reference,omitempty"`
}

type StaticCandidatePoolReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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

type DocumentCategoryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

