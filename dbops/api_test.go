package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func clearTables(){
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M){
	 clearTables()
	 m.Run()
	 clearTables()
}


func TestUserWorkFlow(t *testing.T){
	t.Run("add",testAddUser)
	t.Run("get",testGetUser)
	t.Run("del",testDelUser)
	t.Run("reget",testReGetUser)
}

func testAddUser(t *testing.T){
	err := AddUserCredential("fff","111")
	if err != nil{
		t.Error(err)
	}
}

func testGetUser(t *testing.T){
	pwd,err := GetUsercredential("fff")
	if err != nil{
		t.Errorf("Error!")
	}
	t.Log(pwd)
}

func testDelUser(t *testing.T){
	err :=DeleteUser("fff","111")
	if err != nil{
		t.Errorf("Error!")
	}
}

func testReGetUser(t *testing.T){
	pwd,err := GetUsercredential("fff")
	if err != nil{
		t.Errorf("Error!")
	}
	t.Log(pwd)
}


func TestComments(t *testing.T){
	clearTables()
	t.Run("adduser",testAddUser)
	t.Run("addcomments",testAddComments)
	t.Run("Listcomments",testListComments)
}

func testAddComments(t *testing.T){
	vid := "12345"
	aid := 1
	content := "i like it"
	err := AddNewComments(vid,aid,content)
	if  err!=nil{
		t.Errorf("cuole")
	}
}
func testListComments(t *testing.T){
	vid := "12345"
	from := 1514764800
	to,_ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000,10))
	res ,err := ListComments(vid,from,to)
	if  err!=nil{
		t.Error(err)
	}
	for i, ele := range res{
		fmt.Println(i,ele)
	}
	}