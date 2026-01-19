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
	ApplicantID               ApplicantID             `json:"Applicant_ID"`
	PersonalData              PersonalData            `json:"Personal_Data"`
	QualificationData         QualificationData       `json:"Qualification_Data"`
	RecruitingData            RecruitingData          `json:"Recruiting_Data"`
	ResumeData                ApplicantDataResumeData `json:"Resume_Data"`
	BackgroundCheckData       BackgroundCheckData     `json:"Background_Check_Data"`
	ExternalIntegrationIDData ApplicantReference      `json:"External_Integration_ID_Data"`
	DocumentFieldResultData   DocumentFieldResultData `json:"Document_Field_Result_Data"`
}

type BackgroundCheckData struct {
	StatusDate      string             `json:"Status_Date"`
	StatusReference ApplicantReference `json:"Status_Reference"`
	StatusComment   ApplicantID        `json:"Status_Comment"`
}

type DocumentFieldResultData struct {
	FieldReference ApplicantReference `json:"Field_Reference"`
	Value          ApplicantID        `json:"Value"`
}

type PersonalData struct {
	UniversalID             ApplicantID        `json:"Universal_ID"`
	NameData                NameData           `json:"Name_Data"`
	PersonalInformationData string             `json:"Personal_Information_Data"`
	IdentificationData      IdentificationData `json:"Identification_Data"`
	ContactData             ContactData        `json:"Contact_Data"`
	TobaccoUse              bool               `json:"Tobacco_Use"`
}

type ContactData struct {
	AddressData          AddressData          `json:"Address_Data"`
	PhoneData            PhoneData            `json:"Phone_Data"`
	EmailAddressData     EmailAddressData     `json:"Email_Address_Data"`
	InstantMessengerData InstantMessengerData `json:"Instant_Messenger_Data"`
	WebAddressData       WebAddressData       `json:"Web_Address_Data"`
}

type AddressData struct {
	CountryReference        ApplicantReference `json:"Country_Reference"`
	LastModified            string             `json:"Last_Modified"`
	AddressLineData         ApplicantID        `json:"Address_Line_Data"`
	Municipality            ApplicantID        `json:"Municipality"`
	CountryCityReference    ApplicantReference `json:"Country_City_Reference"`
	SubmunicipalityData     ApplicantID        `json:"Submunicipality_Data"`
	CountryRegionReference  ApplicantReference `json:"Country_Region_Reference"`
	CountryRegionDescriptor ApplicantID        `json:"Country_Region_Descriptor"`
	SubregionData           ApplicantID        `json:"Subregion_Data"`
	PostalCode              ApplicantID        `json:"Postal_Code"`
	UsageData               UsageData          `json:"Usage_Data"`
	NumberOfDays            int64              `json:"Number_of_Days"`
	MunicipalityLocal       ApplicantID        `json:"Municipality_Local"`
	AddressReference        ApplicantReference `json:"Address_Reference"`
	AddressID               ApplicantID        `json:"Address_ID"`
}

type UsageData struct {
	TypeData                TypeData           `json:"Type_Data"`
	UseForReference         ApplicantReference `json:"Use_For_Reference"`
	UseForTenantedReference ApplicantReference `json:"Use_For_Tenanted_Reference"`
	Comments                ApplicantID        `json:"Comments"`
}

type TypeData struct {
	TypeReference ApplicantReference `json:"Type_Reference"`
}

type EmailAddressData struct {
	EmailAddress   ApplicantID        `json:"Email_Address"`
	EmailComment   ApplicantID        `json:"Email_Comment"`
	UsageData      UsageData          `json:"Usage_Data"`
	EmailReference ApplicantReference `json:"Email_Reference"`
	ID             ApplicantID        `json:"ID"`
}

type InstantMessengerData struct {
	InstantMessengerAddress       ApplicantID        `json:"Instant_Messenger_Address"`
	InstantMessengerTypeReference ApplicantReference `json:"Instant_Messenger_Type_Reference"`
	InstantMessengerComment       ApplicantID        `json:"Instant_Messenger_Comment"`
	UsageData                     UsageData          `json:"Usage_Data"`
	InstantMessengerReference     ApplicantReference `json:"Instant_Messenger_Reference"`
	ID                            ApplicantID        `json:"ID"`
}

