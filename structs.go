package gx

type Catalog struct {
	Pages []CatalogPage
}

type CatalogPage struct {
	PageNumber int    `json:"page"`
	Threads    []Post `json:"threads"`
}

type Post struct {
	No              int    `json:"no"`
	LastModified    int    `json:"last_modified"`
	Comment         string `json:"com"`
	Cyclical        string `json:"cyclical"`
	Time            int    `json:"time"`
	Filesize        int    `json:"fsize"`
	Filename        string `json:"filename"`
	Images          int    `json:"images"`
	Locked          int    `json:"locked"`
	OmittedImages   int    `json:"omitted_images"`
	OmittedPosts    int    `json:"omitted_posts"`
	Name            string `json:"name"`
	Replies         int    `json:"replies"`
	ReplyTo         int    `json:"resto"`
	Sticky          int    `json:"sticky"`
	MD5             string `json:"md5"`
	Extension       string `json:"ext"`
	RenamedFilename string `json:"tim"`
	ThumbnailHeight int    `json:"tn_h"`
	ThumbnailWidth  int    `json:"tn_w"`
	ImageHeight     int    `json:"h"`
	ImageWidth      int    `json:"w"`
	ExtraFiles      []File `json:"extra_files"`
}

type File struct {
	Extension       string `json:"ext"`
	Filename        string `json:"file"`
	Filesize        int    `json:"fsize"`
	Height          int    `json:"h"`
	Width           int    `json:"w"`
	ThumbnailHeight int    `json:"tn_h"`
	ThumbnailWidth  int    `json:"tn_w"`
}

type Thread struct {
	Posts []Post `json:"posts"`
}
