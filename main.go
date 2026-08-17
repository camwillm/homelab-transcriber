package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func main() {
	http.HandleFunc("/transcribe", HandleTranscribe)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}

func HandleTranscribe(w http.ResponseWriter, r *http.Request){
	//here we only allow post request
	if r.Method != "POST"{
		fmt.Fprintf(w, "Use POST method\n")
		return 
	}
	//parse the file
	file, header, err := r.FormFile("video")
	if err != nil{
		fmt.Fprintf(w, "Error reading file: %v\n", err)
		return
	}
	defer file.Close()

	//save thefile using os to the disk
	out, err := os.Create(header.Filename)
	if err != nil{
		fmt.Fprintf(w, "Error saving file: %v\n", err)
		return
	}
	defer out.Close()
	//Copy the uploadedfile to the disk
	_, err = io.Copy(out, file)
	if err != nil{
		fmt.Fprintf(w, "Error Writing File: %v\n", err)
		return
	}

	fmt.Fprintf(w, "Recieved and saved the file: %s\n", header.Filename)
}
