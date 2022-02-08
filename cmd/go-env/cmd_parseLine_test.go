package goenv

import "testing"

const chave = "key"
const valor = "value"
const lineValueValido = chave + "=" + valor
const lineValueValorVazio = chave + "="
const lineValueChaveVazio = "=" + valor
const lineValueTudoVazio = "="

func TestParseLineChaveValorValido(t *testing.T) {
	c, v, err := parseLine(lineValueValido)

	if err != nil {
		t.Error("Não foi possível fazer o parse da linha")
	}

	if c != chave {
		t.Errorf("Esperado %s, encontrado %s", chave, c)
	}

	if v != valor {
		t.Errorf("Esperado %s, encontrado %s", valor, v)
	}
}

func TestParseLineValorVazio(t *testing.T) {
	c, v, err := parseLine(lineValueValorVazio)

	if err != nil {
		t.Error("Não foi possível fazer o parse da linha")
	}

	if c != chave {
		t.Errorf("Esperado %s, encontrado %s", chave, c)
	}

	if v != "" {
		t.Errorf("Esperado %s, encontrado %s", "", v)
	}
}

func TestParseLineChaveVazio(t *testing.T) {
	_, _, err := parseLine(lineValueChaveVazio)

	if err == nil {
		t.Error("Um erro deveria ter sido disparado")
	}
}

func TestParseLineTudoVazio(t *testing.T) {
	_, _, err := parseLine(lineValueTudoVazio)

	if err == nil {
		t.Error("Um erro deveria ter sido disparado")
	}
}

func TestParseLineVazio(t *testing.T) {
	_, _, err := parseLine("")

	if err == nil {
		t.Error("Um erro deveria ter sido disparado")
	}
}
