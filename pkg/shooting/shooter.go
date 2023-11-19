package shooting

import (
	"fmt"

	"github.com/kaitoz11/minaris/pkg/utils/setting"
)

type Shooter struct {
	gun      *Gun
	reader   Reader
	workflow *Config
	config   *setting.ShootingConfig
}

func NewShooterWorkflow(wf *Config) *Shooter {
	shooter := new(Shooter)
	shooter.workflow = wf

	shooter.reader = new(RawReader)

	shooter.gun = new(Gun)
	return shooter
}

func (s *Shooter) UseConfig(cf *setting.ShootingConfig) *Shooter {
	s.config = cf
	s.gun.Init(cf.Proxy)

	return s
}

func (s *Shooter) UseReader(rd Reader) *Shooter {
	s.reader = rd
	s.reader.init()
	return s
}

func (s *Shooter) ShootWorkflow() {
	for index, step := range s.workflow.Workflows {
		fmt.Printf("[%d] %s\n", index+1, step.Name)

		if step.Fact != nil {
			// fmt.Println("WTfact", step.Fact.Raw)
			fmt.Printf("--- BEFORE %s\n", s.gun.PrintFacts())
			step.Fact.Execute(s.gun, s.reader)
			fmt.Printf("--- AFTER %s\n", s.gun.PrintFacts())
		} else if step.Shoot != nil {
			fmt.Println("HolyShoot")
			fmt.Printf("--- BEFORE SHOOT %s\n", s.gun.PrintFacts())
			step.Shoot.Execute(s.gun, s.reader)
			fmt.Printf("--- AFTER %s\n", s.gun.PrintFacts())
		} else {
			fmt.Println("Default")
		}

		// fmt.Println(s.gun.facts)

		// for k, v := range []interfaces {
		// 	fmt.Printf("-- Module [%s] \n", k)
		// }
	}

	// requests, err := ReadFromFile()
}
