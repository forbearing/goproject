package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	stringFunc()
}

func notes() {
	/*
		len()									统计字符串的长度，按字节 len(str)
		r := []rune(str)						字符串遍历，同时处理有中文的问题
		n,err := strconv.Atoi("12")				字符串转整数
		str = strconv.Itoa(12345)				整数转字符串
		var bytes = []byte("hello go")			字符串转 []byte
		str = string([]byte{97,98,99})			[]byte 转字符串
		str=strconv.FormatInt(123,2)// 2->8,16
				10进制转 2，8，16 进制，strconv.FormatInt 返回对应的字符串
		strings.Contains("seafood", "foo")
				查找子串是否在指定的字符串中
		strings.Count("ceheese", "e")
				统计一个字符串中有几个指定的子串
		fmt.Println(strings.EqualFold("abc", "Abc"))
				不区分大小写的字符串比较(==是区分字母大小写的)
		strings.Index("NLT_abc", "abc")
				返回子串在字符串第一次出现的 index 值
		strings.LastIndex("go golang", "go")
				返回子串在字符串最后一次出现的 index
		strings.Replace("go go hello", "go", "go 语言", n)
				将指定的子串替换成另外一个子串，n 可以指定你希望替换几个，如果 n=-1 表示全部替换
		strings.Split("hello,world,ok", ",")
				按照指定的某个字符，为分割标识符，将一个字符串拆分成字符串数组
		strigns.ToLower("Go"), strings.ToUpper("Go")
				将字符串的字母进行大小写转换
		strings.TrimSpace(" tn a long sdlf ")
				将字符串左右两边的空格去掉
		strings.Trim("!hello!", " !")
				将字符串左右两边指定的字符去掉, "!" 和 " "
		strings.TrimLeft("!hello!"," !")
				将字符串左边指定的字符去掉, "!" 和 " "
		strings.TrimRight("!hello!", " !")
				将右边 ! 和 " " 去掉
		strings.HasPrefix("ftp://192.168.10.1", "ftp")
				判断字符串是否指定的字符串开头
		strings.HasSuffix("NLT_abc.jpg", "jpg")
				判断i字符串是否以指定的字符串结束
	*/
}
func stringFunc() {
	// 1. 打印字符串长度
	var str1 = "hello上海go北京"
	fmt.Println(len(str1))

	// 2. 遍历字符串
	str2 := []rune(str1)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i])
	}
	fmt.Println()

	// 3. 字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Println(n)
	}

	// 4. 整数转字符串
	str3 := strconv.Itoa(12345)
	fmt.Printf("str=%v str=%T\n", str3, str3)

	// 5. 字符串转 []byte
	byte1 := []byte("hello go")
	fmt.Printf("bytes=%v\n", byte1)

	// 6. [] byte 转字符串
	str4 := string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str4)

	// 7. 10进制转换2进制、8进制、16进制
	str := strconv.FormatInt(123, 2)
	fmt.Println("123 对应的二进制:", str)
	str = strconv.FormatInt(123, 8)
	fmt.Println("123 对应的八进制:", str)
	str = strconv.FormatInt(123, 16)
	fmt.Println("123 对应的十六进制:", str)

	// 8. 查找子串是否在指定的字符串中
	b := strings.Contains("hello go", "helo")
	fmt.Println(b)

	// 9. 统计一个字符串有几个指定的子串
	num := strings.Count("ceheese", "e")
	fmt.Println(num)

	// 10. 不区分大小写的字符串比较(== 是区分大小写的)
	b = strings.EqualFold("abc", "ABc")
	fmt.Println(b)

	// 11. 返回子串在字符串第一次出现的 index 值
	strings.Index("NLT_abc", "abc")

	// 12. 返回子串在字符串最后一次出现的 index，如果没有返回 -1
	index := strings.LastIndex("go golang", "go")
	fmt.Println(index)

	// 13. 将指定的子串替换成另外一个子串
	str3 = "go go hello go"
	str = strings.Replace(str3, "go", "golang", 1)
	fmt.Println(str)

	// 14. 按照指定的某个字符，为分隔标识，将一个字符串拆分为字符串数组
	str3 = "hello,world,go,lang"
	strArr := strings.Split(str3, ",")
	fmt.Printf("strArr=%v\n", strArr)
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}

	// 15. 大小写转换
	str3 = "Go"
	str = strings.ToLower(str3)
	fmt.Println(str3, str)
	str = strings.ToUpper(str3)
	fmt.Println(str3, str)

	// 16. 将字符串两边的空格去掉
	str3 = " hello golang  "
	str = strings.TrimSpace(str3)
	fmt.Printf("%q\n", str3)
	fmt.Printf("%q\n", str)

	// 17. 将字符串左右两边指定的字符去掉, 同时去掉 " " 和 "!"
	str3 = "! hello! "
	str = strings.Trim(str3, "! ")
	fmt.Printf("%q\n", str3)
	fmt.Printf("%q\n", str)

	// 18. 将左边或右边的指定的字符去掉

	// 19. 判断字符串是否以指定的字符串开头
	b = strings.HasPrefix("ftp://192.168.0.1", "ftp")
	if b {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// 20. 判断字符串是否以指定的字符串结尾
	b = strings.HasSuffix("file.conf", "conf")
	if b {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
