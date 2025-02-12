package services

import "github.com/idiarso/belajar-git/src/api/models"

var internships []models.Internship
var companies []models.Company

func AddInternship(companyName string, studentID int, startDate string, endDate string, status string) models.Internship {
    internship := models.Internship{CompanyName: companyName, StudentID: studentID, StartDate: startDate, EndDate: endDate, Status: status}
    internships = append(internships, internship)
    return internship
}

func GetInternships() []models.Internship {
    return internships
}

func GetInternshipByID(id int) (models.Internship, bool) {
    for _, internship := range internships {
        if internship.ID == id {
            return internship, true
        }
    }
    return models.Internship{}, false
}
