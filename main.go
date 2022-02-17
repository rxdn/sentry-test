package main

import (
	"flag"
	"fmt"
	"github.com/getsentry/sentry-go"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var logger = log.Default()

func main() {
	dsn := flag.String("dsn", "", "Sentry DSN string")
	flag.Parse()

	if *dsn == "" {
		log.Fatalln("No DSN provided, exiting") // exits with status 1
	}

	statusCode := testConnection(*dsn)
	log.Printf("HTTP connection established with status code %d\n", statusCode)

	testSentry(*dsn)
	log.Println("Test complete, please check your Sentry instance to verify that the error was logged")
}

// exit if failure
// returns http status code
func testConnection(dsn string) int {
	client := &http.Client{
		Timeout: time.Second * 3,
	}

	res, err := client.Get(dsn)
	if err != nil {
		log.Fatalf("Connection test failed: %s\n", err.Error())
	}

	return res.StatusCode
}

func testSentry(dsn string) {
	transport := sentry.NewHTTPSyncTransport()

	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
		Debug: true,
		Transport: transport,
	})

	if err != nil {
		log.Fatalf("Error initialising sentry client: %s\n", err.Error())
	}

	rand.Seed(time.Now().UnixNano())
	errorString := fmt.Errorf("test-%d", rand.Intn(100_000))

	log.Printf("Attempting to log error with message %s\n", errorString)
	sentry.CaptureException(errorString)
	log.Println("Error successfully logged")

	sentry.Flush(time.Second * 10) // Shouldn't need due to sync transport
}
