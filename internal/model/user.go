package model

/**
* user model
* sql:
* CREATE DATABASE IF NOT EXISTS test default charset utf8mb4;
* create table user (id int primary key auto_increment,name varchar(200)) engine=innodb;
* 模拟数据插入
* mysql> insert into user (name) values("xiaoming");
   Query OK, 1 row affected (0.11 sec)

   mysql> insert into user (name) values("hello");
   Query OK, 1 row affected (0.04 sec)
*/
// User model
type User struct {
	ID   int64  `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"type:varchar(200)"`
}

// TableName userservice table.
func (User) TableName() string {
	return "user"
}
