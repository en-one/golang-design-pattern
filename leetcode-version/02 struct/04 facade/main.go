package main

import "fmt"

type IUser interface {
	Login(phone int, code int) (*User, error)
	Register(phone int, code int) (*User, error)
}

//门面模式（外观模式）：提供一个统一接口 里面有多次调用
type IUserFacade interface {
	LoginOrRegister(phone int, code int) error
}

type User struct {
	Name string
}

type UserService struct{}

func (u UserService) Login(phone int, code int) (*User, error) {
	//校验操作
	return &User{Name: "test login"}, nil
}

func (u UserService) Register(phone int, code int) (*User, error) {
	//校验操作
	//创建用户
	return &User{Name: "test register"}, nil
}

func (u UserService) LoginOrRegister(phone int, code int) (*User, error) {
	user, err := u.Login(phone, code)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, err
	}
	return u.Register(phone, code)
}

func main() {
	service := UserService{}
	user, err := service.Login(13001010101, 1234)
	fmt.Println(user, err)

	user, err = service.LoginOrRegister(13001010101, 1234)
	fmt.Println(user, err)

	user, err = service.Register(13001010101, 1234)
	fmt.Println(user, err)

	user, err = service.LoginOrRegister(13001010101, 1234)
	fmt.Println(user, err)
}
