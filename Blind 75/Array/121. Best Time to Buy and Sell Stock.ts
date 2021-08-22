function maxProfit(prices: number[]): number {
  let buyPrice = prices[0];
  let highestProfit = 0;

  for (let price of prices) {
    if (price < buyPrice) {
      buyPrice = price;
    } else {
      highestProfit = Math.max(price - buyPrice, highestProfit);
    }
  }
  return highestProfit;
}
