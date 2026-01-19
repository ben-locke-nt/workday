// Generated from ../model/get_applicants_request/get_applicants_request.go
package get_applicants_request

import "encoding/xml"

type GetApplicantsRequest struct {
    XMLName xml.Name `xml:"wd:Get_Applicants_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    RequestReferences *RequestReferences `xml:"wd:Request_References,omitempty"`
    RequestCriteria *RequestCriteria `xml:"wd:Request_Criteria,omitempty"`
    ResponseFilter *ResponseFilter `xml:"wd:Response_Filter,omitempty"`
    ResponseGroup *ResponseGroup `xml:"wd:Response_Group,omitempty"`
}

type RequestReferences struct {
    ApplicantReference *ApplicantReference `xml:"wd:Applicant_Reference,omitempty"`
}

type ApplicantReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Id struct {
    Value *string `xml:",chardata"`
    Type *string `xml:"wd:type,attr,omitempty"`
}

type RequestCriteria struct {
    WorkerReference *WorkerReference `xml:"wd:Worker_Reference,omitempty"`
    FormerWorkerReference *FormerWorkerReference `xml:"wd:Former_Worker_Reference,omitempty"`
    EmailAddress *EmailAddress `xml:"wd:Email_Address,omitempty"`
    FieldAndParameterCriteriaData *FieldAndParameterCriteriaData `xml:"wd:Field_And_Parameter_Criteria_Data,omitempty"`
}

type WorkerReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type FormerWorkerReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type EmailAddress struct {
    Value *string `xml:",chardata"`
}

type FieldAndParameterCriteriaData struct {
    ProviderReference *ProviderReference `xml:"wd:Provider_Reference,omitempty"`
}

type ProviderReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ResponseFilter struct {
    AsOfEffectiveDate *AsOfEffectiveDate `xml:"wd:As_Of_Effective_Date,omitempty"`
    AsOfEntryDatetime *AsOfEntryDatetime `xml:"wd:As_Of_Entry_DateTime,omitempty"`
    Page *Page `xml:"wd:Page,omitempty"`
    Count *Count `xml:"wd:Count,omitempty"`
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
    IncludeReference *IncludeReference `xml:"wd:Include_Reference,omitempty"`
    IncludePersonalInformation *IncludePersonalInformation `xml:"wd:Include_Personal_Information,omitempty"`
    ShowAllPersonalInformation *ShowAllPersonalInformation `xml:"wd:Show_All_Personal_Information,omitempty"`
    IncludeRecruitingInformation *IncludeRecruitingInformation `xml:"wd:Include_Recruiting_Information,omitempty"`
    IncludeQualificationProfile *IncludeQualificationProfile `xml:"wd:Include_Qualification_Profile,omitempty"`
    IncludeResume *IncludeResume `xml:"wd:Include_Resume,omitempty"`
    IncludeBackgroundCheck *IncludeBackgroundCheck `xml:"wd:Include_Background_Check,omitempty"`
    IncludeExternalIntegrationIdData *IncludeExternalIntegrationIdData `xml:"wd:Include_External_Integration_ID_Data,omitempty"`
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

