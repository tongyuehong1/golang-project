package article
import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)
type Article struct {
	Classes string
	Title   string
	Time    time.Time
	Brief   string
	Article string
	Status  bool
}
func (this *Article) Insert(article Article) error {
	o := orm.NewOrm()
	sql := "INSERT INTO Article(classes,Title,Brief,Article) VALUES(?,?,?)"
	values := []interface{}{article.Classes, article.Title, article.Brief, article.Article}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()

	return err
}
//func (this *Article) get(Title string) (uint64, error) {
//	o := orm.NewOrm()
//	var id uint64
//	err := o.Raw("SELECT id FROM Article WHERE Title=?").QueryRow(&id)
//	return id, err
//}
func (this *Article) Update(title string, article string) error {
	o := orm.NewOrm()
	sql := "UPDATE Article SET Article=? WHERE Title=? AND Status=? LIMIT 1"
	values := []interface{}{article, title, true}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()

	return err
}
func (this *Article) Delete(title string) error {
	o := orm.NewOrm()
	sql := "UPDATE Article SET Status=? WHERE Title=? LIMIT 1"
	values := []interface{}{true, title}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()

	return err
}