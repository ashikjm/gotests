package base

import (
	"github.com/gorilla/mux"
	"net/http"
	//"fmt"
	//"fmt"
	//"fmt"
	"fmt"
)


//Handler.go here


//func getAllMaintenanceHandler(r *http.Request) ([]byte, *appError) {
//	a := []byte("Success accessing this")
//	//(*writer).Write(a)
//	return a, nil
//}





// Main.go here
type appError struct {
	code    int
	message string
}

type appHandler func(*http.Request) ([]byte, *appError)

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, err := fn(r)
	if err != nil {
		http.Error(w, `{ "error": "`+err.message+`" }`, err.code)
	} else {
		w.Write(body)
	}
}

func Main() {
	err := DBConnect()
	if err != nil {
		fmt.Errorf("Couldn't connect to DB")
	}

	err = migrate_db()
	if err != nil {
		fmt.Errorf("Couldn't connect to DB")
	}

	router := mux.NewRouter().StrictSlash(true)
	//router.Path("/maintenance/").Methods("GET").Handler(appHandler(GetAllMaintenanceHandlerTest))
	//router.Path("/maintenance/").Methods("POST").Handler(appHandler(CreateMaintenanceHandlerTest))
	//router.Path("/maintenance/{id}").Methods("GET").Handler(appHandler(GetMaintenanceHandlerTest))
	router.Path("/maintenance/{id}").Methods("GET").Handler(appHandler(GetMaintenanceHandler))
	router.Path("/maintenance/").Methods("POST").Handler(appHandler(CreateMaintenanceHandler))
	router.Path("/maintenance/{id}/schedule/").Methods("POST").Handler(appHandler(ScheduleMaintenanceHandler))
	router.Path("/maintenance/{id}/start/").Methods("POST").Handler(appHandler(StartMaintenanceHandler))
	router.Path("/maintenance/{id}/complete/").Methods("POST").Handler(appHandler(CompleteMaintenanceHandler))
	http.ListenAndServe("0.0.0.0:8000", router)

}
