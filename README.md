# TP : GoLog Analyzer

## Membres du groupe

- Matteo Vecchione
- Yamil Issa
- Florent Paris

## Description

GoLog Analyzer est un outil en ligne de commande écrit en Go, conçu pour aider les administrateurs système à analyser des fichiers de logs provenant de différentes sources (serveurs, applications, etc.). Il permet d'exécuter les analyses de manière concurrente, de gérer les erreurs proprement, et de centraliser les résultats sous forme de résumé JSON.

Ce projet a été réalisé dans le cadre d’un TP visant à renforcer les compétences suivantes :

- Gestion de la concurrence avec goroutines et WaitGroups
- Création d’un CLI modulaire avec Cobra
- Manipulation de fichiers JSON en entrée et en sortie
- Gestion des erreurs personnalisées avec errors.Is et errors.As
- Structuration d’un projet Go avec des packages internes (internal/)

## Fonctionnement du programme

### Étapes d’analyse

Pour chaque fichier de log défini dans le fichier config.json, l’outil :

1. Vérifie que le fichier est lisible.
2. Valide le type de fichier via une liste blanche.
3. Simule un traitement via un time.Sleep aléatoire (50–200ms).
4. Possède une probabilité de 10% d’échouer avec une erreur de parsing.
5. Remonte les résultats dans un rapport JSON

### Erreurs prises en charge

- UnsupportedFileTypeError : le type de log n’est pas dans la liste des types autorisés.
- FileAccessError : le fichier est introuvable ou inaccessible.
- ParsingError : erreur simulée de parsing lors de l’analyse.

## Commandes CLI

L'outil est basé sur [Cobra](https://github.com/spf13/cobra) pour structurer ses commandes.

Syntaxe de la commande analyze:

```sh
loganizer analyze --config config.json [--output report.json]
```

| Flag              | Description                                  |
| ----------------- | -------------------------------------------- |
| `--config` / `-c` | Chemin vers le fichier JSON de configuration |
| `--output` / `-o` | Chemin du rapport exporté au format JSON     |

## Développement et exécution

Installer les hooks Git (optionnel) :

```sh
sh setup_git_hook.sh
```

Lancer le CLI :

```sh
make clean run ARGUMENTS="analyze --config config.json"
```

Exporter les résultats dans un fichier :

```sh
make clean run ARGUMENTS="analyze --config config.json --output report.json"
```
