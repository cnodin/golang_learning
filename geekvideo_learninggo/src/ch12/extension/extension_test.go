package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
	
}

func (p *Pet) Speak()  {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {

}

func (p *Dog) Speak()  {
	fmt.Println("Wang")
}

func TestExtension(t *testing.T) {
	p := new(Dog)
	p.Speak()
}
