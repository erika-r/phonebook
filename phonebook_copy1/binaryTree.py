class Node:
	
	def __init__(self, name,number,address):
		self.name = name
		self.number = number
		self.address = address
		self.left = None
		self.right = None

	def __str__(self):
		print("\nName: {}\nAddress: {}\nNumber: {}\n".format(self.name,self.address,self.number))

class Tree:
	# 
	def __init__(self):
		self.root = None

	#garbage collector will clean up for us
	def reset(self):
		self.root = None

	def printTree(self):
		if self.root is not None:
			self._printTree(self.root)

	def _printTree(self, node):
		if node is not None:
			self._printTree(node.left)
			print("\nName: {}\nAddress: {}\nNumber: {}\n".format(node.name,node.address,node.number))
			self._printTree(node.right)

	def find(self, node, value,status):
		while node is not None:

			#get appropriate value
			if value.isdecimal():
				node_value = node.number
			else:
				node_value = node.name
			

			#when value is found
			if value == node_value:
				if status:
					print(("\nName: {}\nAddress: {}\nNumber: {}\n".format(node.name,node.address,node.number)))
				return node

			#check left side
			elif value < node_value:
				node = node.left

			#check right side
			elif value > node_value:
				node = node.right
		
		#node not found
		if status:
			print("{} not found\n".format(value))
		return None
 

	#add based on number or name
	def add(self,name,number,address,sort):
		if self.root is None:
			self.root = Node(name,number,address)
		else:
			self._add(name,number,address,self.root,sort)

	def _add(self, name,number,address,node,sort):
		
		#define what to sort and compare by
		if sort == "number":
			value = number
			node_value = node.number
		elif sort == "name":
			value = name
			node_value = node.name

		if value < node_value:
			if node.left is not None:
				self._add(name,number,address,node.left,sort)
			else:
				node.left = Node(name,number,address)
		else:
			if node.right is not None:
				self._add(name,number,address,node.right,sort)
			else:
				node.right = Node(name,number,address)


	def remove(self,node,value):
		self._remove(self.root,value)
		# self.printTree()

	def _remove(self,node,value):
		#determine what to search by

		if node is not None:
			if value.isdecimal():
				node_value = node.number
			else:
				node_value = node.name
		else:
			return node

		if value < node_value:
			node.left = self._remove(node.left,value)
		
		elif value > node_value:
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

			tmp_value = node_value
			#node = Node(tmp.name,tmp.number,tmp.address)
			node.right = self._remove(node.right,tmp_value)
		return node

