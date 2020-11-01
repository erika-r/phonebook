
package phonebook

import (
	"phonebook_copy2/binaryTree"
	"strconv"
)

type Phonebook struct
{
	nameTree binaryTree.Bst
	numberTree binaryTree.Bst
}

func (p *Phonebook) Add(number string, name string,address string) {
	p.nameTree.InsertName(name,number,address)
	p.numberTree.InsertNumber(name,number,address)
}

//true used to decide when to print values or not
func (p *Phonebook) Find(value string) {
	if _, err := strconv.Atoi(value); err == nil {
		p.numberTree.FindNumber(value,true)
	} else {
		p.nameTree.FindName(value,true)
	}
}

func (p *Phonebook) Remove(value string) {
	if _, err := strconv.Atoi(value); err == nil {
		name := p.numberTree.FindNumber(value,false)
		p.numberTree.RemoveNumber(value)
		p.nameTree.RemoveName(name)
	} else {
		number := p.numberTree.FindName(value,false)
		p.nameTree.RemoveName(value)
		p.nameTree.RemoveNumber(number)
	}
}

func (p *Phonebook) ShowAll() {
	p.nameTree.ShowBook()
}