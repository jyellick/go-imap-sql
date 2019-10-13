// Code generated by easyjson for marshaling/unmarshaling. Patched by hand.

package imapsql

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

func easyjsonUnmarshalCachedHeader(in *jlexer.Lexer, out map[string][]string) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('{')
		if !in.IsDelim('}') {
			out = make(map[string][]string)
		} else {
			out = nil
		}
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v1 []string
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				in.Delim('[')
				if v1 == nil {
					if !in.IsDelim(']') {
						v1 = make([]string, 0, 4)
					} else {
						v1 = []string{}
					}
				} else {
					v1 = (v1)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					v1 = append(v1, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
			out[key] = v1
			in.WantComma()
		}
		in.Delim('}')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonMarshalCachedHeader(out *jwriter.Writer, in map[string][]string) {
	if in == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v3First := true
		for v3Name, v3Value := range in {
			if v3First {
				v3First = false
			} else {
				out.RawByte(',')
			}
			out.String(string(v3Name))
			out.RawByte(':')
			if v3Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
				out.RawString("null")
			} else {
				out.RawByte('[')
				for v4, v5 := range v3Value {
					if v4 > 0 {
						out.RawByte(',')
					}
					out.String(string(v5))
				}
				out.RawByte(']')
			}
		}
		out.RawByte('}')
	}
}
