package account

import (
	"github.com/evanchen/bamboo/glog"
	"github.com/evanchen/bamboo/pto"
	"github.com/evanchen/bamboo/pto/ptohandler"
)

type Account struct {
	Urs string
	Uid int64
}

var g_accounts = make(map[string]*Account)
