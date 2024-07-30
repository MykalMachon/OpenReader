package ingest

import (
	"os"
	"path"
	"testing"

	"github.com/mykalmachon/openreader/pkg/database"
)

func TestRunIngest(t *testing.T) {
	database, err := database.Init(":memory:")

	if err != nil {
		t.Errorf("Failed to init testing database %v", err)
	}

	packageDir, err := os.Getwd()
	if err != nil {
		t.Errorf("Failed to fetch working directory via os.Getwd(): %v", err)
	}

	ingester := IngestService{
		Db:       database,
		FilePath: path.Join(packageDir, "../../internal/testdata"),
	}

	err = ingester.RunIngest()

	if err != nil {
		t.Errorf("Failed to run ingest without throwing an error: %v", err)
	}
}
