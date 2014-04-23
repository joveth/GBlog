package models
import (
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"time"
)
type History struct {
	Year int
	Blogs []Blog
}
func (dao *Dao) InsertHistory(history *History) error {
	historyCollection := dao.session.DB(DbName).C(HistoryCollection)
	err := historyCollection.Insert(history)
	if err != nil {
		revel.WARN.Printf("Unable to save History: %v error %v", history, err)
	}
	return err
}
func (dao *Dao) FindHistory() []History{
	historyCollection := dao.session.DB(DbName).C(HistoryCollection)
	his := []History{}
	query := historyCollection.Find(bson.M{}).Sort("-year")
	query.All(&his)
	return his
}
func (dao *Dao) RemoveAll() error{
	historyCollection := dao.session.DB(DbName).C(HistoryCollection)
	_,err := historyCollection.RemoveAll(bson.M{})
	if err != nil {
		revel.WARN.Printf("Unable to RemoveAll: error %v",  err)
	}
	return err
}
func (dao *Dao) CreateAllHistory() {
	dao.RemoveAll();
	var end int = time.Now().Year();
	for i:=BaseYear;i<=end;i++{
		history := new(History);
		history.Year = i;
		dao.InsertHistory(history);
	}
}