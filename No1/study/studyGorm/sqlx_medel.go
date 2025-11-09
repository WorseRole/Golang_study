package studygorm

import "github.com/jmoiron/sqlx"

// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，
// 并将结果映射到一个自定义的 Employee 结构体切片中。

// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，
// 并将结果映射到一个 Employee 结构体中。

type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

// 初始化数据库连接
func InitDB(dataSourceName string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// init Employee
func InitEmployeeData(db *sqlx.DB) error {
	_, err := db.Exec(`create table if not exists employees (
		id int auto_increment primary key,
		name varchar(100) not null,
		department varchar(100) not null,
		salary int not null
	)`)
	if err != nil {
		return err
	}
	// 插入数据
	_, err = db.Exec(`insert into employees (name, department, salary) values 
		('张三', '技术部', 15000),
		('王五', '技术部', 18000),
		('李四', '技术部', 25000),
		('赵六', '销售部', 5000),
		('钱七', '人事部', 11000),
		('孙八', '技术部', 8000)
	`)
	return err
}
