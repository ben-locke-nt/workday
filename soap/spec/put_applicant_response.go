package recruiting

import "encoding/xml"

type Id struct {
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"
}

type ApplicantReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type CountryReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type TitleReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type SalutationReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type PrefixData struct {
    TitleReference *TitleReference `xml:"Title_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    TitleDescriptor *string `xml:"Title_Descriptor,omitempty"
    SalutationReference *SalutationReference `xml:"Salutation_Reference,omitempty"
}

type LocalNameDetailData struct {
    Com.workday/bsvc}localname *string `xml:"wd:com.workday/bsvc}Local_Name,attr"
    Com.workday/bsvc}localscript *string `xml:"wd:com.workday/bsvc}Local_Script,attr"
    FirstName *string `xml:"First_Name,omitempty"
    MiddleName *string `xml:"Middle_Name,omitempty"
    LastName *string `xml:"Last_Name,omitempty"
    SecondaryLastName *string `xml:"Secondary_Last_Name,omitempty"
    FirstName2 *string `xml:"First_Name_2,omitempty"
    MiddleName2 *string `xml:"Middle_Name_2,omitempty"
    LastName2 *string `xml:"Last_Name_2,omitempty"
    SecondaryLastName2 *string `xml:"Secondary_Last_Name_2,omitempty"
}

type SocialSuffixReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type AcademicSuffixReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type HereditarySuffixReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type HonorarySuffixReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type ProfessionalSuffixReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type ReligiousSuffixReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type RoyalSuffixReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type SuffixData struct {
    SocialSuffixReference *SocialSuffixReference `xml:"Social_Suffix_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    SocialSuffixDescriptor *string `xml:"Social_Suffix_Descriptor,omitempty"
    AcademicSuffixReference *AcademicSuffixReference `xml:"Academic_Suffix_Reference,omitempty"
    HereditarySuffixReference *HereditarySuffixReference `xml:"Hereditary_Suffix_Reference,omitempty"
    HonorarySuffixReference *HonorarySuffixReference `xml:"Honorary_Suffix_Reference,omitempty"
    ProfessionalSuffixReference *ProfessionalSuffixReference `xml:"Professional_Suffix_Reference,omitempty"
    ReligiousSuffixReference *ReligiousSuffixReference `xml:"Religious_Suffix_Reference,omitempty"
    RoyalSuffixReference *RoyalSuffixReference `xml:"Royal_Suffix_Reference,omitempty"
}

type NameDetailData struct {
    Com.workday/bsvc}formattedname *string `xml:"wd:com.workday/bsvc}Formatted_Name,attr"
    Com.workday/bsvc}reportingname *string `xml:"wd:com.workday/bsvc}Reporting_Name,attr"
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    PrefixData *PrefixData `xml:"Prefix_Data,omitempty"
    FirstName *string `xml:"First_Name,omitempty"
    MiddleName *string `xml:"Middle_Name,omitempty"
    LastName *string `xml:"Last_Name,omitempty"
    SecondaryLastName *string `xml:"Secondary_Last_Name,omitempty"
    TertiaryLastName *string `xml:"Tertiary_Last_Name,omitempty"
    LocalNameDetailData *LocalNameDetailData `xml:"Local_Name_Detail_Data,omitempty"
    Com.workday/bsvc}localName *string `xml:"wd:com.workday/bsvc}Local_Name,attr"`
    Com.workday/bsvc}localScript *string `xml:"wd:com.workday/bsvc}Local_Script,attr"`
    SuffixData *SuffixData `xml:"Suffix_Data,omitempty"
    FullNameForSingaporeAndMalaysia *string `xml:"Full_Name_for_Singapore_and_Malaysia,omitempty"
}

type LegalNameData struct {
    NameDetailData *NameDetailData `xml:"Name_Detail_Data,omitempty"
    Com.workday/bsvc}formattedName *string `xml:"wd:com.workday/bsvc}Formatted_Name,attr"`
    Com.workday/bsvc}reportingName *string `xml:"wd:com.workday/bsvc}Reporting_Name,attr"`
}

type PreferredNameData struct {
    NameDetailData *NameDetailData `xml:"Name_Detail_Data,omitempty"
    Com.workday/bsvc}formattedName *string `xml:"wd:com.workday/bsvc}Formatted_Name,attr"`
    Com.workday/bsvc}reportingName *string `xml:"wd:com.workday/bsvc}Reporting_Name,attr"`
}

type NameTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type AdditionalNameData struct {
    NameDetailData *NameDetailData `xml:"Name_Detail_Data,omitempty"
    Com.workday/bsvc}formattedName *string `xml:"wd:com.workday/bsvc}Formatted_Name,attr"`
    Com.workday/bsvc}reportingName *string `xml:"wd:com.workday/bsvc}Reporting_Name,attr"`
    NameTypeReference *NameTypeReference `xml:"Name_Type_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
}

type NameData struct {
    LegalNameData *LegalNameData `xml:"Legal_Name_Data,omitempty"
    PreferredNameData *PreferredNameData `xml:"Preferred_Name_Data,omitempty"
    AdditionalNameData *AdditionalNameData `xml:"Additional_Name_Data,omitempty"
}

type MaritalStatusReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
    Com.workday/bsvc}parentId *string `xml:"wd:com.workday/bsvc}parent_id,attr"`
    Com.workday/bsvc}parentType *string `xml:"wd:com.workday/bsvc}parent_type,attr"`
}

type ReligionReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type DisabilityReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type DisabilityGradeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type DisabilityCertificationAuthorityReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type DisabilityCertificationBasisReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type WorkerDocumentReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type DisabilityStatusReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type DisabilityStatusData struct {
    DisabilityReference *DisabilityReference `xml:"Disability_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    DisabilityStatusDate *string `xml:"Disability_Status_Date,omitempty"
    DisabilityDateKnown *string `xml:"Disability_Date_Known,omitempty"
    DisabilityEndDate *string `xml:"Disability_End_Date,omitempty"
    DisabilityGradeReference *DisabilityGradeReference `xml:"Disability_Grade_Reference,omitempty"
    DisabilityDegree *string `xml:"Disability_Degree,omitempty"
    DisabilityRemainingCapacity *string `xml:"Disability_Remaining_Capacity,omitempty"
    DisabilityCertificationAuthorityReference *DisabilityCertificationAuthorityReference `xml:"Disability_Certification_Authority_Reference,omitempty"
    DisabilityCertificationAuthority *string `xml:"Disability_Certification_Authority,omitempty"
    DisabilityCertifiedAt *string `xml:"Disability_Certified_At,omitempty"
    DisabilityCertificationId *string `xml:"Disability_Certification_ID,omitempty"
    DisabilityCertificationBasisReference *DisabilityCertificationBasisReference `xml:"Disability_Certification_Basis_Reference,omitempty"
    DisabilitySeverityRecognitionDate *string `xml:"Disability_Severity_Recognition_Date,omitempty"
    DisabilityFteTowardQuota *string `xml:"Disability_FTE_Toward_Quota,omitempty"
    DisabilityWorkRestrictions *string `xml:"Disability_Work_Restrictions,omitempty"
    DisabilityAccommodationsRequested *string `xml:"Disability_Accommodations_Requested,omitempty"
    DisabilityAccommodationsProvided *string `xml:"Disability_Accommodations_Provided,omitempty"
    DisabilityRehabilitationRequested *string `xml:"Disability_Rehabilitation_Requested,omitempty"
    DisabilityRehabilitationProvided *string `xml:"Disability_Rehabilitation_Provided,omitempty"
    Note *string `xml:"Note,omitempty"
    WorkerDocumentReference *WorkerDocumentReference `xml:"Worker_Document_Reference,omitempty"
    DisabilityStatusReference *DisabilityStatusReference `xml:"Disability_Status_Reference,omitempty"
}

type EthnicityReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type RaceEthnicityDetailsReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type EthnicityVisualSurveyReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type AboriginalIndigenousIdentificationReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type AboriginalIndigenousIdentificationDetailsReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type HukouRegionReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type HukouSubregionReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type HukouTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type NativeRegionReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type PoliticalAffiliationReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type SocialBenefitsLocalityReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type RelativeNameReferenceReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type RelativeTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type RelativeNameInformationData struct {
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"
    RelativeNameReferenceReference *RelativeNameReferenceReference `xml:"Relative_Name_Reference_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    RelativeTypeReference *RelativeTypeReference `xml:"Relative_Type_Reference,omitempty"
    NameDetailData *NameDetailData `xml:"Name_Detail_Data,omitempty"
    Com.workday/bsvc}formattedName *string `xml:"wd:com.workday/bsvc}Formatted_Name,attr"`
    Com.workday/bsvc}reportingName *string `xml:"wd:com.workday/bsvc}Reporting_Name,attr"`
}

type SexualOrientationAndGenderIdentityReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type GenderReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type ReportingGenderReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type VeteransPreferenceReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type VeteransPreferenceForRifReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type SelectiveServiceRegistrationReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type ActiveMilitaryUniformedServiceReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type UniformServiceReserveStatusReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type Field1Reference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type Field2Reference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type Field3Reference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type CountrySpecificSection1Data struct {
    Field1Reference *Field1Reference `xml:"Field_1_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    Field2Reference *Field2Reference `xml:"Field_2_Reference,omitempty"
    Field3Reference *Field3Reference `xml:"Field_3_Reference,omitempty"
}

type CountrySpecificSection2Data struct {
    Field1Reference *Field1Reference `xml:"Field_1_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    Field2Reference *Field2Reference `xml:"Field_2_Reference,omitempty"
    Field3Reference *Field3Reference `xml:"Field_3_Reference,omitempty"
}

type CountrySpecificSection3Data struct {
    Field1Reference *Field1Reference `xml:"Field_1_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    Field2Reference *Field2Reference `xml:"Field_2_Reference,omitempty"
    Field3Reference *Field3Reference `xml:"Field_3_Reference,omitempty"
}

type CountrySpecificSectionData struct {
    CountrySpecificSection1Data *CountrySpecificSection1Data `xml:"Country_Specific_Section_1_Data,omitempty"
    CountrySpecificSection2Data *CountrySpecificSection2Data `xml:"Country_Specific_Section_2_Data,omitempty"
    CountrySpecificSection3Data *CountrySpecificSection3Data `xml:"Country_Specific_Section_3_Data,omitempty"
}

type CountryPersonalInformationData struct {
    MaritalStatusReference *MaritalStatusReference `xml:"Marital_Status_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    MaritalStatusDate *string `xml:"Marital_Status_Date,omitempty"
    ReligionReference *ReligionReference `xml:"Religion_Reference,omitempty"
    DisabilityStatusData *DisabilityStatusData `xml:"Disability_Status_Data,omitempty"
    EthnicityReference *EthnicityReference `xml:"Ethnicity_Reference,omitempty"
    RaceEthnicityDetailsReference *RaceEthnicityDetailsReference `xml:"Race_Ethnicity_Details_Reference,omitempty"
    EthnicityVisualSurveyReference *EthnicityVisualSurveyReference `xml:"Ethnicity_Visual_Survey_Reference,omitempty"
    HispanicOrLatino *string `xml:"Hispanic_or_Latino,omitempty"
    HispanicOrLatinoVisualSurvey *string `xml:"Hispanic_or_Latino_Visual_Survey,omitempty"
    AboriginalIndigenousIdentificationReference *AboriginalIndigenousIdentificationReference `xml:"Aboriginal_Indigenous_Identification_Reference,omitempty"
    AboriginalIndigenousIdentificationDetailsReference *AboriginalIndigenousIdentificationDetailsReference `xml:"Aboriginal_Indigenous_Identification_Details_Reference,omitempty"
    HukouRegionReference *HukouRegionReference `xml:"Hukou_Region_Reference,omitempty"
    HukouSubregionReference *HukouSubregionReference `xml:"Hukou_Subregion_Reference,omitempty"
    HukouLocality *string `xml:"Hukou_Locality,omitempty"
    HukouPostalCode *string `xml:"Hukou_Postal_Code,omitempty"
    HukouTypeReference *HukouTypeReference `xml:"Hukou_Type_Reference,omitempty"
    LocalHukou *string `xml:"Local_Hukou,omitempty"
    NativeRegionReference *NativeRegionReference `xml:"Native_Region_Reference,omitempty"
    NativeRegionDescriptor *string `xml:"Native_Region_Descriptor,omitempty"
    PersonnelFileAgencyForPerson *string `xml:"Personnel_File_Agency_for_Person,omitempty"
    PoliticalAffiliationReference *PoliticalAffiliationReference `xml:"Political_Affiliation_Reference,omitempty"
    SocialBenefitsLocalityReference *SocialBenefitsLocalityReference `xml:"Social_Benefits_Locality_Reference,omitempty"
    RelativeNameInformationData *RelativeNameInformationData `xml:"Relative_Name_Information_Data,omitempty"
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"`
    SexualOrientationAndGenderIdentityReference *SexualOrientationAndGenderIdentityReference `xml:"Sexual_Orientation_and_Gender_Identity_Reference,omitempty"
    GenderReference *GenderReference `xml:"Gender_Reference,omitempty"
    ReportingGenderReference *ReportingGenderReference `xml:"Reporting_Gender_Reference,omitempty"
    VeteransPreferenceReference *VeteransPreferenceReference `xml:"Veterans_Preference_Reference,omitempty"
    VeteransPreferenceForRifReference *VeteransPreferenceForRifReference `xml:"Veterans_Preference_For_RIF_Reference,omitempty"
    SelectiveServiceRegistrationReference *SelectiveServiceRegistrationReference `xml:"Selective_Service_Registration_Reference,omitempty"
    DisabledVeteranLeaveDateReference *string `xml:"Disabled_Veteran_Leave_Date_Reference,omitempty"
    ActiveMilitaryUniformedServiceReference *ActiveMilitaryUniformedServiceReference `xml:"Active_Military_Uniformed_Service_Reference,omitempty"
    UniformServiceReserveStatusReference *UniformServiceReserveStatusReference `xml:"Uniform_Service_Reserve_Status_Reference,omitempty"
    CountrySpecificSectionData *CountrySpecificSectionData `xml:"Country_Specific_Section_Data,omitempty"
}

