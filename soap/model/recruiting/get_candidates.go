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

type GetCandidatesRequest struct {
	XMLName xml.Name `xml:"wd:Get_Candidates_Request,omitempty"`
	*model.XMLTopLevel
	References *model.WorkdayRequestReferencesList
	Pagination *model.PaginationFilter
}

type CandidateReference struct {
	XMLName   xml.Name `xml:"wd:Candidate_Reference,omitempty"`
	Reference *model.WorkdayObjectID
}

func NewGetCandidatesRequest() *GetCandidatesRequest {
	return &GetCandidatesRequest{
		XMLTopLevel: model.NewXMLTopLevel(RecruitingServiceVersion),
	}
}

func NewGetCandidateByWorkdayIDRequest(workdayID string) *GetCandidatesRequest {
	request := NewGetCandidatesRequest()
	request.References = &model.WorkdayRequestReferencesList{
		References: []any{
			&CandidateReference{
				Reference: model.NewWorkdayObjectID(workdayID),
			},
		},
	}
	request.Pagination = &model.PaginationFilter{
		Count: 1,
	}
	return request
}

type GetCandidatesResponse struct {
	Content []byte `xml:",innerxml"`
}
