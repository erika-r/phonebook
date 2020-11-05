
class Contact:

	def __init__(self,name,number,address):
		self.name = name
		self.number = number
		self.address = address

	def __str__(self):
		print("\nName: {}\nAddress: {}\nNumber: {}\n".format(self.name,self.address,self.number))

class Node:
	
	def __init__(self, contact):
		self.contact = contact
		self.left = None
		self.right = None

class Tree:

	def __init__(self):
		self.root = None

	def printTree(self):
		if self.root is not None:
			self._printTree(self.root)

	def _printTree(self, node):
		if node is not None:
			self._printTree(node.left)
			print("\nName: {}\nAddress: {}\nNumber: {}\n".format(node.contact.name,node.contact.address,node.contact.number))
			self._printTree(node.right)

################ INHERITANCE ################
class NameTree(Tree):

	def find(self, node, name,status):
		while node is not None:
			
			#when name is found
			if name == node.contact.name:
				if status:
					print(("\nName: {}\nAddress: {}\nNumber: {}\n".format(node.contact.name,node.contact.address,node.contact.number)))
				return node

			#check left side
			elif name < node.contact.name:
				node = node.left

			#check right side
			elif name > node.contact.name:
				node = node.right
		
		#node not found
		if status:
			print("{} not found\n".format(name))
		return None


	def add(self,contact):
		if self.root is None:
			self.root = Node(contact)
		else:
			self._add(self.root,contact)

	def _add(self,node,contact):
		if contact.name < node.contact.name:
			if node.left is not None:
				self._add(node.left,contact)
			else:
				node.left = Node(contact)
		
		else:
			if node.right is not None:
				self._add(node.right,contact)
			else:
				node.right = Node(contact)


	def remove(self,node,value):
		self._remove(self.root,value)

	def _remove(self,node,value):
		if node is None:
			return node

		elif value < node.contact.name:
			node.left = self._remove(node.left,value)
		
		elif value > node.contact.name:
			node.right = self._remove(node.right,value)      

		else:    

			if not node.left:
				tmp = node.right
				node = None
				return tmp

			if not node.right:
				tmp = node.left
				node = None
				return tmp

			tmp = node.right
			while tmp and tmp.left:
				tmp = tmp.left

			tmp_value = node.contact.name
			#node = Node(tmp.name,tmp.number,tmp.address)
			node.right = self._remove(node.right,tmp_value)
		return node


class NumberTree(Tree):

	def find(self, node, number,status):
		while node is not None:
			
			#when number is found
			if number == node.contact.number:
				if status:
					print(("\nName: {}\nAddress: {}\nNumber: {}\n".format(node.contact.name,node.contact.address,node.contact.number)))
				return node

			#check left side
			elif number < node.contact.number:
				node = node.left

			#check right side
			elif number > node.contact.number:
				node = node.right
		
		#node not found
		if status:
			print("{} not found\n".format(number))
		return None
 

	def add(self,contact):
		if self.root is None:
			self.root = Node(contact)
		else:
			self._add(self.root,contact)

	def _add(self,node,contact):
		if contact.number < node.contact.number:
			if node.left is not None:
				self._add(node.left,contact)
			else:
				node.left = Node(contact)
		
		else:
			if node.right is not None:
				self._add(node.right,contact)
			else:
				node.right = Node(contact)


	def remove(self,node,value):
		self._remove(self.root,value)

	def _remove(self,node,value):
		if node is None:
			return node

		elif value < node.contact.number:
			node.left = self._remove(node.left,value)
		
		elif value > node.contact.number:
			node.right = self._remove(node.right,value)      

		else:    

			if not node.left:
				tmp = node.right
				node = None
				return tmp

			if not node.right:
				tmp = node.left
				node = None
				return tmp

			tmp = node.right
			while tmp and tmp.left:
				tmp = tmp.left

			tmp_value = node.contact.number
			#node = Node(tmp.name,tmp.number,tmp.address)
			node.right = self._remove(node.right,tmp_value)
		return node
