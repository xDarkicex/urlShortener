package controllers

import (
	"strings"

	"github.com/xDarkicex/EncodeBase64"
	models "github.com/xDarkicex/urlShortener/Models"
)

var (
	I     models.Index
	url   *models.URL
	count = I.GetLength()
)

func Shorten(original string) interface{} {
	count = I.Increment()
	url.PUT(count, original)
	return "https://r01.ce/" + encoder.IntToBase64(count)
}

func GET(u string) models.URL {
	return url.FindByID(strings.Split(u, "/")[3])
}

// var entry = make(map[interface{}]string)

// func (d Datastore) Set(id int, original string) Datastore {
// 	encoded := encoder.Base64(id)
// 	var url = URL{ID: id, Original: original, New: "https://gro.de/" + encoded}
// 	json, _ := json.Marshal(url)
// 	fmt.Println(string(json))
// 	entry[url.ID] = string(json)
// 	return d
// }

// func exists(k string) bool {
// 	if len(entry[k]) > 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func (d Datastore) Update(hash string, newURL string) Datastore {
// 	entry[hash] = newURL
// 	return d
// }

// func (d Datastore) Shorten(k string) string {
// 	return ""
// }

// func (d Datastore) Get(k int) string {
// 	var url URL
// 	if err := json.Unmarshal([]byte(entry[k]), &url); err != nil {
// 		fmt.Println("error unmarshalling", err)
// 	}

// 	return url.New
// }
