package folder_test

import (
	"testing"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

var def_org_id = uuid.FromStringOrNil(folder.DefaultOrgID)

var sampleData = []folder.Folder{
		{"creative-scalphunter",  def_org_id,  "creative-scalphunter"},
		{"clear-arclight", def_org_id,  "creative-scalphunter.clear-arclight"},
		{"bursting-lionheart", def_org_id,  "creative-scalphunter.clear-arclight.bursting-lionheart"},
		{"noble-vixen",  uuid.FromStringOrNil("somethingrandom"),  "noble-vixen"},
	}
var GetFoldersByOrgID_tests = [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{"incorrect orgID", def_org_id, sampleData, []folder.Folder{
				{"creative-scalphunter",  def_org_id,  "creative-scalphunter"},
				{"clear-arclight", def_org_id,  "creative-scalphunter.clear-arclight"},
				{"bursting-lionheart", def_org_id,  "creative-scalphunter.clear-arclight.bursting-lionheart"},

			},
		},

	}
var GetAllChildFolders_tests = [...]struct {
		name    string
		fname 	string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{"normal case", "creative-scalphunter", def_org_id, sampleData, []folder.Folder{
				{"clear-arclight", def_org_id,  "creative-scalphunter.clear-arclight"},
				{"bursting-lionheart", def_org_id,  "creative-scalphunter.clear-arclight.bursting-lionheart"},

			},
		},

	}
// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	for _, tt := range GetFoldersByOrgID_tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(get, tt.want)
			
		})
	}
}
func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	for _, tt := range GetAllChildFolders_tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			f := folder.NewDriver(tt.folders)
			get := f.GetAllChildFolders(def_org_id, tt.fname)
			assert.Equal(get, tt.want)
			
		})
	}
}