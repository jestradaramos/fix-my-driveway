package domain

// Repo will hold the functions that any DB trying to be used by this app needs to have these functions implemented
type Repo interface {
	AddUser(*User) error
	UpdateUser(*User) error
	GetUserByEmail(string) (*User, error) 
	// AddArrivalTime()
	// UpdateArrivalTime()
	// DeleteArrivalTime()
	// AddDepartureTime()
	// UpdateDepartureTime()
	// DeleteDepartureTime()
}
