    def remove(self,node,value):
        self._remove(self.root,value)

    def _remove(self,node,value):
        #determine what to search by
        if value.isdecimal():
            node_value = node.number
        else:
            node_value = node.name
        
        if node is None:
            return None

        if value < node_value:
            node.left = self._remove(node.left,value)
            return node
        
        if value > node_value:
            node.right = self._remove(node.right,value)  
            return node      
            
        if not node.left and not node.right:
            node = None
            return node

        if not node.left:
           node = node.right
           return node

        if not node.right:
           node = node.left
           return node
        
        tmp = node.right

        while tmp and tmp.left:
            tmp = tmp.left

        node = Node(tmp.name,tmp.number,tmp.address)
        node.right = self._remove(node.right,node.number)

    def remove(self,node,value):
        self._remove(self.root,value)
        # self.printTree()

    def _remove(self,node,value):
        #determine what to search by
        if node is not None:
            # print("not none")
            if value.isdecimal():
                node_value = node.number
            else:
                node_value = node.name
        else:
            # print("None")
            return node

        if value < node_value:
            node.left = self._remove(node.left,value)
            return node
        
        if value > node_value:
            node.right = self._remove(node.right,value)  
            return node      
            
        if not node.left and not node.right:
            node = None
            return node

        if not node.left:
           node = node.right
           return node

        if not node.right:
           node = node.left
           return node
        
        tmp = node.right

        while tmp and tmp.left:
            tmp = tmp.left

        tmp_value = node_value
        node = Node(tmp.name,tmp.number,tmp.address)
        node.right = self._remove(node.right,tmp_value)


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
    def find(self, value,status):
        return self._find(value, self.root,status)

    def _find(self, value,node,status):

        if value.isdecimal():
            node_value = node.number
        else:
            node_value = node.name

        if value == node_value:
            if status:
                print(("\nName: {}\nAddress: {}\nNumber: {}\n".format(node.name,node.address,node.number)))
            return node
        
        elif (value < node_value and node.left is not None):
           self._find(value, node.left,status)
        elif (value > node_value and node.right is not None):
           self._find(value, node.right,status)

        if node.left is None and node.right is None:
            if status:
                print("{} not found\n".format(value))
            return node


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
