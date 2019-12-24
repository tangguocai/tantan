package main

import (
	. "tantan/dbhelper"
	"tantan/model"
)

func init(){
	PgInit()
}

func main(){
	defer func() {
		Pg.Conn.Close()  // 程序停止后，关闭连接
	}()
	userInfo := &model.UserInfo{}
	models := []interface{}{userInfo}
	if err := Pg.CreateTabel(models);err != nil{
		panic(err)
	}
}
