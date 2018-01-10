package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/tongyuehong1/golang-project/application/forum/common"
	"time"
)

func init() {
	orm.RegisterModel(new(Article))
}

type ArticleServiceProvider struct {
}

var ArticleServer *ArticleServiceProvider

type Article struct {
	ArticleId uint64    `orm:"column(id);pk"         json:"id"`
	Title     string    `orm:"column(title)"         json:"title"`
	Category  string    `orm:"column(category)"      json:"category"`
	Content   string    `orm:"column(content)"       json:"content"`
	Author    string    `orm:"column(author)"        json:"author"`
	Created   time.Time `orm:"column(created)"       json:"created"`
	Brief     string    `orm:"column(brief)"         json:"brief"`
	Status    uint8     `orm:"column(status)"        json:"status"`
}
type ShowArticle struct {
	ArticleId uint64    `orm:"column(id);pk"         json:"id"`
	Title     string    `orm:"column(title)"         json:"title"`
	Category  string    `orm:"column(category)"      json:"category"`
	Content   string    `orm:"column(content)"       json:"content"`
	Author    string    `orm:"column(author)"        json:"author"`
	Created   time.Time `orm:"column(created)"       json:"created"`
	Brief     string    `orm:"column(brief)"         json:"brief"`
	Status    uint8     `orm:"column(status)"        json:"status"`
}

func (this *ArticleServiceProvider) Insert(article Article) error {
	o := orm.NewOrm()
	article.Created = time.Now()
	sql := "INSERT INTO forum.article(id,title,category,content,author,created,brief,status) VALUES(?,?,?,?,?,?,?,?)"
	values := []interface{}{article.ArticleId, article.Title, article.Category, article.Content, article.Author, article.Created, article.Brief, common.NormalArticle}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	return err
}

func (this *ArticleServiceProvider) Change(article Article) error {
	o := orm.NewOrm()
	sql := "UPDATE forum.article SET title=?,category=?,content=?,author=?,created=?,brief=?,status=? WHERE id=? LIMIT 1"
	values := []interface{}{article.Title, article.Category, article.Content, article.Author, article.Created, article.Brief, article.Status, article.ArticleId}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	return err
}

func (this *ArticleServiceProvider) Recommend(title string) error {
	o := orm.NewOrm()
	sql := "UPDATE forum.article SET status=? WHERE id=? LIMIT 1"
	values := []interface{}{common.RecommendArticle, title}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	return err
}

func (this *ArticleServiceProvider) GetArticle(category string) ([]ShowArticle, error) {
	var showarticle []ShowArticle
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM forum.article WHERE category=? AND status!=?", category, 2).QueryRows(&showarticle)
	return showarticle, err
}

func (this *ArticleServiceProvider) AllArticle() ([]ShowArticle, error) {
	var showarticle []ShowArticle
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM forum.article WHERE status!=?", 2).QueryRows(&showarticle)
	return showarticle, err
}

func (this *ArticleServiceProvider) DeleteArticle(title string) error {
	o := orm.NewOrm()
	sql := "UPDATE forum.article SET status=? WHERE title=? LIMIT 1"
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
