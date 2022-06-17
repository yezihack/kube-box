package pkg

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	con *sql.DB
)

type MySQLHandlerInfo struct {
	Host     string // 连接主机名
	Port     int    // 端口号
	User     string // 用户名称
	Password string // 密码
	Database string // 数据库名称
}

type MySQLHandler struct {
	db       *sql.DB
	database string
}

func NewMySQLHandler() *MySQLHandler {
	return &MySQLHandler{}
}

// 连接数据库
func (m *MySQLHandler) Connect(info *MySQLHandlerInfo) (err error) {
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		info.User, info.Password, info.Host, info.Port,
	)
	log.Printf("Connect Info:%s\n", conStr)
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		return
	}
	// 设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	// 设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	// 赋值对象DB
	m.db = db
	// set database
	m.database = info.Database
	return
}

// Ping
func (m *MySQLHandler) Ping() bool {
	err := m.db.Ping()
	if err != nil {
		log.Printf("MySQL: Ping is error, err:%v\n", err)
		return false
	}
	return true
}

// 创建数据库
func (m *MySQLHandler) CreateDBName() (err error) {
	sqlFormat := "CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci"
	sqlFormat = fmt.Sprintf(sqlFormat, m.database)
	log.Printf("SQL: %s, database: %s\n", sqlFormat, m.database)
	_, err = m.db.Exec(sqlFormat)
	if err != nil {
		return
	}
	return
}

// 关闭数据库
func (m *MySQLHandler) Close() error {
	return m.db.Close()
}
