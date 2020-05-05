package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAMinusB(t *testing.T) {
	bp := &BashProfiler{}
	tests := []struct {
		Name             string
		Error            string
		A                []string
		B 				 []string
		Res 			 []string
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
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			res := bp.aMinusB(tt.A, tt.B)

			assert.Equal(t, len(res), len(tt.Res))
			for i := 0; i < len(res); i++ {
				assert.Equal(t, tt.Res[i], res[i])
			}
		})
	}
}

func TestSplitDeleted(t *testing.T){
	bp := &BashProfiler{}
	tests := []struct {
		Name             string
		Error            string
		RawBash          string
		DeletedArray 	 []string
		BashArray 		 []string
	}{
		{
			Name: "Normal section and deleted section",
			RawBash: "alias b=echo \"bash_profile\"\r\n" +
				"alias a=echo \"bash_profile\"\r\n" +
				"commandC() {\n\techo \"bash_profile\"\n}\r\n" +
				"commandD() {\n\techo \"bash_profile\"\n}\r\n" +
				"#Deleted\r\n" +
				"alias e=echo \"bash profile_delete\"\r\n" +
				"commandF() {\n\techo \"bash_profile_delete\"\n}",
			BashArray: []string{"alias b=echo \"bash_profile\"",
				"alias a=echo \"bash_profile\"",
				"commandC() {\n\techo \"bash_profile\"\n}",
				"commandD() {\n\techo \"bash_profile\"\n}",
			},
			DeletedArray: []string{"alias e=echo \"bash profile_delete\"",
				"commandF() {\n\techo \"bash_profile_delete\"\n}",
			},
		},
		{
			Name: "Normal section no deleted section",
			RawBash: "alias b=echo \"bash_profile\"\r\n" +
				"alias a=echo \"bash_profile\"\r\n" +
				"commandC() {\n\techo \"bash_profile\"\n}\r\n" +
				"commandD() {\n\techo \"bash_profile\"\n}",
			BashArray: []string{"alias b=echo \"bash_profile\"",
				"alias a=echo \"bash_profile\"",
				"commandC() {\n\techo \"bash_profile\"\n}",
				"commandD() {\n\techo \"bash_profile\"\n}",
			},
		},
		{
			Name: "No Normal section deleted section present",
			RawBash: "#Deleted\r\n" +
				"alias e=echo \"bash profile_delete\"\r\n" +
				"commandF() {\n\techo \"bash_profile_delete\"\n}",
			DeletedArray: []string{"alias e=echo \"bash profile_delete\"",
				"commandF() {\n\techo \"bash_profile_delete\"\n}",
			},
		},
		{
			Name: "No Normal section and no deleted section",
			BashArray: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			bashRes, deletedRes, err := bp.splitDeleted(tt.RawBash)
			if err != nil {
				assert.Equal(t, tt.Error, err.Error())
				return
			}

			assert.Equal(t, len(tt.BashArray), len(bashRes))
			for i := 0; i < len(tt.BashArray); i++ {
				assert.Equal(t, tt.BashArray[i], bashRes[i])
			}

			assert.Equal(t, len(tt.DeletedArray), len(deletedRes))
			for i := 0; i < len(tt.DeletedArray); i++ {
				assert.Equal(t, tt.DeletedArray[i], deletedRes[i])
			}

			assert.Equal(t, true, tt.Error=="")
		})
	}
}

