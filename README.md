# go-hypixel [![Build Status](https://travis-ci.org/maxikg/go-hypixel.svg)](https://travis-ci.org/maxikg/go-hypixel) [![GoDoc](https://godoc.org/github.com/maxikg/go-hypixel/hypixel?status.svg)](https://godoc.org/github.com/maxikg/go-hypixel/hypixel)

go-hypixel is a Go client library for interacting with the [Hypixel API](https://api.hypixel.net/).

Please note that this is my first go project. My intention is to get familiar with go. Please excuse any insult of the
aesthetics of the go lang. Please let me know what you think about it.

## Installation

Before you can install the software you need to [install go](https://golang.org/doc/install). The `GOPATH` environment
variable is required too.

The installation of the library (should be) is pretty simple:

```
go get github.com/maxikg/go-hypixel/hypixel
```

Updates can be done using the `-u` option on the `go get`: `go get -u github.com/maxikg/go-hypixel/hypixel`

After the library is installed the import should be resolved.

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

## Supported API features

 * `/key` - Statistic and basic information about the used API key
 * `/findGuild` - Gets a guilds id by a player name, a player uuid or a guild name
 * `/guild` - Gets a guild by its id
 * `/friends` - Get the friend list of a player
 * `/player` - Get various information about a player
 * `/session` - Get the current minigame server of a player including some additional information about it

## ToDo

 * Add tests

## License

This library is distributed under the GNU General Public License version 3. A copy is shipped within the
[LICENSE.txt](/LICENSE.txt) file.