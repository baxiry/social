const url = "http://localhost:8080"

// putData put form data to server
function putData() {
    let elmInput = $("#update-name");

    var data = new FormData();
    data.append("username", elmInput.value);
    fetch(url+"/upacount", { method: "POST", body: data }); 
}

// createInput creates new input element
function creatInput(id) {
    let elmInput = $("#"+id);
    elmInput.innerHTML = '<input type="text" placeholder='+id+' id='+id+'>'
    elmInput.innerHTML += '<button class="btn btn-outline-primary" name=update >update</button>'
}

// $ my awsome javascript framework
function $(element) {
    return document.querySelector(element)
}
