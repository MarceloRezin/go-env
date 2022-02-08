package goenv

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func Read(envDir string) (map[string]string, error) {
	scanner, err := getScannerFromFileDir(envDir)
	if err != nil {
		return nil, errors.New("Não foi possível ler o arquivo.")
	}

	envVars := make(map[string]string)
	for scanner.Scan() {
		chave, valor, err := parseLine(scanner.Text())
		if err != nil {
			return nil, err
		}

		envVars[chave] = valor
	}

	return envVars, nil
}

func getScannerFromFileDir(fileDir string) (*bufio.Scanner, error) {
	file, err := os.Open(fileDir)
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(file), nil
}

func parseLine(line string) (string, string, error) {
	chaveValor := strings.Split(line, "=")

	if len(chaveValor) == 2 {
		chave := chaveValor[0]
		if chave == "" {
			return "", "", errors.New("A chave não pode ser vazia")
		}

		return chave, chaveValor[1], nil
	} else {
		return "", "", errors.New("Conjunto chave valor inválido.")
	}
}
