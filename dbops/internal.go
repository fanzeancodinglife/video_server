package dbops

import (
	"../defs"
	"database/sql"
	"log"
	"strconv"
	"sync"
)

func InsertSession(sid string, ttl int64,uname string)error{
	ttlstr := strconv.FormatInt(ttl,10)
	stmtIns, err := dbConn.Prepare("INSERT  into sessions (session_id,TTL,login_name) values (?,?,?)")
	if err != nil{
		return err
	}
	_, err = stmtIns.Exec(sid,ttlstr,uname)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}


func RetireveSession(sid string)(*defs.SimpleSession,error){
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("select TTL,login_name from sessions Where session_id = ?")
	if err != nil{
		return  nil,err
	}
	var ttl string
	var uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl,&uname)
	if err != nil && err != sql.ErrNoRows{
		return nil,err
	}
	if res,err := strconv.ParseInt(ttl,10,64);err == nil{
		ss.TTL = res
		ss.Username = uname
	}else {
		return nil,err
	}
	defer stmtOut.Close()
	return ss, nil
}

func RetriveAllSessions()(*sync.Map,error){
	m := &sync.Map{}
	stmtOut,err := dbConn.Prepare("select * from sessions")
	if err != nil{
		return nil,err
	}
	rows, err := stmtOut.Query()
	if err != nil{
		return nil,err
	}
	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if er := rows.Scan(&id,&ttlstr,&login_name); er != nil{
			log.Printf("error:%s",er)
			break
		}
		if ttl,err1 := strconv.ParseInt(ttlstr,10,64);err1 == nil{
			ss := &defs.SimpleSession{Username:login_name,TTL:ttl}
			m.Store(id,ss)
			log.Printf("session id :%s,ttl:%d",id,ttl)
		}
	}
	defer stmtOut.Close()
	return m,nil
}

func DeleteSession(sid string)error{
	stmtDel, err := dbConn.Prepare("delete from sessions where session_id = ?")
	if err != nil{
		log.Printf("%s",err)
		return err
	}
	_,err = stmtDel.Query(sid)
	if err != nil{
		log.Printf("%s",err)
		return err
	}

	defer stmtDel.Close()
	return nil
}