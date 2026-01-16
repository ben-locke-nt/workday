// Copyright 2025 Nametag Inc.
//
// All information contained herein is the property of Nametag Inc.. The
// intellectual and technical concepts contained herein are proprietary, trade
// secrets, and/or confidential to Nametag, Inc. and may be covered by U.S.
// and Foreign Patents, patents in process, and are protected by trade secret or
// copyright law. Reproduction or distribution, in whole or in part, is
// forbidden except by express written permission of Nametag, Inc.

package client

import (
	"time"
)

// NametagIDVCustomObject represents the Nametag IDV custom business object in Workday.
// Though multiple custom business objects must be created in Workday for the different
// business object types, they will all have the exact same fields, so on our side we
// can — dare I say, should! — represent them with a single struct.
// TODO - needs to be JSON for REST and XML for SOAP. Currently only JSON is implemented.
type NametagIDVCustomObject struct {
	RequestedAt                        *time.Time `json:"requestedAt,omitempty"`                        // display order a
	Status                             *string    `json:"status,omitempty"`                             // display order b
	AssertedName                       *string    `json:"assertedName,omitempty"`                       // display order c
	VerifiedName                       *string    `json:"verifiedName,omitempty"`                       // display order d
	AssertedBirthdate                  *time.Time `json:"assertedBirthDate,omitempty"`                  // display order e
	VerifiedBirthdate                  *time.Time `json:"verifiedBirthDate,omitempty"`                  // display order f
	AssertedAndVerifiedNamesMatch      *bool      `json:"assertedAndVerifiedNamesMatch,omitempty"`      // display order g
	AssertedAndVerifiedBirthDatesMatch *bool      `json:"assertedAndVerifiedBirthDatesMatch,omitempty"` // display order h
	VerifiedAt                         *time.Time `json:"verifiedAt,omitempty"`                         // display order i
	NametagID                          *string    `json:"nametagId,omitempty"`                          // display order j
	WorkdayID                          *string    `json:"workdayId,omitempty"`                          // display order k
}

// WID is required to serialize the Workday ID field correctly in REST JSON payloads
type WID struct {
	ID string `json:"id,omitempty"`
}

// NametagWorkerIDVCustomObject represents the Nametag IDV custom business object for a Worker in Workday
type NametagWorkerIDVCustomObject struct {
	ID *WID `json:"worker,omitempty"`
	NametagIDVCustomObject
}

// NametagProspectIDVCustomObject represents the Nametag IDV custom business object for a Prospect in Workday
type NametagProspectIDVCustomObject struct {
	ID *WID `json:"prospect,omitempty"`
	NametagIDVCustomObject
}
