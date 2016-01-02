# go-hypixel [![Build Status](https://travis-ci.org/maxikg/go-hypixel.svg)](https://travis-ci.org/maxikg/go-hypixel) [![GoDoc](https://godoc.org/github.com/maxikg/go-hypixel/hypixel?status.svg)](https://godoc.org/github.com/maxikg/go-hypixel/hypixel)

go-hypixel is a Go client library for interacting with the [Hypixel API](https://api.hypixel.net/).

Please note that this is my first go project. My intention is to get familiar with go. Please excuse any insult of the
aesthetics of the go lang. Please let me know what you think about it.

## Usage

First import the library:

```go
import "github.com/maxikg/go-hypixel/hypixel"
```

Construct a new Hypixel client to access the various aspects of their public api. For example, to show an api keys
statistics:

```go
client := hypixel.NewClient("your-api-key", nil)
profile, err := client.KeyInfo()
```

## ToDo

 * Implementing friends
 * Implementing guild
 * Implementing player
 * Implementing session

## License

This library is distributed under the GNU General Public License version 3. A copy is shipped within the
[LICENSE.txt](/LICENSE.txt) file.