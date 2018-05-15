const t1 = () => {
    var arrayF = []

    for (var i = 0; i < 4; i++) {
        arrayF[i] = function () {
            console.log(i)
        }
    }

    for (let f of arrayF) {
        f()
    }
}

const t2 = () => {
    for (let i = 0; i < 100000; i++) {
        console.log(i)
    }
}
