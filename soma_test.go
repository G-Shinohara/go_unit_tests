package main

import "testing"

func TestShoudSumCorrect(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6
	if teste != resultado {

		t.Error("Valor esperado", resultado, "Valor retornado:", teste)
	}
}

func TestShoudSumIncorrect(t *testing.T) {
    teste := soma(3, 2, 1)
    resultado := 7

    if teste != resultado {
        t.Error("Valor esperado:", resultado, "Valor retornado", teste)
    }
}

func TestShoudMultCorrect(t *testing.T) {
    teste := multiplica(10, 10)
    resultado := 100

    if teste != resultado {
        t.Error("Valor esperado:", resultado, "Valor retornado", teste)
    }
}

func TestShoudMultIncorrect(t *testing.T) {
    teste := multiplica(10, 10)
    resultado := 2560

    if teste != resultado {
        t.Error("Valor esperado:", resultado, "Valor retornado", teste)
    }
}

func TestShoudSubtractCorrect(t *testing.T) {
    teste := subtrai(10, 3)
    resultado := -7

    if teste != resultado {
        t.Error("Valor esperado:", resultado, "Valor retornado", teste)
    }
}

func TestShoudSubtractIncorrect(t *testing.T) {
    teste := subtrai(10, 3)
    resultado := 10

    if teste != resultado {
        t.Error("Valor esperado:", resultado, "Valor retornado", teste)
    }
}

func TestShoudDivideCorrect(t *testing.T) {
    teste := divide(2, 10)
    resultado := 5

    if teste != resultado {
        t.Error("Valor esperado:", resultado, "Valor retornado", teste)
    }
}

func TestShoudDivideIncorrect(t *testing.T) {
    teste := divide(2, 10)
    resultado := 10

    if teste != resultado {
        t.Error("Valor esperado:", resultado, "Valor retornado", teste)
    }
}

