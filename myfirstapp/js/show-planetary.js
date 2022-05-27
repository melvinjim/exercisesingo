planetInformation()

function planetInformation(){
    let url = 'http://localhost:3000/planetary';

    const api = new XMLHttpRequest();
    api.open('GET', url, true);
    api.send();

    api.onreadystatechange = function(){
        if (this.status == 200 && this.readyState == 4) {
            let datos = JSON.parse(this.responseText);
            document.getElementById("date").innerHTML = "Date: " +  datos.date;
            document.getElementById("explanation").innerHTML = "Explanation: " +  datos.explanation;
            document.getElementById("hdurl").innerHTML = "Hdurl: " +  datos.hdurl;
            document.getElementById("mediaType").innerHTML = "MediaType: " +  datos.media_type;
            document.getElementById("serviceVersion").innerHTML = "ServiceVersion: " +  datos.service_version;
            document.getElementById("title").innerHTML = "Title: " +  datos.title;
            document.getElementById("URL").innerHTML = "URL: " +  datos.url;
        }
    }
}