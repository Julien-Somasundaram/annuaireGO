# TP1 - Annuaire (en Go)

## EQUIPE
- SOMASUNDARAM Julien
- TAMIMOUL Chuaibe

## 📌 Présentation

Ce projet est un annuaire simple développé en langage **Go**
Il permet de gérer une liste de contacts (nom + téléphone) depuis la ligne de commande, avec une **sauvegarde automatique dans un fichier JSON**.

---

## ⚙️ Fonctionnalités

- ✅ Ajouter un contact
- ✅ Rechercher un contact
- ✅ Modifier un contact
- ✅ Supprimer un contact
- ✅ Lister tous les contacts
- ✅ Sauvegarde automatique dans `annuaire.json`

---

## 🧠 Concepts Go utilisés

- `map[string]contact` pour stocker les contacts
- `flag` pour récupérer les arguments en ligne de commande
- `os.ReadFile`, `os.WriteFile` pour lire/écrire le fichier
- `encoding/json` pour convertir les données en JSON
- Fonctions + struct + tests unitaires

---

## 🚀 Utilisation

```bash
# Ajouter un contact
go run main.go --action ajouter --nom "Charlie" --tel "0811223344"

# Rechercher un contact
go run main.go --action rechercher --nom "Charlie"

# Lister tous les contacts
go run main.go --action lister

# Modifier un contact
go run main.go --action modifier --nom "Charlie" --tel "0999888777"

# Supprimer un contact
go run main.go --action supprimer --nom "Charlie"
