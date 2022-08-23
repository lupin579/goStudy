package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Result struct {
	ID       int
	Name     string
	Password string
}

func LoadingTable(db *gorm.DB) {
	db.AutoMigrate(&Result{})
}

func GormOriginQuerySingle(db *gorm.DB, result *Result) {
	db.Raw("select id,name,password from results").Scan(result) //Scan会执行至少一次的解引用操作再判断是否为结构体，再结构体内是否有指针（若有则再次解引用），但是如果传入的参数不是指针而是结构体则会报错
	fmt.Printf("%v\n", result)
}

/*
上下两者的区别：
上面的result是Result类型结构体的指针（地址），解引用后得到Result类型的结构体
下面的result是切片即Slice类型的结构体，如果下面Scan的参数直接传result的话就相当于传入了一个结构体会报错，传入&result即Slice结构体的地址后，首先对&result解引用得到result这个Slice类型的结构体（里面存储的是指向第一个底层数组元素的指针（第一个底层数组元素的地址值），切片的容量和切片的大小）后，
再对该结构体的第一个属性（指向底层数组（此处底层数组每一个数据块存储的就是上面的Result类型的结构体）的指针）进行解引用，才能够得到我们想要的（和上面的result类型结构体一样的）Result类型的结构体
*/

func GormOriginQueryMany(db *gorm.DB, result []Result) {
	db.Raw("select id,name,password from results").Scan(&result) //此处只能传入&result，因为[]Result是切片类型，切片类型的本质是一个结构体，该结构体的第一个属性为一个指向对应类型数组的指针；如果传入的是result则相当于传入了一个Slice类型的结构体，结构体是无法进行解引用操作的（能进行解引用是slice的一个指针属性）所以会报错
	for _, result01 := range result {
		fmt.Printf("%v\n", result01)
	}
}

func GormCreate() {
	dsn := "root:xuan0623@tcp(127.0.0.1:3306)/contract?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect")
	}
	db.AutoMigrate(&Product{})
}

func GormInsert() {
	dsn := "root:xuan0623@tcp(127.0.0.1:3306)/contract?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect")
	}
	db.Create(&Product{
		Code:  "100",
		Price: 100,
	})
}

func GormQuery() {
	dsn := "root:xuan0623@tcp(127.0.0.1:3306)/contract?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect")
	}
	product := Product{}
	db.First(&product, 1)
	fmt.Printf("%v\n", &product)
}

func GormQuery01() {
	dsn := "root:xuan0623@tcp(127.0.0.1:3306)/contract?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect")
	}
	product := Product{}
	db.First(&product, "code=?", "100")
	fmt.Println(product)
}

func Update() {
	dsn := "root:xuan0623@tcp(127.0.0.1:3306)/contract?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect")
	}
	product := Product{}
	db.First(&product, "code=?", "100")
	fmt.Println(product)
	db.Model(&product).Update("Price", 333)
	fmt.Println(product)
}

func del() {
	dsn := "root:xuan0623@tcp(127.0.0.1:3306)/contract?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect")
	}
	// product := Product{}
	db01 := db.Delete(&Product{}, 4)
	fmt.Println(*db01)
}

func del01() {
	dsn := "root:xuan0623@tcp(127.0.0.1:3306)/contract?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect")
	}
	db.Delete(&Product{}, "code = ?", "10010")
}

func linkDatabase() *gorm.DB {
	dsn := "root:xuan0623@tcp(127.0.0.1:3306)/contract?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect")
	}
	return db
}
