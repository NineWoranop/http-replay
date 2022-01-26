package file

import (
	"fmt"
	"http-replay/cmd"
)

type Counter struct {
	Count uint
	Max   uint
}

func (self Counter) CurrentValue() uint {
	return self.Count
}

func (self *Counter) Increment() {
	self.Count++
	self.Count %= self.Max
}

type HanlderConfig struct {
	Counter    *Counter
	passedOnce bool
	Path       string
	TotalFiles uint
	AutoRepeat bool
}

func ToHanlderConfig(args *cmd.CmdArguments) *HanlderConfig {
	counter := Counter{Count: 0, Max: args.TotalFiles}
	return &HanlderConfig{Counter: &counter, Path: args.Path, TotalFiles: args.TotalFiles, AutoRepeat: args.AutoRepeat}
}

func (conf HanlderConfig) CurrentFilePath() string {
	var filename string = fmt.Sprintf("%06d.dat", conf.Counter.CurrentValue()+1)
	var filepath string = conf.Path + filename
	return filepath
}

func (conf HanlderConfig) NextFilePath() string {
	var filename string = fmt.Sprintf("%06d.dat", conf.NextCounterValue()+1)
	var filepath string = conf.Path + filename
	return filepath
}

func (conf HanlderConfig) CanIncrease() bool {
	if conf.AutoRepeat {
		return (conf.TotalFiles > 1)
	} else if conf.TotalFiles > 1 {
		return ((conf.CurrentIndex() + 1) < conf.TotalFiles)
	} else {
		return false
	}
}

func (conf HanlderConfig) increase(c uint) uint {
	result := c + 1
	result %= conf.TotalFiles
	return result
}

func (conf HanlderConfig) Next() {
	conf.passedOnce = true
	conf.Counter.Increment()
}

func (conf HanlderConfig) NextCounterValue() uint {
	return conf.increase(uint(conf.Counter.CurrentValue()))
}

func (conf HanlderConfig) PassedOnce() bool {
	return conf.passedOnce
}

func (conf HanlderConfig) CurrentIndex() uint {
	return uint(conf.Counter.CurrentValue())
}
