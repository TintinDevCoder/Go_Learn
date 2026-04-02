package main

// goto 语句
// goto 语句可以无条件地转移到函数内的指定标签处执行
// 标签由标识符后跟冒号组成，标签必须在同一函数内定义
// goto 语句可以用来跳出多层循环，或者在某些情况下替代 break 和 continue 语句

func main() {
	// 使用 goto 跳出多层循环
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				goto Exit // 跳转到 Exit 标签处
			}
			println(i, j)
		}
	}
Exit:
	println("Exited from nested loops")

	// 使用 goto 替代 break 和 continue
	for i := 0; i < 5; i++ {
		if i == 2 {
			goto Skip // 跳转到 Skip 标签处，跳过当前迭代
		}
		println(i)
	}
Skip:
	println("Skipped iteration when i == 2")

}
