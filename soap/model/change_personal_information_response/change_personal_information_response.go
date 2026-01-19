// Generated from ../model/change_personal_information_response/change_personal_information_response.go
package change_personal_information_response

import "encoding/xml"

type ChangePersonalInformationResponse struct {
    XMLName xml.Name `xml:"Change_Personal_Information_Response"`
    Version *string `xml:"version,attr,omitempty"`
    PersonalInformationChangeEventReference *PersonalInformationChangeEventReference `xml:"Personal_Information_Change_Event_Reference,omitempty"`
    PersonalInformationData *PersonalInformationData `xml:"Personal_Information_Data,omitempty"`
}

type PersonalInformationChangeEventReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Id struct {
    Value *string `xml:",chardata"`
    Type *string `xml:"type,attr,omitempty"`
    ParentId *string `xml:"parent_id,attr,omitempty"`
    ParentType *string `xml:"parent_type,attr,omitempty"`
}

type PersonalInformationData struct {
    DateOfBirth *DateOfBirth `xml:"Date_of_Birth,omitempty"`
    BirthCountryReference *BirthCountryReference `xml:"Birth_Country_Reference,omitempty"`
    BirthRegionReference *BirthRegionReference `xml:"Birth_Region_Reference,omitempty"`
    GenderReference *GenderReference `xml:"Gender_Reference,omitempty"`
    ReportingGenderReference *ReportingGenderReference `xml:"Reporting_Gender_Reference,omitempty"`
    DisabilityInformationData *DisabilityInformationData `xml:"Disability_Information_Data,omitempty"`
    MaritalStatusReference *MaritalStatusReference `xml:"Marital_Status_Reference,omitempty"`
    CitizenshipReference *CitizenshipReference `xml:"Citizenship_Reference,omitempty"`
    PrimaryNationalityReference *PrimaryNationalityReference `xml:"Primary_Nationality_Reference,omitempty"`
    AdditionalNationalityReference *AdditionalNationalityReference `xml:"Additional_Nationality_Reference,omitempty"`
    EthnicityReference *EthnicityReference `xml:"Ethnicity_Reference,omitempty"`
    RaceEthnicityDetailsReference *RaceEthnicityDetailsReference `xml:"Race_Ethnicity_Details_Reference,omitempty"`
    HispanicOrLatino *HispanicOrLatino `xml:"Hispanic_or_Latino,omitempty"`
    VisualSurveyEthnicityReference *VisualSurveyEthnicityReference `xml:"Visual_Survey_Ethnicity_Reference,omitempty"`
    HispanicOrLatinoForVisualSurvey *HispanicOrLatinoForVisualSurvey `xml:"Hispanic_or_Latino_for_Visual_Survey,omitempty"`
    AboriginalIndigenousIdentificationReference *AboriginalIndigenousIdentificationReference `xml:"Aboriginal_Indigenous_Identification_Reference,omitempty"`
    AboriginalIndigenousIdentificationDetailsReference *AboriginalIndigenousIdentificationDetailsReference `xml:"Aboriginal_Indigenous_Identification_Details_Reference,omitempty"`
    ReligionReference *ReligionReference `xml:"Religion_Reference,omitempty"`
    HukouRegionReference *HukouRegionReference `xml:"Hukou_Region_Reference,omitempty"`
    HukouSubregionReference *HukouSubregionReference `xml:"Hukou_Subregion_Reference,omitempty"`
    HukouLocality *HukouLocality `xml:"Hukou_Locality,omitempty"`
    HukouPostalCode *HukouPostalCode `xml:"Hukou_Postal_Code,omitempty"`
    HukouTypeReference *HukouTypeReference `xml:"Hukou_Type_Reference,omitempty"`
    NativeRegionReference *NativeRegionReference `xml:"Native_Region_Reference,omitempty"`
    PersonnelFileAgency *PersonnelFileAgency `xml:"Personnel_File_Agency,omitempty"`
    MilitaryInformationData *MilitaryInformationData `xml:"Military_Information_Data,omitempty"`
    PoliticalAffiliationReference *PoliticalAffiliationReference `xml:"Political_Affiliation_Reference,omitempty"`
    DateOfDeath *DateOfDeath `xml:"Date_of_Death,omitempty"`
    CityOfBirth *CityOfBirth `xml:"City_of_Birth,omitempty"`
    CityOfBirthReference *CityOfBirthReference `xml:"City_of_Birth_Reference,omitempty"`
    MaritalStatusDate *MaritalStatusDate `xml:"Marital_Status_Date,omitempty"`
    LastMedicalExamDate *LastMedicalExamDate `xml:"Last_Medical_Exam_Date,omitempty"`
    LastMedicalExamValidTo *LastMedicalExamValidTo `xml:"Last_Medical_Exam_Valid_To,omitempty"`
    MedicalExamNotes *MedicalExamNotes `xml:"Medical_Exam_Notes,omitempty"`
    BloodTypeReference *BloodTypeReference `xml:"Blood_Type_Reference,omitempty"`
    UsesTobacco *UsesTobacco `xml:"Uses_Tobacco,omitempty"`
    SocialBenefitsLocalityReference *SocialBenefitsLocalityReference `xml:"Social_Benefits_Locality_Reference,omitempty"`
    LgbtIdentificationReference *LgbtIdentificationReference `xml:"LGBT_Identification_Reference,omitempty"`
    SexualOrientationReference *SexualOrientationReference `xml:"Sexual_Orientation_Reference,omitempty"`
    GenderIdentityReference *GenderIdentityReference `xml:"Gender_Identity_Reference,omitempty"`
    PronounReference *PronounReference `xml:"Pronoun_Reference,omitempty"`
    RelativeNameInformationData *RelativeNameInformationData `xml:"Relative_Name_Information_Data,omitempty"`
    VeteransPreferenceForRifReference *VeteransPreferenceForRifReference `xml:"Veterans_Preference_For_RIF_Reference,omitempty"`
    SelectiveServiceRegistrationReference *SelectiveServiceRegistrationReference `xml:"Selective_Service_Registration_Reference,omitempty"`
    VeteransPreferenceReference *VeteransPreferenceReference `xml:"Veterans_Preference_Reference,omitempty"`
    ActiveMilitaryUniformedServiceReference *ActiveMilitaryUniformedServiceReference `xml:"Active_Military_Uniformed_Service_Reference,omitempty"`
    DisabledVeteranLeaveDateReference *DisabledVeteranLeaveDateReference `xml:"Disabled_Veteran_Leave_Date_Reference,omitempty"`
    UniformServiceReserveStatusReference *UniformServiceReserveStatusReference `xml:"Uniform_Service_Reserve_Status_Reference,omitempty"`
    NonCountrySpecificSectionData *NonCountrySpecificSectionData `xml:"Non_Country_Specific_Section_Data,omitempty"`
    CountrySpecificSectionData *CountrySpecificSectionData `xml:"Country_Specific_Section_Data,omitempty"`
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

type GenderReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ReportingGenderReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DisabilityInformationData struct {
    ReplaceAll *string `xml:"Replace_All,attr,omitempty"`
    DisabilityStatusInformationData *DisabilityStatusInformationData `xml:"Disability_Status_Information_Data,omitempty"`
}

type DisabilityStatusInformationData struct {
    Delete *string `xml:"Delete,attr,omitempty"`
    DisabilityStatusReference *DisabilityStatusReference `xml:"Disability_Status_Reference,omitempty"`
    DisabilityStatusData *DisabilityStatusData `xml:"Disability_Status_Data,omitempty"`
}

type DisabilityStatusReference struct {
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

type MaritalStatusReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type CitizenshipReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PrimaryNationalityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type AdditionalNationalityReference struct {
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

type HispanicOrLatino struct {
    Value *string `xml:",chardata"`
}

type VisualSurveyEthnicityReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type HispanicOrLatinoForVisualSurvey struct {
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

type ReligionReference struct {
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

type NativeRegionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type PersonnelFileAgency struct {
    Value *string `xml:",chardata"`
}

type MilitaryInformationData struct {
    ReplaceAll *string `xml:"Replace_All,attr,omitempty"`
    MilitaryServiceInformationData *MilitaryServiceInformationData `xml:"Military_Service_Information_Data,omitempty"`
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

type PoliticalAffiliationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DateOfDeath struct {
    Value *string `xml:",chardata"`
}

type CityOfBirth struct {
    Value *string `xml:",chardata"`
}

type CityOfBirthReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
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
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type UsesTobacco struct {
    Value *string `xml:",chardata"`
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

type RelativeNameInformationData struct {
    ReplaceAll *string `xml:"Replace_All,attr,omitempty"`
    RelativeNameData *RelativeNameData `xml:"Relative_Name_Data,omitempty"`
}

type RelativeNameData struct {
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

type VeteransPreferenceForRifReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type SelectiveServiceRegistrationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type VeteransPreferenceReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type ActiveMilitaryUniformedServiceReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type DisabledVeteranLeaveDateReference struct {
    Value *string `xml:",chardata"`
}

type UniformServiceReserveStatusReference struct {
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

