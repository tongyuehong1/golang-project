package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/tongyuehong1/golang-project/application/forum/common"
	"github.com/tongyuehong1/golang-project/libs/logger"
)

func init() {
	orm.RegisterModel(new(Comment))
}

type CommentServiceProvider struct {
}

var CommentServer *CommentServiceProvider

type Comment struct {
	CommentId uint64    `orm:"column(id);pk"      json:"id"`
	UserId    uint64    `orm:"column(userId)"     json:"userId"`
	ArtTitle  string    `orm:"column(artTitle)"   json:"artTitle"`
	Comment   string    `orm:"column(comment)"    json:"comment"`
	Created   time.Time `orm:"column(created)"    json:"created"`
	Status    uint8     `orm:"column(status)"     json:"status"`
}
type ShowComment struct {
	Comment  []Comment
	UserName []string
}

// 添加评论
func (this *CommentServiceProvider) Comment(comment Comment) error {
	o := orm.NewOrm()
	comment.Created = time.Now()
	sql := "INSERT INTO forum.comment(id, userId, artTitle, comment, created, status) VALUES(?,?,?,?,?,?)"
	values := []interface{}{comment.CommentId, comment.UserId, comment.ArtTitle, comment.Comment, comment.Created, common.NormalComment}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()

	return err
}

// 获取评论
func (this *CommentServiceProvider) GetComment(artTitle string) (ShowComment, error) {
	o := orm.NewOrm()
	var show ShowComment

	_, err := o.Raw("SELECT * FROM forum.comment WHERE artTitle=?", artTitle).QueryRows(&show.Comment)
	if err != nil {
		return show, err
	}
	for _, comment := range show.Comment {
		userName, err := CommentServer.GetUserName(comment.UserId)
		if err != nil {
			logger.Logger.Error("ERROR:", err)
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
