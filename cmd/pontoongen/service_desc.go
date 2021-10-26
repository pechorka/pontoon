package main

type serviceDesc struct {
	name string
	doc  string

	filename          string
	pkg               string
	serviceStructName string

	handlers []hdlDesc
}

type hdlDesc struct {
	op          string
	path        string
	description string // TODO fill
	inout       hdlTypesDesc
}

type hdlTypesDesc struct {
	inType            *typeDesc
	hasResponseWriter bool
	outType           *typeDesc
}
