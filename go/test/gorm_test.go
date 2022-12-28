package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestGorm(t *testing.T) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"119129",
		"175.178.218.166",
		3306,
		"pxjCourseSystem",
	)
	fmt.Printf("dsn: %v", dsn)
	db, _ := gorm.Open("mysql", dsn)

	db.AutoMigrate(&Product{})
	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	db.Delete(&product, 1)
}
