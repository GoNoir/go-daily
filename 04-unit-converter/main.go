//Unit Converter — Two input fields (kilometers to miles, or similar).
//As you type in one, the other updates automatically. 
//Used: Interface, two-way data binding and number formatting.

package main

//import

type Converter interface {
	convert(length float32) float32
}

type Kilometers struct {
	length float32
}

func (k Kilometers) convert(length float32) float32 {
	return length * 0.62
	
}

type Miles struct {
	length float32
}

func (m Miles) convert(length float32) float32 {
	return length * 1.61 
}







//func main() {

