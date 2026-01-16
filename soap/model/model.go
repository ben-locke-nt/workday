// Copyright 2025 Nametag Inc.
//
// All information contained herein is the property of Nametag Inc.. The
// intellectual and technical concepts contained herein are proprietary, trade
// secrets, and/or confidential to Nametag, Inc. and may be covered by U.S.
// and Foreign Patents, patents in process, and are protected by trade secret or
// copyright law. Reproduction or distribution, in whole or in part, is
// forbidden except by express written permission of Nametag, Inc.

package model

import (
	"bytes"
	"encoding/xml"
)

const WorkdayXMLNamespace = "urn:com.workday/bsvc"

type XMLTopLevel struct {
	XMLNamespace string `xml:"xmlns:wd,attr"`
	Version      string `xml:"wd:version,attr"`
}

func NewXMLTopLevel(version string) *XMLTopLevel {
	return &XMLTopLevel{
		XMLNamespace: WorkdayXMLNamespace,
		Version:      version,
	}
}

func SerializeWithXMLHeader(v any) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(`<?xml version="1.0" encoding="utf-8"?>`)
	enc := xml.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	if err := enc.Flush(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type WorkdayObjectID struct {
	XMLName string `xml:"wd:ID"`
	Type    string `xml:"wd:type,attr,omitempty"`
	Value   string `xml:",chardata"`
}

func NewWorkdayObjectID(workdayID string) *WorkdayObjectID {
	return &WorkdayObjectID{
		Type:  "WID",
		Value: workdayID,
	}
}

type WorkdayRequestReferencesList struct {
	XMLName    string `xml:"wd:Request_References"`
	References []any
}

type PaginationFilter struct {
	XMLName string `xml:"wd:Response_Filter"`
	Count   int    `xml:"wd:Count,omitempty"`
	Page    int    `xml:"wd:Page,omitempty"`
}
