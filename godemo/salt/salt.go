package salt

import (
	"fmt"
	"gitlab.com/NebulousLabs/fastrand"
)

//输入用户的密码，然后返回经过sha512 和argon2 + salt的组合后的密码。

//加密分为3个等级 1 使用sha1+salt 2 使用sha512+salt 3 使用sha512> salt+Argon2

const (
	Level1 = 1 // 第一等级
	Level2 = 2 // 第二等级
	Level3 = 3 // 第三等级
)

type encryptValue struct {
	level    int      // 等级
	password string   //用户密码
	sha1     [20]byte //sha1 加密
	sha512   [64]byte //sha512 加密
	salt     string   // 盐
	result   string   // 输出的结果
}

//  Enter password and encryption level，return the Encrypted password
func NewEncrypt(passworld string, level int) (string, error) {
	switch level {
	case Level1:
		return encrypt1(passworld)
	case Level2:
		return encrypt2(passworld)
	case Level3:
		return encrypt3(passworld)
	default:
		return "", fmt.Errorf("The parameter you entered is incorrect,please use goo.Level1 goo.Level2 goo.Level3")
	}
	return "", fmt.Errorf("system error")
}

//内部使用的，用户生成一个encrypt1的东西
func encrypt1(password string) (string, error) {
	e := new(encryptValue)
	e.sha1Method()
	err := e.saltMethod()
	if err != nil {
		return "", nil
	}
	return e.result, nil
}

//内部使用的，用户生成一个encrypt1的东西
func encrypt2(password string) (string, error) {
	e := new(encryptValue)
	e.sha512Method()
	err := e.saltMethod()
	if err != nil {
		return "", nil
	}
	return e.result, nil
}

//内部使用的，用户生成一个encrypt3的东西
func encrypt3(password string) (string, error) {
	e := new(encryptValue)
	e.sha512Method()

	err := e.saltMethod()
	if err != nil {
		return "", nil
	}
	err = e.argon2Method()
	if err != nil {
		return "", nil
	}
	return e.result, nil
}

//sha1
func (e *encryptValue) sha1Method() {

}

// sha512
func (e *encryptValue) sha512Method() {

}

// salt
func (e *encryptValue) saltMethod() error {

fastrand

	return nil
}

// argon2
func (e *encryptValue) argon2Method() error {
	return nil
}
