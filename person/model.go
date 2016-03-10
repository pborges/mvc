package person

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Notes     []string
	errors    []error
}

func (this *Person) AddNote() {
	this.Notes = append(this.Notes, "")
}

func (this *Person) RemoveNote(id int) {
	this.Notes = append(this.Notes[:id], this.Notes[id + 1:]...)
}

func (this *Person) LeroyJenkins() {
	this.FirstName = "Leeerrooyy"
	this.LastName = "Jennnkinssssss"
}

func (this *Person)Errors() []error {
	return this.errors
}