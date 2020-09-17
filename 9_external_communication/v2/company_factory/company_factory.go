package company_factory

import (
	"sergio/unit-testing/9_external_communication/v2/user"
)

func Create(data map[string]interface{}) *user.Company {
	domainName := data["domainName"].(string)
	numberOfEmployees := data["numberOfEmployees"].(int)
	return user.NewCompany(domainName, numberOfEmployees)
}
