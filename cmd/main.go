package main

import (
	"github.com/wkharold/adminspike/admins"
	_ "github.com/wkharold/adminspike/admins/drivers/cassandra"
	_ "github.com/wkharold/adminspike/admins/drivers/inmemory"
)

func main() {
	im, _ := admins.Using("cassandra")
	im.FindByEmail("dobbs@sierramadre.io")
}
