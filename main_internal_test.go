package main

import "testing"

func TestAjouterEtRechercher(t *testing.T) {
	annuaire = make(map[string]contact)
	ajouter("Alice", "1234")
	result := rechercher("Alice")
	expected := "Contact trouvé : Alice - 1234"
	if result != expected {
		t.Errorf("rechercher = %v; want %v", result, expected)
	}
}

func TestAjouterDoublon(t *testing.T) {
	annuaire = make(map[string]contact)
	ajouter("Bob", "4567")
	result := ajouter("Bob", "9999")
	expected := "Ce contact existe déjà."
	if result != expected {
		t.Errorf("ajouter doublon = %v; want %v", result, expected)
	}
}

func TestSupprimer(t *testing.T) {
	annuaire = make(map[string]contact)
	ajouter("Charlie", "0000")
	result := supprimer("Charlie")
	expected := "Contact supprimé."
	if result != expected {
		t.Errorf("supprimer = %v; want %v", result, expected)
	}
}

func TestModifier(t *testing.T) {
	annuaire = make(map[string]contact)
	ajouter("David", "1111")
	result := modifier("David", "2222")
	expected := "Contact modifié."
	if result != expected {
		t.Errorf("modifier = %v; want %v", result, expected)
	}
}
