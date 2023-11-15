package shooting

import "fmt"

// type ModuleStrategy interface {
// 	Execute(*Gun)
// }

// type Module struct {
// 	moduleType ModuleStrategy
// }

// func NewModule(moduleType ModuleStrategy) *Module {
// 	return &Module{moduleType: moduleType}
// }

// func (m *Module) Run(g *Gun) {
// 	m.moduleType.Execute(g)
// }

type Config struct {
	Name      string   `yaml:"name"`
	Workflows []Module `yaml:"workflows"`
}

type Module struct {
	Name string `yaml:"name"`

	// Add Modules
	Fact  *FactModule  `yaml:"fact,omitempty"`
	Shoot *ShootModule `yaml:"shoot,omitempty"`
}

type FactModule struct {
	Raw  []string `yaml:"raw,omitempty"`
	File string   `yaml:"file,omitempty"`
}

func (f *FactModule) Execute(g *Gun, r Reader) {
	if f.Raw != nil {
		g.AddRawFacts(f.Raw)
		fmt.Println(g.facts)
	} else {
		f.loadFile(r)
	}
}

func (f *FactModule) loadFile(r Reader) {
	// TODO: add file loader
}

type ShootModule struct {
	File string `yaml:"file,omitempty"`
}

func (s *ShootModule) Execute(g *Gun, r Reader) {
	requests, err := r.readCustomFile(s.File)
	if err != nil {
		fmt.Println(err)
	}
	g.Load(requests)
	g.Shoot()

}
