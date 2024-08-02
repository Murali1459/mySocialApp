package database

import "github.com/beego/beego/v2/client/orm"

type Follow struct {
	Id         int `orm:"column(id)"`
	FollowerId int `orm:"column(follower_id)"`
	FolloweeId int `orm:"column(followee_id)"`
}

func init() {
	orm.RegisterModel(new(Follow))
}

func (f *Follow) TableName() string {
	return "follow"
}

func (d *Db) AddNewFollower(f Follow) error {
	o := orm.NewOrm()
	_, err := o.Insert(&f)
	return err
}

func (d *Db) UnfollowUser(followerId, followeeId int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Follow)).Filter("follower_id", followerId).Filter("followee_id", followeeId).Delete()
	return err
}

func (d *Db) CheckForFollow(followerId, followeeId int) bool {
	o := orm.NewOrm()
	return o.QueryTable(new(Follow)).Filter("follower_id", followerId).Filter("followee_id", followeeId).Exist()
}

func (d *Db) GetAllFollowersForUser(userId int) ([]Follow, error) {
	o := orm.NewOrm()
	allFollowers := []Follow{}
	_, err := o.QueryTable(new(Follow)).Filter("followee_id", userId).All(&allFollowers)
	return allFollowers, err
}
