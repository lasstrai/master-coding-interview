/**
 * @param {number[]} height
 * @returns {number}
 */

const maxArea = function(height) {
  let i = 0;
  let j = height.length - 1;
  let max = 0;
  while (i < j) {
    let calculation = (j - i) * Math.min(height[j], height[i]);
    if (calculation > max) {
      max = calculation;
    }
    if (height[i] < height[j]) {
      i++;
    }
    else {
      j--;
    }
  }
  return max;
}

export default maxArea;