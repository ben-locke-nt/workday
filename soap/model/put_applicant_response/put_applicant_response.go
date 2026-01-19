// Generated from ../model/put_applicant_response/put_applicant_response.go
package put_applicant_response

import "encoding/xml"

type PutApplicantResponse struct {
    XMLName xml.Name `xml:"Put_Applicant_Response"`
    Version *string `xml:"version,attr,omitempty"`
    ApplicantReference *ApplicantReference `xml:"Applicant_Reference,omitempty"`
    ExceptionsResponseData *ExceptionsResponseData `xml:"Exceptions_Response_Data,omitempty"`
}

type ApplicantReference struct {
    Descriptor *string `xml:"Descriptor,attr,omitempty"`
    Id *Id `xml:"ID,omitempty"`
}

type Id struct {
    Value *string `xml:",chardata"`
    Type *string `xml:"type,attr,omitempty"`
}

type ExceptionsResponseData struct {
    ExceptionsData *ExceptionsData `xml:"Exceptions_Data,omitempty"`
}

type ExceptionsData struct {
    ExceptionData *ExceptionData `xml:"Exception_Data,omitempty"`
}

type ExceptionData struct {
    Classification *Classification `xml:"Classification,omitempty"`
    Message *Message `xml:"Message,omitempty"`
}

type Classification struct {
    Value *string `xml:",chardata"`
}

type Message struct {
    Value *string `xml:",chardata"`
}

