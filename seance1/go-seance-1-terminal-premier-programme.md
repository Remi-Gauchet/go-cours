# Séance 1 — Le terminal et le premier programme

**Application CLI & Parsing en Go — Bachelor 1 — Ycode A2627_4658**
**Séance 1 · 2 h**
**Prérequis : aucun. Cette séance part de zéro.**

Vous n'avez jamais programmé ? C'est parfait, ce cours est fait pour vous. Chaque mot technique est expliqué la première fois qu'il apparaît.

À la fin du module, vous aurez construit un vrai petit logiciel : **textfix**, un outil qui corrige automatiquement un fichier texte. Aujourd'hui, on découvre les outils de base : le terminal, et un tout premier programme.

---

## Ce que vous saurez faire en fin de séance

Ouvrir un terminal et vous déplacer dans vos dossiers.
Créer un programme en Go et le lancer.
Comprendre un message d'erreur simple.
Ranger une information dans une variable et l'afficher.

---

## Déroulé

| Partie | Contenu | Durée |
|---|---|---|
| 1 | C'est quoi, un programme ? | 10 min |
| 2 | Le terminal | 20 min |
| — | Exercice 1 | 20 min |
| 3 | Le premier programme | 15 min |
| — | Exercice 2 | 25 min |
| 4 | Les variables | 15 min |
| — | Exercice 3 | 15 min |

---

# Partie 1 — C'est quoi, un programme ?

## Une recette de cuisine

Un programme, c'est une **liste d'instructions** que l'ordinateur exécute dans l'ordre, une par une. Exactement comme une recette : « casser les œufs », puis « ajouter la farine », puis « mélanger ».

La différence : l'ordinateur ne devine rien. Si une instruction est mal écrite, même d'une seule lettre, il s'arrête et vous le signale.

## Un langage

L'ordinateur ne comprend pas le français. On lui parle dans un **langage de programmation**. Il en existe des centaines : Python, JavaScript, Java… Dans ce module, on utilise **Go**.

Go a été créé par Google en 2009. Il est apprécié parce qu'il est simple à lire et très rapide.

## Le code et la compilation

Le texte que vous écrivez s'appelle le **code**. Avant de l'exécuter, Go le **compile** : il le traduit dans le langage de la machine.

Si votre code contient une faute, la compilation échoue et Go vous indique **la ligne** du problème. C'est une aide, pas une punition.

---

# Partie 2 — Le terminal

## Ce que c'est

Le **terminal** est une fenêtre dans laquelle on tape des commandes au clavier, au lieu de cliquer avec la souris.

Ça paraît austère au début, mais c'est l'outil de base de tous les développeurs. Et notre projet `textfix` sera justement un programme qu'on lance depuis le terminal.

## L'ouvrir dans VS Code

Ouvrez VS Code, puis le menu **Terminal → Nouveau terminal**. Une zone apparaît en bas de la fenêtre.

Vous y voyez une ligne qui ressemble à ceci :

```
PS C:\Users\lea>
```

C'est l'**invite de commande**. Elle vous indique **dans quel dossier vous êtes**, et attend une commande.

## Les dossiers et les chemins

Vos fichiers sont rangés dans des dossiers, eux-mêmes rangés dans d'autres dossiers. Le **chemin** d'un fichier, c'est son adresse : la liste des dossiers à traverser pour l'atteindre.

```
C:\Users\lea\Documents\cours\notes.txt
```

## Les cinq commandes à connaître

| Commande | Ce qu'elle fait |
|---|---|
| `pwd` | Affiche le dossier où vous êtes |
| `ls` | Liste le contenu du dossier |
| `mkdir nom` | Crée un dossier |
| `cd nom` | Entre dans un dossier |
| `cd ..` | Remonte d'un dossier |

Elles fonctionnent sous Windows, sous Mac et sous Linux.

## Deux astuces qui font gagner beaucoup de temps

**La touche Tab** complète automatiquement un nom de dossier. Tapez `cd Doc` puis Tab : le terminal écrit `Documents` pour vous.

**La flèche du haut** rappelle les commandes précédentes. Inutile de tout retaper.

---

## ✅ Exercice 1 — 20 min

**1.1** Ouvrez un terminal dans VS Code. Tapez `pwd` et notez le dossier où vous êtes.

**1.2** Tapez `ls`. Retrouvez-vous des noms de dossiers que vous connaissez ?

**1.3** Allez dans votre dossier Documents avec `cd`.

**1.4** Créez un dossier `go-cours` avec `mkdir`, puis entrez dedans.

**1.5** Remontez d'un dossier avec `cd ..`, puis revenez dans `go-cours` en utilisant la touche Tab.

**1.6** Tapez `go version`. Si vous voyez un numéro de version, Go est bien installé. Sinon, appelez l'enseignant.

**Vérification** : `pwd` affiche un chemin qui se termine par `go-cours`, et `go version` affiche un numéro.

---

# Partie 3 — Le premier programme

## Préparer le projet

Dans le dossier `go-cours`, créez un dossier `bonjour` et entrez dedans :

```bash
mkdir bonjour
cd bonjour
go mod init bonjour
```

La commande `go mod init` crée un fichier `go.mod`. C'est la **fiche d'identité** du projet : son nom et la version de Go. Sans ce fichier, Go refuse de lancer le programme. On le fait **une seule fois** par projet.

## Ouvrir le dossier dans VS Code

Menu **Fichier → Ouvrir le dossier**, puis choisissez `bonjour`. Créez un fichier nommé `main.go`.

## Le code

Recopiez exactement ceci :

```go
package main

import "fmt"

func main() {
	fmt.Println("Bonjour !")
}
```

## Ligne par ligne

`package main` — ce fichier est le point de départ d'un programme. Toujours en première ligne.

