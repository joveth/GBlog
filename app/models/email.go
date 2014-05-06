package models
import (
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"time"
	"crypto/md5"
	"io"
	"fmt"
)
type EmailObj struct{
	Email string
	ImgUrl string
	CDate time.Time
}
func (dao *Dao) InsertEmail(emailObj *EmailObj) error {
	emailCollection := dao.session.DB(DbName).C(EmailCollection)
	emailObj.CDate = time.Now();
	h := md5.New()
	io.WriteString(h, emailObj.Email)
  	emailObj.ImgUrl = fmt.Sprintf("%x", h.Sum(nil))
    fmt.Println(emailObj)
	_,err := emailCollection.Upsert(bson.M{"email": emailObj.Email}, emailObj)
	if err != nil {
		revel.WARN.Printf("Unable to save EmailObj: %v error %v", emailObj, err)
	}
	return err
}
func (dao *Dao) FindAllEmails() []EmailObj{
	emailCollection := dao.session.DB(DbName).C(EmailCollection)
	emails := []EmailObj{}
	query := emailCollection.Find(bson.M{}).Sort("-cdate")
	query.All(&emails)
	return emails
}
