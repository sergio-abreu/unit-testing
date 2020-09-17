package v3_test

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	v3 "sergio/unit-testing/7_functional/v3"
	"testing"
	"time"
)

func Test_A_new_file_is_created_when_the_current_file_overflows(t *testing.T) {
	g := NewGomegaWithT(t)
	sut := v3.NewAuditManager(3)
	files := []v3.FileContent{
		{Filename: "audit-1.txt", Lines: nil},
		{Filename: "audit-2.txt", Lines: []string{
			"Peter; 2019-04-06T16:30:00Z",
			"Jane; 2019-04-06T16:40:00Z",
			"Jack; 2019-04-06T17:00:00Z",
		}},
	}

	update := sut.AddRecord(
		files, "Alice", time.Date(2019, 04, 06, 18, 0, 0, 0, time.UTC))

	g.Expect(update).Should(
		MatchAllFields(Fields{
			"Filename": Equal("audit-3.txt"),
			"Content": Equal("Alice; 2019-04-06T18:00:00Z"),
		}))
}
