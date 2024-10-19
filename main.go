package main

import (
	// "fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	// "github.com/gofrs/uuid"
)

func main() {
	// orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	// res := folder.GetAllFolders()

	// // example usage
	// folderDriver := folder.NewDriver(res)
	// // orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	// contains, error := folderDriver.GetAllChildFolders(orgID, name)

	// // folder.PrettyPrint(contains)
	// // fmt.Printf("\n Folders for orgID: %s", orgID)
	// // folder.PrettyPrint(orgFolder)
	res := folder.GenerateData()

    folder.PrettyPrint(res)

    folder.WriteSampleData(res)
}
