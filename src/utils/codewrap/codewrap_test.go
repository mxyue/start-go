package codewrap

import (
	"errors"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	errMSG := "test message"
	err := Error(Codes.NotFound, errMSG)
	code, message := Parse(err.Error())
	fmt.Println(message)
	if code != Codes.NotFound || message != errMSG {
		t.Error("parse 出错")
	}
}

func TestParse_1(t *testing.T) {
	errMSG := "test message 2"
	err := errors.New(errMSG)
	code, message := Parse(err.Error())
	if code != 0 || message != errMSG {
		t.Error("parse 出错")
	}
}
