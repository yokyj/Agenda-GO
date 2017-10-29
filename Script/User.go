package User

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

const userItemsFilePath string = "User.json"
const currentUserFilePath string = "Current.txt"
const md5ExePath string = "./MD5"

type userItem struct {
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
	// 初始化
	userItems = make(map[string](userItem))
	CurrentUser = nil
	readJSON()
}

// 新建一个userItem，并返回指针
func newUser(name string, password string,
	email string, phoneNumber string) *userItem {
	newItem := new(userItem)
	newItem.name = name
	newItem.hashPassword = hashFunc(password)
	newItem.email = email
	newItem.phoneNumber = phoneNumber
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
	return nil
}

// LoginUser : 登录用户
// 如果用户名不存在，则返回err = errors.New("name")
// 或者用户名密码不正确，则返回err = errors.New("password")
func LoginUser(name string, password string) error {
	tempUser, nameok := userItems[name]
	// 账号错误
	if !nameok {
		return errors.New("ERROR:The user's name not exists")
	}

	passwordok := tempUser.hashPassword == hashFunc(password)
	// 密码错误
	if !passwordok {
		return errors.New("ERROR:The user's password is wrong")
	}

	// 成功登录
	CurrentUser = &tempUser
	return nil
}

// LogoutUser : 登出用户，如果当前没有用户登录，则返回err
func LogoutUser() error {
	if !IsLogin() {
		return errors.New("ERROR:No registered user")
	}

	CurrentUser = nil
	return nil
}

// ListUsers : 列出当前所有用户名，邮箱，密码并组合成字符串返回
// 如果当前没有用户登录，返回err
func ListUsers() (string, error) {
	if !IsLogin() {
		return "", errors.New("ERROR:No registered user")
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
			i, user.name, user.email, user.phoneNumber)
		outputStr += nextStr
		i++
	}
	// 输出结尾
	outputStr += "All user listed as follow.\n"
	return outputStr, nil
}

// DeleteUser : 删除当前登录用户，删除后当前登录用户置为nil
// 如果当前没有用户登录，返回err
func DeleteUser() error {
	if !IsLogin() {
		return errors.New("ERROR:No registered user")
	}

	delete(userItems, CurrentUser.name)
	CurrentUser = nil
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
	return CurrentUser.name
}

func readJSON() {
	// 解析userItems
	b1, err1 := ioutil.ReadFile(userItemsFilePath)
	if err1 == nil {
		json.Unmarshal(b1, userItems)
	}

	// 解析CurrentUser
	b2, err2 := ioutil.ReadFile(currentUserFilePath)
	if err2 == nil {
		CurrentUser = new(userItem)
		json.Unmarshal(b2, *CurrentUser)
	}
}

func writeJSON() {
	// 写入userItems
	b1, err1 := json.Marshal(userItems)
	if err1 == nil {
		ioutil.WriteFile(userItemsFilePath, b1, 0644)
	}

	// 写入CurrentUser
	if CurrentUser == nil {
		return
	}
	b2, err2 := json.Marshal(*CurrentUser)
	if err2 == nil {
		ioutil.WriteFile(userItemsFilePath, b2, 0644)
	}
}
