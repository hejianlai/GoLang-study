package main

import (
	`database/sql`
	`encoding/json`
	`errors`
	`fmt`
	_ "github.com/go-sql-driver/mysql"
	`strings`
)

// 数据库配置
const   (
	userName = "root"
	password = "123456"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "test"
)
// db数据库连接池
var DB *sql.DB

type User struct {
	id int64
	name string
	age int8
	sex int8
	phone string
}
// 注意方法名大写，就是public
func InitDB()  {
	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName,":",password,"@tcp(",ip,":",port,")/",dbName,"?charset=utf8"},"")
	// 打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB ,_ = sql.Open("mysql",path)
	// 设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	// 设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connect success!!")
}
// 查询操作
func Query()  {
	var  user  User
	rows, e := DB.Query("select * from user where id in (1)")
	if e == nil{
		errors.New("query incur error!!")
	}
	for rows.Next(){
		e := rows.Scan(user.sex, user.phone,user.name,user.id,user.age)
		if e != nil{
			fmt.Println(json.Marshal(user))
		}
	}
	rows.Close()
	DB.QueryRow("select * from user where id=3").Scan(user.age,user.id,user.name,user.phone,user.sex)
	stmt, e := DB.Prepare("select * from user where id=?")
	query, e := stmt.Query(1)
	query.Scan()
}
func main()  {
	InitDB()
	Query()
}
