package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var renderer *render.Render

func init() {
	//렌더러 생성
	renderer = render.New()
}

func main() {
	//라우터 생성
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderer.HTML(w, http.StatusOK, "index", map[string]string{"title": "Simple Chat!"})
	}).Methods("GET")

	//negroni 미들웨어 생성
	n := negroni.Classic()

	//negroni에 router를 핸들러로 등록
	n.UseHandler(router)

	//웹서버 실행
	n.Run(":3000")
}
