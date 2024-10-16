/**
 * @param {string} s
 * @return {string}
 */

function reverseVowels(s) {
  let i = 0;
  let j = s.length - 1;
  let vowels = new Map([
    ["a", true], 
    ["e", true],
    ["i", true], 
    ["o", true], 
    ["u", true]
  ]);
  while (i < j) {
    let cond1 = vowels.has(s[i].toLowerCase());
    let cond2 = vowels.has(s[j].toLowerCase());
    if (cond1 && cond2) {
      let temp = s[i];
      s[i] = s[j];
      s[j] = temp
      i++;
      j--;
    }
    if (!cond1) i++;
    if (!cond2) j--;
  }
  return s;
}

export default reverseVowels;