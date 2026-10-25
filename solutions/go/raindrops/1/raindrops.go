package raindrops

import "fmt"

func Convert(number int) string {
    var outputstr string
	if number%3==0 {
        outputstr+= "Pling"
    }
    if number%5==0{
        outputstr+= "Plang"
    }
    if number%7==0{
        outputstr+= "Plong"
    }
    if outputstr==""{
        outputstr=fmt.Sprintf("%d",number)
    }
    return outputstr
}
