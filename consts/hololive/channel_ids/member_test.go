package consts

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestMemberDefinitionsPreserveLegacyRouting(t *testing.T) {
	groups := []struct {
		name    string
		members []Member
	}{
		{"Whole", Whole},
		{"Gen0s", Gen0s},
		{"Gen1s", Gen1s},
		{"Gamerses", Gamerses},
		{"Gen2s", Gen2s},
		{"Gen3s", Gen3s},
		{"Gen4s", Gen4s},
		{"Gen5s", Gen5s},
		{"Holoxes", Holoxes},
		{"ReGrosses", ReGrosses},
		{"FlowGlows", FlowGlows},
		{"AsobiMawaritais", AsobiMawaritais},
		{"IdGen1s", IDGen1s},
		{"IdGen2s", IDGen2s},
		{"IdGen3s", IDGen3s},
		{"Mythes", Mythes},
		{"Promises", Promises},
		{"Jutsices", Jutsices},
		{"Advents", Advents},
		{"Councils", Councils},
	}

	data, err := os.ReadFile("testdata/legacy_members.tsv")
	if err != nil {
		t.Fatal(err)
	}
	wantByGroup := make(map[string][]Member)
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		fields := strings.Split(line, "\t")
		if len(fields) != 3 {
			t.Fatalf("invalid fixture row %q", line)
		}
		wantByGroup[fields[0]] = append(wantByGroup[fields[0]], Member{
			YouTubeChannelID: fields[1], DiscordChannelID: fields[2],
		})
	}
	if len(wantByGroup) != len(groups) {
		t.Fatalf("fixture groups = %d, definitions = %d", len(wantByGroup), len(groups))
	}
	for _, group := range groups {
		if !reflect.DeepEqual(group.members, wantByGroup[group.name]) {
			t.Errorf("%s routing changed: got %+v, want %+v", group.name, group.members, wantByGroup[group.name])
		}
	}
	if len(SearchGroups) != len(groups)-1 {
		t.Fatalf("search groups = %d, want %d", len(SearchGroups), len(groups)-1)
	}
	for i, group := range groups[:len(groups)-1] {
		if !reflect.DeepEqual(SearchGroups[i], group.members) {
			t.Errorf("search group %d differs from %s", i, group.name)
		}
	}
	allGroups := append(append([][]Member(nil), SearchGroups...), Councils)
	if err := ValidateMembers(allGroups); err != nil {
		t.Fatalf("member definitions are invalid: %v", err)
	}
}

func TestValidateMembers(t *testing.T) {
	valid := Member{YouTubeChannelID: "youtube-a", DiscordChannelID: "discord-a"}
	tests := []struct {
		name   string
		groups [][]Member
		want   string
	}{
		{"valid", [][]Member{{valid}, {{YouTubeChannelID: "youtube-b", DiscordChannelID: "discord-b"}}}, ""},
		{"empty YouTube ID", [][]Member{{{DiscordChannelID: "discord-a"}}}, "empty channel ID"},
		{"empty Discord ID", [][]Member{{{YouTubeChannelID: "youtube-a"}}}, "empty channel ID"},
		{"duplicate YouTube ID", [][]Member{{valid}, {{YouTubeChannelID: "youtube-a", DiscordChannelID: "discord-b"}}}, "duplicate YouTube channel ID"},
		{"duplicate Discord ID", [][]Member{{valid}, {{YouTubeChannelID: "youtube-b", DiscordChannelID: "discord-a"}}}, "duplicate Discord channel ID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMembers(tt.groups)
			if tt.want == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("error = %v, want %q", err, tt.want)
			}
		})
	}
}
