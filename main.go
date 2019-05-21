package main

import "github.com/unrolled/render"

var renderer *render.Render

func init() {
	//렌더러 생성
	renderer = render.New()
}

func main() {
	router := httprouter.New()

}
