package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"

	//"github.com/fanzeancodinglife/video_server/api/defs"
	"../defs"
	"../utils"
)

//func openConn() *sql.DB{
//	dbConn ,err :=sql.Open("mysql","root:1234@tcp(localhost:3306)/video_server?charset=utf8i")
//	if err!=nil{
//		panic(err.Error())
//	}
//	return dbConn
//}

func AddUserCredential(loginName string,pwd string) error{
	stmtIns,err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)")
	if err != nil{
		return err
	}
	_,err =stmtIns.Exec(loginName,pwd)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUsercredential(loginName string)(string,error){
	stmtOut,err := dbConn.Prepare("select pwd from users where login_name = ?")
	if err != nil{
		log.Printf("%s",err)
		return "",err
	}
	var pwd string
	err =stmtOut.QueryRow(loginName).Scan(&pwd)
	// 注意这里的错误处理go
	if err!=nil && err !=sql.ErrNoRows {
		return "" ,err
	}
	defer stmtOut.Close()
	return pwd,nil
}

func DeleteUser(loginName string,pwd string) error{
	stmtDel,err := dbConn.Prepare("delete from users where login_name = ? and pwd = ?")
	if err != nil{
		log.Printf("%s",err)
		return err
	}
	_,err =stmtDel.Exec(loginName,pwd)
	if err != nil{
		log.Printf("%s",err)
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewVideo(aid int,name string)(*defs.VideoInfo, error){
	vid, err := utils.NewUUID()
	if err != nil{
		return nil,err
	}
	// 页面显示的时间
	t := time.Now()
	// 必须用这个数字串格式化
	ctime := t.Format("Jan 02 2006,15:04:05")
	stmtIns ,err := dbConn.Prepare(`INSERT into video_info (id,author_id,name,display_ctime) 
	values(?,?,?,?)`)
	if err != nil{
		return nil,err
	}
	_,err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil{
		return nil,err
	}
	res := &defs.VideoInfo{Id:vid,AuthorId:aid, Name:name ,DisplayCtime:ctime}
	defer stmtIns.Close()
	return res,nil
}

func GetVideoInfo(vid string)(*defs.VideoInfo,error){
	stmtOut, err := dbConn.Prepare("select author_id,name,display_ctime from video_info")
	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows{
		return nil,err
	}
	defer stmtOut.Close()
	res := &defs.VideoInfo{Id:vid,AuthorId:aid, Name:name ,DisplayCtime:dct}
	return res, nil
	}

func DeleteVideoInfo(vid string) error{
	stmtDel ,err := dbConn.Prepare("delete  from video_info where id = ?")
	if err != nil{
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil{
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewComments(vid string, aid int, content string) error{
	id,err := utils.NewUUID()
	if err != nil{
		return err
	}
	stmtIns, err := dbConn.Prepare("insert into comments (id,video_id,author_id,content) values (?,?,?,?)")
	if err != nil{
		return err
	}
	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}

func ListComments(vid string,from,to int)([] *defs.Comment, error){
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name, comments.content
	from comments INNER join users on comments.author_id = users.id where comments.video_id = ? and comments.time >from_unixtime(?)
	and comments.time <=from_unixtime(?)`)
	if err != nil{
		return nil,err
	}
	var res []*defs.Comment
	row,err := stmtOut.Query(vid,from,to)
	if err != nil{
		return res,err
	}
	for row.Next(){
		var id, name, content string
		if err := row.Scan(&id,&name,&content); err != nil {
			return res,err
		}
		c := &defs.Comment{Id:id,Author:name,Content:content}
		res = append(res,c)
	}
	defer stmtOut.Close()
	return res,nil
}