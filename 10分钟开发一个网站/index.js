const Date = (year,month) => {
    if (year <=0 || month <= 0) {
        return "error"
    }
    var result = []
    if (month === 12) {
        for (let i = 1;i <= 12; i++) {
            if (i >= 10) {
                result.push(year + '-'+ i)
            } else {
                result.push(year + '-'+ '0'+i)
            }
    }
    return result
        } else{
        for (let i = month; i <=12;i++ ){
            if (i >= 10) {
                result.push((year - 1) + '-'+ i)
            } else {
                result.push((year - 1) + '-'+ '0'+i)
            }
        }
            for (let i =1;i<= month;i ++ ){
                if (i >= 10) {
                    result.push(year + '-'+ i)
                } else {
                    result.push(year + '-'+ '0'+i)
                }

            }
                return result
        }

}

console.log(age(2015,12))
console.log(age(2015,11))
console.log(age(2015,10))
console.log(age(2015,9))
console.log(age(2015,8))
console.log(age(2015,7))
console.log(age(2015,6))

console.log(age(2015,5))
console.log(age(2015,4))
console.log(age(2015,3))
console.log(age(2015,2))

console.log(age(2015,1))
console.log(age(2015,0))
