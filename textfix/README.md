# textfix
Programme en GO servant à transformant un fichier texte en appliquant plusieurs modifications, insérées dans le texte en tant que marqueurs.

## Installation
Teléchargez le dossier textfix.
Placez votre fichier .txt à la racine du projet.
Dans votre éditeur de texte, donnez en argument le nom du fichier à modifier, suivi du nom du fichier à créer

## Utilisation
Exemple : 

```bash
go run . entree.txt sortie .txt
```
(ici **entree.txt** est le texte d'entrée à modifier
et **sortie.txt** et le résultat obtenu)

## Les règles
|    Marqueur    |              Effet           |  Exemple   | Résultat |
|----------------|------------------------------|------------|----------|
|     `(up)`     |       tout en majuscule      | chat (up)  |   CHAT   |
|     `(low)`    |       tout en minuscule      | CHAT (low) |   chat   |
|     `(cap)`    | première lettre en majuscule | chat (cap) |   Chat   |

## Les messages d'erreur
Impossible de lire le fichier
Impossible d'écrire le fichier
Veuillez entrer le fichier d'entrée puis le fichier de sortie.
Veillez à ce que le nom du fichier d'entrée et de sortie soient différents.


## Les tests
Lancez l'input suivant :
go run . tests/01-majuscules.txt sortie.txt
Le fichier sortie.txt doit avoir le même contenu que le fichier 01-majuscules.attendu.txt

Vous pouvez reproduire le test avec un autre fichier en suivant le même schéma.


## Organisation du code
fichiers.go lit le fichier donné, et en crée un nouveau
regles.go définit et applique les différentes modifications
main.go lance le programme et appelle les autres fonctions

## Limites
Ne gère pas la ponctuation.
Les marqueurs doivent impérativement être séparés d'un espace avant et après votre texte.

## Auteur
Rémi GAUCHET