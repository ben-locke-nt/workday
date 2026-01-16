// Copyright 2025 Nametag Inc.
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

/*
<bsvc:Employee_Personal_Info_Get xmlns:bsvc="urn:com.workday/bsvc" bsvc:As_Of_Date="2008-09-29" bsvc:As_Of_Moment="2014-09-18T23:18:33" bsvc:version="string">
	<bsvc:Employee_Reference>
		<bsvc:Integration_ID_Reference bsvc:Descriptor="string">
			<bsvc:ID bsvc:System_ID="string">string</bsvc:ID>
		</bsvc:Integration_ID_Reference>
	</bsvc:Employee_Reference>
</bsvc:Employee_Personal_Info_Get>
*/

type IntegrationIDReference struct {
	SystemID string `xml:"System_ID,chardata,omitempty"`
}

type EmployeeReference struct {
	IntegrationIDReference *IntegrationIDReference `xml:"Integration_ID_Reference,omitempty"`
}

type GetEmployeePersonalInformationRequest struct {
	XMLName      xml.Name           `xml:"wd:Employee_Personal_Info_Get"`
	XMLNamespace string             `xml:"xmlns:wd,attr"`
	Version      string             `xml:"wd:version,attr"`
	Reference    *EmployeeReference `xml:"wd:Employee_Reference,omitempty"`
}

func NewGetEmployeePersonalInformationRequest() *GetEmployeePersonalInformationRequest {
	return &GetEmployeePersonalInformationRequest{
		XMLNamespace: model.WorkdayXMLNamespace,
		Version:      HumanResourcesServiceVersion,
	}
}

func NewGetEmployeePersonalInformationByWorkdayIDRequest(workdayID string) *GetEmployeePersonalInformationRequest {
	request := NewGetEmployeePersonalInformationRequest()
	request.Reference = &EmployeeReference{
		IntegrationIDReference: &IntegrationIDReference{
			SystemID: "WID=" + workdayID,
		},
	}
	return request
}


