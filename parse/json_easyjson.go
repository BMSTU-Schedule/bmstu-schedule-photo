// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package parse

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson42239ddeDecodeBmstuSchedulePhotoParse(in *jlexer.Lexer, out *Groups) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Groups, 0, 8)
			} else {
				*out = Groups{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *Group
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(Group)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson42239ddeEncodeBmstuSchedulePhotoParse(out *jwriter.Writer, in Groups) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Groups) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson42239ddeEncodeBmstuSchedulePhotoParse(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Groups) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson42239ddeEncodeBmstuSchedulePhotoParse(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Groups) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson42239ddeDecodeBmstuSchedulePhotoParse(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Groups) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson42239ddeDecodeBmstuSchedulePhotoParse(l, v)
}
func easyjson42239ddeDecodeBmstuSchedulePhotoParse1(in *jlexer.Lexer, out *Group) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
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
		case "url":
			out.URL = string(in.String())
		case "group":
			out.GroupName = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson42239ddeEncodeBmstuSchedulePhotoParse1(out *jwriter.Writer, in Group) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URL))
	}
	{
		const prefix string = ",\"group\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.GroupName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Group) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson42239ddeEncodeBmstuSchedulePhotoParse1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Group) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson42239ddeEncodeBmstuSchedulePhotoParse1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Group) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson42239ddeDecodeBmstuSchedulePhotoParse1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Group) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson42239ddeDecodeBmstuSchedulePhotoParse1(l, v)
}
