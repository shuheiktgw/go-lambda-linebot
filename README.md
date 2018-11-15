go-lambda-linebot
====
[![GitHub release](http://img.shields.io/github/release/shuheiktgw/go-lambda-linebot.svg?style=flat-square)](https://github.com/shuheiktgw/go-lambda-linebot/releases/latest)
[![CircleCI](https://circleci.com/gh/shuheiktgw/go-lambda-linebot.svg?style=svg)](https://circleci.com/gh/shuheiktgw/go-lambda-linebot)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

`go-lambda-linebot` provides utility functions to parse AWS Lambda Events and turn them into `linebot.Event`.  

## Example

### `ParseAPIGatewayProxyRequest`
`ParseAPIGatewayProxyRequest` parses `events.APIGatewayProxyRequest` and turns it into an array of `linebot.Event`. The implementation of `ParseAPIGatewayProxyRequest` follows [webhook.go](https://github.com/line/line-bot-sdk-go/blob/master/linebot/webhook.go)  in [line/line-bot-sdk-go](https://github.com/line/line-bot-sdk-go).

``` go
package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shuheiktgw/go-lambda-linebot/parser"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	lineEvents, err := parser.ParseAPIGatewayProxyRequest(os.Getenv("CHANNEL_SECRET"), &request)
	
	# Do something useful...
}

func main() {
	lambda.Start(handler)
}
```

### `ParseSNSEvent`
`ParseSNSEvent` parses `events.SNSEvent` and turns it inti `linebot.Event`. `ParseSNSEvent` assumes that you provide a single `linebot.Event` in the Message field of SNS.

``` go
package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shuheiktgw/go-lambda-linebot/parser"
)

func handler(request event events.SNSEvent) error {
	lineEvent, err := parser.ParseSNSEvent(&request)
	
	# Do something useful...
}

func main() {
	lambda.Start(handler)
}
```

## Install

``` bash
$ go get github.com/shuheiktgw/go-lambda-linebot/parser
```

## Author
[Shuhei Kitagawa](https://github.com/shuheiktgw)