package error_handler

import (
	"fmt"
	"net/http"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func ResponseBadRequest(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "error",
		"code":    http.StatusBadRequest,
		"message": msg,
	}
}

func ResponseOkNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"code":    http.StatusOK,
		"message": msg,
	}
}

func ResponseOkWithData(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"code":    http.StatusOK,
		"message": msg,
		"data":    data,
	}
}

func ResponseBadGetway(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "error",
		"code":    http.StatusGatewayTimeout,
		"message": msg,
	}
}
