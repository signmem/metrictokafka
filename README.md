# metrictokafka  

> 模拟 falcon metric 监控数据写入 kafka  
> 主要用于测试写入 falcon 主程序监控数据至 kafka  
> 不把监控数据直接写入 falcon-agent 避免过度依赖链路，无法排查链路故障信息   


# kafka 测试方法  
> 直接通过下面命令查询 kafka 消费记录   

```
2.2.2.2 falcon.falcon-alarm.health      port=2222       1713270376      1.000   GAUGE   60
3.3.3.3 falcon.falcon-hbs.health        port=3333       1713270381      1.000   GAUGE   60
1.1.1.2 falcon.falcon-api.health        port=1111       1713270381      1.000   GAUGE   60
3.3.3.2 falcon.falcon-hbs.health        port=3333       1713270381      1.000   GAUGE   60
1.1.1.1 falcon.falcon-api.health        port=1111       1713270386      1.000   GAUGE   60
1.1.1.2 falcon.falcon-api.health        port=1111       1713270386      1.000   GAUGE   60
2.2.2.2 falcon.falcon-alarm.health      port=2222       1713270386      1.000   GAUGE   60
2.2.2.3 falcon.falcon-alarm.health      port=2222       1713270386      1.000   GAUGE   60
3.3.3.1 falcon.falcon-hbs.health        port=3333       1713270386      1.000   GAUGE   60
```

# 写入条件  
> metric item 达到 70 个  
> 或满足一分钟一次 time.ticker   
