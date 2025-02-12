package models

type Attendance struct {
    ID        int    `json:"id"`
    UserID    int    `json:"user_id"`
    Timestamp string `json:"timestamp"`
    Status    string `json:"status"`
}
