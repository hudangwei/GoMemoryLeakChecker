# GoMemoryLeakChecker

## 内存泄漏定位工具

想法来源于此篇文章[实战Go内存泄露](https://segmentfault.com/a/1190000019222661)

## 实现思路
* 对正在运行的程序(go语言编写)，注入shellcode代码(采集pprof数据)
  
  优点是对正在运行的程序直接进行内存泄漏定位，无需修改程序源码，无需重新编译和重启

* 定期采集pprof数据，通过图表方式，更直观显示出差异变化，从而定位出存在疑似问题的代码


## 参考文献
* [Golang specific ELF reader/parser](https://github.com/sitano/goelf)
* [JIT code-generation in Go](https://github.com/nelhage/gojit)
* [JIT compiler in Golang](https://medium.com/kokster/writing-a-jit-compiler-in-golang-964b61295f)
* [golang动态加载原生代码思路](https://www.cnblogs.com/dearplain/p/8145985.html)
* [golang动态补丁实现](https://github.com/bouk/monkey)