`import "fmt"` — on utilise une boîte à outils déjà écrite, nommée `fmt`, qui sert à afficher du texte.

`func main() { … }` — c'est la **fonction principale**. Quand on lance le programme, Go exécute ce qu'il y a entre les accolades `{` et `}`.

`fmt.Println("Bonjour !")` — affiche le texte entre guillemets, puis passe à la ligne. `Println` se lit « print line » : « afficher une ligne ».

## Lancer le programme

Dans le terminal :

```bash
go run .
```

Le point `.` veut dire « le dossier où je suis ». Vous devez voir :

```
Bonjour !
```

## Les erreurs, vos alliées

Si vous oubliez le guillemet fermant, Go affiche :

```
./main.go:6:25: newline in string
```

`main.go:6:25` signifie : fichier `main.go`, **ligne 6**, colonne 25. Allez voir cette ligne : l'erreur est là, ou juste avant.

**Lire le message d'erreur, c'est la moitié du travail.** Ne paniquez pas devant un message en anglais : cherchez d'abord le numéro de ligne.

---

## ✅ Exercice 2 — 25 min

**2.1** Créez le projet `bonjour` comme expliqué, et lancez-le.

**2.2** Ajoutez deux lignes pour afficher votre prénom et ce que vous apprenez :

```
Bonjour !
Je m'appelle Léa.
J'apprends le Go.
```

**2.3** Supprimez volontairement un guillemet. Lancez. Notez le numéro de ligne indiqué, puis corrigez.

**2.4** Écrivez `fmt.println` avec un `p` minuscule. Lancez et lisez l'erreur. En Go, les majuscules comptent.

**2.5** Supprimez la ligne `import "fmt"`. Lancez et lisez l'erreur. Remettez-la.

**Vérification** : votre programme affiche trois lignes, et vous savez retrouver la ligne fautive à partir d'un message d'erreur.

---

# Partie 4 — Les variables

## Une boîte avec une étiquette

Une **variable** est une boîte dans laquelle on range une information, avec un nom collé dessus pour la retrouver.

```go
prenom := "Léa"
```

Cette ligne crée une boîte nommée `prenom`, et range dedans le texte `"Léa"`.

Le symbole `:=` veut dire « crée la boîte et mets ceci dedans ».

## Utiliser une variable

```go
fmt.Println(prenom)
```

affiche `Léa`. On écrit le nom de la variable **sans guillemets** : avec des guillemets, Go afficherait le mot « prenom ».

## Deux sortes d'informations

| Sorte | Exemple | Nom en Go |
|---|---|---|
| Du texte | `"Léa"` | `string` |
| Un nombre entier | `19` | `int` |

Le texte s'écrit **entre guillemets**, les nombres **sans**. On appelle ces sortes des **types**.

## Afficher plusieurs choses

```go
age := 19
fmt.Println("Âge :", age)
```

affiche `Âge : 19`. La virgule sépare les éléments, et Go met un espace entre eux.

## Changer le contenu de la boîte

```go
age = age + 1
```

On utilise `=` tout seul, **sans les deux-points**, parce que la boîte existe déjà. On ne la crée pas, on change son contenu.

## Coller du texte

```go
phrase := prenom + " habite à " + ville + "."
```

Le `+` entre deux textes les colle bout à bout. Attention aux espaces : il faut les écrire vous-même à l'intérieur des guillemets.

---

## ✅ Exercice 3 — 15 min

**3.1** Créez un nouveau projet `identite`, comme pour `bonjour`.

**3.2** Créez trois variables : `prenom`, `ville`, et `age`.

**3.3** Affichez-les, une par ligne, avec une étiquette :

```
Prénom : Léa
Ville : Paris
Âge : 19
```

**3.4** Ajoutez 1 à l'âge, puis affichez : `L'année prochaine, j'aurai 20 ans.`

**3.5** Construisez avec `+` la phrase `Léa habite à Paris.` et affichez-la.

**Vérification** : le programme affiche les cinq lignes, et changer la valeur d'une variable change bien l'affichage.

---

# À retenir

## Le terminal

`pwd` où je suis, `ls` ce qu'il y a, `mkdir` créer, `cd` entrer, `cd ..` remonter.
Tab pour compléter, flèche du haut pour rappeler une commande.

## Le programme

`go mod init nom` une fois par projet.
`package main`, `import "fmt"`, `func main() { … }` : le squelette de tout programme.
`go run .` pour lancer.
Une erreur indique toujours un numéro de ligne.

## Les variables

`:=` crée une variable, `=` change son contenu.
Le texte entre guillemets, les nombres sans.
`+` colle deux textes.

## Les erreurs qui coûtent le plus de temps

| Symptôme | Cause |
|---|---|
| `go: command not found` ou commande inconnue | Go n'est pas installé, ou le terminal a été ouvert avant l'installation |
| `go.mod file not found` | `go mod init` oublié, ou terminal dans le mauvais dossier |
| `newline in string` | Un guillemet fermant manquant |
| `undefined: fmt.println` | `println` au lieu de `Println` : la majuscule compte |
| `undefined: fmt` | La ligne `import "fmt"` a été supprimée |
| `declared and not used` | Une variable créée mais jamais utilisée : Go l'interdit |
| `no new variables on left side of :=` | `:=` utilisé sur une variable qui existe déjà : il faut `=` |
| `undefined: age` | `=` utilisé sur une variable qui n'existe pas encore : il faut `:=` |
| Le mot « prenom » s'affiche au lieu de « Léa » | Le nom de la variable est entre guillemets |

---

# Prochaine séance

**Les briques de base (4 h)** : prendre des décisions avec `if`, manipuler du texte, faire des listes, répéter une action, et écrire vos propres fonctions.
