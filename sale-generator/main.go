package main

import (
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/go-kit/kit/log"
	kitcloudwatch "github.com/go-kit/kit/metrics/cloudwatch"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewJSONLogger(os.Stderr)

	// Create the config specifying the Region for the DynamoDB table.
	// If Config.Region is not set the region must come from the shared
	// config or AWS_REGION environment variable.
	awscfg := &aws.Config{}

	// Create the session that the Cloudwatch service will use.
	sess := session.Must(session.NewSession(awscfg))

	cloudwatchAPI := cloudwatch.New(sess)
	cloudwatch := kitcloudwatch.New("cloud-native-tmdeur", cloudwatchAPI)

	requestCount := cloudwatch.NewCounter("counter")
	requestLatency := cloudwatch.NewHistogram("latency")

	var service SaleGeneratorService
	service = saleGeneratorService{}
	service = loggingMiddleware{logger, service}
	service = instrumentingMiddleware{requestCount, requestLatency, service}

	saleHandler := httptransport.NewServer(
		makeSaleEndpoint(service),
		decodeSaleRequest,
		encodeResponse,
	)

	http.Handle("/sale", saleHandler)
	http.ListenAndServe(":8080", nil)
}
