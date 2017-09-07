package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"

	controllers "github.com/xDarkicex/urlShortener/Controllers"
	"github.com/xDarkicex/urlShortener/Datastore"
)

var session *mgo.Session

func init() {
	// Dial mongo Datastore
	err := db.Dial()
	if err != nil {
		log.Fatal(err)
	}
	session = db.Session()
}

func main() {
	fmt.Println(controllers.GET("https://r01.ce/1D"))
	// data := db.Set(3213123124243545463, "http://www.animefreaks.com/longurl/pic/anime/cute001.jpg").Set(2, "https://a2plvcpnl13751.prod.iad2.secureserver.net").Set(3, "https://google.com")
	// fmt.Printf("%s\n%s\n%s\n", data.Get(3213123124243545463), data.Get(2), data.Get(3))
}
