package mysql

import (
	"fmt"
	"log"
	"testing"
)
// 使用UT需要在这个文件夹中拷贝config/config.yaml文件到这个目录，路径也是config/config.yaml
// 初始化测试，定义一个全局变量

func TestMysql(t *testing.T) {
    var mysql = NewMysql()
	// if mysql.Init() != nil {
	    
	// 	t.Error("mysql init failed")
	// }
	defer mysql.Close()
	
	mysql.Insert("INSERT INTO users (username, age) VALUES (?, ?)", "Tom", 20)
	fmt.Println("test")
	mysql.Update("UPDATE users SET age = ? WHERE username = ?", 21, "Tom")
	rows, _ := mysql.Query("SELECT * FROM users")
	for rows.Next() {
		var id int
		var username string
		var age int
		if err := rows.Scan(&id, &username, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, username, age)
	}
	
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	mysql.Delete("DELETE FROM users WHERE username = ?", "Tom")
}