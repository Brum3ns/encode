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
