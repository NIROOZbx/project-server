package fileupload

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/NIROOZbx/project-server/internal/shared/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFileToCloudinary(ctx context.Context,imageData *multipart.FileHeader) (string,error) {

	folderName:="products"

	cld := config.GetConfig().CloudinaryURL

	file, err := imageData.Open()
	if err != nil {
		return "", err
	}

	fmt.Println("File data",file)

	defer file.Close()

	cloud, err := cloudinary.NewFromURL(cld)

	if err != nil {
		log.Println("Failed to initialize Cloudinary", err)
	}

	resp, err := cloud.Upload.Upload(ctx, file,uploader.UploadParams{
		Folder: folderName,
	})

	if err != nil {
		log.Fatalf("Failed to upload image: %v", err)
	}

	return resp.SecureURL,nil

}
