package models

import (
	"testing"
)

type MyModel struct {
	Model
}

func (this *MyModel) TestFunc(t *testing.T) error {
	t.Log("In TestFunc, try to open/close DB")
	if err := this.OpenDB(); err != nil {
		t.Errorf("Failed to open database: %s", err)
		return err
	}
	defer this.CloseDB()
	return nil
}

func TestModel(t *testing.T) {
	t.Log("TestModel")
	mymodel := MyModel{}
	mymodel.TestFunc(t)
}

func (this *MyModel) TestFuncGetAllFields(t *testing.T) error {
	if err := this.OpenDB(); err != nil {
		t.Errorf("Failed to open database: %s", err)
		return err
	}
	defer this.CloseDB()
	t.Log("Get all fields from a struct")
	type Foo struct {
		RunId      int `db:"run_id"`
		Status     string
		StatusCode string `db:"status_code"`
	}
	f := Foo{1, "a", "b"}
	sql, err := this.GenerateInsertSQL(f, "foo", []string{"run_id"})
	t.Log(sql)
	if err != nil {
		t.Errorf("%s", err)
	}
	return err

}
func TestGetAllFields(t *testing.T) {
	mymodel := MyModel{}
	mymodel.TestFuncGetAllFields(t)
}
