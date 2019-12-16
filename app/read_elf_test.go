package main

import (
	"debug/elf"
	"debug/gosym"
	"fmt"
	"testing"
)

func TestReadELFSymtab(t *testing.T) {
	f, err := elf.Open("app")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	symtab, err := f.Section(".gosymtab").Data()
	if err != nil {
		panic(err)
	}

	gopclntab, err := f.Section(".gopclntab").Data()
	if err != nil {
		panic(err)
	}

	pcln := gosym.NewLineTable(gopclntab, f.Section(".text").Addr)
	var tab *gosym.Table
	tab, err = gosym.NewTable(symtab, pcln)
	if err != nil {
		panic(err)
	}

	for _, x := range tab.Funcs {
		fmt.Printf("addr:0x%x\t\tname:%s,\t\n", x.Entry, x.Name)
	}
}
