package mongo

import "github.com/goincremental/dal"

type dummyItem struct {
}

func (d *dummyItem) GetID() dal.ID {
	return nil
}
