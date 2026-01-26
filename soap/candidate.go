package soap

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"workday/soap/model"
	"workday/soap/model/get_candidates_request"
	"workday/soap/model/get_candidates_response"
	"workday/soap/model/get_job_application_additional_data_request"
	"workday/soap/model/put_candidate_request"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

func (s *Client) GetCandidate(permissionReport *model.PermissionCheck, workdayID string) (*get_candidates_response.CandidateData, error) {
	request := get_candidates_request.GetCandidatesRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		RequestReferences: &get_candidates_request.RequestReferences{
			CandidateReference: &get_candidates_request.CandidateReference{
				Id: &get_candidates_request.Id{
					Type:  lo.ToPtr("Candidate_ID"),
					Value: lo.ToPtr(workdayID),
				},
			},
		},
		ResponseFilter: &get_candidates_request.ResponseFilter{
			AsOfEffectiveDate: lo.ToPtr(time.Now().Format("2006-01-02")),
			Count:             lo.ToPtr(1),
		},
		// TODO - why does this break validation?
		// RequestCriteria: &get_candidates_request.RequestCriteria{
		// 	InternalWorkersOnly: lo.ToPtr(false),
		// },
		ResponseGroup: &get_candidates_request.ResponseGroup{
			IncludeReference:      lo.ToPtr(false),
			ExcludeAllAttachments: lo.ToPtr(true),
		},
	}

	var response get_candidates_response.GetCandidatesResponse
	if call := s.call("Recruiting/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return nil, fmt.Errorf("call SOAP service: %w", call.Error)
	}

	return response.ResponseData.Candidate.CandidateData, nil
}

func (s *Client) UpdateCandidateName(permissionReport *model.PermissionCheck, workdayID, firstName, middleName, lastName string) error {
	candidate, err := s.GetCandidate(permissionReport, workdayID)
	if err != nil {
		fmt.Printf("Could not find candidate with Workday ID: %s\n", workdayID)
		return err
	}

	request := put_candidate_request.PutCandidateRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		CandidateReference: &put_candidate_request.CandidateReference{
			Id: &put_candidate_request.Id{
				Type:  lo.ToPtr("WID"),
				Value: lo.ToPtr(workdayID),
			},
		},
		CandidateData: &put_candidate_request.CandidateData{
			NameData: &put_candidate_request.NameData{
				LegalName: &put_candidate_request.LegalName{
					NameDetailData: &put_candidate_request.NameDetailData{
						FirstName: lo.ToPtr(firstName),
						// They don;t accept middle names :shrug:....if configured for whatever reason. how are we going to handle this generally???
						LastName: lo.ToPtr(lastName),
					},
				},
			},
			ContactData: &put_candidate_request.ContactData{
				// The fields allowed/required for Address data is highly locaion-specific.
				// How are we going to handle this generally???
				LocationData: &put_candidate_request.LocationData{
					CountryReference: &put_candidate_request.CountryReference{
						Id: &put_candidate_request.Id{
							Type:  candidate.ContactData.LocationData.CountryReference.Id.Type,
							Value: candidate.ContactData.LocationData.CountryReference.Id.Value,
						},
					},
					AddressLine1: lo.ToPtr("hello street"),
				},
			},
			// JobApplicationData: &JobApplicationData{
			// 	JobAppliedToData: &JobAppliedToData{
			// 		GlobalPersonalInformationData: &GlobalPersonalInformationData{
			// 			DateOfBirth: &DateOfBirth{
			// 				Value: lo.ToPtr("1990-02-12"),
			// 			},
			// 		},
			// 	},
			// },
		},
	}

	var response AnyXML
	if call := s.call("Recruiting/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return fmt.Errorf("call SOAP service: %w", call.Error)
	}

	return nil
}


func (s *Client) GetJobApplicationAdditionalData(permissionReport *model.PermissionCheck, workdayID, candidateID string) error {
	request := get_job_application_additional_data_request.GetJobApplicationAdditionalDataRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		RequestReferences: &get_job_application_additional_data_request.RequestReferences{
			JobApplicationReference: &get_job_application_additional_data_request.JobApplicationReference{
				Id: &get_job_application_additional_data_request.Id{
					Type:  lo.ToPtr("WID"),
					Value: lo.ToPtr(workdayID),
				},
			},
		},
		// RequestCriteria: &get_job_application_additional_data_request.RequestCriteria{
		// 	CandidateReference: &get_job_application_additional_data_request.CandidateReference{
		// 		Id: &get_job_application_additional_data_request.Id{
		// 			Type:  lo.ToPtr("Candidate_ID"),
		// 			Value: lo.ToPtr(candidateID),
		// 		},
		// 	},
		// },
	}

	var response AnyXML
	if call := s.call("Recruiting/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return fmt.Errorf("call SOAP service: %w", call.Error)
	}

	return nil
}

