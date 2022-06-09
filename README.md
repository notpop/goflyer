# goflyer

golangでbitFlyer LightningのHTTP APIとRealtime APIを試しに叩いてみる<br>
https://lightning.bitflyer.com/docs

#### 実装したもの
```
/v1/me/getbalance
```
[getbalance詳細へ](https://lightning.bitflyer.com/docs#%E8%B3%87%E7%94%A3%E6%AE%8B%E9%AB%98%E3%82%92%E5%8F%96%E5%BE%97)<br><br>
```
/v1/ticker
```
[ticker詳細へ](https://lightning.bitflyer.com/docs#ticker)<br><br>
```
lightning_ticker_{product_code}
```
[lightning_ticker詳細へ](https://bf-lightning-api.readme.io/docs/realtime-ticker)<br><br>
```
/v1/me/sendchildorder
```
[sendchildorder詳細へ](https://lightning.bitflyer.com/docs#%E6%96%B0%E8%A6%8F%E6%B3%A8%E6%96%87%E3%82%92%E5%87%BA%E3%81%99)<br><br>
```
/v1/me/getchildorders
```
[getchildorders詳細へ](https://lightning.bitflyer.com/docs#%E6%B3%A8%E6%96%87%E3%81%AE%E4%B8%80%E8%A6%A7%E3%82%92%E5%8F%96%E5%BE%97)<br><br>
