// ex7.16 is a web-based calculator.
//
// Evaluate an expression for http clients, accepting `expr` and `env` as query
// parameters, eg: curl 'localhost:8080/?expr=x*y&env=x=3%20y=4'
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	eval "github.com/torbiak/gopl/ex7.14"
)

func parseEnv(s string) (eval.Env, error) {
	env := eval.Env{}
	assignments := strings.Fields(s)
	for _, a := range assignments {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			return env, fmt.Errorf("bad assignment: %s\n", a)
		}
		ident, valStr := fields[0], fields[1]
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return env, fmt.Errorf("bad value for %s: %s\n", ident, err)
		}
		env[eval.Var(ident)] = val
	}
	return env, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		exprStr := r.FormValue("expr")
		if exprStr == "" {
			http.Error(w, "no expression", http.StatusBadRequest)
			return
		}
		env, err := parseEnv(r.FormValue("env"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		expr, err := eval.Parse(exprStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, expr.Eval(env))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
