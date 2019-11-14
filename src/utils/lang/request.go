package lang

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
)

//Unwrap 从body中解析数据出来
func Unwrap(closer io.ReadCloser) (map[string]interface{}, error) {
	var data map[string]interface{}
	body, err := ioutil.ReadAll(closer)
	err = json.Unmarshal(body, &data)
	if err != nil {
		logrus.Error("[lang.Unwrap] parse request body error: ", err.Error())
		return map[string]interface{}{}, err
	}
	return data, nil
}

//Bind 从body中解析数据出来
func Bind(closer io.ReadCloser, value interface{}) error {
	defer closer.Close()
	body, err := ioutil.ReadAll(closer)
	err = json.Unmarshal(body, &value)
	if err != nil {
		logrus.Error("[lang.Bind] parse request body error: ", err.Error())
		return err
	}
	return nil
}
