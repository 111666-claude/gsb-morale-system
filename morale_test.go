package morale

import "testing"

func TestNewUnitIsFullMorale(t *testing.T) {
	field := New(50, 20, 30, 5, 60000)
	field.Add("u1", 0, 0, false)
	if field.Morale("u1") != 100 {
		t.Fatalf("新单位士气应该是 100：%v", field.Morale("u1"))
	}
}

func TestUnknownUnitMoraleIsZero(t *testing.T) {
	if New(50, 20, 30, 5, 60000).Morale("ghost") != 0 {
		t.Fatal("不存在的单位士气应该是 0")
	}
}

func TestUnknownUnitIsNotBroken(t *testing.T) {
	if New(50, 20, 30, 5, 60000).Broken("ghost") {
		t.Fatal("不存在的单位不该算崩溃")
	}
}

func TestDamageReducesMorale(t *testing.T) {
	field := New(50, 20, 30, 5, 60000)
	field.Add("u1", 0, 0, false)
	if got := field.Damage("u1", 100, 0); got != 90 {
		t.Fatalf("受伤 100 应该掉 10 点：%v", got)
	}
}

func TestScannedStartsAtZero(t *testing.T) {
	if New(50, 20, 30, 5, 60000).Scanned() != 0 {
		t.Fatal("Scanned 初始应该是 0")
	}
}
