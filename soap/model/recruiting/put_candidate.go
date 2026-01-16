// Copyright 2025 Nametag Inc.
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

type CandidateReference struct {
	Reference *model.WorkdayObjectReferenceList `xml:"wd:Candidate_Reference"`
}

type CandidateNameDetail struct {
	FirstName string `xml:"wd:First_Name,omitempty"`
}

type CandidatePreferredName struct {
	NameDetail *CandidateNameDetail `xml:"wd:Name_Detail_Data,omitempty"`
}

type CandidateName struct {
	PreferredName *CandidatePreferredName `xml:"wd:Preferred_Name,omitempty"`
}

type CandidateData struct {
	Name *CandidateName `xml:"wd:Name_Data,omitempty"`
}

type PutCandidateRequest struct {
	XMLName      xml.Name            `xml:"wd:Put_Candidate_Request"`
	XMLNamespace string              `xml:"xmlns:wd,attr"`
	Version      string              `xml:"wd:version,attr,omitempty"`
	Reference    *CandidateReference `xml:"wd:Candidate_Reference,omitempty"`
	Data         *CandidateData      `xml:"wd:Candidate_Data,omitempty"`
	AddOnly      bool                `xml:"wd:Add_Only,omitempty"`
}

func NewPutCandidateRequest() *PutCandidateRequest {
	return &PutCandidateRequest{
		XMLNamespace: model.WorkdayXMLNamespace,
		Version:      RecruitingServiceVersion,
	}
}

func NewPutCandidateRequestType(workdayID, firstName string) *PutCandidateRequest {
	request := NewPutCandidateRequest()
	request.Reference = &CandidateReference{
		Reference: &model.WorkdayObjectReferenceList{
			References: []*model.WorkdayObjectReference{
				model.NewWorkdayIDObjectReference(workdayID),
			},
		},
	}
	request.Data = &CandidateData{
		Name: &CandidateName{
			PreferredName: &CandidatePreferredName{
				NameDetail: &CandidateNameDetail{
					FirstName: firstName,
				},
			},
		},
	}
	request.AddOnly = false
	return request
}
