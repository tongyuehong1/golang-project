package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/tongyuehong1/golang-project/application/forum/common"
	"time"
	"github.com/tongyuehong1/another-golang/beego-training/libs/logger"
)

func init() {
	orm.RegisterModel(new(Article))
}

type ArticleServiceProvider struct {
}

var ArticleServer *ArticleServiceProvider

// 文章基本信息
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

// 文章基本信息显示
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

// 文章其他信息（点赞，收藏）
type ArticleExtraInfo struct {
	Title string
	Key   string
	Value string
}

// 添加文章
func (this *ArticleServiceProvider) Insert(article Article) error {
	o := orm.NewOrm()
	article.Created = time.Now()
	sql := "INSERT INTO forum.article(id,title,category,content,author,created,brief,status) VALUES(?,?,?,?,?,?,?,?)"
	values := []interface{}{article.ArticleId, article.Title, article.Category, article.Content, article.Author, article.Created, article.Brief, common.NormalArticle}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	return err
}

// 修改文章
func (this *ArticleServiceProvider) Change(article Article) error {
	o := orm.NewOrm()
	sql := "UPDATE forum.article SET title=?,category=?,content=?,author=?,created=?,brief=?,status=? WHERE id=? LIMIT 1"
	values := []interface{}{article.Title, article.Category, article.Content, article.Author, article.Created, article.Brief, article.Status, article.ArticleId}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	return err
}

// 推荐文章
func (this *ArticleServiceProvider) Recommend(title string) error {
	o := orm.NewOrm()
	sql := "UPDATE forum.article SET status=? WHERE id=? LIMIT 1"
	values := []interface{}{common.RecommendArticle, title}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	return err
}

// 根据类别获取文章
func (this *ArticleServiceProvider) GetArticle(category string) ([]ShowArticle, error) {
	var showarticle []ShowArticle
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM forum.article WHERE category=? AND status!=?", category, 2).QueryRows(&showarticle)
	return showarticle, err
}

// 获取所有文章
func (this *ArticleServiceProvider) AllArticle() ([]ShowArticle, error) {
	var showarticle []ShowArticle
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM forum.article WHERE status!=?", 2).QueryRows(&showarticle)
	return showarticle, err
}

// 删除文章
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

// 收藏文章（取消收藏）
func (this *ArticleServiceProvider) Collect(title string, userId uint64) error {
	o := orm.NewOrm()
	var value string

	err := o.Raw("SELECT value FROM forum.userextra WHERE id=? AND `key`=? AND value=? LIMIT 1 LOCK IN SHARE MODE", userId, common.KeyCollection, title).QueryRow(&value)

	if err == orm.ErrNoRows {
		// 未收藏，开始收藏
		sql := "INSERT INTO forum.userextra(id,`key`,value)VALUES(?,?,?)"
		values := []interface{}{userId, common.KeyCollection, title}
		raw := o.Raw(sql, values)
		_, err := raw.Exec()

		return err
	} else if err == nil {
		// 已经收藏，取消收藏
		sql := "DELETE FROM forum.userextra WHERE value=? AND id=? AND `key`=? LIMIT 1"
		values := []interface{}{userId, common.KeyCollection, title}
		raw := o.Raw(sql, values)
		_, err := raw.Exec()

		return err
	}
	return err
}

// 显示收藏文章
func (this *ArticleServiceProvider) ShowCollection(userId uint64) ([]Article, error) {
	o := orm.NewOrm()
	var articles []Article
	var collection []string
	_, err := o.Raw("SELECT value FROM forum.userextra WHERE `key`=? AND id=?", common.KeyCollection, userId).QueryRows(&collection)

	if err != nil {
		return articles, err
	}

	logger.Logger.Info("collection:", collection)
	for _, title := range collection {
		article := Article{}
		logger.Logger.Info("title:", title)

		if err != nil {
			return articles, err
		}

		err = o.Raw("SELECT * FROM  forum.article WHERE title=? LIMIT 1 LOCK IN SHARE MODE", title).QueryRow(&article)

		if err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}
