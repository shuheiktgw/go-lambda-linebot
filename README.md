go-lambda-linebot
====
[![GitHub release](http://img.shields.io/github/release/shuheiktgw/go-lambda-linebot.svg?style=flat-square)](https://github.com/shuheiktgw/go-lambda-linebot/releases/latest)
[![CircleCI](https://circleci.com/gh/shuheiktgw/go-lambda-linebot.svg?style=svg)](https://circleci.com/gh/shuheiktgw/go-lambda-linebot)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

`go-lambda-linebot` provides a functionality to parse `event.APIGatewayProxyRequest` and turn it into an array of `linebot.Event`. This is useful especially when you create LINE bot using their [Messaging API](https://developers.line.me/ja/services/messaging-api/) with AWS Lambda and API Gateway. 

## Example

``` go
package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shuheiktgw/go-lambda-linebot/parser"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	lineEvents, err := parser.ParseRequest(os.Getenv("CHANNEL_SECRET"), &request)
	
	# Do something useful...
}

func main() {
	lambda.Start(handler)
}
```

The implementation of `parser.ParseRequest` follows [webhook.go](https://github.com/line/line-bot-sdk-go/blob/master/linebot/webhook.go)  in [line/line-bot-sdk-go](https://github.com/line/line-bot-sdk-go).

## Install

``` bash
$ go get github.com/shuheiktgw/go-lambda-linebot/parser
```

## Author
[Shuhei Kitagawa](https://github.com/shuheiktgw)