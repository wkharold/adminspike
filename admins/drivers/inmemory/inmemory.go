package inmemory

import (
	"fmt"

	"github.com/wkharold/adminspike/admins"
)

type InmemoryDriver struct {
	collection map[string]*admins.Admin
}

func init() {
	driver := InmemoryDriver{collection: map[string]*admins.Admin{}}
	admins.Register("inmemory", driver)
}

func (d InmemoryDriver) Lookup(address string) (*admins.Admin, error) {
	if admin, ok := d.collection[address]; ok {
		return admin, nil
	}
	return nil, fmt.Errorf("No such admin %s", address)
}

func (d InmemoryDriver) Store(admin *admins.Admin) error {
	if admin, ok := d.collection[admin.Address]; ok {
		return fmt.Errorf("No duplicates allowed %+v", admin)
	}
	d.collection[admin.Address] = admin
	return nil
}
