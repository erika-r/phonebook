#!/usr/bin/python
#FINAL

from binaryTree import *
from phonebook import *

def main():

    contacts = Phonebook()

    #adding contacts
    contacts.add("John","03495803","DCU")
    contacts.add("Erika","08123456","13 Fernwood Way")
    contacts.add("Tofu","08123457","15 Fernwood Way")
    contacts.add("Daisy","08123458","17 Fernwood Way")
    contacts.add("Ginger","03495813","DCU")
    contacts.add("Garlic","08124456","12 Fernwood Way")
    contacts.add("Onion","07123457","19 Fernwood Way")
    contacts.add("Pepper","08124458","18 Fernwood Way")

    print("-- All starting contacts --")
    contacts.showAll()

    #finding contacts WORKS
    print("--- Finding contacts ---")
    contacts.find("John")
    contacts.find("Pepper")
    contacts.find("08123456")
    contacts.find("07123457")

    #will not be found
    contacts.find("Rotten Tomatoes")
    

    print(". . . Removing contacts . . .\n")
    contacts.remove("08123457")     #tofu
    contacts.remove("08123458")     #daisy
    contacts.remove("08123456")     #erika
    contacts.remove("Onion")        #07123457
    
    #will not be removeable
    contacts.remove("Rotten Tomatoes")

    print("--- Showing all contacts ---")
    contacts.showAll()

if __name__ == '__main__':
    main()
