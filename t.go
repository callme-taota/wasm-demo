// // main.go
package main

//
//import (
//	"fmt"
//	"syscall/js"
//)
//
//// Go function to be called from JavaScript
//func add(this js.Value, p []js.Value) interface{} {
//	return p[0].Int() + p[1].Int()
//}
//
//func tryGoTypeInt(this js.Value, p []js.Value) interface{} {
//	return 1
//}
//
//type JSON struct {
//	i int
//	b bool
//	f float64
//	s string
//	//iList []int
//	//m     map[string]interface{}
//}
//
//func tryGoTypeStruct(this js.Value, p []js.Value) interface{} {
//	//var j JSON
//	//j.i = 1
//	//j.b = true
//	//j.f = 1.00000001
//	//j.s = "hello"
//	//j.iList = []int{1, 2, 3, 4, 5}
//	m := make(map[string]interface{})
//	m["a"] = 1
//	m["b"] = "hello"
//	m["c"] = true
//	//j.m = m
//	return m
//}
//
//func main1() {
//	c := make(chan struct{})
//	js.Global().Set("add", js.FuncOf(add))
//	js.Global().Set("tryGoTypeInt", js.FuncOf(tryGoTypeInt))
//	js.Global().Set("tryGoTypeStru", js.FuncOf(tryGoTypeStruct))
//	fmt.Println("WASM Go Initialized")
//	<-c
//}
//
////type InlayHint = {
////    content: string | Icon,
////    style: InlayHintStyle,
////    behavior: InlayHintBehavior,
////    position: Position
////}
//
//type InlayHint struct {
//	Content  string   `json:"content"`
//	Style    string   `json:"style"`
//	Behavior string   `json:"behavior"`
//	Position Position `json:"position"`
//}
//
//type Position struct {
//	Column     int
//	LineNumber int
//}
