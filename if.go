// Copyright 2017 The Kirinlabs. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"strconv"
)

type IfThen struct {
	c bool
	t interface{}
	e interface{}
	v interface{}
}

func (i *IfThen) If(c bool) *IfThen {
	i.c = c
	return i
}

func (i *IfThen) Then(v interface{}) *IfThen {
	i.t = v
	return i
}

func (i *IfThen) Else(v interface{}) *IfThen {
	i.e = v
	return i
}

func (i *IfThen) String() string {
	if i.c {
		i.v = i.t
	} else {
		i.v = i.e
	}
	return fmt.Sprintf("%v", i.v)
}

func (i *IfThen) Int() (int64, error) {
	if i.c {
		i.v = i.t
	} else {
		i.v = i.e
	}
	switch i.v.(type) {
	case int64:
		return i.v.(int64), nil
	case int32:
		return int64(i.v.(int32)), nil
	case int:
		return int64(i.v.(int)), nil
	case int8:
		return int64(i.v.(int8)), nil
	case int16:
		return int64(i.v.(int16)), nil
	case float32:
		return int64(i.v.(float32)), nil
	case float64:
		return int64(i.v.(float64)), nil
	case []byte:
		i, err := strconv.ParseInt(string(i.v.([]byte)), 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	case string:
		i, err := strconv.ParseInt(i.v.(string), 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
	return 0, fmt.Errorf("Unsupported type: %v", i.v)
}

func (i *IfThen) Float() (float64, error) {
	if i.c {
		i.v = i.t
	} else {
		i.v = i.e
	}
	switch i.v.(type) {
	case float32:
		return float64(i.v.(float32)), nil
	case float64:
		return i.v.(float64), nil
	case string:
		i, err := strconv.ParseFloat(i.v.(string), 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	case []byte:
		i, err := strconv.ParseFloat(string(i.v.([]byte)), 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
	return 0, fmt.Errorf("Unsupported type: %v", i.v)
}
