package studygorm

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	// ignored      string         // fields that aren't exported are ignored
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	// db.AutoMigrate(&Member{})
	// db.AutoMigrate(&Blog{})
	// db.AutoMigrate(&Blog2{})

	// user := &User{}
	// user.MemberNumber.Valid = true

	birthday := time.Now() // 因为struct中的birthday 为指针格式 所以创建user时也需要 birthday为指针
	user := User{Name: "leoYan", Age: 10, Birthday: &birthday}
	result := db.Create(&user) // 通过数据指针来创建
	fmt.Println(result.RowsAffected)

	// create传指针
	// mem := Member{}
	// db.Create(&mem)
	// fmt.Println(mem.ID)
	// db.Delete(&Member{}, 1)
}
