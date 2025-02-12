package models

type Internship struct {
    ID          int    `json:"id"`
    CompanyName string `json:"company_name"`
    StudentID   int    `json:"student_id"`
    StartDate   string `json:"start_date"`
    EndDate     string `json:"end_date"`
    Status      string `json:"status"`
}

// Company model

type Company struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Location string `json:"location"`
}
