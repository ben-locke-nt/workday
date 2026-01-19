// Copyright 2026 Nametag Inc.
//
// All information contained herein is the property of Nametag Inc.. The
// intellectual and technical concepts contained herein are proprietary, trade
// secrets, and/or confidential to Nametag, Inc. and may be covered by U.S.
// and Foreign Patents, patents in process, and are protected by trade secret or
// copyright law. Reproduction or distribution, in whole or in part, is
// forbidden except by express written permission of Nametag, Inc.

package recruiting

import (
	"encoding/xml"

	"github.com/nametaginc/nt/server/workday/client/.dev/soap/model"
)

type PutApplicantRequest struct {
	XMLName xml.Name `xml:"wd:Put_Applicant_Request,omitempty"`
	*model.XMLTopLevel
	Reference *ApplicantReference `xml:"wd:Applicant_Reference,omitempty"`
	Data      *ApplicantData      `xml:"wd:Applicant_Data,omitempty"`
}

type ApplicantReference struct {
	Reference []*model.WorkdayObjectID
}

func NewPutApplicantRequest() *PutApplicantRequest {
	return &PutApplicantRequest{
		XMLTopLevel: model.NewXMLTopLevel(RecruitingServiceVersion),
	}
}

func NewPutApplicantByWorkdayIDRequest(workdayID string) *PutApplicantRequest {
	request := NewPutApplicantRequest()
	request.Reference = &ApplicantReference{
		Reference: []*model.WorkdayObjectID{model.NewWorkdayObjectID(workdayID)},
	}
	return request
}

type PutApplicantResponse struct {
	Content []byte `xml:",innerxml"`
}

type ApplicantData struct {
	ApplicantID               *string                      `xml:"wd:Applicant_ID,omitempty"`
	PersonalData              *PersonalData                `xml:"wd:Personal_Data,omitempty"`
	QualificationData         *QualificationData           `xml:"wd:Qualification_Data,omitempty"`
	RecruitingData            *RecruitingData              `xml:"wd:Recruiting_Data,omitempty"`
	ResumeData                *ResumeData                  `xml:"wd:Resume_Data,omitempty"`
	BackgroundCheckData       []*BackgroundCheckData       `xml:"wd:Background_Check_Data,omitempty"`
	ExternalIntegrationIDData []*ExternalIntegrationIDData `xml:"wd:External_Integration_ID_Data,omitempty"`
	DocumentFieldResultData   []*DocumentFieldResultData   `xml:"wd:Document_Field_Result_Data,omitempty"`
}

type PersonalData struct {
	UniversalID             *string                    `xml:"wd:Universal_ID,omitempty"`
	NameData                *NameData                  `xml:"wd:Name_Data,omitempty"`
	PersonalInformationData []*PersonalInformationData `xml:"wd:Personal_Information_Data,omitempty"`
	IdentificationData      *IdentificationData        `xml:"wd:Identification_Data,omitempty"`
	ContactData             *ContactData               `xml:"wd:Contact_Data,omitempty"`
	TobaccoUse              *bool                      `xml:"wd:Tobacco_Use,omitempty"`
}

type NameData struct {
	LegalNameData      *LegalNameData        `xml:"wd:Legal_Name_Data,omitempty"`
	PreferredNameData  *PreferredNameData    `xml:"wd:Preferred_Name_Data,omitempty"`
	AdditionalNameData []*AdditionalNameData `xml:"wd:Additional_Name_Data,omitempty"`
}

type LegalNameData struct {
	NameDetailData *NameDetailData `xml:"wd:Name_Detail_Data,omitempty"`
}

type PreferredNameData struct {
	NameDetailData *NameDetailData `xml:"wd:Name_Detail_Data,omitempty"`
}

type AdditionalNameData struct {
	NameDetailData    *NameDetailData    `xml:"wd:Name_Detail_Data,omitempty"`
	NameTypeReference *NameTypeReference `xml:"wd:Name_Type_Reference,omitempty"`
}

type NameDetailData struct {
	FormattedName *string `xml:"wd:Formatted_Name,omitempty"`
	ReportingName *string `xml:"wd:Reporting_Name,omitempty"`
}

type NameTypeReference struct {
	Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
}

