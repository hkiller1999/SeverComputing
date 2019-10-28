package entity

import (
	"encoding/json"
	"fmt"
	"os"
)

const userFile string = "curUser.txt"

type User struct {
	userName, passWord, email, phone string
	isLogIn                          bool
}

func (user *User) Init(username, password, email, phone string) {
	user.userName = username
	user.passWord = password
	user.email = email
	user.phone = phone
	user.isLogIn = false
}

func (user User) Find() bool {
	file, err := os.OpenFile(userFile, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return false
	}
	dec := json.NewDecoder(file)
	for {
		var v = make(map[string]interface{})
		if err := dec.Decode(&v); err != nil {
			if err.Error() == "EOF" {
				break
			}
			return false
		} else if user.userName == v["userName"] {
			return false
		}
	}
	file.Close()
	return true
}

func (user User) LogIn() bool {
	file, err := os.OpenFile(userFile, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return false
	}
	dec := json.NewDecoder(file)
	enc := json.NewEncoder(file)
	var jsonArray []map[string]interface{}
	isFind := false
	for {
		var v = make(map[string]interface{})
		if err := dec.Decode(&v); err != nil {
			if err.Error() == "EOF" {
				break
			}
			return false
		} else {
			if user.userName == v["userName"] {
				if user.passWord == v["passWord"] {
					v["isLogIn"] = true
					isFind = true
				} else {
					return false
				}
			}
		}
		jsonArray = append(jsonArray, v)
	}
	if !isFind {
		return false
	}
	file.Seek(0, 0)
	file.Truncate(0)
	for _, v := range jsonArray {
		enc.Encode(&v)
	}
	file.Close()
	return true
}

func (user User) PrintUser() {
	fmt.Println("User:")
	fmt.Println("	username: ", user.userName)
	fmt.Println("	email: ", user.email)
	fmt.Println("	phone: ", user.phone)
}

func (user User) Register() bool {
	file, err := os.OpenFile(userFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		return false
	}
	enc := json.NewEncoder(file)
	var v = make(map[string]interface{})
	v["userName"] = user.userName
	v["passWord"] = user.passWord
	v["email"] = user.email
	v["phone"] = user.phone
	v["isLogIn"] = false
	if err := enc.Encode(&v); err != nil {
		return false
	}
	file.Close()
	return true
}