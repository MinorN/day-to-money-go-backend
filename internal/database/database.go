package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "pg-for-go-mangosteen"
	port     = 5432
	user     = "mangosteen"
	password = "123456"
	dbname   = "mangosteen_dev"
)

func Connect() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	DB = db

}

type User struct {
	ID       int
	Email    string `gorm:"uniqueIndex"`
	Phone    string
	Address  string
	CreateAt time.Time
	UpdateAt time.Time
}

type Item struct {
	ID         int
	UserID     int
	Amount     int
	HappenedAt time.Time
	CreateAt   time.Time
	UpdateAt   time.Time
}

type Tag struct {
	ID   int
	Name string
}

var models = []any{&User{}, &Item{}, &Tag{}}

func CreateTables() {
	for _, model := range models {
		err := DB.Migrator().CreateTable(model)
		if err != nil {
			log.Println(err)
		}
	}
}

func Migrate() {
	DB.AutoMigrate(models...)
}

func Crud() {
	// // 创建一个User
	// user := User{Email: "test3@qq.com"}
	// tx := DB.Create(&user)
	// log.Println(tx.RowsAffected)
	// log.Println(user)

	// u2 := User{}
	// tx = DB.Find(&u2, 1)
	// u2.Phone = "12345678"
	// tx = DB.Save(&u2)
	// if tx.Error != nil {
	// 	log.Println(tx.Error)
	// } else {
	// 	log.Println(tx.RowsAffected)
	// 	log.Println(u2)
	// }

	users := []User{}
	DB.Find(&users, []int{1, 3, 6})
	log.Println(users)

	u := User{ID: 1}
	tx := DB.Delete(&u)
	if tx.Error != nil {
		log.Println(tx.Error)
	} else {
		log.Println(tx.RowsAffected)
	}
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.Close()
}
