package main

type User struct {
	ID   uint
	Name string
}

func getUser(id uint) {
	logger.info("getUser()")
}

func createUser(user User) {
	logger.info("createUser()")
}

func deleteUser(id uint) {
	logger.info("createUser()")
}
