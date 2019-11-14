package lang

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

//GzipCompress 字符串压缩
func GzipCompress(str string) ([]byte, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(str)); err != nil {
		return nil, err
	}
	if err := gz.Flush(); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

//GzipDeCompress 字符串解压
func GzipDeCompress(b []byte) (string, error) {
	data, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return "", err
	}

	str, err := ioutil.ReadAll(data)
	if err != nil {
		return "", err
	}
	return string(str), nil
}
