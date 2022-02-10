package goenv

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

const SEPARATOR = '='

func Read(envDir string) (map[string]string, error) {
	scanner, err := getScannerFromFileDir(envDir)
	if err != nil {
		return nil, errors.New("Não foi possível ler o arquivo.")
	}

	return scannerFileToEnvMap(scanner)
}

func getScannerFromFileDir(fileDir string) (*bufio.Scanner, error) {
	file, err := os.Open(fileDir)
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(file), nil
}

func scannerFileToEnvMap(scanner *bufio.Scanner) (map[string]string, error) {
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

func parseLine(line string) (string, string, error) {
	indexSeparator := strings.IndexByte(line, SEPARATOR)
	if indexSeparator <= 0 {
		return "", "", errors.New("Conjunto chave valor inválido.")
	}

	chave := line[:indexSeparator]
	if chave == "" {
		return "", "", errors.New("A chave não pode ser vazia")
	}

	valor := line[(indexSeparator + 1):]
	return chave, valor, nil
}
