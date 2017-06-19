package jsoniter

import (
	"encoding/json"
	"github.com/json-iterator/go/require"
	"math/big"
	"testing"
)

func Test_read_map(t *testing.T) {
	should := require.New(t)
	iter := ParseString(ConfigDefault, `{"hello": "world"}`)
	m := map[string]string{"1": "2"}
	iter.ReadVal(&m)
	copy(iter.buf, []byte{0, 0, 0, 0, 0, 0})
	should.Equal(map[string]string{"1": "2", "hello": "world"}, m)
}

func Test_read_map_of_interface(t *testing.T) {
	should := require.New(t)
	iter := ParseString(ConfigDefault, `{"hello": "world"}`)
	m := map[string]interface{}{"1": "2"}
	iter.ReadVal(&m)
	should.Equal(map[string]interface{}{"1": "2", "hello": "world"}, m)
	iter = ParseString(ConfigDefault, `{"hello": "world"}`)
	should.Equal(map[string]interface{}{"hello": "world"}, iter.Read())
}

func Test_map_wrapper_any_get_all(t *testing.T) {
	should := require.New(t)
	any := Wrap(map[string][]int{"Field1": {1, 2}})
	should.Equal(`{"Field1":1}`, any.Get('*', 0).ToString())
}

func Test_write_val_map(t *testing.T) {
	should := require.New(t)
	val := map[string]string{"1": "2"}
	str, err := MarshalToString(val)
	should.Nil(err)
	should.Equal(`{"1":"2"}`, str)
}

func Test_slice_of_map(t *testing.T) {
	should := require.New(t)
	val := []map[string]string{{"1": "2"}}
	str, err := MarshalToString(val)
	should.Nil(err)
	should.Equal(`[{"1":"2"}]`, str)
	val = []map[string]string{}
	should.Nil(UnmarshalFromString(str, &val))
	should.Equal("2", val[0]["1"])
}

func Test_encode_int_key_map(t *testing.T) {
	should := require.New(t)
	val := map[int]string{1: "2"}
	str, err := MarshalToString(val)
	should.Nil(err)
	should.Equal(`{"1":"2"}`, str)
}

func Test_decode_int_key_map(t *testing.T) {
	should := require.New(t)
	var val map[int]string
	should.Nil(UnmarshalFromString(`{"1":"2"}`, &val))
	should.Equal(map[int]string{1: "2"}, val)
}

func Test_encode_TextMarshaler_key_map(t *testing.T) {
	should := require.New(t)
	f, _, _ := big.ParseFloat("1", 10, 64, big.ToZero)
	val := map[*big.Float]string{f: "2"}
	str, err := MarshalToString(val)
	should.Nil(err)
	should.Equal(`{"1":"2"}`, str)
}

func Test_decode_TextMarshaler_key_map(t *testing.T) {
	should := require.New(t)
	var val map[*big.Float]string
	should.Nil(UnmarshalFromString(`{"1":"2"}`, &val))
	str, err := MarshalToString(val)
	should.Nil(err)
	should.Equal(`{"1":"2"}`, str)
}

func Test_map_key_with_escaped_char(t *testing.T) {
	type Ttest struct {
		Map map[string]string
	}
	var jsonBytes = []byte(`
	{
	    "Map":{
		"k\"ey": "val"
	    }
	}`)
	should := require.New(t)
	{
		var obj Ttest
		should.Nil(json.Unmarshal(jsonBytes, &obj))
		should.Equal(map[string]string{"k\"ey": "val"}, obj.Map)
	}
	{
		var obj Ttest
		should.Nil(Unmarshal(jsonBytes, &obj))
		should.Equal(map[string]string{"k\"ey": "val"}, obj.Map)
	}
}

func Test_encode_map_with_sorted_keys(t *testing.T) {
	should := require.New(t)
	m := map[string]interface{}{
		"3": 3,
		"1": 1,
		"2": 2,
	}
	bytes, err := json.Marshal(m)
	should.Nil(err)
	output, err := ConfigCompatibleWithStandardLibrary.MarshalToString(m)
	should.Nil(err)
	should.Equal(string(bytes), output)
}

func Test_decode_map_of_raw_message(t *testing.T) {
	should := require.New(t)
	type RawMap map[string]*json.RawMessage
	b := []byte("{\"test\":[{\"key\":\"value\"}]}")
	var rawMap RawMap
	should.Nil(Unmarshal(b, &rawMap))
	should.Equal(`[{"key":"value"}]`, string(*rawMap["test"]))
	type Inner struct {
		Key string `json:"key"`
	}
	var inner []Inner
	Unmarshal(*rawMap["test"], &inner)
	should.Equal("value", inner[0].Key)
}

func Test_encode_map_of_raw_message(t *testing.T) {
	should := require.New(t)
	type RawMap map[string]*json.RawMessage
	value := json.RawMessage("[]")
	rawMap := RawMap{"hello": &value}
	output, err := MarshalToString(rawMap)
	should.Nil(err)
	should.Equal(`{"hello":[]}`, output)
}

