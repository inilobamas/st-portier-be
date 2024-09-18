package services

import (
	"errors"
	"st-portier-be/config"
	"st-portier-be/models"
)

// GetAllCompanies retrieves all companies from the database
func GetAllCompanies() ([]models.Company, error) {
	var companies []models.Company

	// Query to get all companies
	if err := config.DB.Find(&companies).Error; err != nil {
		return nil, err
	}

	return companies, nil
}

// CreateCompany creates a new company in the database
func CreateCompany(company *models.Company) error {
	if err := config.DB.Create(&company).Error; err != nil {
		return err
	}
	return nil
}

// GetCompanyByID retrieves a company by its ID
func GetCompanyByID(companyID int) (*models.Company, error) {
	var company models.Company
	if err := config.DB.First(&company, "id = ?", companyID).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

// GetCompanyByIDAndUserCompany retrieves a company by ID if it matches the user's company
func GetCompanyByIDAndUserCompany(userCompanyID int, companyID int) (*models.Company, error) {
	var company models.Company
	if err := config.DB.First(&company, "id = ? AND company_id = ?", companyID, userCompanyID).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

// UpdateCompany updates the details of a company
func UpdateCompany(companyID int, updatedCompany *models.Company) error {
	var company models.Company

	// Find the company by ID
	if err := config.DB.First(&company, "id = ?", companyID).Error; err != nil {
		return errors.New("company not found")
	}

	// Update the company's details
	company.Name = updatedCompany.Name
	company.Description = updatedCompany.Description

	// Save the updated company
	if err := config.DB.Save(&company).Error; err != nil {
		return err
	}
	return nil
}

// DeleteCompany deletes a company from the database
func DeleteCompany(companyID int) error {
	var company models.Company

	// Find the company by ID
	if err := config.DB.First(&company, "id = ?", companyID).Error; err != nil {
		return errors.New("company not found")
	}

	// Delete the company
	if err := config.DB.Delete(&company).Error; err != nil {
		return err
	}
	return nil
}
