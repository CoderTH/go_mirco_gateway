package public

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
)

//密码加盐
func GenSaltPassword(salt string, password string) string {
	s1 := sha256.New()
	s1.Write([]byte(password))
	str1 := fmt.Sprintf("%X", s1.Sum(nil))
	s2 := sha256.New()
	s2.Write([]byte(str1 + salt))
	return fmt.Sprintf("%X", s2.Sum(nil))

}

//MD5 md5加密
func MD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Obj2Json(s interface{}) string {
	bts, _ := json.Marshal(s)
	return string(bts)
}
func InStringSlice(slice []string,str string) bool{
	for  _,item:=range slice{
		if str==item{
			return true
		}
	}
	return false
}
