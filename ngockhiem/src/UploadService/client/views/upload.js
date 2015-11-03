console.log("Hello");

var form = document.getElementById('uploadform');

form.addEventListener("submit", function(e) {
	e.preventDefault();
	var input = document.getElementById("file");
	if (input.files.length === 0) {

	} else {
		var file = input.files[0];
		console.log(file.type);
		if (file.type.match(/image.*/)) {
			xhr = new XMLHttpRequest();
			xhr.open("POST", "http://localhost:4000/upload");
			var formData = new FormData();
			formData.append("file", file);
			xhr.send(formData);
			xhr.onprogress = function(e) {
				if (e.lengthComputable) {
					var progress = document.getElementById("progress");
					var percentComplete = (e.loaded / e.total) * 100;
					progress.innerHTML = percentComplete + "%";
				} else {
					console.log("Cant compute");
				}
			}
			xhr.onload = function(e) {
				var img = document.getElementById("img");
				img.src = this.response;
				img.style.maxWidth = "300px";
				img.style.maxHeight = "300px";
				console.log(this.response);
			};
		} else {
			alert("Please choose an image file");
		}
	}
});