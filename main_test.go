package main

import (
	"fmt"
	"testing"
)

func Test_sign(t *testing.T) {
	fmt.Println(sign("2015063000000001", "1435660288", "12345678", "apple"))
}
