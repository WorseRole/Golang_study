package studygorm

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// 转账函数	账户 A 向账户 B 转账 100 元
func TransferMoney(db *gorm.DB, fromAccountId uint, toAccountId uint, amount float64) error {
	// 开始事物
	// var tx *gorm.DB = db.Begin()
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// 延迟函数处理，在函数结束后执行
	// 如果发生错误则回滚，否则则提交
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. 先检查账户 A 是否存在
	accountFrom := Accounts{}
	// tx.Where("id = ?", fromAccountId).Find(&account) // 用Find 不存在会返回0 用First 不存在会返回错误
	resultFrom := tx.Where("id = ?", fromAccountId).First(&accountFrom)
	if resultFrom.Error != nil {
		tx.Rollback()
		return fmt.Errorf("转钱的账户不存在: %v", resultFrom.Error)
	}
	// 1.1 再判断 A 账户 钱够不
	if accountFrom.Balance < amount {
		tx.Rollback()
		return errors.New("账户余额不足")
	}

	// 2. 再检查 B 账户是否存在
	accountTo := Accounts{}
	// db.Where("id = ?", fromAccountId).Find(&account) // 用Find 不存在会返回0 用First 不存在会返回错误
	resultTo := tx.Where("id = ?", toAccountId).First(&accountTo)
	if resultTo.Error != nil {
		tx.Rollback()
		return fmt.Errorf("转钱的账户不存在: %v", resultTo.Error)
	}

	// 3. 转钱出去 accountFrom  A账户 少 amount
	resultFrom = tx.Model(&accountFrom).Where("id = ? ", accountFrom.ID).Update("balance", accountFrom.Balance-amount)
	if resultFrom.Error != nil {
		tx.Rollback()
		return fmt.Errorf("扣款失败: %v", resultFrom.Error)
	}

	// 4. 转钱进来 accountTo B 账户 多 amount
	resultTo = tx.Model(&accountTo).Where("id = ?", accountTo.ID).Update("balance", accountTo.Balance+amount)
	if resultTo.Error != nil {
		tx.Rollback()
		return fmt.Errorf("存款失败: %v", resultTo.Error)
	}

	// 5. 记录交易信息
	transaction := Transactions{FromAccountId: fromAccountId, Amount: amount, ToAccountId: toAccountId}
	resultTransaction := tx.Create(&transaction)
	if resultTransaction.Error != nil {
		tx.Rollback()
		return fmt.Errorf("记录交易失败: %v", resultTransaction.Error)
	}

	// 6. 最后提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交事务失败: %v", err)
	}

	return nil
}
