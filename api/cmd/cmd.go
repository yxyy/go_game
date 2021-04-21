package cmd

import "flag"

var (
	Port *string
	Config *string

)

func init()  {
	Port = flag.String("p","","port")
	Config = flag.String("c","./conf","confige")
	flag.Parse()
}
