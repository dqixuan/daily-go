## redis持久方式
有两种：RDB和AOF
```
1、RDB方式 快照方式保存某时间点的数据库
通过SAVE或BGSAVE进行持久化， 
save命令会阻塞Redis服务器，直到命令执行完成
BGSAVE命令，会开启子进程，在子进程中执行BGSAVE命令，不影响服务器执行其他操作(SAVE, BGSAVE, BGREWRITEAOF除外）

type redisServer struct {
    saveparam saveparms; // 包含时间，次数
    t_time lastsave;     // 上次持久化时间
    int dirty;           // 上次持久化之后，修改数据库次数
}

周期函数定时扫描，满足条件，则执行持久化

2、AOF方式  追加命令行的方式持久化

修改数据库的命令保存到aof_buf缓存，以某种策略同步到磁盘
保存策略
always:     aof_buf中有内容就同步磁盘，安全性高，最多丢失1次数据，频繁进行IO操作，性能低
everysec:   每秒钟保存一次，最多丢失1s的数据，性能较always提高，子线程操作，不阻塞。
no:         redis不控制保存的时间点，操作系统控制什么时间保存

everysec：write操作主进程执行，会阻塞；save操作由子线程执行，不阻塞主进程。
always、no: write、save均由主线程执行，会阻塞。

AOF重写功能
不是对AOF文件的重写，而是保存数据库中的现有字段。开启子进程执行，不影响主进程处理其他命令。
AOF重写缓存，保存重写过程中对数据库的修改操作。
重写完成后，子进程通知主进程，主进程调用信号处理函数，将AOF重写缓存的内容保存到新的AOF文件中，并替换原有AOF文件。
重写触发条件：
```

## redis底层字符串结构SDS(simple dynamic string)
SDS结构体如下
```cgo
struct sdshdr {
    // 字符串长度
    int len;
    // 剩余字节长度
    int free;
    // 底层字符数组，用于保存字符串
    char buf[];
}
```
        Redis自定义字符串类型，优化了C语言处理字符串的逻辑。与c语言字符串相比优势：
        1、O(1)时间复杂度获取字符串长度
        2、杜绝字符串溢出，执行相关操作前，会会检查SDS空间是否足够
        3、减少修改字符串时带来的内存重分配次数。
            内存预分配：字符串大小小于1M，分配 2M+1byte; >= 1M，分配 len+1M+1byte；
            惰性空间释放：字符串缩短时，不会立即回收空间，free字段记录，等待将来使用。
        4、c字符串只能保存文本，sds保存二进制文件，可以是图片、音频、视频等
        5、sds兼容部分c语言字符串的操作

## 字典类型
字典是数据结构如下
```cgo
struct dict {
    ...
    // 两个哈希表
    dictht ht[2];
    // rehash索引
    int rehashidx;
    ...
}
struct dictht {
    // 哈希表数组
    dictEntry **table;
    // 哈希表大小
    unsigned long size;
    // 哈希表掩码 size-1
    unsigned long sizemask;
    // 哈希表已有节点数量
    unsigned long used;
}
```
        字典底层主要使用两个hashtable, ht[0]平时使用， ht[1]仅在rehash的时候使用
        负载因子 load factor 控制，哈希表是扩容还是缩容   load factor = ht[0].used / ht[0].size 
        扩容：未执行save/bgsave  factor>=1 执行扩容； 执行save/bgsave  factor >= 5  新容量 ht[0].used *2 大于该值的第一个2次幂
        缩容：factor < 0.1  大于等于 ht[0].used的第一个2次幂
        rehash过程不是一次性、全量进行的，而是分批次，渐进式
        rehash过程：
          1、为ht[1]分配空间，让字典同时持有两个hashtable
          2、将rehashidx置为0
          3、每当程序执行增、删、改、查操作，程序除了执行相应命令外，还会将ht[0]在rehashidx索引的所有键值对rehash到ht[1], rehashidx+1
          4、重复步骤3直到ht[0]的数据为空，释放ht[0],将ht[1]设置为ht[0]，ht[1]设置为空哈希表, rehashidx置为-1
        rehash过程中的增、删、改、查：
            查：先在ht[0]查，再在ht[1]
            增：在ht[1]增加键值对
            删、改: 类似查操作
        

## sorted set  有序集合
