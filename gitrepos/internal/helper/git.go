package helper

import "gitrepos/internal/exec"

const gitBin = "git"

// Git git cmd line
type Git struct {
	Bin  string   `json:"bin"`
	Dir  string   `json:"dir"`
	Args []string `json:"args"`
}

func NewGit() (*Git, error) {
	g := new(Git)
	s, err := exec.Find(gitBin)
	if err != nil {
		g.Bin = s
	}
	return g, err
}

func (g *Git) SetDir(dir string) {
	g.Dir = dir
}

func (g *Git) SetArgs(args []string) {
	g.Args = args
}

// Clone clone to target path
func (g *Git) Clone(url string, named string) error {
	c := make([]string, 0, len(g.Args)+4)
	c = append(c, g.Bin, "clone")
	c = append(c, g.Args...)
	c = append(c, url, named)
	return exec.Run(c, g.Dir)
}

// Fetch fetch all new code
func (g *Git) Fetch(projectPath string) error {
	c := make([]string, 0, len(g.Args)+3)
	c = append(c, g.Bin, "fetch")
	c = append(c, g.Args...)
	c = append(c, projectPath)
	return exec.Run(c, g.Dir)
}
