// Package stringutil contains utility for working with string
package stringutil

func Reverse(s string) string {
	return reverseTwo(s)
}
/*
go build
		go build reverse.go reverseTwo.go
		won't produce an output file
go install
		will place the package inside the pkg directory.


*/


