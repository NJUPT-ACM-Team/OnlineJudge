package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Language struct {
	LangId      int64 `db:"lang_id"`
	Language    string
	OptionValue string `db:"option_value"`
	Compiler    string
	OJIdFK      int64 `db:"oj_id_fk"`
}

type LanguageModel struct {
	Model
}

func NewLanguageModel() *LanguageModel {
	return &LanguageModel{Model{Table: "Languages"}}
}

func (this *LanguageModel) Insert(tx *sqlx.Tx, lang *Language) (int64, error) {
	last_insert_id, err := this.InlineInsert(tx, lang, nil, []string{"lang_id"})
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}

func (this *LanguageModel) QueryById(tx *sqlx.Tx, id int64, required []string, excepts []string) (*Language, error) {
	lang := Language{}
	str_fields, err := this.GenerateSelectSQL(lang, required, excepts)
	// fmt.Println(str_fields)
	if err != nil {
		return nil, err
	}
	if err := tx.Get(&lang, fmt.Sprintf("SELECT %s FROM %s WHERE lang_id=?", str_fields, this.Table), id); err != nil {
		return nil, err
	}
	return &lang, nil
}
