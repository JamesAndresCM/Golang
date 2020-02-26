package user

import "time"

type User struct {
  Id int
  Name string
  LastName string
  DateActive time.Time
  Status bool
}

func (this *User) Initialize(id int, name string, lastName string, date time.Time, status bool){
  this.Id = id
  this.Name = name
  this.LastName = lastName
  this.DateActive = date
  this.Status = status
}

func (this *User) ToS() (complete string) {
  complete = this.Name + " " + this.LastName
  return 
}
