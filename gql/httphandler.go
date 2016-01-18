package gql

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

func QueryHandler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/graphql" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := graphql.Do(
			graphql.Params{
				Schema:        schema,
				RequestString: string(body),
			},
		)

		if result.HasErrors() {
			w.WriteHeader(http.StatusBadRequest)
			for _, err := range result.Errors {
				w.Write([]byte(err.Error() + "\n"))
			}
			return
		}

		js, err := json.Marshal(result.Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(js); err != nil {
			log.Println(err)
		}
	}
}
