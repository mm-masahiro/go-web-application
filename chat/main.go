package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
					<title>チャット</title>
				</head>
				<body>
					Let's Chat!!
				</body>
			</html>
		`))
	})

	if err := http.ListenAndServe(":8800", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