type PersonalInformationForCountryData struct {
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    CountryPersonalInformationData *CountryPersonalInformationData `xml:"Country_Personal_Information_Data,omitempty"
}

type CountryOfBirthReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type CountryRegionOfBirthReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type CityOfBirthReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
    Com.workday/bsvc}parentId *string `xml:"wd:com.workday/bsvc}parent_id,attr"`
    Com.workday/bsvc}parentType *string `xml:"wd:com.workday/bsvc}parent_type,attr"`
}

type CitizenshipStatusReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type PrimaryNationalityReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type AdditionalNationalitiesReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type BloodTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type StatusReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type MilitaryServiceTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type MilitaryRankReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type MilitaryServiceReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type MilitaryDischargeTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type MilitaryServiceData struct {
    StatusReference *StatusReference `xml:"Status_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    DischargeDate *string `xml:"Discharge_Date,omitempty"
    StatusBeginDate *string `xml:"Status_Begin_Date,omitempty"
    MilitaryServiceTypeReference *MilitaryServiceTypeReference `xml:"Military_Service_Type_Reference,omitempty"
    MilitaryRankReference *MilitaryRankReference `xml:"Military_Rank_Reference,omitempty"
    Notes *string `xml:"Notes,omitempty"
    MilitaryServiceReference *MilitaryServiceReference `xml:"Military_Service_Reference,omitempty"
    MilitaryDischargeTypeReference *MilitaryDischargeTypeReference `xml:"Military_Discharge_Type_Reference,omitempty"
}

type SexualOrientationReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type GenderIdentityReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type PronounReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type NonCountrySpecificSection1Data struct {
    Field1Reference *Field1Reference `xml:"Field_1_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    Field2Reference *Field2Reference `xml:"Field_2_Reference,omitempty"
    Field3Reference *Field3Reference `xml:"Field_3_Reference,omitempty"
}

type NonCountrySpecificSection2Data struct {
    Field1Reference *Field1Reference `xml:"Field_1_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    Field2Reference *Field2Reference `xml:"Field_2_Reference,omitempty"
    Field3Reference *Field3Reference `xml:"Field_3_Reference,omitempty"
}

type NonCountrySpecificSection3Data struct {
    Field1Reference *Field1Reference `xml:"Field_1_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    Field2Reference *Field2Reference `xml:"Field_2_Reference,omitempty"
    Field3Reference *Field3Reference `xml:"Field_3_Reference,omitempty"
}

type NonCountrySpecificSectionData struct {
    NonCountrySpecificSection1Data *NonCountrySpecificSection1Data `xml:"Non_Country_Specific_Section_1_Data,omitempty"
    NonCountrySpecificSection2Data *NonCountrySpecificSection2Data `xml:"Non_Country_Specific_Section_2_Data,omitempty"
    NonCountrySpecificSection3Data *NonCountrySpecificSection3Data `xml:"Non_Country_Specific_Section_3_Data,omitempty"
}

