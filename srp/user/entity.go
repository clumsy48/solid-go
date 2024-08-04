package user

// user definition
type User struct  {
  fullName string
  email string
}

// initialie user
func (user *User) CreateUser(fullName string, email string) {
   user.fullName = fullName
   user.email = email
}

func (user User) GetFullName() string {
  return user.fullName
}

func (user User) GetEmail() string {
  return user.email
}



