//22:28pm 1/11/2020
package main

import . "phonebook_copy2/phonebook"
import "fmt"


func main() {

	//NAME TREE
    contacts := &Phonebook{}   //create tree

    contacts.Add("087123456789", "John Doe","DCU")
    contacts.Add("087987654321", "Jane Doe","DCU")
    contacts.Add("087246810121", "John Smith","DCU")   
    contacts.Add("087369121518", "Jane Smith","DCU")

    fmt.Printf("\n-- Showing Contacts --\n")
    contacts.ShowAll()

    fmt.Printf("\n-- Finding values --\n")
    contacts.Find("John Doe")
    contacts.Find("Jane Doe")
    contacts.Find("087369121518")
    contacts.Find("Unknown")  //will not be found

    fmt.Printf("\n-- Removing Values --\n")
    contacts.Remove("087369121518")
    contacts.Remove("087369121513")

    fmt.Printf("\n-- Showing Contacts --\n")
    contacts.ShowAll() //works

}