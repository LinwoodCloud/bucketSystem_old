package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Help\n" +
			"BUCKET_SYSTEM_REPO_URL:      The env url of the github project\n\n" +
			"validate <IssueID>           Validate if the issue is correct\n" +
			"export                       Export the static api")
	}
}
