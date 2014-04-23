package models
import (
	"labix.org/v2/mgo"
)
const (
	DbName = "ominds"
	BlogCollection = "blogs"
	CommentCollection = "gb_comments"
	MessageCollection = "gb_messages"
	HistoryCollection = "gb_historys"
	EmailCollection = "gb_emails"
	BaseYear = 2014
)
type Dao struct {
	session *mgo.Session
}
func NewDao() (*Dao, error){
	//session, err := mgo.Dial("mongodb://omind:jov.omind@widmore.mongohq.com:10000/ominds")
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return &Dao{session}, nil
}
func (d *Dao) Close() {
	d.session.Close()
}