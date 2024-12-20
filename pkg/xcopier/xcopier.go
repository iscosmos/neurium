package xcopier

import (
	"github.com/jinzhu/copier"
	"time"
)

var (
	CopierOption = copier.Option{
		Converters: []copier.TypeConverter{
			copier.TypeConverter{
				SrcType: time.Time{},
				DstType: int64(0),
				Fn: func(src interface{}) (dst interface{}, err error) {
					if t, ok := src.(time.Time); ok {
						return t.Unix(), nil
					}
					return nil, err
				},
			},
		},
	}
)
