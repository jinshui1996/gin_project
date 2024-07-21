package service

import (
	"fmt"
	"log"
	"gin_project/utils/cache/lru"
	"gin_project/utils/cache"
	"gin_project/utils/mysql"
)

type Agent struct {
	getter Getter
	cache  *cache.Cache
}

// 为Agent生成Get函数
func (a *Agent) Get(key string) (lru.Value, error) {
	if v, ok := a.cache.Get(key); ok {
		log.Default().Println("get from cache")
		return v, nil
	}
	if v, err := a.getter.Get(key); err == nil {
		a.cache.Add(key, v)
		return v, nil
	} else {
		return nil, err
	}
}

func NewAgent(getter Getter, maxBytes int64) *Agent {
    return &Agent{getter: getter, cache: cache.NewCache(maxBytes)}
}

// 生成Getter的接口
type Getter interface {
	Get(key string) (lru.Value, error)
}

// 生成一个Getter的实现结构体
type GetterFunc struct {
	mydb *mysql.Mysql
}

// 实现GetterFunc的init方法
func (f *GetterFunc) init() {
    f.mydb = mysql.NewMysql()
}

// 实现Getter接口的Get方法
func (f *GetterFunc) Get(key string) (lru.Value, error) {
	rows, err := f.mydb.Query("SELECT username, age FROM users where username = ?", key)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var age int
		err := rows.Scan(&name, &age)
		if err != nil {
			panic(err)
		}
		return lru.String(name), nil
	}
    return lru.String(""), fmt.Errorf("no such key")
}

// 获取一个GetterFunc的实例
func NewGetterFunc() *GetterFunc {
	f := &GetterFunc{}
	f.init()
    return f
}
