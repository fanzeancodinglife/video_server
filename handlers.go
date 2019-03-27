package main

import (
	"encoding/json"
	"./defs"
	"./dbops"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"./session"
)

func CreateUser(w http.ResponseWriter, r *http.Request,p httprouter.Params){
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody);err != nil{
		sendErrorResponse(w,defs.ErrorRequestBodyParseFaild)
		return
	}
	if err := dbops.AddUserCredential(ubody.Username,ubody.Pwd);err!=nil{
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}
	id := session.GenrateNewSessionId(ubody.Username)
	su := &defs.SignUp{Success:true,SessionId:id}
	if resp, err := json.Marshal(su);err != nil {
		sendErrorResponse(w,defs.ErrorInterfnalFault)
		return
	}else {
		sendNormalResponse(w,string(resp),201)
	}
		}
func Login (w http.ResponseWriter, r *http.Request,p httprouter.Params){
	_, _ = io.WriteString(w,p.ByName("user_name"))
}