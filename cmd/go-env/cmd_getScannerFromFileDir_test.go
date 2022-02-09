package goenv

import "testing"

func TestFileDirInvalido(t *testing.T) {
	_, err := getScannerFromFileDir("")
	if err == nil {
		t.Error("Um erro deveria ter sido disparado.")
	}
}

func TestFileDirInexistente(t *testing.T) {
	_, err := getScannerFromFileDir("./env")
	if err == nil {
		t.Error("Um erro deveria ter sido disparado.")
	}
}

func TestFileDirValida(t *testing.T) {
	scanner, err := getScannerFromFileDir("./test/.envFileTest")
	if err != nil {
		t.Error("Não foi possível ler o arquivo env.")
	}

	if scanner == nil {
		t.Error("Esperado um scanner diferente de nulo")
	}
}