// Workday custom-object payload namespace, per tenant.
// Matches what you see in `soap/test.xml`.
const WorkdayNametagCustomNamespace = "urn:com.workday/tenants/nametag_dpt1/data/custom"

// NametagIdvCustomObject builds XML like:
// <custom:nametagIdvJobApplication xmlns:custom="..." wd:Descriptor="..."> ... </custom:nametagIdvJobApplication>
type NametagIdvCustomObject struct {
	XMLName xml.Name `xml:"custom:nametagIdvJobApplication" json:"-"`

	// Ensure nested `custom:` elements/attrs serialize with the same prefix.
	XMLNSCustom string `xml:"xmlns:custom,attr,omitempty" json:"-"`

	// Workday uses `wd:` attrs on custom objects.
	Descriptor *string `xml:"wd:Descriptor,attr,omitempty" json:"-"`

	NametagIdvJobApplicationReference *NametagIdvJobApplicationReference `xml:"custom:nametagIdvJobApplicationReference,omitempty" json:"-"`

	VerifiedName      *string `xml:"custom:verifiedName,omitempty" json:"verifiedName,omitempty"`
	Status            *string `xml:"custom:status,omitempty" json:"status,omitempty"`
	RequestedAt       *string `xml:"custom:requestedAt,omitempty" json:"requestedAt,omitempty"`
	SourceDocument    *string `xml:"custom:sourceDocument,omitempty" json:"sourceDocument,omitempty"`
	LastUpdatedAt     *string `xml:"custom:lastUpdatedAt,omitempty" json:"lastUpdatedAt,omitempty"`
	VerifiedAt        *string `xml:"custom:verifiedAt,omitempty" json:"verifiedAt,omitempty"`
	VerifiedAddress   *string `xml:"custom:verifiedAddress,omitempty" json:"verifiedAddress,omitempty"`
	VerifiedBirthDate *string `xml:"custom:verifiedBirthDate,omitempty" json:"verifiedBirthDate,omitempty"`
	NametagIdvId      *string `xml:"custom:nametagIdvId,omitempty" json:"nametagIdvId,omitempty"`
}

type NametagIdvJobApplicationReference struct {
	IDs []*CustomID `xml:"custom:ID,omitempty"`
}

type CustomID struct {
	Type  *string `xml:"custom:type,attr,omitempty"`
	Value *string `xml:",chardata"`
}

// Minimal, request-local structs so we can correctly model the schema's `xsd:any`:
// `Business_Object_Additional_Data` contains arbitrary XML *directly*, not wrapped
// in a `<wd:AnyElement>` placeholder.
type putJobApplicationAdditionalDataRequest struct {
	XMLName xml.Name `xml:"wd:Put_Job_Application_Additional_Data_Request"`

	XMLNamespace *string `xml:"xmlns:wd,attr,omitempty"`
	Version      *string `xml:"wd:version,attr,omitempty"`

	JobApplicationCustomObjectData *putJobApplicationCustomObjectData `xml:"wd:Job_Application_Custom_Object_Data,omitempty"`
}

type putJobApplicationCustomObjectData struct {
	JobApplicationReference      *jobApplicationReference        `xml:"wd:Job_Application_Reference,omitempty"`
	BusinessObjectAdditionalData *businessObjectAdditionalDataAny `xml:"wd:Business_Object_Additional_Data,omitempty"`
}

type jobApplicationReference struct {
	Descriptor *string `xml:"wd:Descriptor,attr,omitempty"`
	IDs        []*wdID `xml:"wd:ID,omitempty"`
}

type wdID struct {
	Type  *string `xml:"wd:type,attr,omitempty"`
	Value *string `xml:",chardata"`
}

type businessObjectAdditionalDataAny struct {
	Content string `xml:",innerxml"`
}

func marshalInnerXML(elems ...any) (string, error) {
	var b strings.Builder
	for _, elem := range elems {
		raw, err := xml.Marshal(elem)
		if err != nil {
			return "", err
		}
		b.Write(raw)
	}
	return b.String(), nil
}

