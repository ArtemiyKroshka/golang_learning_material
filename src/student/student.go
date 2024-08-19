package student

type Student interface {
	Study(string)
}

type Engineears struct {
	Subject string
}

type Doctors struct {
	Subject string
}

func (e *Engineears) Study(s string) {
	e.Subject = s
}

func (d *Doctors) Study(s string) {
	d.Subject = s
}
