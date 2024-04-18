package links

import (
	"reflect"
	"testing"
)

func Test_localLinks(t *testing.T) {
	type args struct {
		source []byte
	}
	tests := []struct {
		name                  string
		args                  args
		wantLocalDestinations []string
	}{
		{
			name: "No local links",
			args: args{
				source: []byte("# Page C\n\nAll by myself. I don't link to any other page.\n"),
			},
			wantLocalDestinations: nil,
		},
		{
			name: "Multiple local links",
			args: args{
				source: []byte("# Page A\n\nCheck out [Page B](/pageB) and again [Page C](/pageC)!\n\nOk!?"),
			},
			wantLocalDestinations: []string{"/pageB", "/pageC"},
		},
		{
			name: "No images",
			args: args{
				source: []byte("[foo](/bar.jpg)"),
			},
			wantLocalDestinations: nil,
		},
		{
			name: "No external links",
			args: args{
				source: []byte("[foo](https://example.com/bar/foo)"),
			},
			wantLocalDestinations: nil,
		},
		{
			name: "Remove trailing slash",
			args: args{
				source: []byte("ok, here is [foo](/foo/)"),
			},
			wantLocalDestinations: []string{"/foo"},
		},
		// {
		// 	name: "Href link",
		// 	args: args{
		// 		source: []byte(`<a href="/test">foo</a>`),
		// 	},
		// 	wantLocalDestinations: []string{"/test"},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLocalDestinations := LocalLinks(tt.args.source)
			if !reflect.DeepEqual(gotLocalDestinations, tt.wantLocalDestinations) {
				t.Errorf("localLinks() = %#v, want %#v", gotLocalDestinations, tt.wantLocalDestinations)
			}
		})
	}
}
