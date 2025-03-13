package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type UserNotFoundError struct {
	Message string
}

func (e *UserNotFoundError) Error() string {
	return e.Message
}

type Users struct {
	LastName  string
	FirstName string
	Surname   string
}

type Specializations struct {
	Name string
}

type Visits struct {
	UserInfo           Users
	SpecializationInfo Specializations
	VisitDate          time.Time
}

var specializationArray = []string{"Kardiolog", "Nevropatolog", "Pediatr", "Dermatolog", "Oftalmolog"}
var specializationsMap = make(map[string]Specializations)
var visitsMap = make(map[Users][]Visits)
var usersArray = make([]Users, 0)

func main() {
	var input string
	for _, specialization := range specializationArray {
		specializationsMap[specialization] = Specializations{Name: specialization}
	}
	for {
		fmt.Println("1-Save   2-GetHistory   3-GetLastVisit")
		fmt.Scanf("%s", &input)
		if input == "q" {
			break
		}
		operation, err := strconv.Atoi(input)
		if err != nil {
			clearScreen()
			fmt.Println("Noto'g'ri amal kiritdingiz")
			continue
		}
		switch operation {
		case 1:
			func_Save()
		case 2:
			func_GetHistory()
		case 3:
			func_GetLastVisit()
		default:
			clearScreen()
			fmt.Println("Noto'g'ri amal kiritdingiz")
		}
	}
}

func func_GetLastVisit() {
	user, err := selectUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	visits := visitsMap[user]
	fmt.Println(visits[len(visits)-1].VisitDate.Format("2006-01-02"))
}

func func_GetHistory() {
	user, err := selectUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	visits := visitsMap[user]
	for _, visit := range visits {
		fmt.Println(visit.SpecializationInfo.Name, visit.VisitDate.Format("2006-01-02"))
	}
}

func func_Save() {
	var lastName, firstName, surname string
	fmt.Println("F.I.O ni kiriting (Familiya Ism Sharif): ")
	fmt.Scan(&lastName, &firstName, &surname)

	user, found := findUser(lastName, firstName, surname)
	if !found {
		user = Users{
			LastName:  lastName,
			FirstName: firstName,
			Surname:   surname,
		}
		usersArray = append(usersArray, user)
	}

	var index int
	for i, specialization := range specializationArray {
		if (i+1)%3 == 0 {
			fmt.Printf("%d - %s  \n", i+1, specialization)
		} else {
			fmt.Printf("%d - %s  ", i+1, specialization)
		}
	}

	fmt.Print("\nShifokor mutaxassisligini tanlang: ")
	fmt.Scan(&index)
	if index < 1 || index > len(specializationArray) {
		fmt.Println("Xato: Noto'g'ri tanlov!")
		return
	}

	fmt.Print("Tashrif sanasini kiriting Yil-Oy-Kun (2006-01-02): ")
	var dateStr string
	fmt.Scan(&dateStr)
	visitDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Println("Xato: noto'g'ri sana formati! YYYY-MM-DD shaklida kiriting.")
		return
	}

	visit := Visits{
		UserInfo:           user,
		SpecializationInfo: specializationsMap[specializationArray[index-1]],
		VisitDate:          visitDate,
	}
	visitsMap[user] = append(visitsMap[user], visit)
}

func findUser(lastName, firstName, surname string) (Users, bool) {
	for _, u := range usersArray {
		if u.LastName == lastName && u.FirstName == firstName && u.Surname == surname {
			return u, true
		}
	}
	return Users{}, false
}

func selectUser() (Users, error) {
	fmt.Println("Bemorni tanlang: ")
	for i, user := range usersArray {
		fmt.Printf("%d - %s-%s-%s\n", i+1, user.LastName, user.FirstName, user.Surname)
	}
	var index int
	fmt.Scan(&index)
	if index < 1 || index > len(usersArray) {
		return Users{}, &UserNotFoundError{Message: "user not found"}
	}
	return usersArray[index-1], nil
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
