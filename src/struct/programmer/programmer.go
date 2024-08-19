package programmer

type Programmer struct {
	Name                string
	Age                 int
	salary              float64
	ProgrammingLanguage string
	Company
}

type Company struct {
	Name string
}

func NewProgrammer(name string, age int, salary float64, proglang string, company Company) *Programmer {
	return &Programmer{Name: name, Age: age, salary: salary, ProgrammingLanguage: proglang, Company: company}
}
