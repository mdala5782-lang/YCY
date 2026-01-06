package main

import (
	"fmt"
	"net/http"
)

func StartHTTP() {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "# HELP cpu_usage CPU usage percentage")
		fmt.Fprintln(w, "# TYPE cpu_usage gauge")
		fmt.Fprintf(w, "cpu_usage %.2f\n", GetCPU())

		fmt.Fprintln(w, "# HELP mem_usage Memory usage percentage")
		fmt.Fprintln(w, "# TYPE mem_usage gauge")
		fmt.Fprintf(w, "mem_usage %.2f\n", GetMem())

		fmt.Fprintln(w, "# HELP disk_usage Disk usage percentage")
		fmt.Fprintln(w, "# TYPE disk_usage gauge")
		fmt.Fprintf(w, "disk_usage %.2f\n", GetDisk())
	})

	http.ListenAndServe(":9100", nil)
}
