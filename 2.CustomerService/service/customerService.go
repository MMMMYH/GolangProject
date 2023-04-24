package service

import (
	"Project/2.CustomerService/model"
)

// 对Customer进行操作,增删改查
type CustomerService struct {
	customers []model.Customer
	//含有多少个客户
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 18, "12345678901", "123@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (c *CustomerService) List() []model.Customer {
	return c.customers
}

func (c *CustomerService) Add(customer model.Customer) bool {
	c.customerNum++
	customer.Id = c.customerNum
	c.customers = append(c.customers, customer)
	return true
}

func (c *CustomerService) Delete(id int) bool {
	index := c.FindById(id)
	if index == -1 {
		return false
	}
	c.customers = append(c.customers[:index], c.customers[index+1:]...)
	return true
}

func (c *CustomerService) FindById(id int) int {
	index := 1
	for i := 0; i < len(c.customers); i++ {
		if c.customers[i].Id == id {
			index = i
		}
	}
	return index
}

func (c *CustomerService) Update(index int, customer model.Customer) bool {
	if index < 0 || index >= len(c.customers) {
		return false
	}
	c.customers[index] = customer
	return true
}
