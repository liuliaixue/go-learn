var arrayF =[]

for (var i = 0; i< 4 ; i++) {
    arrayF[i] = function() {
        console.log(i)
    }
}

for(let f of arrayF) {
    f()
}