package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    mds "github.com/catman/go_api/models"
    "fmt"
    "time"
    "math/rand"
    "log"
)
var db  *gorm.DB

func main() {
    r := gin.Default()
    r.GET("/users", func(c *gin.Context) {
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
    })

    r.GET("/", func(c *gin.Context) {
        var product mds.Product
        product.Code = "L1212"
        samsam := db.First(&product)
        fmt.Println(samsam)
        c.String(200, "Hello world!%v", samsam)
    })
    r.Run(":30000")
}

func init(){
    // 初期化時にDBと接続
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()
    db.AutoMigrate(&mds.User{}, &mds.Item{})
    db.AutoMigrate(&mds.Product{})
    db.Create(&mds.Product{Code: "L1212", Price: 1000})
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
        items[i] = NewItem(i, name)
    }

    return items
}
func NewItem(name string) mds.Item {
    rand.Seed(time.Now().UnixNano())
    sample := mds.Item{
        UserID: 1,
        Name:  name,
        Score: rand.Intn(10) + 1,
    }
    log.Printf("俺だよ%v", sample)
    return sample
}
