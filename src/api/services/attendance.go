package services

import "github.com/idiarso/belajar-git/api/src/api/models"

var attendanceRecords []models.Attendance

func MarkAttendance(userID int, status string) models.Attendance {
    record := models.Attendance{UserID: userID, Status: status}
    attendanceRecords = append(attendanceRecords, record)
    return record
}

func GetAttendance() []models.Attendance {
    return attendanceRecords
}
