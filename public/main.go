package main

import (
	"fmt"
	"syscall/js"
)

var c chan bool

func init() {
	fmt.Println("Hello, WebAssembly!")
	c := make(chan int)
	
	js.Global().Set("convertPixels", js.FuncOf(convertPixels))
	js.Global().Set("adjustBrightness", js.FuncOf(adjustBrightness))
	println("Done.. done.. done...")

	// tells the channel we created in init() to "stop".
	<-c
}

func main() {

}


func convertPixels(this js.Value, input []js.Value) interface{} {

	var retArray = make([]uint8,input[1].Int())
	for i := 0; i < input[1].Int(); i += 4 {

		lightness := uint8((input[0].Index(i).Int()+ input[0].Index(i+1).Int() + input[0].Index(i+2).Int())/3)
		
		retArray[i], retArray[i+1], retArray[i+2], retArray[i+3] = lightness, lightness, lightness, byte(input[0].Index(i+3).Int())
		

	}
	arrayConstructor := js.Global().Get("Uint8ClampedArray")
	dataJS := arrayConstructor.New(input[1].Int())
	js.CopyBytesToJS(dataJS, retArray)
	return dataJS
}

func adjustBrightness(this js.Value, input []js.Value) interface{}{
	var avg int =0
	for i := 0; i < input[1].Int(); i += 4 {
		avg+= (input[0].Index(i).Int() + input[0].Index(i).Int() + input[0].Index(i).Int())/3
	}
	avg*=4
	avg/=input[1].Int()
	var retArray = make([]uint8,input[1].Int())
	var adj uint8 = 0
	if avg<100 {
		adj = uint8(100 - avg)
	}
	for i := 0; i < input[1].Int(); i += 4 {
		retArray[i] = byte(input[0].Index(i).Int())
		if retArray[i]+adj<255{
			retArray[i] = retArray[i] + adj
		}
		retArray[i+1] = byte(input[0].Index(i+1).Int())
		if retArray[i+1]+adj<255{
			retArray[i+1] = retArray[i+1] + adj
		}
		retArray[i+2] = byte(input[0].Index(i+2).Int())
		if retArray[i+2]+adj<255{
			retArray[i+2] = retArray[i+2] + adj
		}
			
		retArray[i+3] = byte(input[0].Index(i+3).Int())
		
	}
	arrayConstructor := js.Global().Get("Uint8ClampedArray")
	dataJS := arrayConstructor.New(input[1].Int())
	js.CopyBytesToJS(dataJS, retArray)
	return dataJS

}





