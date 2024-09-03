
/*
 errors.As is used to determine if an error can be assigned to a specific type, not to compare two error types. It helps you:

Check Type: Verify if the error is of a specific type or implements a certain interface.
Extract Error: Retrieve the underlying error if it matches the specified type.
Key Points about errors.As
Type Matching: errors.As checks if the error value can be assigned to the type of the target variable. This is useful for checking if an error is of a specific type or if it implements a particular interface.

Error Extraction: If the error matches the type, errors.As assigns the error to the target variable, allowing you to work with it in a type-specific manner.
*/
 package main

import (
    "errors"
    "fmt"
)

// Define custom error types
type ErrNetworkIssue struct {
    Msg string
}

func (e *ErrNetworkIssue) Error() string {
    return fmt.Sprintf("network issue: %s", e.Msg)
}

type ErrInvalidResponse struct {
    Code int
    Msg  string
}

func (e *ErrInvalidResponse) Error() string {
    return fmt.Sprintf("invalid response: %d - %s", e.Code, e.Msg)
}

// Simulate a function that returns errors
func processRequest() error {
    // Simulate an error of type ErrInvalidResponse
    return &ErrInvalidResponse{Code: 404, Msg: "Not Found"}
}

func main() {
    err := processRequest()
    
    if err != nil {
        var netErr *ErrNetworkIssue
        var respErr *ErrInvalidResponse

        if errors.As(err, &netErr) {
            fmt.Println("Handled network issue:", netErr)
        } else if errors.As(err, &respErr) {
            fmt.Printf("Handled invalid response: Code %d, Message %s\n", respErr.Code, respErr.Msg)
        } else {
            fmt.Println("Unhandled error:", err)
        }
    }
}
