package marshallers

import (
	"io"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MarshalObjectId(id primitive.ObjectID) graphql.Marshaler {
	data, err := id.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write(data)
	})
}

func UnmarshalObjectId(v interface{}) (primitive.ObjectID, error) {
	switch v := v.(type) {
	case string:
		return primitive.ObjectIDFromHex(v)
	default:
		panic("Value isn't an ObjectId")
	}
}
