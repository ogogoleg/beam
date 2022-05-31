package utils

import (
	pb "beam.apache.org/playground/backend/internal/api/v1"
	"beam.apache.org/playground/backend/internal/logger"
	"path/filepath"
	"strings"
)

const (
	javaMainMethod        = "public static void main(String[] args)"
	goMainMethod          = "func main()"
	pythonMainMethod      = "if __name__ == '__main__'"
	scioMainMethod        = "def main(cmdlineArgs: Array[String])"
	defaultJavaFileName   = "main.java"
	defaultGoFileName     = "main.go"
	defaultPythonFileName = "main.py"
	defaultScioFileName   = "main.scala"
	javaExt               = ".java"
	goExt                 = ".go"
	pythonExt             = ".py"
	scioExt               = ".scala"
)

func GetCodeName(name string, sdk pb.Sdk) string {
	if name == "" {
		logger.Error("The name of code file is empty. Will be used default value")
		switch sdk {
		case pb.Sdk_SDK_JAVA:
			return defaultJavaFileName
		case pb.Sdk_SDK_GO:
			return defaultGoFileName
		case pb.Sdk_SDK_PYTHON:
			return defaultPythonFileName
		case pb.Sdk_SDK_SCIO:
			return defaultScioFileName
		}
	}
	return getCorrectName(name, sdk)
}

func getCorrectName(name string, sdk pb.Sdk) string {
	ext := filepath.Ext(name)
	switch sdk {
	case pb.Sdk_SDK_JAVA:
		return getCorrectNameOrDefault(ext, javaExt, defaultJavaFileName, name)
	case pb.Sdk_SDK_GO:
		return getCorrectNameOrDefault(ext, goExt, defaultGoFileName, name)
	case pb.Sdk_SDK_PYTHON:
		return getCorrectNameOrDefault(ext, pythonExt, defaultPythonFileName, name)
	case pb.Sdk_SDK_SCIO:
		return getCorrectNameOrDefault(ext, scioExt, defaultScioFileName, name)
	}
	return name
}

func getCorrectNameOrDefault(actualExt, correctExt, defaultFileName, name string) string {
	if actualExt == "" {
		logger.Error("The name of code does not have extension. Will be used default value")
		return defaultFileName
	}
	if actualExt != correctExt {
		logger.Error("The name of code has wrong extension. Will be used correct extension according to sdk")
		return name[0:len(name)-len(actualExt)] + correctExt
	} else {
		return name
	}
}

func IsCodeMain(code string, sdk pb.Sdk) bool {
	switch sdk {
	case pb.Sdk_SDK_JAVA:
		return strings.Contains(code, javaMainMethod)
	case pb.Sdk_SDK_GO:
		return strings.Contains(code, goMainMethod)
	case pb.Sdk_SDK_PYTHON:
		return strings.Contains(code, pythonMainMethod)
	case pb.Sdk_SDK_SCIO:
		return strings.Contains(code, scioMainMethod)
	}
	return false
}
