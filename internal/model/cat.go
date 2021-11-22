package model

import "encoding/json"

type Cat struct {
	Id   string `bson:"_id,omitempty" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type Cats []*Cat

type User struct {
	Id       int    `bson:"_id,omitempty" json:"id,omitempty" db:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func FromJson(val []byte) Cat {
	cat := Cat{}
	err := json.Unmarshal(val, &cat)
	if err != nil {
		panic(err)
	}
	return cat
}
func ToJson(v *Cat) []byte {
	json, err := json.Marshal(&v)
	if err != nil {
		panic(err)
	}
	return json
}
