package models

type Configuration struct {
	Objects  []Object `json:"objects"`
	Format   string   `json:"file_format"`
	FileName string   `json:"file_name"`
}
