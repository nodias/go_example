package main

import (
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var renderer *render.Render

const (
	//애플리케이션에서 사용할 세션의 키 정보
	sessionKey    = "simple_chat_session"
	sessionSecret = "simple_chat_session_secret"
)

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
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		renderer.HTML(w, http.StatusOK, "login", nil)
	}).Methods("GET")
	router.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		sessions.GetSession(r).Delete(keyCurrentUser)
		http.Redirect(w, r, "/login", http.StatusFound)
	}).Methods("GET")

	//negroni 미들웨어 생성
	n := negroni.Classic()
	store := cookiestore.New([]byte(sessionSecret))
	n.Use(sessions.Sessions(sessionKey, store))

	//negroni에 router를 핸들러로 등록
	n.UseHandler(router)

	//웹서버 실행
	n.Run(":3000")
}
