package manager

import (
	"fmt"
	"strconv"
)

type FormValue struct {
	Value any
}

func (fv FormValue) GetInt() int {
	v, _ := strconv.Atoi(fv.GetString())
	return v
}

func (fv FormValue) GetString() string {
	return fmt.Sprintf("%v", fv.Value)
}