package main

//
// const str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_"
//
// var by = []byte(str)
//
// const n = 10
//
// func BenchmarkBytes(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		var chars strings.Builder
// 		chars.Grow(n)
// 		for i := 0; i < n; i++ {
// 			chars.WriteByte(str[rand.Intn(64)]) //nolint:gosec
// 		}
// 	}
// }
//
// func BenchmarkString(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		var chars strings.Builder
// 		chars.Grow(n)
// 		for i := 0; i < n; i++ {
// 			chars.WriteByte(by[rand.Intn(64)]) //nolint:gosec
// 		}
// 	}
// }
//

/*
goos: linux
goarch: amd64
pkg: ok
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkBytes-6         4975813               210.0 ns/op            16 B/op          1 allocs/op
BenchmarkString-6        6414460               188.7 ns/op            16 B/op          1 allocs/op
PASS
ok      ok      3.734s
*/
