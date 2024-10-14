/**
 * @param {number[]} flowerbed
 * @param {number} n
 * @returns {boolean}
 */

const canPlaceFlowers = function(flowerbed, n) {
  for (let i = 0; i < flowerbed.length; i++) {
    if (flowerbed[i] === 0) {
      if (flowerbed[i - 1] === 0 && flowerbed[i + 1] === 0) {
        flowerbed[i] = 1;
        n--;
      }
      if (i === 0 && flowerbed[i + 1] === 0) {
        flowerbed[i] = 1;
        n--;
      }
      if (i === flowerbed.length - 1 && flowerbed[i - 1] === 0) {
        flowerbed[i] = 1;
        n--;
      }
    }
  }
  return n ===  0;
}

export default canPlaceFlowers;