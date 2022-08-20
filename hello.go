package main

import (
	"fmt"
)

func test01() {
	//声明时没有指定数组元素的值，默认为零值
	var arr [5]int
	fmt.Println(arr) //[0 0 0 0 0]

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr) //[1 2 3 0 0]
}

func test02() {
	//直接在声明时对数组进行初始化
	var arr1 = [5]int{15, 20, 25, 30, 35}
	fmt.Println(arr1) //[15 20 25 30 35]

	//使用短声明
	arr2 := [5]int{15, 20, 25, 30, 35}
	fmt.Println(arr2) //[15 20 25 30 35]

	//部分初始化，未初始化部分为零值
	arr3 := [5]int{15, 20}
	fmt.Println(arr3) //[15 20 0 0 0]

	//可以通过指定索引，方便地对数组某几个元素赋值
	arr4 := [5]int{1: 100, 4: 200}
	fmt.Println(arr4) //[0 100 0 0 200]

	//直接使用...让编译器为我们计算该数组的长度
	arr5 := [...]int{15, 20, 25, 30, 35}
	fmt.Println(arr5) //[15 20 25 30 35]
}

func test03() {
	//特别注意数组的长度是类型的一部分，所以[3]int和[5]int是不同的类型
	arr1 := [3]int{15, 20, 25}
	arr2 := [5]int{15, 20, 25, 30, 35}
	fmt.Printf("type of arr1 %T\n", arr1) //type of arr1 [3]int
	fmt.Printf("type of arr2 %T\n", arr2) //type of arr2 [5]int
	//要用printf不能用println
}

func test04() {
	//定义多维数组
	arr := [3][2]string{
		{"1", "tokisaki"},
		{"2", "Tokisaki"},
		{"3", "toKisaki!"}}
	fmt.Println(arr) //[[1 tokisaki] [2 Tokisaki] [3 toKisaki!]]
}

func arrByValue() {
	arr := [...]string{"Go语言极简一本通", "Go语言微服务架构核心22讲", "从0到Go语言微服务架构师"}
	copy := arr
	copy[0] = "Golang"
	fmt.Println(arr)
	fmt.Println("-----------------")

	fmt.Println(copy)
}

func sliceTest() {
	// var slice01 []int
	slice01 := make([]int, 5, 5)
	fmt.Print(slice01)
}

func sliceTest01() {
	var slice01 = []int{1, 11, 1, 1, 1}
	fmt.Print(slice01)
}

func sliceTest02() {
	slice01 := []int{1, 121, 212, 1, 3, 5, 2}
	fmt.Print(slice01)
}

func sliceTest03() {
	slice01 := [][]int{{1, 2}, {2, 2}, {4, 65, 4}, {1, 2, 95}}
	fmt.Print(slice01)
}

func sliceTest04() {
	arr := []int{1, 21, 15, 4, 56, 2} //创建切片时若未规定其容量，则默认设置为切片指向的首元素到底层数组的末尾元素（包含末尾元素）的个数为容量，例见sliceTest05
	slice01 := arr[1:5]               //[21 15 4 56]
	fmt.Println(slice01)
	fmt.Println(cap(arr))
	fmt.Printf("arr=%p\n", arr)
	fmt.Printf("&arr=%p\n", &arr)
	fmt.Println(cap(arr)) //每次扩容都会使容量扩张为*切片*（非底层数组）原容量的两倍（cap>1024后变为1.25倍），
	//若扩容之后切片容量未大于底层数组的容量，则切片的指针依旧指向数组的原本指向元素（不一定是首元素，应该是切片的起始位置在数组中的地址）的地址，若扩容后大于数组容量，
	//则开辟新内存作为新数组并将原数组数据拷贝过来，并让切片的指针指向新数组的原本指向的元素的地址
	//---------------------------------------------------------------------------------------------------------------------------------------------
	// arr = append(arr, 3)//如果切片的容量小于1024个元素，那么扩容的时候slice的cap就乘以2；一旦元素个数超过1024个元素，增长因子就变成1.25，即每次增加原来容量的四分之一。
	// 如果扩容之后，还没有触及原数组的容量，那么，切片中的指针指向的位置，就还是原数组，如果扩容之后，超过了原数组的容量，
	//那么，Go就会开辟一块新的内存，把原来的值拷贝过来，这种情况丝毫不会影响到原数组。
	//---------------------------------------------------------------------------------------------------------------------------------------------
	fmt.Println(cap(arr))
	fmt.Printf("arr=%p\n", arr)
	// fmt.Println(len(arr))
	// fmt.Println(cap(arr))
	// arr[1] = 3333
	// fmt.Println(slice01)
	// var ptr = new(string)
	// *ptr = "string"
	// fmt.Print(ptr, " ", &ptr, " ", *ptr) //ptr:数据块存储的地址（即数据块的内容） &ptr：数据块的地址 *ptr:数据块存储的地址（即内容）指向的位置的数据块的内容
}

