package domain

// TimeEntry used for holding the information regarding peoples arrival and departure
type TimeEntry struct {
	userID    int    `json:"user_id,omitempty"`
	date      string `json:"date,omitempty"`
	entryType string `json:"entry_type,omitempty"`
}
