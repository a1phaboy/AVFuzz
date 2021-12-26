package main

import "github.com/gookit/color"



func Showbanner(){

}

func ErrorLog(info string){
	color.RGBStyleFromString("255,0,0").Println(info)
}

func InfoLog(info string){
	color.RGBStyleFromString("0,245,255").Println(info)
}

func SuccessLog(info string){
	color.RGBStyleFromString("0,255,0").Println(info)
}
