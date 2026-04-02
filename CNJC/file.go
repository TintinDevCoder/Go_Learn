package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

/*
Go 语言文件处理
在 Go 语言中，文件处理是一个非常重要的功能，它允许我们读取、写入和操作文件。
无论是处理配置文件、日志文件，还是进行数据持久化，文件处理都是不可或缺的一部分。

Go 语言提供了丰富的标准库来支持文件处理，包括文件的打开、关闭、读取、写入、追加和删除等操作。
*/

// os 包
// 文件创建
func file_test1() {
	// 创建文件，如果文件已存在会被截断（清空）
	// 在模块的根目录下创建一个名为 test.txt 的文件
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // 确保文件关闭

	log.Println("文件创建成功")
}

/*
文件的打开与关闭
在 Go 语言中，我们使用 os 包来打开和关闭文件。
os.Open 函数用于打开一个文件，并返回一个 *os.File 类型的文件对象。
打开文件后，我们通常需要调用 Close 方法来关闭文件，以释放系统资源。
*/
func file_test2() {
	// 从模块根目录下打开一个名为 example.txt 的文件
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File opened successfully!")
}

/*
文件的读取
Go 语言提供了多种读取文件的方式，包括逐行读取、一次性读取整个文件等。
我们可以使用 bufio 包来逐行读取文件，或者使用 ioutil 包来一次性读取整个文件。
*/
// 逐行读取文件
func file_test3() {
	// file, err := os.Create("example.txt")
	file, err := os.Open("example.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// 在函数结束离开的那一刻，执行这个操作
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// 当调用 Scan() 时，Scanner 不会只从硬盘读几个字节，而是会“贪婪”地读取一大块数据（默认通常是 64KB）放进 buf
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// 一次性读取整个文件
func file_test4() {
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(content))
}

/*
文件的写入
Go 语言也提供了多种写入文件的方式，包括逐行写入、一次性写入等。
*/
func file_Write() {
	// 我们可以使用 os 包来创建和写入文件。
	// 方式1：直接写入字符串
	file, err := os.Create("write1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString("直接写入字符串\n")

	// 方式2：写入字节切片
	data := []byte("写入字节切片\n")
	file.Write(data)

	// 方式3：使用fmt.Fprintf格式化写入
	fmt.Fprintf(file, "格式化写入: %d\n", 123)

	// 逐行写入文件
	file2, err2 := os.Create("output.txt")
	if err2 != nil {
		fmt.Println("Error creating file:", err2)
		return
	}
	defer file2.Close()
	// bufio.Writer 不会你每写一个字符就跑去写一次硬盘（因为硬盘 IO 极慢）。
	// 它内部有一个缓冲区（默认通常是 4096 字节）。
	// 当你调用 writer.WriteString("Hello") 时，数据其实只是被存在了内存里的这个“仓库”中
	// 如果仓库没满，数据就会一直待在内存里。Flush方法强制将缓冲区中剩余的所有数据，一次性推送到真正的目的地
	// Go 的 bufio 包中： file.Close() 是系统资源的关闭，它不知道上面还有没有 bufio 包装层。 writer 只是一个内存对象，它没有 Close() 方法（因为它不直接持有系统资源）。

	writer := bufio.NewWriter(file2)
	fmt.Fprintln(writer, "Hello, World!1")
	writer.Flush() // 调用 writer.Flush() 确保所有数据都被写入文件
	fmt.Fprintln(writer, "Hello, World!2")
	writer.Flush()

	// 一次性写入文件
	content := []byte("Hello, World!")
	err3 := ioutil.WriteFile("output.txt", content, 0644)
	if err3 != nil {
		fmt.Println("Error writing file:", err3)
		return
	}

	fmt.Println("File written successfully!")
}

// 文件的追加写入
func file_AppendWrite() {
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("Appended text\n"); err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

	fmt.Println("Text appended successfully!")
}

// 文件的删除
// 在 Go 语言中，我们可以使用 os.Remove 函数来删除文件。
func file_Remove() {
	os.Create("output2.txt") // 创建一个文件，方便后续删除
	err := os.Remove("output2.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File deleted successfully!")
}

// 文件信息与操作
func file_Data() {
	// 获取文件信息
	fileInfo, err1 := os.Stat("test.txt")
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("文件名:", fileInfo.Name())
	fmt.Println("文件大小:", fileInfo.Size(), "字节")
	fmt.Println("权限:", fileInfo.Mode())
	fmt.Println("最后修改时间:", fileInfo.ModTime())
	fmt.Println("是目录吗:", fileInfo.IsDir())

	// 检查文件是否存在
	if _, err2 := os.Stat("test.txt"); os.IsNotExist(err2) {
		fmt.Println("文件不存在")
	} else {
		fmt.Println("文件存在")
	}

	// 重命名和移动文件
	err3 := os.Rename("old.txt", "new.txt")
	if err3 != nil {
		log.Fatal(err3)
	}
	log.Println("重命名成功")
}

// 目录操作
func directory_Function() {
	// 创建目录
	// 创建单个目录
	err := os.Mkdir("newdir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// 递归创建多级目录
	err = os.MkdirAll("path/to/newdir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// 读取目录内容
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		info, _ := entry.Info()
		fmt.Printf("%-20s %8d %v\n",
			entry.Name(),
			info.Size(),
			info.ModTime().Format("2006-01-02 15:04:05"))
	}

	// 删除空目录
	err = os.Remove("emptydir")
	if err != nil {
		log.Fatal(err)
	}

	// 递归删除目录及其内容
	err = os.RemoveAll("path/to/dir")
	if err != nil {
		log.Fatal(err)
	}
}

// 高级文件操作
func extention_File_Function() {
	// 文件复制
	srcFile, err := os.Open("source.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create("destination.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	bytesCopied, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("复制完成，共复制 %d 字节", bytesCopied)

}
func main() {
	file_Data()
}
