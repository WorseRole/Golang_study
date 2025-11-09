package main

import (
	// 必须导入 MySQL 驱动，使用下划线 _ 表示只执行初始化
	studygorm "Golang_study/No1/study/studyGorm"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"

	"gorm.io/gorm"
)

func main() {
	//连接数据库：如果有密码，使用您的密码
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Printf("连接失败，%v \n", err)
	// 	log.Println("提示: 可以尝试其他连接方式")
	// 	return
	// }
	// log.Println("GORM + MYSQL 连接成功！")
	/*
	   *

	   	// var version string
	   	// db.Raw("SELECT VERSION()").Scan(&version)
	   	// log.Printf("Mysql 版本:%s \n", version)

	   	// 创建表
	   	// studygorm.Run(db)
	   	// studygorm.RunQuestion1(db)

	   	// 1. 插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"
	   	// var student studygorm.Students = studygorm.Students{Name: "张三", Age: 20, Grade: "三年级"}
	   	// result := db.Create(&student)
	   	// if result.Error != nil {
	   	// 	log.Printf("插入失败: %v \n", result.Error)
	   	// } else {
	   	// 	fmt.Printf("插入成功,学生ID: %d \n", student.ID)
	   	// }

	   	// 2. 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	   	// var studentLists []studygorm.Students
	   	// db.Where("age > ?", 18).Find(&studentLists)
	   	// fmt.Printf("年龄大于18岁的学生共 %d 人 \n", len(studentLists))
	   	// for i, stu := range studentLists {
	   	// 	fmt.Printf("i:%d, ID:%d, 姓名:%s, 年龄:%d, 年级:%s \n", i+1, stu.ID, stu.Name, stu.Age, stu.Grade)
	   	// }

	   	// 3. 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	   	// result := db.Model(&studygorm.Students{}).Where("name = ?", "张三").Update("grade", "四年级")
	   	// if result.Error != nil {
	   	// 	log.Printf("更新失败: %v \n", result.Error)
	   	// } else {
	   	// 	fmt.Printf("更新成功，影响行数: %d \n", result.RowsAffected)
	   	// }

	   	// 4. 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	   	// student := studygorm.Students{Name: "张十三", Age: 13, Grade: "幼儿园"}
	   	// db.Create(&student)
	   	// var studentLists []studygorm.Students
	   	// db.Where("age < ?", 15).Find(&studentLists)
	   	// for _, stu := range studentLists {
	   	// 	fmt.Printf("id: %d, name:%s, age:%d, grade:%s \n", stu.ID, stu.Name, stu.Age, stu.Grade)
	   	// }
	   	// // 删除
	   	// result := db.Where("age < ?", 15).Delete(&studygorm.Students{})
	   	// if result.Error != nil {
	   	// 	log.Printf("删除失败:%v \n", result.Error)
	   	// } else {
	   	// 	fmt.Printf("删除成功 删除 %d 行 \n", result.RowsAffected)
	   	// }
	   	// db.Where("age < ?", 15).Find(&studentLists)
	   	// for _, stu := range studentLists {
	   	// 	fmt.Printf("id: %d, name:%s, age:%d, grade:%s \n", stu.ID, stu.Name, stu.Age, stu.Grade)
	   	// }

	   *
	*/

	// 初始化 表
	// studygorm.InitCreateTableAccounts(db)
	// studygorm.InitCreateTableTransactions(db)
	// InitAccounts(db)
	// 账户 A = 1 ; B = 2
	// var A uint = 1
	// var B uint = 3
	// var amount float64 = 10
	// result := studygorm.TransferMoney(db, A, B, float64(amount))
	// if result != nil {
	// 	log.Printf("result: %s \n", result)
	// }

	/**
		// 连接sqlx
		db, err := studygorm.InitDB(dsn)
		if err != nil {
			log.Fatal("数据库连接失败:", err)
		}
		defer db.Close()
		fmt.Println("sqlx 数据库连接成功")
		// err = studygorm.InitEmployeeData(db)
		// if err != nil {
		// 	log.Fatal("创建测试数据失败:", err)
		// }
		fmt.Println("测试数据准备完成")
		// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
		var employees []studygorm.Employee = make([]studygorm.Employee, 0)
		department := "技术部"
		employees, err = studygorm.QueryEmployeesByDepartment(db, department)
		if err != nil {
			log.Printf("查询失败:%v \n", err)
		} else {
			fmt.Printf("%s的员工信息:\n", department)
			for _, emp := range employees {
				fmt.Printf("员工ID: %d, Name:%s, Department:%s, Salary:%d \n", emp.ID, emp.Name, emp.Department, emp.Salary)
			}
		}
		// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
		emp, err := studygorm.QueryEmployeeBySalaryMax(db)
		if err != nil {
			log.Printf("查询失败:%v \n", err)
		} else {
			fmt.Println("工资最高的员工信息:")
			fmt.Printf("员工ID: %d, Name:%s, Department:%s, Salary:%d \n", emp.ID, emp.Name, emp.Department, emp.Salary)
		}
	**/

	//
	db, err := studygorm.InitBookDB(dsn)
	if err != nil {
		log.Printf("数据库连接失败:%v \n", err)
		return
	}
	// err = studygorm.InitBookTableAndTestData(db)
	// if err != nil {
	// 	log.Printf("初始化数据失败:%v \n", err)
	// 	return
	// }
	price, _ := decimal.NewFromString("50.00")

	books, err := studygorm.QueryBookByPrice(db, price)
	if err != nil {
		log.Printf("数据库根据价格查询大于：%s 查询失败:%v \n", price, err)
		return
	}
	for _, book := range books {
		fmt.Printf("book:(ID: %d, Author: %s, Title: %s, Price: %v )\n", book.ID, book.Author, book.Title, book.Price)
	}
}

func InitAccounts(db *gorm.DB) {
	accounts := []studygorm.Accounts{}
	account1 := studygorm.Accounts{ID: 1, Balance: 100}
	account2 := studygorm.Accounts{ID: 2, Balance: 200}
	accounts = append(accounts, account1)
	accounts = append(accounts, account2)

	for _, account := range accounts {
		db.Create(&account)
	}

}
