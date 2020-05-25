package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"sync"
	"time"
)

func GetMsgId()int64{
	return time.Now().UnixNano()
}
func Md5V3(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func AppReturn(code,status int,action,msg string,msgtoken interface{})map[string]interface{}{
	lock := new(sync.RWMutex)
	lock.Lock()
	defer lock.Unlock()
	appResp := make(map[string]interface{})
	appResp["code"] = code
	appResp["msg"]= msg
	appResp["isOnline"] = status
	appResp["action"] = action
	appResp["msgToken"] = msgtoken
	return appResp
}