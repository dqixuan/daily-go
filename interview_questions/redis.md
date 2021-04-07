## redis持久方式
有两种：RDB和AOF
```
1、RDB方式
通过SAVE或BGSAVE进行持久化， 
save命令会阻塞Redis服务器，直到命令执行完成
BGSAVE命令，会开启子进程，在子进程中执行BGSAVE命令，不影响服务器执行其他操作(SAVE, BGSAVE, BGREWRITEAOF除外）


```