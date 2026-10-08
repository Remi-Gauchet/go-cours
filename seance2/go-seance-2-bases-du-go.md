# Séance 2 — Les briques de base

**Application CLI & Parsing en Go — Bachelor 1 — Ycode A2627_4658**
**Séance 2 · 4 h**
**Prérequis : séance 1**

En séance 1, votre programme affichait toujours la même chose. Aujourd'hui, il apprend à **décider**, à **manipuler du texte**, à **répéter** une action, et à **recevoir des informations** au moment où on le lance.

Ce sont les cinq briques avec lesquelles on construira `textfix`. Prenez le temps de bien les comprendre : tout le reste du module repose dessus.

---

## Ce que vous saurez faire en fin de séance

Faire un choix avec `if` et `else`.
Transformer du texte : majuscules, minuscules, recherche.
Ranger plusieurs valeurs dans une liste et la parcourir.
Écrire votre propre fonction.
Lire les mots tapés après le nom du programme.

---

## Déroulé

| Partie | Contenu | Durée |
|---|---|---|
| 1 | Prendre une décision : `if` | 15 min |
| — | Exercice 1 | 30 min |
| 2 | Manipuler du texte | 15 min |
| — | Exercice 2 | 30 min |
| 3 | Les listes et les boucles | 20 min |
| — | Exercice 3 | 35 min |
| 4 | Les fonctions | 15 min |
| — | Exercice 4 | 35 min |
| 5 | Les arguments du programme | 10 min |
| — | Exercice 5 | 35 min |

---

## Organiser vos exercices

Aujourd'hui, vous allez écrire cinq petits programmes. Un programme Go a **un seul** `func main()` par dossier. On crée donc un dossier par exercice, dans un même projet :

```
seance2/
├── go.mod
├── conditions/
│   └── main.go
├── chaines/
│   └── main.go
├── listes/
│   └── main.go
├── fonctions/
│   └── main.go
└── crier/
    └── main.go
```

Création du projet, une seule fois :

```bash
mkdir seance2
cd seance2
go mod init seance2
```

Puis, pour lancer l'exercice d'un dossier :

```bash
go run ./conditions
```

Le `./conditions` veut dire : « le programme qui se trouve dans le dossier `conditions` ».

---

# Partie 1 — Prendre une décision : `if`

## Le principe

`if` veut dire « si ». On exécute un bloc de code **seulement si** une condition est vraie.

```go
age := 17

if age >= 18 {
	fmt.Println("Tu es majeur.")
} else {
	fmt.Println("Tu es mineur.")
}
```

`else` veut dire « sinon ». Ce bloc s'exécute quand la condition est fausse.

## Les comparaisons

| Symbole | Signification |
|---|---|
| `==` | est égal à |
| `!=` | est différent de |
| `>` | est plus grand que |
| `<` | est plus petit que |
| `>=` | est plus grand ou égal à |
| `<=` | est plus petit ou égal à |

**Attention au double égal.** `=` range une valeur dans une variable. `==` compare deux valeurs. Ce n'est pas la même chose.

## Plusieurs cas

```go
if note >= 16 {
	fmt.Println("Très bien")
} else if note >= 10 {
	fmt.Println("Admis")
} else {
	fmt.Println("Non admis")
}
```

Go teste les conditions **dans l'ordre**, et s'arrête à la première qui est vraie.

## Comparer du texte

```go
if motDePasse == "go2026" {
	fmt.Println("Accès autorisé")
}
```

La comparaison est exacte : `"Go2026"` avec une majuscule n'est pas égal à `"go2026"`.

---

## ✅ Exercice 1 — 30 min

Dans `conditions/main.go` :

**1.1** Créez une variable `age`. Affichez « Tu es majeur. » ou « Tu es mineur. ».

**1.2** Changez la valeur de `age` et relancez, pour vérifier les deux cas.

**1.3** Créez une variable `note`. Affichez « Très bien » à partir de 16, « Admis » à partir de 10, « Non admis » en dessous.

