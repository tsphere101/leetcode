const ROMAN_NUMBER = { I: 1, V: 5, X: 10, L: 50, C: 100, D: 500, M: 1000 };

/**
 * @param {string} s
 * @return {number}
 */
let romanToInt = function (s) {
    let result = 0;
    for (let i = 0; i < s.length; i++) {
        if (typeof s[i + 1] === "undefined") return result + ROMAN_NUMBER[s[i]];
        if (ROMAN_NUMBER[s[i]] < ROMAN_NUMBER[s[i + 1]]) {
            result -= ROMAN_NUMBER[s[i]];
            continue;
        }
        result += ROMAN_NUMBER[s[i]];
    }

    return result;
};
