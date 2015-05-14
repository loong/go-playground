Ravelin Code Test
=================

## Summary
We need an HTTP server that will accept any POST request (JSON) from muliple clients' websites. Each request forms part of a struct (for that particular visitor) that will be printed to the terminal when the struct is fully complete. 

For the JS part of the test please feel free to use any libraries that may help you **but please only use the Go standard library for the backend**.

## Frontend (JS)
Insert JavaScript into the index.html (supplied) that captures and posts data every time one of the below events happens; this means you will be posting multiple times per visitor. Assume only one resize occurs.

  - if the screen resizes, the before and after dimensions
  - copy & paste (for each field)
  - time taken from the 1st character typed to clicking the submit button

### Example JSON Requests
```
{
  "eventType": "copyAndPaste",
  "websiteUrl": "https://ravelin.com",
  "sessionId": "123123-123123-123123123",
  "pasted": true,
  "formId": "inputCardNumber"
}

{
  "eventType": "timeTaken",
  "websiteUrl": "https://ravelin.com",
  "sessionId": "123123-123123-123123123",
  "time": 72, // seconds
}

...

```

## Backend (Go)

The Backend must to:

1. Create a Server
2. Accept POST requests in JSON format similar to those specified above
3. Map the JSON requests to relevant sections of the data struct (specified below)
4. Print the struct for each stage of its construction
5. Also print the struct when it is complete (i.e. when the form submit button has been clicked)

### Go Struct
```
type Data struct {
	WebsiteUrl         string
	SessionId          string
	ResizeFrom         Dimension
	ResizeTo           Dimension
	CopyAndPaste       map[string]bool // map[fieldId]true
	FormCompletionTime int // Seconds
}

type Dimension struct {
	Width  string
	Height string
}
```




