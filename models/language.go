package models

import (
	"fmt"
)

type Language struct {
	LangId      int `db:"lang_id"`
	Language    string
	OptionValue string `db:"option_value"`
	Compiler    string
	OJIdFK      int `db:"oj_id_fk"`
}

type LanguageModel struct {
	Model
}

func NewLanguageModel() *LanguageModel {
	return &LanguageModel{Model{Table: "Languages"}}
}

func (this *LanguageModel) Insert(lang *Language) (int64, error) {
	if err := this.OpenDB(); err != nil {
		return 0, err
	}
	defer this.CloseDB()
	last_insert_id, err := this.InlineInsert(lang, nil, []string{"lang_id"})
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}

func (this *LanguageModel) QueryById(id int, required []string, excepts []string) (*Language, error) {
	if err := this.OpenDB(); err != nil {
		return nil, err
	}
	defer this.CloseDB()
	lang := Language{}
	str_fields, err := this.GenerateSelectSQL(lang, required, excepts)
	// fmt.Println(str_fields)
	if err != nil {
		return nil, err
	}
	if err := this.DB.Get(&lang, fmt.Sprintf("SELECT %s FROM %s WHERE lang_id=?", str_fields, this.Table), id); err != nil {
		return nil, err
	}
	return &lang, nil
}
