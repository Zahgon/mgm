package crud

import (
	"github.com/kamva/mgm/v3"
)

type book struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}

func newBook(name string, pages int) *book { _ = "STUB: not implemented"; return nil }
