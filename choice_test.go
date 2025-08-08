package choice4go

import (
	"context"
	"testing"
	"time"
)

func TestChoiceCSD(t *testing.T) {
	libDir := "./dependency/libs"
	libName := "EMQuantAPI"
	user := "rdrk0006"
	pass := "ji848857"

	instance, err := NewChoice(
		libDir, libName, "",
	)
	if err != nil {
		t.Fatal(err)
	}

	if err = instance.Start(context.TODO(), user, pass); err != nil {
		t.Fatal(err)
	}
	defer instance.Stop()

	results, err := instance.Csd(
		[]string{"000300.SH"},
		[]string{"OPEN", "CLOSE", "HIGH", "LOW", "VOLUME", "AMOUNT", "PRECLOSE", "CHANGE"},
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2024, 12, 31, 0, 0, 0, 0, time.Local),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range results.Iter() {
		t.Logf("%+v", v)
	}
}
