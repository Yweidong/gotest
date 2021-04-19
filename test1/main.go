package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	_ "github.com/pkg/errors"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return
	}
	var username string

	// 因为查询单条数据时, 可能返回var ErrNoRows = errors.New("sql: no rows in result set")该种错误信息
	// 而这属于正常错误,不需要向上层抛出
	err = db.QueryRow("select username from user ").Scan(&username)

	switch  {
	case err == sql.ErrNoRows:
	case err != nil: //而当err == nil 的时候 代表执行方法的时候出错，有可能sql语句写错，这时候需要抛出

		 errors.Wrap(err,"sql is error")
	}



}
