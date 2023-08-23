package encoder

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html"
	"net/url"
	"strings"
)

type Encoders struct {
	Valid []string
}

func NewEncoder() *Encoders {
	return &Encoders{
		Valid: GetMethodNames(Encoders{}),
	}
}

// //////////////// Supported Encoders //////////////// //

func (e *Encoders) Upper(s string) string {
	return strings.ToUpper(s)
}

func (e *Encoders) Lower(s string) string {
	return strings.ToLower(s)
}

func (e *Encoders) Html(s string) string {
	return html.EscapeString(s)
}

func (e *Encoders) Htmle(s string) string {
	return strings.ReplaceAll(html.EscapeString(s), "&#34;", "&quot;")
}

func (e *Encoders) Url(s string) string {
	return url.QueryEscape(s)
}

func (e *Encoders) Urldouble(s string) string {
	return strings.ReplaceAll(url.QueryEscape(s), "%", "%25")
}

func (e *Encoders) Base64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func (e *Encoders) Base32(s string) string {
	return base32.StdEncoding.EncodeToString([]byte(s))
}

func (e *Encoders) Unicode(s string) string {
	var str string
	for _, r := range s {
		str += "\\u" + fmt.Sprintf("%U", r)[2:]
	}
	return str
}

func (e *Encoders) Unicodeplus(s string) string {
	var str string
	for _, r := range s {
		str += "U+" + fmt.Sprintf("%U", r)[2:]
	}
	return str
}

func (e *Encoders) Octal(s string) string {
	var str string
	for _, r := range s {
		str += "\\" + fmt.Sprintf("%o", r)
	}
	return str
}

func (e *Encoders) Hexdec(s string) string {
	s = fmt.Sprintf("%d", []byte(s))
	return s[1 : len(s)-1]
}

func (e *Encoders) Xhexdec(s string) string {
	var str string
	s = fmt.Sprintf("%d", []byte(s))
	for _, i := range strings.Split(s[1:len(s)-1], " ") {
		str += fmt.Sprintf("\\x%v", i)
	}
	return str
}

func (e *Encoders) Hex(s string) string {
	return hex.EncodeToString([]byte(s))
}

func (e *Encoders) Xhex(s string) string {
	var str string
	for _, r := range s {
		str += fmt.Sprintf("\\x%x", string(r))
	}
	return str
}

func (e *Encoders) Binary(s string) string {
	var l []string
	for _, i := range strings.Split(s, "") {
		l = append(l, fmt.Sprintf("%.08b", []byte(i))[1:9])
	}
	return strings.Join(l, " ")
}
