# go-featurebase

go-featurebase is a Go client for [Featurebase](https://www.featurebase.com/).


### Installation

```
go get github.com/rachithrr/go-featurebase
```

### Usage

```go
package main

import (
	"fmt"

	fb "github.com/rachithrr/go-featurebase"
)

func main() {
	client := fb.NewClient(&fb.Options{
        # if using local instance
		Host:     "localhost",
		Port:     "10101",

        # if using cloud instance
		QueryURL: "https://query.featurebase.com/v2/databases/{database_id}/query/sql", 
		APIKey:   "{api_key}",
	})

	// Create table
	resp, err := client.Query("CREATE TABLE test (_id INT, name STRING)")
	handleError(err, resp)

	// Insert row  
	resp, err = client.Query("INSERT INTO test VALUES (1, 'test')")
	handleError(err, resp)

	// Select rows
	resp, err = client.Query("SELECT * FROM test")
	handleError(err, resp)

	fmt.Printf("%v \n%v \n%v\n", resp.Schema, resp.Data, resp.ExecutionTime)
}

func handleError(err error, resp *fb.Response) {
	if err != nil {
		fmt.Println(err)
		return
	}
	
	if resp.Error != "" {
		fmt.Println(resp.Error)
		return
	}
}
```
