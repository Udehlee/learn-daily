## overview


Learned how to copy data from one source to another(destination ) using io.Copy() in Go. In this project, I have used it to copy uploaded files to a TCP connection.


## implementation


conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Copy the file data to the TCP connection
	_, err = io.Copy(conn, uploadedFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
