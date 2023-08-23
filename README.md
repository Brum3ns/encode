Simple script that takes input from stdin and encodes each given input. The script can encode the same input multiple times and use worker pools with threads to handle large inputs without problems.

The script is particularly useful for modifying general wordlists or payloads.

## Install 
```bash
go install -v github.com/Brum3ns/encode/cmd/encode@latest
```

## Options
| Option | Description |
|-|-|
| -d | The decoder to be used for each input given by stdin |
| -e | The encoder to be used for each input given by stdin |
| -t | Threads to use (Default: 42) |

| Methods | Supports |
|-|-|
|  Base32 | encoder/decoder |
|  Base64 | encoder/decoder |
|  Binary | encoder/decoder |
|  Hex | encoder/decoder |
|  Hexdec | encoder/decoder |
|  Html | encoder/decoder |
|  Htmle | encoder/decoder |
|  Lower | encoder |
|  Octal | encoder/decoder |
|  Unicode | encoder/decoder |
|  Unicodeplus | encoder/decoder |
|  Upper | encoder |
|  Url | encoder/decoder |
|  Urldouble | encoder/decoder |
|  Xhex | encoder/decoder |
|  Xhexdec | encoder/decoder |


## Example Of Usage
```bash
cat wordlist.txt | encode -e url
```

## Adding a New Encoder
> !Note : **The function MUST start with an upper-case and continue with lower-case letters/digits**
> You can read more about this in the code itself.

You add a custom encoder/decoder within this file:
```bash
pkg/encoder/custom.go
```

### Encoder/Decoder Template

> Encoder
```golang
func (e *Encoders) Name(s string) string {
	var str string
	//Your code...
    //...
	return str
}
```
> Decoder
```golang
func (d *Decoders) Name(s string) string {
	var str string
	//Your code...
    //...
	return str
}
```
