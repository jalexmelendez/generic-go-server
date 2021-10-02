package main

type User struct {
	id       int64
	mail     string
	password string
}

type UserMethods interface{}

func getUser() interface{} {
	Res := []interface{}{}
	Res = append(Res, User{id: 1, mail: "example", password: "example"})
	return Res
}
