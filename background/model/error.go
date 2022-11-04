package model

import "fmt"
type ClientError struct {
	Code    int    `json:"name"`
	Message string `json:"message"`
}
func (err ClientError) Error() string {
    return fmt.Sprintf("Status code: %d, message: %s", err.Code, err.Message)
}
