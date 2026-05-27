package aggregate

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type book struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string             `json:"name" bson:"name"`
	Pages            int                `json:"pages" bson:"pages"`
	AuthorID         primitive.ObjectID `json:"author_id" bson:"author_id"`
}

func newBook(name string, pages int, authID primitive.ObjectID) *book {
	_ = "STUB: not implemented"
	return nil
}

type author struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
}

func newAuthor(name string) *author { _ = "STUB: not implemented"; return nil }