type PhoneData struct {
	CountryISOCode           string             `json:"Country_ISO_Code"`
	InternationalPhoneCode   ApplicantID        `json:"International_Phone_Code"`
	PhoneNumber              ApplicantID        `json:"Phone_Number"`
	PhoneExtension           ApplicantID        `json:"Phone_Extension"`
	PhoneDeviceTypeReference ApplicantReference `json:"Phone_Device_Type_Reference"`
	UsageData                UsageData          `json:"Usage_Data"`
	PhoneReference           ApplicantReference `json:"Phone_Reference"`
	ID                       ApplicantID        `json:"ID"`
}

type WebAddressData struct {
	WebAddress          ApplicantID        `json:"Web_Address"`
	WebAddressComment   ApplicantID        `json:"Web_Address_Comment"`
	UsageData           UsageData          `json:"Usage_Data"`
	WebAddressReference ApplicantReference `json:"Web_Address_Reference"`
	ID                  ApplicantID        `json:"ID"`
}

type IdentificationData struct {
	NationalID   NationalID   `json:"National_ID"`
	GovernmentID GovernmentID `json:"Government_ID"`
	VisaID       VisaID       `json:"Visa_ID"`
	PassportID   PassportID   `json:"Passport_ID"`
	LicenseID    LicenseID    `json:"License_ID"`
	CustomID     CustomID     `json:"Custom_ID"`
}

type CustomID struct {
	CustomIDReference       ApplicantReference `json:"Custom_ID_Reference"`
	CustomIDData            CustomIDData       `json:"Custom_ID_Data"`
	CustomIDSharedReference ApplicantReference `json:"Custom_ID_Shared_Reference"`
}

type CustomIDData struct {
	ID                      ApplicantID        `json:"ID"`
	IDTypeReference         ApplicantReference `json:"ID_Type_Reference"`
	IssuedDate              string             `json:"Issued_Date"`
	ExpirationDate          string             `json:"Expiration_Date"`
	OrganizationIDReference ApplicantReference `json:"Organization_ID_Reference"`
	CustomDescription       ApplicantID        `json:"Custom_Description"`
}

type GovernmentID struct {
	GovernmentIDReference       ApplicantReference `json:"Government_ID_Reference"`
	GovernmentIDData            TIDData            `json:"Government_ID_Data"`
	GovernmentIDSharedReference ApplicantReference `json:"Government_ID_Shared_Reference"`
}

type TIDData struct {
	ID                         ApplicantID         `json:"ID"`
	IDTypeReference            ApplicantReference  `json:"ID_Type_Reference"`
	CountryReference           ApplicantReference  `json:"Country_Reference"`
	IssuedDate                 string              `json:"Issued_Date"`
	ExpirationDate             string              `json:"Expiration_Date"`
	VerificationDate           string              `json:"Verification_Date"`
	TIDDataVerifiedByReference *ApplicantReference `json:"Verified_by_Reference,omitempty"`
	VerifiedByReference        *ApplicantReference `json:"Verified_By_Reference,omitempty"`
}

type LicenseID struct {
	LicenseIDReference       ApplicantReference `json:"License_ID_Reference"`
	LicenseIDData            LicenseIDData      `json:"License_ID_Data"`
	LicenseIDSharedReference ApplicantReference `json:"License_ID_Shared_Reference"`
}

type LicenseIDData struct {
	ID                      ApplicantID        `json:"ID"`
	IDTypeReference         ApplicantReference `json:"ID_Type_Reference"`
	CountryReference        ApplicantReference `json:"Country_Reference"`
	CountryRegionReference  ApplicantReference `json:"Country_Region_Reference"`
	CountryRegionDescriptor ApplicantID        `json:"Country_Region_Descriptor"`
	AuthorityReference      ApplicantReference `json:"Authority_Reference"`
	LicenseClass            ApplicantID        `json:"License_Class"`
	IssuedDate              string             `json:"Issued_Date"`
	ExpirationDate          string             `json:"Expiration_Date"`
	VerificationDate        string             `json:"Verification_Date"`
	VerifiedByReference     ApplicantReference `json:"Verified_By_Reference"`
}

