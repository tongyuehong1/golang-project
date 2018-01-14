package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/tongyuehong1/golang-project/application/blog/common"
	"time"

	"github.com/tongyuehong1/golang-project/libs/logger"
)

func init() {
	orm.RegisterModel(new(Article))
}

func (this *Article) TableName() string {
	return "article"
}

type Article struct {
	Id      int       `orm:"column(id);pk"    json:"id"`
	Classes string    `orm:"column(classes)"  json:"classes"`
	Title   string    `orm:"column(title)"    json:"title"`
	Created time.Time `orm:"column(created)"  json:"created"`
	Brief   string    `orm:"column(brief)"    json:"brief"`
	Article string    `orm:"column(article)"  json:"article"`
	Status  int       `orm:"column(status)"   json:"status"`
}

type Show struct {
	Id      int
	Classes string
	Title   string
	Created time.Time
	Brief   string
	Article string
	Status  int
}

type ArticleServiceProvider struct {
}

var ArticleServer *ArticleServiceProvider

func (this *ArticleServiceProvider) Insert(article Article) error {
	o := orm.NewOrm()
	article.Created = time.Now()
	logger.Logger.Info("article:", article)
	sql := "INSERT INTO article.article(classes,title,brief,article,status) VALUES(?,?,?,?,?)"
	values := []interface{}{article.Classes, article.Title, article.Brief, article.Article, common.NormalArticle}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	return err
}

func (this *ArticleServiceProvider) UpdateArticle(title string, article string) error {
	o := orm.NewOrm()
	sql := "UPDATE article.article SET article=? WHERE Title=? AND status=? LIMIT 1"
	values := []interface{}{article, title, true}
	raw := o.Raw(sql, values)
	result, err := raw.Exec()
	if err == nil {
		if row, _ := result.RowsAffected(); row == 0 {
			return common.ErrNotFound
		}
	}

	return err
}

func (this *ArticleServiceProvider) UpdateTitle(title string, changetitle string) error {
	o := orm.NewOrm()
	sql := "UPDATE article.article SET title=? WHERE title=? AND status=? LIMIT 1"
	values := []interface{}{changetitle, title, true}
	raw := o.Raw(sql, values)
	result, err := raw.Exec()
	if err == nil {
		if row, _ := result.RowsAffected(); row == 0 {
			return common.ErrNotFound
		}
	}

	return err
}

func (this *ArticleServiceProvider) UpdateBrief(title string, brief string) error {
	o := orm.NewOrm()
	sql := "UPDATE article.article SET brief=? WHERE title=? AND status=? LIMIT 1"
	values := []interface{}{brief, title, true}
	raw := o.Raw(sql, values)
	result, err := raw.Exec()
	if err == nil {
		if row, _ := result.RowsAffected(); row == 0 {
			return common.ErrNotFound
		}
	}

	return err
}

func (this *ArticleServiceProvider) Delete(title string) error {
	o := orm.NewOrm()
	sql := "UPDATE Article SET status=? WHERE title=? LIMIT 1"
	values := []interface{}{common.RemovedArticle, title}
	raw := o.Raw(sql, values)
	result, err := raw.Exec()
	if err == nil {
		if row, _ := result.RowsAffected(); row == 0 {
			return common.ErrNotFound
		}
	}

	return err
}

func (this *ArticleServiceProvider) Get(classes string) ([]Show, error) {
	var show []Show
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM article.article WHERE classes=? AND status=?", classes, common.NormalArticle).QueryRows(&show)
	return show, err
}
