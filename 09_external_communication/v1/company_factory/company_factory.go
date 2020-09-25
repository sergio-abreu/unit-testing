package company_factory

import (
	"sergio/unit-testing/09_external_communication/v1/user"
)

func Create(data map[string]interface{}) *user.Company {
	domainName := data["domainName"].(string)
	numberOfEmployees := data["numberOfEmployees"].(int)
	return user.NewCompany(domainName, numberOfEmployees)
}
