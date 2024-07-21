package mysql

// 生成读取mysql的类
import (
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

type Mysql struct {
    DB *sql.DB
	
}

func (m *Mysql) Init() error {
	fmt.Println("Init Mysql")
    return m.Connect()
}

func NewMysql() *Mysql {
	mysql := &Mysql{}
	mysql.Init()
    return mysql
}

func (m *Mysql) Connect() error {
    db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
    if err != nil {
		log.Fatal(err)
        return err
    }
    m.DB = db
    return nil
}

func (m *Mysql) Close() error {
	if (m.DB == nil) {
	    
		return fmt.Errorf("mysql DB is nil")
	}
    return m.DB.Close()
}

// 查询命令
func (m *Mysql) Query(query string, args ...interface{}) (*sql.Rows, error) {
    return m.DB.Query(query, args...)
}

// 执行命令
func (m *Mysql) Exec(query string, args ...interface{}) (sql.Result, error) {  
	return m.DB.Exec(query, args...)
}

// 向数据库插入数据
func (m *Mysql) Insert(query string, args ...interface{}) (int64, error) {
	if (m.DB == nil) {
	    
		return 0, fmt.Errorf("mysql DB is nil")
	}
	result, err := m.DB.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return result.LastInsertId()
}

// 更新数据库数据
func (m *Mysql) Update(query string, args ...interface{}) (int64, error) {
	result, err := m.DB.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// 删除数据库数据
func (m *Mysql) Delete(query string, args ...interface{}) (int64, error) {
	result, err := m.DB.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// 生成上面函数的例子
// func main() {
//     mysql := NewMysql()
//     err := mysql.Connect()
//     if err != nil {
//         panic(err)
//     }
//     defer mysql.Close()
// 	 // mysql.Insert("INSERT INTO users (name, age) VALUES (?, ?)", "Tom", 20)
// 	 // mysql.Update("UPDATE users SET age = ? WHERE name = ?", 21, "Tom")
// 	  // mysql.Delete("DELETE FROM users WHERE name = ?", "Tom")
// 	    rows, err := mysql.Query("SELECT * FROM users")
// 		 if err != nil {
		    
// 			panic(err)
// 		}
// 		defer rows.Close()
// 		 for rows.Next() {
		    
// 			var name string
// 			var age int
// 			err := rows.Scan(&name, &age)
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Println(name, age)
// 		}
// 				err = rows.Err()
// 				if err != nil {
				    
// 						panic(err)
// 				}
// }


// CREATE TABLE users (
//     id INT AUTO_INCREMENT PRIMARY KEY,
//     username VARCHAR(50),
//     age INT
// );