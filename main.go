package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	)

type middleWareHandler struct {
	r *httprouter.Router
}

// 返回一个接口
func NewMiddleWareHandler(r *httprouter.Router) http.Handler{
	m :=middleWareHandler{}
	m.r = r
	return m
}
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	// 检查session
	ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
}
func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	router.POST("/user",CreateUser)
	router.POST("/user/:user_name",Login)

	return router
}


func main(){
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	_ = http.ListenAndServe(":8090", mh)
}
