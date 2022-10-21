class Roman {
    static ROMAN_NUMBER = { I: 1, V: 5, X: 10, L: 50, C: 100, D: 500, M: 1000 };
    constructor(s) {
        this.roman = new String(s);
        this.roman.charValue = function (index) {
            return Roman.ROMAN_NUMBER[this[index].toUpperCase()];
        };
    }
    toInt() {
        let r = this.roman;
        let result = 0;
        for (let i = 0; i < this.roman.length; i++) {
            if (typeof r[i + 1] === "undefined") {
                result += r.charValue(i);
                break;
            }
            if (r.charValue(i) < r.charValue(i + 1)) {
                result -= r.charValue(i);
                continue;
            }
            result += r.charValue(i);
        }
        return result;
    }
}
/**
 * @param {string} s
 * @return {number}
 */
let romanToInt = function (s) {
    let r = new Roman(s);
    return r.toInt();
};

console.log(romanToInt(""));
