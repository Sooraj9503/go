package Model


type Student struct {
    RollNo      uint   `gorm:"primary_key"`
    Name        string `gorm:"type:varchar(50)"`
    Address     string `gorm:"type:varchar(100)"`
    ParentName  string `gorm:"type:varchar(50)"`
    Standard    uint   `gorm:"type:integer unsigned"`
    DOB         string `gorm:"type:date"`
    ContactNo   string `gorm:"type:varchar(15)"`
}