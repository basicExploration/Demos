## 也算是回顾也算是重新学习shell
- 获取进程的id 利用名字 pgrep -f name， 如果是可执行文件的话 pidof name 这样 也可以额  ps -ef |grep go|grep -v grep|awk '{print$2}'
 其中 awk是一种在Linux中运用的编程语言 这样就输出了第二个选项 就是进程id  grep -v grep 是将grep自己的干扰项去除。
