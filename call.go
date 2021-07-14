package main

import (
	"encoding/json"
	"fmt"
	"github.com/djumanoff/amqp"
	users2 "github.com/kirigaikabuto/dockerComposeRabbitCluster/users"
)


func main() {
	rabbitConfig := amqp.Config{
		Host:     "localhost",
		Port:     5673,
		LogLevel: 5,
		User: "admin",
		Password: "admin",
	}
	sess := amqp.NewSession(rabbitConfig)
	err := sess.Connect()
	if err != nil {
		panic(err)
		return
	}
	clt, err := sess.Client(amqp.ClientConfig{})
	if err != nil {
		panic(err)
		return
	}
	userForCreate := &users2.User{
		Username:  "123",
		Password:  "123",
		FirstName: "123",
		LastName:  "123",
		Avatar:    "123",
	}
	createUserJson, err := json.Marshal(userForCreate)
	if err != nil {
		panic(err)
		return
	}
	_, err = clt.Call("create_user", amqp.Message{Body: createUserJson})
	if err != nil {
		panic(err)
		return
	}
	users := []users2.User{}
	response, err := clt.Call("list_users", amqp.Message{})
	if err != nil {
		panic(err)
		return
	}
	err = json.Unmarshal(response.Body, &users)
	if err != nil {
		panic(err)
		return
	}
	for _, v := range users {
		fmt.Println(v.Id, v.Username)
	}
}
