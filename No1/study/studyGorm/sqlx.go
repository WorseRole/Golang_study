package studygorm

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，
// 并将结果映射到一个自定义的 Employee 结构体切片中。
func QueryEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee

	query := "select id, name, department, salary from employees where department = ?"
	err := db.Select(&employees, query, department)
	if err != nil {
		return nil, fmt.Errorf("查询员工为%s 的部门失败: %v", department, err)
	}
	return employees, nil
}

// 这里只能拿一个
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，
// 并将结果映射到一个 Employee 结构体中。

func QueryEmployeeBySalaryMax(db *sqlx.DB) (*Employee, error) {
	employee := Employee{}

	query := "select * from employees order by salary desc limit 1"
	err := db.Get(&employee, query)
	if err != nil {
		return nil, fmt.Errorf("查询最高工资失败%v", err)
	}
	return &employee, nil
}
