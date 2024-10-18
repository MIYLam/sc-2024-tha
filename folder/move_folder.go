package folder

import("errors")

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folders := f.folders

	var new_parent Folder;
	

	for _, f := range folders{
		if f.Name == dst{
			new_parent = f;
		}
	}
	if new_parent.Name == "" {
		return folders, errors.New("destination not found")
	}else{
		return nil, nil 
	}


	return []Folder{}, nil
}
