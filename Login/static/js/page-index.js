window.onload = function(){
	var ph = document.getElementById("phone");
	var phoneReg = /^1[3-9][0-9]{9}$/;
	var btn = document.getElementById("su");
	btn.onclick = function(){
		if(phoneReg.test(ph.value)){
			alert("正确");
		}
		else{
			alert("手机号输入有误");
		}
		
	}
}