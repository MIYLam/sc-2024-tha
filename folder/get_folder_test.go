package folder_test

import (
	"errors"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

var def_org_id = uuid.FromStringOrNil(folder.DefaultOrgID)

var sampleData = []folder.Folder{
	{"creative-scalphunter", def_org_id, "creative-scalphunter"},
	{"clear-arclight", def_org_id, "creative-scalphunter.clear-arclight"},
	{"topical-micromax", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax"},
	{"bursting-lionheart", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart"},
	{"striking-black-panther", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.striking-black-panther"},
	{"advanced-professor-monster", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.advanced-professor-monster"},
	{"assuring-red-shift", uuid.FromStringOrNil("notorgid"), "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.assuring-red-shift"},
	{"merry-mega-man", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.merry-mega-man"},
	{"patient-red-wolf", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf"},
	{"coherent-night-nurse", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.coherent-night-nurse"},
	{"smashing-raphael", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.smashing-raphael"},
	{"gentle-tempest", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.gentle-tempest"},
	{"famous-rescue", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue"},
	{"crucial-mister-sinister", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue.crucial-mister-sinister"},
	{"flexible-iron-man", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue.flexible-iron-man"},

	}
var GetFoldersByOrgID_tests = [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{"normal case: alternate orgID", uuid.FromStringOrNil("notorgid"), sampleData, []folder.Folder{
				{"assuring-red-shift", uuid.FromStringOrNil("notorgid"), "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.assuring-red-shift"},
		}},
		{"normal case", def_org_id, sampleData, []folder.Folder{
			{"creative-scalphunter", def_org_id, "creative-scalphunter"},
			{"clear-arclight", def_org_id, "creative-scalphunter.clear-arclight"},
			{"topical-micromax", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax"},
			{"bursting-lionheart", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart"},
			{"striking-black-panther", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.striking-black-panther"},
			{"advanced-professor-monster", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.advanced-professor-monster"},
			{"merry-mega-man", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.merry-mega-man"},
			{"patient-red-wolf", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf"},
			{"coherent-night-nurse", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.coherent-night-nurse"},
			{"smashing-raphael", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.smashing-raphael"},
			{"gentle-tempest", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.gentle-tempest"},
			{"famous-rescue", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue"},
			{"crucial-mister-sinister", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue.crucial-mister-sinister"},
			{"flexible-iron-man", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue.flexible-iron-man"},

			},
		},

	}
var GetAllChildFolders_tests = [...]struct {
		name    string
		fname 	string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
		exp_err	error
	}{
		{"normal case", "creative-scalphunter", def_org_id, sampleData, []folder.Folder{
				{"clear-arclight", def_org_id, "creative-scalphunter.clear-arclight"},
				{"topical-micromax", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax"},
				{"bursting-lionheart", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart"},
				{"striking-black-panther", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.striking-black-panther"},
				{"advanced-professor-monster", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.advanced-professor-monster"},
				{"merry-mega-man", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.merry-mega-man"},
				{"patient-red-wolf", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf"},
				{"coherent-night-nurse", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.coherent-night-nurse"},
				{"smashing-raphael", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.smashing-raphael"},
				{"gentle-tempest", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.gentle-tempest"},
				{"famous-rescue", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue"},
				{"crucial-mister-sinister", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue.crucial-mister-sinister"},
				{"flexible-iron-man", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue.flexible-iron-man"},

			}, nil,
		},
		{"normal case: picking up on different orgID", "bursting-lionheart", def_org_id, sampleData, []folder.Folder{
				{"striking-black-panther", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.striking-black-panther"},
				{"advanced-professor-monster", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.advanced-professor-monster"},
				{"merry-mega-man", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.merry-mega-man"},
				}, nil, 
		},
		{"normal case: leaf node", "merry-mega-man", def_org_id, sampleData, []folder.Folder{
				}, nil, 
		},
		{"edge case: target is not in org", "assuring-red-shift", def_org_id, sampleData, []folder.Folder{
				}, errors.New("folder does not exist"), 
		},
		{"normal case: picking up on different orgID", "bursting-lionheart", def_org_id, sampleData, []folder.Folder{
				{"striking-black-panther", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.striking-black-panther"},
				{"advanced-professor-monster", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.advanced-professor-monster"},
				{"merry-mega-man", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart.merry-mega-man"},
				}, nil,
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
			return_folder, return_error := f.GetAllChildFolders(def_org_id, tt.fname)
			
			
			if return_folder == nil{ 
				assert.Error(return_error)
				assert.Equal(return_error, tt.exp_err)
				
			}
		})
	}
}