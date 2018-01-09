package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(Comment))
}

type CommentServiceProvider struct {
}

var CommentServer *CommentServiceProvider

type Comment struct {
	CommentId uint64	`orm:"column(id);pk"`
	UserId    uint64
	ArticleId uint64    `orm:"column(articleId)"`
	Comment   string	`orm:"column(comment)"`
	Created   time.Time
	Status    uint8
}
type ShowComment struct {
	Comment  []Comment
	UserName []string
}

// 添加评论
func (this *CommentServiceProvider) Comment(userId, articleId uint64, comment string) error {
	o := orm.NewOrm()

	sql := "INSERT INTO forum.comment(id, articleId, comment) VALUES(?,?,?)"
	values := []interface{}{userId, articleId, comment}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()

	return err
}

// 获取评论
func (this *CommentServiceProvider) GetComment(ArticleId uint64) (ShowComment, error) {
	o := orm.NewOrm()
	var show ShowComment

	_, err := o.Raw("SELECT * FROM forum.comment WHERE articleId=?, ArticleId").QueryRows(&show.Comment)
	if err != nil {
		return show, err
	}
	for _, comment := range show.Comment {
		userName, err := CommentServer.GetUserName(comment.UserId)
		if err != nil {
			return show, err
		}
		show.UserName = append(show.UserName, userName)
	}
	return show, nil
}

// 获取用户名
func (this *CommentServiceProvider) GetUserName(UserId uint64) (string, error) {
	o := orm.NewOrm()
	var name string

	err := o.Raw("SELECT name FROM forum.user WHERE id=? LIMIT 1", UserId).QueryRow(&name)

	return name, err
}
