# go-test-catalog
Creates test catalog from a given directory

This code helps create test catalog for go tests from the test docs/comments

While writing go tests, write the docs or comments in the below format and run the code (refer How to run section). 
 - This will generate an HTML file where you will be able to view the catalog for all the go tests. 
 - It also generates a json file with details about the tests. You can use this json file to build your own dashboards 
 - It also hosts the test catalog on your localhost on port 8080. Access at http://localhost:8080

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
A screen grab from the generated HTML file / hosted service
![test_catalog.png](templates%2Ftest_catalog.png)

Corresponding json data
```
[
   {
      "metadata": {
         "dir": "catalog",
         "package": "catalog",
         "file": "catalog_test.go"
      },
      "testName": "Test_Catalog",
      "description": "\nThis is a sample description of what the test does\n\n",
      "scenario": "\nGiven I have some conditions\nWhen I do task A\nThen I expect task to be completed successfully\nWhen I disable task A\nAnd I enable task B\nThen I expect task A to be disabled and task B to be executed successfully\n\n",
      "types": [
         "e2e"
      ],
      "components": [
         "tasks"
      ]
   }
]
```

# How to run
```go run main.go```

OR

```go run main.go /path/to/directory```
