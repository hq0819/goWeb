package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func main() {
	DSN := `yjyasd:ioklhguh@tcp(10.1.167.23:3306)/bwrp?parseTime=True`
	//DSN := `heqin:123456@tcp(120.48.45.173:3306)/go?parseTime=True`
	config := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	}
	db, err := gorm.Open(mysql.Open(DSN), &config)
	if err != nil {
		fmt.Println("数据库连接失败！")

	}
	//selectTest(db)

	i := make([]map[string]interface{}, 10, 10)
	db.Debug().Model(&LogCount{}).Find(&i)
	for _, m := range i {
		if m != nil {
			fmt.Println(m)

		}
	}
	/*	//db.Debug().Create(&LogCount{"李四",10})
		c := new(LogCount)
		db.Where(&LogCount{MapperMethod: "张三"}).Find(c)
		db.Debug().Where(LogCount{Counts: 10}).Where(LogCount{MapperMethod: "李四"}).Find(c)
		fmt.Println(*c)*/

}

type LogDetail struct {
	Id           int32     `gorm:"column:id;primary_key;auto_increment;type:int(10)"`
	Sqls         string    `gorm:"column:sqls;type:text"`
	MapperMethod string    `gorm:"column:mapperMethod;type:varchar(200)"`
	Parameters   string    `gorm:"column:parameters;type:varchar(2000)"`
	CostTime     int32     `gorm:"column:costTime;type:int(10)"`
	CreateTime   time.Time `gorm:"column:createTime"`
}

func (LogCount) TableName() string {
	return "bw_sql_log_counts"
}
func (LogDetail) TableName() string {
	return "bw_sql_log_detail"
}

type LogCount struct {
	MapperMethod string `gorm:"column:mapperMethod;type:varchar(200)"`
	Counts       int32  `gorm:"column:counts;type:int(10)"`
}

/*
func selectTest(db *gorm.DB) {
	userinfo := make([]User,50,50)

	db.Raw("select * from bw_user").Scan(&userinfo)
	sort.Slice(userinfo, func(i, j int) bool {
		return userinfo[i].Id-userinfo[j].Id<0
	})
	for _, user := range userinfo {
		fmt.Println(user)
	}
	fmt.Println("############################")
	user := NewUser("GAOLIQING")
	db.Table("bw_user").Where(user).Debug().Find(user)
	fmt.Println(user)



}
*/

/*type User struct {
	Id int32 `gorm:"column:id;primary_key"`
	Username string `gorm:"column:user_name"`
	Passwd string `gorm:"column:password"`
	CreateTime time.Time `gorm:"column:create_time"`
	UdpateDate time.Time `gorm:"column:update_time"`
	Founder string `gorm:"column:founder"`
	Modifier string `gorm:"column:modifier"`

}*/

/*
func NewUser(name string) *User {
	u := new(User)
	u.Username=name
	return u
}*/

/*type Student struct {
	ID int32 `gorm:"column:StuId;primary_key;auto_increment;"`
	Stuname string `gorm:"column:stuname;type:varchar(20);not null"`
	Age int8 `gorm:"column:age;type:int"`
	Email string `gorm:"column:email;type:varchar(16)"`
	Address string `gorm:"column:address;type:varchar(50)"`
}

func NewStudent(stuName string,age int8,email string,address string) *Student {
	return &Student{ID:1,Stuname: stuName,Age: age,Email: email,Address: address}
}

func (s Student)TableName()string  {
	return "t_student"
}
*/
