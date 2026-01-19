// Generated from ../model/change_legal_name_request/change_legal_name_request.go
package change_legal_name_request

import "encoding/xml"

type ChangeLegalNameRequest struct {
    XMLName xml.Name `xml:"wd:Change_Legal_Name_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    BusinessProcessParameters *BusinessProcessParameters `xml:"wd:Business_Process_Parameters,omitempty"`
    ChangeLegalNameData *ChangeLegalNameData `xml:"wd:Change_Legal_Name_Data,omitempty"`
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

type ChangeLegalNameData struct {
    PersonReference *PersonReference `xml:"wd:Person_Reference,omitempty"`
    UniversalIdReference *UniversalIdReference `xml:"wd:Universal_ID_Reference,omitempty"`
    EffectiveDate *EffectiveDate `xml:"wd:Effective_Date,omitempty"`
    NameData *NameData `xml:"wd:Name_Data,omitempty"`
}

type PersonReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type UniversalIdReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type EffectiveDate struct {
    Value *string `xml:",chardata"`
}

type NameData struct {
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

