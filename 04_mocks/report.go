package _4_mocks

type IDatabase interface {
	GetNumberOfUsers() (int, error)
}

func NewReportController(database IDatabase) ReportController {
	return ReportController{database: database}
}

type ReportController struct {
	database IDatabase
}

func (r ReportController) CreteReport() (report Report, err error) {
	report.NumberOfUsers, err = r.database.GetNumberOfUsers()
	return
}

type Report struct {
	NumberOfUsers int
}
