package v2_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	v2 "sergio/unit-testing/07_functional/v2"
	"sergio/unit-testing/07_functional/v2/mocks"
	"testing"
	"time"
)

func Test_A_new_file_is_created_when_the_current_file_overflows(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mocks.NewMockIFileSystem(ctrl)
	mock.EXPECT().
		GetFiles("/tmp/audits").
		Return([]string{
			`/tmp/audits/audit-1.txt`,
			`/tmp/audits/audit-2.txt`,
		})
	mock.EXPECT().
		ReadAllLines(`/tmp/audits/audit-2.txt`).
		Return([]string{
			"Peter; 2019-04-06T16:30:00Z",
			"Jane; 2019-04-06T16:40:00Z",
			"Jack; 2019-04-06T17:00:00Z",
		}, nil)
	sut := v2.NewAuditManager(3, "/tmp/audits", mock)

	mock.EXPECT().
		WriteText("/tmp/audits/audit-3.txt", "Alice; 2019-04-06T18:00:00Z").
		Times(1)

	err := sut.AddRecord("Alice", time.Date(2019, 04, 06, 18, 0, 0, 0, time.UTC))

	g.Expect(err).Should(
		Not(HaveOccurred()))
}
