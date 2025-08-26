# Conversion:

- Type casting, often known as a conversion.

- Type Conversion in Golang is the process of converting a value of one type to a value of another type. Go is a statically-typed language, which means that types are determined at compile-time, rather than at runtime. Therefore, type conversions must be explicit and cannot be done implicitly.

- Go has several built-in functions for type conversion, such as float64(x), int(x), string(x), etc. These functions take the value of x as an argument and return a value of the specified type.

- General Syntax of Type Conversion in Golang:

```go
v := typeName(otherTypeValue)
e.g. i := int(32.987) // casting to integer
```

What is the Need for Type Conversion?
The need for type conversion in Go arises when we need to perform operations or assign values of one type to variables of another type. Go is a statically typed language, which means that the type of a variable is determined at compile-time, and it cannot change at runtime. This can make it difficult to perform operations or assign values between different types of variables.

Type conversions are also used to ensure that values being passed to functions or methods match the expected data type.

There are two types of type conversion:

1. Implicit Type Conversion
2. Explicit Type Conversion

Implicit Type Conversion: Golang does not support implicit type conversion. You will encounter an error if you attempt to assign an int variable to a float variable.

Explicit Type Conversion: This Type Conversion in Golang requires you to explicitly specify the target data type. This is done by using the syntax <target_type>(value). For example, to convert an int value to a float64, you would use the syntax float64(x), where x is an int variable.

You can use type conversion to convert between basic types like int, float, string, bool, etc, and you can also convert between structs, arrays, and pointers to other types.

It's worth mentioning that the Go compiler will not allow you to convert between different types that are not compatible. For example, you can't convert a string to an int, we need to use a library.

Usage:
There are several situations where type conversion is commonly used in Go:

- Assigning values: When assigning a value of one data type to a variable of another data type, an explicit type conversion may be required.
- Function calls: When passing a value to a function or method that expects a specific data type, an explicit type conversion may be required.
- Arithmetic operations: When performing mathematical operations on values of different data types, an explicit type conversion may be required. For example, when dividing an int by a float64, the int value must be converted to a float64.
- Interface implementation: When a struct implements an interface, type conversion may be required to convert the struct to the interface type.
- JSON encoding and decoding: When encoding or decoding JSON data, type conversion may be required to convert between JSON data types and Go data types.
- Database interaction: When interacting with a database, type conversion may be required to convert between Go data types and database data types.