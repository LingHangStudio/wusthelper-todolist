package dao

import (
	"fmt"
	"github.com/yitter/idgenerator-go/idgen"
	"os"
	"testing"
	"time"
	"wusthelper-todolist-service/app/conf"
	"wusthelper-todolist-service/app/model"
)

var (
	dao       *Dao
	legacyDao *Dao
)

func TestMain(m *testing.M) {
	options := idgen.NewIdGeneratorOptions(8)
	idgen.SetIdGenerator(options)

	if err := conf.Init(); err != nil {
		panic(err)
	}

	conf.Conf.Database.PrintSql = false
	dao = New(conf.Conf)

	conf.Conf.Database.DbName = "mp_wusthelper"
	legacyDao = New(conf.Conf)

	m.Run()
	os.Exit(0)
}

type WechatUser struct {
	WechatOpenid string    `gorm:"wechat_openid,omitempty"`
	Nickname     string    `gorm:"nickname,omitempty"`
	Gender       int32     `gorm:"gender,omitempty"`
	City         string    `gorm:"city,omitempty"`
	Province     string    `gorm:"province,omitempty"`
	Country      string    `gorm:"country,omitempty"`
	AvatarUrl    string    `gorm:"avatar_url,omitempty"`
	AddTime      time.Time `gorm:"add_time,omitempty"`
}

type QQUser struct {
	QqOpenid  string    `gorm:"qq_openid,omitempty"`
	Nickname  string    `gorm:"nickname,omitempty"`
	Gender    int32     `gorm:"gender,omitempty"`
	City      string    `gorm:"city,omitempty"`
	Province  string    `gorm:"province,omitempty"`
	Country   string    `gorm:"country,omitempty"`
	AvatarUrl string    `gorm:"avatar_url,omitempty"`
	AddTime   time.Time `gorm:"add_time,omitempty"`
}

func (*Dao) TableName() string {
	return "wechat_user"
}

type User struct {
	Stuid             string    `json:"stuid,omitempty"`
	StuName           string    `json:"stu_name,omitempty"`
	QqOpenid          string    `json:"qq_openid,omitempty"`
	WechatOpenid      string    `json:"wechat_openid,omitempty"`
	Sex               string    `json:"sex,omitempty"`
	ClassName         string    `json:"className,omitempty"`
	Major             string    `json:"major,omitempty"`
	College           string    `json:"college,omitempty"`
	AddTime           time.Time `json:"add_time,omitempty"`
	LastLoginTime     string    `json:"last_login_time,omitempty"`
	LastLoginPlatform string    `json:"last_login_platform,omitempty"`
}

func TestUserTransfer(t *testing.T) {
	legacyUndergrad := make([]User, 0)
	err := legacyDao.db.Table("user").Find(&legacyUndergrad)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("total:")
	fmt.Println(len(legacyUndergrad))

	count := 0
	for _, legacyStudent := range legacyUndergrad {
		now := time.Now()
		student := model.Student{
			Sid:        legacyStudent.Stuid,
			Name:       legacyStudent.StuName,
			College:    legacyStudent.College,
			Major:      legacyStudent.Major,
			Clazz:      legacyStudent.ClassName,
			CreateTime: legacyStudent.AddTime,
			UpdateTime: now,
			Deleted:    0,
		}

		_, err := dao.db.InsertOne(student)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if legacyStudent.WechatOpenid != "" {
			user := model.UserBasic{
				Sid:        legacyStudent.Stuid,
				UpdateTime: now,
			}
			_, err := dao.db.Where("oid = ?", legacyStudent.WechatOpenid).MustCols("sid").Update(user)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		if legacyStudent.QqOpenid != "" {
			user := model.UserBasic{
				Sid:        legacyStudent.Stuid,
				UpdateTime: now,
			}
			_, err := dao.db.Where("oid = ?", legacyStudent.QqOpenid).MustCols("sid").Update(user)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		count++
	}

	fmt.Println("processed:")
	fmt.Println(count)
}
