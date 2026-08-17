package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
	"strings"
)

func main() {
	http.HandleFunc("/transcribe", handleTranscribe)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleTranscribe(w http.ResponseWriter, r *http.Request){
	// Only allow POST requests
	if r.Method != "POST"{
		fmt.Fprintf(w, "Use POST method\n")
		return 
	}

	// Parse the file
	file, header, err := r.FormFile("video")
	if err != nil{
		fmt.Fprintf(w, "Error reading file: %v\n", err)
		return
	}
	defer file.Close()

	// Save the file to disk
	out, err := os.Create(header.Filename)
	if err != nil{
		fmt.Fprintf(w, "Error saving file: %v\n", err)
		return
	}
	defer out.Close()

	// Copy the uploaded file to disk
	_, err = io.Copy(out, file)
	if err != nil{
		fmt.Fprintf(w, "Error writing file: %v\n", err)
		return
	}

	// Check if video file exists
	if _, err := os.Stat(header.Filename); err != nil {
		fmt.Fprintf(w, "Video file not found: %v\n", err)
		return
	}

	// Call Whisper to transcribe
	baseFilename := strings.TrimSuffix(header.Filename, ".mov")
	transcriptFile := baseFilename + ".json"
	cmd := exec.Command("/Users/cameronwilliams/.local/bin/whisper", header.Filename, "--output_format", "json", "--output_dir", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	err = cmd.Run()
	if err != nil{
		fmt.Fprintf(w, "Error transcribing: %v\n", err)
		return
	}

	fmt.Printf("Whisper finished. Looking for: %s\n", transcriptFile)

	// Wait for transcript file to be created
	for i := 0; i < 60; i++ {
		if _, err := os.Stat(transcriptFile); err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	// Read the transcript JSON
	data, err := os.ReadFile(transcriptFile)
	if err != nil {
		fmt.Fprintf(w, "Error reading transcript: %v\n", err)
		return
	}

	// Delete the video file
	os.Remove(header.Filename)

	// Return the transcript
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}