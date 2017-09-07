package db

import mgo "gopkg.in/mgo.v2"

// Session is our database session
var _session *mgo.Session

// Dial dials for dialing mongo server
func Dial() error {
	var err error
	_session, err = mgo.Dial("localhost")
	if err != nil {
		return err
	}
	// SetMode is an important feature when making a connection to mongodb, monotonic is my prefered middle ground setting.
	_session.SetMode(mgo.Monotonic, true)
	return nil
}

// Session func will allow you to clone this private session to all aready of our application.
func Session() *mgo.Session {
	return _session.Clone()
}
