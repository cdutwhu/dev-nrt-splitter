package splitter

import (
	"testing"
)

func Test_csvSplitter(t *testing.T) {
	type args struct {
		csvfile    string
		outdir     string
		categories []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				csvfile: "../in/system_reports/systemDomainScores.csv",
				outdir:  "out",
				// categories: []string{"School", "Domain", "YrLevel"},
				categories: []string{"Domain", "School", "YrLevel"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := csvSplitter(tt.args.csvfile, tt.args.outdir, tt.args.categories...); (err != nil) != tt.wantErr {
				t.Errorf("csvSplitter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