type NationalID struct {
	NationalIDReference       ApplicantReference `json:"National_ID_Reference"`
	NationalIDData            IDData             `json:"National_ID_Data"`
	NationalIDSharedReference ApplicantReference `json:"National_ID_Shared_Reference"`
}

type IDData struct {
	ID                  ApplicantID        `json:"ID"`
	IDTypeReference     ApplicantReference `json:"ID_Type_Reference"`
	CountryReference    ApplicantReference `json:"Country_Reference"`
	IssuedDate          string             `json:"Issued_Date"`
	ExpirationDate      string             `json:"Expiration_Date"`
	VerificationDate    string             `json:"Verification_Date"`
	Series              *ApplicantID       `json:"Series,omitempty"`
	IssuingAgency       *ApplicantID       `json:"Issuing_Agency,omitempty"`
	VerifiedByReference ApplicantReference `json:"Verified_By_Reference"`
}

type PassportID struct {
	PassportIDReference       ApplicantReference `json:"Passport_ID_Reference"`
	PassportIDData            TIDData            `json:"Passport_ID_Data"`
	PassportIDSharedReference ApplicantReference `json:"Passport_ID_Shared_Reference"`
}

type VisaID struct {
	VisaIDReference       ApplicantReference `json:"Visa_ID_Reference"`
	VisaIDData            IDData             `json:"Visa_ID_Data"`
	VisaIDSharedReference ApplicantReference `json:"Visa_ID_Shared_Reference"`
}

type NameData struct {
	LegalNameData      LegalNameData      `json:"Legal_Name_Data"`
	PreferredNameData  PreferredNameData  `json:"Preferred_Name_Data"`
	AdditionalNameData AdditionalNameData `json:"Additional_Name_Data"`
}

type AdditionalNameData struct {
	NameDetailData    NameDetailData     `json:"Name_Detail_Data"`
	NameTypeReference ApplicantReference `json:"Name_Type_Reference"`
}

type NameDetailData struct {
	CountryReference                ApplicantReference  `json:"Country_Reference"`
	PrefixData                      PrefixData          `json:"Prefix_Data"`
	FirstName                       ApplicantID         `json:"First_Name"`
	MiddleName                      ApplicantID         `json:"Middle_Name"`
	LastName                        ApplicantID         `json:"Last_Name"`
	SecondaryLastName               ApplicantID         `json:"Secondary_Last_Name"`
	TertiaryLastName                ApplicantID         `json:"Tertiary_Last_Name"`
	LocalNameDetailData             LocalNameDetailData `json:"Local_Name_Detail_Data"`
	SuffixData                      SuffixData          `json:"Suffix_Data"`
	FullNameForSingaporeAndMalaysia ApplicantID         `json:"Full_Name_for_Singapore_and_Malaysia"`
}

type LocalNameDetailData struct {
	FirstName          ApplicantID `json:"First_Name"`
	MiddleName         ApplicantID `json:"Middle_Name"`
	LastName           ApplicantID `json:"Last_Name"`
	SecondaryLastName  ApplicantID `json:"Secondary_Last_Name"`
	FirstName2         ApplicantID `json:"First_Name_2"`
	MiddleName2        ApplicantID `json:"Middle_Name_2"`
	LastName2          ApplicantID `json:"Last_Name_2"`
	SecondaryLastName2 ApplicantID `json:"Secondary_Last_Name_2"`
}

type PrefixData struct {
	TitleReference      ApplicantReference `json:"Title_Reference"`
	TitleDescriptor     ApplicantID        `json:"Title_Descriptor"`
	SalutationReference ApplicantReference `json:"Salutation_Reference"`
}

type SuffixData struct {
	SocialSuffixReference       ApplicantReference `json:"Social_Suffix_Reference"`
	SocialSuffixDescriptor      ApplicantID        `json:"Social_Suffix_Descriptor"`
	AcademicSuffixReference     ApplicantReference `json:"Academic_Suffix_Reference"`
	HereditarySuffixReference   ApplicantReference `json:"Hereditary_Suffix_Reference"`
	HonorarySuffixReference     ApplicantReference `json:"Honorary_Suffix_Reference"`
	ProfessionalSuffixReference ApplicantReference `json:"Professional_Suffix_Reference"`
	ReligiousSuffixReference    ApplicantReference `json:"Religious_Suffix_Reference"`
	RoyalSuffixReference        ApplicantReference `json:"Royal_Suffix_Reference"`
}

