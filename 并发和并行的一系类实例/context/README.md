# 主要测试4个情况
- withCancel
- withDeadline
- withTimeout
- sync.WaitGroup

其中前三个输入context包，第四个WaitGroup属于sync包，sync包的waitGroup解决了
前三个里面无法获取"子"gorutine结束时间的问题。

每个问题是一个文件夹，每个文件夹有一个main.go文件，每个文件夹都是一种
情况的发生（或者是模拟发生）

- 第一个是主goruntine（goruntine并没有主次之分，不过context树把他们分为了上下子包的这种关系）
可以立即取消子goruntine的情况
- 第二个是等待时间点之后才能执行取消任务的情况
- 第三个是等待时间之后并且时间是时间间隔比如过了5分钟后，这就不是时间点了。