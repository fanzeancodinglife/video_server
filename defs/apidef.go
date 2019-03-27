package defs

//request
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}
type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}

type SimpleSession struct {
	Username string
	TTL int64
}

type SignUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}