package user

import (
	"errors"
	"fmt"
	"server_login/my_db"
)

type UserInfo struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user *UserInfo) Login() (err error) {
	if user.Name == "" {
		return errors.New("invaild user name")
	}
	if user.Password == "" {
		return errors.New("invaild user password")
	}

	if user.check() {
		return
	}
	return errors.New("user login failed. user name or password error")
}

func (user *UserInfo) Register() (err error) {

	if user.Name == "" {
		return errors.New("invaild user name")
	}

	if user.Password == "" {
		return errors.New("invaild user password")
	}

	if user.Email == "" {
		return errors.New("invaild user Email")
	}

	if user.isExsit() {
		return errors.New("user name or email has existed")
	}

	tmpErr := user.insertRowData()
	if tmpErr != nil {
		return tmpErr
	}
	return
}

func (user *UserInfo) insertRowData() (err error) {
	sqlStr := "insert into userinfo(name, password, email) values(?,?,?)"
	_, tmpErr := my_db.AuthDb.Exec(sqlStr, user.Name, user.Password, user.Email)
	if tmpErr != nil {
		// fmt.Printf("insert failed err:%v\n", err)
		logErr("insert failed", tmpErr)
		return tmpErr
	}
	return
}

func (user *UserInfo) check() bool {
	sqlStr := "select name, password, email from userinfo where name = ? or email = ?"
	var u []UserInfo
	err := my_db.AuthDb.Select(&u, sqlStr, user.Name, user.Email)
	if err != nil {
		logErr("check() | data base", err)
		return false
	}
	if len(u) != 1 || (u[0].Password != user.Password) {
		logErr("User info err : Multi same name or same email", err)
		return false
	}
	return true
}

func (user *UserInfo) isExsit() bool {
	sqlStr := "select name, password, email from userinfo where name = ? or email = ?"
	var u []UserInfo
	err := my_db.AuthDb.Select(&u, sqlStr, user.Name, user.Email)
	if err != nil {
		logErr("get failed", err)
		// ？？停止插入
		return true
	}
	if len(u) > 0 {
		return true
	}
	return false
}

func logErr(errInfo string, err error) {
	fmt.Printf("%v, err : %v\n", errInfo, err)
}
