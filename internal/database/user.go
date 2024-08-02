package database

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Id         int    `orm:"column(id)"`
	UserName   string `orm:"column(userName)"`
	Email      string `orm:"column(email)"`
	Password   string `orm:"column(password_hash)"`
	ProfilePic string `orm:"column(profile_picture_url)"`
	Bio        string `orm:"column(bio)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "user"
}

func (d *Db) GetUserById(id int) (User, error) {
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (d *Db) AddUser(u *User) error {
	o := orm.NewOrm()
	_, err := o.Insert(u)
	if err != nil {
		return err
	}
	return nil
}

func (d *Db) GetUserByName(n string) (User, error) {
	o := orm.NewOrm()
	u := User{}
	err := o.QueryTable(new(User)).Filter("userName", n).One(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}
