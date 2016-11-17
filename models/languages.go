package models

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

func (this *LanguageModel) Insert(lang *Language) (int, error) {
	if err := this.OpenDB(); err != nil {
		return 0, err
	}
	defer this.CloseDB()
	last_insert_id, err := this.InlineInsert(lang, []string{"lang_id"})
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}
