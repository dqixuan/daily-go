# 常见面试题

## 拷贝大切片一定比小切片代价大吗？
   否，拷贝所有切片的代价是相同的。Slice结构如下
   ```go
    type slice struct {
    	array unsafe.Pointer // 指向底层数组的指针
    	len   int            // 长度   
    	cap   int            // 容量
    }
  ```
  拷贝切片只是复制这个三个常量， 所以拷贝大切片和小切片的代价是相同的。

## 翻转含有中文、数字、英文字母的字符串
   代码如下：
   
   ```go
    func reverse(s string) string { 
        runeSlice := []rune(s)
        for l, r:=0, len(runeSlice)-1; l < r; {
            runeSlice[l], runeSlice[r] = runeSlice[r], runeSlice[l]
            l++
            r--
        }
        return string(runeSlice)
    }
   ```

[注]：一个英文或数字，对应一个字节；汉字对应三个字节，测试代码如下
    
```go
    s1 := "abc"
    s2 := "abc丁"
    fmt.Println(len(s1)) // 3
    fmt.Println(len(s2)) // 6
  ```
    
## map不初始化会怎么样，初始化长度和不初始化长度区别
三种定义map的形式
```
var m map[key]value
m := make(map[key]value)
m := make(map[key]value, 10)
```
第一种，声明map但未初始化
   写入数据，panic：assignment to entry in nil map; 
   读取数据不会报错, m[key] 为value对应的0值，如int:0, string:空字符串，[]string:[]等
第二种，声明并初始化map,
    
## make和new的区别
   1、make 分配并初始化类型所需的内存空间和结构，并返回类型本身；new分配类型所需的内存空间，返回指向内存的指针
   2、相同点：都是go的内置函数，都可用于分配内存空间
   3、make仅支持slice、map、channel三种类型的创建，初始化内置的数据结构
   4、make能传多个参数， new只能传一个类型参数
   5、分配内存在哪？栈或堆？  不一定，对象大小、是否发生内存逃逸？
   
## Golang GC过程
![image](https://github.com/dqixuan/daily-go/blob/main/Image/question_readme/golang_gc.png)
```
   流程图如下：
第一步：扫描根对象。 1.8之前，全局变量和开启写屏障需要STW，G stack只需要停止该协程
第二步：标记。算法：三色标记法。
第三步：重新标记。重新扫描全局变量，以及修改过的协程栈。 需要STW。
第四部：按标记结果清除。
```
三色标记法：





