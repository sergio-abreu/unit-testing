package company_factory

import (
	"sergio/unit-testing/08_code_types/v3/user"
)

func Create(data map[string]interface{}) *user.Company {
	domainName := data["domainName"].(string)
	numberOfEmployees := data["numberOfEmployees"].(int)
	return user.NewCompany(domainName, numberOfEmployees)
}
