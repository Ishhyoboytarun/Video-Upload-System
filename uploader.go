package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const (
	chunkSize   = 1024 * 1024     // 1MB chunk size
	maxAttempts = 3               // maximum number of upload attempts
	retryDelay  = 5 * time.Second // delay between upload attempts
)

func main() {
	// assume the file to upload is named "video.mp4"
	file, err := os.Open("video.mp4")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return
	}
	defer file.Close()

	fileSize, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		fmt.Printf("error getting file size: %v\n", err)
		return
	}

	// generate a random session ID to identify the upload
	sessionId := fmt.Sprintf("%x", rand.Int63())

	var uploadedBytes int64

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("Uploading attempt #%d...\n", attempt)

		if err := resumeUpload(sessionId, file, fileSize, uploadedBytes); err != nil {
			fmt.Printf("error uploading file: %v\n", err)
			uploadedBytes = 0 // reset uploaded bytes for the next attempt
			time.Sleep(retryDelay)
		} else {
			fmt.Println("Upload complete!")
			return
		}
	}

	fmt.Println("Upload failed after maximum number of attempts.")
}

func resumeUpload(sessionId string, file *os.File, fileSize, uploadedBytes int64) error {
	// pretend that network connectivity is unreliable by dropping the connection
	// after 50% upload progress
	if uploadedBytes > fileSize/2 {
		fmt.Println("Network connection lost, retrying...")
		return io.ErrUnexpectedEOF
	}

	// seek to the next chunk to upload
	if _, err := file.Seek(uploadedBytes, io.SeekStart); err != nil {
		return fmt.Errorf("error seeking file: %v", err)
	}

	// read the next chunk from the file
	buffer := make([]byte, chunkSize)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return fmt.Errorf("error reading file: %v", err)
	}

	// simulate the upload progress
	time.Sleep(500 * time.Millisecond)

	// pretend that the upload failed randomly to test the retry mechanism
	if rand.Float32() < 0.1 {
		fmt.Println("Upload failed, retrying...")
		return io.ErrUnexpectedEOF
	}

	// update the uploaded bytes and print progress
	uploadedBytes += int64(n)
	fmt.Printf("Uploaded %d/%d bytes (%.2f%%)\n", uploadedBytes, fileSize, float64(uploadedBytes)/float64(fileSize)*100)

	// if all bytes have been uploaded, return nil to signal success
	if uploadedBytes >= fileSize {
		return nil
	}

	// recursively call resumeUpload with the updated uploadedBytes
	return resumeUpload(sessionId, file, fileSize, uploadedBytes)
}

/*
This implementation uses a recursive function resumeUpload to handle each upload attempt. 
It simulates unreliable network connectivity by dropping the connection after 50% upload 
progress, and randomly failing the upload with a 10% probability to test the retry mechanism. 
The function also prints progress information to the console, and returns an error if the 
upload fails. The main function calls resumeUpload up to maxAttempts times, with a delay.
*/
