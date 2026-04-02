package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件的基本操作
// 文件在程序中是以流的方式操作的。
// 流：数据在数据源和程序之间经历的路径  输入流：读文件，文件向程序流动  输出流：写文件，程序向文件流动
// os.File封装了所有文件相关操作，File是一个结构体

// 打开文件
func OpenFile() {
	// os.Open()函数用于打开一个文件，并返回一个*os.File类型的文件对象和一个错误对象。
	// file是文件指针，也称为文件句柄
	file, err := os.Open("D:\\dsy\\Go_Learn\\SGG\\advanced\\file\\example1.txt")

	// 关闭文件
	// defer关键字用于确保在函数结束时执行某些操作，无论函数是正常返回还是发生错误。
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			// 处理关闭文件时的错误，例如记录日志等
			println("关闭文件时发生错误：", closeErr)
		} else {
			println("文件已成功关闭")
			file.Close()
		}
	}()

	if err != nil {
		// 处理错误
		panic(err)
	}
	// 继续进行文件操作，例如读取文件内容等
}

// 读文件，两种方式：带缓冲的读取和一次性读取
// os.Read()函数用于从文件中读取数据，参数是一个字节切片，返回读取的字节数和一个错误对象。
// os.ReadAt()函数用于从文件的指定位置读取数据，参数是一个字节切片和一个偏移量，返回读取的字节数和一个错误对象。
func readFile() {
	// 打开文件
	file, err1 := os.Open("D:\\dsy\\Go_Learn\\SGG\\advanced\\file\\example1.txt")
	if err1 != nil {
		fmt.Println(err1)
	}
	// 关闭文件
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			// 处理关闭文件时的错误，例如记录日志等
			println("关闭文件时发生错误：", closeErr)
		} else {
			println("文件已成功关闭")
			file.Close()
		}
	}()

	// 带缓冲的读取
	// bufio.NewReader()函数用于创建一个带缓冲的读取器，参数是一个io.Reader类型的对象，返回一个*bufio.Reader类型的读取器对象。
	// Reader类型默认缓冲大小4096字节，一次性读取4096字节的数据，如果数据超过缓冲大小，则会分多次读取。
	// bufio.Reader.ReadString()函数用于从读取器中读取数据，参数是一个分隔符，返回读取的字符串和一个错误对象。
	// bufio.Reader.ReadLine()函数用于从读取器中读取一行数据，返回读取的字节切片、是否已到达行尾和一个错误对象。
	reader := bufio.NewReader(file)
	// 循环读取文件内容
	for {
		// 读到一个换行符就结束依次
		str, err := reader.ReadString('\n')
		// 如果读取到文件末尾，err会返回io.EOF错误，表示已经没有更多的数据可供读取了。
		if err == io.EOF {
			if str != "" {
				fmt.Print(str)
			}
			break
		}
		if err != nil {
			// 处理错误，例如记录日志等
			println("读取文件时发生错误：", err)
			break
		}
		fmt.Print(str)
	}
	fmt.Println()
	fmt.Println("------------------------------")

	// ReadFile()函数一次性读取文件的所有内容到内存，不需要文件打开，其将打开操作和读取封装在一起。
	// 接收文件路径作为参数，返回文件内容的字节切片和一个错误对象。
	// 用于一次性读取整个文件的内容，参数是文件路径，返回文件内容的字节切片和一个错误对象。
	// 适合于小文件的读取，如果文件过大，可能会导致内存不足的问题。
	bytes, err2 := os.ReadFile("D:\\dsy\\Go_Learn\\SGG\\advanced\\file\\example1.txt")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(string(bytes))
	}
}

