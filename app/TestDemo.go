package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func main() {

	p := test5("sun", 19)
	fmt.Println(p)
	p.setAge(102)
	fmt.Println(p.age)
	_sort()

}

func _sort() {

	map1 := make(map[int]string, 5)
	map1[2] = "www.topgoer.com"
	map1[3] = "rpc.topger.com"
	map1[1] = "ceshi"

	index := []int{}
	for k, _ := range map1 {
		index = append(index, k)
	}
	sort.Ints(index)
	for i := 0; i < len(map1); i++ {
		fmt.Println(map1[index[i]])
	}
}

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	// ce = append(ce, student{3, "xiaowang", 56})
	// return ce
}
func test9() {
	var ce []student //定义一个切片类型的结构体
	ce = []student{
		student{1, "xiaoming", 22},
		student{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)
}
func test8() {
	s1 := Student{
		ID:     1,
		Gender: "女",
		name:   "pprof",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"女"}
}

type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func test7() {

	var d dog
	d.Age = 12
	d.say()
	d.pao()

	fmt.Println(d)
	d2 := dog{Age: 14}
	marshal, err := json.Marshal(d2)
	if err != nil {
		fmt.Println("json object failed")
		return
	}

	fmt.Println("-->", marshal)

}

type Animal struct {
	name string
}

func (a *Animal) say() {
	fmt.Println("wanwanwang")
}

type dog struct {
	Age     int8
	*Animal //通过嵌套匿名的结构体实现继承
}

func (d *dog) pao() {
	fmt.Println("paopaopao")
}

type MyInt int

func test6() {
	var m MyInt
	fmt.Println(m)
	m.say()
	m = 100
	fmt.Println(m)
}

func (m MyInt) say() {
	fmt.Println("ays")
}

func (p *persion) setAge(age int8) {
	p.age = age
}

func test5(name string, age int8) *persion {
	return &persion{name: name, age: age}
}

func test4() {

	m := make(map[string]*persion)
	per := []persion{
		{name: "name1", age: 12},
		{name: "name2", age: 13},
	}
	for k, p := range per {
		fmt.Println(k)
		fmt.Println(&p)
		m[p.name] = &p
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

func test3() {

	var p1 persion
	p1.name = "sunjunhong"
	p1.age = 11
	fmt.Println(p1.name)
	fmt.Println(p1.age)
}

type persion struct {
	name string
	city string
	age  int8
}

func test2() {

	var mapSlice = make(map[string][]string, 3)
	fmt.Println(mapSlice)
	key := "中国"
	value, ok := mapSlice[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	mapSlice[key] = value
	fmt.Println(mapSlice)
}

func tests() {
	var mapSlice = make([]map[string]string, 3)
	for k, v := range mapSlice {
		fmt.Println("index:%d value:%v \n", k, v)
	}

	fmt.Println("after print")
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "张三"
	mapSlice[0]["pass"] = "pass"

	for k, v := range mapSlice {
		fmt.Println("index:%d, value:%d", k, v)
		fmt.Println(v["name"])
		fmt.Println(v["pass"])
	}
}
