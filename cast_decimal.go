package casts

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func MustDecimal(v any) decimal.Decimal {
	if v == nil {
		return decimal.Decimal{}
	}

	d, err := decimal.NewFromString(fmt.Sprintf("%v", v))
	if err != nil {
		fmt.Println(err)
		return decimal.Decimal{}
	}
	return d
}

func Decimal(v any) (d decimal.Decimal, err error) {
	if v == nil {
		err = fmt.Errorf("cast to decimal: target is nil")
		return
	}

	d, err = decimal.NewFromString(fmt.Sprintf("%v", v))
	if err != nil {
		return
	}
	return
}

func DecimalDefault(v any, dv decimal.Decimal) decimal.Decimal {
	if v == nil {
		return dv
	}

	d, err := decimal.NewFromString(fmt.Sprintf("%v", v))
	if err != nil {
		return dv
	}
	return d
}