**1.4** Testez les notes 18, 10, 9 et 16. Le résultat est-il celui attendu à chaque fois ?

**1.5** Créez une variable `motDePasse`. Affichez « Accès autorisé » si elle vaut `"go2026"`, « Accès refusé » sinon.

**Vérification** : chaque programme donne la bonne réponse pour toutes les valeurs testées, y compris les valeurs pile à la limite, comme 10 et 16.

---

# Partie 2 — Manipuler du texte

## La boîte à outils `strings`

Go fournit une boîte à outils pour le texte, appelée `strings`. On l'importe comme `fmt` :

```go
import (
	"fmt"
	"strings"
)
```

Quand on importe plusieurs boîtes, on les écrit entre parenthèses, une par ligne.

## Les outils utiles

| Outil | Ce qu'il fait | Exemple | Résultat |
|---|---|---|---|
| `strings.ToUpper` | Majuscules | `strings.ToUpper("salut")` | `"SALUT"` |
| `strings.ToLower` | Minuscules | `strings.ToLower("SALUT")` | `"salut"` |
| `strings.Contains` | Contient-il ce texte ? | `strings.Contains("bonjour", "jour")` | `true` |
| `strings.TrimSpace` | Retire les espaces au début et à la fin | `strings.TrimSpace("  oui  ")` | `"oui"` |
| `strings.ReplaceAll` | Remplace un texte par un autre | `strings.ReplaceAll("le chat", "le", "un")` | `"un chat"` |

## Vrai ou faux

`strings.Contains` renvoie `true` (vrai) ou `false` (faux). Ce type de valeur s'appelle un **booléen**, `bool` en Go. On peut s'en servir directement dans un `if` :

```go
if strings.Contains(phrase, "urgent") {
	fmt.Println("Message urgent !")
}
```

## Un outil ne modifie pas la variable

```go
mot := "salut"
strings.ToUpper(mot)
fmt.Println(mot)        // affiche toujours "salut"
```

`strings.ToUpper` **fabrique un nouveau texte**, il ne change pas l'ancien. Pour garder le résultat, il faut le ranger :

```go
mot = strings.ToUpper(mot)
```

---

## ✅ Exercice 2 — 30 min

Dans `chaines/main.go` :

**2.1** Créez une variable `mot` contenant `"Bonjour"`. Affichez-la en majuscules, puis en minuscules.

**2.2** Affichez si `mot` contient `"jour"`, puis si elle contient `"soir"`.

**2.3** Créez la variable `"  trop   d'espaces   ici  "` et affichez-la sans les espaces du début et de la fin. Les espaces du milieu disparaissent-ils ?

**2.4** Transformez `"le chat et le chien"` en `"un chat et un chien"`.

**2.5** Écrivez un `if` qui affiche « Message urgent ! » si une phrase contient le mot `"urgent"`.

**Vérification** : chaque ligne affichée correspond au résultat attendu, et vous savez expliquer pourquoi les espaces du milieu restent en 2.3.

---

# Partie 3 — Les listes et les boucles

## Une liste

Une **liste** range plusieurs valeurs dans une seule variable. En Go, on l'appelle une **slice**.

```go
fruits := []string{"pomme", "banane", "cerise"}
```

`[]string` veut dire « une liste de textes ».

## Lire un élément

Chaque élément a un **numéro**, appelé **index**. On commence à compter à **zéro**.

| Index | 0 | 1 | 2 |
|---|---|---|---|
| Valeur | pomme | banane | cerise |

```go
fruits[0]       // "pomme"
fruits[2]       // "cerise"
len(fruits)     // 3, le nombre d'éléments
```

Demander `fruits[3]` fait **planter** le programme : il n'y a pas d'élément numéro 3.

## Ajouter un élément

```go
fruits = append(fruits, "kiwi")
```

`append` ajoute à la fin, et renvoie la nouvelle liste. Il faut toujours écrire `fruits =` devant, sinon l'ajout est perdu.

## Parcourir la liste : la boucle `for`

Une **boucle** répète une action pour chaque élément :

