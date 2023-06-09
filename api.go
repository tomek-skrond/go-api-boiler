package main

///////////// TUTAJ API /////////////
import (
	"encoding/json"
	"net/http"
)

// typowana funkcja, bierze responsewritera i request jako arg i zwraca typ error
type apiFunc func(http.ResponseWriter, *http.Request) error

// funkcja tworzaca implementacje dla typow funkcji apiFunc (error handling + funkcjonalnosc)
func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil { //jezeli wystapi error w funkcji o typie apiFunc, idz dalej (zawsze wystapi bo apiFunc musi zwracac err)
			if e, ok := err.(apiError); ok { //probuje konwertowac zmienna err do typu apiError, jesli sie uda - przejdz dalej
				writeJSON(w, e.Status, e) // wypisuje JSON'a do responseWriter'a + status + any (err msg)
				return
			}
			writeJSON(w, http.StatusInternalServerError, apiError{Err: "internal server error"})
		}
	}
}

// writeJSON -> zwraca JSON z odpowiednim statusem + naglowkiem
func writeJSON(w http.ResponseWriter, status int, v any) error { //response writer -> konstruuje odpowiedzi HTTP
	w.WriteHeader(status)                              //wpisuje status Request'a do response writera
	w.Header().Add("Content-Type", "application/json") //dodawnanie naglowka http
	return json.NewEncoder(w).Encode(v)                //zwracanie JSON'a (na poczatku tworzenie encodera z ResponseWriter potem kodowanie jasona do strumienia)
}

// API //

// User by ID Handling
func handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return ErrInvalidMethod //!!! Mozna zwrocic apiError poniewaz ten typ implementuje interfejs builtin.error
	}

	user := User{}

	if !user.Valid {
		return ErrUserInvalid
	}

	return writeJSON(w, http.StatusOK, user)
}
