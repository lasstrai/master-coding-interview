/**
 * @param {string} s
 * @returns {string}
 */

const reverseWords = function(s) {
  let word = "";
  let result = "";
  let flag = false;
  let i = 0;
  while (i < s.length) {
    if (s[i] !== " ") {
      word += s[i];
      flag = true;
    }
    else {
      result = word + result;
      word = "";
      if (flag) {
        result = " " + result;
        flag = false;
      }
    }
    i++;
  }
  return word + result;
}

console.log(reverseWords("the sky is blue"));

export default reverseWords;