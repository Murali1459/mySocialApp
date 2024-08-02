package database

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/spf13/cast"
)

type Post struct {
	Id       int    `orm:"column(id)"`
	UserId   int    `orm:"column(user_id)"`
	Content  string `orm:"column(content)"`
	ImageUrl string `orm:"column(image_url)"`
}

func init() {
	orm.RegisterModel(new(Post))
}

func (p *Post) TableName() string {
	return "post"
}

func (d *Db) GetPostById(id int) (Post, error) {
	o := orm.NewOrm()
	post := Post{Id: id}
	err := o.Read(&post)
	return post, err
}

func (d *Db) AddNewPost(newPost Post) error {
	o := orm.NewOrm()
	id, err := o.Insert(&newPost)
	newPost.Id = cast.ToInt(id)
	return err
}

func (d *Db) GetAllPostsForUser(userId int) ([]Post, error) {
	o := orm.NewOrm()
	posts := []Post{}
	_, err := o.QueryTable(new(Post)).Filter("userId", userId).All(&posts)
	return posts, err
}
