# Generics

- [官方提案](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)
- [官方博客](https://go.dev/blog/intro-generics)

Generics are the most significant change to Go since the release of Go 1.

Go 是一门强类型语言，意味着程序中的每个变量和值都有某种特定的类型。现在有泛型了，不代表你一定要用。在没有泛型的时候，我们有三个选择：

1. [代码生成](./docs/generator_generics.md)
2. 反射
3. 冗余的代码

**Write code, don't design types.**

泛型的意思就是，先编写好数据结构和函数，再指定其中的类型。使用泛型后，函数拥有一种新的参数，称为“类型参数”。对于类型参数，我们会说实例化而不是调用，因为相关操作完全在编译阶段而不是在运行时发生。类型参数具有限制条件，限制允许的类型实参集。

```go
func name[TypeLabel Constraints](...) {
    ...
}
```

## 泛型约束

[`constraints`](golang.org/x/exp/constraints) 包中定义了一组用于类型参数（泛型）的约束条件。实际上 Go 会进行类型推断，即编译器会通过普通参数的类型推导出类型参数。不过，跟 Go 中其他类型自动推导类似，有些情况是无法自动推导的，这时候必须手动指定实际的类型参数。

### any

Go 语言自身实现，表示没有任何约束，`any` 是 `interface{}` 的别名，但是，注意与 `interface{}` 这样的任意类型区分开, 泛型中的类型，在函数内部并不需要任何**类型断言**和**反射**的工作，在编译期就可以确定具体的类型。

```go
type any = interface{}
```

引入 `any` 关键字，让泛型修饰符更短更清爽，在项目中使用 `gofmt` 进行批量修改。

```shell
all: gofmt -w -r 'interface{} -> any' ./...
```

### [comparable](https://pkg.go.dev/builtin@master#comparable)

Go 语言本身实现，是一个 `interface` 表示可比较（相等与否的比较）。

所有的可比较类型：

- booleans
- numbers
- strings
- pointers
- channels
- interfaces
- arrays of comparable types
- structs whose fields are all comparable types

**只能作为泛型的参数类型，不能作为变量的类型**。

### `~`

如下约束 `~string` 表示的是，支持 `string` 类型以及底层是 `string` 类型的类型。因此，`MyString` 类型也可以作为下面 `add` 函数的类型参数。

```go
func add[T ~string](x, y T) T {}

type MyString string

func demo(){
    a, b := MyString("a"), MyString("b")
    s := add(a, b)
    fmt.Println(s)
}
```

### 接口约束

定义一个 `interface` 其中包含对应的约束，称为类型列表，使用 `|` 分隔，有或的含义。

为了方便，官方提供了一个新的包 [`constraints`](golang.org/x/exp/constraints) ，预定义了一些接口约束。

```go
// Package constraints defines a set of useful constraints to be used
// with type parameters.
package constraints

// 整型相关类型约束
// Signed is a constraint that permits any signed integer type.
type Signed interface {
 ~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
type Unsigned interface {
 ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint that permits any integer type.
type Integer interface {
 Signed | Unsigned
}

// 浮点型约束
// Float is a constraint that permits any floating-point type.
type Float interface {
 ~float32 | ~float64
}

// 负数类型约束
// Complex is a constraint that permits any complex numeric type.
type Complex interface {
 ~complex64 | ~complex128
}

// 支持排序的类型约束（支持大小比较的类型）
// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
type Ordered interface {
 Integer | Float | ~string
}
```

匿名接口约束，类型约束比较少的时候使用。

```go
func add[T int|float64](a,b T) T {}
```

除了可以基于基本数据类型做约束，也可以像常规的接口一样，通过实现方法的形式来约束。

```go
type customConstraint interface {
    ~int
    String() string
}
```

所有实现 `func String() string` 方法并且底层类型是 `int` 的类型都满足上述约束。

## 注意点

1. 类型参数不能用在方法上，只能用在函数上。
2. 如果约束是 channel，对泛型不能用 `make()` 函数。

## 使用泛型

> - 工欲善其事，必先利其器，君子藏器于身，待时而动。
> - It is essential to have good tools， but it is also essential that the tools should be used in the right way.

如果多次编写完全相同的代码，各个版本之间唯一的区别是代码使用不同的类型，请考虑是否使用类型参数，不要过早的引入泛型。

> 使用反射，编写费劲，运行时慢，也没有静态的类型检查，使用 `reflect` 包的过程非常复杂。

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

### 何时应该使用泛型

我们先编写函数，稍后，当你清晰地发现可以使用类型参数时再轻松地添加。

1. 对语言中定义的特殊类型（切片、映射和通道）进行操作的函数，并且函数代码没有对特殊类型中元素的类型做出任何特定假设，那么使用类型参数可能会很有用。
2. 通用数据结构，类似切片或映射，但是没有内置到语言中，我们可以使用特定的元素类型进行编写或者使用接口类型，将特定元素类型替换为参数类型，生成更通用的数据结构。将接口类型替换为类型参数，通常可以更高效地存储数据，在某些情况下，使用类型参数而不是接口类型，代码可以避免了类型断言，而且可以在编译时就进行全面的类型检查。

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

3. 当不同的类型需要实现一些通用方法，而针对各种类型的实现看起来都相同时，使用类型参数是合理的做法。

```go
// SliceFn implements sort.Interface for any slice type.
type SliceFn[T any] struct{
    s []T
    cmp func(T, T) bool
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

### 何时不应该使用泛型

1. Go 具有接口类型，接口类型已经允许某种泛型编程，例如 `io.Reader` 接口提供了一种通用机制，用于从包含信息（如文件）或生成信息（如随机数生成器）的任何值中读取数据，对于某个类型的值，如果只需要对该值调用一个方法，请使用接口类型而不是类型参数。
2. 当不同的类型使用一个共同的方法时，考虑该方法的实现，如果一个方法的实现对于所有类型都相同，则使用类型参数，相反，如果每种类型的实现各不相同，使用不同的方法，不要使用类型参数。例如，从文件读取的实现与从随机数生成器读取的实现完全不同，这意味着要编写两种不同的读取方法，两种方法都不应该使用类型参数。
3. 如果某些操作必须支持甚至没有方法的类型，那么接口类型就不起作用，并且如果每种类型的操作都不同，请使用发射，例如 JSON 编码包，我们不要求我们编码的每个类型都支持 marshal JSON 方法，因此不能使用接口类型，但是对帧数类型进行编码与对结构体类型进行编码完全不同，因此不应该使用类型参数，所以标准库中使用的是反射，就是实现起来有点复杂。
4. 当类型参数使我们的代码更复杂时，不应该考虑泛型，毕竟没有泛型，Go 也已经存在十多年了。如果编写通用函数或结构时发现类型参数不会使我们的代码更清晰，应该重新考虑要不要用泛型。
