package users

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
)

type UsersAmqpEndpoints struct {
	usersStore UsersStore
}

func NewUsersAmqpEndpoints(u UsersStore) UsersAmqpEndpoints {
	return UsersAmqpEndpoints{usersStore: u}
}

func (u *UsersAmqpEndpoints) ListUsers() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		users, err := u.usersStore.List()
		if err != nil {
			panic(err)
		}
		dataJson, err := json.Marshal(users)
		if err != nil {
			panic(err)
		}
		return &amqp.Message{Body: dataJson}
	}
}

func (u *UsersAmqpEndpoints) CreateUser() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		jsonData := message.Body
		var user *User
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			panic(err)
		}
		newUser, err := u.usersStore.Create(user)
		if err != nil {
			panic(err)
		}
		dataJson, err := json.Marshal(newUser)
		return &amqp.Message{Body: dataJson}
	}
}
