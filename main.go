package main

import (
	"fmt"

	"cloud.google.com/go/spanner"
	_ "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

func main() {
	_ = spanner.Statement{}
	fmt.Println("foo")
}
