package main

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(-1, errors.New("not exist"))
	m.EXPECT().Get(gomock.Any()).Return(-1, errors.New("not exist"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}

	if v := GetFromDB(m, "Sam"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}
