package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8089", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		panic(err)
	}
}
