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

# Launch firestore emulator on port from FIRESTORE_EMULATOR_HOST env or use default port is 8082
FIRESTORE_FULL_ADDRESS="${FIRESTORE_EMULATOR_HOST:-"localhost:8082"}"
echo "FIRESTORE_FULL_ADDRESS:" "${FIRESTORE_FULL_ADDRESS}"
FIRESTORE_PORT="${FIRESTORE_FULL_ADDRESS##*:}"
PID=$(lsof -t -i :"${FIRESTORE_PORT}" -s tcp:LISTEN)
if [ -z "$PID" ]; then
  echo "Starting mock Firestore server on port from FIRESTORE_EMULATOR_HOST env or use default port is 8082"
  nohup gcloud beta emulators firestore start \
    --host-port="${FIRESTORE_FULL_ADDRESS}" \
    >/tmp/mock-firestore-logs &
else
  echo "There is an instance of Firestore already running on port from FIRESTORE_EMULATOR_HOST env or use default port is 8082"
fi
