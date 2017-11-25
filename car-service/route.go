package car

// import (
// 	"encoding/json"
// 	"net/http"
// )


// type route struct {
// 	model Model
// }

// func (r route) GetOne (w http.ResponseWriter, r* http.Request) {
// 	var req OneRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return 
// 	}

// 	res := r.model.One(req)
	
// 	if err := json.NewEncoder(w).Encode(res); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return 
// 	}
// }

// func NewRoute (model Model) *route {
// 	return &route{
// 		model: model,
// 	}
// }