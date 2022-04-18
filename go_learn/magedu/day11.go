package magedu

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Sql1(deviceName, conn string) {
	db, err := sql.Open(deviceName, conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	fmt.Println(db.Ping()) //数据库连接性测试
	//	fmt.Println(db.Exec(`
	//		create table if not exists testGo(
	//		    id bigint primary key auto_increment,
	//		    name varchar(32) not null default '' comment '名字',
	//		    status int not null default 0 comment '状态'
	//		) engine=innodb default charset utf8mb4;
	//`))
	db.Exec(`update testGo set status = 1`)
}
