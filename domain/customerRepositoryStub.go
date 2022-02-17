package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Özgür", City: "İstanbul", DateofBirth: "20-09-1989", Status: "1"},
		{Id: "1002", Name: "Rob", City: "New Delhi", DateofBirth: "20-09-1990", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
