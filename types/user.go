package types

type User struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
}

func (u User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}