func varDemo() {
	var demo01 int = 3
	var demo02 = 3
	demo03 := 3
	print(demo01, demo02, demo03)
}

func sliceTest05() {
	arr := [6]int{1, 21, 15, 4, 56, 2}
	slice01 := arr[:] //创建切片截取整个数组相当于slice01 := arr[0:6]
	slice03 := arr[2:3]
	slice04 := arr[2:6]
	//slice02 := arr    //创建新数组并拷贝旧数组过来（值传递）
	// fmt.Println(arr, "\n", slice01)
	// fmt.Println("-------------------------------------")
	// arr[0] = 333
	// fmt.Println(arr, "\n", slice01, "\n", slice02)//[333 21 15 4 56 2] [333 21 15 4 56 2] [1 21 15 4 56 2]
	fmt.Println(slice01)
	fmt.Println(cap(arr))
	fmt.Println(cap(slice01))
	fmt.Println("---------------------")
	fmt.Println(slice03)
	fmt.Println("---------------------")
	fmt.Println(slice04)
}

/*
map的四种定义方式
*/
func mapTest01() {
	var maps01 map[int]string = map[int]string{
		1: "name",
		2: "gender",
		3: "height",
	}
	var maps02 = map[int]string{
		1: "外币",
		2: "歪比",
		3: "外壁",
	}
	maps03 := map[int]string{
		1: "巴伯",
		2: "八波",
		3: "八伯",
	}
	maps04 := make(map[int]string)
	fmt.Printf("maps01:%v\n", maps01) //maps01:map[1:name 2:gender 3:height]
	fmt.Printf("maps02:%v\n", maps02) //maps02:map[1:外币 2:歪比 3:外壁]
	fmt.Printf("maps03:%v\n", maps03) //maps03:map[1:巴伯 2:八波 3:八伯]
	fmt.Printf("maps04:%v\n", maps04) //maps04:map[]
}

/*
增删改查、遍历、判存在、获取长度
*/
func mapTest02() {
	var maps01 map[int]string = map[int]string{
		1: "name",
		2: "gender",
		3: "height",
	}
	maps01[4] = "width"
	delete(maps01, 1)
	maps01[2] = "3333333"
	temp := maps01[3]
	fmt.Println(temp) //height
	value, ok := maps01[4]
	fmt.Println(value, ok) //width
	value01, ok01 := maps01[5]
	fmt.Println(value01, ok01) // false
	/*
		key:2, value:3333333
		key:3, value:height
		key:4, value:width
	*/
	for key, value := range maps01 {
		fmt.Printf("key:%d, value:%s\n", key, value)
	}
	fmt.Println(len(maps01)) //3
}

/*
map是引用类型的验证
*/
func mapTest03() {
	var maps map[int]string = map[int]string{
		1: "name",
		2: "gender",
		3: "height",
	}
	newMaps := maps
	fmt.Printf("maps:%v\n", maps)
	fmt.Printf("newMaps:%v\n", newMaps)
	fmt.Println("------------------------------------")
	for key, _ := range newMaps {
		maps[key] += "333"
	}
	fmt.Printf("maps01:%v\n", maps)
	fmt.Printf("newMaps01:%v\n", newMaps)
}

