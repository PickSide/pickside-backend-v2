package data

type Group struct {
	ID               uint64 `json:"id"`
	CoverPhoto       string `json:"coverPhoto"`
	Description      string `json:"description"`
	Members          string `json:"members"`
	Name             string `json:"name"`
	RequiresApproval bool   `json:"requiresApproval"`
	Visibility       string `json:"visibility"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
	OrganizerID      uint64 `json:"organizerId"`
	SportID          uint64 `json:"sportId"`
}
