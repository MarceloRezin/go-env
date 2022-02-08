package goenv

import "testing"

func TestReadFileInvalido(t *testing.T) {
	_, err := Read("")
	if err == nil {
		t.Error("Um erro deveria ter sido disparado.")
	}
}

func TestReadFileInexistente(t *testing.T) {
	_, err := Read("./env")
	if err == nil {
		t.Error("Um erro deveria ter sido disparado.")
	}
}

func TestReadFileValida(t *testing.T) {
	envVars, err := Read("./test/.envFileTest")
	if err != nil {
		t.Error("Não foi possível ler o arquivo env.")
	}

	v := envVars[chave]
	if v != valor {
		t.Errorf("Esperado %s, encontrado %s", valor, v)
	}

}