func ptrTest01() {
	x := "tokisaki"
	ptr01 := &x
	fmt.Println(ptr01)  //输出存储的地址
	fmt.Println(*ptr01) //输出存储的地址对应的位置的数据块内存储的内容
	fmt.Println("-----------------------")
	ptr02 := new(string)
	*ptr02 = "kurumi"
	fmt.Println(ptr02)
	fmt.Println(*ptr02)
	fmt.Println("-----------------------")
	x01 := "tokisaki"
	var ptr03 *string
	ptr03 = &x01
	var ptr04 *string = &x01
	fmt.Println(ptr03)
	fmt.Println(*ptr03)
	fmt.Println("-----------------------")
	fmt.Println(ptr04)
	fmt.Println(*ptr04)
	fmt.Println("-----------------------")
	var ptr05 *string
	fmt.Println(ptr05)
	//fmt.Println(*ptr05)panic: runtime error: invalid memory address or nil pointer dereference
}

func ptrTest02(value *int) {
	*value = 333
}

// type Demo struct {
// 	name     string
// 	age      int
// 	identity string
// }

//创建匿名结构体和赋值（函数外部）
var Demo struct {
	name, identity string
	age            int
}

func structSpawn02() {
	Demo.name = "ks"
	Demo.identity = "mywife"
	Demo.age = 18
	fmt.Printf("A girl named %s, is %s, aged %v.\n", Demo.name, Demo.identity, Demo.age)
}

//创建命名的结构体和实例化赋值
type Demo01 struct {
	name01, identity01 string
	age01              int
}

func structSpawn() {
	demo01 := Demo01{
		name01:     "kurummi",
		identity01: "mywife",
		age01:      18,
	}
	fmt.Printf("A girl named %s, is %s, aged %v.\n", demo01.name01, demo01.identity01, demo01.age01)
}

//创建匿名结构体和实例化赋值（函数内部）
func structSpawn01() {
	demo02 := struct {
		name02, identity02 string
		age02              int
	}{
		name02:     "tokisaki",
		identity02: "mywife",
		age02:      18,
	}
	fmt.Printf("A girl named %s, is %s, aged %v.\n", demo02.name02, demo02.identity02, demo02.age02)
}

/*
匿名字段（即未命名的属性）（默认将名字设为属性的类型）
*/
type Demo04 struct {
	string
	int
}

func DemoTest04() {
	demo04 := Demo04{}
	demo04.int = 18
	demo04.string = "ks"
	fmt.Println("demo04的int:", demo04.int)
	fmt.Println("demo04的string:", demo04.string)
}

func StructApp(msg *struct {
	name string
	age  int
}) {
	fmt.Println((*msg).age) //相当于msg.age结构体（无论匿名还是命名）自动帮忙解引用
}

/*
嵌套结构体
*/
type Demo05 struct {
	demo01 Demo01
	demo04 Demo04
}

func DemoTest05() {
	tdemo01 := Demo01{
		name01:     "ks",
		identity01: "mw",
		age01:      18,
	}
	tdemo04 := Demo04{
		string: "ks",
		int:    18,
	}
	demo05 := Demo05{
		demo01: tdemo01,
		demo04: tdemo04,
	}
	fmt.Println(demo05) //{{ks mw 18} {ks 18}}
	fmt.Println(demo05.demo01.age01)
}

type Demo06 struct {
	name string
	Demo04
}

func DemoTest06() {
	demo06 := Demo06{
		name: "kurummi",
	}
	demo06.Demo04 = Demo04{
		string: "ks",
		int:    18,
	}
	fmt.Println(demo06.string, demo06.int) //相当于demo06.Demo04.string和demo06.Demo04.int（提升字段）
}

/*
结构体比较
 如果结构体的全部成员都是可以比较的，
 那么结构体也是可以比较的，
 那样的话两个结构体将可以使用 == 或 != 运算符进行比较。
 可以通过==运算符或 DeeplyEqual()函数比较两个结构相同的类型并包含相同的字段值。
 在 Go 中无法在结构体内部定义方法，这一点与 C 语言类似。
 但是将结构体类型设为方法的接收体类型（若想通过方法修改传入的接收体参数的值需要传入指针(demo *Demo)（即传入要修改的结构体对象的地址（&）））
*/

type Cat struct {
	color string
	name  string
}

