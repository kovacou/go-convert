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
func TimeE(v any) (t time.Time, err error) {
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
func Time(v any) time.Time {
	t, _ := TimeE(v)
	return t
}

// DateTimeE cast input to types.DateTime and return the associated error.
func DateTimeE(v any) (dt types.DateTime, err error) {
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
func DateTime(v any) types.DateTime {
	dt, _ := DateTimeE(v)
	return dt
}

// DateE cast input to types.Date and return the associated error.
func DateE(v any) (d types.Date, err error) {
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
func Date(v any) types.Date {
	d, _ := DateE(v)
	return d
}

// DateYearMonthE cast input to types.DateYearMonth and return the associated error.
func DateYearMonthE(v any) (d types.DateYearMonth, err error) {
	switch val := v.(type) {
	case time.Time:
		d = types.NewDateYearMonth(val)
	case *time.Time:
		if val != nil {
			d = types.NewDateYearMonth(*val)
		}
	case types.Date:
		d = types.NewDateYearMonth(val.Time)
	case *types.Date:
		if val != nil {
			d = types.NewDateYearMonth(val.Time)
		}
	case types.DateTime:
		d = types.NewDateYearMonth(val.Time)
	case *types.DateTime:
		if val != nil {
			d = types.NewDateYearMonth(val.Time)
		}
	case types.DateYearMonth:
		d = val
	case *types.DateYearMonth:
		if val != nil {
			d = types.NewDateYearMonth(val.Time)
		}
	default:
		err = errors.New("could not convert to types.DateTime")
	}

	return
}

// DateYearMonth cast input to types.DateYearMonth and return default value if error.
func DateYearMonth(v any) types.DateYearMonth {
	d, _ := DateYearMonthE(v)
	return d
}
