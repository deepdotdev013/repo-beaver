package contracts

type InitConfig struct {
	ProjectName string
	ModulePath  string
}

type FileTemplate struct {
	Tmpl string
	Dest string
}
