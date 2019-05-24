//The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility) 
//
// 
//P   A   H   N
//A P L S I I G
//Y   I   R
// 
//
// And then read line by line: "PAHNAPLSIIGYIR" 
//
// Write the code that will take a string and make this conversion given a number of rows: 
//
// 
//string convert(string s, int numRows); 
//
// Example 1: 
//
// 
//Input: s = "PAYPALISHIRING", numRows = 3
//Output: "PAHNAPLSIIGYIR"
// 
//
// Example 2: 
//
// 
//Input: s = "PAYPALISHIRING", numRows = 4
//Output: "PINALSIGYAHRPI"
//Explanation:
//
//P     I    N
//A   L S  I G
//Y A   H R
//P     I 
//

/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function(s, numRows) {
    var string = s;
    var num = numRows;
    var data = [];
    var col = true;
    var st = null;
    if (numRows === 1 || s === '') {
        return s;
    }
    var sort = getSort(numRows);
    for (var i = 0; i < string.length; i++) {
        st = sort[i % sort.length];
        if (!data[st]) {
            data[st] = [];
        }
        data[st].push(string[i])

    }
    var res = '';
    for (var j = 0; j < data.length; j++) {
        res += data[j].join('');
    }
    return res;
};

var getSort = (num) => {
    var data = [];
    for (var i = 0; i < 2 * (num - 1); i ++)  {
        if (i <= (num - 1)) {
            data[i] = i % num;
        } else {
            data[i] = num - (i % num) - 2;
        }
    }
    return data;
}