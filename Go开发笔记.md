### Delve

- ```
  Goland Debug报错:
  	Version of Delve is too old for this version of Go (maximum supported version 1.12, suppress this error with --check-go-version=false)
  	即Delve版本过旧或者压根没有Deve
  ```

  - 解决措施：下载一个高版本或新的Delve,并在goland中配置其路径

    1. ```
       go get -u github.com/go-delve/delve/cmd/dlv	//下载dlv到本地
       ```

    2. 用everything搜索delve.exe找到其路径，并克隆到Go安装路径的bin目录下

    3. 打开 `Hele->Edit Customer Properties`,若提示文件不存在，点击创建。然后在新加一行 

       ```
       dlv.path=Go的安装路径/bin/dlv.exe
       ```

        重启就可以了

## map使用及知识点

```go
	//定义map,int为key类型，string为值类型
	var s1 map[int]string

	//map为引用类型，如果不进行初始化则为空，即nil，如果不进行初始化就往map中添加元素，则会报错
	fmt.Println(s1 == nil) //true

	//所以map定义之后还需要分配内存空间来进行初始化，类似于java中的new操作，所以使用make来分配内存空间进行初始化
	s1 = make(map[int]string, 10) //第一个参数为map类型，第二个参数为map容量，超出容量会动态扩容

	//初始化之后即可为往map中存储元素
	s1[3] = "李华"
	fmt.Println(s1)

	//根据key获取value
	v1 := s1[3]
	fmt.Println(v1)

	//判断key是否存在并返回value
	v2, ok := s1[4]
	if ok {
		fmt.Println(v2)
	} else {
		fmt.Println("key不存在")
		//string不是引用类型，空值不是nil，而是""这样一个空字符串
		fmt.Printf("value类型为:%T,是否为空:%v,值为:%v\n", v2, v2 == "", v2)
	}

	//map的遍历
	//map中元素无序，无法用普通for循环遍历，需要使用for range遍历
	for k3, v3 := range s1 {
		fmt.Println(k3, v3)
	}
	//只遍历key
	for k4 := range s1 {
		fmt.Println(k4)
	}
	//只遍历value
	for _, v5 := range s1 {
		fmt.Println(v5)
	}

	//删除键值对
	//func delete(m map[Type]Type1, key Type)
	//内置函数delete按照指定的键将元素从映射中删除。若m为nil或无此元素，delete不进行操作。
	delete(s1, 3)
	fmt.Println(s1)

	//测试printOrderly函数
	//初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	//初始化map
	s2 := make(map[string]int, 20)
	//往map中填充元素
	var key string //将key与value放在for循环外避免重复定义
	var value int
	for i := 0; i < 20; i++ {
		key = fmt.Sprintf("第%02d", i) // %02d : 不足两位的整数用0补足，超出或等于两位的整数正常输出
		value = rand.Intn(20)         //生产0~19的整数
		s2[key] = value
	}
	//调用printOrderly函数
	printOrderly(s2)
}

//将map中的元素按序打印
func printOrderly(s2 map[string]int) {
	//初始化一个容量为s2长度的切片
	s3 := make([]string, 0, len(s2))
	//遍历s2，将key存入切片中
	for k := range s2 {
		s3 = append(s3, k)
	}
	//为切片中的key排序
	sort.Strings(s3)
	//输出
	for _, v := range s3 {
		fmt.Println(v, s2[v])
	}

}
```