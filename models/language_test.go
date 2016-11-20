package models

import "testing"

/*
func TestLanguageInsert(t *testing.T) {
	lm := NewLanguageModel()
	lang := Language{
		Language:    "c++",
		OptionValue: "1",
		Compiler:    "g++4.9",
		OJIdFK:      1,
	}
	id, err := lm.Insert(&lang)
	if err != nil {
		t.Errorf("Failed to insert lang, %s", err)
	}
	t.Log("last insert id: ", id)

}
*/

func TestLanuageQuery(t *testing.T) {
	lm := NewLanguageModel()
	lang, err := lm.QueryById(2, nil, nil)
	if err != nil {
		t.Errorf("Failed to query, %s", err)
	}
	t.Log(lang)
}
