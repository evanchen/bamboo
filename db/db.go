package db

import (
	"github.com/evanchen/bamboo/etc"
	"github.com/evanchen/bamboo/glog"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"log"
)

var g_session *mgo.Session
var g_db *mgo.Database
var logger = glog.New("log/db.log")

func Init() {
	ret, dbAddr := etc.GetConfigStr("db_addr")
	if !ret {
		log.Fatalf("[db.Init] failed to get db_addr config")
	}
	ret, dbname := etc.GetConfigStr("db_name")
	if !ret {
		log.Fatalf("[db.Init] failed to get db_name config")
	}
	ret, user := etc.GetConfigStr("db_user")
	if !ret {
		log.Fatalf("[db.Init] failed to get db_user config")
	}
	ret, pwd := etc.GetConfigStr("db_pwd")
	if !ret {
		log.Fatalf("[db.Init] failed to get db_pwd config")
	}
	dInfo := &mgo.DialInfo{
		Addrs:     []string{dbAddr},
		Timeout:   0, //block
		Database:  dbname,
		Username:  user,
		Password:  pwd,
		Source:    dbname,
		PoolLimit: 4096,
		Direct:    false,
	}
	g_session, err := mgo.DialWithInfo(dInfo)
	if err != nil {
		log.Fatalf("[db.Init] failed to dial db: %v\n, error :%s", dInfo, err.Error())
	}
	g_session.SetSafe(&mgo.Safe{})
	g_db = g_session.DB(dbname)
	logger.Info("[db.Init] connected to db server: %v\n", dInfo)
}

func NewSession() *mgo.Session {
	return g_session.Copy()
}
