package domain

// Repo will hold the functions that any DB trying to be used by this app needs to have these functions implemented
type Repo interface {
	AddUser(User) (*User, error)
	UpdateUser(User) (*User, error)
	// AddArrivalTime()
	// UpdateArrivalTime()
	// DeleteArrivalTime()
	// AddDepartureTime()
	// UpdateDepartureTime()
	// DeleteDepartureTime()
}
