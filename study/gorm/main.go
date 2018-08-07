package main

import (
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

type (
	// User present plateform user
	User struct {
		ID             int       `json:"id,omitempty" gorm:"primary_key"`
		Name           string    `json:"name,omitempty" gorm:"UNIQUE_INDEX"`
		Displayname    string    `json:"displayname,omitempty"`
		Password       string    `json:"password,omitempty"`
		Email          string    `json:"email,omitempty"`
		Phone          string    `json:"phone,omitempty"`
		LoginFrequency int       `json:"loginFrequency,omitempty"`
		Active         int8      `json:"active,omitempty" gorm:"DEFAULT:1"`
		APIToken       string    `json:"apiToken,omitempty"`
		Role           int32     `json:"role,omitempty" gorm:"DEFAULT:1"`
		Teams          []Team    `json:"teams,omitempty" gorm:"many2many:user_teams;"`
		Spaces         []Space   `json:"spaces,omitempty" gorm:"many2many:user_spaces;"`
		CreatedAt      time.Time `json:"create_at,omitempty"`
		UpdatedAt      time.Time `json:"update_at,omitempty"`
	}

	// Team present plateform team
	Team struct {
		ID          int       `json:"id,omitempty" gorm:"primary_key"`
		Name        string    `json:"name,omitempty" gorm:"UNIQUE_INDEX"`
		Description string    `json:"description,omitempty"`
		CreatorID   int       `json:"creatorID,omitempty"`
		Spaces      []Space   `json:"spaces,omitempty"`
		Users       []User    `json:"users,omitempty"  gorm:"many2many:user_teams;"`
		CreatedAt   time.Time `json:"create_at,omitempty"`
		UpdatedAt   time.Time `json:"update_at,omitempty"`
	}

	// Space present plateform namespace
	Space struct {
		ID          int       `json:"id,omitempty" gorm:"primary_key"`
		Name        string    `json:"name,,omitempty" gorm:"UNIQUE_INDEX"`
		Description string    `json:"description,omitempty"`
		Type        int       `json:"type,omitempty"` // 1 personal namespace 2 team's namespace
		Users       []User    `json:"users,omitempty" gorm:"many2many:user_spaces;"`
		TeamID      int       `json:"teamID"`
		CreatedAt   time.Time `json:"create_at,omitempty"`
	}

	Pagination struct {
		Size int
		Num  int
	}

	Cluster struct {
		ID      string `json:"id,omitempty" gorm:"primary_key"`
		Content string `json:"content,omitempty" sql:"type:text"`
	}
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open("mysql", "root:b355dfda0dd5867c@tcp(10.10.233.116:32326)/paas?timeout=30s&loc=Local&parseTime=true")
	if err != nil {
		glog.Fatalf("init mysql connection err: %v", err)
	}
	db.LogMode(true)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < 20; i++ {
		glog.Info(db.Exec("INSERT INTO clusters (content )(select content from clusters)").Error)
	}
	time.Sleep(time.Second * 36000)
}

// Retrieve query user by params and pagination
func (u *User) Retrieve(params map[string]interface{}, page Pagination) (list []User, err error) {
	err = db.Where(params).Limit(page.Size).Offset(page.Size * page.Num).Find(&list).Error
	total := 0
	db.Model(u).Where(params).Count(&total)
	glog.Info(total)
	return
}

// Create create user
func (u *User) Create() error {
	return db.Create(u).Error
}

// Update update user by id
func (u *User) Update() error {
	return db.Model(u).Omit("id").Set("gorm:save_associations", false).Update(u).Error
}

// Delete delete user by id
func (u *User) Delete() error {
	return db.Delete(u).Error
}

// BatchDelete batch delete by param
func (u *User) BatchDelete(params map[string]interface{}) error {
	return db.Where(params).Delete(u).Error
}
