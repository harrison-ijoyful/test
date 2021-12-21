package main

import (
	"fmt"
	"time"
    //V2
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	UserName     string = "root"
	Password     string = "12345678"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "test"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

type User struct {
    //gorm為model的tag標籤，v2版的auto_increment要放在type裡面，v1版是放獨立定義
	ID        int64     `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Username  string    `gorm:"type:varchar(20) NOT NULL;" json:"username,omitempty"`
	Password  string    `gorm:"type:varchar(100) NOT NULL;" json:"password,omitempty"`
	Status    int32     `gorm:"type:int(5);" json:"status,omitempty"`
	CreatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}


func main() { 
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", UserName, Password, Addr, Port, Database)    
	//連接MySQL
	conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}  
    
	db, err1 := conn.DB()
	if err1 != nil {
		fmt.Println("get db failed:", err)
		return
	}
	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)	

	conn.Debug().AutoMigrate(&User{})
    
    migrator := conn.Migrator()
    has := migrator.HasTable(&User{})
	//has := migrator.HasTable("GG")
	if !has {
		fmt.Println("table not exist")
	}

	user := User{Username: "tester", Password: "12333", Status: 1}
    //
	result := conn.Debug().Create(&user)
	if result.Error != nil {
		fmt.Println("Create failt")
	}
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number failt")
	}

	users := []User{{Username: "tester", Password: "12333", Status: 1}, {Username: "gger", Password: "132333", Status: 1}, {Username: "ininder", Password: "12333", Status: 1}}

	result = conn.Debug().Create(&users)
	if result.Error != nil {
		fmt.Println("Create failt")
	}
	fmt.Println("result.RowsAffected:", result.RowsAffected)

	var u User
	var au []User
	res := conn.Debug().Find(&au)
	fmt.Println(res.RowsAffected)
	//SELECT * FROM `users`
	conn.Debug().First(&u)
	fmt.Println(u)
	//SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	conn.Debug().Take(&u)
	fmt.Println(u)
	//SELECT * FROM `users` WHERE `users`.`id` = 1 LIMIT 1
	conn.Debug().Last(&u)
	fmt.Println(u)

}