package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAMinusB(t *testing.T) {
	bp := &BashProfiler{}
	tests := []struct {
		Name             string
		Error            string
		A                []string
		B []string
		Res []string
	}{
		{
			Name: "Both lists populated",
			A: []string{"1","2","3","4","5"},
			B: []string{"1","3","8"},
			Res: []string{"2","4","5"},
		},
		{
			Name: "Only A is present",
			A: []string{"1","2","3","4","5"},
			Res: []string{"1","2","3","4","5"},
		},
		{
			Name: "Only A is populated",
			A: []string{"1","2","3","4","5"},
			B: []string{},
			Res: []string{"1","2","3","4","5"},
		},
		{
			Name: "Only B is present",
			B: []string{"1","2","3","4","5"},
			Res: []string{},
		},
		{
			Name: "Only B is populated",
			A: []string{},
			B: []string{"1","2","3","4","5"},
			Res: []string{},
		},
		{
			Name: "Shuffled",
			A: []string{"1","6","a","g","c","d"},
			B: []string{"6","1"},
			Res: []string{"a","c","d","g"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			res := bp.AMinusB(tt.A, tt.B)

			assert.Equal(t, len(res), len(tt.Res))
			for i := 0; i < len(res); i++ {
				assert.Equal(t, tt.Res[i], res[i])
			}
		})
	}
}

//func TestPull(t *testing.T) {
//	bp := &BashProfiler{}
//	tests := []struct {
//		Name             string
//		Error            string
//	}{
//
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.Name, func(t *testing.T) {
//			res := bp.AMinusB(tt.A, tt.B)
//
//			assert.Equal(t, len(res), len(tt.Res))
//			for i := 0; i < len(res); i++ {
//				assert.Equal(t, tt.Res[i], res[i])
//			}
//		})
//	}
//}