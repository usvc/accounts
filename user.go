package main

type User struct {
	Uuid     string `json:"uuid"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ErrorUserNotFound struct {
	error
}

var user = User{}

func (user *User) GetByUuid(uuid string) (User, error) {
	logger.info("getUser()")
	return User{
		Uuid:     uuid,
		Email:    "todo@todo.com",
		Username: "username",
	}, nil
}

func (user *User) Create(newUser User) (User, error) {
	logger.info("createUser()")
	return User{
		Uuid:     "uuid",
		Email:    "todo@todo.com",
		Username: "username",
	}, nil
}

func (user *User) Delete(id uint) error {
	logger.info("createUser()")
	return nil
}
