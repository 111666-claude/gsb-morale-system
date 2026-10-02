# morale-system

士气系统：受伤扣减、指挥官减半、阵亡波及、按时间恢复与崩溃滞回，只用 Go 标准库。

```
go run ./cmd/morale-system --sample=break
go run ./cmd/morale-system --sample=rally
go run ./cmd/morale-system --sample=recover
go run ./cmd/morale-system --sample=kill
go run ./cmd/morale-system --sample=scan
go test ./...
go vet ./...
```

## 口径（README 为准）

- **受伤**：士气减去 `amount/10`（整除）；附近（`leaderRadius` 内）有别的存活指挥官时，
  这次扣减减半（整除）。
- **阵亡波及**：单位阵亡时，半径内的存活单位额外扣 15 点士气；半径外不扣。
- **钳制**：士气永远在 `[0,100]`。
- **自然恢复**：距上次受伤满 `recoverMs` 毫秒恢复 `recoverPer` 点，
  按真实时间结算，和 `Tick` 调用次数无关。
- **崩溃滞回**：士气不大于 `breakAt` 时进入崩溃；已经崩溃的单位只有士气不小于
  `rallyAt` 才恢复；`breakAt` 与 `rallyAt` 之间保持原状态。
- **代价**：范围查询不许遍历全部单位，`Scanned` 不随单位数线性增长。
- **规模**：10 万单位、每秒 10 万次结算，单次摊还 O(半径内单位数)。

## 输出契约（不改格式）

```
broken=.. morale=..
morale=..
b=..
scanned=..
```

## 目录

```
morale.go              受伤、阵亡、恢复与滞回
morale_test.go         单元测试
cmd/morale-system/     命令行入口
```
