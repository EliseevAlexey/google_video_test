package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

var DbSchema = "local"

type Config struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type User struct {
	ID        uint `gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     *string
	Age       uint8
	Surname   *string
	Config    Config `gorm:"type:json"`
}

func (User) TableName() string {
	return fmt.Sprintf("%s.users", DbSchema)
}

func main() {
	os.Setenv("DB_SCHEMA", "local")

	dsn := "host=localhost port=25432 user=postgres password=postgres_password dbname=postgres"
	db, _ := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: os.Getenv("DB_SCHEMA"),
			},
		},
	)
	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}

	user := User{Name: "Alex", Age: 36, Config: Config{Key: "SomeKey2", Value: "SomeValue2"}}
	result := db.Create(&user)

	fmt.Printf(`UserID=%d, Error=%s, AffectedRows=%d`, user.ID, result.Error, result.RowsAffected)

	var result2 User
	db.Model(User{Age: 36}).First(&result2)
}
