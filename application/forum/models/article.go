package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/tongyuehong1/golang-project/application/forum/common"
	"github.com/tongyuehong1/golang-project/libs/logger"
	"time"
	"unicode/utf8"
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
func (this *ArticleServiceProvider) Insert(article Article, name string) (int64,error) {
	o := orm.NewOrm()
	user, _ := UserServer.GetUserId(name)
	var id int64
	lasttime, err := UserServer.GetLastTime(user)
	//present := time.Now()
	//a, _ := time.ParseDuration("5m")
	now := time.Now()
	m, _ := time.ParseDuration("5m")
	m1 := lasttime.Add(m)
	//fmt.Println(m1)
	subM := now.Sub(m1)
	if subM.Minutes() > 5{
		article.Created = time.Now()
		article.Author = name
		sql := "INSERT INTO forum.article(title,category,content,author,created,brief,status) VALUES(?,?,?,?,?,?,?)"
		values := []interface{}{article.Title, article.Category, article.Content, article.Author, article.Created, article.Brief, common.NormalArticle}
		raw := o.Raw(sql, values)
		resu,err:=raw.Exec()
		if err != nil {
			return 0,err
		}
		id, _ = resu.LastInsertId()
	} else {
		logger.Logger.Info("添加出现错误：", "超出五分钟")
		return 0, err
	}
	_, err = ArticleServer.InsertLastTime(user)
	return id,err
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

// 搜索文章
func (this *ArticleServiceProvider) SearchArticle(title string) ([]Article, error) {
	o := orm.NewOrm()
	var articles []Article
	var str string
	for len(title) > 0 {
		r, size := utf8.DecodeRuneInString(title)
		title = title[size:]

		str += string(r) + "%"
	}
	str = "%" + str
	_, err := o.Raw("SELECT * FROM forum.article WHERE title LIKE ?", str).QueryRows(&articles)
	return articles, err
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

// 获取文章id
func (this *ArticleServiceProvider) GetArticleId(title string) (uint64, error) {
	o := orm.NewOrm()
	var articleId uint64

	err := o.Raw("SELECT id FROM forum.article WHERE title = ?", title).QueryRow(&articleId)

	return articleId, err
}

// 收藏文章（取消收藏）
func (this *ArticleServiceProvider) Collect(articleId uint64, userId uint64) error {
	o := orm.NewOrm()
	var value string

	err := o.Raw("SELECT value FROM forum.userextra WHERE id=? AND `key`=? AND value=? LIMIT 1 LOCK IN SHARE MODE", userId, common.KeyCollection, articleId).QueryRow(&value)

	if err == orm.ErrNoRows {
		// 未收藏，开始收藏
		sql := "INSERT INTO forum.userextra(id,`key`,value)VALUES(?,?,?)"
		values := []interface{}{userId, common.KeyCollection, articleId}
		raw := o.Raw(sql, values)
		_, err := raw.Exec()

		return err
	} else if err == nil {
		// 已经收藏，取消收藏
		sql := "DELETE FROM forum.userextra WHERE value=? AND id=? AND `key`=? LIMIT 1"
		values := []interface{}{articleId, userId, common.KeyCollection}
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

	for _, title := range collection {
		article := Article{}
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

func (this *ArticleServiceProvider) InsertLastTime(userid uint64)  (int64,error) {
	o := orm.NewOrm()
	//var created time.Time
	//var last time.Time
	//err := o.Raw("SELECT created FROM forum.article WHERE author=? ORDER BY created DESC", userid).QueryRow(&created)
	//
	//if err != nil {
	//	return  err
	//} else {
	//	fmt.Println("InsertLastTime err:",err)
	//}
	//o.Using("forum")
	//sr := last.Unix()
	//lasttime := time.Unix(sr, 0).Format("2006-01-02 15:04:05")

	//userextra := UserExtra{Key: "KeyLastInsert", UserID: userid, Value: lasttime}
	//isCreated, id, err := o.ReadOrCreate(&userextra, "Key", "UserID")
	//if err == nil {
	//	if !isCreated {
	//		sql := "UPDATE forum.userextra SET value=? WHERE id=? AND `key`=?"
	//		values := []interface{}{last, userid, common.KeyLastInsert}
	//		raw := o.Raw(sql, values)
	//		_, err = raw.Exec()
	//		return id, err
	//	}
	//}
	//sql := "UPDATE forum.userextra SET value=? WHERE userid=? AND `key`=?"
	//values := []interface{}{time.Now().Format("2006-01-02 15:04:05"), userid, common.KeyLastInsert}
	//raw := o.Raw(sql, values)
	//_, err := raw.Exec()
	var last UserExtra
	o.Raw("SELECT * FROM forum.userextra WHERE userid=? AND `key`=?",userid, common.KeyLastInsert).QueryRow(&last)
	last.Value = time.Now().Format("2006-01-02 15:04:05")
	//var u  UserExtra=UserExtra{Key: common.KeyLastInsert, UserID: userid, Value: time.Now().Format("2006-01-02 15:04:05")}
	id,err := o.Update(&last, "value")
	return  id,err
}
