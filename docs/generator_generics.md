# 基于代码生成的泛型

泛型编程主要解决的问题是，*静态类型语言有类型，相关的算法或数据处理的程序会因为类型不同而需要复制一份，这会导致数据和算法功能耦合。*这样的话，在编程的时候就不用关心数据的类型，只要关心处理的逻辑。

*Go语言的代码生成主要是用来解决编程泛型的问题。*

## 类型检查

目前，Go语言不支持泛型，使用`interface{}`这样来实现泛型，在实际使用时需要进行类型检查（Type Assert 或者 Reflection）。

### Type Assert

对某个变量进行`.(type)`的转型操作，返回两个值，`variable`和`error`。

* variable：表示被转换好的类型

* error：表示如果不能转换类型，则报错

## Go generator

手动实现编译时的类型转换。

1. 一个函数模板，在里面设置好相应的占位符
2. 一个脚本，用于按规则来替换文本并生成新的代码（不用手写，有开源的库）:[genny](https://github.com/cheekybits/genny)，[generic](https://github.com/taylorchu/generic)，[gengen](https://github.com/joeshaw/gengen)，[gen](https://github.com/clipperhouse/gen)
3. 一行注释代码

在项目目录下执行`go generate`命令，就能直接生成对应类型的函数。

*具体生成的类型，等参数都是通过注释中的内容来控制的*。

[示例代码](../examples/generator_generics/template/container.tmp.go)。
