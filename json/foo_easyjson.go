package json

import (
	json "encoding/json"

	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

var _ = json.RawMessage{} // suppress unused package warning

func easyjson_decode_github_com_jcbwlkr_studies_json_Foo(in *jlexer.Lexer, out *Foo) {
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Bar":
			out.Bar = in.String()
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjson_encode_github_com_jcbwlkr_studies_json_Foo(out *jwriter.Writer, in *Foo) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Bar\":")
	out.String(in.Bar)
	out.RawByte('}')
}
func (v *Foo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson_encode_github_com_jcbwlkr_studies_json_Foo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}
func (v *Foo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson_encode_github_com_jcbwlkr_studies_json_Foo(w, v)
}
func (v *Foo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson_decode_github_com_jcbwlkr_studies_json_Foo(&r, v)
	return r.Error()
}
func (v *Foo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson_decode_github_com_jcbwlkr_studies_json_Foo(l, v)
}
