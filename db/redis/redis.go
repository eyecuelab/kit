package redis

import (
	"github.com/eyecuelab/kit/config"
	"github.com/eyecuelab/kit/log"
	r "github.com/go-redis/redis"
	"github.com/spf13/cobra"
)

var Client *r.Client

func init() {
	cobra.OnInitialize(connectDB)
}

func connectDB() {
	url := config.RequiredString("redis_url")
	opts, err := r.ParseURL(url)
	log.Check(err)
	Client = r.NewClient(opts)
}
