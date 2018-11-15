// Copyright 2016 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package parser

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"

	"github.com/line/line-bot-sdk-go/linebot"
)

var testProxyRequest = `{
    "events": [
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "message",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            },
            "message": {
                "id": "325708",
                "type": "text",
                "text": "Hello, world"
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "message",
            "timestamp": 1462629479859,
            "source": {
                "type": "group",
                "groupId": "u206d25c2ea6bd87c17655609a1c37cb8",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            },
            "message": {
                "id": "325708",
                "type": "text",
                "text": "Hello, world"
            }
        }
    ]
}
`

var expectedEvents = []*linebot.Event{
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       linebot.EventTypeMessage,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &linebot.EventSource{
			Type:   linebot.EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Message: &linebot.TextMessage{
			ID:   "325708",
			Text: "Hello, world",
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       linebot.EventTypeMessage,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &linebot.EventSource{
			Type:    linebot.EventSourceTypeGroup,
			UserID:  "u206d25c2ea6bd87c17655609a1c37cb8",
			GroupID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Message: &linebot.TextMessage{
			ID:   "325708",
			Text: "Hello, world",
		},
	},
}

func TestParseAPIGatewayProxyRequest_Success(t *testing.T) {
	// Create a right signature
	mac := hmac.New(sha256.New, []byte("rightSecret"))
	mac.Write([]byte(testProxyRequest))

	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{"X-Line-Signature": base64.StdEncoding.EncodeToString(mac.Sum(nil))},
		Body:    testProxyRequest,
	}

	got, err := ParseAPIGatewayProxyRequest("rightSecret", &request)
	if err != nil {
		t.Fatalf("ParseAPIGatewayProxyRequest returns unexpected error: %s", err)
	}

	if !reflect.DeepEqual(got, expectedEvents) {
		t.Errorf("ParseAPIGatewayProxyRequest returns unexpected events: want: %s, got %s", expectedEvents, got)
	}
}

func TestParseAPIGatewayProxyRequest_Fail(t *testing.T) {
	// Create a wrong signature
	mac := hmac.New(sha256.New, []byte("wrongSecret"))
	mac.Write([]byte(testProxyRequest))

	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{"X-Line-Signature": base64.StdEncoding.EncodeToString(mac.Sum(nil))},
		Body:    testProxyRequest,
	}

	if _, got := ParseAPIGatewayProxyRequest("rightSecret", &request); got != linebot.ErrInvalidSignature {
		t.Fatalf("ParseAPIGatewayProxyRequest returns unexpected error: want: %s, got: %s", linebot.ErrInvalidSignature, got)
	}
}
