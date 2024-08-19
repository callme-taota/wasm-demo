package main

import (
	"embed"
	"wasm-example/utils"
)

//go:embed spx.a
var spxA embed.FS

//go:embed Library/spxSource
var spxS embed.FS

//func main() {
//	c := make(chan struct{})
//	//js function register
//	js.Global().Set("startSPXTypesAnalyser", js.FuncOf(StartSPXTypesAnalyserJS))
//	fmt.Println("WASM Go Initialized")
//	<-c
//}
//
//func StartSPXTypesAnalyserJS(this js.Value, p []js.Value) interface{} {
//	fileName := p[0].String()
//	fileCode := p[1].String()
//	fmt.Println("fileName: ", fileName)
//	fmt.Println("fileCode: ", fileCode)
//	return utils.StartSPXTypesAnalyser(fileName, fileCode)
//}

func main() {
	utils.StartSPXTypesAnalyser("test.spx", `onStart => {
	flag := true
	for flag {
		onMsg "die", => {
			flag = false
		}
		glide -877, 180, 3
		setXYpos -240, 180
	}}`, spxA, spxS)
	//	utils.StartSPXIGOP("test.spx", `onStart => {
	//flag := true
	//for flag {
	//	onMsg "die", => {
	//		flag = false
	//	}
	//	glide -877, 180, 3
	//	setXYpos -240, 180
	//}}`)
}
