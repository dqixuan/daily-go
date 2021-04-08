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

##C语言字符串和Redis定制的SDS(simple dynamic string)的区别