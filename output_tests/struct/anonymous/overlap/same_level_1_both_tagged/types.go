package test

type Embedded1 struct {
	F1 int32 `json:"F1"`
}

type Embedded2 struct {
	F1 int32 `json:"F1"`
}

type typeForTest struct {
	Embedded1
	Embedded2
}
