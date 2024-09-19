package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://google.com/", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer func(res *http.Response) {
		if err = res.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}(res)

	if _, err = io.CopyBuffer(os.Stdout, res.Body, nil); err != nil {
		panic(err)
	}
}
