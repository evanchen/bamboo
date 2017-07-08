package account

import (
	"github.com/evanchen/bamboo/glog"
	"github.com/evanchen/bamboo/pto"
	"github.com/evanchen/bamboo/pto/ptohandler"
	"sync"
)

type Account struct {
	Urs string
	Uid int64
}

var g_accounts sync.Map

func SetAccount(acc *Account) {
	g_accounts.Store(acc.Uid,acc)
}

func GetAccount(uid int64) *Account {
	acc,_ := g_accounts.Load(uid)
	return acc
}

