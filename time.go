// Copyright © 2021 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package convert

import (
	"errors"
	"time"

	"github.com/kovacou/go-types"
)

// TimeE cast input to time.Time and return the associated error.
func TimeE(v interface{}) (t time.Time, err error) {
	switch val := v.(type) {
	case time.Time:
		t = val
	case *time.Time:
		if val != nil {
			t = *val
		}
	case types.Date:
		t = val.Time
	case *types.Date:
		if val != nil {
			t = val.Time
		}
	case types.DateTime:
		t = val.Time
	case *types.DateTime:
		if val != nil {
			t = val.Time
		}
	case types.DateYearMonth:
		t = val.Time
	case *types.DateYearMonth:
		if val != nil {
			t = val.Time
		}
	default:
		err = errors.New("could not convert to time.Time")
	}

	return
}

// Time cast input to time.Time and return default value if error.
func Time(v interface{}) time.Time {
	t, _ := TimeE(v)
	return t
}

// DateTimeE cast input to types.DateTime and return the associated error.
func DateTimeE(v interface{}) (dt types.DateTime, err error) {
	switch val := v.(type) {
	case time.Time:
		dt = types.NewDateTime(val)
	case *time.Time:
		if val != nil {
			dt = types.NewDateTime(*val)
		}
	case types.Date:
		dt = types.NewDateTime(val.Time)
	case *types.Date:
		if val != nil {
			dt = types.NewDateTime(val.Time)
		}
	case types.DateTime:
		dt = val
	case *types.DateTime:
		if val != nil {
			dt = types.NewDateTime(val.Time)
		}
	case types.DateYearMonth:
		dt = types.NewDateTime(val.Time)
	case *types.DateYearMonth:
		if val != nil {
			dt = types.NewDateTime(val.Time)
		}
	default:
		err = errors.New("could not convert to types.DateTime")
	}

	return
}

// DateTime cast input to types.DateTime and return default value if error.
func DateTime(v interface{}) types.DateTime {
	dt, _ := DateTimeE(v)
	return dt
}

// DateE cast input to types.Date and return the associated error.
func DateE(v interface{}) (d types.Date, err error) {
	switch val := v.(type) {
	case time.Time:
		d = types.NewDate(val)
	case *time.Time:
		if val != nil {
			d = types.NewDate(*val)
		}
	case types.Date:
		d = val
	case *types.Date:
		if val != nil {
			d = types.NewDate(val.Time)
		}
	case types.DateTime:
		d = types.NewDate(val.Time)
	case *types.DateTime:
		if val != nil {
			d = types.NewDate(val.Time)
		}
	case types.DateYearMonth:
		d = types.NewDate(val.Time)
	case *types.DateYearMonth:
		if val != nil {
			d = types.NewDate(val.Time)
		}
	default:
		err = errors.New("could not convert to types.DateTime")
	}

	return
}

// Date cast input to types.Date and return default value if error.
func Date(v interface{}) types.Date {
	d, _ := DateE(v)
	return d
}
