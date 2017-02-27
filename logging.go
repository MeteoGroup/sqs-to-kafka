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
  "github.com/go-kit/kit/log"
  "os"
)

var LOG = log.NewJSONLogger(os.Stdout)

func logAndPanic(err error) {
  if err != nil {
    LOG.Log(
      "level", "FATAL",
      "message", err,
    )
    panic(err)
  }
}

func logInfo(message interface{}, keyvals ...interface{}) {
  LOG.Log(
    append(keyvals,
      "level", "INFO",
      "message", message,
    )...
  )
}

func logError(err error) {
  if err != nil {
    LOG.Log(
      "level", "ERROR",
      "message", err,
    )
  }
}