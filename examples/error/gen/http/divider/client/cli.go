// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// divider HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/goa/examples/error/design -o
// $(GOPATH)/src/goa.design/goa/examples/error

package client

import (
	"fmt"
	"strconv"

	dividersvc "goa.design/goa/examples/error/gen/divider"
)

// BuildIntegerDividePayload builds the payload for the divider integer_divide
// endpoint from CLI flags.
func BuildIntegerDividePayload(dividerIntegerDivideA string, dividerIntegerDivideB string) (*dividersvc.IntOperands, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(dividerIntegerDivideA, 10, 64)
		a = int(v)
		if err != nil {
			err = fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(dividerIntegerDivideB, 10, 64)
		b = int(v)
		if err != nil {
			err = fmt.Errorf("invalid value for b, must be INT")
		}
	}
	if err != nil {
		return nil, err
	}
	payload := &dividersvc.IntOperands{
		A: a,
		B: b,
	}
	return payload, nil
}

// BuildDividePayload builds the payload for the divider divide endpoint from
// CLI flags.
func BuildDividePayload(dividerDivideA string, dividerDivideB string) (*dividersvc.FloatOperands, error) {
	var err error
	var a float64
	{
		a, err = strconv.ParseFloat(dividerDivideA, 64)
		if err != nil {
			err = fmt.Errorf("invalid value for a, must be FLOAT64")
		}
	}
	var b float64
	{
		b, err = strconv.ParseFloat(dividerDivideB, 64)
		if err != nil {
			err = fmt.Errorf("invalid value for b, must be FLOAT64")
		}
	}
	if err != nil {
		return nil, err
	}
	payload := &dividersvc.FloatOperands{
		A: a,
		B: b,
	}
	return payload, nil
}