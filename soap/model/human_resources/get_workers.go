// Copyright 2026 Nametag Inc.
//
// All information contained herein is the property of Nametag Inc.. The
// intellectual and technical concepts contained herein are proprietary, trade
// secrets, and/or confidential to Nametag, Inc. and may be covered by U.S.
// and Foreign Patents, patents in process, and are protected by trade secret or
// copyright law. Reproduction or distribution, in whole or in part, is
// forbidden except by express written permission of Nametag, Inc.

package human_resources

import (
	"encoding/xml"

	"github.com/nametaginc/nt/server/workday/client/.dev/soap/model"
)

type GetWorkersRequest struct {
	XMLName xml.Name `xml:"wd:Get_Workers_Request"`
	*model.XMLTopLevel
	References *model.WorkdayRequestReferencesList
	Pagination *model.PaginationFilter
}

type WorkerReference struct {
	XMLName   xml.Name `xml:"wd:Worker_Reference"`
	Reference *model.WorkdayObjectID
}

func NewGetWorkersRequest() *GetWorkersRequest {
	return &GetWorkersRequest{
		XMLTopLevel: model.NewXMLTopLevel(HumanResourcesServiceVersion),
	}
}

func NewGetWorkerByWorkdayIDRequest(workdayID string) *GetWorkersRequest {
	request := NewGetWorkersRequest()
	request.References = &model.WorkdayRequestReferencesList{
		References: []any{
			&WorkerReference{
				Reference: model.NewWorkdayObjectID(workdayID),
			},
		},
	}
	request.Pagination = &model.PaginationFilter{
		Count: 1,
	}
	return request
}

type GetWorkersResponse struct {
}
