/**
 * @param {number[]} flowerbed
 * @param {number} n
 * @returns {boolean}
 */

const canPlaceFlowers = function(flowerbed, n) {
  let prev, next;
  let i = 0;
  while (i < flowerbed.length && n > 0) {
    if (flowerbed[i] === 0) {
      prev = i === 0 || flowerbed[i - 1] === 0;
      next = i === flowerbed.length - 1 || flowerbed[i + 1] === 0;
      if (prev && next) {
        flowerbed[i] = 1;
        n--;
      }
    }
    i++;
  }
  return n === 0;
}

export default canPlaceFlowers;