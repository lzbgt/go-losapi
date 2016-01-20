package cmd

import "fmt"

type Log struct {
	verbose bool
}

var log *Log = &Log{}

func (p *Log) Verbose(v bool) {
	p.verbose = v
	//fmt.Println("log set to verbose: ", p.verbose)
}

func (p *Log) V(args ...interface{}) {
	//fmt.Println("v: ", p.verbose)
	if p.verbose {
		fmt.Println(args...)
	}
}

func (p *Log) Info(args ...interface{}) {
	fmt.Println(args...)

}

func (p *Log) Vf(shape string, args ...interface{}) {
	if p.verbose {
		fmt.Printf(shape, args...)
	}
}

func (p *Log) Infof(shape string, args ...interface{}) {
	fmt.Printf(shape, args...)
}

func (p *Log) Debug(args ...interface{}) {
	log.V(args...)
}

func (p *Log) Debugf(shape string, args ...interface{}) {
	log.Vf(shape, args)
}
