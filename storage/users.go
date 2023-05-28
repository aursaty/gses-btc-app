package storage

import (
	"gses2-btc-app/utils"
	"io"
	"log"
	"os"
	"strings"
)

// CheckForEmail function for check is email in mail storage
func CheckForEmail(email string) bool {
	// try to open file
	fi, err := os.Open(utils.EmailsFile)
	if err != nil {
		panic(err)
	}
	// close fi
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	// go throw the data, check for email
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		s := string(buf[:n])
		return strings.Contains(s, email)
	}
	return false
}

// AddEmail add an email at the end of the emails list
func AddEmail(email string) {
	// open output file
	// os.O_APPEND and os.O_WRONLY os.O_APPEND is to control behaviour in this case to Append mode so it
	// does not need to maintain a file pointer. os.O_WRONLY is to state the file in this case write.
	fo, err := os.OpenFile(utils.EmailsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//fo, err := os.Create("local/emails.csv")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// write a chunk
	if _, err := fo.WriteString(email + "\n"); err != nil {
		panic(err)
	}
}
