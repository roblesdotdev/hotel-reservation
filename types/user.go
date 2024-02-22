package types

type User struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string `bson:"firstname" json:"firstname"`
	LastName  string `bson:"lastname" json:"lastname"`
}
