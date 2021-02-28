package main

import (
	"log"
	"math/big"
	"net/http"

	"github.com/teejays/clog"
	"github.com/teejays/gopi/mux"
)

func main() {
	clog.LogLevel = 0
	routes := []mux.Route{
		{
			Method:      http.MethodGet,
			Version:     1,
			Path:        "api/power",
			HandlerFunc: PowerHandler,
		},
	}
	log.Fatal(mux.StartServer("0.0.0.0", 8080, routes, nil, []mux.MiddlewareFunc{mux.LoggerMiddleware}, nil))

}

// PowerHandler accepts a payload of the form:
// {
//		"Base": 4,
// 		"Power": 2,
// }
// And returns the power 2 of base 4, i.e. 4^2 = 16, in the form
// {
//		"Data": 16,
//		"Error": nil,
// }
func PowerHandler(w http.ResponseWriter, r *http.Request) {

	// Get the request
	type payload struct {
		Base  float64
		Power int
	}
	var p payload
	err := mux.UnmarshalJSONFromRequest(r, &p)
	if err != nil {
		mux.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}
	clog.Debugf("%+v", p)
	result := power(p.Base, p.Power)

	mux.WriteStandardResponse(w, result)

}

func power(x float64, n int) string {
	var xB = big.NewFloat(x)
	var r = powerBig(xB, n)
	return r.Text('f', 10)
}

func powerBig(x *big.Float, n int) *big.Float {
	var r = big.NewFloat(1)
	for i := 0; i < n; i++ {
		r = r.Mul(r, x)
		clog.Debugf("Power calc %d: %s", i, r.Text('f', 10))
	}
	return r
}
