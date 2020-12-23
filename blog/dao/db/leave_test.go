package db

import (
	"github.com/zhangliying0113/gin/blog/model"
	"testing"
)

func init() {
	// parseTime=true 将 mysql 的时间类型自动解析为 go 结构体中的时间类型，不加报错
	dns := "root:123456@tcp(localhost:3306)/blog?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestGetLeaveList(t *testing.T) {
	leaveList, err := GetLeaveList()
	if err != nil {
		t.Errorf("get leave list failed, err:%v\n", err)
		return
	}
	for _, v := range leaveList{
		t.Logf("category: %#v\n", v)
	}
}

func TestInsertLeave(t *testing.T) {
	leave := &model.Leave{}
	leave.Content = "小白"
	leave.Email = "17111232@qq.com"
	leave.Content = "this is an question"
	err := InsertLeave(leave)
	if err != nil {
		t.Errorf("insert leave failed, err:%v\n", err)
		return
	}
	t.Logf("insert leave success")
}
