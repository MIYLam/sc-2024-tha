package folder_test

import (
	"errors"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	// "github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

var MoveFolderTests = [...]struct {
		name    string
		cname 	string
		dname	string
		folders []folder.Folder
		want    []folder.Folder
		exp_err	error
	}{
		{"normal case", "bursting-lionheart", "patient-red-wolf", sampleData, []folder.Folder{
			{"creative-scalphunter", def_org_id, "creative-scalphunter"},
			{"clear-arclight", def_org_id, "creative-scalphunter.clear-arclight"},
			{"topical-micromax", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax"},
			{"bursting-lionheart", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.bursting-lionheart"},
			{"striking-black-panther", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.bursting-lionheart.striking-black-panther"},
			{"advanced-professor-monster", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.bursting-lionheart.advanced-professor-monster"},
			{"merry-mega-man", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.bursting-lionheart.merry-mega-man"},
			{"patient-red-wolf", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf"},
			{"coherent-night-nurse", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.coherent-night-nurse"},
			{"smashing-raphael", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.smashing-raphael"},
			{"gentle-tempest", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.gentle-tempest"},
			{"famous-rescue", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue"},
			{"crucial-mister-sinister", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue.crucial-mister-sinister"},
			{"flexible-iron-man", def_org_id, "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue.flexible-iron-man"},
		}, nil,
	},{
		"different orgID", "bursting-lionheart", "assuring-red-shift", sampleData, nil, errors.New("different orgID"),
		},
		
	{
		"nested children", "bursting-lionheart", "advanced-professor-monster", sampleData, nil, errors.New("destination is child of target"),
		},
	{
		"duplicate", "bursting-lionheart", "bursting-lionheart", sampleData, nil, errors.New("cannot move folder to itself"),
		},
	}

func Test_MoveFolder(t *testing.T) {
	t.Parallel()
	for _, tt := range MoveFolderTests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			f := folder.NewDriver(tt.folders)
			return_folder, return_error := f.MoveFolder(tt.cname, tt.dname)
			
			
			if return_folder == nil{ 
				assert.Error(return_error)
				assert.Equal(return_error, tt.exp_err)
				
			}
		})
	}
}
