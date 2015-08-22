# Retry

A simple go package for retrying execution of a function that generates error. The time delay between errors and the number of retires until give up can be manually specified. 

## Usage

Suppose there is a function `HttpFetch` that fetches an html document from the web and returns two arguments, the string document and an error.
Sometimes due to temporary network failure, the request may fail and on retrying it after some duration, it request completes sucessfully.
To implement the retrying mechanism using this package:

```go
  
var doc string

retry.Do(func() (err error) {
  doc, err = HttpFetch("http://google.com")
  return err
})

``` 

## Options

You can manually specify the number of times to retry executing the given function until it returns the error. You can also manually specifiy weather or not to panic in case of error.
The time duration to sleep before retrying can also be changed from its default value.

```go

retry.NumRetries = 5              // default: 15
retry.Delay = 10 * time.Second    // default: 2 seconds
retry.PanicEnabled = true         // default: false

```
