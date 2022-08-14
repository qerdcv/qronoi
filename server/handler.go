package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/qerdcv/voronoi/voronoi"
)

func index(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if kv := q.Get("kw"); kv != "" {
		wd, _ := strconv.ParseBool(q.Get("with_dots"))
		v := voronoi.New(kv, wd)
		w.Header().Set("Content-Type", "image/png")
		if err := v.Export(w); err != nil {
			if wErr := writeError(w, err.Error(), http.StatusInternalServerError); wErr != nil {
				log.Printf("ERROR: %s\n", wErr.Error())
			}
		}
		return
	}

	if err := writeError(w, "key word cannot be empty", http.StatusBadRequest); err != nil {
		log.Printf("ERROR: %s\n", err.Error())
	}
}

func writeError(w http.ResponseWriter, errMsg string, code int) error {
	b, err := json.Marshal(map[string]string{"message": errMsg})
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err = w.Write(b); err != nil {
		return fmt.Errorf("writer write: %w", err)
	}

	return nil
}
