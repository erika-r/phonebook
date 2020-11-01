
package binaryTree

import "fmt"

//node struct
type Bstnode struct {
    Name string   //will be used to arrange Bst
    Number string
    Address string
    left  *Bstnode
    right *Bstnode
}

//binary search tree struct
type Bst struct {
    root *Bstnode
}

//reset tree  resetting root value
func (b *Bst) Reset() {
    b.root = nil
}

//print whole tree
func (b *Bst) ShowBook() {
    b.ShowBookRec(b.root)
}

func (b *Bst) ShowBookRec(node *Bstnode) {
    if node != nil {
        b.ShowBookRec(node.left)
        
        name := node.Name
        address := node.Address
        number := node.Number
        fmt.Printf("\nName: %s\n",name)
        fmt.Printf("Address: %s\n",address)
        fmt.Printf("Number: %s\n",number)
        
        b.ShowBookRec(node.right)
    }
}

///////////////////// NUMBER TREE /////////////////////

//Insert number into tree  number
func (b *Bst) InsertNumber(name string, number string,address string) {
    b.InsertNumberRec(b.root,name,number,address)
}

//recursive InsertNumber
func (b *Bst) InsertNumberRec(node *Bstnode,name string, number string,address string) *Bstnode {
    if b.root == nil {
        b.root = &Bstnode{
            Name: name, Number: number, Address: address,
        }
        return b.root
    }
    if node == nil {
        return &Bstnode{Name: name,Number: number, Address: address,}
    }
    if number <= node.Number {
        node.left = b.InsertNumberRec(node.left, name,number,address)
    }
    if number > node.Number {
        node.right = b.InsertNumberRec(node.right, name,number,address)
    }
    return node
}


//Find value in tree using number
func (b *Bst) FindNumber(number string,print bool) string {    //return Name
    node := b.FindNumberRec(b.root, number)
    if node == nil {
        fmt.Printf("\nNumber %s not found\n",number)
        return ""
    } else {
        name := node.Name
        address := node.Address
        if print == true {
            fmt.Printf("\nNumber %s found\n", number)
            fmt.Printf("Name: %s\n",name)
            fmt.Printf("Address: %s\n",address)            
        }
        return name  
    }
}


//recursive FindNumber
func (b *Bst) FindNumberRec(node *Bstnode, number string) *Bstnode {
    if node == nil {
        return nil
    }
    if node.Number == number {
        return node
    }
    if number < node.Number {
        return b.FindNumberRec(node.left, number)
    }
    return b.FindNumberRec(node.right, number)
}


// remove number from tree 
func (b *Bst) RemoveNumber(number string) {
    RemoveNumberRec(b.root, number)
}


//recursive function to remove name
func RemoveNumberRec(node *Bstnode, number string) *Bstnode {
    if node == nil {
        return nil
    }
    if number < node.Number {
        node.left = RemoveNumberRec(node.left, number)
        return node
    }
    if number > node.Number {
        node.right = RemoveNumberRec(node.right, number)
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
    leftmostrightside := node.right
    for {
        //find smallest value on the right side
        if leftmostrightside != nil && leftmostrightside.left != nil {
            leftmostrightside = leftmostrightside.left
        } else {
            break
        }
    }
    node.Name = leftmostrightside.Name
    node.Number = leftmostrightside.Number
    node.Address = leftmostrightside.Address
    node.right = RemoveNumberRec(node.right, node.Number)
    return node
}



///////////////////// NAME TREE /////////////////////

//Insert name into tree  name
func (b *Bst) InsertName(name string, number string,address string) {
    b.InsertNameRec(b.root,name,number,address)
}

//recursive InsertName
func (b *Bst) InsertNameRec(node *Bstnode,name string, number string,address string) *Bstnode {
    if b.root == nil {
        b.root = &Bstnode{
            Name: name, Number: number, Address: address,
        }
        return b.root
    }
    if node == nil {
        return &Bstnode{Name: name,Number: number, Address: address,}
    }
    if name <= node.Name {
        node.left = b.InsertNameRec(node.left, name, number, address)
    }
    if name > node.Name {
        node.right = b.InsertNameRec(node.right, name, number, address)
    }
    return node
}


//find value in tree using name
func (b *Bst) FindName(name string,print bool) string {    //return number
    node := b.FindNameRec(b.root, name)
    if node == nil {
        fmt.Printf("\nName %s not found\n",name)
        return ""
    } else {
        number := node.Number
        address := node.Address
        if print == true {
            fmt.Printf("\n%s found in contacts\n", name)
            fmt.Printf("Number: %s\n",number)
            fmt.Printf("Address: %s\n",address)  
        }
        return number
    }
}


//recursive FindName
func (b *Bst) FindNameRec(node *Bstnode, name string) *Bstnode {
    if node == nil {
        return nil
    }
    if node.Name == name {
        return node
    }
    if name < node.Name {
        return b.FindNameRec(node.left, name)
    }
    return b.FindNameRec(node.right, name)
}



// remove name from tree 
func (b *Bst) RemoveName(name string) *Bstnode {
    return RemoveNameRec(b.root, name)
}


//recursive function to remove name
func RemoveNameRec(node *Bstnode, name string) *Bstnode {
    if node == nil {
        return nil
    }
    if name < node.Name {
        node.left = RemoveNameRec(node.left, name)
        return node
    }
    if name > node.Name {
        node.right = RemoveNameRec(node.right, name)
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
    leftmostrightside := node.right
    for {
        //find smallest value on the right side
        if leftmostrightside != nil && leftmostrightside.left != nil {
            leftmostrightside = leftmostrightside.left
        } else {
            break
        }
    }
    node.Name = leftmostrightside.Name
    node.Number = leftmostrightside.Number
    node.Address = leftmostrightside.Address
    node.right = RemoveNameRec(node.right, node.Name)
    return node
}