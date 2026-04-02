package main

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (s *Student) say() string {
	return fmt.Sprintf("student name: %s gender: %s age: %d id: %d score: %f", s.name, s.gender, s.age, s.id, s.score)
}

type Visitor struct {
	name  string
	age   int
	money int
}

func (v *Visitor) String() string {
	return fmt.Sprintf("visitor name: %s age: %d money: %d", v.name, v.age, v.money)
}
func VisitorConstruct(name string, age int) *Visitor {
	money := 20
	if age < 18 {
		money = 0
	}
	return &Visitor{
		name:  name,
		age:   age,
		money: money,
	}
}
func main() {
	s := Student{name: "dd", gender: "male", age: 18, id: 1, score: 100.5}
	say := s.say()
	fmt.Println(say)

	fmt.Printf("请输入姓名：")
	var name string = ""
	fmt.Scanf("%s", &name)
	fmt.Printf("请输入年龄：")
	var age int = 0
	fmt.Scanf("%d", &age)
	if name == "" || age <= 0 || age >= 200 {
		fmt.Println("输入有误")
	} else {
		v := VisitorConstruct(name, age)
		fmt.Println(v)
	}
}