type PersonalInformationData struct {
	PersonalInformationForCountryData []*PersonalInformationForCountryData `xml:"wd:Personal_Information_For_Country_Data,omitempty"`
	BirthDate                         *string                              `xml:"wd:Birth_Date,omitempty"`
	DateOfDeath                       *string                              `xml:"wd:Date_of_Death,omitempty"`
	CountryOfBirthReference           *CountryReference                    `xml:"wd:Country_of_Birth_Reference,omitempty"`
	CountryRegionOfBirthReference     *CountryRegionReference              `xml:"wd:Country_Region_of_Birth_Reference,omitempty"`
	RegionOfBirthDescriptor           *string                              `xml:"wd:Region_of_Birth_Descriptor,omitempty"`
	CityOfBirth                       *string                              `xml:"wd:City_of_Birth,omitempty"`
	CityOfBirthReference              *CityReference                       `xml:"wd:City_of_Birth_Reference,omitempty"`
	CitizenshipStatusReference        []*CitizenshipStatusReference        `xml:"wd:Citizenship_Status_Reference,omitempty"`
	PrimaryNationalityReference       *NationalityReference                `xml:"wd:Primary_Nationality_Reference,omitempty"`
	AdditionalNationalitiesReference  []*NationalityReference              `xml:"wd:Additional_Nationalities_Reference,omitempty"`
	LastMedicalExamValidTo            *string                              `xml:"wd:Last_Medical_Exam_Valid_To,omitempty"`
	LastMedicalExamDate               *string                              `xml:"wd:Last_Medical_Exam_Date,omitempty"`
	MedicalExamNotes                  *string                              `xml:"wd:Medical_Exam_Notes,omitempty"`
	BloodTypeReference                *BloodTypeReference                  `xml:"wd:Blood_Type_Reference,omitempty"`
	MilitaryServiceData               []*MilitaryServiceData               `xml:"wd:Military_Service_Data,omitempty"`
	SexualOrientationReference        []*SexualOrientationReference        `xml:"wd:Sexual_Orientation_Reference,omitempty"`
	GenderIdentityReference           []*GenderIdentityReference           `xml:"wd:Gender_Identity_Reference,omitempty"`
	PronounReference                  []*PronounReference                  `xml:"wd:Pronoun_Reference,omitempty"`
	NonCountrySpecificSectionData     *NonCountrySpecificSectionData       `xml:"wd:Non_Country_Specific_Section_Data,omitempty"`
}

// --- Begin full XML struct descent ---
type PersonalInformationForCountryData struct {
	CountryReference               *CountryReference                 `xml:"wd:Country_Reference,omitempty"`
	CountryPersonalInformationData []*CountryPersonalInformationData `xml:"wd:Country_Personal_Information_Data,omitempty"`
}

type CountryReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type CountryPersonalInformationData struct {
	// Add all relevant fields as needed
}

type CountryRegionReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type CityReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type CitizenshipStatusReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type NationalityReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type BloodTypeReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type MilitaryServiceData struct {
	StatusReference                *StatusReference                `xml:"wd:Status_Reference,omitempty"`
	DischargeDate                  *string                         `xml:"wd:Discharge_Date,omitempty"`
	StatusBeginDate                *string                         `xml:"wd:Status_Begin_Date,omitempty"`
	MilitaryServiceTypeReference   *MilitaryServiceTypeReference   `xml:"wd:Military_Service_Type_Reference,omitempty"`
	MilitaryRankReference          *MilitaryRankReference          `xml:"wd:Military_Rank_Reference,omitempty"`
	Notes                          *string                         `xml:"wd:Notes,omitempty"`
	MilitaryServiceReference       *MilitaryServiceReference       `xml:"wd:Military_Service_Reference,omitempty"`
	MilitaryDischargeTypeReference *MilitaryDischargeTypeReference `xml:"wd:Military_Discharge_Type_Reference,omitempty"`
}

type StatusReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type MilitaryServiceTypeReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type MilitaryRankReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type MilitaryServiceReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type MilitaryDischargeTypeReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type SexualOrientationReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type GenderIdentityReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type PronounReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type NonCountrySpecificSectionData struct {
	NonCountrySpecificSection1Data *NonCountrySpecificSection1Data `xml:"wd:Non_Country_Specific_Section_1_Data,omitempty"`
	NonCountrySpecificSection2Data *NonCountrySpecificSection2Data `xml:"wd:Non_Country_Specific_Section_2_Data,omitempty"`
	NonCountrySpecificSection3Data *NonCountrySpecificSection3Data `xml:"wd:Non_Country_Specific_Section_3_Data,omitempty"`
}

type NonCountrySpecificSection1Data struct {
	// Add all relevant fields as needed
}

