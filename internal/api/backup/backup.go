package backup

import (
	"backup-service/internal/app"
	"encoding/json"
	"fmt"
	"net/http"
)

type API struct {
	App *app.App
}

func BackupHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/dbs":
		{
			fmt.Fprintf(w, "BackupHandler called for path: %s\n", path)
		}
	default:
		{
			http.Error(w, "Not Found", http.StatusNotFound)
		}

	}
}

func (api *API) GetDBs(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusOK)
	print("GetDBs called\n")
	data, err := json.Marshal(api.App.Containers)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshalling data: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("%s", data)
	fmt.Fprintf(w, "%s", data)
}
