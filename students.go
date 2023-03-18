package database

type Person struct {
	ID      uint   // уникальный идентификатор
	Name    string // имя
	Surname string // фамилия
	Phone   string // номер телефона
}

type Student struct {
	Person                        // студент является человеком, поэтому нам нужны все поля структуры Person
	GroupID        uint           // идентификатор группы, в котором учится студент
	Grades         map[string]int // баллы по неделям
	TopicsToWorkOn []string       // темы, над которыми студенту нужно поработать
	Attendance     map[int]int    // присутствие на занятиях по дням месяца
}

var Students = []Student{
	{
		Person: Person{
			ID:      16,
			Name:    "Alisher",
			Surname: "Yoqubov",
			Phone:   "987654315",
		},
		GroupID: 3,
		Grades: map[string]int{
			"week1": 86,
			"week2": 80,
			"week3": 33,
			"week4": 1,
		},
		TopicsToWorkOn: nil,
		Attendance: map[int]int{
			13: 0,
			14: 0,
			15: 0,
			16: 0,
			17: 0,
			18: 0,
		},
	},
	{
		Person: Person{
			ID:      17,
			Name:    "Amir",
			Surname: "Arifjonov",
			Phone:   "987654316",
		},
		GroupID: 3,
		Grades: map[string]int{
			"week1": 76,
			"week2": 68,
			"week3": 33,
			"week4": 1,
		},
		TopicsToWorkOn: []string{
			"methods",
		},
		Attendance: map[int]int{
			13: 0,
			14: 0,
			15: 0,
			16: 0,
			17: 0,
			18: 0,
		},
	},
	{
		Person: Person{
			ID:      18,
			Name:    "Behruz",
			Surname: "Shodiev",
			Phone:   "987654317",
		},
		GroupID: 3,
		Grades: map[string]int{
			"week1": 68,
			"week2": 62,
			"week3": 33,
			"week4": 1,
		},
		TopicsToWorkOn: []string{
			"slice",
			"map",
			"logic",
			"interface",
		},
		Attendance: map[int]int{
			13: 0,
			14: 0,
			15: 0,
			16: 0,
			17: 0,
			18: 0,
		},
	},
	{
		Person: Person{
			ID:      19,
			Name:    "Farhod",
			Surname: "Olimzoda",
			Phone:   "987654318",
		},
		GroupID: 3,
		Grades: map[string]int{
			"week1": 88,
			"week2": 86,
			"week3": 33,
			"week4": 1,
		},
		TopicsToWorkOn: nil,
		Attendance: map[int]int{
			13: 0,
			14: 0,
			15: 0,
			16: 0,
			17: 0,
			18: 0,
		},
	},
	{
		Person: Person{
			ID:      20,
			Name:    "Foziljon",
			Surname: "Muminov",
			Phone:   "987654319",
		},
		GroupID: 3,
		Grades: map[string]int{
			"week1": 85,
			"week2": 88,
			"week3": 33,
			"week4": 1,
		},
		TopicsToWorkOn: []string{
			"little problem with slides",
			"a little Modules and Packages",
		},
		Attendance: map[int]int{
			13: 0,
			14: 0,
			15: 0,
			16: 0,
			17: 0,
			18: 0,
		},
	},
	{
		Person: Person{
			ID:      21,
			Name:    "Alijon",
			Surname: "Boboyorov",
			Phone:   "987654320",
		},
		GroupID: 3,
		Grades: map[string]int{
			"week1": 89,
			"week2": 74,
			"week3": 33,
			"week4": 1,
		},
		TopicsToWorkOn: []string{
			"struct",
			"methods",
			"loop",
			"map",
		},
		Attendance: map[int]int{
			13: 0,
			14: 0,
			15: 0,
			16: 0,
			17: 0,
			18: 0,
		},
	},
	{
		Person: Person{
			ID:      22,
			Name:    "Khurshed",
			Surname: "Rahimov",
			Phone:   "987654321",
		},
		GroupID: 3,
		Grades: map[string]int{
			"week1": 96,
			"week2": 99,
			"week3": 33,
			"week4": 1,
		},
		TopicsToWorkOn: []string{
			"struct",
			"method",
			"interface",
		},
		Attendance: map[int]int{
			13: 0,
			14: 0,
			15: 0,
			16: 0,
			17: 0,
			18: 0,
		},
	},
	{
		Person: Person{
			ID:      23,
			Name:    "Mehrdod",
			Surname: "Rahmatov",
			Phone:   "987654322",
		},
		GroupID: 3,
		Grades: map[string]int{
			"week1": 96,
			"week2": 94,
			"week3": 33,
			"week4": 1,
		},
		TopicsToWorkOn: []string{
			"Parallel programming and Goroutines",
			"ORM",
			"Joining Tables",
			"Introduction to PostgreSQL",
		},
		Attendance: map[int]int{
			13: 0,
			14: 0,
			15: 0,
			16: 0,
			17: 0,
			18: 0,
		},
	},
	{
		Person: Person{
			ID:      24,
			Name:    "Muhammad",
			Surname: "Hoshimov",
			Phone:   "987654323",
		},
		GroupID: 3,
		Grades: map[string]int{
			"week1": 83,
			"week2": 56,
			"week3": 33,
			"week4": 1,
		},
		TopicsToWorkOn: nil,
		Attendance: map[int]int{
			13: 0,
			14: 0,
			15: 0,
			16: 0,
			17: 0,
			18: 0,
		},
	},
	{
		Person: Person{
			ID:      25,
			Name:    "Muhammad",
			Surname: "Khujaev",
			Phone:   "987654324",
		},
		GroupID: 3,
		Grades: map[string]int{
			"week1": 91,
			"week2": 88,
			"week3": 33,
			"week4": 1,
		},
		TopicsToWorkOn: []string{
			"Modules&Packages",
		},
		Attendance: map[int]int{
			13: 0,
			14: 0,
			15: 0,
			16: 0,
			17: 0,
			18: 0,
		},
	},
	{
		Person: Person{
			ID:      26,
			Name:    "Nozim",
			Surname: "Safarov",
			Phone:   "987654325",
		},
		GroupID: 3,
		Grades: map[string]int{
			"week1": 99,
			"week2": 90,
			"week3": 33,
			"week4": 1,
		},
		TopicsToWorkOn: nil,
		Attendance: map[int]int{
			13: 0,
			14: 0,
			15: 0,
			16: 0,
			17: 0,
			18: 0,
		},
	},
	{
		Person: Person{
			ID:      27,
			Name:    "Sunatullo",
			Surname: "Gafurov",
			Phone:   "987654326",
		},
		GroupID: 3,
		Grades: map[string]int{
			"week1": 88,
			"week2": 97,
			"week3": 33,
			"week4": 1,
		},
		TopicsToWorkOn: nil,
		Attendance: map[int]int{
			13: 0,
			14: 0,
			15: 0,
			16: 0,
			17: 0,
			18: 0,
		},
	},
	{
		Person: Person{
			ID:      28,
			Name:    "Tamim",
			Surname: "Orif",
			Phone:   "939132003",
		},
		GroupID: 3,
		Grades: map[string]int{
			"week1": 89,
			"week2": 71,
			"week3": 33,
			"week4": 1,
		},
		TopicsToWorkOn: []string{
			"Pointers",
			"Packages",
		},
		Attendance: map[int]int{
			13: 0,
			14: 0,
			15: 0,
			16: 0,
			17: 0,
			18: 0,
		},
	},
	{
		Person: Person{
			ID:      29,
			Name:    "Zohira",
			Surname: "Furmal",
			Phone:   "987654328",
		},
		GroupID: 3,
		Grades: map[string]int{
			"week1": 85,
			"week2": 74,
			"week3": 33,
			"week4": 1,
		},
		TopicsToWorkOn: nil,
		Attendance: map[int]int{
			13: 0,
			14: 0,
			15: 0,
			16: 0,
			17: 0,
			18: 0,
		},
	},
}
