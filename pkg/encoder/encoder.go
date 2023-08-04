package encoder

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html"
	"net/url"
	"reflect"
	"strings"
)

// /////////////////////////////////////////////////////////////////////////////////// //
// !IMPORTANT
// ----------
//	ALL reciver functions (methods) inside the [struct]ure 'Encoders' are supported
//	encoders. No function should be a reciver unless it's a encoder. All encoder
//	functions *MUST* start with an uppercase and continue with lowercase.
// /////////////////////////////////////////////////////////////////////////////////// //

type Encoders struct {
	Valid []string
}

func New() *Encoders {
	return &Encoders{
		Valid: set_EncoderNames(),
	}
}

func IsValid(l []string) bool {
	encode := reflect.TypeOf(&Encoders{})
	for _, name := range l {
		for i := 0; i < encode.NumMethod(); i++ {
			if name == encode.Method(i).Name {
				break
			}
			if i == encode.NumMethod()-1 {
				return false
			}
		}
	}
	return true
}

func ListEncoders() string {
	var v = "  ---"
	encode := reflect.TypeOf(&Encoders{})
	for i := 0; i < encode.NumMethod(); i++ {
		name := encode.Method(i).Name
		v += fmt.Sprintf("\n  %d. %s", (i + 1), name)
	}
	v += "\n  ---\n"
	return v
}

func set_EncoderNames() []string {
	var lst []string
	encode := reflect.TypeOf(&Encoders{})
	for i := 0; i < encode.NumMethod(); i++ {
		method := encode.Method(i)
		lst = append(lst, method.Name)
	}
	return lst
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

func (e *Encoders) Ascii(s string) string {
	str := fmt.Sprintf("%d", []byte(s))
	return str[1 : len(str)-1]
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

// Add custom encoders below. The *name of the function* will be added to the *encoder list* when you run: "encoder -h" next time.
// !IMPORTANT : See the note at the top of this code (package)
/* func (e *Encoders) Name(s string) string {
	var str string
	//Your code...
	//...
	return str
} */
