# Finding the type of an object

There is the `typeof` in C; there is `typeof` in JavaScript; and there is `type()` in Python to see what type has an object.

In Go, there is not `typeof`, nor `type` in the basic syntax, but there are three ways to get an object's type:

## 1. String formatting

The simplest way.

```go
n := 3.14
fmt.Printf("%T", n)
// float64
```

## 2. Type assertion

We have to prepare it a little bit:

1. Use .(type) inside type switch.
2. Type switch can be done on an interface value only.

```go
var n interface{}
n = 3.14
switch n.(type) {
    case int:
        fmt.Println("int")
    case float32:
        fmt.Println("float32")
    case float64:
        fmt.Println("float64")
    default:
        fmt.Println("unknown")
}
// float64
```

## 3. Reflect package's TypeOf()

An import is needed from the standard library. Here is that we missed in the basic syntax. :)

```go
n := 3.14
fmt.Println(reflect.TypeOf(n).String())
// float64
```

[Based on Stacoverflow comment of Grzegorz Luczywo](https://stackoverflow.com/questions/20170275/how-to-find-the-type-of-an-object-in-go/27160765#27160765).