```go
for i, fruit := range fruits {
	fmt.Println(i+1, "-", fruit)
}
```

À chaque tour, `i` contient l'index et `fruit` contient la valeur. On affiche `i+1` pour numéroter à partir de 1, ce qui est plus naturel à lire.

```
1 - pomme
2 - banane
3 - cerise
4 - kiwi
```

## Quand on n'a pas besoin de l'index

```go
for _, fruit := range fruits {
	fmt.Println(fruit)
}
```

Le tiret bas `_` veut dire « je ne m'en sers pas ». Go interdit les variables inutilisées, d'où ce symbole.

## Découper une phrase en mots

C'est l'outil le plus important pour notre projet :

```go
mots := strings.Fields("le   chat   dort")
// ["le", "chat", "dort"]
```

`strings.Fields` coupe à chaque espace, en ignorant les espaces en trop. Le résultat est une liste.

Et l'inverse, recoller une liste en phrase :

```go
phrase := strings.Join(mots, " ")
// "le chat dort"
```

---

## ✅ Exercice 3 — 35 min

Dans `listes/main.go` :

**3.1** Créez la liste `fruits` avec trois fruits. Affichez le premier et le nombre d'éléments.

**3.2** Ajoutez un quatrième fruit avec `append`, puis affichez toute la liste.

**3.3** Parcourez la liste avec `for` et affichez chaque fruit, numéroté à partir de 1.

**3.4** Créez une nouvelle liste vide, `var enMajuscules []string`, et remplissez-la avec les fruits écrits en majuscules.

**3.5** Découpez la phrase `"le   chat   dort"` avec `strings.Fields`, affichez le nombre de mots, puis recollez les mots avec un seul espace.

**Vérification** : la liste en majuscules contient bien tous les fruits, et la phrase recollée n'a plus d'espaces en trop.

---

# Partie 4 — Les fonctions

## Pourquoi

Une **fonction** est un morceau de code auquel on donne un nom, pour pouvoir le réutiliser. Vous en utilisez déjà : `fmt.Println` et `strings.ToUpper` sont des fonctions écrites par d'autres.

Aujourd'hui, vous écrivez les vôtres.

## Une fonction qui fait une action

```go
func direBonjour(prenom string) {
	fmt.Println("Bonjour", prenom, "!")
}
```

`prenom string` est le **paramètre** : l'information dont la fonction a besoin, avec son type.

On l'**appelle** depuis `main` :

```go
direBonjour("Léa")
direBonjour("Sam")
```

## Une fonction qui renvoie un résultat

```go
func doubler(nombre int) int {
	return nombre * 2
}
```

Le `int` après les parenthèses indique le type du **résultat**. `return` renvoie ce résultat à celui qui a appelé la fonction.

```go
resultat := doubler(21)
fmt.Println(resultat)       // 42
```

## Une fonction qui renvoie vrai ou faux

```go
func estLong(mot string) bool {
	return len(mot) > 6
}
```

`len(mot) > 6` vaut `true` ou `false` : la fonction renvoie directement ce résultat.

## Où écrire les fonctions

En dehors de `main`, au même niveau, dans le même fichier. L'ordre n'a pas d'importance : Go les trouve où qu'elles soient.

---

## ✅ Exercice 4 — 35 min

Dans `fonctions/main.go` :

**4.1** Écrivez `direBonjour(prenom string)` et appelez-la deux fois avec des prénoms différents.

**4.2** Écrivez `doubler(nombre int) int` et affichez `doubler(21)`.

**4.3** Écrivez `crier(texte string) string`, qui renvoie le texte en majuscules suivi de ` !`. Rangez le résultat dans une variable, puis affichez-la.

**4.4** Écrivez `estLong(mot string) bool`, qui renvoie `true` si le mot a plus de 6 lettres. Testez avec `"chat"` et `"ordinateur"`.

**4.5** Utilisez `estLong` dans un `if` pour afficher « mot long » ou « mot court ».

