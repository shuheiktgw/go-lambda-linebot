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
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/line/line-bot-sdk-go/linebot"
)

// ParseAPIGatewayProxyRequest turns events.APIGatewayProxyRequest into array of linebot.Event
func ParseAPIGatewayProxyRequest(channelSecret string, r *events.APIGatewayProxyRequest) ([]*linebot.Event, error) {
	body := []byte(r.Body)

	if !validateSignature(channelSecret, r.Headers["X-Line-Signature"], body) {
		return nil, linebot.ErrInvalidSignature
	}

	request := &struct {
		Events []*linebot.Event `json:"events"`
	}{}

	if err := json.Unmarshal(body, request); err != nil {
		return nil, err
	}

	return request.Events, nil
}

// ParseSNSEvent turns events.SNSEvent into linebot.Event
func ParseSNSEvent(r *events.SNSEvent) (*linebot.Event, error) {
	var event linebot.Event
	if err := json.Unmarshal([]byte(r.Records[0].SNS.Message), &event); err != nil {
		return nil, err
	}

	return &event, nil
}

func validateSignature(channelSecret, signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}

	hash := hmac.New(sha256.New, []byte(channelSecret))
	_, err = hash.Write(body)
	if err != nil {
		return false
	}

	return hmac.Equal(decoded, hash.Sum(nil))
}
