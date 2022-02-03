package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

var AllowedFileExtensions = []string{"csv"}

func IsAllowedFile(fileExt string) bool {
	for _, val := range AllowedFileExtensions {
		if strings.ToLower(fileExt) == val {
			return true
		}
	}
	return false
}

func SaveLogFile(m []*multipart.FileHeader) (string, error) {
	var err error
	var filename, filepath string

	for i := range m {
		chunks := strings.Split(m[i].Filename, ".")
		if IsAllowedFile(chunks[1]) {
			// Set the upload path for the logfile
			// TODO: Remove hardcoded destination.
			uploadPath := fmt.Sprintf("uploads/csv/")

			// Test for the existence of the upload directory.
			_, err := os.Stat(uploadPath)
			if os.IsNotExist(err) {
				// The upload directory does not exist. Create it.
				err = os.Mkdir(uploadPath, 0755)
				if err != nil && strings.Contains(err.Error(), "file exists") {
					return filepath, err
				}
			}

			// Alter the filename to include timestamp
			filename = fmt.Sprintf("%s_%d.%s", chunks[0], time.Now().Unix(), chunks[1])
			filepath = uploadPath + filename
			filehandle, err := os.Create(filepath)
			if err != nil {
				return filepath, err
			}

			// Close the destination file handle on function return
			defer filehandle.Close()

			// Limit access restrictions
			defer os.Chmod(filepath, (os.FileMode)(0755))
			if err != nil {
				return filepath, err
			}

			// Copy the uploaded file to the destination file
			file, err := m[i].Open()
			defer file.Close()

			if _, err := io.Copy(filehandle, file); err != nil {
				return filepath, err
			}
			continue
		}
	}
	return filepath, err
}
