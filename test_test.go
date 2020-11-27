package main

import testing2 "testing"

func test11() string {
	var s string
	for i := 0; i < 1000; i++ {
		s += "a"
	}
	return s
}

func BenchmarkTest(b *testing2.B) {
	for i := 0; i < b.N; i++ {
		test11()
	}
}

// go test -run= C:\GoProgram\src\tt\goNotes\test_test.go -bench=. -cpu=6 -benchtime="3s" -timeout="5s" -benchmem
// BenchmarkTest-6            32199            112591 ns/op          530341 B/op        999 allocs/op
