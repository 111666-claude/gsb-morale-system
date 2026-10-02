// Package morale 是士气系统：受伤扣减、指挥官减半、阵亡波及、按时间恢复与崩溃滞回。
// 缺陷：崩溃判定用严格小于、没有滞回、恢复按 Tick 次数、阵亡波及不看距离、
// 士气不钳制、范围查询扫描全部单位。
package morale

import "math"

// Unit 是一个带士气的单位。
type Unit struct {
	ID     string
	X      float64
	Y      float64
	Leader bool
	Alive  bool
	Morale int
	Broken bool
}

// Field 是士气场。
type Field struct {
	radius      float64
	breakAt     int
	rallyAt     int
	recoverPer  int
	recoverMs   int64
	units       map[string]*Unit
	lastRecover map[string]int64
	scanned     int
}

// New 建士气场。
func New(leaderRadius float64, breakAt, rallyAt, recoverPer int, recoverMs int64) *Field {
	return &Field{radius: leaderRadius, breakAt: breakAt, rallyAt: rallyAt,
		recoverPer: recoverPer, recoverMs: recoverMs,
		units: map[string]*Unit{}, lastRecover: map[string]int64{}}
}

// Add 加一个单位。
func (f *Field) Add(id string, x, y float64, leader bool) {
	f.units[id] = &Unit{ID: id, X: x, Y: y, Leader: leader, Alive: true, Morale: 100}
}

// Damage 造成一次伤害，返回当前士气。
// 缺陷：不看附近指挥官、士气不钳制、崩溃状态不更新。
func (f *Field) Damage(id string, amount int, nowMs int64) int {
	unit := f.units[id]
	if unit == nil {
		return 0
	}
	unit.Morale -= amount / 10
	f.lastRecover[id] = nowMs
	return unit.Morale
}

// Heal 恢复一部分士气，返回当前士气。
func (f *Field) Heal(id string, amount int) int {
	unit := f.units[id]
	if unit == nil {
		return 0
	}
	unit.Morale += amount
	return unit.Morale
}

// Kill 让一个单位阵亡，返回被波及的单位数。
// 缺陷：不看距离，所有存活单位都扣。
func (f *Field) Kill(id string, nowMs int64) int {
	f.scanned += len(f.units)
	victim := f.units[id]
	if victim == nil || !victim.Alive {
		return 0
	}
	victim.Alive = false
	affected := 0
	for _, unit := range f.units {
		if !unit.Alive {
			continue
		}
		unit.Morale -= 15
		affected++
	}
	return affected
}

// Tick 结算恢复。
// 缺陷：每次调用都加固定值，和真实时间无关。
func (f *Field) Tick(nowMs int64) {
	for _, unit := range f.units {
		if unit.Alive {
			unit.Morale += f.recoverPer
		}
	}
}

// Morale 是当前士气。
func (f *Field) Morale(id string) int {
	if unit := f.units[id]; unit != nil {
		return unit.Morale
	}
	return 0
}

// Broken 是当前是否崩溃。
// 缺陷：用严格小于判定，没有滞回。
func (f *Field) Broken(id string) bool {
	unit := f.units[id]
	if unit == nil {
		return false
	}
	return unit.Morale < f.breakAt
}

// Distance 是两个单位的直线距离（样例核对用）。
func (f *Field) Distance(a, b string) float64 {
	first, second := f.units[a], f.units[b]
	if first == nil || second == nil {
		return 0
	}
	return math.Hypot(first.X-second.X, first.Y-second.Y)
}

// Scanned 是累计扫描量（规模观测）。
func (f *Field) Scanned() int { return f.scanned }
