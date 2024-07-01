package domain

const (
	maxNameLength = 50
	minNameLength = 1
)

type User struct {
	uuid       UUID
	firstName  string
	lastName   string
	email      Email
	credential *Credential
}

func NewUser(
	uuid UUID,
	firstName string,
	lastName string,
	email Email,
) (*User, error) {
	if !isValidName(firstName) {
		return nil, &ValidationError{Msg: "first name is too long"}
	}
	if !isValidName(lastName) {
		return nil, &ValidationError{Msg: "last name is too long"}
	}

	return &User{
		uuid:      uuid,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
	}, nil
}

func (u *User) UUID() UUID {
	return u.uuid
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) FullName() string {
	return u.firstName + " " + u.lastName
}

func (u *User) Email() Email {
	return u.email
}

func (u *User) Credential() *Credential {
	return u.credential
}

func (u *User) SetFirstName(firstName string) error {
	if !isValidName(firstName) {
		return &ValidationError{Msg: "first name is too long"}
	}
	u.firstName = firstName
	return nil
}

func (u *User) SetLastName(lastName string) error {
	if !isValidName(lastName) {
		return &ValidationError{Msg: "last name is too long"}
	}
	u.lastName = lastName
	return nil
}

func (u *User) SetEmail(email Email) {
	u.email = email
}

func (u *User) SetCredential(credential *Credential) {
	u.credential = credential
}

func isValidName(name string) bool {
	return minNameLength <= len(name) && len(name) <= maxNameLength
}