**Vérification** : toutes vos fonctions sont écrites en dehors de `main`, et chacune donne le résultat attendu.

---

# Partie 5 — Les arguments du programme

## Ce que c'est

Quand on lance un programme, on peut écrire des mots après son nom :

```bash
go run ./crier attention au départ
```

Ces mots s'appellent des **arguments**. Le programme les reçoit dans une liste nommée `os.Args`, qui vient de la boîte à outils `os` (pour « operating system », le système).

## Ce que contient `os.Args`

| Index | Contenu |
|---|---|
| `os.Args[0]` | Le nom du programme lui-même |
| `os.Args[1]` | `attention` |
| `os.Args[2]` | `au` |
| `os.Args[3]` | `départ` |

**Le premier élément est toujours le nom du programme.** Les vrais arguments commencent à l'index 1.

Pour obtenir seulement les arguments :

```go
mots := os.Args[1:]
```

`[1:]` veut dire « à partir de l'index 1, jusqu'à la fin ».

## Vérifier avant de lire

Si l'utilisateur ne tape aucun mot, `os.Args[1]` n'existe pas, et le programme plante. On vérifie donc d'abord :

```go
if len(os.Args) < 2 {
	fmt.Println("Erreur : il faut au moins un mot.")
	os.Exit(1)
}
```

`os.Exit(1)` arrête immédiatement le programme. Le nombre 1 signale au système qu'il y a eu un problème ; 0 voudrait dire que tout s'est bien passé.

---

## ✅ Exercice 5 — 35 min

Dans `crier/main.go`, écrivez un programme qui crie les mots qu'on lui donne :

```bash
go run ./crier attention au départ
```

```
ATTENTION AU DÉPART !
```

**5.1** Récupérez les arguments avec `os.Args[1:]`.

**5.2** Recollez-les en une phrase avec `strings.Join`.

**5.3** Affichez la phrase en majuscules, suivie de ` !`.

**5.4** Sans argument, le programme affiche un message d'erreur, un exemple d'utilisation, et s'arrête avec `os.Exit(1)`.

**5.5** Testez avec des mots accentués, puis sans aucun mot.

**Vérification** : `go run ./crier bonjour élise` affiche `BONJOUR ÉLISE !`, et le lancement sans mot affiche un message clair au lieu de planter.

---

# À retenir

## Décider

`if` condition `{ … } else { … }`.
`==` compare, `=` range une valeur.

## Le texte

`strings.ToUpper`, `ToLower`, `Contains`, `TrimSpace`, `ReplaceAll`.
Un outil fabrique un nouveau texte : il faut ranger le résultat.

## Les listes

`[]string{…}` crée une liste. Les index commencent à 0.
`append` ajoute, et il faut récupérer son résultat.
`for i, x := range liste` parcourt la liste.
`strings.Fields` découpe une phrase en mots, `strings.Join` les recolle.

## Les fonctions

`func nom(parametre type) typeDuResultat { return … }`.
Elles s'écrivent en dehors de `main`.

## Les arguments

`os.Args[0]` est le nom du programme, les arguments commencent à 1.
On vérifie leur nombre avant de les lire.

## Les erreurs qui coûtent le plus de temps

| Symptôme | Cause |
|---|---|
| `main redeclared in this block` | Deux `func main()` dans le même dossier |
| `index out of range` | Lecture d'un élément qui n'existe pas, souvent `os.Args[1]` sans argument |
| La liste reste vide | `append(liste, x)` sans `liste =` devant |
| Le texte n'est pas en majuscules | `strings.ToUpper(mot)` sans ranger le résultat |
| `"strings" imported and not used` | Boîte à outils importée mais pas encore utilisée |
| `missing return` | Fonction qui annonce un résultat mais n'a pas de `return` |
| `cannot use … (untyped string constant) as int value` | Un texte passé là où la fonction attend un nombre |

---

# Prochaine séance

**Fichiers et erreurs (4 h)** : on démarre le projet `textfix`. Votre programme lira un fichier texte, le nettoiera, et écrira le résultat dans un autre fichier.
