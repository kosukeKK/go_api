package users

import (
	"github.com/gin-gonic/gin"
	mds "github.com/catman/go_api/models"
	"github.com/jinzhu/gorm"
	"time"
	"math/rand"
)

var db  *gorm.DB
func Handler(c *gin.Context) {
	var (
		users   []mds.User
		items   []mds.Item
		jsonMap map[string]interface{} = make(map[string]interface{})
	)

	db.Find(&users)

	for i, user := range users {
		db.Model(&user).Related(&items)
		user.Items = items
		users[i] = user
	}

	jsonMap["users"] = users
	c.JSON(200, jsonMap)
}

func init(){
	// 初期化時にDBと接続
	db, _ = gorm.Open("sqlite3", "test.db")
	db.AutoMigrate(&mds.User{}, &mds.Item{})
	itemNames := []string{
		"ゲーム1",
		"ゲーム2",
		"ゲーム3",
		"ゲーム4",
		"ゲーム5",
		"ゲーム6",
		"ゲーム6",
		"ゲーム7",
	}

	userNames := []string{
		"ユーザー1",
		"ユーザー2",
		"ユーザー3",
		"ユーザー4",
		"ユーザー5",
		"ユーザー6",
		"ユーザー7",
		"ユーザー8",
		"ユーザー9",
		"ユーザー10",
	}

	users := CreateUsers(userNames, itemNames)

	count := 0
	db.Table("users").Count(&count)
	if count == 0 {
		for _, user := range users {
			db.Create(&user)
		}
	}
}

func CreateUsers(userNames []string, itemSlice []string) []mds.User {
	users := make([]mds.User, len(userNames))

	for i, name := range userNames {
		users[i] = NewUser(name, itemSlice)
	}

	return users
}

func NewUser(name string, itemSlice []string) mds.User {
	return mds.User{
		Name:  name,
		Items: CreateItems(itemSlice),
	}
}

func CreateItems(itemSlice []string) []mds.Item {
	items := make([]mds.Item, len(itemSlice))

	for i, name := range itemSlice {
		items[i] = NewItem(name)
	}

	return items
}
func NewItem(name string) mds.Item {
	rand.Seed(time.Now().UnixNano())

	return mds.Item{
		Name:  name,
		Score: rand.Intn(10) + 1,
	}
}