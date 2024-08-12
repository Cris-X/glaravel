package glaravel

const version = "1.0.0"

type Glaravel struct {
	AppName string
	Debug   bool
	Version string
}

func (g *Glaravel) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handler", "models", "logics", "migrations", "views", "public", "logs", "middleware", "tmp", "routers"},
	}

	err := g.Init(pathConfig)
	if err != nil {
		return err
	}

	return nil
}

func (g *Glaravel) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// create folder if it doesn't exist
		err := g.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}
