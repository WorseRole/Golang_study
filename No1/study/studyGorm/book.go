package studygorm

import (
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
type Books struct {
	ID     int             `db:"id"`
	Title  string          `db:"title"`
	Author string          `db:"author"`
	Price  decimal.Decimal `db:"price"`
}

func InitBookDB(dataSourceName string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitBookTableAndTestData(db *sqlx.DB) error {
	_, err := db.Exec(`create table if not exists books (
		id int auto_increment primary key,
		title varchar(200) not null,
		author varchar(100) not null,
		price decimal(15, 2)
	)`)
	if err != nil {
		return err
	}

	// 插入一批测试数据
	_, err = db.Exec(`insert into books (id, title, author, price) values 
			(1, 'Go语言编程', '张三', 59.99),
    		(2, '数据库原理', '李四', 79.50),
    		(3, 'Web开发实战', '王五', 89.00)
	`)
	return err
}

// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
func QueryBookByPrice(db *sqlx.DB, price decimal.Decimal) ([]Books, error) {
	var books []Books = []Books{}

	query := "select * from books where price > ?"
	err := db.Select(&books, query, price)
	if err != nil {
		return nil, err
	} else {
		return books, nil
	}

}
