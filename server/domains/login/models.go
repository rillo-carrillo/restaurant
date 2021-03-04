package login

//UserResponse when created return id
type UserResponse struct {
	ID       uint
	Username string
}

//StatusResponse only reply with a true or false
type StatusResponse struct {
	Status bool
}
