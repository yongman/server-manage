#A tool used to alloc machine for redis,redisproxy,memcached,etc

目录说明：
> * command：对不同需求进行资源分配的入口
    redis: redis服务相关操作
    redisproxy: redisproxy相关操作
    memcache: TODO
    print: 打印相关操作

> * modules:
    alloc:资源分配模块
    commit：结果提交模块
    db：数据结构定义和数据库访问模块
    fmtoutput：格式化输出模块
    update：信息更新模块
    log：日志模块

> * utils：公共函数和常量定义


```
server-manage redis  alloc -m 1G -bj 2 -nj 2 -hz 1
```
作用：分别在bj，nj，hz分配指定数目的机器用做redis。

```
server-manage redis list -n 10
```
作用：列出最近的10条提交，无n参数会列出全部的提交结果

```
server-manage redis drop -cid <commitID/all>
```
作用：撤销指定的提交或全部提交，撤销后系统会将盒子数量释放，并添加给机器

```
server-manage redis update -host <hostname> -b1g <k> -b5g <m> -b10g <n>
```
作用：修正指定机器上内存盒子的数量

```
server-manage redisproxy -bj 2 -bj 3 -hz 1
```
作用：在对应的region选择指定数量的机器用作redisproxy

```
server-manage update -s <servicefile> -m <machinefile>
```
作用：用于更新系统中的机器和服务信息

```
server-manage print -t <service/machine> -p <the used percentage(0--100)>
```
作用：打印machine信息或service信息

```
server-manage ban -host <hostname> -s <redis|redisproxy|memcache|...>
```
作用：对hostname机器对某个服务做封禁

```
server-manage unban -host <hostname> -s <redis|redisproxy|memcache|...>
```
作用：对hostname机器对某个服务解封


可以添加其他服务的分配策略
