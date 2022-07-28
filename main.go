package main

import (
	"altip/utils"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getAlternativeIps(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addr := vars["ip"]
	prefix := vars["prefix"]

	if prefix != "" {
		prefix = prefix + "://"
	}

	if utils.IsValidIp(addr) == false {
		_, _ = fmt.Fprintf(w, "error: invalid ip")
		return
	}
	tokens := utils.Tokenize(addr)

	_, _ = fmt.Fprintf(w, "%s%d\n", prefix, (tokens[0]<<24)|(tokens[1]<<16)|(tokens[2]<<8)|tokens[3])
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, utils.SimpleTransform("0x%02X", tokens))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, utils.SimpleTransform("%04o", tokens))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, utils.SimpleTransform("0x%010X", tokens))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, utils.SimpleTransform("%010o", tokens))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, utils.ConditionalTransform(3, "%d", "0x%02X", tokens))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, utils.ConditionalTransform(2, "%d", "0x%02X", tokens))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, utils.ConditionalTransform(1, "%d", "0x%02X", tokens))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, utils.ConditionalTransform(3, "%d", "%04o", tokens))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, utils.ConditionalTransform(2, "%d", "%04o", tokens))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, utils.ConditionalTransform(1, "%d", "%04o", tokens))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, utils.TransformLeftShift(2, "0x%02X", "%d", tokens))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, utils.TransformLeftShift(2, "%04o", "%d", tokens))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, fmt.Sprintf("0x%02X.%d", tokens[0], (tokens[1]<<16)|(tokens[2]<<8)|tokens[3]))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, fmt.Sprintf("%04o.%d", tokens[0], (tokens[1]<<16)|(tokens[2]<<8)|tokens[3]))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, utils.ConditionalTransform(2, "%04o", "0x%02X", tokens))
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, utils.ConditionalTransform(1, "%04o", "0x%02X", tokens))

	result := ""
	for i := 0; i < 2; i++ {
		if i >= 1 {
			result += fmt.Sprintf("%04o.", tokens[i])
			result += fmt.Sprintf("%d", (tokens[2]<<8)|tokens[3])
		} else {
			result += fmt.Sprintf("0x%02X.", tokens[i])
		}
	}
	_, _ = fmt.Fprintf(w, "%s%s\n", prefix, result)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{ip}", getAlternativeIps)
	router.HandleFunc("/{ip}/{prefix}", getAlternativeIps)

	log.Fatal(http.ListenAndServe(":8066", router))
}
