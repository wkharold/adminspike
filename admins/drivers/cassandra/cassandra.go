package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/wkharold/adminspike/admins"
)

type CassandraDriver struct {
	cluster *gocql.ClusterConfig
}

func init() {
	cluster := gocql.NewCluster("10.100.89.72")
	cluster.Keyspace = "adminspike"
	cluster.ProtoVersion = 3

	driver := CassandraDriver{cluster: cluster}
	admins.Register("cassandra", driver)
}

func (d CassandraDriver) Lookup(address string) (*admins.Admin, error) {
	session, err := d.cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	admin := &admins.Admin{}

	err = session.Query("select name, address from admin_from_email where address = ?", address).Scan(admin.Name, admin.Address)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (d CassandraDriver) Store(admin *admins.Admin) error {
	session, err := d.cluster.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	var name string

	if err := session.Query("select name from admin_from_email where address = ?", admin.Address).Scan(&name); err != nil {
		return fmt.Errorf("Duplicate admins not allowed")
	}

	if err := session.Query("insert into admin_from_email (address, name) values (?, ?)", admin.Address, admin.Name).Exec(); err != nil {
		return err
	}

	return nil
}
