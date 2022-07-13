// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package conf

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/paashzj/gutil"
	"os"
	"strings"
)

var (
	PulsarHost             = gutil.GetEnvStr("PULSAR_HOST", "localhost")
	PulsarPort             = gutil.GetEnvInt("PULSAR_PORT", 6650)
	PulsarTopic            = os.Getenv("PULSAR_TOPIC")
	PulsarSubscriptionName = os.Getenv("PULSAR_SUBSCRIPTION_NAME")
	pulsarSubscriptionType = os.Getenv("PULSAR_SUBSCRIPTION_TYPE")
)

func GetPulsarSubscriptionType() pulsar.SubscriptionType {
	pulsarSubscriptionType = strings.ToLower(pulsarSubscriptionType)
	if pulsarSubscriptionType == "shared" {
		return pulsar.Shared
	} else if pulsarSubscriptionType == "keyshared" {
		return pulsar.KeyShared
	} else if pulsarSubscriptionType == "exclusive" {
		return pulsar.Exclusive
	} else {
		return pulsar.Failover
	}
}
