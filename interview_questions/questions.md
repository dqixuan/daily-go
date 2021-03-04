# 常见面试题

## 拷贝大切片一定比小切片代价大吗？
   否，拷贝所有切片的代价是相同的。Slice结构如下```
    type slice struct {
    	array unsafe.Pointer // 指向底层数组的指针
    	len   int            // 长度   
    	cap   int            // 容量
    }
    ```
   拷贝切片只是复制这个三个常量， 所以拷贝大切片和小切片的代价是相同的。

## 翻转含有中文、数字、英文字母的字符串ds
   代码如下：
   
   ```
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

