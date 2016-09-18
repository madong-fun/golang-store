package file

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

/**
 * 根据文件内容做md5加密
 */
func Encrypt(fileName string) (result string) {

	file, err := os.Open(fileName)

	if err != nil {
		log.Panic(err)
	}

	info, err := file.Stat()
	if err != nil {
		log.Panic(err)
	}

	len := info.Size()
	var b []byte = make([]byte, len)
	reader := bufio.NewReader(file)
	n, err := reader.Read(b)
	if err != nil {
		fmt.Println(err)
	}

	if n > 0 {
		h := md5.New()
		h.Write(b)
		p := h.Sum(nil)
		result = hex.EncodeToString(p)
	}

	return

}