type NonCountrySpecificSection2Data struct {
	// Add all relevant fields as needed
}

type NonCountrySpecificSection3Data struct {
	// Add all relevant fields as needed
}

type NationalID struct {
	Delete              *bool                `xml:"wd:Delete,attr,omitempty"`
	NationalIDReference *NationalIDReference `xml:"wd:National_ID_Reference,omitempty"`
	NationalIDData      *NationalIDData      `xml:"wd:National_ID_Data,omitempty"`
}

type NationalIDReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type NationalIDData struct {
	// Add all relevant fields as needed
}

type GovernmentID struct {
	Delete *bool `xml:"wd:Delete,attr,omitempty"`
	// Add all relevant fields as needed
}

type VisaID struct {
	Delete *bool `xml:"wd:Delete,attr,omitempty"`
	// Add all relevant fields as needed
}

type PassportID struct {
	Delete *bool `xml:"wd:Delete,attr,omitempty"`
	// Add all relevant fields as needed
}

type LicenseID struct {
	Delete *bool `xml:"wd:Delete,attr,omitempty"`
	// Add all relevant fields as needed
}

type CustomID struct {
	Delete *bool `xml:"wd:Delete,attr,omitempty"`
	// Add all relevant fields as needed
}

type AddressData struct {
	FormattedAddress             *string `xml:"wd:Formatted_Address,omitempty"`
	AddressFormatType            *string `xml:"wd:Address_Format_Type,omitempty"`
	DefaultedBusinessSiteAddress *bool   `xml:"wd:Defaulted_Business_Site_Address,omitempty"`
	Delete                       *bool   `xml:"wd:Delete,attr,omitempty"`
	DoNotReplaceAll              *bool   `xml:"wd:Do_Not_Replace_All,attr,omitempty"`
	EffectiveDate                *string `xml:"wd:Effective_Date,omitempty"`
	// Add all relevant fields as needed
}

type PhoneData struct {
	AreaCode                         *string `xml:"wd:Area_Code,omitempty"`
	TenantFormattedPhone             *string `xml:"wd:Tenant_Formatted_Phone,omitempty"`
	InternationalFormattedPhone      *string `xml:"wd:International_Formatted_Phone,omitempty"`
	PhoneNumberWithoutAreaCode       *string `xml:"wd:Phone_Number_Without_Area_Code,omitempty"`
	NationalFormattedPhone           *string `xml:"wd:National_Formatted_Phone,omitempty"`
	E164FormattedPhone               *string `xml:"wd:E164_Formatted_Phone,omitempty"`
	WorkdayTraditionalFormattedPhone *string `xml:"wd:Workday_Traditional_Formatted_Phone,omitempty"`
	Delete                           *bool   `xml:"wd:Delete,attr,omitempty"`
	DoNotReplaceAll                  *bool   `xml:"wd:Do_Not_Replace_All,attr,omitempty"`
	// Add all relevant fields as needed
}

type EmailAddressData struct {
	Delete          *bool `xml:"wd:Delete,attr,omitempty"`
	DoNotReplaceAll *bool `xml:"wd:Do_Not_Replace_All,attr,omitempty"`
	// Add all relevant fields as needed
}

type InstantMessengerData struct {
	Delete          *bool `xml:"wd:Delete,attr,omitempty"`
	DoNotReplaceAll *bool `xml:"wd:Do_Not_Replace_All,attr,omitempty"`
	// Add all relevant fields as needed
}

type WebAddressData struct {
	Delete          *bool `xml:"wd:Delete,attr,omitempty"`
	DoNotReplaceAll *bool `xml:"wd:Do_Not_Replace_All,attr,omitempty"`
	// Add all relevant fields as needed
}

type ExternalJobHistory struct {
	JobHistoryReference *JobHistoryReference `xml:"wd:Job_History_Reference,omitempty"`
	JobHistoryData      []*JobHistoryData    `xml:"wd:Job_History_Data,omitempty"`
}

type JobHistoryReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type JobHistoryData struct {
	// Add all relevant fields as needed
}

