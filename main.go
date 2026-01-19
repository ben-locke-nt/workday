// Copyright 2025 Nametag Inc.
//
// All information contained herein is the property of Nametag Inc.. The
// intellectual and technical concepts contained herein are proprietary, trade
// secrets, and/or confidential to Nametag, Inc. and may be covered by U.S.
// and Foreign Patents, patents in process, and are protected by trade secret or
// copyright law. Reproduction or distribution, in whole or in part, is
// forbidden except by express written permission of Nametag, Inc.

package main

import (
	"log"
	"workday/soap"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	soapClient, err := soap.New()
	if err != nil {
		log.Fatalf("Error creating SOAP client: %v", err)
	}

	// To get perms for the design doc:
	// View Integration System Security Group (Unconstrained)

	// To change scopes for api client:
	// view api clients -> edit

	// Get_Workers security:
	// https://doc.workday.com/admin-guide/en-us/integrations/workday-web-services-and-integration-ids/workday-soap-api-guidelines-and-troubleshooting/get-workers-wws.html#:~:text=Title%3A%20Get%20Workers%20Product%20Area(s)%3A%20Integration%20Get%20Workers%20The%20Get_Workers%20web%20service%20is%20used%20to%20retrieve%20worker%20details%20from%20Workday%2C%20and%20supports%20granular%20transaction%20detection%20through%20the%20use%20of%20the%20%3Cwd%3ATransaction_Log_Criteria_Data%3E%20section.
	//soapClient.GetWorker("b3c518ac84321001c1c8ed9c98790000")

	// Get_Candidates security:
	// https://doc.workday.com/admin-guide/en-us/integrations/workday-web-services-and-integration-ids/workday-soap-api-guidelines-and-troubleshooting/get-candidate-wws.html?toc=17.8.6
	//soapClient.GetCandidate("fd9e9f710bc881556faa7574ac0fd306")
	
	// Get_Applicants security:
	// 
	//soapClient.GetApplicant("7687e33f159501842af83fedc4564407")

	//soapClient.UpdateWorkerName("b3c518ac84321001c1c8ed9c98790000", "Benjamin", "Billy", "Locke")

	//soapClient.UpdateCandidateName("fd9e9f710bc881556faa7574ac0fd306", "Sabriña", "Marie", "Valdez")

	soapClient.UpdateApplicant("7687e33f159501842af83fedc4564407", "Amanda", "Willy", "Lockehart")
}

// Applicant == Pre-hire
// Prospect == Candidate without application

// Just an applicant
// Candidate: Name=Amanda Lockhart; Workday ID=7687e33f15958135d3352f70be567c03; Candidate_ID=CAND-1318
// Applicant: Name=Amanda Lockhart; Workday ID=7687e33f159501842af83fedc4564407; Applicant_ID=A01692

// Employee applied to internal job posting
// Employee:  Name=Leigh Freidman; Workday ID=2f0532a0d3b84c0dadc842094ccc1411; Employee_ID=21229
// Candidate: Name=Leigh Freidman; Workday ID=58ec81dd9b8990668d1537aa76410b91; Candidate_ID=CAND-1036
// Applicant: Name=Leigh Freidman; Workday ID=4c952fc67fa747ce8ef83a8f9803ae82; Applicant_ID=A01404

// Prospsect: Sabrina Valdez
// Workday ID=fd9e9f710bc881556faa7574ac0fd306; Candidate_ID=CAND-1349

// Greg Weiss (Referral) For: R-00225 Director, Technical Operations
// Candidate: Workday ID=8828a84896248192f4886d9ad4556514; Candidate_ID=CAND-1245
// Prospect:  Workday ID=8828a848962481e3ab0edd62f8549c57; Candidate_ID=CAND-1321
// Job Application: Workday ID=8828a848962481631711499bd5558e14; Job_Application_ID=JOB_APPLICATION-11-385

// Employee:  Ben Locke
// Workday ID=b3c518ac84321001c1c8ed9c98790000; Employee_ID=21603
// Existing nametagIdvForWorker: Workday ID=f5d0dde74ced30010bed10cec0a90000