type LegalNameData struct {
	NameDetailData string `json:"Name_Detail_Data"`
}

type PreferredNameData struct {
	NameDetailData NameDetailData `json:"Name_Detail_Data"`
}

type QualificationData struct {
	ExternalJobHistory        ExternalJobHistory        `json:"External_Job_History"`
	Competency                Competency                `json:"Competency"`
	Certification             Certification             `json:"Certification"`
	Training                  Training                  `json:"Training"`
	AwardAndActivity          AwardAndActivity          `json:"Award_and_Activity"`
	OrganizationMembership    OrganizationMembership    `json:"Organization_Membership"`
	Education                 Education                 `json:"Education"`
	WorkExperience            WorkExperience            `json:"Work_Experience"`
	Language                  Language                  `json:"Language"`
	InternalProjectExperience InternalProjectExperience `json:"Internal_Project_Experience"`
}

type AwardAndActivity struct {
	AwardAndActivityReference ApplicantReference   `json:"Award_and_Activity_Reference"`
	AwardAndActivityData      AwardAndActivityData `json:"Award_and_Activity_Data"`
}

type AwardAndActivityData struct {
	AwardAndActivityID            ApplicantID        `json:"Award_and_Activity_ID"`
	RemoveAwardAndActivity        bool               `json:"Remove_Award_and_Activity"`
	AwardAndActivityTypeReference ApplicantReference `json:"Award_and_Activity_Type_Reference"`
	Title                         ApplicantID        `json:"Title"`
	SponsorIssuer                 ApplicantID        `json:"Sponsor_Issuer"`
	StartDate                     string             `json:"Start_Date"`
	EndDate                       string             `json:"End_Date"`
	Description                   ApplicantID        `json:"Description"`
	URL                           ApplicantID        `json:"URL"`
	RelatedPositionReference      ApplicantReference `json:"Related_Position_Reference"`
}

type Certification struct {
	CertificationReference ApplicantReference `json:"Certification_Reference"`
	CertificationData      CertificationData  `json:"Certification_Data"`
}

type CertificationData struct {
	CertificationID          ApplicantID              `json:"Certification_ID"`
	RemoveCertification      bool                     `json:"Remove_Certification"`
	CertificationReference   ApplicantReference       `json:"Certification_Reference"`
	CountryReference         ApplicantReference       `json:"Country_Reference"`
	CertificationName        ApplicantID              `json:"Certification_Name"`
	Issuer                   ApplicantID              `json:"Issuer"`
	CertificationNumber      ApplicantID              `json:"Certification_Number"`
	IssuedDate               string                   `json:"Issued_Date"`
	ExpirationDate           string                   `json:"Expiration_Date"`
	ExaminationScore         ApplicantID              `json:"Examination_Score"`
	ExaminationDate          string                   `json:"Examination_Date"`
	SpecialtyAchievementData SpecialtyAchievementData `json:"Specialty_Achievement_Data"`
	WorkerDocumentData       MentData                 `json:"Worker_Document_Data"`
}

type SpecialtyAchievementData struct {
	SpecialtyReference    ApplicantReference `json:"Specialty_Reference"`
	StartDate             string             `json:"Start_Date"`
	EndDate               string             `json:"End_Date"`
	SubspecialtyReference ApplicantReference `json:"Subspecialty_Reference"`
}

type MentData struct {
	FileName                  ApplicantID        `json:"File_Name"`
	Comment                   ApplicantID        `json:"Comment"`
	File                      string             `json:"File"`
	DocumentCategoryReference ApplicantReference `json:"Document_Category_Reference"`
	ContentType               *ApplicantID       `json:"Content_Type,omitempty"`
}

