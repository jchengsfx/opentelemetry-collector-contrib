// Copyright 2019 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package honeycombexporter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/consumer/pdata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestSpanAttributesToMap(t *testing.T) {

	spanAttrs := []pdata.AttributeMap{
		pdata.NewAttributeMap().InitFromMap(map[string]pdata.AttributeValue{
			"foo": pdata.NewAttributeValueString("bar"),
		}),
		pdata.NewAttributeMap().InitFromMap(map[string]pdata.AttributeValue{
			"foo": pdata.NewAttributeValueInt(1234),
		}),
		pdata.NewAttributeMap().InitFromMap(map[string]pdata.AttributeValue{
			"foo": pdata.NewAttributeValueBool(true),
		}),
		pdata.NewAttributeMap().InitFromMap(map[string]pdata.AttributeValue{
			"foo": pdata.NewAttributeValueDouble(0.3145),
		}),
		pdata.NewAttributeMap(),
	}

	wantResults := []map[string]interface{}{
		{"foo": "bar"},
		{"foo": int64(1234)},
		{"foo": true},
		{"foo": 0.3145},
		{},
		{},
	}

	for i, attrs := range spanAttrs {
		got := spanAttributesToMap(attrs)
		want := wantResults[i]
		for k := range want {
			if got[k] != want[k] {
				t.Errorf("Got: %+v, Want: %+v", got[k], want[k])
			}
		}
	}
}

func TestTimestampToTime(t *testing.T) {
	var t1 time.Time
	emptyTime := timestampToTime(pdata.TimestampUnixNano(0))
	if t1 != emptyTime {
		t.Errorf("Expected %+v, Got: %+v\n", t1, emptyTime)
	}

	t2 := time.Now()
	seconds := t2.UnixNano() / 1000000000
	nowTime := timestampToTime(pdata.TimestampToUnixNano(
		&timestamppb.Timestamp{
			Seconds: seconds,
			Nanos:   int32(t2.UnixNano() - (seconds * 1000000000)),
		}))

	if !t2.Equal(nowTime) {
		t.Errorf("Expected %+v, Got %+v\n", t2, nowTime)
	}
}

func TestStatusCode(t *testing.T) {
	status := pdata.NewSpanStatus()
	assert.Equal(t, int32(pdata.StatusCodeUnset), getStatusCode(status), "empty")

	status.SetCode(pdata.StatusCodeError)
	assert.Equal(t, int32(pdata.StatusCodeError), getStatusCode(status), "error")

	status.SetCode(pdata.StatusCodeOk)
	assert.Equal(t, int32(pdata.StatusCodeOk), getStatusCode(status), "ok")
}

func TestStatusMessage(t *testing.T) {
	status := pdata.NewSpanStatus()
	assert.Equal(t, "STATUS_CODE_UNSET", getStatusMessage(status), "empty")

	status.SetMessage("custom message")
	assert.Equal(t, "custom message", getStatusMessage(status), "custom")
}