// PutJobApplicationAdditionalData sets `Business_Object_Additional_Data` to one-or-more custom objects,
// exactly like the `custom:nametagIdvJobApplication` entries in `soap/test.xml`.
func (s *Client) PutJobApplicationAdditionalData(permissionReport *model.PermissionCheck, jobApplicationWID string, customObjects ...NametagIdvCustomObject) error {
	// Default the namespace so callers don't have to remember it.
	for i := range customObjects {
		if customObjects[i].XMLNSCustom == "" {
			customObjects[i].XMLNSCustom = WorkdayNametagCustomNamespace
		}
	}

	innerElems := make([]any, 0, len(customObjects))
	for i := range customObjects {
		innerElems = append(innerElems, customObjects[i])
	}

	inner, err := marshalInnerXML(innerElems...)
	if err != nil {
		return fmt.Errorf("marshal custom objects inner XML: %w", err)
	}

	request := putJobApplicationAdditionalDataRequest{
		XMLNamespace: lo.ToPtr(WorkdayXMLNamespace),
		Version:      lo.ToPtr(WorkdayAPIVersion),
		JobApplicationCustomObjectData: &putJobApplicationCustomObjectData{
			JobApplicationReference: &jobApplicationReference{
				IDs: []*wdID{
					{
						Type:  lo.ToPtr("WID"),
						Value: lo.ToPtr(jobApplicationWID),
					},
				},
			},
			BusinessObjectAdditionalData: &businessObjectAdditionalDataAny{
				Content: inner,
			},
		},
	}

	var response AnyXML
	if call := s.call("Recruiting/"+WorkdayAPIVersion, &request, &response); call.Error != nil {
		fmt.Printf("SOAP request error: %v\n", call.Error)
		return fmt.Errorf("call SOAP service: %w", call.Error)
	}

	return nil
}

func (s *Client) RESTPutJobApplicationAdditionalData(jobApplicationWID string, customObjectID string) error {

	path := fmt.Sprintf("/customObjects/%s/%s;nametagIdvId=%s", "nametagIdvJobApplication", jobApplicationWID, customObjectID)
	
	url := "https://impl-services1.wd12.myworkday.com/ccx/api/v1/nametag_dpt1" + path

	body := `{
		"verifiedName": "testymctestface"
	}`

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader([]byte(body)))
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	fmt.Println(url)

	resp, err := s.credentialedClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	fmt.Println(resp.StatusCode)

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}

	fmt.Printf("body: %s\n", string(bodyBytes))

	return nil
}


func (s *Client) RESTPropspects(id string) error {

	// Nope :(


	path := fmt.Sprintf("/prospects/%s", id)
	
	url := "https://impl-services1.wd12.myworkday.com/ccx/api/v1/nametag_dpt1" + path

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	fmt.Println(url)

	resp, err := s.credentialedClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	fmt.Println(resp.StatusCode)

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}

	fmt.Printf("body: %s\n", string(bodyBytes))

	return nil
}

type WID struct {
	ID string `json:"id,omitempty"`
}

type NametagIdvCustomObjectJobApplication struct {
	ID *WID `json:"jobApplication,omitempty"`
	NametagIdvCustomObject
}

func (s *Client) RESTPostJobApplicationAdditionalData(jobApplicationWID string) error {

	path := fmt.Sprintf("/customObjects/%s/%s", "nametagIdvJobApplication", jobApplicationWID)
	url := "https://impl-services1.wd12.myworkday.com/ccx/api/v1/nametag_dpt1" + path

	idv := NametagIdvCustomObjectJobApplication{
		ID: &WID{
			ID: jobApplicationWID,
		},
		NametagIdvCustomObject: NametagIdvCustomObject{
			VerifiedName: lo.ToPtr("testymctestface"),
			Status: lo.ToPtr("test"),
			RequestedAt: lo.ToPtr("2021-01-01"),
			SourceDocument: lo.ToPtr("test"),
			LastUpdatedAt: lo.ToPtr("2021-01-01"),
			VerifiedAt: lo.ToPtr("2021-01-01"),
			VerifiedAddress: lo.ToPtr("test"),
			VerifiedBirthDate: lo.ToPtr("1999-01-01"),
			NametagIdvId: lo.ToPtr(uuid.New().String()[:8]),
		},
	}

	body, err := json.Marshal(idv)
	if err != nil {
		return fmt.Errorf("marshal idv: %w", err)
	}

	fmt.Printf("body: %s\n", string(body))

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	fmt.Println(url)

	resp, err := s.credentialedClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	fmt.Println(resp.StatusCode)

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}

	fmt.Printf("body: %s\n", string(bodyBytes))

	return nil
}