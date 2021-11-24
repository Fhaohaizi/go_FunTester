package test

import (
	"fmt"
	"funtester/funhttp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

type Funtester struct {
	gorm.Model
	Name string
	Age  int
}

func TestMy(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root123456@(localhost:3306)/funhttp?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println(err)
		fmt.Println("mysql conntect err")
		return
	}

	// 迁移 schema
	db.AutoMigrate(&Funtester{})

	//// Create
	db.Create(&Funtester{Name: "D42", Age: funhttp.RandomInt(1000)})
	//
	//// Read
	var f Funtester
	db.First(&f, 34) // 根据整形主键查找
	//db.First(&f, "age = ?", "21") // 查找 code 字段值为 D42 的记录

	fmt.Println(f.ID)

	//// Update - 将 product 的 price 更新为 200
	db.Model(&f).Where("name = ?", "D42").Update("Age", funhttp.RandomInt(1000))
	//// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - 删除 product
	db.Model(&f).Where("name = ?", "D42").Delete(&f)
	db.Delete(&f, 1)

}
