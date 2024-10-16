/**
 * @param {string} s
 * @param {string} t
 * @returns {boolean}
 */

const backspaceCompare = function(s, t) {
  
  function nextValidIndex(str, index) {
    let backSpace = 0;
    while (index >= 0) {
      if (str.charAt(index) === "#") {
        backSpace++;
        index--;
      }
      else if (backSpace > 0 ) {
        backSpace--;
        index--;
      }
      else {
        return index;
      }
    }
    return -1;
  }

  let i = s.length - 1;
  let j = t.length - 1;
  while (i >= 0 || j >= 0) {
    i = nextValidIndex(s, i);
    j = nextValidIndex(t, j);
    if (i>= 0 && j >= 0 & s.charAt(i) !== t.charAt(j)) {
      return false;
    }
    if ((i >= 0 && j < 0) || (i < 0 && j >= 0)) {
      return false;
    } 
    i--;
    j--;
  }
  return true;
}

export default backspaceCompare;