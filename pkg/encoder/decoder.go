package encoder

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type Decoders struct {
	Valid []string
}

func NewDecoder() *Decoders {
	return &Decoders{
		Valid: SetNames(Decoders{}),
	}
}

// //////////////// Supported Decoders //////////////// //

func (d *Decoders) Html(s string) string {
	return html.UnescapeString(s)
}

func (d *Decoders) Htmle(s string) string {
	return html.UnescapeString(strings.ReplaceAll(s, "&quot;", "&#34;"))
}

func (d *Decoders) Url(s string) string {
	str, err := url.QueryUnescape(s)
	if err != nil {
		return s
	}
	return str
}

func (d *Decoders) Urldouble(s string) string {
	str, err := url.QueryUnescape(strings.ReplaceAll(s, "%25", "%"))
	if err != nil {
		return s
	}
	return str
}

func (d *Decoders) Base64(s string) string {
	str, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return s
	}
	return string(str[:])
}

func (d *Decoders) Base32(s string) string {
	str, err := base32.StdEncoding.DecodeString(s)
	if err != nil {
		return s
	}
	return string(str[:])
}

func (d *Decoders) Unicode(s string) string {
	var str string
	json.Unmarshal([]byte(fmt.Sprintf(`"%v"`, s)), &str)
	return str
}

func (d *Decoders) Unicodeplus(s string) string {
	var str string
	for r, v := range map[string]string{"u+": "\\u", "U+": "\\u"} {
		s = strings.ReplaceAll(s, r, v)
	}
	json.Unmarshal([]byte(fmt.Sprintf(`"%v"`, s)), &str)
	return str
}

func (d *Decoders) Octal(s string) string {
	return s
}

func (e *Decoders) Hexdec(s string) string {
	var str string
	for _, i := range strings.Split(s, " ") {
		c, err := strconv.Atoi(i)
		if err != nil {
			str += i
		} else {
			str += string(c)
		}
	}
	return str
}

func (e *Decoders) Xhexdec(s string) string {
	var str string
	for r, v := range map[string]string{"\\x": " ", "\\X": " "} {
		s = strings.ReplaceAll(s, r, v)
	}
	for _, i := range strings.Split(s, " ") {
		c, err := strconv.Atoi(i)
		if err != nil {
			str += i
		} else {
			str += string(c)
		}
	}
	return str
}

func (d *Decoders) Hex(s string) string {
	bytes, _ := hex.DecodeString(s)
	return string(bytes[:])
}

func (d *Decoders) Xhex(s string) string {
	for r, v := range map[string]string{"\\x": "", "\\X": ""} {
		s = strings.ReplaceAll(s, r, v)
	}
	bytes, _ := hex.DecodeString(s)
	return string(bytes[:])
}

func (d *Decoders) Binary(s string) string {
	var str string
	s = strings.ReplaceAll(s, " ", "")
	for len(s) > 0 {
		if len(s) >= 8 {
			v, err := strconv.ParseUint(s[:8], 2, 64)
			if err == nil {
				str += string(v)
			} else {
				log.Fatal("Invalid binary input given")
			}
			s = s[8:]
		} else {
			break
		}
	}
	return str
}
