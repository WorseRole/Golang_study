package studygorm

import "gorm.io/gorm"

// 假设有两个表：
// accounts 表（包含字段 id 主键， balance 账户余额）
// 和
// transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。

type Accounts struct {
	ID      uint    `gorm:"primaryKey"`
	Balance float64 `gorm:"not null;default:0"`
}

type Transactions struct {
	ID            uint    `gorm:"primaryKey"`
	FromAccountId uint    `gorm:"not null"`
	ToAccountId   uint    `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
}

func InitCreateTableAccounts(db *gorm.DB) {
	db.AutoMigrate(&Accounts{})
}

func InitCreateTableTransactions(db *gorm.DB) {
	db.AutoMigrate(&Transactions{})
}
