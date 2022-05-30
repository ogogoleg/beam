#!/bin/bash
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Launch firestore emulator on port 8082
PID=$(lsof -t -i :8082 -s tcp:LISTEN)
if [ -z "$PID" ]; then
  echo "Starting mock Firestore server on port 8082"
  nohup gcloud beta emulators firestore start \
    --host-port=127.0.0.1:8082 \
    >/tmp/mock-firestore-logs &
else
  echo "There is an instance of Firestore already running on port 8082"
fi
