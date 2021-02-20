package pattern

import "testing"

func TestFactory(t *testing.T) {
	penFactory := GetPenFactory("pencil")
	if penFactory != nil {
		t.Log(penFactory.Create().Write())
	}
	penFactory = GetPenFactory("brushPen")
	if penFactory != nil {
		t.Log(penFactory.Create().Write())
	}
}