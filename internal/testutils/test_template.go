package testutils

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/RedHatInsights/sources-api-go/internal/testutils/parser"
	"github.com/RedHatInsights/sources-api-go/util"
)

// SkipIfNotRunningIntegrationTests is a helper function which skips a test if the integration tests don't want to be
// run.
func SkipIfNotRunningIntegrationTests(t *testing.T) {
	if !parser.RunningIntegrationTests {
		t.Skip("Skipping integration test")
	}
}

func NotFoundTest(t *testing.T, rec *httptest.ResponseRecorder) {
	if rec.Code != 404 {
		t.Error(fmt.Sprintf("Wrong return code: expected 404, got %d", rec.Code))
	}

	var out util.ErrorDocument
	err := json.Unmarshal(rec.Body.Bytes(), &out)
	if err != nil {
		t.Error("Failed unmarshaling output")
	}

	if len(out.Errors) == 0 {
		t.Error("Error message is empty")
	}

	for _, src := range out.Errors {
		if !strings.HasSuffix(src.Detail, "not found") {
			t.Error(fmt.Sprintf("Wrong error message: expected suffix 'not found' in '%s'", src.Detail))
		}
		if src.Status != "404" {
			t.Error(fmt.Sprintf("Wrong error status: expected 404, got %s", src.Status))
		}
	}
}
