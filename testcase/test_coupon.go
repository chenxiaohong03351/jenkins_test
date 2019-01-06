package testcase

import (
	"reflect"
	"fmt"
)

type  Coupon struct{
	Amount int
	Desc string
	Name string
	ValidDate string
	StartTime string
	EndTime string
	CouponBatchNo string
	CouponCode string
}
func (c Coupon) Equel(compactc Coupon)  bool {
	s := reflect.ValueOf(&compactc).Elem()
	t := reflect.ValueOf(&c).Elem()
	for i :=0 ; i < s.NumField(); i++ {
		f := s.Field(i).Interface()
		fmt.Println(f)
		v := t.Field(i).Interface()
		if f != v {
			return  false
		}
	}
	return true
}