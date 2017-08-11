package mongo

import (
	"github.com/eyecuelab/kit/log"
	"github.com/spf13/cobra"
	"gopkg.in/mgo.v2"
)

var (
	MDb *mgo.Session
)

func init() {
	cobra.OnInitialize(connect)
}

func connect() {
	var err error
	MDb, err = mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}

	MDb.SetMode(mgo.Monotonic, true)
}
