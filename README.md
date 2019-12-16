# GoMemoryLeakChecker

## 内存泄漏定位工具

想法来源于此篇文章[实战Go内存泄露](https://segmentfault.com/a/1190000019222661)

## 实现思路
* 对正在运行的程序(go语言编写)，在linux环境下使用ptrace等系列调试函数，注入执行pprof.Lookup("heap")(采集pprof数据)
  
  优点是对正在运行的程序直接进行内存泄漏定位，无需修改程序源码，无需重新编译和重启

* 定期采集pprof数据，通过图表方式，更直观显示出差异变化，从而定位出存在疑似问题的代码

## 参考文献
* [实现一个 Golang 调试器（第一部分）](https://studygolang.com/articles/12553)
* [实现一个 Golang 调试器（第二部分）](https://studygolang.com/articles/12794)
* [实现一个 Golang 调试器（第三部分）](https://studygolang.com/articles/12804)
