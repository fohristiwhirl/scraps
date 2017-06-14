package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	SERVER = "127.0.0.1"
	PORT = 8000
	FILES_DIR = "./files"
)

var known_codes = map[int]string{200: "OK", 403: "Forbidden", 404: "Not Found"}

func main() {
	server_and_port := fmt.Sprintf("%s:%d", SERVER, PORT)
	http.HandleFunc("/", main_handler)
	http.ListenAndServe(server_and_port, nil)
}

func send_error(writer http.ResponseWriter, code int) {
	writer.WriteHeader(code)
	writer.Write([]byte(fmt.Sprintf("%d %s", code, known_codes[code])))
}

func main_handler(writer http.ResponseWriter, request * http.Request) {

	req_path_clean := strings.Trim(request.URL.Path, "\n\r\t /")
	req_path_list := strings.Split(req_path_clean, "/")

	local_path_list := append([]string{FILES_DIR}, req_path_list...)
	local_filename := filepath.Join(local_path_list...)

	if strings.Contains(local_filename, "..") {		// Important safety check. But there is still danger
		send_error(writer, 403)						// if the files contain a softlink to outside...
		return
	}

	infile, err := os.Open(local_filename)

	if err != nil {
		err_string := fmt.Sprintf("%v", err)
		if strings.Contains(err_string, "cannot find the file") {
			send_error(writer, 404)
		} else {
			send_error(writer, 403)
		}
		return
	}

	buffer := make([]byte, 1024)

	for {

		chars, err := infile.Read(buffer)

		if chars > 0 {
			writer.Write(buffer[0:chars])
		}

		// Note that the first call to writer.Write triggers a header write, including auto-detecting Content-Type.

		if err != nil {
			return
		}
	}
}
