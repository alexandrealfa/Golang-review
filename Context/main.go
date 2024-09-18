package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Minute)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "get", "http://google.com/", nil)

	if err != nil {
		fmt.Println(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	io.CopyBuffer(os.Stdout, res.Body, nil)
}
