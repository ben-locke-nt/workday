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

type JobApplicationReferences struct {
	References *model.WorkdayObjectReferenceList `xml:"wd:Job_Application_Reference"`
}

type GetJobApplicationRequest struct {
	XMLName      xml.Name                  `xml:"wd:Get_Job_Application_Additional_Data_Request"`
	XMLNamespace string                    `xml:"xmlns:wd,attr"`
	Version      string                    `xml:"wd:version,attr"`
	References   *JobApplicationReferences `xml:"wd:Request_References,omitempty"`
	Pagination   *model.PaginationFilter   `xml:"wd:Response_Filter,omitempty"`
}

func NewGetJobApplicationRequest() *GetJobApplicationRequest {
	return &GetJobApplicationRequest{
		XMLNamespace: model.WorkdayXMLNamespace,
		Version:      RecruitingServiceVersion,
	}
}

func NewGetJobApplicationRequestByWorkdayID(workdayID string) *GetJobApplicationRequest {
	request := NewGetJobApplicationRequest()
	request.References = &JobApplicationReferences{
		References: &model.WorkdayObjectReferenceList{
			References: []*model.WorkdayObjectReference{
				model.NewWorkdayIDObjectReference(workdayID),
			},
		},
	}
	request.Pagination = &model.PaginationFilter{
		Count: 1,
	}
	return request
}
