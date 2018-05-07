package dao

import "database/sql"

func OpenDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/blog_go")
}
