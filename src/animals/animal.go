package animals

type Animal interface {
	Speak() string

	GetName() string
}

type Bird struct {
	Name string
}

func (b *Bird) Speak() string {
	return "chirp chirp!!"
}

func (b *Bird) GetName() string {
	return b.Name
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "woof woof"
}

func (d Dog) GetName() string {
	return d.Name
}
