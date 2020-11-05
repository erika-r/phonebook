
package main

import "fmt"
import "strconv"

//binary search tree struct
type Bst struct {
    root *Bstnode
}

type Bstnode struct {
    person *Contact
    left *Bstnode
    right *Bstnode
}

type Contact struct {
    Name string
    Number string
    Address string
}

type Phonebook struct {
	nameTree *Bst
	numberTree *Bst
}


// print whole tree
func (b *Bst) ShowBook() {
    b.ShowBookRec(b.root)
}

func (b *Bst) ShowBookRec(node *Bstnode) {
    if node != nil {
        b.ShowBookRec(node.left)
        
        name := node.person.Name
        address := node.person.Address
        number := node.person.Number
        fmt.Printf("\nName: %s\n",name)
        fmt.Printf("Address: %s\n",address)
        fmt.Printf("Number: %s\n",number)
        
        b.ShowBookRec(node.right)
    }
}



//Insert into tree, sortBy differentiates between name and number
func (b *Bst) Insert(being *Contact,sortBy string) {
    b.InsertRec(b.root,being,sortBy)
}

//recursive Insert
func (b *Bst) InsertRec(node *Bstnode,being *Contact,sortBy string) *Bstnode {
    if b.root == nil {
        b.root = &Bstnode{person:being}
        return b.root
    }
    if node == nil {
        return &Bstnode{person:being}
    }

    //use sortBy
    if sortBy == "name" {
        if being.Name < node.person.Name {
            node.left = b.InsertRec(node.left,being,"name")
        } else {
            node.right = b.InsertRec(node.right,being,"name")
        }
    } else {
        if being.Number < node.person.Number {
            node.left = b.InsertRec(node.left,being,"number")
        } else {
            node.right = b.InsertRec(node.right,being,"number")
        }
    }
    return node
}



//Find value in tree, sortBy tells the function what to sort by
func (b *Bst) Search(value string,sortBy string,print bool) string {    //return Name
    node := b.SearchRec(b.root,value,sortBy)
    if node == nil {
        fmt.Printf("\n%s not found\n",value)
        return ""
    } else {
        name := node.person.Name
        number := node.person.Number
        address := node.person.Address
        if print == true {
            fmt.Printf("\n%s found\n", value)
            fmt.Printf("Name: %s\n",name)
            fmt.Printf("Number: %s\n",number)
            fmt.Printf("Address: %s\n",address)            
        }

        if sortBy == "name" {
            return node.person.Number  
        } else {
            return node.person.Name
        }
    }
}


//recursive Search
func (b *Bst) SearchRec(node *Bstnode, value string,sortBy string) *Bstnode {
    if node == nil {
        return nil
    }

    if sortBy == "name" {
        if node.person.Name == value {
            return node
        } else if node.person.Name > value {
            return b.SearchRec(node.left,value,"name")
        } else if node.person.Name < value {
            return b.SearchRec(node.right,value,"name")
        }
    } else if sortBy == "number" {
        if node.person.Number == value {
            return node
        } else if node.person.Number > value {
            return b.SearchRec(node.left,value,"number")
        } else if node.person.Number < value {
            return b.SearchRec(node.right,value,"number")
        }
    }
    return node
}


// remove value from tree 
func (b *Bst) Remove(value string,sortBy string) {
    b.RemoveRec(b.root, value,sortBy)
}


//recursive function to remove value
func (b *Bst) RemoveRec(node *Bstnode,value string,sortBy string) *Bstnode {
    if node == nil {
        return nil
    } else {
        var nodeVal string

        if sortBy == "name" {
            nodeVal = node.person.Name
        } else if sortBy == "number" {
            nodeVal = node.person.Number
        }

        if value < nodeVal {
            node.left = b.RemoveRec(node.left, value,sortBy)
            return node
        }
        if value > nodeVal {
            node.right = b.RemoveRec(node.right, value,sortBy)
            return node
        }
        if node.left == nil && node.right == nil {
            node = nil
            return nil
        }
        if node.left == nil {
            node = node.right
            return node
        }
        if node.right == nil {
            node = node.left
            return node
        }
        tmp := node.right
        //find smallest value on the right side
        for {   //infinite loop until it finds the right node
            if tmp != nil && tmp.left != nil {
                tmp = tmp.left
            } else {
                break
            }
        }
        
        node.person.Name = tmp.person.Name
        node.person.Number = tmp.person.Number
        node.person.Address = tmp.person.Address

        if sortBy == "name" {
            node.right = b.RemoveRec(node.right, node.person.Name,"name")
        } else {
            node.right = b.RemoveRec(node.right, node.person.Number,"number")
        }

        return node
    }

}

// PHONEBOOK FUNCTIONS 
func (p *Phonebook) Add(number string, name string,address string) {
    being := &Contact{Name:name,Number:number,Address:address}
    p.nameTree.Insert(being,"name")
	p.numberTree.Insert(being,"number")
}


func (p *Phonebook) ShowAll() {
	p.nameTree.ShowBook()       //only one needs to be shown
}


//true used to decide when to print values or not
func (p *Phonebook) Find(value string) {
	if _, err := strconv.Atoi(value); err == nil {
		p.numberTree.Search(value,"number",true)
	} else {
		p.nameTree.Search(value,"name",true)
	}
}

func (p *Phonebook) Delete(value string) {
	if _, err := strconv.Atoi(value); err == nil {
		name := p.numberTree.Search(value,"number",false)
		p.numberTree.Remove(value,"number")
		p.nameTree.Remove(name,"name")
	} else {
		number := p.numberTree.Search(value,"name",false)
		p.nameTree.Remove(value,"name")
		p.nameTree.Remove(number,"number")
	}
}



func main() {

	//NAME TREE
    contacts := &Phonebook{nameTree:&Bst{},numberTree:&Bst{}}   //create tree

    contacts.Add("087123456789", "John Doe","ITT")
    contacts.Add("089987654321", "Jane Doe","DIT")
    contacts.Add("087246810121", "John Smith","ITB")   
	contacts.Add("089369121518", "Jane Smith","Maynooth")
	contacts.Add("088456789012", "Johnny","UCD")
	contacts.Add("082654321987", "Erika","DCU")

    fmt.Printf("\n-- Showing Contacts --\n")
    contacts.ShowAll()

    fmt.Printf("\n-- Finding values --\n")
    contacts.Find("John Doe")
    contacts.Find("Jane Doe")
    contacts.Find("087369121518")
    
    //will not be found
    contacts.Find("Unknown")  

    fmt.Printf("\n-- Deleting Values --\n")
	contacts.Delete("087369121518")
	
	//will not be found
    contacts.Delete("087369121513")

    fmt.Printf("\n-- Showing Contacts --\n")
    contacts.ShowAll()

}