//返回值为Cat的话，返回的是一个Cat结构体
//相当于是将该Cat里的color和name属性复制一份给了main里的接收结构体Cat（即有两个结构体，main里面的Cat和create函数里的Cat是两个不同的结构体（地址不同），只是这两个结构体的属性的值相同而已（因为main里Cat的属性值，是用create里的Cat的属性值复制过去的））
//返回值为*Cat的话，返回的是Cat结构体的地址（即指针）,
//main函数里的接收结构体接收到的是这个结构体的地址（即只有一份Cat结构体，main里面的Cat是指针保存的是create函数里的Cat的地址）
func catCreate(col string) Cat {
	cat := Cat{
		color: col}
	fmt.Printf("true address is %p\n", &cat)
	return cat
}

/*
接口
*/
type Ark interface {
	Dark(time int) string
	Dawn(time int) string
}

type Hat interface {
	Dark(time int) string
	Dawn(time int) string
}

type day struct {
	ark Ark
}

func (d *day) Dark(time int) string {
	fmt.Println("it is nighttime now!")
	return "dark time"
}

func (d *day) Dawn(time int) string {
	fmt.Println("it is daytime now!")
	return "light time"
}

func interDemo01() {
	d := new(day)
	var ark Ark = d
	fmt.Printf("%p\n", &ark) //除该行外，其余行地址均为0xc000088220（由此可见，创建新的接口类型是申请一块新的内存将结构体指针转换为对应接口类型后存入该内存中）
	d.ark = d
	d.ark.Dark(3)
	fmt.Printf("%p\n", d)
	fmt.Printf("%p\n", &(d.ark)) //创建一个结构体后，所有属性均存放在该地址中
	fmt.Printf("%p\n", d)
	//ark.Dawn(3)
}

/*
接口（类型断言）
首先创建一个实现接口方法的结构体的指针
再将该指针赋值给一个interface{}类型的变量（即将结构体的指针类型转换为接口类型（当该结构体指针实现了接口的所有方法时即可完成转换））
再使用类型断言obj.(T)——obj为转换后的接口类型，T为将要转换为的类型（有两个返回值：转换为的类型变量，是否转换成功）（只有转换前的obj实现了T接口的所有方法才可以转换成功）
通过接收转换后的接口类型即可使用接收体对这个接口类型的方法进行操作
*/
type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type Bird struct {
}

type Pig struct {
}

func (bird *Bird) Fly() {
	fmt.Println("bird can fly")
}

func (bird *Bird) Walk() {
	fmt.Println("bird can walk")
}

func (pig *Pig) Walk() {
	fmt.Println("pig can walk")
}

func interDemo() {
	/*
		animals :=map[string]interface{}{
			"bird":new(Bird),
			"pig":new(Pig),
		}
	*/
	//	for name,obj:=range animals{
	//
	//	}
	var obj interface{} = new(Pig)
	fly, isFlyer := obj.(Flyer)
	if isFlyer {
		fly.Fly()
	}
	walk, isWalker := obj.(Walker)
	if isWalker {
		walk.Walk()
	}
	walkA, isWalkerA := obj.(Walker)
	if isWalkerA {
		walkA.Walk()
	}
	var objB interface{} = new(Bird)
	flyB, isFlyerB := objB.(Flyer)
	if isFlyerB {
		flyB.Fly()
	}
	walkB, isWalkerB := objB.(Walker)
	if isWalkerB {
		walkB.Walk()
	}
}

func main() {
	interDemo01()
	/*
		f := catCreate//将函数当做值传递给一个变量（相当与给函数起了一个别名）调用时f()和catCreate()都可以调用该方法
		f("sss")
	*/
	/*
		cat := catCreate("wihte")
		fmt.Printf("true address is %p\n", &cat)
		fmt.Println("cat's color is", cat.color)
	*/
	//test01()
	//test02()
	//test03()
	//test04()
	//arrByValue()
	//sliceTest04()
	//varDemo()
	//sliceTest05()
	//mapTest01()
	//mapTest02()
	//mapTest03()
	//ptrTest01()
	/*
		demo := 1
		fmt.Println(demo)
		p := &demo
		fmt.Println(*p)
		ptrTest02(p)
		fmt.Println(*p)
	*/
	//structSpawn02()
	/*
		ptr := &Demo01{"ks", "mywife", 18}
		fmt.Println("ptr name: ", (*ptr).name01)
		fmt.Println("ptr name: ", ptr.name01) //已自动解引用来获得结构体的name01属性
	*/
	//DemoTest04()
	//DemoTest05()
	//DemoTest06()
}
