package test

type Struct struct {
	String string
	Int    int32
	Float  float64
	Struct struct {
		X string
	}
	Slice []string
	Map   map[string]string
}

type typeForTest map[string]*Struct
