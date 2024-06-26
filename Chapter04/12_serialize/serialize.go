package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {

	eur, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}
	t := time.Date(2017, 11, 20, 11, 20, 10, 0, eur)

	// json.Marshaler interface
	b, err := t.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println("Serialized as RFC 3339:", string(b))
	t2 := time.Time{}
	err = t2.UnmarshalJSON(b)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deserialized from RFC 3339:", t2)

	// Serialize as epoch
	epoch := t.Unix()
	fmt.Println("Serialized as Epoch:", epoch)

	// Deserialize epoch
	jsonStr := fmt.Sprintf("{ \"created\":%d }", epoch)
	data := struct {
		Created int64 `json:"created"`
	}{}

	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(err)
	}

	deserialized := time.Unix(data.Created, 0)
	fmt.Println("Deserialized from Epoch:", deserialized)

}
