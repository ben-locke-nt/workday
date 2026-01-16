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
)

// ResponseResults represents pagination and result summary information
type ResponseResults struct {
	TotalResults *float64 `xml:"wd:Total_Results,omitempty"`
	TotalPages   *float64 `xml:"wd:Total_Pages,omitempty"`
	PageResults  *float64 `xml:"wd:Page_Results,omitempty"`
	Page         *float64 `xml:"wd:Page,omitempty"`
}

// UniversalIdentifierReference represents a universal identifier reference
type UniversalIdentifierReference struct {
	Descriptor string `xml:"wd:Descriptor,attr,omitempty"`
	ID         []ID   `xml:"wd:ID"`
}

// WorkerData represents comprehensive worker data
type WorkerData struct {
	WorkerID    *string `xml:"wd:Worker_ID,omitempty"`
	UserID      *string `xml:"wd:User_ID,omitempty"`
	UniversalID *string `xml:"wd:Universal_ID,omitempty"`
	// Note: PersonalData and other nested structures would be added here
	// This is a simplified version - the full structure would include all nested elements
}

// Worker represents a worker in the response
type Worker struct {
	WorkerReference              *WorkerReference              `xml:"wd:Worker_Reference,omitempty"`
	WorkerDescriptor             *string                       `xml:"wd:Worker_Descriptor,omitempty"`
	UniversalIdentifierReference *UniversalIdentifierReference `xml:"wd:Universal_Identifier_Reference,omitempty"`
	WorkerData                   *WorkerData                   `xml:"wd:Worker_Data,omitempty"`
}

// ResponseData represents the response data containing workers
type ResponseData struct {
	Worker []Worker `xml:"wd:Worker"`
}

// GetWorkersResponse represents the root element of the Get_Workers_Response XML
type GetWorkersResponse struct {
	XMLName           xml.Name           `xml:"wd:Get_Workers_Response"`
	XMLNamespace      string             `xml:"xmlns:wd,attr"`
	Version           string             `xml:"wd:version,attr"`
	RequestReferences *RequestReferences `xml:"wd:Request_References,omitempty"`
	RequestCriteria   *RequestCriteria   `xml:"wd:Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilter    `xml:"wd:Response_Filter,omitempty"`
	ResponseGroup     *ResponseGroup     `xml:"wd:Response_Group,omitempty"`
	ResponseResults   *ResponseResults   `xml:"wd:Response_Results,omitempty"`
	ResponseData      *ResponseData      `xml:"wd:Response_Data,omitempty"`
}
