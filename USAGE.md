# How to Use dcode

dcode is all about decoding json where the top level is an _object_. That means don't use it to decode things like this:

```json
1
```

Or this:

```json
[1, 2, 3, "a", "b", "c"]
```

Only use it to decode things like this:

```json
{
    "a": {
        "b": {
            "one": 1,
            "two": 2
        }
    }
}
```

I'll show some usage of dcode below, starting with basics and then moving to more advanced.

### Basic - Decoding An Integer Into A `int`

Let's say you have this JSON:

```json
{
    "one": 1
}
```

And you want to decode the value of `"one"` (an `int`) into a Go variable.

Here's how you'd do it with dcode:

```go
decoder := Field("one", Int())
var i int
// this returns an error, so check it in real life!
Decode(jsonBytes, decoder, &i)
```

### Intermediate - Drilling Down

Let's move to a more advanced example. Say you have this JSON:

```json
{
    "a": {
        "b": {
            "one": 1,
            "two": 2
        }
    },
    "b": {
        "100": "oneHundred"
    }
}
```

And you want to get the value of `a.b.one` (that is, the value of the `"one"` key under the `"b"` object, which is under the `"a"` object).

Here's how you'd do it in dcode:

```go
decoder := Field("a", Field("b", Field("c", Int())))
var i int
Decode(decoder, jsonBytes, &i)
```

See those nested calls to `Field`? We can make that simpler with an API that looks like the builder pattern (but is completely functional!):

```go
decoder := First("a").Then("b").Then("c").Into(Int())
var i int
Decode(decoder, jsonBytes, &i)
```

#### Check Out That `Int()`!

IN this and the basic example, we passed `Int()` into our `Field` and `Into` functions. `Int()` returns a decoder that knows how to decode some JSON into a Go `int`. dcode provides lots of these functions for various different primitive types. 

You can compose these types together, but don't go crazy. If you're trying to build up lots of values, you're basically gonna start duplicating `json.Unmarshal`, so just build a `struct` with JSON annotations and use the tried-and-true [`encoding/json`](https://godoc.org/encoding/json) package.

### Advanced - Passing Decoders Around

Decoders are just `func`s, and they don't have any side effects. In other words, if you have a decoder, it'll decode the same JSON the same way, every single time. This is called a _pure function_ in FP lingo.

Sounds obvious, doesn't it? It's really nice if you've got a big codebase and you want to share decoding logic. Instead of creating a new `struct` just to fetch some data out of some JSON, and relying on other people to call `json.Unmarshal`, you can instead pass a decoder to another function and let them call `Decode` as many times as they like.

Let's say you want to build a JSON-based RPC system, where the payloads look like this:

```json
{
    "method_name": "some_function",
    "arguments": {
        "param1": "val1",
        "param2": "val2"
    }
}
```

And you want to build a function that determines if there are any arguments to the method call. You could do this:

```go
paramDecoder := Field("arguments", Field("param1", String()))
if Exists(paramDecoder, jsonBytes) {
    callWithParams(jsonBytes)
} else {
    callWithoutParams(jsonBytes)
}
```