type Competency struct {
	CompetencyLevelBehaviorReference  ApplicantReference `json:"Competency_Level_Behavior_Reference"`
	CompetencyLevelBehaviorDescriptor ApplicantID        `json:"Competency_Level_Behavior_Descriptor"`
	AssessmentComment                 ApplicantID        `json:"Assessment_Comment"`
	AssessedOn                        string             `json:"Assessed_On"`
	AssessedByPersonReference         ApplicantReference `json:"Assessed_By_Person_Reference"`
	AssessedByWorkerDescriptor        ApplicantID        `json:"Assessed_By_Worker_Descriptor"`
	CompetencyReference               ApplicantReference `json:"Competency_Reference"`
	CompetencyDescriptor              ApplicantID        `json:"Competency_Descriptor"`
}

type Education struct {
	EducationReference ApplicantReference `json:"Education_Reference"`
	EducationData      EducationData      `json:"Education_Data"`
}

type EducationData struct {
	EducationID               ApplicantID        `json:"Education_ID"`
	RemoveEducation           bool               `json:"Remove_Education"`
	CountryReference          ApplicantReference `json:"Country_Reference"`
	SchoolReference           ApplicantReference `json:"School_Reference"`
	SchoolName                ApplicantID        `json:"School_Name"`
	SchoolTypeReference       ApplicantReference `json:"School_Type_Reference"`
	Location                  ApplicantID        `json:"Location"`
	DegreeReference           ApplicantReference `json:"Degree_Reference"`
	DegreeCompletedReference  ApplicantReference `json:"Degree_Completed_Reference"`
	DateDegreeReceived        string             `json:"Date_Degree_Received"`
	FieldOfStudyReference     ApplicantReference `json:"Field_Of_Study_Reference"`
	GradeAverage              ApplicantID        `json:"Grade_Average"`
	FirstYearAttended         string             `json:"First_Year_Attended"`
	LastYearAttended          string             `json:"Last_Year_Attended"`
	IsHighestLevelOfEducation bool               `json:"Is_Highest_Level_of_Education"`
	FirstDayAttended          string             `json:"First_Day_Attended"`
	LastDayAttended           string             `json:"Last_Day_Attended"`
	EducationAttachmentData   MentData           `json:"Education_Attachment_Data"`
}

type ExternalJobHistory struct {
	JobHistoryReference ApplicantReference `json:"Job_History_Reference"`
	JobHistoryData      JobHistoryData     `json:"Job_History_Data"`
}

type JobHistoryData struct {
	JobHistoryID                    ApplicantID        `json:"Job_History_ID"`
	RemoveJobHistory                bool               `json:"Remove_Job_History"`
	JobTitle                        ApplicantID        `json:"Job_Title"`
	Company                         ApplicantID        `json:"Company"`
	JobHistoryCompanyReference      ApplicantReference `json:"Job_History_Company_Reference"`
	StartDate                       string             `json:"Start_Date"`
	EndDate                         string             `json:"End_Date"`
	ResponsibilitiesAndAchievements ApplicantID        `json:"Responsibilities_and_Achievements"`
	Location                        ApplicantID        `json:"Location"`
	JobReference                    ApplicantID        `json:"Job_Reference"`
	Contact                         ApplicantID        `json:"Contact"`
}

type InternalProjectExperience struct {
	InternalProjectExperienceReference ApplicantReference            `json:"Internal_Project_Experience_Reference"`
	InternalProjectExperienceData      InternalProjectExperienceData `json:"Internal_Project_Experience_Data"`
}

type InternalProjectExperienceData struct {
	InternalProjectExperienceID     ApplicantID `json:"Internal_Project_Experience_ID"`
	RemoveInternalProjectExperience bool        `json:"Remove_Internal_Project_Experience"`
	InternalProject                 ApplicantID `json:"Internal_Project"`
	Description                     ApplicantID `json:"Description"`
	StartDate                       string      `json:"Start_Date"`
	EndDate                         string      `json:"End_Date"`
	ProjectLeader                   ApplicantID `json:"Project_Leader"`
}

type Language struct {
	LanguageReference         ApplicantReference `json:"Language_Reference"`
	RemoveLanguage            bool               `json:"Remove_Language"`
	NativeLanguage            bool               `json:"Native_Language"`
	LanguageAbility           LanguageAbility    `json:"Language_Ability"`
	AssessedOn                string             `json:"Assessed_On"`
	Note                      ApplicantID        `json:"Note"`
	AssessedByWorkerReference ApplicantReference `json:"Assessed_by_Worker_Reference"`
}

