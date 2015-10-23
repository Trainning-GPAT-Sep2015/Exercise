const Node=(function(){
	var valueSymbol=Symbol('value');
	var leftSymbol=Symbol('left');
	var rightSymbol=Symbol('right');

	function Node(value,left,right){
		this[valueSymbol]=value;
		this[leftSymbol]=left;
		this[rightSymbol]=right;
	};

	Node.prototype.getValue = function() {
		return this[valueSymbol];
	}

	Node.prototype.getLeft = function() {
		return this[leftSymbol];
	}

	Node.prototype.getRight = function() {
		return this[rightSymbol];
	}

	Node.prototype.setValue = function(value){
		this[valueSymbol]=value;
	}

	Node.prototype.setLeft = function(node){
		this[leftSymbol]=node;
	}

	Node.prototype.setRight = function(node){
		this[rightSymbol]=node;
	}

	return Node;
})();

function insert(root,value){
	if (root===null){
		return new Node(value,null,null);
	} else if (value<root.getValue()){
		root.setLeft(insert(root.getLeft(),value));
		return root;
	} else{
		root.setRight(insert(root.getRight(),value));
		return root;
	}
}

function inTraverse(root){
	if (root.getLeft()!==null){
		inTraverse(root.getLeft());
	}
	console.log(root.getValue());
	if (root.getRight()!==null){
		inTraverse(root.getRight());
	}
}

function *walk(root){
	if (root===null){
		return;
	}
	yield* walk(root.getLeft());
	yield root.getValue();
	yield* walk(root.getRight());
}

var tree1=null;
var tree2=null;
const a1=[6,1,7,10,9,4,3,12];
const a2=[7,6,9,3,4,12,1,10];

for (const val of a1){
	tree1=insert(tree1,val)
}
for (const val of a2){
	tree2=insert(tree2,val)
}
console.log("Infix BST:");
console.log("BST 1:");
inTraverse(tree1);
console.log("BST 2:")
inTraverse(tree2);

var i1=walk(tree1);
var i2=walk(tree2);
var j1=i1.next();
var j2=i2.next();
var check=(j1.value===j2.value);
while (check && !j1.done && !j2.done){
	j1=i1.next();
	j2=i2.next();
	check=(j1.value===j2.value);
}
if (i1.next().done!==i2.next().done){
	check=false;
}
console.log("Equivalent: "+check);