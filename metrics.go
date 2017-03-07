/*
Copyright © 2017 MeteoGroup Deutschland GmbH

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
 */
package main

import (
  "net/http"
  "github.com/prometheus/client_golang/prometheus/promhttp"
  "github.com/prometheus/client_golang/prometheus"
)

var messageCounter = prometheus.NewCounterVec(prometheus.CounterOpts{Name:"messages", Help: "Messages handled"}, []string{"state"})
var kafkaOffsets = prometheus.NewGaugeVec(prometheus.GaugeOpts{
  Name:"kafka_offset", Help: "Last known Kafka offsets",
  ConstLabels: prometheus.Labels{"direction": "produced"},
}, []string{"topic", "partition"})

func startPrometheusHttpExporter() {
  if metricsAddress == "" {
    return
  }
  go func() {
    http.Handle("/prometheus", promhttp.Handler())
    http.ListenAndServe(metricsAddress, nil)
  }()
}

func init() {
  prometheus.MustRegister(
    messageCounter,
    kafkaOffsets,
  )
}