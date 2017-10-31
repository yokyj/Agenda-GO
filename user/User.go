<<<<<<< HEAD
package user

import (
	"Agenda-GO/mylog"
=======
package main

import (
>>>>>>> 015ad6e43e4e87e3f14180abcf8579d3401169b1
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

<<<<<<< HEAD
const userItemsFilePath string = "./Json/UserItems.json"
const currentUserFilePath string = "./Json/Current.txt"
=======
const userItemsFilePath string = "data/UserItems.json"
const currentUserFilePath string = "data/Current.txt"
>>>>>>> 015ad6e43e4e87e3f14180abcf8579d3401169b1

type userItem struct {
	// 用户名字
	Name string
	// hash过的密码
	HashPassword string
	// 注册用的邮箱
	Email string
	// 注册用的电话号码
	PhoneNumber string
}

func init() {
	// 初始化
	userItems = make(map[string](userItem))
	CurrentUser = nil
	readJSON()
}

<<<<<<< HEAD
=======
// 测试部分
/*
func main() {
	//IsLogInOutTest()
	RegisterUserTest()
	//IsRegisteredUserTest()
	//DeleteUserTest()
}

func IsLogInOutTest() {
	fmt.Println("Test IsLogin():")
	fmt.Println(IsLogin())
	LoginUser("huziang", "123456")
	fmt.Println(IsLogin())
	LogoutUser()
	fmt.Println(IsLogin())
}
func RegisterUserTest() {
	fmt.Println("Test RegisterUser():")
	err := RegisterUser("DJG", "123456", "3254266353@qq.com", "13719316539")
	fmt.Println(err)
	//err = RegisterUser("houhongxiao", "123645", "325423@qq.com", "12321331213")
	//fmt.Println(err)
	//LoginUser("DJG", "123456")
	//ListUsers()
}

func IsRegisteredUserTest() {
	fmt.Println("Test IsRegisteredUser():")
	fmt.Println(IsRegisteredUser("houhongxiao"))
}

func DeleteUserTest() {
	err := LoginUser("huziang", "123456")
	fmt.Println(err)
	DeleteUser()
	err = LoginUser("houhongxiao", "123645")
	fmt.Println(err)
	ListUsers()
	LogoutUser()
}
*/

>>>>>>> 015ad6e43e4e87e3f14180abcf8579d3401169b1
// 新建一个userItem，并返回指针
func newUser(name string, password string,
	email string, phoneNumber string) *userItem {
	newItem := new(userItem)
	newItem.Name = name
	newItem.HashPassword = hashFunc(password)
	newItem.Email = email
	newItem.PhoneNumber = phoneNumber
	return newItem
}

// 用于密码hash的函数
func hashFunc(hashString string) string {
	// 进行md5加密
	h := md5.New()
	h.Write([]byte(hashString))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// 储存user的map集合
var userItems map[string](userItem)

// CurrentUser : currentUser是当前User，如果没有登录为nil
var CurrentUser *userItem

// IsLogin : 判断当前有没有用户登录，并不是很必要
func IsLogin() bool {
	return CurrentUser != nil
}

// RegisterUser : 注册用户，如果用户名一样，则返回err
func RegisterUser(name string, password string,
	email string, phoneNumber string) error {
	_, ok := userItems[name]
	if ok {
		return errors.New("ERROR:The user has registered")
	}
	userItems[name] = *newUser(name, password, email, phoneNumber)
<<<<<<< HEAD

	writeJSON()
	mylog.AddLog("", "RegisterUser", "", userItems[name].String())
=======
	writeJSON()
>>>>>>> 015ad6e43e4e87e3f14180abcf8579d3401169b1
	return nil
}

// LoginUser : 登录用户
// 如果用户名不存在，则返回err1
// 或者用户名密码不正确，则返回err2
func LoginUser(name string, password string) error {
	if IsLogin() {
		return errors.New("ERROR:Please logout at first")
	}
	tempUser, nameok := userItems[name]
	// 账号错误
	if !nameok {
		return errors.New("ERROR:The user's name not exists")
	}

	passwordok := tempUser.HashPassword == hashFunc(password)
	// 密码错误
	if !passwordok {
		return errors.New("ERROR:The user's password is wrong")
	}

	// 成功登录
	CurrentUser = &tempUser
	writeJSON()
<<<<<<< HEAD
	mylog.AddLog(GetLogonUsername(), "LoginUser", "", "")
=======
>>>>>>> 015ad6e43e4e87e3f14180abcf8579d3401169b1
	fmt.Println("Hi " + name + "!")
	return nil
}

// LogoutUser : 登出用户，如果当前没有用户登录，则返回err
func LogoutUser() error {
	if !IsLogin() {
		return errors.New("ERROR:No registered user")
	}

	CurrentUser = nil
	writeJSON()
<<<<<<< HEAD
	mylog.AddLog(GetLogonUsername(), "LogoutUser", "", "")
=======
>>>>>>> 015ad6e43e4e87e3f14180abcf8579d3401169b1
	fmt.Println("Logout successfully!")
	return nil
}

// ListUsers : 列出当前所有用户名，邮箱，密码并组合成字符串返回
// 如果当前没有用户登录，返回err
func ListUsers() error {
	if !IsLogin() {
		return errors.New("ERROR:No registered user")
	}

	outputStr := ""
	i := 1
	// 输出标题
	nextStr := fmt.Sprintf("%-7s|%-12s|%-17s|%-12s\n",
		"No", "Name", "Email", "Phone")
	outputStr += nextStr
	// 依次输出map中的所有值
	for _, user := range userItems {
		nextStr := fmt.Sprintf("%-7d|%-12s|%-17s|%-12s\n",
			i, user.Name, user.Email, user.PhoneNumber)
		outputStr += nextStr
		i++
	}
	// 输出结尾
	outputStr += "All user listed as follow.\n"
	fmt.Printf("%s", outputStr)
<<<<<<< HEAD
	mylog.AddLog(GetLogonUsername(), "ListUsers", "", "")
=======
>>>>>>> 015ad6e43e4e87e3f14180abcf8579d3401169b1
	return nil
}

// DeleteUser : 删除当前登录用户，删除后当前登录用户置为nil
// 如果当前没有用户登录，返回err
func DeleteUser() error {
	if !IsLogin() {
		return errors.New("ERROR:No registered user")
	}

	delete(userItems, CurrentUser.Name)
<<<<<<< HEAD
	mylog.AddLog(GetLogonUsername(), "DeleteUser", (*CurrentUser).String(), "")
=======
>>>>>>> 015ad6e43e4e87e3f14180abcf8579d3401169b1
	CurrentUser = nil
	writeJSON()
	return nil
}

// IsRegisteredUser 判断当前姓名的用户是否注册
func IsRegisteredUser(name string) bool {
	_, ok := userItems[name]
	return ok
}

// GetLogonUsername 得到当前已登录用户的姓名，如果没有登录，返回""
func GetLogonUsername() string {
	if !IsLogin() {
		return ""
	}
	return CurrentUser.Name
}

func readJSON() {
	// 解析userItems
	b1, err1 := ioutil.ReadFile(userItemsFilePath)
	if err1 == nil {
		json.Unmarshal(b1, &userItems)
	}

	// 解析CurrentUser
	b2, err2 := ioutil.ReadFile(currentUserFilePath)
	if err2 == nil {
		CurrentUser = new(userItem)
		json.Unmarshal(b2, CurrentUser)
	}
}

func writeJSON() {
	// 写入userItems
	b1, err1 := json.Marshal(userItems)
<<<<<<< HEAD

=======
>>>>>>> 015ad6e43e4e87e3f14180abcf8579d3401169b1
	if err1 == nil {
		if _, err := os.Open(userItemsFilePath); err != nil {
			os.Create(userItemsFilePath)
		}
<<<<<<< HEAD
		ioutil.WriteFile(userItemsFilePath, b1, 0755)
	}

	// 写入CurrentUser
	if CurrentUser == nil {
		os.Remove(currentUserFilePath)
=======
		ioutil.WriteFile(userItemsFilePath, b1, 0644)
	}

	// 写入CurrentUser

	if CurrentUser == nil {
		if _, err := os.Open(currentUserFilePath); err == nil {
			os.Remove(currentUserFilePath)
		}
>>>>>>> 015ad6e43e4e87e3f14180abcf8579d3401169b1
		return
	}
	b2, err2 := json.Marshal(*CurrentUser)
	if err2 == nil {
		if _, err := os.Open(currentUserFilePath); err != nil {
			os.Create(currentUserFilePath)
		}
<<<<<<< HEAD
		ioutil.WriteFile(currentUserFilePath, b2, 0755)
	}
}

func (u userItem) String() string {
	return "Name:" + u.Name + "  Email:" + u.Email + "  Phone:" + u.PhoneNumber
=======
		ioutil.WriteFile(currentUserFilePath, b2, 0644)
	}

>>>>>>> 015ad6e43e4e87e3f14180abcf8579d3401169b1
}
