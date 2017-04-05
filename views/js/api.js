function checkError(error){
    if (error.response.status === 401){
        self.location = "/login"
    }else if (error.response.status === 500){
        alert("服务器错误，请稍候重试！")
    }
}

function getParams(){
    
}