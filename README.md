# dcode

This repository lets you decode JSON using a functional API. Some of the time, when you deal with JSON, **you don't know the "shape" of it ahead of time**. In most of these cases, you end up decoding the incoming data into a `map[string]interface{}`. And once you decode, you have to start dealing with the map of untyped interfaces. That usually means either checking type assertions as you drill down the JSON objects, or writing a lot of struct boilerplate (examples to come).

## Usage

Instead, using dcode, we can access the value of `"c"` using this syntax:

```go
decoder := Field("a", Field("b", Field("c", Int)))
var i int
decoder.Decode(incomingBytes, &i)
```

Or, if you're nesting really deep into JSON objects, the "builder" syntax (to come) is cleaner:

```go
decoder := First("a").Then("b").Then("c").Into(Int)
var i int
decoder.Decode(incomingBytes, &i)
```

Notice the `Int` in both of those examples? Those are decoders for primitive types. If you want to decode to a complex type, you can compose them together (example to come). But, in that case, you are probably already likely to know the structure of the JSON, and have a `struct` already available to decode into. Read on for advice on when to use the tried-and-true `struct` decoder and when to use dcode.

## Use Cases

As said in the introduction, dcode is great for cases where you either don't know the shape of the JSON you're ingesting, or when you want to just drill down to get one or two values from a deeply nested JSON object. In the former case, you can avoid decoding into a `map[string]interface{}` and in the latter case, you can avoid writing lots of boilerplate struct code.

But there are a few other great uses for dcode:

1. Passing JSON decoding _logic_ to a function, rather than a `struct`
   - A `Decoder` carries the minimal amount of logic to extract JSON data, where a `struct` may have more fields than the function may need. Passing a `Decoder` encourages loose coupling between functions
   - If you're writing multiple different versions of the same struct (i.e. a full user profile, just their email & name, etc...), consider using Decoders
2. Serializing `Decoder`s (more to come)

And finally, as I mentioned above, combining lots of separate decoders together to populate a `struct` is an antipattern. If you're trying to decode JSON into a struct, use [`encoding/json`](https://godoc.org/encoding/json).

## Notes

This library follows patterns from the [Elm language JSON decoding library](https://guide.elm-lang.org/effects/json.html).

