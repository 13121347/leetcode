package Option

import (
	"fmt"
	"testing"
)

func Test_loadOptions(t *testing.T) {
	//testOptions := Options{1,1,1}

	op1 := func(opts *Options) { opts.a = 2 }
	op2 := func(opts *Options) { opts.b = 2 }
	op3 := func(opts *Options) { opts.c = 2 }

	testOpts := loadOptions(op1, op2, op3)
	fmt.Println(testOpts)

}
