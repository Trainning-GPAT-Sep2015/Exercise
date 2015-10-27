// Input string
var string = '5 8 2 6 3 0 1';
// Templet array to store input string
var arr = string.split(' ');
// Struct of a node of a binary tree
function node(value){
	this.value = value;
	this.left = null;
	this.right = null;
};
//Struct of a binary tree
function BinaryTree() {
	this.root = null;	
};
//Create a binary tree
BinaryTree.prototype.add = function(value) {
	//current root of the binary tree
	var root = this.root;
	//if the current binary tree is null, set the value as the root of the tree
	if(!root) {
		//create new node with the value input
		//the set the root of the bTree as this new node
		this.root = new node(value);
		return;
	}
	var current = root;
	var newNode = new node(value);

	while(current) {
		if(value < current.value) {
			if(!current.left) {	// if left child is not null, set new left child
				current.left = newNode;
				break;
			} else {
				current = current.left;
			}
		} else {
			if(!current.right) {	// if right child is not null, set new right child
				current.right = newNode;
				break;
			} else {
				current = current.right;
			}
		}	
	}
};

function inorder(node){
	// if(node) {
	// 	return '(' + inorder(node.left) + node.value.toString() + inorder(node.right) + ')';
	// }	
	var left;
	var right;
	if (node){
		if(node.left) {
			left = '(' + inorder(node.left) + ')';
		}
		else {
			left = '';
		}
		if(node.right) {
			right = '(' + inorder(node.right) + ')';
		}
		else {
			right = '';
		}
		return left + node.value + right;
	}
}

BinaryTree.prototype.inorderTraverse = function() {
	if(this.root) {
		return inorder(this.root);
	}
	else {
		return;
	}
}

// Print out the input array
console.log('Input array: ');
for (var i in arr) {
	console.log(i + ': ' + arr[i]);
};

// create a binary tree
var bTree = new BinaryTree(null);
for (var i in arr) {
	bTree.add(arr[i]);
};
//var arr = [];
//Print out binary tree
console.log(bTree);
console.log(bTree.inorderTraverse());
//console.log(arr);