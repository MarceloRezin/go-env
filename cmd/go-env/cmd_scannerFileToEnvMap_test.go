package goenv

import (
	"bufio"
	"io"
	"testing"
)

type mockScanner struct {
	data string
	done bool
}

func (r *mockScanner) Read(p []byte) (n int, err error) {
	copy(p, []byte(r.data))
	if r.done {
		return 0, io.EOF
	}
	r.done = true
	return len([]byte(r.data)), nil
}

func TestEnvFileVazioVazio(t *testing.T) {
	s := bufio.NewScanner(&mockScanner{"", false})

	envMap, err := scannerFileToEnvMap(s)
	if err != nil {
		t.Error("Não deveria ter sido disparado um erro")
	}

	tamanho := len(envMap)
	if tamanho != 0 {
		t.Errorf("Tamano do map esperado %d, mas encontrado %d", 0, tamanho)
	}
}

func TestEnvFileUmaLinha(t *testing.T) {
	s := bufio.NewScanner(&mockScanner{lineValueValido, false})

	envMap, err := scannerFileToEnvMap(s)
	if err != nil {
		t.Error("Não deveria ter sido disparado um erro")
	}

	tamanho := len(envMap)
	if tamanho != 1 {
		t.Errorf("Tamano do map esperado %d, mas encontrado %d", 1, tamanho)
	}

	valorRecuperado := envMap[chave]
	if valorRecuperado != valor {
		t.Errorf("Valor esperado %s, mas encontrado %s", valor, valorRecuperado)
	}
}

func TestEnvFileDuasLinhas(t *testing.T) {
	mockEnvFile := lineValueValido + "\n" + lineValueValido2

	s := bufio.NewScanner(&mockScanner{mockEnvFile, false})

	envMap, err := scannerFileToEnvMap(s)
	if err != nil {
		t.Error("Não deveria ter sido disparado um erro")
	}

	tamanho := len(envMap)
	if tamanho != 2 {
		t.Errorf("Tamano do map esperado %d, mas encontrado %d", 2, tamanho)
	}

	valorRecuperadoPrimeraLinha := envMap[chave]
	if valorRecuperadoPrimeraLinha != valor {
		t.Errorf("Valor esperado %s, mas encontrado %s", valor, valorRecuperadoPrimeraLinha)
	}

	valorRecuperadoSegundaLinha := envMap[chave2]
	if valorRecuperadoSegundaLinha != valor2 {
		t.Errorf("Valor esperado %s, mas encontrado %s", valor2, valorRecuperadoSegundaLinha)
	}
}
