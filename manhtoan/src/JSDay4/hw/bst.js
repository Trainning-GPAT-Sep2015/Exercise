class BinaryTree {
    constructor(value) {
        this.root = new Node(value);
    }

    insertNode(value) {
        this._insertNodeRecursive(this.root, value);
    }

    _insertNodeRecursive(subroot, value) {
        if (value < subroot.value) {
            if (subroot.left === null) {
                subroot.left = new Node(value);
            } else {
                this._insertNodeRecursive(subroot.left, value);
            }
        } else if (value > subroot.value) {
            if (subroot.right === null) {
                subroot.right = new Node(value);
            } else {
                this._insertNodeRecursive(subroot.right, value);
            }
        }
    }

    * traverseLNR(subroot) {
        if (subroot.left !== null) {
            yield * this.traverseLNR(subroot.left);
        }
        yield subroot.value;
        if (subroot.right !== null) {
            yield * this.traverseLNR(subroot.right);
        }
    }

}

class Node {
    constructor(value) {
        this.left = null;
        this.right = null;
        this.value = value;
    }
}

// var newnode = new Node(null, 2, null);

var tree1 = new BinaryTree(2);

// tree.insertNode(2);
tree1.insertNode(1);
tree1.insertNode(3);

var tree2 = new BinaryTree(2);

// tree.insertNode(2);
tree2.insertNode(1);
tree2.insertNode(4);


var tra1 = tree1.traverseLNR(tree1.root);
var tra2 = tree2.traverseLNR(tree2.root);

var flag = true;

// console.log(tra1.next().value);
// console.log(tra2.next().value);
// console.log(tra1.next().value);
// console.log(tra2.next().value);
// console.log(tra1.next().value);
// console.log(tra2.next().value);
// console.log(tra1.next().value);
// console.log(tra2.next().value);

console.log("Start compare");

for(var i = 0; i < 3; i++) {
    if(tra1.next().value !== tra2.next().value) {
        flag = false;
        break;
    }
}

if (flag) {
    console.log("Equal");
} else {
    console.log("Not Equal");
}