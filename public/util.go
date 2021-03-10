package public

import (
	"crypto/sha256"
	"fmt"
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
