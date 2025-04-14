package main

import (
	"testing"
	"bytes"
	"fmt"
)

func TestSoma(t *testing.T) {
	resultado := soma(112, 10)
	esperado := 122

	if resultado != esperado {
		t.Errorf("Resultado incorreto: esperado %d, obtido %d", esperado, resultado)
	}
}

func TestRun(t *testing.T) {
	// Redireciona a saída padrão para um buffer
	var buf bytes.Buffer
	fmtOrig := fmt.Println
	defer func() { fmt.Println = fmtOrig }() // restaura ao final do teste

	fmt.Println = func(a ...any) (n int, err error) {
		return fmt.Fprint(&buf, a...)
	}

	run()

	saida := buf.String()
	esperado := "122"

	if saida != esperado {
		t.Errorf("Saída incorreta: esperada %q, obtida %q", esperado, saída)
	}
}
