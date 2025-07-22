package main


type Shape interface{
	Parameter (height int, width int) float32 
}

type Rectangle struct{}

func (Rectangle) Parameter(height int, width int) float32{
	return float32(2*(height + width))
}

type Circle struct{}

func (Circle) Parameter(height int, width int) float32{
	pie := 3.14
	return float32(pie * float64(height) * float64(width))
}





func main(){

}