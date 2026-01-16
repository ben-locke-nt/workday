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

	"github.com/joho/godotenv"

	"github.com/nametaginc/nt/server/workday/client/.dev/soap"
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

	//soapClient.Human_Resources__Get_Worker("b3c518ac84321001c1c8ed9c98790000")

	//soapClient.Recruiting__Get_Candidate("8828a84896248192f4886d9ad4556514")
	//soapClient.Recruiting__Put_Candidate("8828a84896248192f4886d9ad4556514", "Gregg")
	//soapClient.Recruiting__Get_Job_Application("8828a848962481631711499bd5558e14")

	soapClient.GetWorker("b3c518ac84321001c1c8ed9c98790000")

	//soapClient.Recruiting_Get_Candidate("8828a84896248192f4886d9ad4556514")
	//soapClient.Recruiting_Get_Candidate("8828a848962481e3ab0edd62f8549c57")
	//soapClient.Human_Resources__Get_Workers()
	//soapClient.Recruiting__Put_Applicant()

	// restClient, err := rest.NewRESTClient()
	// if err != nil {
	// 	log.Fatalf("Error creating REST client: %v", err)
	// }

	//restClient.GetAWorker("b3c518ac84321001c1c8ed9c98790000")
	//restClient.GetWorkerCustomObjects("b3c518ac84321001c1c8ed9c98790000")
	//restClient.GetWorkerCustomObject("f5d0dde74ced30010bed10cec0a90000")
	//restClient.CreateWorkerCustomObject("b3c518ac84321001c1c8ed9c98790000", time.Now().AddDate(0, 0, -3))
	//restClient.CreateWorkerCustomObject("b3c518ac84321001c1c8ed9c98790000", time.Now().AddDate(0, 0, -2))
	//restClient.CreateWorkerCustomObject("b3c518ac84321001c1c8ed9c98790000", time.Now().AddDate(0, 0, -1))

	// restClient.CreateWorkerCustomObject(&workday.NametagWorkerIDVCustomObject{
	// 	WorkerID: &workday.WorkdayID{ID: "b3c518ac84321001c1c8ed9c98790000"},
	// 	NametagIDVCustomObject: workday.NametagIDVCustomObject{
	// 		RequestedAt:                        lo.ToPtr(time.Now()),
	// 		Status:                             lo.ToPtr("Pending"),
	// 		AssertedName:                       lo.ToPtr("Ben Locke"),
	// 		AssertedBirthdate:                  lo.ToPtr(time.Date(1987, time.December, 15, 0, 0, 0, 0, time.UTC)),
	// 	},
	// })

	// restClient.CreateWorkerCustomObject(&workday.NametagWorkerIDVCustomObject{
	// 	WorkerID: &workday.WorkdayID{ID: "b3c518ac84321001c1c8ed9c98790000"},
	// 	NametagIDVCustomObject: workday.NametagIDVCustomObject{
	// 		RequestedAt:                        lo.ToPtr(time.Now().AddDate(0, 0, -3)),
	// 		VerifiedAt:                         lo.ToPtr(time.Now().AddDate(0, 0, -2)),
	// 		Status:                             lo.ToPtr("Done"),
	// 		AssertedName:                       lo.ToPtr("Ben Locke"),
	// 		VerifiedName:                       lo.ToPtr("Judson Benton Locke"),
	// 		AssertedBirthdate:                  lo.ToPtr(time.Date(1987, time.December, 15, 0, 0, 0, 0, time.UTC)),
	// 		VerifiedBirthdate:                  lo.ToPtr(time.Date(1987, time.December, 15, 0, 0, 0, 0, time.UTC)),
	// 		AssertedAndVerifiedNamesMatch:      lo.ToPtr(true),
	// 		AssertedAndVerifiedBirthDatesMatch: lo.ToPtr(true),
	// 	},
	// })

	// restClient.CreateWorkerCustomObject(&workday.NametagWorkerIDVCustomObject{
	// 	WorkerID: &workday.WorkdayID{ID: "b3c518ac84321001c1c8ed9c98790000"},
	// 	NametagIDVCustomObject: workday.NametagIDVCustomObject{
	// 		RequestedAt:                        lo.ToPtr(time.Now().AddDate(0, 0, -5)),
	// 		VerifiedAt:                         lo.ToPtr(time.Now().AddDate(0, 0, -4)),
	// 		Status:                             lo.ToPtr("Done"),
	// 		AssertedName:                       lo.ToPtr("Ben Locke"),
	// 		VerifiedName:                       lo.ToPtr("Judson Benton Locke"),
	// 		AssertedBirthdate:                  lo.ToPtr(time.Date(1987, time.December, 15, 0, 0, 0, 0, time.UTC)),
	// 		VerifiedBirthdate:                  lo.ToPtr(time.Date(1987, time.December, 15, 0, 0, 0, 0, time.UTC)),
	// 		AssertedAndVerifiedNamesMatch:      lo.ToPtr(true),
	// 		AssertedAndVerifiedBirthDatesMatch: lo.ToPtr(true),
	// 	},
	// })

	// restClient.CreateWorkerCustomObject(&workday.NametagWorkerIDVCustomObject{
	// 	WorkerID: &workday.WorkdayID{ID: "b3c518ac84321001c1c8ed9c98790000"},
	// 	NametagIDVCustomObject: workday.NametagIDVCustomObject{
	// 		RequestedAt:       lo.ToPtr(time.Now()),
	// 		Status:            lo.ToPtr("Pending"),
	// 		AssertedName:      lo.ToPtr("Ben Locke"),
	// 		AssertedBirthdate: lo.ToPtr(time.Date(1987, time.December, 15, 0, 0, 0, 0, time.UTC)),
	// 	},
	// })

	//restClient.GetJobApplicationCustomObject("8828a84896248192f4886d9ad4556514")
}

// Applicant == Pre-hire

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
