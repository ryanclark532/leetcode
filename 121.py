from typing import List


def maxProfit(prices: List[int]) -> int:
    buy_price = prices[0]
    profit = 0

    for n in prices[1:]:
        if n > buy_price:
            buy_price = n

        profit = max(profit, n-buy_price)
    return  profit



print(maxProfit([1,2, 4]))




