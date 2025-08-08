package choice4go

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

func TestChoiceCSD(t *testing.T) {
	libDir := "./dependency/libs"
	libName := "EMQuantAPI"
	user := "rdrk0006"
	pass := "ji848857"

	choice, err := NewChoice(
		libDir, libName, "",
	)
	if err != nil {
		t.Fatal(err)
	}

	if err = choice.Start(
		context.TODO(), user, pass,
		NewStartOptions().
			ForceLogin().
			LogLevel(slog.LevelDebug),
	); err != nil {
		t.Fatal(err)
	}
	defer choice.Stop()

	if results, err := choice.Csd(
		[]string{"000300.SH"},
		[]string{"OPEN", "CLOSE", "HIGH", "LOW", "VOLUME", "AMOUNT", "PRECLOSE", "CHANGE"},
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2024, 12, 31, 0, 0, 0, 0, time.Local),
		nil,
	); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range results.Iter() {
			t.Logf("%+v", v)
		}
	}

	// if results, err := choice.TradeDates(
	// 	time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
	// 	time.Date(2024, 12, 31, 0, 0, 0, 0, time.Local),
	// 	nil,
	// ); err != nil {
	// 	t.Fatal(err)
	// } else {
	// 	for _, v := range results.Iter() {
	// 		t.Logf("%+v", v)
	// 	}
	// }

	if results, err := choice.Csd(
		[]string{"000002.SZ", "300059.SZ"},
		[]string{"OPEN", "HIGH", "LOW", "CLOSE"},
		time.Date(2016, 1, 10, 0, 0, 0, 0, time.Local),
		time.Date(2016, 4, 13, 0, 0, 0, 0, time.Local),
		NewCsdOptions().
			Period(Daily).
			Adjust(NoAdjusted).
			Currency(CurrCNY).
			BondType(BondDirty),
	); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range results.Iter() {
			t.Logf("%+v", v)
		}
	}
}