// 写文件操作
func writeFile() {
	// OpenFile()函数用于打开一个文件，并返回一个*os.File类型的文件对象和一个错误对象。
	// 第一个参数是文件路径
	// 第二个参数是文件打开模式，可以是以下几种模式的组合：
	//   os.O_WRONLY：以只写模式打开文件，如果文件不存在则返回错误。
	//   os.O_RDONLY：以只读模式打开文件，如果文件不存在则返回错误。
	//   os.O_RDWR：以读写模式打开文件，如果文件不存在则返回错误。
	//   os.O_CREATE：如果文件不存在则创建文件。
	//   os.O_APPEND：以追加模式打开文件，写入的数据会追加到文件末尾。
	//   os.O_TRUNC：如果文件存在且以只写模式打开，则清空文件。
	// 第三个参数是文件权限，只有在创建文件时才需要指定，常用的权限有：
	//   0644：表示文件所有者具有读写权限，其他用户具有只读权限。
	//   0755：表示文件所有者具有读写执行权限，其他用户具有只读和执行权限。
	//   0777：表示所有用户都具有读写执行权限。
	file1, err1 := os.OpenFile("D:\\dsy\\Go_Learn\\SGG\\advanced\\file\\example2.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err1 != nil {
		fmt.Println(err1)
	}

	// 普通写入方式
	// 我们可以使用 os 包来创建和写入文件。
	// 方式1：直接写入字符串
	file1.WriteString("直接写入字符串\n")

	// 方式2：写入字节切片
	data := []byte("写入字节切片\n")
	file1.Write(data)

	// 方式3：使用fmt.Fprintf格式化写入
	fmt.Fprintf(file1, "格式化写入: %d\n", 123)
	// 一次性写入文件
	content := []byte("Hello, World!")
	err3 := ioutil.WriteFile("output.txt", content, 0644)
	if err3 != nil {
		fmt.Println("Error writing file:", err3)
		return
	} else {
		fmt.Println("File written successfully!")

	}

	// 带缓存的写入方式（逐行写入）
	// bufio.NewWriter()函数用于创建一个带缓冲的写入器，参数是一个io.Writer类型的对象，返回一个*bufio.Writer类型的写入器对象。
	// Writer类型默认缓冲大小4096字节，一次性写入4096字节的数据，如果数据超过缓冲大小，则会分多次写入。
	// writer的WriteString()函数用于向文件中写入字符串，参数是一个字符串，返回写入的字节数和一个错误对象。
	str := "这是一个写入文件的测试字符串。\n"
	writer := bufio.NewWriter(file1)
	for i := 0; i < 5; i++ {
		writer.WriteString(str) // 写入文件
	}

	// 因为writer带缓存，因此实际还未写入磁盘，因此需要调用Flush()函数将缓冲区中的数据写入文件，否则数据会丢失。
	// Flush()函数用于将缓冲区中的数据写入文件，参数是一个错误对象，如果写入成功则返回nil，否则返回一个错误对象。
	err4 := writer.Flush()
	if err4 != nil {
		// 处理错误，例如记录日志等
		println("刷新缓冲区时发生错误：", err4)
	}

	// Write()函数用于向文件中写入数据，参数是一个字节切片，返回写入的字节数和一个错误对象。

	// 关闭文件
	defer func() {
		if closeErr := file1.Close(); closeErr != nil {
			// 处理关闭文件时的错误，例如记录日志等
			println("关闭文件时发生错误：", closeErr)
		} else {
			println("文件已成功关闭")
			file1.Close()
		}
	}()

}

// 文件拷贝
func copyFile() {
	// 1、使用ReadFile函数和WriteFile函数实现文件复制，适合于小文件的复制，如果文件过大，可能会导致内存不足的问题。
	src := "D:\\dsy\\Go_Learn\\SGG\\advanced\\file\\example1.txt"
	dst := "D:\\dsy\\Go_Learn\\SGG\\advanced\\file\\example3.txt"

	// 读取源文件内容
	content, err := os.ReadFile(src)
	if err != nil {
		fmt.Println("读取源文件时发生错误：", err)
		return
	}
	// 将内容写入目标文件
	err = os.WriteFile(dst, content, 0644)
	if err != nil {
		fmt.Println("写入有目标文件时发生错误：", err)
		return
	}

	// 2、使用Copy函数实现文件拷贝，函数接收Writer和Reader作为输入，返回一个和一个error
	srcfile, err := os.Open("D:\\dsy\\Go_Learn\\SGG\\advanced\\file\\example1.txt") // 源目标文件路径
	if err != nil {
		fmt.Println("打开源文件时发生错误：", err)
		return
	}
	dstfile, err := os.OpenFile("D:\\dsy\\Go_Learn\\SGG\\advanced\\file\\example4.txt", os.O_WRONLY|os.O_CREATE, 0666) // 目标文件
	if err != nil {
		fmt.Println("创建目标文件时发生错误：", err)
		return
	}
	// 关闭文件
	defer func() {
		if closeErr := srcfile.Close(); closeErr != nil {
			// 处理关闭文件时的错误，例如记录日志等
			println("关闭源文件时发生错误：", closeErr)
		} else {
			println("源文件已成功关闭")
			srcfile.Close()
		}
		if closeErr := dstfile.Close(); closeErr != nil {
			// 处理关闭文件时的错误，例如记录日志等
			println("关闭目标文件时发生错误：", closeErr)
		} else {
			println("目标文件已成功关闭")
			dstfile.Close()
		}
	}()
	// 获取Reader
	reader := bufio.NewReader(srcfile)
	// 获取到Writer
	writer := bufio.NewWriter(dstfile)
	// 拷贝文件
	io.Copy(writer, reader)
}

// 判断文件或目录是否存在
// os.Stat()方法用于判断文件或文件夹是否存在，如果存在则返回文件或文件夹的信息，如果不存在则返回一个错误。
func isExist() {
	// 判断文件是否存在
	fileData, err := os.Stat("D:\\dsy\\Go_Learn\\SGG\\advanced\\file\\example1.txt")
	if err != nil { // 是否存在错误
		if os.IsNotExist(err) { // 是否是文件或目录不存在错误
			fmt.Println("文件不存在")
		} else {
			fmt.Println("获取文件信息时发生错误：", err)
		}
	}
	fmt.Println("文件存在")
	name := fileData.Name()
	fmt.Println(name)
}
func main() {
	//readFile()
	//writeFile()
	copyFile()
	//isExist()
}
