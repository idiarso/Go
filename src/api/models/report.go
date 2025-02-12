package models

// AttendanceReport represents the structure of an attendance report

type AttendanceReport struct {
    UserID    int    `json:"user_id"`
    UserName  string `json:"user_name"`
    AttendanceCount int `json:"attendance_count"`
}

// StudentPerformanceReport represents the structure of a student performance report

type StudentPerformanceReport struct {
    StudentID int    `json:"student_id"`
    Name      string `json:"name"`
    AverageScore float64 `json:"average_score"`
    TotalAttendance int `json:"total_attendance"`
}