type PersonalInformationData struct {
    PersonalInformationForCountryData *PersonalInformationForCountryData `xml:"Personal_Information_For_Country_Data,omitempty"
    BirthDate *string `xml:"Birth_Date,omitempty"
    DateOfDeath *string `xml:"Date_of_Death,omitempty"
    CountryOfBirthReference *CountryOfBirthReference `xml:"Country_of_Birth_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    CountryRegionOfBirthReference *CountryRegionOfBirthReference `xml:"Country_Region_of_Birth_Reference,omitempty"
    RegionOfBirthDescriptor *string `xml:"Region_of_Birth_Descriptor,omitempty"
    CityOfBirth *string `xml:"City_of_Birth,omitempty"
    CityOfBirthReference *CityOfBirthReference `xml:"City_of_Birth_Reference,omitempty"
    CitizenshipStatusReference *CitizenshipStatusReference `xml:"Citizenship_Status_Reference,omitempty"
    PrimaryNationalityReference *PrimaryNationalityReference `xml:"Primary_Nationality_Reference,omitempty"
    AdditionalNationalitiesReference *AdditionalNationalitiesReference `xml:"Additional_Nationalities_Reference,omitempty"
    LastMedicalExamValidTo *string `xml:"Last_Medical_Exam_Valid_To,omitempty"
    LastMedicalExamDate *string `xml:"Last_Medical_Exam_Date,omitempty"
    MedicalExamNotes *string `xml:"Medical_Exam_Notes,omitempty"
    BloodTypeReference *BloodTypeReference `xml:"Blood_Type_Reference,omitempty"
    MilitaryServiceData *MilitaryServiceData `xml:"Military_Service_Data,omitempty"
    SexualOrientationReference *SexualOrientationReference `xml:"Sexual_Orientation_Reference,omitempty"
    GenderIdentityReference *GenderIdentityReference `xml:"Gender_Identity_Reference,omitempty"
    PronounReference *PronounReference `xml:"Pronoun_Reference,omitempty"
    NonCountrySpecificSectionData *NonCountrySpecificSectionData `xml:"Non_Country_Specific_Section_Data,omitempty"
}

type NationalIdReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type IdTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type VerifiedByReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type NationalIdData struct {
    Id *string `xml:"ID,omitempty"
    IdTypeReference *IdTypeReference `xml:"ID_Type_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"
    IssuedDate *string `xml:"Issued_Date,omitempty"
    ExpirationDate *string `xml:"Expiration_Date,omitempty"
    VerificationDate *string `xml:"Verification_Date,omitempty"
    Series *string `xml:"Series,omitempty"
    IssuingAgency *string `xml:"Issuing_Agency,omitempty"
    VerifiedByReference *VerifiedByReference `xml:"Verified_By_Reference,omitempty"
}

type NationalIdSharedReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type NationalId struct {
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"
    NationalIdReference *NationalIdReference `xml:"National_ID_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    NationalIdData *NationalIdData `xml:"National_ID_Data,omitempty"
    NationalIdSharedReference *NationalIdSharedReference `xml:"National_ID_Shared_Reference,omitempty"
}

type GovernmentIdReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type GovernmentIdData struct {
    Id *string `xml:"ID,omitempty"
    IdTypeReference *IdTypeReference `xml:"ID_Type_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"
    IssuedDate *string `xml:"Issued_Date,omitempty"
    ExpirationDate *string `xml:"Expiration_Date,omitempty"
    VerificationDate *string `xml:"Verification_Date,omitempty"
    VerifiedByReference *VerifiedByReference `xml:"Verified_by_Reference,omitempty"
}

type GovernmentIdSharedReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type GovernmentId struct {
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"
    GovernmentIdReference *GovernmentIdReference `xml:"Government_ID_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    GovernmentIdData *GovernmentIdData `xml:"Government_ID_Data,omitempty"
    GovernmentIdSharedReference *GovernmentIdSharedReference `xml:"Government_ID_Shared_Reference,omitempty"
}

type VisaIdReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type VisaIdData struct {
    Id *string `xml:"ID,omitempty"
    IdTypeReference *IdTypeReference `xml:"ID_Type_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"
    IssuedDate *string `xml:"Issued_Date,omitempty"
    ExpirationDate *string `xml:"Expiration_Date,omitempty"
    VerificationDate *string `xml:"Verification_Date,omitempty"
    VerifiedByReference *VerifiedByReference `xml:"Verified_By_Reference,omitempty"
}

type VisaIdSharedReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type VisaId struct {
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"
    VisaIdReference *VisaIdReference `xml:"Visa_ID_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    VisaIdData *VisaIdData `xml:"Visa_ID_Data,omitempty"
    VisaIdSharedReference *VisaIdSharedReference `xml:"Visa_ID_Shared_Reference,omitempty"
}

type PassportIdReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type PassportIdData struct {
    Id *string `xml:"ID,omitempty"
    IdTypeReference *IdTypeReference `xml:"ID_Type_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"
    IssuedDate *string `xml:"Issued_Date,omitempty"
    ExpirationDate *string `xml:"Expiration_Date,omitempty"
    VerificationDate *string `xml:"Verification_Date,omitempty"
    VerifiedByReference *VerifiedByReference `xml:"Verified_By_Reference,omitempty"
}

type PassportIdSharedReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type PassportId struct {
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"
    PassportIdReference *PassportIdReference `xml:"Passport_ID_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    PassportIdData *PassportIdData `xml:"Passport_ID_Data,omitempty"
    PassportIdSharedReference *PassportIdSharedReference `xml:"Passport_ID_Shared_Reference,omitempty"
}

type LicenseIdReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type CountryRegionReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type AuthorityReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type LicenseIdData struct {
    Id *string `xml:"ID,omitempty"
    IdTypeReference *IdTypeReference `xml:"ID_Type_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"
    CountryRegionReference *CountryRegionReference `xml:"Country_Region_Reference,omitempty"
    CountryRegionDescriptor *string `xml:"Country_Region_Descriptor,omitempty"
    AuthorityReference *AuthorityReference `xml:"Authority_Reference,omitempty"
    LicenseClass *string `xml:"License_Class,omitempty"
    IssuedDate *string `xml:"Issued_Date,omitempty"
    ExpirationDate *string `xml:"Expiration_Date,omitempty"
    VerificationDate *string `xml:"Verification_Date,omitempty"
    VerifiedByReference *VerifiedByReference `xml:"Verified_By_Reference,omitempty"
}

type LicenseIdSharedReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type LicenseId struct {
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"
    LicenseIdReference *LicenseIdReference `xml:"License_ID_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    LicenseIdData *LicenseIdData `xml:"License_ID_Data,omitempty"
    LicenseIdSharedReference *LicenseIdSharedReference `xml:"License_ID_Shared_Reference,omitempty"
}

type CustomIdReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type OrganizationIdReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type CustomIdData struct {
    Id *string `xml:"ID,omitempty"
    IdTypeReference *IdTypeReference `xml:"ID_Type_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    IssuedDate *string `xml:"Issued_Date,omitempty"
    ExpirationDate *string `xml:"Expiration_Date,omitempty"
    OrganizationIdReference *OrganizationIdReference `xml:"Organization_ID_Reference,omitempty"
    CustomDescription *string `xml:"Custom_Description,omitempty"
}

type CustomIdSharedReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type CustomId struct {
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"
    CustomIdReference *CustomIdReference `xml:"Custom_ID_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    CustomIdData *CustomIdData `xml:"Custom_ID_Data,omitempty"
    CustomIdSharedReference *CustomIdSharedReference `xml:"Custom_ID_Shared_Reference,omitempty"
}

type IdentificationData struct {
    NationalId *NationalId `xml:"National_ID,omitempty"
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"`
    GovernmentId *GovernmentId `xml:"Government_ID,omitempty"
    VisaId *VisaId `xml:"Visa_ID,omitempty"
    PassportId *PassportId `xml:"Passport_ID,omitempty"
    LicenseId *LicenseId `xml:"License_ID,omitempty"
    CustomId *CustomId `xml:"Custom_ID,omitempty"
}

type AddressLineData struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}Type,attr"
}

type CountryCityReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
    Com.workday/bsvc}parentId *string `xml:"wd:com.workday/bsvc}parent_id,attr"`
    Com.workday/bsvc}parentType *string `xml:"wd:com.workday/bsvc}parent_type,attr"`
}

type SubmunicipalityData struct {
    Com.workday/bsvc}addresscomponentname *string `xml:"wd:com.workday/bsvc}Address_Component_Name,attr"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}Type,attr"
}

type SubregionData struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}Type,attr"
}

type TypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type TypeData struct {
    Com.workday/bsvc}primary *string `xml:"wd:com.workday/bsvc}Primary,attr"
    TypeReference *TypeReference `xml:"Type_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
}

type UseForReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type UseForTenantedReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type UsageData struct {
    Com.workday/bsvc}public *string `xml:"wd:com.workday/bsvc}Public,attr"
    TypeData *TypeData `xml:"Type_Data,omitempty"
    Com.workday/bsvc}primary *string `xml:"wd:com.workday/bsvc}Primary,attr"`
    UseForReference *UseForReference `xml:"Use_For_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    UseForTenantedReference *UseForTenantedReference `xml:"Use_For_Tenanted_Reference,omitempty"
    Comments *string `xml:"Comments,omitempty"
}

type AddressReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type AddressData struct {
    Com.workday/bsvc}formattedaddress *string `xml:"wd:com.workday/bsvc}Formatted_Address,attr"
    Com.workday/bsvc}addressformattype *string `xml:"wd:com.workday/bsvc}Address_Format_Type,attr"
    Com.workday/bsvc}defaultedbusinesssiteaddress *string `xml:"wd:com.workday/bsvc}Defaulted_Business_Site_Address,attr"
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"
    Com.workday/bsvc}donotreplaceall *string `xml:"wd:com.workday/bsvc}Do_Not_Replace_All,attr"
    Com.workday/bsvc}effectivedate *string `xml:"wd:com.workday/bsvc}Effective_Date,attr"
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    LastModified *string `xml:"Last_Modified,omitempty"
    AddressLineData *AddressLineData `xml:"Address_Line_Data,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}Type,attr"`
    Municipality *string `xml:"Municipality,omitempty"
    CountryCityReference *CountryCityReference `xml:"Country_City_Reference,omitempty"
    SubmunicipalityData *SubmunicipalityData `xml:"Submunicipality_Data,omitempty"
    Com.workday/bsvc}addressComponentName *string `xml:"wd:com.workday/bsvc}Address_Component_Name,attr"`
    CountryRegionReference *CountryRegionReference `xml:"Country_Region_Reference,omitempty"
    CountryRegionDescriptor *string `xml:"Country_Region_Descriptor,omitempty"
    SubregionData *SubregionData `xml:"Subregion_Data,omitempty"
    PostalCode *string `xml:"Postal_Code,omitempty"
    UsageData *UsageData `xml:"Usage_Data,omitempty"
    Com.workday/bsvc}public *string `xml:"wd:com.workday/bsvc}Public,attr"`
    NumberOfDays *string `xml:"Number_of_Days,omitempty"
    MunicipalityLocal *string `xml:"Municipality_Local,omitempty"
    AddressReference *AddressReference `xml:"Address_Reference,omitempty"
    AddressId *string `xml:"Address_ID,omitempty"
}

type PhoneDeviceTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type PhoneReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type PhoneData struct {
    Com.workday/bsvc}areacode *string `xml:"wd:com.workday/bsvc}Area_Code,attr"
    Com.workday/bsvc}tenantformattedphone *string `xml:"wd:com.workday/bsvc}Tenant_Formatted_Phone,attr"
    Com.workday/bsvc}internationalformattedphone *string `xml:"wd:com.workday/bsvc}International_Formatted_Phone,attr"
    Com.workday/bsvc}phonenumberwithoutareacode *string `xml:"wd:com.workday/bsvc}Phone_Number_Without_Area_Code,attr"
    Com.workday/bsvc}nationalformattedphone *string `xml:"wd:com.workday/bsvc}National_Formatted_Phone,attr"
    Com.workday/bsvc}e164formattedphone *string `xml:"wd:com.workday/bsvc}E164_Formatted_Phone,attr"
    Com.workday/bsvc}workdaytraditionalformattedphone *string `xml:"wd:com.workday/bsvc}Workday_Traditional_Formatted_Phone,attr"
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"
    Com.workday/bsvc}donotreplaceall *string `xml:"wd:com.workday/bsvc}Do_Not_Replace_All,attr"
    CountryIsoCode *string `xml:"Country_ISO_Code,omitempty"
    InternationalPhoneCode *string `xml:"International_Phone_Code,omitempty"
    PhoneNumber *string `xml:"Phone_Number,omitempty"
    PhoneExtension *string `xml:"Phone_Extension,omitempty"
    PhoneDeviceTypeReference *PhoneDeviceTypeReference `xml:"Phone_Device_Type_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    UsageData *UsageData `xml:"Usage_Data,omitempty"
    Com.workday/bsvc}public *string `xml:"wd:com.workday/bsvc}Public,attr"`
    PhoneReference *PhoneReference `xml:"Phone_Reference,omitempty"
    Id *string `xml:"ID,omitempty"
}

type EmailReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type EmailAddressData struct {
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"
    Com.workday/bsvc}donotreplaceall *string `xml:"wd:com.workday/bsvc}Do_Not_Replace_All,attr"
    EmailAddress *string `xml:"Email_Address,omitempty"
    EmailComment *string `xml:"Email_Comment,omitempty"
    UsageData *UsageData `xml:"Usage_Data,omitempty"
    Com.workday/bsvc}public *string `xml:"wd:com.workday/bsvc}Public,attr"`
    EmailReference *EmailReference `xml:"Email_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    Id *string `xml:"ID,omitempty"
}

type InstantMessengerTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type InstantMessengerReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type InstantMessengerData struct {
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"
    Com.workday/bsvc}donotreplaceall *string `xml:"wd:com.workday/bsvc}Do_Not_Replace_All,attr"
    InstantMessengerAddress *string `xml:"Instant_Messenger_Address,omitempty"
    InstantMessengerTypeReference *InstantMessengerTypeReference `xml:"Instant_Messenger_Type_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    InstantMessengerComment *string `xml:"Instant_Messenger_Comment,omitempty"
    UsageData *UsageData `xml:"Usage_Data,omitempty"
    Com.workday/bsvc}public *string `xml:"wd:com.workday/bsvc}Public,attr"`
    InstantMessengerReference *InstantMessengerReference `xml:"Instant_Messenger_Reference,omitempty"
    Id *string `xml:"ID,omitempty"
}

type WebAddressReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type WebAddressData struct {
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"
    Com.workday/bsvc}donotreplaceall *string `xml:"wd:com.workday/bsvc}Do_Not_Replace_All,attr"
    WebAddress *string `xml:"Web_Address,omitempty"
    WebAddressComment *string `xml:"Web_Address_Comment,omitempty"
    UsageData *UsageData `xml:"Usage_Data,omitempty"
    Com.workday/bsvc}public *string `xml:"wd:com.workday/bsvc}Public,attr"`
    WebAddressReference *WebAddressReference `xml:"Web_Address_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    Id *string `xml:"ID,omitempty"
}

