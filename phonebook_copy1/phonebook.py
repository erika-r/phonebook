
from binaryTree import *

class Phonebook():

    def __init__(self):
        self.nameTree = Tree()
        self.numberTree = Tree()

    def add(self,name,number,address):
        self.nameTree.add(name,number,address,"name")
        self.numberTree.add(name,number,address,"number")
    
    def find(self,value):
        if value.isdecimal():
            self.numberTree.find(self.numberTree.root,value,True)        #will return name
        else:
            self.nameTree.find(self.nameTree.root,value,True)       #will return number
    
    def remove(self,value):       
        if value.isdecimal():       #if value is a number 
            node = self.numberTree.find(self.numberTree.root,value,False)   #will be node with all info if not None
            if node is not None:
                name = node.name 
                self.numberTree.remove(self.numberTree.root,value)   #return name
                self.nameTree.remove(self.nameTree.root,name)
            else:
                print("{} not found, cannot be removed from contacts\n".format(value))
        else:
            node = self.nameTree.find(self.nameTree.root,value,False)  #will be node with all info if not None
            if node is not None:
                number = node.number
                self.nameTree.remove(self.numberTree.root,value)   #return name
                self.numberTree.remove(self.nameTree.root,number)
            else:
                print("{} not found, cannot be removed from contacts\n".format(value))
    
    def showAll(self):
        self.nameTree.printTree()        #only one needs to be printed since they are the same

