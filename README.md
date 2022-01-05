# Generics

Generics are the most significant change to Go since the release of Go 1.

## Using Generics

**Write code, don't design types.**

先编写数据结构和函数，再指定其中的类型。使用泛型后，函数拥有一种新的参数，称为“类型参数”。对于类型参数，我们会说实例化而不是调用，因为相关操作完全在编译阶段而不是在运行时发生。

类型参数具有限制条件，限制允许的类型实参集。

> 使用反射，编写费劲，运行时慢，也没有静态的类型检查，使用 reflect 包的过程非常复杂，

```go
// This works for maps of type map[string]int.
func MapKeys(m map[string]int) []string{
    var s []string
    for k := range m {
        s = append(s, k)
    }
    return s
}

// This works for maps of any type.
// comparable 是 K 的限制条件，可以认为是该类型参数的元类型。
func MapKeys[K comparable, V any](m map[K]V)[]K{
    var s []K
    for k := range m {
        s = append(s, k)
    }
    return s
}
```

### when to use generics

首先编写函数，稍后，当你清晰地看到可以使用类型参数时再轻松地添加，

1. Functions that work on slices，maps and channels of any element type.
    对该语言中定义的特殊类型进行操作的函数：切片、映射和通道，如果函数具有这些类型的参数，并且函数代码没有对元素类型做出任何特定假设，那么使用类型参数可能会很有用。
2. General purpose data structures.
    2.1 When operating on type parameters, prefer functions to methods.
通用数据结构，类似切片或映射，但是没有内置到语言中，可以使用特定的元素类型进行编写或者使用接口类型，将特定元素类型替换为参数类型，生成更通用的数据结构。
将接口类型替换为类型参数，通常可以更高效地存储数据，在某些情况下，使用类型参数而不是接口类型，代码可以避免了类型断言，而且可以在编译时就进行全面的类型检查。
```go
type tree[T any] struct {
    cmp func(T, T) int
    root *leaf[T]
}

type leaf[T any] struct {
    val T // The value is stored directly in each leaf as a T, not an interface{}.
    left, right *leaf[T]
}
```

3. when a method looks the same for all types.
当不同的类型需要实现一些通用方法，而针对各种类型的实现看起来都相同时，使用类型参数时合理的做法。
```go
// SliceFn implements sort.Interface for any slice type.
type SliceFn[T any] struct{
    s []T
    cmp func（T, T) bool
}

func (s SliceFn[T]) Len() int {return len(s.s)}
func (s SliceFn[T]) Swap(i, j int){ s.s[i],s.s[j] = s.s[j],s.s[i] }
func (s SliceFn[T]) Less(i, j int) bool { return s.cmp(s.[i], s.[j]) }

// SortFn uses SliceFN to sort a slice using a function.
// This is similar to sort.Slice, but the comparison function uses values rather than indexes.
func SortFn[T any](s []T, cmp func(T, T) bool){
    sort.Sort(SliceFn[T]{s, cmp})
}
```

### when not to use generics

1. When just calling a method on the type argument.
Go 具有接口类型，接口类型已经允许某种泛型编程，例如io.Reader接口提供了一种通用机制，用于从包含信息（如文件）或生成信息（如随机数生成器）的任何值中读取数据，对于某个类型的值，如果只需要对该值调用一个方法，请使用接口类型而不是类型参数。
2. When the implementation of a common method is different for each type.
当不同的类型使用一个共同的方法时，考虑该方法的实现，如果一个方法的实现对于所有类型都相同，则使用类型参数，相反，如果每种类型的实现各不相同，使用不同的方法，不要使用类型参数。例如，从文件读取的实现与从随机数生成器读取的实现完全不同，这意味着要编写两种不同的读取方法，两种方法都不应该使用类型参数。
3. When the operation is different for each type, even without a method.
如果某些操作必须支持甚至没有方法的类型，那么接口类型就不起作用，并且如果每种类型的操作都不同，请使用发射，例如JSON编码包，我们不要求我们编码的每个类型都支持marshal JSON方法，因此不能使用接口类型，但是对帧数类型进行编码与对结构体类型进行编码完全不同，因此不应该使用类型参数，所以标准库中使用的是反射，就是实现起来有点复杂。


### guideline

- Avoid boilerplate
如果多次编写完全相同的代码，各个版本之间唯一的区别是代码使用不同的类型，请考虑是否使用类型参数。
    - Corollary：don't use type parameters prematurely；wait until you are about to write boilerplate code.
    在你注意到多次编写完全相同的代码之前，应该避免使用类型参数

It is essential to have good tools， but it is also essential that the tools should be used in the right way.