class Node{
	constructor(value=0,left=null,right=null){
		this.value=value;
		this.left=left;
		this.right=right;
	}
}

function insert(root,value){
	if (root===null){
		root=new Node(value);
		return root;
	} else if (value<root.value){
		root.left=insert(root.left,value);
		return root;
	} else{
		root.right=insert(root.right,value);
		return root;
	}

}

function inTraverse(root){
	if (root.left!==null){
		inTraverse(root.left);
	}
	console.log(root.value);
	if (root.right!==null){
		inTraverse(root.right);
	}
}

function treeString(root){
	if (root===null){
		return '';
	} else{
		const st_left=root.left===null ? '' : '('+treeString(root.left)+')';
		const st_right=root.right===null ? '' : '('+treeString(root.right)+')';
		return st_left+root.value.toString()+st_right;
	}
}

function getLevel(node,root){
	if (node===root || node===null){
		return 0;
	} else if (node.value<root.value){
		return 1+getLevel(node,root.left);
	} else{
		return 1+getLevel(node,root.right);
	}
}

function inTraverseBonus(node,root){
	var st='';
	for (var i=0;i<getLevel(node,root);i++){
		st=st+'-';
	}
	st=st+node.value.toString();
	console.log(st);
	if (node.left!==null){
		inTraverseBonus(node.left,root);
	}
	if (node.right!==null){
		inTraverseBonus(node.right,root);
	}
}

var treeRoot=null;
const n=10;
var a=Array(n);
for (var i=0;i<n;i++){
	a[i]=Math.floor(Math.random()*20)+1;
	treeRoot=insert(treeRoot,a[i]);
}

console.log("Insertions to BST:");
console.log(a);
console.log("Infix traversal:");
console.log(treeString(treeRoot));
console.log("Infix traversal bonus:");
inTraverseBonus(treeRoot,treeRoot);