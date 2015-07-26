package admins

import "fmt"

type Admins interface {
	FindByEmail(address string) (*Admin, error)
	Save(admin Admin) error
}

type Driver interface {
	Lookup(address string) (*Admin, error)
	Store(admin *Admin) error
}

type Admin struct {
	Name    string
	Address string
}

type Collection struct {
	driver Driver
}

var (
	drivers map[string]Driver = map[string]Driver{}
)

func Register(name string, driver Driver) error {
	if driver == nil {
		return fmt.Errorf("Can't register a nil driver")
	}

	if _, dup := drivers[name]; dup {
		return fmt.Errorf("Driver already registered %s", name)
	}

	drivers[name] = driver
	return nil
}

func Using(driver string) (*Collection, error) {
	if d, ok := drivers[driver]; ok {
		return &Collection{driver: d}, nil
	}
	return nil, fmt.Errorf("No such driver %s", driver)
}

func (c Collection) FindByEmail(address string) (*Admin, error) {
	admin, err := c.driver.Lookup(address)
	if err != nil {
		return nil, fmt.Errorf("Can't find %s: %+v", address, err)
	}
	return admin, nil
}

func (c *Collection) Save(admin *Admin) error {
	return c.driver.Store(admin)
}
