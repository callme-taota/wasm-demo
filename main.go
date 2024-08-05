package main

import "wasm-example/utils"

//func main() {
//	c := make(chan struct{})
//	//js function register
//	js.Global().Set("startSPXTypesAnalyser", js.FuncOf(utils.StartSPXTypesAnalyser_JS))
//	//js.Global.Set("startSPXTypesAnalyser", utils.StartSPXTypesAnalyserJS)
//	fmt.Println("WASM Go Initialized")
//	<-c
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
}}`)
}
