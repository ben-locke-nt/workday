package model

type PermissionCheck struct {
	CanGetWorkerLegalName    bool `json:"can_get_worker_legal_name"`
	CanGetCandidateLegalName bool `json:"can_get_candidate_legal_name"`
	CanGetApplicantLegalName bool `json:"can_get_applicant_legal_name"`

	CanGetWorkerHomeAddress    bool `json:"can_get_worker_home_address"`
	CanGetCandidateHomeAddress bool `json:"can_get_candidate_home_address"`
	CanGetApplicantHomeAddress bool `json:"can_get_applicant_home_address"`

	CanGetWorkerBirthDate    bool `json:"can_get_worker_birth_date"`
	CanGetCandidateBirthDate bool `json:"can_get_candidate_birth_date"`
	CanGetApplicantBirthDate bool `json:"can_get_applicant_birth_date"`

	CanUpdateWorkerLegalName    bool `json:"can_update_worker_legal_name"`
	CanUpdateCandidateLegalName bool `json:"can_update_candidate_legal_name"`
	CanUpdateApplicantLegalName bool `json:"can_update_applicant_legal_name"`

	CanUpdateWorkerHomeAddress    bool `json:"can_update_worker_home_address"`
	CanUpdateCandidateHomeAddress bool `json:"can_update_candidate_home_address"`
	CanUpdateApplicantHomeAddress bool `json:"can_update_applicant_home_address"`

	CanUpdateWorkerBirthDate    bool `json:"can_update_worker_birth_date"`
	CanUpdateCandidateBirthDate bool `json:"can_update_candidate_birth_date"`
	CanUpdateApplicantBirthDate bool `json:"can_update_applicant_birth_date"`

	CanGetWorkerEmail    bool `json:"can_get_worker_email"`
	CanGetCandidateEmail bool `json:"can_get_candidate_email"`
	CanGetApplicantEmail bool `json:"can_get_applicant_email"`

	CanGetWorkerPhone    bool `json:"can_get_worker_phone"`
	CanGetCandidatePhone bool `json:"can_get_candidate_phone"`
	CanGetApplicantPhone bool `json:"can_get_applicant_phone"`
}
