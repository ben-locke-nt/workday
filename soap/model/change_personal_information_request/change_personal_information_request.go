// Generated from ../model/change_personal_information_request/change_personal_information_request.go
package change_personal_information_request

import "encoding/xml"

type ChangePersonalInformationRequest struct {
    XMLName xml.Name `xml:"wd:Change_Personal_Information_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    BusinessProcessParameters *BusinessProcessParameters `xml:"wd:Business_Process_Parameters,omitempty"`
    ChangePersonalInformationData *ChangePersonalInformationData `xml:"wd:Change_Personal_Information_Data,omitempty"`
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

type ChangePersonalInformationData struct {
    PersonReference *PersonReference `xml:"wd:Person_Reference,omitempty"`
    UniversalIdReference *UniversalIdReference `xml:"wd:Universal_ID_Reference,omitempty"`
    CountryReference *CountryReference `xml:"wd:Country_Reference,omitempty"`
    PersonalInformationData *PersonalInformationData `xml:"wd:Personal_Information_Data,omitempty"`
}

type PersonReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type UniversalIdReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CountryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PersonalInformationData struct {
    DateOfBirth *DateOfBirth `xml:"wd:Date_of_Birth,omitempty"`
    BirthCountryReference *BirthCountryReference `xml:"wd:Birth_Country_Reference,omitempty"`
    BirthRegionReference *BirthRegionReference `xml:"wd:Birth_Region_Reference,omitempty"`
    GenderReference *GenderReference `xml:"wd:Gender_Reference,omitempty"`
    ReportingGenderReference *ReportingGenderReference `xml:"wd:Reporting_Gender_Reference,omitempty"`
    DisabilityInformationData *DisabilityInformationData `xml:"wd:Disability_Information_Data,omitempty"`
    MaritalStatusReference *MaritalStatusReference `xml:"wd:Marital_Status_Reference,omitempty"`
    CitizenshipReference *CitizenshipReference `xml:"wd:Citizenship_Reference,omitempty"`
    PrimaryNationalityReference *PrimaryNationalityReference `xml:"wd:Primary_Nationality_Reference,omitempty"`
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
    DateOfDeath *DateOfDeath `xml:"wd:Date_of_Death,omitempty"`
    CityOfBirth *CityOfBirth `xml:"wd:City_of_Birth,omitempty"`
    CityOfBirthReference *CityOfBirthReference `xml:"wd:City_of_Birth_Reference,omitempty"`
    MaritalStatusDate *MaritalStatusDate `xml:"wd:Marital_Status_Date,omitempty"`
    LastMedicalExamDate *LastMedicalExamDate `xml:"wd:Last_Medical_Exam_Date,omitempty"`
    LastMedicalExamValidTo *LastMedicalExamValidTo `xml:"wd:Last_Medical_Exam_Valid_To,omitempty"`
    MedicalExamNotes *MedicalExamNotes `xml:"wd:Medical_Exam_Notes,omitempty"`
    BloodTypeReference *BloodTypeReference `xml:"wd:Blood_Type_Reference,omitempty"`
    UsesTobacco *UsesTobacco `xml:"wd:Uses_Tobacco,omitempty"`
    SocialBenefitsLocalityReference *SocialBenefitsLocalityReference `xml:"wd:Social_Benefits_Locality_Reference,omitempty"`
    LgbtIdentificationReference *LgbtIdentificationReference `xml:"wd:LGBT_Identification_Reference,omitempty"`
    SexualOrientationReference *SexualOrientationReference `xml:"wd:Sexual_Orientation_Reference,omitempty"`
    GenderIdentityReference *GenderIdentityReference `xml:"wd:Gender_Identity_Reference,omitempty"`
    PronounReference *PronounReference `xml:"wd:Pronoun_Reference,omitempty"`
    RelativeNameInformationData *RelativeNameInformationData `xml:"wd:Relative_Name_Information_Data,omitempty"`
    VeteransPreferenceForRifReference *VeteransPreferenceForRifReference `xml:"wd:Veterans_Preference_For_RIF_Reference,omitempty"`
    SelectiveServiceRegistrationReference *SelectiveServiceRegistrationReference `xml:"wd:Selective_Service_Registration_Reference,omitempty"`
    VeteransPreferenceReference *VeteransPreferenceReference `xml:"wd:Veterans_Preference_Reference,omitempty"`
    ActiveMilitaryUniformedServiceReference *ActiveMilitaryUniformedServiceReference `xml:"wd:Active_Military_Uniformed_Service_Reference,omitempty"`
    DisabledVeteranLeaveDateReference *DisabledVeteranLeaveDateReference `xml:"wd:Disabled_Veteran_Leave_Date_Reference,omitempty"`
    UniformServiceReserveStatusReference *UniformServiceReserveStatusReference `xml:"wd:Uniform_Service_Reserve_Status_Reference,omitempty"`
    NonCountrySpecificSectionData *NonCountrySpecificSectionData `xml:"wd:Non_Country_Specific_Section_Data,omitempty"`
    CountrySpecificSectionData *CountrySpecificSectionData `xml:"wd:Country_Specific_Section_Data,omitempty"`
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

type GenderReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ReportingGenderReference struct {
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

type DisabilityStatusReference struct {
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

type MaritalStatusReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CitizenshipReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type PrimaryNationalityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AdditionalNationalityReference struct {
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

type HispanicOrLatino struct {
    Value *string `xml:",chardata"`
}

type VisualSurveyEthnicityReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type HispanicOrLatinoForVisualSurvey struct {
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

type ReligionReference struct {
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

type NativeRegionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
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

type PoliticalAffiliationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DateOfDeath struct {
    Value *string `xml:",chardata"`
}

type CityOfBirth struct {
    Value *string `xml:",chardata"`
}

type CityOfBirthReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type MaritalStatusDate struct {
    Value *string `xml:",chardata"`
}

type LastMedicalExamDate struct {
    Value *string `xml:",chardata"`
}

type LastMedicalExamValidTo struct {
    Value *string `xml:",chardata"`
}

type MedicalExamNotes struct {
    Value *string `xml:",chardata"`
}

type BloodTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type UsesTobacco struct {
    Value *string `xml:",chardata"`
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

type RelativeNameInformationData struct {
    ReplaceAll *string `xml:"wd:Replace_All,attr,omitempty"`
    RelativeNameData *RelativeNameData `xml:"wd:Relative_Name_Data,omitempty"`
}

type RelativeNameData struct {
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

type VeteransPreferenceForRifReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SelectiveServiceRegistrationReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type VeteransPreferenceReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ActiveMilitaryUniformedServiceReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type DisabledVeteranLeaveDateReference struct {
    Value *string `xml:",chardata"`
}

type UniformServiceReserveStatusReference struct {
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

