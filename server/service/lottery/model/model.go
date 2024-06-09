package model

import (
	"database/sql/driver"
	"fmt"
	"github.com/bytedance/sonic"
)

type Int64Array []int64

func (a *Int64Array) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Int64Array: %v", value)
	}
	return sonic.Unmarshal(bytes, a)
}

func (a Int64Array) Value() (driver.Value, error) {
	return sonic.Marshal(a)
}

type Result struct {
	ID      int64      `gorm:"type:bigint;primaryKey;autoIncrement:false"`
	Winners Int64Array `gorm:"type:json"`
}
