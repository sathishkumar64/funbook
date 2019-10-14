package model

//Category is used to define model.
type Category struct {
	CategoryID  string
	Name        string
	Alias       string
	Description string
	active      bool
}

//SubCategory is used to define model.
type SubCategory struct {
	SubCategoryID string
	CategoryID    string
	Name          string
	Alias         string
	Description   string
	active        bool
}