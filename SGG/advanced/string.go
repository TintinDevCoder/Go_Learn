package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串常用的系统函数
func main() {
	var str string = "hello world丁" // ascii码占一个字节，汉字占三个字节
	// 长度
	strlen := len(str) // len函数返回字符串的字节数，而不是字符数
	fmt.Println("字符串str长度:", strlen)

	// 字符串遍历,同时处理中文问题
	r := []rune(str) // 转为rune切片，rune类型占4个字节，可以处理中文
	for i := 0; i < len(r); i++ {
		fmt.Printf("str[%d]=%c\n", i, r[i]) // %c表示输出字符
	}

	// 字符串转为整数
	str2 := "123"
	num, err := strconv.Atoi(str2)
	if err == nil {
		fmt.Println("转换成功！结果是：", num)
	} else {
		fmt.Println("转换失败！", err)
	}

	// 整数转为字符串
	num2 := 456
	str3 := strconv.Itoa(num2)
	fmt.Println("整数转为字符串的结果是：", str3)

	// 字符串转为byte
	bytes := []byte(str)
	fmt.Println("字符串转为byte切片的结果是：", bytes)

	// byte转为字符串
	str4 := string(bytes)
	fmt.Println("byte切片转为字符串的结果是：", str4)

	// 10进制转为2、8、16进制字符串
	num3 := 255
	str5 := strconv.FormatInt(int64(num3), 2) // 转为二进制字符串
	fmt.Println("10进制转为2进制字符串的结果是：", str5)
	str6 := strconv.FormatInt(int64(num3), 8) // 转为八进制字符串
	fmt.Println("10进制转为8进制字符串的结果是：", str6)
	str7 := strconv.FormatInt(int64(num3), 16) // 转为十六进制字符串
	fmt.Println("10进制转为16进制字符串的结果是：", str7)

	// 查找子串是否存在在指定字符串中
	a1 := strings.Contains("hello world", "world") // true
	a2 := strings.Contains("hello world", "abc")   // false
	fmt.Println("hello world字符串中是否包含子串world:", a1)
	fmt.Println("hello world字符串中是否包含子串abc:", a2)

	// 统计一个字符串有几个指定的字符串
	count := strings.Count("hello world", "o") // 2
	fmt.Println("hello world字符串中o出现的次数:", count)

	// 不区分大小写的比较
	a3 := strings.EqualFold("Hello", "hello") // true
	a4 := strings.EqualFold("Hello", "world") // false
	fmt.Println("Hello和hello是否相等:", a3)
	fmt.Println("Hello和world是否相等:", a4)

	// 返回字串在字符串第一次出现的index，如果没有返回-1
	index1 := strings.Index("hello world", "o")   // 4
	index2 := strings.Index("hello world", "abc") // -1
	fmt.Println("o在hello world字符串中第一次出现的index:", index1)
	fmt.Println("abc在hello world字符串中第一次出现的index:", index2)

	// 返回字串在字符串最后一次出现的index，如果没有返回-1
	lastIndex1 := strings.LastIndex("hello world", "o")   // 7
	lastIndex2 := strings.LastIndex("hello world", "abc") // -1
	fmt.Println("o在hello world字符串中最后一次出现的index:", lastIndex1)
	fmt.Println("abc在hello world字符串中最后一次出现的index:", lastIndex2)

	// 将指定的子串替换成其他子串
	// 最后一个参数n表示替换的个数，-1表示替换所有
	str8 := strings.Replace("hello world", "o", "O", -1)
	fmt.Println("将hello world字符串中的o替换成O的结果是：", str8)

	// 将字符串按照指定的分隔符切割成字符串切片
	str9 := "hello,world,go"
	split := strings.Split(str9, ",")
	fmt.Println("将字符串hello,world,go按照逗号切割成字符串切片的结果是：", split)

	// 字符串大小写转换
	str10 := "HELLO WORLD"
	fmt.Println("将字符串HELLO WORLD转换为小写的结果是：", strings.ToLower(str10))
	str11 := "hello world"
	fmt.Println("将字符串hello world转换为大写的结果是：", strings.ToUpper(str11))

	// 去掉字符串两端的空格
	str12 := "  hello world  "
	fmt.Printf("去掉字符串两端空格前：'%s'\n", str12)
	fmt.Printf("去掉字符串两端空格后：'%s'\n", strings.TrimSpace(str12))

	// 去掉字符串两边指定的字符
	str13 := "&&hello world&&"
	fmt.Println("去掉两边的&&字符", strings.Trim(str13, "&"))

	// 去掉左边的指定字符
	fmt.Println("去掉左边的&&字符", strings.TrimLeft(str13, "&"))
	// 去掉右边的指定字符
	fmt.Println("去掉右边的&&字符", strings.TrimRight(str13, "&"))

	// 判断字符串是否以指定的子串开头
	fmt.Println("hello world是否以hello开头:", strings.HasPrefix("hello world", "hello"))

	// 判断字符串是否以指定的子串结尾
	fmt.Println("hello world是否以world结尾:", strings.HasSuffix("hello world", "world"))
}
