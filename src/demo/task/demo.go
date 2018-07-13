package task

import (
	_ "github.com/go-sql-driver/mysql"
	"demo/config"
	"gopkg.in/mgo.v2"
	"time"
)

type Obj struct {
	Id       string         `bson:"_id"`
	Counts   map[string]int `bson:"counts"`
	EditTime time.Time      `bson:"edit_time"`
	ADDTime  time.Time      `bson:"add_time"`
}

func syncStart() {
	// mysql
	//uri := config.Mysql.User + ":" + config.Mysql.Pwd + "@tcp(" + config.Mysql.Host + ":" + config.Mysql.Port + ")/" + config.Mysql.Database
	//
	//db, err := sql.Open("mysql", uri)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
}

// mongo 初始化
func syncInitMongo() *mgo.Session {
	session, err := mgo.Dial(config.Mongo.Uri)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Eventual, true)
	session.SetSocketTimeout(time.Second * 10)
	session.SetSyncTimeout(time.Second * 10)
	return session
}

func SyncDataRun() {
	syncStart()
}
