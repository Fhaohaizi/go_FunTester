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

func init0() {
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
	//drive.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Funtester{})//带参数的迁移

}

type Funtester struct {
	gorm.Model
	Name string
	Age  int
}

// TestSelect1
// @Description: 普通is寻
// @param t
func TestSelect1(t *testing.T) {
	var f Funtester
	drive.First(&f, 34)                //默认id
	last := drive.Last(&f, "age != 1") //添加条件
	fmt.Printf("查询到记录数 %d "+base.LINE, last.RowsAffected)
	fmt.Println(f)
	take := drive.Take(&f) //不指定顺序
	fmt.Println(take.RowsAffected)
}

// TestSelect2
// @Description: 常用查询和处理结果
// @param t
func TestSelect2(t *testing.T) {
	var fs []Funtester
	var f Funtester
	drive.Where("id = ?", 45).First(&f) //另外一种写法
	//fmt.Println(f)
	find := drive.Where("name like ?", "fun%").Find(&fs).Limit(10).Order("id") //多查询条件串联
	rows, _ := find.Rows()                                                     //获取结果
	defer rows.Close()
	for rows.Next() {
		var ff Funtester
		drive.ScanRows(rows, &ff)
		fmt.Println(ff.Age, ff.Name)
	}
	//另外一种写法
	var f1 Funtester
	drive.Where("name LIKE ?", "fun").Or("id = ?", 123).First(&f1)
	fmt.Println(f1)

}

// TestInsert
// @Description: 增加
// @param t
func TestInsert(t *testing.T) {
	value := &Funtester{Name: "FunTester" + futil.RandomStr(10)}
	drive.Create(value)
	drive.Select("name", "age").Create(value)                           //只创建name和age字段的值
	drive.Omit("age", "name").Create(&Funtester{Name: "fds", Age: 122}) //过滤age和name字段创建
	fs := []Funtester{{Name: "fs" + futil.RandomStr(10), Age: 12}, {Name: "fs" + futil.RandomStr(10), Age: 12}}
	drive.Create(&fs) //这里不支持这么操作的
}

// TestUpdate
// @Description: 更新
// @param t
func TestUpdate(t *testing.T) {
	drive.Model(&Funtester{}).Where("id = ?", 241860).Update("name", base.FunTester+"3")
}

// TestDelete
// @Description: 删除
// @param t
func TestDelete(t *testing.T) {
	db := drive.Where("id = ?", 241859).Delete(&Funtester{})
	fmt.Println(db.RowsAffected)
}

// TestSql
// @Description: 直接执行SQL
// @param t
func TestSql(t *testing.T) {
	var funtester []Funtester
	scan := drive.Raw("select * from funtesters where id > 333 limit 10").Scan(&funtester)
	fmt.Println(scan.RowsAffected)
	fmt.Println(funtester)
}

// TestRollBack
// @Description: 事务&回滚
// @param t
func TestRollBack(t *testing.T) {
	funtester := Funtester{Name: base.FunTester, Age: 32232}
	begin := drive.Begin()
	err := begin.Create(&funtester).Error
	if err != nil {
		begin.Rollback()
	}
	begin.Commit()
}
