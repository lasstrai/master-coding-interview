/**
 * @param {number[]} flowerbed
 * @param {number} n
 * @returns {boolean}
 */

const canPlaceFlowers = function(flowerbed, n) {
  let prev, next;
  for (let i = 0; i < flowerbed.length; i++) {
    if (flowerbed[i] === 0) {
      prev = i === 0 || flowerbed[i - 1] === 0;
      next = i === flowerbed.length - 1 || flowerbed[i + 1] === 0;
      if (prev && next) {
        flowerbed[i] = 1;
        n--;
      }
    }
  }
  return n === 0;
}

export default canPlaceFlowers;