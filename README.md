# dcode

_Note: this library is a work in progress. I'd love for you to try it! If you do find issues, please [file an issue](https://github.com/go-functional/dcode/issues/new). Happy hacking!_

This repository lets you decode JSON objects using a functional API. It aims to help you solve these two cases in your everyday programming life:

1. When you don't know the "shape" of the JSON ahead of time
1. When you know there's a huge JSON object coming in, but you only need to get a few values out of it

To accomplish the first case without dcode, you'd usually decode the JSON into a `map[string]interface{}`, and then deal with the untyped `interface{}` values as you traverse your way down into the map to get the value you want.

To accomplish the second case without dcode, you'd usually write a `struct` so that you can decode the values you want, and that's boilerplate code. If the value you want is deeply nested inside the object, this leads to a lot of boilerplate as you build nested structs out.

Interesting? Cool! Check out [USAGE.md](/USAGE.md) for details on using this library.

## Notes

This library follows patterns from the [Elm language JSON decoding library](https://guide.elm-lang.org/effects/json.html).

