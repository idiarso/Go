package services

var attendanceRecords []interface{} // Change to a generic type or define a new struct if needed

func MarkAttendance(userID int, status string) interface{} {
	record := map[string]interface{}{"UserID": userID, "Status": status}
	attendanceRecords = append(attendanceRecords, record)
	return record
}

func GetAttendance() []interface{} {
	return attendanceRecords
}
