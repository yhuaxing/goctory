package goctory

import (
	"log"
	"testing"
)

type NewTestData struct {
	Id   int
	Name string
}

func (n *NewTestData) GetName() string {
	return n.Name
}

func TestObj(t *testing.T) {

	p := NewStruct[NewTestData](
		WithAttr[NewTestData, int]("Id", 1),
		WithAttr[NewTestData, string]("Name", "arthur"),
		WithAttr[NewTestData, int]("Age", 22), //不存在的属性会被忽略
	)
	log.Printf("data:%#v ---- name:%s", p, p.GetName())
	if p.Id != 1 {
		t.Errorf("ID mismatch")
	}
	t.Log(p)
}