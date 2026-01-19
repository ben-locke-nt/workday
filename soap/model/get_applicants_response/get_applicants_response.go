// Generated from ../model/get_applicants_response/get_applicants_response.go
package get_applicants_response

import "encoding/xml"

type GetApplicantsResponse struct {
    XMLName xml.Name `xml:"Get_Applicants_Response"`
    Version *string `xml:"version,attr,omitempty"`
    RequestReferences *RequestReferences `xml:"Request_References,omitempty"`
    RequestCriteria *RequestCriteria `xml:"Request_Criteria,omitempty"`
    ResponseFilter *ResponseFilter `xml:"Response_Filter,omitempty"`
    ResponseGroup *ResponseGroup `xml:"Response_Group,omitempty"`
    ResponseResults *ResponseResults `xml:"Response_Results,omitempty"`
    ResponseData *ResponseData `xml:"Response_Data,omitempty"`
}

type RequestReferences struct {
    ApplicantReference *ApplicantReference `xml:"Applicant_Reference,omitempty"`
}

type ApplicantReference struct {
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
    WorkerReference *WorkerReference `xml:"Worker_Reference,omitempty"`
    FormerWorkerReference *FormerWorkerReference `xml:"Former_Worker_Reference,omitempty"`
    EmailAddress *EmailAddress `xml:"Email_Address,omitempty"`
    FieldAndParameterCriteriaData *FieldAndParameterCriteriaData `xml:"Field_And_Parameter_Criteria_Data,omitempty"`
}

type WorkerReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type FormerWorkerReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type EmailAddress struct {
    Value *string `xml:",chardata"`
}

type FieldAndParameterCriteriaData struct {
    ProviderReference *ProviderReference `xml:"Provider_Reference,omitempty"`
}

type ProviderReference struct {
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
    IncludePersonalInformation *IncludePersonalInformation `xml:"Include_Personal_Information,omitempty"`
    ShowAllPersonalInformation *ShowAllPersonalInformation `xml:"Show_All_Personal_Information,omitempty"`
    IncludeRecruitingInformation *IncludeRecruitingInformation `xml:"Include_Recruiting_Information,omitempty"`
    IncludeQualificationProfile *IncludeQualificationProfile `xml:"Include_Qualification_Profile,omitempty"`
    IncludeResume *IncludeResume `xml:"Include_Resume,omitempty"`
    IncludeBackgroundCheck *IncludeBackgroundCheck `xml:"Include_Background_Check,omitempty"`
    IncludeExternalIntegrationIdData *IncludeExternalIntegrationIdData `xml:"Include_External_Integration_ID_Data,omitempty"`
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
    Applicant *Applicant `xml:"Applicant,omitempty"`
}

type Applicant struct {
    ApplicantReference *ApplicantReference `xml:"Applicant_Reference,omitempty"`
    ApplicantData *ApplicantData `xml:"Applicant_Data,omitempty"`
}

type ApplicantData struct {
    ApplicantId *ApplicantId `xml:"Applicant_ID,omitempty"`
    PersonalData *PersonalData `xml:"Personal_Data,omitempty"`
    QualificationData *QualificationData `xml:"Qualification_Data,omitempty"`
    RecruitingData *RecruitingData `xml:"Recruiting_Data,omitempty"`
    ResumeData *ResumeData `xml:"Resume_Data,omitempty"`
    BackgroundCheckData *BackgroundCheckData `xml:"Background_Check_Data,omitempty"`
    ExternalIntegrationIdData *ExternalIntegrationIdData `xml:"External_Integration_ID_Data,omitempty"`
    DocumentFieldResultData *DocumentFieldResultData `xml:"Document_Field_Result_Data,omitempty"`
}

