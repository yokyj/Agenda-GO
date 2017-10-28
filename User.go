package User

import (
	"errors"
	"fmt"
)

type UserItem struct {
	// 用户名字
	name string
	// hash过的密码
	hashPassword string
	// 注册用的邮箱
	email string
	// 注册用的电话号码
	phoneNumber string
}

func init() {
	readJson()
}

// 新建一个UserItem，并返回指针
func newUser(name string, password string,
	email string, phoneNumber string) *UserItem {
	var newItem *UserItem = new(UserItem)
	newItem.name = name
	newItem.hashPassword = hashFunc(password)
	newItem.email = email
	newItem.phoneNumber = phoneNumber
	return newItem
}

// 用于密码hash的函数
func hashFunc(password string) string {
	return password
}

// 储存user的map集合
var userItems map[string](*UserItem) = make(map[string](*UserItem))

// currentUser是当前User，如果没有登录为nil
var CurrentUser *UserItem = nil

// 判断当前有没有用户登录，并不是很必要
func IsLogin() bool {
	return CurrentUser != nil
}

// 注册用户，如果用户名一样，则返回err
func RegisterUser(name string, password string,
	email string, phoneNumber string) error {
	_, ok := userItems[name]
	if ok {
		return errors.New("The user has registered!")
	}
	userItems[name] = newUser(name, password, email, phoneNumber)
	return nil
}

// 登录用户
// 如果用户名不存在，则返回err = errors.New("name")
// 或者用户名密码不正确，则返回err = errors.New("password")
func LoginUser(name string, password string) error {
	tempUser, nameok := userItems[name]
	// 账号错误
	if !nameok {
		return errors.New("name")
	}

	passwordok := tempUser.hashPassword == hashFunc(password)
	// 密码错误
	if !passwordok {
		return errors.New("password")
	}

	// 成功登录
	CurrentUser = tempUser
	return nil
}

// 登出用户，如果当前没有用户登录，则返回err
func LogoutUser() error {
	if !IsLogin() {
		return errors.New("No registered user!")
	}

	CurrentUser = nil
	return nil
}

// 列出当前所有用户名，邮箱，密码并组合成字符串返回
// 如果当前没有用户登录，返回err
func ListUsers() (string, error) {
	if !IsLogin() {
		return "", errors.New("No registered user!")
	}

	outputStr := ""
	appendStr := ""
	i := 1
	// 输出标题
	appendStr = fmt.Sprintf("%-7s|%-12s|%-17s|%-12s\n",
		"No", "Name", "Email", "Phone")
	outputStr += appendStr
	// 依次输出map中的所有值
	for _, user := range userItems {
		appendStr = fmt.Sprintf("%-7s|%-12s|%-17s|%-12s\n",
			i, user.name, user.email, user.phoneNumber)
		outputStr += appendStr
		i++
	}
	// 输出结尾
	outputStr += "All user listed as follow.\n"
	return outputStr, nil
}

// 删除当前登录用户，删除后当前登录用户置为nil
// 如果当前没有用户登录，返回err
func DeleteUser() error {
  if !IsLogin() {
    return errors.new("No registered user!")
  }

  delete(userItems, CurrentUser.name)
  CurrentUser = nil
  return nil
}

// 判断当前姓名的用户是否注册
func IsRegisteredUser(name string) bool {
  _, ok = userItems[name]
  return ok
}

// 得到当前已登录用户的姓名，如果没有登录，返回""
func GetLogonUsername() string {
  return CurrentUser != nil ? CurrentUser.name : ""
}

func readJson() {

}

func writeJson() {

}
