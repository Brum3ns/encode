# Encoder

Simple script that takes input from stdin and encodes each given input. The script can encode the same input multiple times and use worker pools with threads to handle large inputs without problems.

The script is particularly useful for modifying general wordlists or payloads.

## Install 
```bash
go install -v github.com/Brum3ns/encode/cmd/encode@latest
```

## Example Of Usage
```bash
cat wordlist.txt | encode -e url
```

## Adding a New Encoder
> !Note : **The function MUST start with an uppercase and continue with lowercase**
> You can read more about this in the code itself.


You find all the function and where you can add your own in the following file:
```bash
pkg/encoder/encoder.go
```

### Template function code:
```golang
func (e *Encoders) Name(s string) string {
	var str string
	//Your code...
    //...
	return str
}
```

## Supported Encoders
- Ascii 
- Base32 
- Base64 
- Binary 
- Hex 
- Html 
- Htmle 
- Lower 
- Octal 
- Unicode 
- Unicodeplus 
- Upper 
- Url 
- Urldouble 
- Xhex 
