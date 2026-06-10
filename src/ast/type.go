package ast

type SymbolType struct {
	Name string
}

func (t SymbolType) _type() {}

func (t SymbolType) String() string {
	return t.Name
}

type ArrayType struct {
	Underlying Type
}

func (t ArrayType) _type() {}
