// Generated from ../model/put_applicant_request/put_applicant_request.go
package put_applicant_request

import "encoding/xml"

type PutApplicantRequest struct {
    XMLName xml.Name `xml:"wd:Put_Applicant_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    AddOnly *string `xml:"wd:Add_Only,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    ApplicantReference *ApplicantReference `xml:"wd:Applicant_Reference,omitempty"`
    ApplicantData *ApplicantData `xml:"wd:Applicant_Data,omitempty"`
}

type ApplicantReference struct {
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

type PersonalData struct {
    UniversalId *UniversalId `xml:"wd:Universal_ID,omitempty"`
    NameData *NameData `xml:"wd:Name_Data,omitempty"`
    PersonalInformationData *PersonalInformationData `xml:"wd:Personal_Information_Data,omitempty"`
    IdentificationData *IdentificationData `xml:"wd:Identification_Data,omitempty"`
    ContactData *ContactData `xml:"wd:Contact_Data,omitempty"`
    TobaccoUse *TobaccoUse `xml:"wd:Tobacco_Use,omitempty"`
}

type UniversalId struct {
    Value *string `xml:",chardata"`
}

type NameData struct {
    LegalNameData *LegalNameData `xml:"wd:Legal_Name_Data,omitempty"`
    PreferredNameData *PreferredNameData `xml:"wd:Preferred_Name_Data,omitempty"`
    AdditionalNameData *AdditionalNameData `xml:"wd:Additional_Name_Data,omitempty"`
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
}

type CountryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ContentType struct {
    Value *string `xml:",chardata"`
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
    FieldOfStudyReference *FieldOfStudyReference `xml:"wd:Field_Of_Study_Reference,omitempty"`
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

type EligibleForHireReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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

type FieldReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Value struct {
    Value *string `xml:",chardata"`
}

