to do unit tests in go:
    1. create a testing file '<filename>_test.go'
    2. import 'testing' package within this file
    3. create a function that takes a special parameter from testing (t *testing.T)
    4. within the function, make assertions by using t.Errorf
    5. run tests by using `go test` or `go test -v`