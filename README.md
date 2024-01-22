# go-test-catalog
Creates test catalog from a given directory

This code helps create test catalog for go tests from the test docs/comments

While writing go tests, write the docs or comments in the below format and run the code (refer How to run section). This will generate an HTML file where you will be able to view the catalog for all the go tests. It also generates a json file with details about the tests. You can use this json file to build your own dashboards  

```
/*
# Description
This is a sample description of what the test does

# Scenario
Write details about your test here

# Metadata
TestType: <e2e or upgrade or rollback or ...>
Component: <what component does this test belong to?>

*/
func Test_Something(t *testing.T) {
    ....
}
```

# How to run
```go run main.go```

OR

```go run main.go /path/to/directory```
