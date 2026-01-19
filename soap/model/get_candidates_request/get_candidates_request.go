// Generated from ../model/get_candidates_request/get_candidates_request.go
package get_candidates_request

import "encoding/xml"

type GetCandidatesRequest struct {
    XMLName xml.Name `xml:"wd:Get_Candidates_Request"`
    XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
    Version *string `xml:"wd:version,attr,omitempty"`
    RequestReferences *RequestReferences `xml:"wd:Request_References,omitempty"`
    RequestCriteria *RequestCriteria `xml:"wd:Request_Criteria,omitempty"`
    ResponseFilter *ResponseFilter `xml:"wd:Response_Filter,omitempty"`
    ResponseGroup *ResponseGroup `xml:"wd:Response_Group,omitempty"`
}

type RequestReferences struct {
    SkipNonExistingInstances *string `xml:"wd:Skip_Non_Existing_Instances,attr,omitempty"`
    CandidateReference *CandidateReference `xml:"wd:Candidate_Reference,omitempty"`
}

type CandidateReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type Id struct {
    Value *string `xml:",chardata"`
    Type *string `xml:"wd:type,attr,omitempty"`
}

type RequestCriteria struct {
    TransactionLogCriteriaData *TransactionLogCriteriaData `xml:"wd:Transaction_Log_Criteria_Data,omitempty"`
    FirstName *FirstName `xml:"wd:First_Name,omitempty"`
    LastName *LastName `xml:"wd:Last_Name,omitempty"`
    CandidateEmailAddress *CandidateEmailAddress `xml:"wd:Candidate_Email_Address,omitempty"`
    PreHireReference *PreHireReference `xml:"wd:Pre_Hire_Reference,omitempty"`
    WorkerReference *WorkerReference `xml:"wd:Worker_Reference,omitempty"`
    JobRequisitionReference *JobRequisitionReference `xml:"wd:Job_Requisition_Reference,omitempty"`
    RecruitingStageReference *RecruitingStageReference `xml:"wd:Recruiting_Stage_Reference,omitempty"`
    ApplicantSourceReference *ApplicantSourceReference `xml:"wd:Applicant_Source_Reference,omitempty"`
    CandidateTagReference *CandidateTagReference `xml:"wd:Candidate_Tag_Reference,omitempty"`
    AppliedFrom *AppliedFrom `xml:"wd:Applied_From,omitempty"`
    AppliedThrough *AppliedThrough `xml:"wd:Applied_Through,omitempty"`
    InternalWorkersOnly *InternalWorkersOnly `xml:"wd:Internal_Workers_Only,omitempty"`
    CreatedFrom *CreatedFrom `xml:"wd:Created_From,omitempty"`
    CreatedThrough *CreatedThrough `xml:"wd:Created_Through,omitempty"`
    CriteriaMatchScoreCategoryReference *CriteriaMatchScoreCategoryReference `xml:"wd:Criteria_Match_Score_Category_Reference,omitempty"`
    StaticCandidatePoolReference *StaticCandidatePoolReference `xml:"wd:Static_Candidate_Pool_Reference,omitempty"`
}

type TransactionLogCriteriaData struct {
    TransactionDateRangeData *TransactionDateRangeData `xml:"wd:Transaction_Date_Range_Data,omitempty"`
    TransactionTypeReferences *TransactionTypeReferences `xml:"wd:Transaction_Type_References,omitempty"`
    SubscriberReference *SubscriberReference `xml:"wd:Subscriber_Reference,omitempty"`
}

type TransactionDateRangeData struct {
    UpdatedFrom *UpdatedFrom `xml:"wd:Updated_From,omitempty"`
    UpdatedThrough *UpdatedThrough `xml:"wd:Updated_Through,omitempty"`
    EffectiveFrom *EffectiveFrom `xml:"wd:Effective_From,omitempty"`
    EffectiveThrough *EffectiveThrough `xml:"wd:Effective_Through,omitempty"`
}

type UpdatedFrom struct {
    Value *string `xml:",chardata"`
}

type UpdatedThrough struct {
    Value *string `xml:",chardata"`
}

type EffectiveFrom struct {
    Value *string `xml:",chardata"`
}

type EffectiveThrough struct {
    Value *string `xml:",chardata"`
}

type TransactionTypeReferences struct {
    TransactionTypeReference *TransactionTypeReference `xml:"wd:Transaction_Type_Reference,omitempty"`
}

type TransactionTypeReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type SubscriberReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type FirstName struct {
    Value *string `xml:",chardata"`
}

type LastName struct {
    Value *string `xml:",chardata"`
}

type CandidateEmailAddress struct {
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

type JobRequisitionReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type RecruitingStageReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type ApplicantSourceReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type CandidateTagReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type AppliedFrom struct {
    Value *string `xml:",chardata"`
}

type AppliedThrough struct {
    Value *string `xml:",chardata"`
}

type InternalWorkersOnly struct {
    Value *string `xml:",chardata"`
}

type CreatedFrom struct {
    Value *string `xml:",chardata"`
}

type CreatedThrough struct {
    Value *string `xml:",chardata"`
}

type CriteriaMatchScoreCategoryReference struct {
    Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
    Id *Id `xml:"wd:ID,omitempty"`
}

type StaticCandidatePoolReference struct {
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
    ExcludeAllAttachments *ExcludeAllAttachments `xml:"wd:Exclude_All_Attachments,omitempty"`
}

type IncludeReference struct {
    Value *string `xml:",chardata"`
}

type ExcludeAllAttachments struct {
    Value *string `xml:",chardata"`
}

