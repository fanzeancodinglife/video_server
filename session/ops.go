package session

import (
	"../dbops"
	"../defs"
	"../utils"
	"log"
	"sync"
	"time"
)
/*
ops用来处理系统事务，主要是session
*/
func NowInMilli() int64{
	return time.Now().UnixNano()/100000
}

var sessionMap *sync.Map

func init(){
	sessionMap = &sync.Map{}

}


func LoadSessionFromDB(){
	r, err := dbops.RetriveAllSessions()
	if err != nil{
		return
	}
	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key,ss)
		return true
	})
}

func GenrateNewSessionId(un string) string {
	id,_:= utils.NewUUID()
	ct := NowInMilli()
	ttl := ct + 30*60*1000
	ss := &defs.SimpleSession{Username:un,TTL:ttl}
	sessionMap.Store(id,ss)
	_ = dbops.InsertSession(id,ttl,un)

	return id
}

func IsSessionExpired(sid string) (string,bool){
	ss, ok := sessionMap.Load(sid)
	if ok{
		ct := NowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
		//	delete expired session
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true

}

func deleteExpiredSession(sid string){
	sessionMap.Delete(sid)
	err := dbops.DeleteSession(sid)
	if err != nil{
		log.Print(err)
		return
	}
}

