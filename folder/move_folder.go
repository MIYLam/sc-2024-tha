package folder

import (
	"errors"
	"slices"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folders := f.folders

	var new_parent Folder
	var changed_folder Folder

	if name == dst{
		return nil, errors.New("cannot move folder to itself")
	}
	for _, f := range folders{
		if f.Name == dst{
			new_parent = f;
		} else if f.Name == name{
			changed_folder = f;
		}
	}

	if new_parent.Name == "" {
		return folders, errors.New("destination not found")
		}else if changed_folder.Name == ""{
			return folders, errors.New("target folder not found")
		}
	if new_parent.OrgId != changed_folder.OrgId{
		return nil, errors.New("different orgID")
	}
	original_filepath := changed_folder.Paths 
	
	// If the destination is a child of the target

	dst_slice := strings.Split(new_parent.Paths, ".")
	if slices.Contains(dst_slice, changed_folder.Name){
		return nil, errors.New("destination is child of target")
	}
	// Change the filepath
	for i := 0; i < len(folders); i++{
		file := folders[i]
		if slices.Contains(strings.Split(file.Name, "."), changed_folder.Name){
			folders[i].Paths = strings.Replace(file.Paths, original_filepath, new_parent.Paths + changed_folder.Name, 1)
		}
	}

	return folders, nil
}
