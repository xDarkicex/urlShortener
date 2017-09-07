package models

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/xDarkicex/EncodeBase64"
	"github.com/xDarkicex/urlShortener/Datastore"
)

type URL struct {
	ID       int
	Original string
}

type Index struct {
	ID int
}

func (i *Index) GetLength() int {
	i.ID = 102
	return i.ID
}

func (i *Index) Increment() int {
	i.ID++
	return i.ID
}

func (u *URL) FindByID(id string) URL {
	session := db.Session()
	defer session.Close()
	var url URL
	ID := encoder.Base64ToInt(id)
	if err := session.DB("shortener").C("urls").Find(bson.M{"id": ID}).One(&url); err != nil {
		fmt.Println(err)
	}
	return url
}

func (u *URL) PUT(id int, orginal string) {
	session := db.Session()
	defer session.Close()
	c := session.DB("shortener").C("urls")
	// Insert url data into mongoDB
	err := c.Insert(&URL{
		ID:       id,
		Original: orginal,
	})
	if err != nil {
		fmt.Println(err)
	}
}