type LanguageAbility struct {
	LanguageAbilityData LanguageAbilityData `json:"Language_Ability_Data"`
}

type LanguageAbilityData struct {
	LanguageProficiencyReference ApplicantReference `json:"Language_Proficiency_Reference"`
	LanguageAbilityTypeReference ApplicantReference `json:"Language_Ability_Type_Reference"`
}

type OrganizationMembership struct {
	OrganizationProfessionalAffiliationReference ApplicantReference                      `json:"Organization_Professional_Affiliation_Reference"`
	OrganizationProfessionalAffiliationData      OrganizationProfessionalAffiliationData `json:"Organization_Professional_Affiliation_Data"`
}

type OrganizationProfessionalAffiliationData struct {
	ProfessionalAffiliationID                        ApplicantID        `json:"Professional_Affiliation_ID"`
	RemoveProfessionalAffiliation                    bool               `json:"Remove_Professional_Affiliation"`
	ProfessionalAffiliationReference                 ApplicantReference `json:"Professional_Affiliation_Reference"`
	ProfessionalAffiliation                          ApplicantID        `json:"Professional_Affiliation"`
	ProfessionalAffiliationTypeReference             ApplicantReference `json:"Professional_Affiliation_Type_Reference"`
	Affiliation                                      ApplicantID        `json:"Affiliation"`
	ProfessionalAffiliationRelationshipTypeReference ApplicantReference `json:"Professional_Affiliation_Relationship_Type_Reference"`
	BeginDate                                        string             `json:"Begin_Date"`
	EndDate                                          string             `json:"End_Date"`
	ContactInformationData                           ContactData        `json:"Contact_Information_Data"`
}

type Training struct {
	TrainingReference ApplicantReference `json:"Training_Reference"`
	TrainingData      TrainingData       `json:"Training_Data"`
}

type TrainingData struct {
	TrainingID            ApplicantID        `json:"Training_ID"`
	RemoveTraining        bool               `json:"Remove_Training"`
	Training              ApplicantID        `json:"Training"`
	Description           ApplicantID        `json:"Description"`
	TrainingTypeReference ApplicantReference `json:"Training_Type_Reference"`
	CompletionDate        string             `json:"Completion_Date"`
	TrainingDuration      ApplicantID        `json:"Training_Duration"`
}

type WorkExperience struct {
	ExperienceReference       ApplicantReference `json:"Experience_Reference"`
	RemoveExperience          bool               `json:"Remove_Experience"`
	ExperienceRatingReference ApplicantReference `json:"Experience_Rating_Reference"`
	ExperienceComment         ApplicantID        `json:"Experience_Comment"`
}

type RecruitingData struct {
	ApplicantEnteredDate                string             `json:"Applicant_Entered_Date"`
	ApplicantComments                   ApplicantID        `json:"Applicant_Comments"`
	EligibleForHireReference            ApplicantReference `json:"Eligible_For_Hire_Reference"`
	EligibleForRehireComments           ApplicantID        `json:"Eligible_for_Rehire_Comments"`
	ApplicantHasMarkedAsNoShowReference ApplicantReference `json:"Applicant_Has_Marked_as_No_Show_Reference"`
	ApplicantSourceReference            ApplicantReference `json:"Applicant_Source_Reference"`
	ReferredByWorkerReference           ApplicantReference `json:"Referred_by_Worker_Reference"`
	PositionsConsideredForReference     ApplicantReference `json:"Positions_Considered_for_Reference"`
}

type ApplicantDataResumeData struct {
	Resume Resume `json:"Resume"`
}

type Resume struct {
	ResumeReference ApplicantReference `json:"Resume_Reference"`
	ResumeData      ResumeResumeData   `json:"Resume_Data"`
}

type ResumeResumeData struct {
	FileID   ApplicantID `json:"File_ID"`
	File     string      `json:"File"`
	FileName ApplicantID `json:"FileName"`
	Comments ApplicantID `json:"Comments"`
}

type ApplicantID string

const (
	String ApplicantID = "string"
)
