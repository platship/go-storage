package facade

import (
	"errors"

	"github.com/platship/go-storage"
	"github.com/platship/go-storage/drivers"
)

type Storage struct {
	Active string  `json:"active"`
	Driver *Driver `json:"driver"`
}

func NewStorage() *Storage {
	return &Storage{
		Active: "local",
		Driver: NewDriver(),
	}
}

func (s *Storage) ActiveDriver() (storage.Storage, error) {
	return s.Driver.Get(s.Active)
}

func (s *Storage) CloseAll() {
	for _, item := range s.Driver.Items() {
		_ = item.Close()
	}
}

type Driver struct {
	Local *drivers.Local `json:"local"`
	Oss   *drivers.Oss   `json:"oss"`
}

func NewDriver() *Driver {
	return &Driver{
		Local: &drivers.Local{},
		Oss:   &drivers.Oss{},
	}
}

func (d *Driver) Items() []storage.Storage {
	return []storage.Storage{
		d.Local,
		d.Oss,
	}
}

func (d *Driver) Get(id string) (storage.Storage, error) {
	if id == "" {
		return nil, errors.New("id undefined")
	}
	switch id {
	case "local":
		return d.Local, nil
	case "oss":
		return d.Oss, nil
	}
	return nil, errors.New("driver not found")
}
