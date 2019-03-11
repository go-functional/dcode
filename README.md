# Welcome to `dcode` :tada:

[![Go Report Card](https://goreportcard.com/badge/github.com/go-functional/dcode)](https://goreportcard.com/report/github.com/go-functional/dcode)

_Note: this library is a work in progress. I'd love for you to try it! If you do find issues, please [file an issue](https://github.com/go-functional/dcode/issues/new). Happy hacking!_

Welcome Gophers! dcode (pronounced "decode") is a library that gives you a quick and easy way to quickly traverse JSON objects, target a single field, and pull it out into a variable you specify. All using a familiar API.

dcode tries to help you with these two cases in your everyday programming life:

1. When you don't know the "shape" of the JSON coming into your program
2. When you know there's a huge JSON object coming in, but you only need to get a few values out of it

To accomplish the first case without dcode, you'd usually decode the JSON into a `map[string]interface{}`, and then deal with the untyped `interface{}` values as you traverse your way down into the map to get the value you want.

To accomplish the second case without dcode, you'd usually write a bunch of `struct`s so you can decode the values you want (one per "level" in the JSON).

## Ok, Tell Me More!

We're going to focus on the "let's decode big JSON objects" use case in this section.

Let's say you're a weather buff, and you want to write some Go to get the forecast from [Dark Sky](https://darksky.net) using their [API](https://darksky.net/dev/docs#forecast-request). Awesome!

Here's part of a JSON response from Dark Sky's API (adapted from their docs page):

```json
{
    "latitude": 42.3601,
    "longitude": -71.0589,
    "timezone": "America/New_York",
    "minutely": {
        "summary": "Light rain stopping in 13 min., starting again 30 min. later.",
        "icon": "rain",
        "data": {
            "time": 1509993240,
            "precipIntensity": 0.007,
            "precipIntensityError": 0.004,
            "precipProbability": 0.84,
            "precipType": "rain"
        }
    }
}
```

Let's try to get the precipitation probability and type so we can print out "84% chance of rain."

### Let's Decode This With [`encoding/json`](https://godoc.org/encoding/json)

The most common way you decode JSON using `encoding/json` is to decode a `[]byte` into a `struct`. Here's what the `struct` for this JSON response would look like:

```go
type minutelyResp struct {
    precipType string `json:"precipType"`
    precipProbability float64 `json:"precipProbability"`
}

type forecast struct {
    minutely minutelyResp `json:"minutely"`
}
```

Then, here's how we'd decode the struct:

```go
fcast := new(forecast)
// TODO: deal with the error!
json.Unmarshal(jsonBytes, fcast)
```

Not terrible, but it's a little boilerplate-ey just to get a `string` and a `float64` out of this response. If there were values nested even deeper, we'd have to write more structs to get them.

>Note: If you have to start grabbing a ton more data from the API response JSON, it might make sense to fill out the struct more and grab the data you need.

### Let's Decode The Same JSON With `dcode`

We can decode the same JSON with a little bit less boilerplate, and be clearer about the JSON we're trying to get!

To start, here's how we'd use dcode to do the same decoding:

```go
typeDecoder := First("minutely").Then("precipType").Into(String())
probDecoder := First("minutely").Then("precipProbability").Into(Float64())
var precipType string
var prob float64
// TODO: deal with the errors here!
Decode(typeDecoder, jsonBytes, &precipType) 
Decode(probDecoder, jsonBytes, &prob)
```

It looks a fair bit different than the previous example with `encoding/json`. How is this different? Why is it better? Here are some answers for you?

- No `struct`s to write. Hooray!
- You target _just_ the fields you want. No more mistakes with struct tags :fire:
- You write _how_ to traverse the returned JSON rather than the complete structure you expect. Don't fail because some other field came in slightly differently than you expeced
- Your parsing code tends to be more self-documenting. It almost looks like [JSONPath](http://jsonpath.com/)!
- You can reuse those `*Decoder` values as many times as you want, against any `[]byte`, without creating any new `struct`s or allocating any new memory

### Interesting, Right?

Check out [USAGE.md](/USAGE.md) for more details, and enjoy!

## Notes

This library follows patterns from the [Elm language JSON decoding library](https://guide.elm-lang.org/effects/json.html).

I love that language, and I think their JSON decoder is really well designed. The described the benefits to their decoder pattern in their documenation, and I slighly adapted what they wrote here :wink::

>Fun Fact: I have heard a bunch of stories of folks finding bugs in their server code as they switched from ~~JS~~ `encoding/json` to ~~Elm~~ `dcode`. The decoders people write end up working as a validation phase, catching weird stuff in JSON values. So when ~~NoRedInk~~ you switched from ~~React~~ `encoding/json` to ~~Elm~~ `dcode`, it revealed a couple bugs in their ~~Ruby~~ server code!


