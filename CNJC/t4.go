package main

import "fmt"

/*
*
iota 枚举常量
*/
func main() {
	// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
	// const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	const (
		a2 = iota
		b2
		c2
	)
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	const (
		i2 = 1 << iota
		j2 = 3 << iota
		k2 = 2
		l2
	)
	// iota会在每个const代码块中从0开始计数
	// iota下面的每一行都会使iota计数加1
	// iota下面空白常量的会继承最上面的表达式
	fmt.Println("i2=", i2)
	fmt.Println("j2=", j2)
	fmt.Println("k2=", k2)
	fmt.Println("l2=", l2)

	const (
		i3 = 1 + iota
		j3 = 3 + iota
		k3
		l3
	)
	fmt.Println("i3=", i3)
	fmt.Println("j3=", j3)
	fmt.Println("k3=", k3)
	fmt.Println("l3=", l3)
}
