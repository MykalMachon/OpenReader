package ingest

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"

	"github.com/mykalmachon/openreader/pkg/database"
)

type IngestService struct {
	FilePath string
	Db       *sql.DB
}

type IngestItem struct {
	Author database.Author
	Book   database.Book
}

func (is *IngestService) RunIngest() error {
	// grab all files in the folder
	files, err := os.ReadDir(is.FilePath)

	if err != nil {
		fmt.Println("Failed to parse home directory during ingest", err)
		panic(err)
	}

	for _, file := range files {
		fileData, err := getFileData(file)
		if err != nil {
			return err
		}

		err = ingestFileData(fileData)
		if err != nil {
			return err
		}
	}
	return nil
}

func getFileData(file fs.DirEntry) (*IngestItem, error) {
	// TODO: parse the file using mupdf
	// TODO: store the information in Author and Book structs
	return nil, nil
}

func ingestFileData(item *IngestItem) error {
	// TODO: check if th author exists in db, if not, create it.
	// TODO: check if the book exists in db, if not, create it.
	return nil
}