type ApplicantId struct {
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

type UniversalId struct {
    Value *string `xml:",chardata"`
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

type CountryReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
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

type QualificationData struct {
    ExternalJobHistory *ExternalJobHistory `xml:"External_Job_History,omitempty"`
    Competency *Competency `xml:"Competency,omitempty"`
    Certification *Certification `xml:"Certification,omitempty"`
    Training *Training `xml:"Training,omitempty"`
    AwardAndActivity *AwardAndActivity `xml:"Award_and_Activity,omitempty"`
    OrganizationMembership *OrganizationMembership `xml:"Organization_Membership,omitempty"`
    Education *Education `xml:"Education,omitempty"`
    WorkExperience *WorkExperience `xml:"Work_Experience,omitempty"`
    Language *Language `xml:"Language,omitempty"`
    InternalProjectExperience *InternalProjectExperience `xml:"Internal_Project_Experience,omitempty"`
}

type ExternalJobHistory struct {
    JobHistoryReference *JobHistoryReference `xml:"Job_History_Reference,omitempty"`
    JobHistoryData *JobHistoryData `xml:"Job_History_Data,omitempty"`
}

type JobHistoryReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type JobHistoryData struct {
    JobHistoryId *JobHistoryId `xml:"Job_History_ID,omitempty"`
    RemoveJobHistory *RemoveJobHistory `xml:"Remove_Job_History,omitempty"`
    JobTitle *JobTitle `xml:"Job_Title,omitempty"`
    Company *Company `xml:"Company,omitempty"`
    JobHistoryCompanyReference *JobHistoryCompanyReference `xml:"Job_History_Company_Reference,omitempty"`
    StartDate *StartDate `xml:"Start_Date,omitempty"`
    EndDate *EndDate `xml:"End_Date,omitempty"`
    ResponsibilitiesAndAchievements *ResponsibilitiesAndAchievements `xml:"Responsibilities_and_Achievements,omitempty"`
    Location *Location `xml:"Location,omitempty"`
    JobReference *JobReference `xml:"Job_Reference,omitempty"`
    Contact *Contact `xml:"Contact,omitempty"`
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
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type StartDate struct {
    Value *string `xml:",chardata"`
}

type EndDate struct {
    Value *string `xml:",chardata"`
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
    CompetencyLevelBehaviorReference *CompetencyLevelBehaviorReference `xml:"Competency_Level_Behavior_Reference,omitempty"`
    CompetencyLevelBehaviorDescriptor *CompetencyLevelBehaviorDescriptor `xml:"Competency_Level_Behavior_Descriptor,omitempty"`
    AssessmentComment *AssessmentComment `xml:"Assessment_Comment,omitempty"`
    AssessedOn *AssessedOn `xml:"Assessed_On,omitempty"`
    AssessedByPersonReference *AssessedByPersonReference `xml:"Assessed_By_Person_Reference,omitempty"`
    AssessedByWorkerDescriptor *AssessedByWorkerDescriptor `xml:"Assessed_By_Worker_Descriptor,omitempty"`
    CompetencyReference *CompetencyReference `xml:"Competency_Reference,omitempty"`
    CompetencyDescriptor *CompetencyDescriptor `xml:"Competency_Descriptor,omitempty"`
}

type CompetencyLevelBehaviorReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
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
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type AssessedByWorkerDescriptor struct {
    Value *string `xml:",chardata"`
}

type CompetencyReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CompetencyDescriptor struct {
    Value *string `xml:",chardata"`
}

type Certification struct {
    CertificationReference *CertificationReference `xml:"Certification_Reference,omitempty"`
    CertificationData *CertificationData `xml:"Certification_Data,omitempty"`
}

type CertificationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CertificationData struct {
    CertificationId *CertificationId `xml:"Certification_ID,omitempty"`
    RemoveCertification *RemoveCertification `xml:"Remove_Certification,omitempty"`
    CertificationReference *CertificationReference `xml:"Certification_Reference,omitempty"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"`
    CertificationName *CertificationName `xml:"Certification_Name,omitempty"`
    Issuer *Issuer `xml:"Issuer,omitempty"`
    CertificationNumber *CertificationNumber `xml:"Certification_Number,omitempty"`
    IssuedDate *IssuedDate `xml:"Issued_Date,omitempty"`
    ExpirationDate *ExpirationDate `xml:"Expiration_Date,omitempty"`
    ExaminationScore *ExaminationScore `xml:"Examination_Score,omitempty"`
    ExaminationDate *ExaminationDate `xml:"Examination_Date,omitempty"`
    SpecialtyAchievementData *SpecialtyAchievementData `xml:"Specialty_Achievement_Data,omitempty"`
    WorkerDocumentData *WorkerDocumentData `xml:"Worker_Document_Data,omitempty"`
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
    SpecialtyReference *SpecialtyReference `xml:"Specialty_Reference,omitempty"`
    StartDate *StartDate `xml:"Start_Date,omitempty"`
    EndDate *EndDate `xml:"End_Date,omitempty"`
    SubspecialtyReference *SubspecialtyReference `xml:"Subspecialty_Reference,omitempty"`
}

type SpecialtyReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SubspecialtyReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type WorkerDocumentData struct {
    FileName *FileName `xml:"File_Name,omitempty"`
    Comment *Comment `xml:"Comment,omitempty"`
    File *File `xml:"File,omitempty"`
    DocumentCategoryReference *DocumentCategoryReference `xml:"Document_Category_Reference,omitempty"`
    ContentType *ContentType `xml:"Content_Type,omitempty"`
}

type FileName struct {
    Value *string `xml:",chardata"`
}

type Comment struct {
    Value *string `xml:",chardata"`
}

type File struct {
    Value *string `xml:",chardata"`
}

type DocumentCategoryReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ContentType struct {
    Value *string `xml:",chardata"`
}

type Training struct {
    TrainingReference *TrainingReference `xml:"Training_Reference,omitempty"`
    TrainingData *TrainingData `xml:"Training_Data,omitempty"`
    Value *string `xml:",chardata"`
}

type TrainingReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type TrainingData struct {
    TrainingId *TrainingId `xml:"Training_ID,omitempty"`
    RemoveTraining *RemoveTraining `xml:"Remove_Training,omitempty"`
    Training *Training `xml:"Training,omitempty"`
    Description *Description `xml:"Description,omitempty"`
    TrainingTypeReference *TrainingTypeReference `xml:"Training_Type_Reference,omitempty"`
    CompletionDate *CompletionDate `xml:"Completion_Date,omitempty"`
    TrainingDuration *TrainingDuration `xml:"Training_Duration,omitempty"`
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
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CompletionDate struct {
    Value *string `xml:",chardata"`
}

type TrainingDuration struct {
    Value *string `xml:",chardata"`
}

type AwardAndActivity struct {
    AwardAndActivityReference *AwardAndActivityReference `xml:"Award_and_Activity_Reference,omitempty"`
    AwardAndActivityData *AwardAndActivityData `xml:"Award_and_Activity_Data,omitempty"`
}

type AwardAndActivityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type AwardAndActivityData struct {
    AwardAndActivityId *AwardAndActivityId `xml:"Award_and_Activity_ID,omitempty"`
    RemoveAwardAndActivity *RemoveAwardAndActivity `xml:"Remove_Award_and_Activity,omitempty"`
    AwardAndActivityTypeReference *AwardAndActivityTypeReference `xml:"Award_and_Activity_Type_Reference,omitempty"`
    Title *Title `xml:"Title,omitempty"`
    SponsorIssuer *SponsorIssuer `xml:"Sponsor_Issuer,omitempty"`
    StartDate *StartDate `xml:"Start_Date,omitempty"`
    EndDate *EndDate `xml:"End_Date,omitempty"`
    Description *Description `xml:"Description,omitempty"`
    Url *Url `xml:"URL,omitempty"`
    RelatedPositionReference *RelatedPositionReference `xml:"Related_Position_Reference,omitempty"`
}

type AwardAndActivityId struct {
    Value *string `xml:",chardata"`
}

type RemoveAwardAndActivity struct {
    Value *string `xml:",chardata"`
}

type AwardAndActivityTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
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
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type OrganizationMembership struct {
    OrganizationProfessionalAffiliationReference *OrganizationProfessionalAffiliationReference `xml:"Organization_Professional_Affiliation_Reference,omitempty"`
    OrganizationProfessionalAffiliationData *OrganizationProfessionalAffiliationData `xml:"Organization_Professional_Affiliation_Data,omitempty"`
}

type OrganizationProfessionalAffiliationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type OrganizationProfessionalAffiliationData struct {
    ProfessionalAffiliationId *ProfessionalAffiliationId `xml:"Professional_Affiliation_ID,omitempty"`
    RemoveProfessionalAffiliation *RemoveProfessionalAffiliation `xml:"Remove_Professional_Affiliation,omitempty"`
    ProfessionalAffiliationReference *ProfessionalAffiliationReference `xml:"Professional_Affiliation_Reference,omitempty"`
    ProfessionalAffiliation *ProfessionalAffiliation `xml:"Professional_Affiliation,omitempty"`
    ProfessionalAffiliationTypeReference *ProfessionalAffiliationTypeReference `xml:"Professional_Affiliation_Type_Reference,omitempty"`
    Affiliation *Affiliation `xml:"Affiliation,omitempty"`
    ProfessionalAffiliationRelationshipTypeReference *ProfessionalAffiliationRelationshipTypeReference `xml:"Professional_Affiliation_Relationship_Type_Reference,omitempty"`
    BeginDate *BeginDate `xml:"Begin_Date,omitempty"`
    EndDate *EndDate `xml:"End_Date,omitempty"`
    ContactInformationData *ContactInformationData `xml:"Contact_Information_Data,omitempty"`
}

type ProfessionalAffiliationId struct {
    Value *string `xml:",chardata"`
}

type RemoveProfessionalAffiliation struct {
    Value *string `xml:",chardata"`
}

type ProfessionalAffiliationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ProfessionalAffiliation struct {
    Value *string `xml:",chardata"`
}

type ProfessionalAffiliationTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Affiliation struct {
    Value *string `xml:",chardata"`
}

type ProfessionalAffiliationRelationshipTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type BeginDate struct {
    Value *string `xml:",chardata"`
}

type ContactInformationData struct {
    AddressData *AddressData `xml:"Address_Data,omitempty"`
    PhoneData *PhoneData `xml:"Phone_Data,omitempty"`
    EmailAddressData *EmailAddressData `xml:"Email_Address_Data,omitempty"`
    InstantMessengerData *InstantMessengerData `xml:"Instant_Messenger_Data,omitempty"`
    WebAddressData *WebAddressData `xml:"Web_Address_Data,omitempty"`
}

type Education struct {
    EducationReference *EducationReference `xml:"Education_Reference,omitempty"`
    EducationData *EducationData `xml:"Education_Data,omitempty"`
}

type EducationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type EducationData struct {
    EducationId *EducationId `xml:"Education_ID,omitempty"`
    RemoveEducation *RemoveEducation `xml:"Remove_Education,omitempty"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"`
    SchoolReference *SchoolReference `xml:"School_Reference,omitempty"`
    SchoolName *SchoolName `xml:"School_Name,omitempty"`
    SchoolTypeReference *SchoolTypeReference `xml:"School_Type_Reference,omitempty"`
    Location *Location `xml:"Location,omitempty"`
    DegreeReference *DegreeReference `xml:"Degree_Reference,omitempty"`
    DegreeCompletedReference *DegreeCompletedReference `xml:"Degree_Completed_Reference,omitempty"`
    DateDegreeReceived *DateDegreeReceived `xml:"Date_Degree_Received,omitempty"`
    FieldOfStudyReference *FieldOfStudyReference `xml:"Field_Of_Study_Reference,omitempty"`
    GradeAverage *GradeAverage `xml:"Grade_Average,omitempty"`
    FirstYearAttended *FirstYearAttended `xml:"First_Year_Attended,omitempty"`
    LastYearAttended *LastYearAttended `xml:"Last_Year_Attended,omitempty"`
    IsHighestLevelOfEducation *IsHighestLevelOfEducation `xml:"Is_Highest_Level_of_Education,omitempty"`
    FirstDayAttended *FirstDayAttended `xml:"First_Day_Attended,omitempty"`
    LastDayAttended *LastDayAttended `xml:"Last_Day_Attended,omitempty"`
    EducationAttachmentData *EducationAttachmentData `xml:"Education_Attachment_Data,omitempty"`
}

type EducationId struct {
    Value *string `xml:",chardata"`
}

type RemoveEducation struct {
    Value *string `xml:",chardata"`
}

type SchoolReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SchoolName struct {
    Value *string `xml:",chardata"`
}

type SchoolTypeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DegreeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DegreeCompletedReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DateDegreeReceived struct {
    Value *string `xml:",chardata"`
}

type FieldOfStudyReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
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
    FileName *FileName `xml:"File_Name,omitempty"`
    Comment *Comment `xml:"Comment,omitempty"`
    File *File `xml:"File,omitempty"`
    DocumentCategoryReference *DocumentCategoryReference `xml:"Document_Category_Reference,omitempty"`
}

type WorkExperience struct {
    ExperienceReference *ExperienceReference `xml:"Experience_Reference,omitempty"`
    RemoveExperience *RemoveExperience `xml:"Remove_Experience,omitempty"`
    ExperienceRatingReference *ExperienceRatingReference `xml:"Experience_Rating_Reference,omitempty"`
    ExperienceComment *ExperienceComment `xml:"Experience_Comment,omitempty"`
}

type ExperienceReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type RemoveExperience struct {
    Value *string `xml:",chardata"`
}

type ExperienceRatingReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ExperienceComment struct {
    Value *string `xml:",chardata"`
}

type Language struct {
    LanguageReference *LanguageReference `xml:"Language_Reference,omitempty"`
    RemoveLanguage *RemoveLanguage `xml:"Remove_Language,omitempty"`
    NativeLanguage *NativeLanguage `xml:"Native_Language,omitempty"`
    LanguageAbility *LanguageAbility `xml:"Language_Ability,omitempty"`
    AssessedOn *AssessedOn `xml:"Assessed_On,omitempty"`
    Note *Note `xml:"Note,omitempty"`
    AssessedByWorkerReference *AssessedByWorkerReference `xml:"Assessed_by_Worker_Reference,omitempty"`
}

type LanguageReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type RemoveLanguage struct {
    Value *string `xml:",chardata"`
}

type NativeLanguage struct {
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

type AssessedByWorkerReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type InternalProjectExperience struct {
    InternalProjectExperienceReference *InternalProjectExperienceReference `xml:"Internal_Project_Experience_Reference,omitempty"`
    InternalProjectExperienceData *InternalProjectExperienceData `xml:"Internal_Project_Experience_Data,omitempty"`
}

type InternalProjectExperienceReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type InternalProjectExperienceData struct {
    InternalProjectExperienceId *InternalProjectExperienceId `xml:"Internal_Project_Experience_ID,omitempty"`
    RemoveInternalProjectExperience *RemoveInternalProjectExperience `xml:"Remove_Internal_Project_Experience,omitempty"`
    InternalProject *InternalProject `xml:"Internal_Project,omitempty"`
    Description *Description `xml:"Description,omitempty"`
    StartDate *StartDate `xml:"Start_Date,omitempty"`
    EndDate *EndDate `xml:"End_Date,omitempty"`
    ProjectLeader *ProjectLeader `xml:"Project_Leader,omitempty"`
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
    ApplicantEnteredDate *ApplicantEnteredDate `xml:"Applicant_Entered_Date,omitempty"`
    ApplicantComments *ApplicantComments `xml:"Applicant_Comments,omitempty"`
    EligibleForHireReference *EligibleForHireReference `xml:"Eligible_For_Hire_Reference,omitempty"`
    EligibleForRehireComments *EligibleForRehireComments `xml:"Eligible_for_Rehire_Comments,omitempty"`
    ApplicantHasMarkedAsNoShowReference *ApplicantHasMarkedAsNoShowReference `xml:"Applicant_Has_Marked_as_No_Show_Reference,omitempty"`
    ApplicantSourceReference *ApplicantSourceReference `xml:"Applicant_Source_Reference,omitempty"`
    ReferredByWorkerReference *ReferredByWorkerReference `xml:"Referred_by_Worker_Reference,omitempty"`
    PositionsConsideredForReference *PositionsConsideredForReference `xml:"Positions_Considered_for_Reference,omitempty"`
}

type ApplicantEnteredDate struct {
    Value *string `xml:",chardata"`
}

type ApplicantComments struct {
    Value *string `xml:",chardata"`
}

type EligibleForHireReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type EligibleForRehireComments struct {
    Value *string `xml:",chardata"`
}

type ApplicantHasMarkedAsNoShowReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ApplicantSourceReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ReferredByWorkerReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PositionsConsideredForReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ResumeData struct {
    Resume *Resume `xml:"Resume,omitempty"`
    FileId *FileId `xml:"File_ID,omitempty"`
    File *File `xml:"File,omitempty"`
    Filename *Filename `xml:"FileName,omitempty"`
    Comments *Comments `xml:"Comments,omitempty"`
}

type Resume struct {
    ResumeReference *ResumeReference `xml:"Resume_Reference,omitempty"`
    ResumeData *ResumeData `xml:"Resume_Data,omitempty"`
}

type ResumeReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type FileId struct {
    Value *string `xml:",chardata"`
}

type Filename struct {
    Value *string `xml:",chardata"`
}

type BackgroundCheckData struct {
    StatusDate *StatusDate `xml:"Status_Date,omitempty"`
    StatusReference *StatusReference `xml:"Status_Reference,omitempty"`
    StatusComment *StatusComment `xml:"Status_Comment,omitempty"`
}

type StatusDate struct {
    Value *string `xml:",chardata"`
}

type StatusComment struct {
    Value *string `xml:",chardata"`
}

type ExternalIntegrationIdData struct {
    Id *Id `xml:"ID,omitempty"`
}

type DocumentFieldResultData struct {
    FieldReference *FieldReference `xml:"Field_Reference,omitempty"`
    Value *Value `xml:"Value,omitempty"`
}

type FieldReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Value struct {
    Value *string `xml:",chardata"`
}

