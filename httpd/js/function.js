(function(){
    
    var loginUN    = document.getElementById('loginuser');
    var loginPass  = document.getElementById('loginpass');
    var inputToken = document.getElementById('tokenid'); 
    var regname    = document.getElementById('regname');
    var regpass    = document.getElementById('regpass');
    var memID      = document.getElementById('memID');
    var acceslvl   = document.getElementById('access');

    var loginbtn   = document.getElementById('login');
    var allbtn     = document.getElementById('getall');
    var regbtn     = document.getElementById('reg');

    loginbtn.addEventListener('click',sendlogininfo);
    allbtn.addEventListener('click', queryall)
    regbtn.addEventListener('click', register)

    function sendlogininfo(e) {
        //location.host
        console.log("HOLA!")
        var myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");
        var raw = JSON.stringify({"Name": loginUN.value,"Password":loginPass.value});
        var requestOptions = { method: 'POST', origin: "http://", headers: myHeaders, body: raw, redirect: 'follow'};
        fetch("http://localhost:8085/login/auth", requestOptions).then(function(response) {
                token = response.headers.get('Authorization'); 
                document.getElementById('dtoken').textContent=token
                inputToken.value=token
            });

        e.preventDefault()
    }
    function queryall(e) {
        var myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");
        myHeaders.append("Authorization", inputToken.value)
        var requestOptions = {method: 'GET', origin: "http://localhost", headers: myHeaders, redirect: 'follow'};
        fetch("http://localhost:8085/user/readall", requestOptions).then(function(response) {
            return response.json();
          }).then(function(myJson) {
            console.log(JSON.stringify(myJson,null, 2));
            document.getElementById('json').textContent = JSON.stringify(myJson,null, 2);
          });
        e.preventDefault()
    }
    function register(e){
        var myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");
        myHeaders.append("Authorization", inputToken.value)
        var raw = JSON.stringify([{"Name": regname.value,"Password":regpass.value, "MembershipID": parseInt(memID.value),"AccessLevelsID":parseInt(acceslvl.value)}]);
        console.log(raw)
        var requestOptions = {method: 'POST', origin: "http://localhost", headers: myHeaders, body: raw , redirect: 'follow'};
        fetch("http://localhost:8085/register/", requestOptions).then(
            response => response.text()
            ).then(result => document.getElementById('json').textContent = result).catch(error => console.log('error', error));
        e.preventDefault()
    }

}())