window.onload = function(){
	var ph = document.getElementById("phone").value;
	var phoneReg = /^1[3-9][0-9]{9}$/;
	var btn = document.getElementById("su");
	btn.onclick = function(){
		if(!(phoneReg.test(ph))){
			alert("手机号输入有误");
			return false;
		}
		
	}
}