const state ={
 	items:[]
};

export default state;

export function deleteItem (index) {
	state.items.splice(index,1);

}

export function addItem(text){
	state.items.push(text);
	notify();
};

const callbacks =[];

export function subscribe(cb){
	callbacks.push(cb);
}

function notify(){
	callbacks.forEach((cb)=>{
		try{
			cb('update');
		}catch(e){
			console.log('Error',e);
		}
	});
}