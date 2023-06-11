# go-api-boiler
This is a template project for Go API projects. It's copied from this [yt video](https://www.youtube.com/watch?v=CJfE9kD_i7Q)

Below you will find my attempt of explaining what the functions are doing (if you see an error or a misinterpretation, feel free to create an issue on git).


## Components

### apiFunc `type func(http.ResponseWriter, *http.Request) error`

`apiFunc` is a function signature for all API functions.

Takes a standard input of a HTTP API function (`http.ResponseWriter` and `*http.Request`) and returns an `error`.

```
type apiFunc func(http.ResponseWriter, *http.Request) error
```


### WriteJSON `func (http.ResponseWriter,int,any) error`

`WriteJSON` function assigns:
- an HTTP response Header for the request(ex. 200, 404, 500 etc.),
- a `Content-Type` header with attribute `application/json` (for JSON APIs)

and returns a JSON Encoded response based on the argument of type `any` (`v any`).
`json.NewEncoder(w)` creates a new Encoder and writes to `w http.ResponseWriter`, then `Encode(v)` writes the JSON encoding of v to the stream.


```
func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
```

### MakeHTTPHandler `makeHTTPHandler(apiFunc) http.HandlerFunc`

`MakeHTTPHandler` is creating an implementation for all functions of type `apiFunc`

```
func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil { 
			if e, ok := err.(apiError); ok {
				writeJSON(w, e.Status, e)
				return
			}
			writeJSON(w, http.StatusInternalServerError, apiError{Err: "internal server error"})
		}
	}
}
```
