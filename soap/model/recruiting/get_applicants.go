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

type GetApplicantsRequest struct {
	XMLName xml.Name `xml:"wd:Get_Applicants_Request,omitempty"`
	*model.XMLTopLevel
	References *model.WorkdayRequestReferencesList
	Pagination *model.PaginationFilter
}

func NewGetApplicantsRequest() *GetApplicantsRequest {
	return &GetApplicantsRequest{
		XMLTopLevel: model.NewXMLTopLevel(RecruitingServiceVersion),
	}
}

func NewGetApplicantByWorkdayIDRequest(workdayID string) *GetApplicantsRequest {
	request := NewGetApplicantsRequest()
	request.References = &model.WorkdayRequestReferencesList{
		References: []any{
			&ApplicantReference{
				Reference: model.NewWorkdayObjectID(workdayID),
			},
		},
	}
	request.Pagination = &model.PaginationFilter{
		Count: 1,
	}
	return request
}

type GetApplicantsResponse struct {
	Content []byte `xml:",innerxml"`
}