type Competency struct {
	CompetencyLevelBehaviorReference  *CompetencyLevelBehaviorReference `xml:"wd:Competency_Level_Behavior_Reference,omitempty"`
	CompetencyLevelBehaviorDescriptor *string                           `xml:"wd:Competency_Level_Behavior_Descriptor,omitempty"`
	AssessmentComment                 *string                           `xml:"wd:Assessment_Comment,omitempty"`
	AssessedOn                        *string                           `xml:"wd:Assessed_On,omitempty"`
	AssessedByPersonReference         *AssessedByPersonReference        `xml:"wd:Assessed_By_Person_Reference,omitempty"`
	AssessedByWorkerDescriptor        *string                           `xml:"wd:Assessed_By_Worker_Descriptor,omitempty"`
	CompetencyReference               *CompetencyReference              `xml:"wd:Competency_Reference,omitempty"`
	CompetencyDescriptor              *string                           `xml:"wd:Competency_Descriptor,omitempty"`
}

type CompetencyLevelBehaviorReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type AssessedByPersonReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type CompetencyReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type Certification struct {
	CertificationReference *CertificationReference `xml:"wd:Certification_Reference,omitempty"`
	CertificationData      []*CertificationData    `xml:"wd:Certification_Data,omitempty"`
}

type CertificationReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type CertificationData struct {
	// Add all relevant fields as needed
}

type Training struct {
	TrainingReference *TrainingReference `xml:"wd:Training_Reference,omitempty"`
	TrainingData      []*TrainingData    `xml:"wd:Training_Data,omitempty"`
}

type TrainingReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type TrainingData struct {
	// Add all relevant fields as needed
}

type AwardAndActivity struct {
	AwardAndActivityReference *AwardAndActivityReference `xml:"wd:Award_and_Activity_Reference,omitempty"`
	AwardAndActivityData      *AwardAndActivityData      `xml:"wd:Award_and_Activity_Data,omitempty"`
}

type AwardAndActivityReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type AwardAndActivityData struct {
	// Add all relevant fields as needed
}

type OrganizationMembership struct {
	OrganizationProfessionalAffiliationReference *OrganizationProfessionalAffiliationReference `xml:"wd:Organization_Professional_Affiliation_Reference,omitempty"`
	OrganizationProfessionalAffiliationData      []*OrganizationProfessionalAffiliationData    `xml:"wd:Organization_Professional_Affiliation_Data,omitempty"`
}

type OrganizationProfessionalAffiliationReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type OrganizationProfessionalAffiliationData struct {
	// Add all relevant fields as needed
}

type Education struct {
	EducationReference *EducationReference `xml:"wd:Education_Reference,omitempty"`
	EducationData      []*EducationData    `xml:"wd:Education_Data,omitempty"`
}

type EducationReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type EducationData struct {
	// Add all relevant fields as needed
}

type WorkExperience struct {
	ExperienceReference       *ExperienceReference       `xml:"wd:Experience_Reference,omitempty"`
	RemoveExperience          *bool                      `xml:"wd:Remove_Experience,omitempty"`
	ExperienceRatingReference *ExperienceRatingReference `xml:"wd:Experience_Rating_Reference,omitempty"`
	ExperienceComment         *string                    `xml:"wd:Experience_Comment,omitempty"`
}

type ExperienceReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type ExperienceRatingReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type Language struct {
	LanguageReference         *LanguageReference         `xml:"wd:Language_Reference,omitempty"`
	RemoveLanguage            *bool                      `xml:"wd:Remove_Language,omitempty"`
	NativeLanguage            *bool                      `xml:"wd:Native_Language,omitempty"`
	LanguageAbility           []*LanguageAbility         `xml:"wd:Language_Ability,omitempty"`
	AssessedOn                *string                    `xml:"wd:Assessed_On,omitempty"`
	Note                      *string                    `xml:"wd:Note,omitempty"`
	AssessedByWorkerReference *AssessedByWorkerReference `xml:"wd:Assessed_by_Worker_Reference,omitempty"`
}

type LanguageReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type LanguageAbility struct {
	// Add all relevant fields as needed
}

type AssessedByWorkerReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type InternalProjectExperience struct {
	InternalProjectExperienceReference *InternalProjectExperienceReference `xml:"wd:Internal_Project_Experience_Reference,omitempty"`
	InternalProjectExperienceData      []*InternalProjectExperienceData    `xml:"wd:Internal_Project_Experience_Data,omitempty"`
}

type InternalProjectExperienceReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type InternalProjectExperienceData struct {
	// Add all relevant fields as needed
}

type EligibleForHireReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type ApplicantHasMarkedAsNoShowReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type ApplicantSourceReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type ReferredByWorkerReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type PositionsConsideredForReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type QualificationData struct {
	ExternalJobHistory        []*ExternalJobHistory        `xml:"wd:External_Job_History,omitempty"`
	Competency                []*Competency                `xml:"wd:Competency,omitempty"`
	Certification             []*Certification             `xml:"wd:Certification,omitempty"`
	Training                  []*Training                  `xml:"wd:Training,omitempty"`
	AwardAndActivity          []*AwardAndActivity          `xml:"wd:Award_and_Activity,omitempty"`
	OrganizationMembership    []*OrganizationMembership    `xml:"wd:Organization_Membership,omitempty"`
	Education                 []*Education                 `xml:"wd:Education,omitempty"`
	WorkExperience            []*WorkExperience            `xml:"wd:Work_Experience,omitempty"`
	Language                  []*Language                  `xml:"wd:Language,omitempty"`
	InternalProjectExperience []*InternalProjectExperience `xml:"wd:Internal_Project_Experience,omitempty"`
}

type RecruitingData struct {
	ApplicantEnteredDate                *string                              `xml:"wd:Applicant_Entered_Date,omitempty"`
	ApplicantComments                   *string                              `xml:"wd:Applicant_Comments,omitempty"`
	EligibleForHireReference            *EligibleForHireReference            `xml:"wd:Eligible_For_Hire_Reference,omitempty"`
	EligibleForRehireComments           *string                              `xml:"wd:Eligible_for_Rehire_Comments,omitempty"`
	ApplicantHasMarkedAsNoShowReference *ApplicantHasMarkedAsNoShowReference `xml:"wd:Applicant_Has_Marked_as_No_Show_Reference,omitempty"`
	ApplicantSourceReference            *ApplicantSourceReference            `xml:"wd:Applicant_Source_Reference,omitempty"`
	ReferredByWorkerReference           []*ReferredByWorkerReference         `xml:"wd:Referred_by_Worker_Reference,omitempty"`
	PositionsConsideredForReference     []*PositionsConsideredForReference   `xml:"wd:Positions_Considered_for_Reference,omitempty"`
}

type ResumeData struct {
	Resume []*Resume `xml:"wd:Resume,omitempty"`
}

type Resume struct {
	ResumeReference *ResumeReference  `xml:"wd:Resume_Reference,omitempty"`
	ResumeData      *ResumeDataDetail `xml:"wd:Resume_Data,omitempty"`
}

type ResumeReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type ResumeDataDetail struct {
	// Add all relevant fields as needed
}

type BackgroundCheckData struct {
	StatusDate      *string          `xml:"wd:Status_Date,omitempty"`
	StatusReference *StatusReference `xml:"wd:Status_Reference,omitempty"`
	StatusComment   *string          `xml:"wd:Status_Comment,omitempty"`
}

type ExternalIntegrationIDData struct {
	ID []*ID `xml:"wd:ID,omitempty"`
}

type ID struct {
	SystemID *string `xml:"wd:System_ID,attr,omitempty"`
	Value    *string `xml:",chardata"`
}

type DocumentFieldResultData struct {
	FieldReference *FieldReference `xml:"wd:Field_Reference,omitempty"`
	Value          *string         `xml:"wd:Value,omitempty"`
}

type FieldReference struct {
	Descriptor *string                  `xml:"wd:Descriptor,attr,omitempty"`
	References []*model.WorkdayObjectID `xml:"wd:ID,omitempty"`
}

type IdentificationData struct {
	NationalID   []*NationalID   `xml:"wd:National_ID,omitempty"`
	GovernmentID []*GovernmentID `xml:"wd:Government_ID,omitempty"`
	VisaID       []*VisaID       `xml:"wd:Visa_ID,omitempty"`
	PassportID   []*PassportID   `xml:"wd:Passport_ID,omitempty"`
	LicenseID    []*LicenseID    `xml:"wd:License_ID,omitempty"`
	CustomID     []*CustomID     `xml:"wd:Custom_ID,omitempty"`
}

type ContactData struct {
	AddressData          []*AddressData          `xml:"wd:Address_Data,omitempty"`
	PhoneData            []*PhoneData            `xml:"wd:Phone_Data,omitempty"`
	EmailAddressData     []*EmailAddressData     `xml:"wd:Email_Address_Data,omitempty"`
	InstantMessengerData []*InstantMessengerData `xml:"wd:Instant_Messenger_Data,omitempty"`
	WebAddressData       []*WebAddressData       `xml:"wd:Web_Address_Data,omitempty"`
}

// --- End full XML struct descent ---
