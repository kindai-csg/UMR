package infrastructure_test

import (
	inf "github.com/kindaidensan/UMR/infrastructure"
	"testing"
)

func TestSendMail(t *testing.T) {
	handler := inf.NewMailHandler()
	err := handler.SendMail("test@kindai.ac.jp", "test", "test\r\nmail")
	if err != nil {
		t.Errorf("faild: send mail")
		t.Errorf(err.Error())
	}
}