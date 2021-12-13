package test

import (
	"fmt"
	"funtester/base"
	"funtester/futil"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"testing"
	"time"
)

var drive *gorm.DB

func init() {
	var err error

	drive, err = gorm.Open("mysql", "root:root123456@(localhost:3306)/funtester?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println(err)
		log.Fatalln("mysql conntect err")
	}
	drive.DB().SetMaxOpenConns(200)
	drive.DB().SetConnMaxLifetime(10 * time.Second)
	drive.DB().SetConnMaxIdleTime(10 * time.Second)
	drive.DB().SetMaxIdleConns(20)
	// 迁移 schema
	drive.AutoMigrate(&Funtester{})
	//注意： AutoMigrate 会创建表，缺少的外键，约束，列和索引，并且会更改现有列的类型（如果其大小、精度、是否为空可更改）。但 不会 删除未使用的列，以保护您的数据。
	//db.AutoMigrate(&User{}, &Product{}, &Order{})
	drive.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Tester{})

}

type Funtester struct {
	gorm.Model
	Name string
	Age  int
}

type Tester struct {
	gorm.Model
	Name string
	Age  int
}

func TestSelect1(t *testing.T) {
	var f Funtester
	drive.First(&f, 34)
	last := drive.Last(&f, "age != 1")
	fmt.Printf("查询到记录数 %d "+base.LINE, last.RowsAffected)
	fmt.Println(f)
	take := drive.Take(&f) //不指定顺序
	fmt.Println(take.RowsAffected)
}

func TestInsert(t *testing.T) {
	drive.Create(&Funtester{Name: "D42", Age: futil.RandomInt(1000)})
	drive.Create(&Tester{Name: "测试" + futil.RandomStr(10), Age: futil.RandomInt(100)})
}

func TestMy(t *testing.T) {

	//// Read
	var f Funtester
	drive.First(&f, 34) // 根据整形主键查找
	//db.First(&f, "age = ?", "21") // 查找 code 字段值为 D42 的记录

	fmt.Println(f.ID)
	rows, _ := drive.Model(&f).Where("age > ?", 20).Rows()
	//rows.
	rows.Close()
	//// Update - 将 product 的 price 更新为 200
	drive.Model(&f).Where("name = ?", "D42").Update("Age", futil.RandomInt(1000))
	//// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - 删除 product
	drive.Model(&f).Where("name = ?", "D42").Delete(&f)
	drive.Delete(&f, 1)
	drive.Close()
}
