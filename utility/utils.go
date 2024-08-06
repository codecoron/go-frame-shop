package utility

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
	"math/rand"
	"time"
)

// EncryptPassword 密码加密
func EncryptPassword(password, salt string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}

func GetOrderNum() (number string) {
	rand.Seed(time.Now().UnixNano())
	number = gconv.String(time.Now().UnixNano()) + gconv.String(rand.Intn(1000))
	return
}
