package utils

type HackeronePrograms struct {
	Data  []HackeroneProgram     `json:"data"`
	Links HackeroneProgramsLinks `json:"links"`
}

type HackeroneProgramsLinks struct {
	Self string `json:"self"`
	Next string `json:"next"`
}

type HackeroneProgram struct {
	ID            string                        `json:"id"`
	Type          string                        `json:"type"`
	Attributes    HackeroneProgramAttributes    `json:"attributes"`
	Relationships HackeroneProgramRelationShips `json:"relationships"`
}

type HackeroneProgramRelationShips struct {
	StructuredScopes HackeroneProgramStructuredScopes `json:"structured_scopes"`
}

type HackeroneProgramStructuredScopes struct {
	Data []HackeroneProgramStructuredScopesData `json:"data"`
}

type HackeroneProgramStructuredScopesData struct {
	ID         string                                         `json:"id"`
	Type       string                                         `json:"type"`
	Attributes HackeroneProgramStructuredScopesDataAttributes `json:"attributes"`
}

type HackeroneProgramStructuredScopesDataAttributes struct {
	AssetType                  string `json:"asset_type"`
	AssetIdentifier            string `json:"asset_identifier"`
	EligibleForBounty          bool   `json:"eligible_for_bounty"`
	EligibleForSubmission      bool   `json:"eligible_for_submission"`
	Instruction                string `json:"instruction"`
	MaxSeverity                string `json:"max_severity"`
	CreatedAt                  string `json:"created_at"`
	UpdatedAt                  string `json:"updated_at"`
	ConfidentialityRequirement string `json:"confidentiality_requirement"`
	IntegrityRequirement       string `json:"integrity_requirement"`
	AvailabilityRequirement    string `json:"availability_requirement"`
}

type HackeroneProgramAttributes struct {
	Handle                          string `json:"handle"`
	Name                            string `json:"name"`
	Currency                        string `json:"currency"`
	ProfilePicture                  string `json:"profile_picture"`
	SubmissionState                 string `json:"submission_state"`
	TriageActive                    bool   `json:"triage_active"`
	State                           string `json:"state"`
	StartedAcceptingAt              string `json:"started_accepting_at"`
	NumberOfReportsForUser          int    `json:"number_of_reports_for_user"`
	NumberOfValidReportsForUser     int    `json:"number_of_valid_reports_for_user"`
	LastInvitationAcceptedAtForUser string `json:"last_invitation_accepted_at_for_user"`
	Bookmarked                      bool   `json:"bookmarked"`
	AllowsBountySplitting           bool   `json:"allows_bounty_splitting"`
	OffersBounties                  bool   `json:"offers_bounties"`
}
