package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"fmt"
)

func init() {
	orm.RegisterModel(new(Article))
}

func (this *Article) TableName() string {
	return "article"
}

type Article struct {
	Id      int
	Classes string
	Title   string
	Created    time.Time
	Brief   string
	Article string
	Status  bool
}
type ArticleServiceProvider struct {
}

var ArticleServer *ArticleServiceProvider

func (this *ArticleServiceProvider) Insert(article Article) error {
	fmt.Println("a:", article)
	o := orm.NewOrm()
	article.Created = time.Now()
	sql := "INSERT INTO article.article(classes,title,brief,article,status) VALUES(?,?,?,?,?)"
	values := []interface{}{article.Classes, article.Title, article.Brief, article.Article, article.Status}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	fmt.Println("err:", err)
	//fmt.Println("sdfg")
	//o.Using("article")
	//fmt.Println(o.Driver())
	//_, err := o.Insert(&article)
	//fmt.Println("err:", err)

	return err
}

func (this *ArticleServiceProvider) Update(title string, article string) error {
	o := orm.NewOrm()
	sql := "UPDATE article SET Article=? WHERE Title=? AND Status=? LIMIT 1"
	values := []interface{}{article, title, true}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()

	return err
}

func (this *ArticleServiceProvider) Delete(title string) error {
	o := orm.NewOrm()
	sql := "UPDATE Article SET Status=? WHERE Title=? LIMIT 1"
	values := []interface{}{false, title}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()

	return err
}

//func (this *ArticleServiceProvider) Get() {
//	o := orm.NewOrm()
//	err := o.Raw("SELECT * WHERE STATUS = false LIMIT 1 LOCK IN SHARE MODE")
//	return
//}
