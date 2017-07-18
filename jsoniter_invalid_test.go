package jsoniter

import (
	"github.com/stretchr/testify/require"
	"testing"
	"encoding/json"
	"io"
)

func Test_missing_object_end(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Metric string                 `json:"metric"`
		Tags   map[string]interface{} `json:"tags"`
	}
	obj := TestObject{}
	should.NotNil(UnmarshalFromString(`{"metric": "sys.777","tags": {"a":"123"}`, &obj))
}

func Test_missing_array_end(t *testing.T) {
	should := require.New(t)
	should.NotNil(UnmarshalFromString(`[1,2,3`, &[]int{}))
}

func Test_invalid_any(t *testing.T) {
	should := require.New(t)
	any := Get([]byte("[]"))
	should.Equal(Invalid, any.Get(0.3).ValueType())
	// is nil correct ?
	should.Equal(nil, any.Get(0.3).GetInterface())

	any = any.Get(0.3)
	should.Equal(false, any.ToBool())
	should.Equal(int(0), any.ToInt())
	should.Equal(int32(0), any.ToInt32())
	should.Equal(int64(0), any.ToInt64())
	should.Equal(uint(0), any.ToUint())
	should.Equal(uint32(0), any.ToUint32())
	should.Equal(uint64(0), any.ToUint64())
	should.Equal(float32(0), any.ToFloat32())
	should.Equal(float64(0), any.ToFloat64())
	should.Equal("", any.ToString())

	should.Equal(Invalid, any.Get(0.1).Get(1).ValueType())
}

func Test_invalid_struct_input(t *testing.T) {
	should := require.New(t)
	type TestObject struct{}
	input := []byte{54, 141, 30}
	obj := TestObject{}
	should.NotNil(Unmarshal(input, &obj))
}

func Test_invalid_slice_input(t *testing.T) {
	should := require.New(t)
	type TestObject struct{}
	input := []byte{93}
	obj := []string{}
	should.NotNil(Unmarshal(input, &obj))
}

func Test_invalid_array_input(t *testing.T) {
	should := require.New(t)
	type TestObject struct{}
	input := []byte{93}
	obj := [0]string{}
	should.NotNil(Unmarshal(input, &obj))
}

func Test_double_negative(t *testing.T) {
	should := require.New(t)
	var v interface{}
	should.NotNil(json.Unmarshal([]byte(`--2`), &v))
	var vFloat64 float64
	should.NotNil(UnmarshalFromString(`--2`, &vFloat64))
	var vFloat32 float32
	should.NotNil(UnmarshalFromString(`--2`, &vFloat32))
	var vInt int
	should.NotNil(UnmarshalFromString(`--2`, &vInt))
	iter := ParseString(ConfigDefault, `--2`)
	iter.Skip()
	should.NotEqual(io.EOF, iter.Error)
	should.NotNil(iter.Error)
}

func Test_leading_zero(t *testing.T) {
	should := require.New(t)
	var v interface{}
	should.NotNil(json.Unmarshal([]byte(`01`), &v))
	var vFloat64 float64
	should.NotNil(UnmarshalFromString(`01`, &vFloat64))
	var vFloat32 float32
	should.NotNil(UnmarshalFromString(`01`, &vFloat32))
	var vInt int
	should.NotNil(UnmarshalFromString(`01`, &vInt))
	iter := ParseString(ConfigDefault, `01,`)
	iter.Skip()
	should.NotEqual(io.EOF, iter.Error)
	should.NotNil(iter.Error)
}

func Test_empty_as_number(t *testing.T) {
	should := require.New(t)
	iter := ParseString(ConfigDefault, `,`)
	iter.ReadFloat64()
	should.NotEqual(io.EOF, iter.Error)
	should.NotNil(iter.Error)
	iter = ParseString(ConfigDefault, `,`)
	iter.ReadFloat32()
	should.NotEqual(io.EOF, iter.Error)
	should.NotNil(iter.Error)
}