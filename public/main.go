package main

import (
	"fmt"
	"syscall/js"
	"reflect"
	"unsafe"
	"math"
)

var c chan bool

func init() {
	fmt.Println("Hello, WebAssembly!")
	c := make(chan int)
	
	js.Global().Set("convertPixels", js.FuncOf(convertPixels))
	js.Global().Set("convertPixels2", js.FuncOf(convertPixels2))
	js.Global().Set("adjustBrightness", js.FuncOf(adjustBrightness))
	js.Global().Set("initializeWasmMemory", js.FuncOf(initializeWasmMemory))
	println("Done.. done.. done...")

	// tells the channel we created in init() to "stop".
	<-c
}

func main() {

}


func convertPixels(this js.Value, input []js.Value) interface{} {

	var retArray = make([]uint8,input[1].Int())
	js.CopyBytesToGo(retArray, input[0])
	for i := 0; i < input[1].Int(); i += 4 {
		red := float64(retArray[i])
		green := float64(retArray[i])
		blue := float64(retArray[i])
		lightness := uint8(math.Floor(0.299*red + 0.587*green + 0.114*blue))
		
		retArray[i], retArray[i+1], retArray[i+2], retArray[i+3] = lightness, lightness, lightness, byte(input[0].Index(i+3).Int())
		

	}
	arrayConstructor := js.Global().Get("Uint8ClampedArray")
	dataJS := arrayConstructor.New(input[1].Int())
	js.CopyBytesToJS(dataJS, retArray)
	return dataJS
}


func convertPixels2(this js.Value, input []js.Value) interface{} {

	var len = input[1].Int()

	sliceHeader := &reflect.SliceHeader{
		Data: uintptr(input[0].Int()),
		Len:  len,
		Cap:  len,
	}

	var ptr = (*[]uint8)(unsafe.Pointer(sliceHeader))
	for i := 0; i < len; i+=4 {
		fmt.Println(uint8((*ptr)[i]))
		lightness:= (uint8((*ptr)[i]) + uint8((*ptr)[i+1]) + uint8((*ptr)[i+2]))
		(*ptr)[i], (*ptr)[i+1], (*ptr)[i+2] = lightness, lightness, lightness
	}
	return 1;

}




func adjustBrightness(this js.Value, input []js.Value) interface{}{
	var imgArray = make([]uint8,input[1].Int())
	js.CopyBytesToGo(imgArray, input[0])
	var avg int =0
	for i := 0; i < input[1].Int(); i += 4 {
		avg+= int((imgArray[i] + imgArray[i+1] + imgArray[i+2])/3)
	}
	avg*=4
	avg/=input[1].Int()
	var retArray = make([]uint8,input[1].Int())
	var adj uint8 = 0
	if avg<64 {
		adj = uint8(64 - avg)
	}
	for i := 0; i < input[1].Int(); i += 4 {
		retArray[i] = byte(imgArray[i])
		if retArray[i]+adj<255{
			retArray[i] = retArray[i] + adj
		}
		retArray[i+1] = byte(imgArray[i+1])
		if retArray[i+1]+adj<255{
			retArray[i+1] = retArray[i+1] + adj
		}
		retArray[i+2] = byte(imgArray[i+2])
		if retArray[i+2]+adj<255{
			retArray[i+2] = retArray[i+2] + adj
		}
			
		retArray[i+3] = byte(imgArray[i+3])
		
	}
	arrayConstructor := js.Global().Get("Uint8ClampedArray")
	dataJS := arrayConstructor.New(input[1].Int())
	js.CopyBytesToJS(dataJS, retArray)
	return dataJS

}

func initializeWasmMemory(this js.Value, args []js.Value) interface{} {

	var ptr *[]uint8
	goArrayLen := args[0].Int()

	goArray := make([]uint8, goArrayLen)
	ptr = &goArray

	boxedPtr := unsafe.Pointer(ptr)
	boxedPtrMap := map[string]interface{}{
		"internalptr": boxedPtr,
	}
	return js.ValueOf(boxedPtrMap)
}





