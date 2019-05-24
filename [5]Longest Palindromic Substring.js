//Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000. 
//
// Example 1: 
//
// 
//Input: "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.
// 
//
// Example 2: 
//
// 
//Input: "cbbd"
//Output: "bb"
// 
//

var getLongest = function (s, i) {
    var str = s[i];
    var diff = 1;
    while (s[i - diff] && s[i + diff] && s[i - diff] === s[i + diff]) {
        str = s[i + diff] + str + s[i + diff];
        diff++;
    }
    var diff = 1;
    var str2 = '';
    while (s[i - diff + 1] && s[i + diff] && s[i - diff + 1] === s[i + diff]) {
        str2 = s[i + diff] + str2 + s[i + diff];
        diff++;
    }
    return str.length > str2.length ? str : str2;
}

/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function(s) {
    var target;
    var string = '';
    for (var i = 0; i < s.length; i++) {
        var temp = getLongest(s, i);
        if (temp.length > string.length) {
            string = temp;
        }
        // if (string.length > s.length - i) {
        //     return string;
        // }
    }
    return string;

};
longestPalindrome('cbbd')