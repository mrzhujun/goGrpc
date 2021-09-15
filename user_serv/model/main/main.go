package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"mxshop_servs/user_serv/model"
	"os"
	"strings"
	"time"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {

	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("pass", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(newPassword)
	passwordInfo := strings.Split(newPassword, "$")
	fmt.Println(passwordInfo)
	check := password.Verify("pass", passwordInfo[2], passwordInfo[3], options)
	fmt.Println(check)
	return
	dsn := "root:root@tcp(127.0.0.1:3306)/user_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	//i定义一个表结构，将表结构直接生产对呀的表 - migrations
	_ = db.AutoMigrate(&model.User{})

}
