package services

import (
	"errors"
	"st-portier-be/config"
	"st-portier-be/models"
)

// CreateEmployee creates a new employee in the database
func CreateEmployee(employee *models.Employee) error {
	if err := config.DB.Create(&employee).Error; err != nil {
		return err
	}
	return nil
}

// GetEmployeeByID fetches an employee by their ID
func GetEmployeeByID(employeeID int) (*models.Employee, error) {
	var employee models.Employee
	if err := config.DB.First(&employee, employeeID).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

// GetAllEmployeesByCompany fetches all employees for a specific company
func GetAllEmployeesByCompany(companyID int) ([]models.Employee, error) {
	var employees []models.Employee
	if err := config.DB.Where("company_id = ?", companyID).Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

// GetAllEmployees fetches all employees (for Super Admin)
func GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	if err := config.DB.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

// UpdateEmployee updates an employee's details in the database
func UpdateEmployee(employeeID int, updatedEmployee *models.Employee) error {
	var employee models.Employee

	// Find the employee by ID
	if err := config.DB.First(&employee, employeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	// Update the employee details
	employee.Name = updatedEmployee.Name
	employee.Email = updatedEmployee.Email
	employee.Phone = updatedEmployee.Phone
	employee.CompanyID = updatedEmployee.CompanyID

	// Save the updated employee
	if err := config.DB.Save(&employee).Error; err != nil {
		return err
	}
	return nil
}

// DeleteEmployee deletes an employee from the database
func DeleteEmployee(employeeID int) error {
	var employee models.Employee

	// Find the employee by ID
	if err := config.DB.First(&employee, employeeID).Error; err != nil {
		return errors.New("employee not found")
	}

	// Delete the employee
	if err := config.DB.Delete(&employee).Error; err != nil {
		return err
	}
	return nil
}