type ContactData struct {
    AddressData *AddressData `xml:"Address_Data,omitempty"
    Com.workday/bsvc}formattedAddress *string `xml:"wd:com.workday/bsvc}Formatted_Address,attr"`
    Com.workday/bsvc}addressFormatType *string `xml:"wd:com.workday/bsvc}Address_Format_Type,attr"`
    Com.workday/bsvc}defaultedBusinessSiteAddress *string `xml:"wd:com.workday/bsvc}Defaulted_Business_Site_Address,attr"`
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"`
    Com.workday/bsvc}doNotReplaceAll *string `xml:"wd:com.workday/bsvc}Do_Not_Replace_All,attr"`
    Com.workday/bsvc}effectiveDate *string `xml:"wd:com.workday/bsvc}Effective_Date,attr"`
    PhoneData *PhoneData `xml:"Phone_Data,omitempty"
    Com.workday/bsvc}areaCode *string `xml:"wd:com.workday/bsvc}Area_Code,attr"`
    Com.workday/bsvc}tenantFormattedPhone *string `xml:"wd:com.workday/bsvc}Tenant_Formatted_Phone,attr"`
    Com.workday/bsvc}internationalFormattedPhone *string `xml:"wd:com.workday/bsvc}International_Formatted_Phone,attr"`
    Com.workday/bsvc}phoneNumberWithoutAreaCode *string `xml:"wd:com.workday/bsvc}Phone_Number_Without_Area_Code,attr"`
    Com.workday/bsvc}nationalFormattedPhone *string `xml:"wd:com.workday/bsvc}National_Formatted_Phone,attr"`
    Com.workday/bsvc}e164FormattedPhone *string `xml:"wd:com.workday/bsvc}E164_Formatted_Phone,attr"`
    Com.workday/bsvc}workdayTraditionalFormattedPhone *string `xml:"wd:com.workday/bsvc}Workday_Traditional_Formatted_Phone,attr"`
    EmailAddressData *EmailAddressData `xml:"Email_Address_Data,omitempty"
    InstantMessengerData *InstantMessengerData `xml:"Instant_Messenger_Data,omitempty"
    WebAddressData *WebAddressData `xml:"Web_Address_Data,omitempty"
}

type PersonalData struct {
    UniversalId *string `xml:"Universal_ID,omitempty"
    NameData *NameData `xml:"Name_Data,omitempty"
    PersonalInformationData *PersonalInformationData `xml:"Personal_Information_Data,omitempty"
    IdentificationData *IdentificationData `xml:"Identification_Data,omitempty"
    ContactData *ContactData `xml:"Contact_Data,omitempty"
    TobaccoUse *string `xml:"Tobacco_Use,omitempty"
}

type JobHistoryReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type JobHistoryCompanyReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type JobHistoryData struct {
    JobHistoryId *string `xml:"Job_History_ID,omitempty"
    RemoveJobHistory *string `xml:"Remove_Job_History,omitempty"
    JobTitle *string `xml:"Job_Title,omitempty"
    Company *string `xml:"Company,omitempty"
    JobHistoryCompanyReference *JobHistoryCompanyReference `xml:"Job_History_Company_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    StartDate *string `xml:"Start_Date,omitempty"
    EndDate *string `xml:"End_Date,omitempty"
    ResponsibilitiesAndAchievements *string `xml:"Responsibilities_and_Achievements,omitempty"
    Location *string `xml:"Location,omitempty"
    JobReference *string `xml:"Job_Reference,omitempty"
    Contact *string `xml:"Contact,omitempty"
}

type ExternalJobHistory struct {
    JobHistoryReference *JobHistoryReference `xml:"Job_History_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    JobHistoryData *JobHistoryData `xml:"Job_History_Data,omitempty"
}

type CompetencyLevelBehaviorReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type AssessedByPersonReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type CompetencyReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type Competency struct {
    CompetencyLevelBehaviorReference *CompetencyLevelBehaviorReference `xml:"Competency_Level_Behavior_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    CompetencyLevelBehaviorDescriptor *string `xml:"Competency_Level_Behavior_Descriptor,omitempty"
    AssessmentComment *string `xml:"Assessment_Comment,omitempty"
    AssessedOn *string `xml:"Assessed_On,omitempty"
    AssessedByPersonReference *AssessedByPersonReference `xml:"Assessed_By_Person_Reference,omitempty"
    AssessedByWorkerDescriptor *string `xml:"Assessed_By_Worker_Descriptor,omitempty"
    CompetencyReference *CompetencyReference `xml:"Competency_Reference,omitempty"
    CompetencyDescriptor *string `xml:"Competency_Descriptor,omitempty"
}

type CertificationReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type SpecialtyReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type SubspecialtyReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type SpecialtyAchievementData struct {
    SpecialtyReference *SpecialtyReference `xml:"Specialty_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    StartDate *string `xml:"Start_Date,omitempty"
    EndDate *string `xml:"End_Date,omitempty"
    SubspecialtyReference *SubspecialtyReference `xml:"Subspecialty_Reference,omitempty"
}

type DocumentCategoryReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type WorkerDocumentData struct {
    FileName *string `xml:"File_Name,omitempty"
    Comment *string `xml:"Comment,omitempty"
    File *string `xml:"File,omitempty"
    DocumentCategoryReference *DocumentCategoryReference `xml:"Document_Category_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    ContentType *string `xml:"Content_Type,omitempty"
}

type CertificationData struct {
    CertificationId *string `xml:"Certification_ID,omitempty"
    RemoveCertification *string `xml:"Remove_Certification,omitempty"
    CertificationReference *CertificationReference `xml:"Certification_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"
    CertificationName *string `xml:"Certification_Name,omitempty"
    Issuer *string `xml:"Issuer,omitempty"
    CertificationNumber *string `xml:"Certification_Number,omitempty"
    IssuedDate *string `xml:"Issued_Date,omitempty"
    ExpirationDate *string `xml:"Expiration_Date,omitempty"
    ExaminationScore *string `xml:"Examination_Score,omitempty"
    ExaminationDate *string `xml:"Examination_Date,omitempty"
    SpecialtyAchievementData *SpecialtyAchievementData `xml:"Specialty_Achievement_Data,omitempty"
    WorkerDocumentData *WorkerDocumentData `xml:"Worker_Document_Data,omitempty"
}

type Certification struct {
    CertificationReference *CertificationReference `xml:"Certification_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    CertificationData *CertificationData `xml:"Certification_Data,omitempty"
}

type TrainingReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type TrainingTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type TrainingData struct {
    TrainingId *string `xml:"Training_ID,omitempty"
    RemoveTraining *string `xml:"Remove_Training,omitempty"
    Training *string `xml:"Training,omitempty"
    Description *string `xml:"Description,omitempty"
    TrainingTypeReference *TrainingTypeReference `xml:"Training_Type_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    CompletionDate *string `xml:"Completion_Date,omitempty"
    TrainingDuration *string `xml:"Training_Duration,omitempty"
}

type Training struct {
    TrainingReference *TrainingReference `xml:"Training_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    TrainingData *TrainingData `xml:"Training_Data,omitempty"
}

type AwardAndActivityReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type AwardAndActivityTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type RelatedPositionReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type AwardAndActivityData struct {
    AwardAndActivityId *string `xml:"Award_and_Activity_ID,omitempty"
    RemoveAwardAndActivity *string `xml:"Remove_Award_and_Activity,omitempty"
    AwardAndActivityTypeReference *AwardAndActivityTypeReference `xml:"Award_and_Activity_Type_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    Title *string `xml:"Title,omitempty"
    SponsorIssuer *string `xml:"Sponsor_Issuer,omitempty"
    StartDate *string `xml:"Start_Date,omitempty"
    EndDate *string `xml:"End_Date,omitempty"
    Description *string `xml:"Description,omitempty"
    Url *string `xml:"URL,omitempty"
    RelatedPositionReference *RelatedPositionReference `xml:"Related_Position_Reference,omitempty"
}

type AwardAndActivity struct {
    AwardAndActivityReference *AwardAndActivityReference `xml:"Award_and_Activity_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    AwardAndActivityData *AwardAndActivityData `xml:"Award_and_Activity_Data,omitempty"
}

type OrganizationProfessionalAffiliationReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type ProfessionalAffiliationReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type ProfessionalAffiliationTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type ProfessionalAffiliationRelationshipTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type ContactInformationData struct {
    AddressData *AddressData `xml:"Address_Data,omitempty"
    Com.workday/bsvc}formattedAddress *string `xml:"wd:com.workday/bsvc}Formatted_Address,attr"`
    Com.workday/bsvc}addressFormatType *string `xml:"wd:com.workday/bsvc}Address_Format_Type,attr"`
    Com.workday/bsvc}defaultedBusinessSiteAddress *string `xml:"wd:com.workday/bsvc}Defaulted_Business_Site_Address,attr"`
    Com.workday/bsvc}delete *string `xml:"wd:com.workday/bsvc}Delete,attr"`
    Com.workday/bsvc}doNotReplaceAll *string `xml:"wd:com.workday/bsvc}Do_Not_Replace_All,attr"`
    Com.workday/bsvc}effectiveDate *string `xml:"wd:com.workday/bsvc}Effective_Date,attr"`
    PhoneData *PhoneData `xml:"Phone_Data,omitempty"
    Com.workday/bsvc}areaCode *string `xml:"wd:com.workday/bsvc}Area_Code,attr"`
    Com.workday/bsvc}tenantFormattedPhone *string `xml:"wd:com.workday/bsvc}Tenant_Formatted_Phone,attr"`
    Com.workday/bsvc}internationalFormattedPhone *string `xml:"wd:com.workday/bsvc}International_Formatted_Phone,attr"`
    Com.workday/bsvc}phoneNumberWithoutAreaCode *string `xml:"wd:com.workday/bsvc}Phone_Number_Without_Area_Code,attr"`
    Com.workday/bsvc}nationalFormattedPhone *string `xml:"wd:com.workday/bsvc}National_Formatted_Phone,attr"`
    Com.workday/bsvc}e164FormattedPhone *string `xml:"wd:com.workday/bsvc}E164_Formatted_Phone,attr"`
    Com.workday/bsvc}workdayTraditionalFormattedPhone *string `xml:"wd:com.workday/bsvc}Workday_Traditional_Formatted_Phone,attr"`
    EmailAddressData *EmailAddressData `xml:"Email_Address_Data,omitempty"
    InstantMessengerData *InstantMessengerData `xml:"Instant_Messenger_Data,omitempty"
    WebAddressData *WebAddressData `xml:"Web_Address_Data,omitempty"
}

type OrganizationProfessionalAffiliationData struct {
    ProfessionalAffiliationId *string `xml:"Professional_Affiliation_ID,omitempty"
    RemoveProfessionalAffiliation *string `xml:"Remove_Professional_Affiliation,omitempty"
    ProfessionalAffiliationReference *ProfessionalAffiliationReference `xml:"Professional_Affiliation_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    ProfessionalAffiliation *string `xml:"Professional_Affiliation,omitempty"
    ProfessionalAffiliationTypeReference *ProfessionalAffiliationTypeReference `xml:"Professional_Affiliation_Type_Reference,omitempty"
    Affiliation *string `xml:"Affiliation,omitempty"
    ProfessionalAffiliationRelationshipTypeReference *ProfessionalAffiliationRelationshipTypeReference `xml:"Professional_Affiliation_Relationship_Type_Reference,omitempty"
    BeginDate *string `xml:"Begin_Date,omitempty"
    EndDate *string `xml:"End_Date,omitempty"
    ContactInformationData *ContactInformationData `xml:"Contact_Information_Data,omitempty"
}

type OrganizationMembership struct {
    OrganizationProfessionalAffiliationReference *OrganizationProfessionalAffiliationReference `xml:"Organization_Professional_Affiliation_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    OrganizationProfessionalAffiliationData *OrganizationProfessionalAffiliationData `xml:"Organization_Professional_Affiliation_Data,omitempty"
}

type EducationReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type SchoolReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type SchoolTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type DegreeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type DegreeCompletedReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type FieldOfStudyReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type EducationAttachmentData struct {
    FileName *string `xml:"File_Name,omitempty"
    Comment *string `xml:"Comment,omitempty"
    File *string `xml:"File,omitempty"
    DocumentCategoryReference *DocumentCategoryReference `xml:"Document_Category_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
}

type EducationData struct {
    EducationId *string `xml:"Education_ID,omitempty"
    RemoveEducation *string `xml:"Remove_Education,omitempty"
    CountryReference *CountryReference `xml:"Country_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    SchoolReference *SchoolReference `xml:"School_Reference,omitempty"
    SchoolName *string `xml:"School_Name,omitempty"
    SchoolTypeReference *SchoolTypeReference `xml:"School_Type_Reference,omitempty"
    Location *string `xml:"Location,omitempty"
    DegreeReference *DegreeReference `xml:"Degree_Reference,omitempty"
    DegreeCompletedReference *DegreeCompletedReference `xml:"Degree_Completed_Reference,omitempty"
    DateDegreeReceived *string `xml:"Date_Degree_Received,omitempty"
    FieldOfStudyReference *FieldOfStudyReference `xml:"Field_Of_Study_Reference,omitempty"
    GradeAverage *string `xml:"Grade_Average,omitempty"
    FirstYearAttended *string `xml:"First_Year_Attended,omitempty"
    LastYearAttended *string `xml:"Last_Year_Attended,omitempty"
    IsHighestLevelOfEducation *string `xml:"Is_Highest_Level_of_Education,omitempty"
    FirstDayAttended *string `xml:"First_Day_Attended,omitempty"
    LastDayAttended *string `xml:"Last_Day_Attended,omitempty"
    EducationAttachmentData *EducationAttachmentData `xml:"Education_Attachment_Data,omitempty"
}

type Education struct {
    EducationReference *EducationReference `xml:"Education_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    EducationData *EducationData `xml:"Education_Data,omitempty"
}

type ExperienceReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type ExperienceRatingReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type WorkExperience struct {
    ExperienceReference *ExperienceReference `xml:"Experience_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    RemoveExperience *string `xml:"Remove_Experience,omitempty"
    ExperienceRatingReference *ExperienceRatingReference `xml:"Experience_Rating_Reference,omitempty"
    ExperienceComment *string `xml:"Experience_Comment,omitempty"
}

type LanguageReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type LanguageProficiencyReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type LanguageAbilityTypeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type LanguageAbilityData struct {
    LanguageProficiencyReference *LanguageProficiencyReference `xml:"Language_Proficiency_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    LanguageAbilityTypeReference *LanguageAbilityTypeReference `xml:"Language_Ability_Type_Reference,omitempty"
}

type LanguageAbility struct {
    LanguageAbilityData *LanguageAbilityData `xml:"Language_Ability_Data,omitempty"
}

type AssessedByWorkerReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type Language struct {
    LanguageReference *LanguageReference `xml:"Language_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    RemoveLanguage *string `xml:"Remove_Language,omitempty"
    NativeLanguage *string `xml:"Native_Language,omitempty"
    LanguageAbility *LanguageAbility `xml:"Language_Ability,omitempty"
    AssessedOn *string `xml:"Assessed_On,omitempty"
    Note *string `xml:"Note,omitempty"
    AssessedByWorkerReference *AssessedByWorkerReference `xml:"Assessed_by_Worker_Reference,omitempty"
}

type InternalProjectExperienceReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type InternalProjectExperienceData struct {
    InternalProjectExperienceId *string `xml:"Internal_Project_Experience_ID,omitempty"
    RemoveInternalProjectExperience *string `xml:"Remove_Internal_Project_Experience,omitempty"
    InternalProject *string `xml:"Internal_Project,omitempty"
    Description *string `xml:"Description,omitempty"
    StartDate *string `xml:"Start_Date,omitempty"
    EndDate *string `xml:"End_Date,omitempty"
    ProjectLeader *string `xml:"Project_Leader,omitempty"
}

type InternalProjectExperience struct {
    InternalProjectExperienceReference *InternalProjectExperienceReference `xml:"Internal_Project_Experience_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    InternalProjectExperienceData *InternalProjectExperienceData `xml:"Internal_Project_Experience_Data,omitempty"
}

type QualificationData struct {
    ExternalJobHistory *ExternalJobHistory `xml:"External_Job_History,omitempty"
    Competency *Competency `xml:"Competency,omitempty"
    Certification *Certification `xml:"Certification,omitempty"
    Training *Training `xml:"Training,omitempty"
    AwardAndActivity *AwardAndActivity `xml:"Award_and_Activity,omitempty"
    OrganizationMembership *OrganizationMembership `xml:"Organization_Membership,omitempty"
    Education *Education `xml:"Education,omitempty"
    WorkExperience *WorkExperience `xml:"Work_Experience,omitempty"
    Language *Language `xml:"Language,omitempty"
    InternalProjectExperience *InternalProjectExperience `xml:"Internal_Project_Experience,omitempty"
}

type EligibleForHireReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type ApplicantHasMarkedAsNoShowReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type ApplicantSourceReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type ReferredByWorkerReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type PositionsConsideredForReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type RecruitingData struct {
    ApplicantEnteredDate *string `xml:"Applicant_Entered_Date,omitempty"
    ApplicantComments *string `xml:"Applicant_Comments,omitempty"
    EligibleForHireReference *EligibleForHireReference `xml:"Eligible_For_Hire_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    EligibleForRehireComments *string `xml:"Eligible_for_Rehire_Comments,omitempty"
    ApplicantHasMarkedAsNoShowReference *ApplicantHasMarkedAsNoShowReference `xml:"Applicant_Has_Marked_as_No_Show_Reference,omitempty"
    ApplicantSourceReference *ApplicantSourceReference `xml:"Applicant_Source_Reference,omitempty"
    ReferredByWorkerReference *ReferredByWorkerReference `xml:"Referred_by_Worker_Reference,omitempty"
    PositionsConsideredForReference *PositionsConsideredForReference `xml:"Positions_Considered_for_Reference,omitempty"
}

type ResumeReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
}

type ResumeData struct {
    Resume *Resume `xml:"Resume,omitempty"
}

type Resume struct {
    ResumeReference *ResumeReference `xml:"Resume_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    ResumeData *ResumeData `xml:"Resume_Data,omitempty"
}

type BackgroundCheckData struct {
    StatusDate *string `xml:"Status_Date,omitempty"
    StatusReference *StatusReference `xml:"Status_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    StatusComment *string `xml:"Status_Comment,omitempty"
}

type ExternalIntegrationIdData struct {
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}systemId *string `xml:"wd:com.workday/bsvc}System_ID,attr"`
}

type FieldReference struct {
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"
    Id *Id `xml:"ID,omitempty"
    Com.workday/bsvc}type *string `xml:"wd:com.workday/bsvc}type,attr"`
    Com.workday/bsvc}parentId *string `xml:"wd:com.workday/bsvc}parent_id,attr"`
    Com.workday/bsvc}parentType *string `xml:"wd:com.workday/bsvc}parent_type,attr"`
}

type DocumentFieldResultData struct {
    FieldReference *FieldReference `xml:"Field_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    Value *string `xml:"Value,omitempty"
}

type ApplicantData struct {
    ApplicantId *string `xml:"Applicant_ID,omitempty"
    PersonalData *PersonalData `xml:"Personal_Data,omitempty"
    QualificationData *QualificationData `xml:"Qualification_Data,omitempty"
    RecruitingData *RecruitingData `xml:"Recruiting_Data,omitempty"
    ResumeData *ResumeData `xml:"Resume_Data,omitempty"
    BackgroundCheckData *BackgroundCheckData `xml:"Background_Check_Data,omitempty"
    ExternalIntegrationIdData *ExternalIntegrationIdData `xml:"External_Integration_ID_Data,omitempty"
    DocumentFieldResultData *DocumentFieldResultData `xml:"Document_Field_Result_Data,omitempty"
}

type PutApplicantRequest struct {
    XMLName xml.Name `xml:"wd:Put_Applicant_Request"`
    XMLNamespace string `xml:"xmlns,attr"`
    Com.workday/bsvc}addonly *string `xml:"wd:com.workday/bsvc}Add_Only,attr"
    Com.workday/bsvc}version *string `xml:"wd:com.workday/bsvc}version,attr"
    ApplicantReference *ApplicantReference `xml:"Applicant_Reference,omitempty"
    Com.workday/bsvc}descriptor *string `xml:"wd:com.workday/bsvc}Descriptor,attr"`
    ApplicantData *ApplicantData `xml:"Applicant_Data,omitempty"
}

