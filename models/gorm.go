package models

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func InitDB() {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(mysql_host:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")

	var err error
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"

	fmt.Println("* Opening Mysql database...")
	db, err = gorm.Open("mysql", CONNECT)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("* Mysql database opened!!")
	db.LogMode(true)
	InitTables(db)
}

func TruncateTables(db2 *gorm.DB) {
	rows, err := db2.Raw("show tables").Rows()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			panic(err.Error())
		}
		db.Exec("TRUNCATE TABLE " + table)
	}
}

func InitTables(db2 *gorm.DB) {
	tx := db2.Begin()
	TruncateTables(tx)
	createdAt, err := time.Parse("2006-01-02 15:04:05", "2019-08-30 00:47:40")
	if err != nil {
		log.Fatalln(err.Error())
	}
	birth, err := time.Parse("2006-01-02", "2019-07-30")
	if err != nil {
		log.Fatalln(err.Error())
	}
	initDishRecords := Dish{
		ID:        4,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		DishName:  "カレー",
		StoreName: "マサラ亭",
	}
	initFollowsRecords := FollowsDBModel{
		ID:        1,
		UserID:    1,
		FollowID:  1,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}
	initPostsRecords := Post{
		ID:           1,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
		UserId:       1,
		DishId:       4,
		Comment:      "hogehoge",
		ImageAddress: "img/dishes/42.jpg",
	}
	initUsersRecords := User{
		ID:          0,
		CreatedAt:   createdAt,
		UpdatedAt:   createdAt,
		Name:        "hogehoge",
		Biography:   "ニンニクなしで、野菜マシ、アブラ、カラメで～",
		Birth:       birth,
		Country:     "Jap",
		Prefecture:  "Tokyo",
		IconAddress: "img/gazou/abc/123.jpg",
	}
	tx.Table("Dishes").Create(&initDishRecords)
	tx.Table("Follows").Create(&initFollowsRecords)
	tx.Table("Posts").Create(&initPostsRecords)
	tx.Table("Users").Create(&initUsersRecords)
	//tx.Exec("INSERT INTO `Dishes` (`id`, `created_at`, `updated_at`, `deleted_at`, `dish_name`, `store_name`)VALUES(1, '2019-08-30 00:47:40', '2019-08-30 00:47:40', NULL, 'カレー', 'マサラ亭'),(2, '2019-08-30 00:47:47', '2019-08-30 00:47:47', NULL, 'パスタ', 'jig.jp'),(3, '2019-08-30 00:49:12', '2019-08-30 00:49:50', NULL, 'フランス産フォアグラのテリーヌ　トリュフとブッフサレ　リ・ド・ヴォとレンズ豆のガトー仕立て', '日吉で一番ウマいのはあびすけさんだと思うので、当店営業時間外はぜひともあびすけさんをご利用ください。 '),(4, '2019-08-30 00:49:12', '2019-08-30 00:49:50', NULL, 'フランス産フォアグラのテリーヌ　トリュフとブッフサレ　リ・ド・ヴォとレンズ豆のガトー仕立て', 'hogehoge'),(5, '2019-08-30 00:49:12', '2019-08-30 00:50:10', NULL, 'hogehoge', '日吉で一番ウマいのはあびすけさんだと思うので、当店営業時間外はぜひともあびすけさんをご利用ください。 ');")
	//tx.Exec("INSERT INTO `Follows` (`id`, `user_id`, `follow_id`, `created_at`, `updated_at`)VALUES(1, 1, 1, '2019-08-23 08:47:14', '2019-08-23 08:47:14'),(2, 1, 3, '2019-08-23 08:47:20', '2019-08-23 08:47:20'),(4, 1, 2, '2019-08-26 01:41:21', '2019-08-26 01:41:21'),(11, 1, 3, '2019-08-26 06:45:38', '2019-08-26 06:45:38'),(12, 1, 3, '2019-08-26 07:04:41', '2019-08-26 07:04:41'),(13, 1, 3, '2019-08-26 07:06:27', '2019-08-26 07:06:27'),(14, 1, 3, '2019-08-26 07:29:35', '2019-08-26 07:29:35'),(15, 1, 3, '2019-08-26 07:30:02', '2019-08-26 07:30:02');")
	//tx.Exec("INSERT INTO `Posts` (`id`, `user_id`, `created_at`, `updated_at`, `dish_id`, `comment`, `image_address`)VALUES(1, 1, '2019-08-29 05:01:31', '2019-08-29 05:26:10', 4, 'hogehoge', 'img/dishes/42.jpg'),(7, 2, '2019-08-29 17:12:03', '2019-08-30 00:33:21', 4, 'it was delicious!', ''),(8, 2, '2019-08-29 17:12:03', '2019-08-30 00:51:55', 4, '', ''),(9, 2, '2019-08-29 17:12:03', '2019-08-30 00:54:23', 4, NULL, 'img/gazou/abc/123.jpg');")
	//tx.Exec("INSERT INTO `Users` (`id`, `name`, `created_at`, `updated_at`, `biography`, `birth`, `country`, `prefecture`, `icon_address`)VALUES(1, 'hogehoge', '2019-08-27 07:36:43', '2019-08-27 07:46:18', NULL, '2019-07-30', 'Jap', 'Fukui', NULL),(2, 'hugahgua', '2019-08-27 07:36:43', '2019-08-28 05:18:44', NULL, '2019-07-30', 'Jap', 'Tokyo', NULL),(3, 'hugahgua', '2019-08-27 07:36:43', '2019-08-30 00:53:52', 'ニンニクなしで、野菜マシ、アブラ、カラメで～', '2019-07-30', 'Jap', 'Tokyo', 'img/gazou/abc/123.jpg'),(4, 'hugahgua', '2019-08-27 07:36:43', '2019-08-30 00:53:52', '', '2019-07-30', 'Jap', 'Tokyo', 'img/gazou/abc/123.jpg');")
	if tx.Error != nil {
		fmt.Println("\x1b[31mstarting transition failed.\x1b[0m")
		tx.Rollback()
		return
	}
	tx.Commit()
}

func GetDB() (*gorm.DB, error) {
	if db == nil {
		return nil, errors.New("database reference is null")
	}
	return db, nil
}

func Finalize() error {
	err := db.Close()
	return err
}
