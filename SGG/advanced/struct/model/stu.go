package model

import "fmt"

type stu struct {
	name  string
	score float64
}

func (s *stu) String() string {
	return fmt.Sprintf("student name: %s score: %f", s.name, s.score)
}

// 因为stu结构体首字母小写，因此是只能在model使用
// 但是我们可以通过工厂模式来创建stu对象，并且在其他包中使用
func NewStu(name string, score float64) *stu {
	return &stu{name: name, score: score}
}

// 字段小写，只能在本包访问，通过方法访问
func (s *stu) GetName() string {
	return s.name
}
func (s *stu) SetName(name string) {
	s.name = name
}
func (s *stu) GetScore() float64 {
	return s.score
}
func (s *stu) SetScore(score float64) {
	s.score = score
}
