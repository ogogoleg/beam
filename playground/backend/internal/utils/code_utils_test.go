// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	pb "beam.apache.org/playground/backend/internal/api/v1"
	"testing"
)

func TestGetCodeName(t *testing.T) {
	type args struct {
		name string
		sdk  pb.Sdk
	}
	tests := []struct {
		name           string
		args           args
		expectedResult string
	}{
		{
			name: "Get code name when name is empty and sdk is JAVA",
			args: args{
				name: "",
				sdk:  pb.Sdk_SDK_JAVA,
			},
			expectedResult: defaultJavaFileName,
		},
		{
			name: "Get code name when name is empty and sdk is GO",
			args: args{
				name: "",
				sdk:  pb.Sdk_SDK_GO,
			},
			expectedResult: defaultGoFileName,
		},
		{
			name: "Get code name when name is empty and sdk is PYTHON",
			args: args{
				name: "",
				sdk:  pb.Sdk_SDK_PYTHON,
			},
			expectedResult: defaultPythonFileName,
		},
		{
			name: "Get code name when name is empty and sdk is SCIO",
			args: args{
				name: "",
				sdk:  pb.Sdk_SDK_SCIO,
			},
			expectedResult: defaultScioFileName,
		},
		{
			name: "Get code name when name is a random string and sdk is JAVA",
			args: args{
				name: "MOCK_NAME",
				sdk:  pb.Sdk_SDK_JAVA,
			},
			expectedResult: defaultJavaFileName,
		},
		{
			name: "Get code name when name has wrong extension and sdk is JAVA",
			args: args{
				name: "MOCK_NAME.py",
				sdk:  pb.Sdk_SDK_JAVA,
			},
			expectedResult: "MOCK_NAME.java",
		},
		{
			name: "Get code name when name is a random string and sdk is GO",
			args: args{
				name: "MOCK_NAME",
				sdk:  pb.Sdk_SDK_GO,
			},
			expectedResult: defaultGoFileName,
		},
		{
			name: "Get code name when name has wrong extension and sdk is GO",
			args: args{
				name: "MOCK_NAME.py",
				sdk:  pb.Sdk_SDK_GO,
			},
			expectedResult: "MOCK_NAME.go",
		},
		{
			name: "Get code name when name is a random string and sdk is PYTHON",
			args: args{
				name: "MOCK_NAME",
				sdk:  pb.Sdk_SDK_PYTHON,
			},
			expectedResult: defaultPythonFileName,
		},
		{
			name: "Get code name when name has wrong extension and sdk is PYTHON",
			args: args{
				name: "MOCK_NAME.java",
				sdk:  pb.Sdk_SDK_PYTHON,
			},
			expectedResult: "MOCK_NAME.py",
		},
		{
			name: "Get code name when name is a random string and sdk is SCIO",
			args: args{
				name: "MOCK_NAME",
				sdk:  pb.Sdk_SDK_SCIO,
			},
			expectedResult: defaultScioFileName,
		},
		{
			name: "Get code name when name has wrong extension and sdk is SCIO",
			args: args{
				name: "MOCK_NAME.java",
				sdk:  pb.Sdk_SDK_SCIO,
			},
			expectedResult: "MOCK_NAME.scala",
		},
		{
			name: "Get code name when name is correct and sdk is JAVA",
			args: args{
				name: "MOCK_NAME.java",
				sdk:  pb.Sdk_SDK_JAVA,
			},
			expectedResult: "MOCK_NAME.java",
		},
		{
			name: "Get code name when name is correct and sdk is GO",
			args: args{
				name: "MOCK_NAME.go",
				sdk:  pb.Sdk_SDK_GO,
			},
			expectedResult: "MOCK_NAME.go",
		},
		{
			name: "Get code name when name is correct and sdk is PYTHON",
			args: args{
				name: "MOCK_NAME.py",
				sdk:  pb.Sdk_SDK_PYTHON,
			},
			expectedResult: "MOCK_NAME.py",
		},
		{
			name: "Get code name when name is correct and sdk is SCIO",
			args: args{
				name: "MOCK_NAME.scala",
				sdk:  pb.Sdk_SDK_SCIO,
			},
			expectedResult: "MOCK_NAME.scala",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualResult := GetCodeName(tt.args.name, tt.args.sdk)
			if actualResult != tt.expectedResult {
				t.Errorf("GetCodeName() actual result is not equal to the expected result")
			}
		})
	}
}

func TestIsCodeMain(t *testing.T) {
	type args struct {
		code string
		sdk  pb.Sdk
	}
	tests := []struct {
		name           string
		args           args
		expectedResult bool
	}{
		{
			name: "Is code main when code doesn't have main method and sdk is JAVA",
			args: args{
				code: "MOCK_CODE",
				sdk:  pb.Sdk_SDK_JAVA,
			},
			expectedResult: false,
		},
		{
			name: "Is code main when code doesn't have main method and sdk is GO",
			args: args{
				code: "MOCK_CODE",
				sdk:  pb.Sdk_SDK_GO,
			},
			expectedResult: false,
		},
		{
			name: "Is code main when code doesn't have main method and sdk is PYTHON",
			args: args{
				code: "MOCK_CODE",
				sdk:  pb.Sdk_SDK_PYTHON,
			},
			expectedResult: false,
		},
		{
			name: "Is code main when code doesn't have main method and sdk is SCIO",
			args: args{
				code: "MOCK_CODE",
				sdk:  pb.Sdk_SDK_SCIO,
			},
			expectedResult: false,
		},
		{
			name: "Is code main when code has main method and sdk is JAVA",
			args: args{
				code: "MOCK_CODEpublic static void main(String[] args)MOCK_CODE",
				sdk:  pb.Sdk_SDK_JAVA,
			},
			expectedResult: true,
		},
		{
			name: "Is code main when code has main method and sdk is GO",
			args: args{
				code: "MOCK_CODEfunc main()MOCK_CODE",
				sdk:  pb.Sdk_SDK_GO,
			},
			expectedResult: true,
		},
		{
			name: "Is code main when code has main method and sdk is PYTHON",
			args: args{
				code: "MOCK_CODEif __name__ == '__main__'MOCK_CODE",
				sdk:  pb.Sdk_SDK_PYTHON,
			},
			expectedResult: true,
		},
		{
			name: "Is code main when code has main method and sdk is SCIO",
			args: args{
				code: "MOCK_CODEdef main(cmdlineArgs: Array[String])MOCK_CODE",
				sdk:  pb.Sdk_SDK_SCIO,
			},
			expectedResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualResult := IsCodeMain(tt.args.code, tt.args.sdk)
			if actualResult != tt.expectedResult {
				t.Errorf("IsCodeMain() actual result is not equal to the expected result")
			}
		})
	}
}
