package app

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"gopkg.in/oauth2.v3/store"
	"net/http"
)

var (
	Store *sessions.FilesystemStore
)

func Init() error {
	Store = sessions.NewFilesystemStore("", []byte("something-very-secret"))
	gob.Register(map[string]interface{}{})
	http.HandleFunc("/todo", viewRecord)
	return nil
}

func viewRecord(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(store.ClientStore)
	if err := datastore.Get(c, key, record); err != nil {
		http.Error(w, err.Error(), 500)
	}
	if err := datastore.Get(c, key, record); err != nil {
		return
	}
	return
}
