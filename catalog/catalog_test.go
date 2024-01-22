package catalog

import (
	"fmt"
	"testing"
)

/*
# Description
This is a sample description of what the test does

# Scenario
Given I have some conditions
When I do task A
Then I expect task to be completed successfully
When I disable task A
And I enable task B
Then I expect task A to be disabled and task B to be executed successfully

# Metadata
TestType: e2e
Component: Tasks

Some other comments
*/
func Test_Catalog(t *testing.T) {
	fmt.Println("executing test catalog")
}
