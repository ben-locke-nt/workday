// Generated from ../model/put_candidate_response/put_candidate_response.go
package put_candidate_response

import "encoding/xml"

type PutCandidateResponse struct {
    XMLName xml.Name `xml:"Put_Candidate_Response"`
    Version *string `xml:"version,attr,omitempty"`
    CandidateReference *CandidateReference `xml:"Candidate_Reference,omitempty"`
    CandidateJobApplicationData *CandidateJobApplicationData `xml:"Candidate_Job_Application_Data,omitempty"`
}

type CandidateReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Id struct {
    Value *string `xml:",chardata"`
    Type *string `xml:"type,attr,omitempty"`
}

type CandidateJobApplicationData struct {
    JobApplicationReference *JobApplicationReference `xml:"Job_Application_Reference,omitempty"`
    JobRequisitionReference *JobRequisitionReference `xml:"Job_Requisition_Reference,omitempty"`
}

type JobApplicationReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type JobRequisitionReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

