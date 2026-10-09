# Séance 3 — Fichiers et erreurs : le projet commence

**Application CLI & Parsing en Go — Bachelor 1 — Ycode A2627_4658**
**Séance 3 · 4 h**
**Prérequis : séances 1 et 2**

Vous avez maintenant toutes les briques de base. Aujourd'hui, on démarre le vrai projet du module : **textfix**.

À la fin de la séance, votre programme lira un fichier texte, supprimera les espaces en trop, et écrira le résultat dans un nouveau fichier. Et surtout, il saura réagir proprement quand quelque chose se passe mal.

En dernière partie, vous appliquerez tout cela au projet sur lequel vous serez évalués, **CampusLink** : votre programme importera une liste de salles depuis un fichier CSV.

---

## Ce que vous saurez faire en fin de séance

Comprendre comment Go signale qu'une opération a échoué.
Lire le contenu d'un fichier.
Écrire du texte dans un fichier.
Afficher des messages d'erreur clairs.
Répartir votre code dans plusieurs fichiers.
Transformer un fichier CSV en une liste de salles, et signaler les lignes invalides.

---

## Déroulé

| Partie | Contenu | Durée |
|---|---|---|
| 1 | Le projet `textfix` | 15 min |
| 2 | Quand quelque chose échoue | 15 min |
| — | Exercice 1 | 25 min |
| 3 | Lire un fichier | 15 min |
| — | Exercice 2 | 25 min |
| 4 | Écrire un fichier | 10 min |
| — | Exercice 3 | 25 min |
| 5 | Des messages d'erreur clairs | 10 min |
| — | Exercice 4 | 25 min |
| — | Exercice 5 — synthèse | 15 min |
| 6 | CampusLink : importer des salles | 20 min |
| — | Exercice 6 — CampusLink | 40 min |

---

# Partie 1 — Le projet `textfix`

## Ce qu'il fera, à la fin du module

On le lance avec deux noms de fichiers :

```bash
go run . entree.txt sortie.txt
```

Il lit `entree.txt`, le corrige, et écrit le résultat dans `sortie.txt`.

Le texte d'entrée contient des **marqueurs** entre parenthèses. Chaque marqueur modifie le mot qui est juste avant lui, puis disparaît :

| Entrée | Sortie |
|---|---|
| `bonjour (up) tout le monde` | `BONJOUR tout le monde` |
| `CE MOT (low) crie` | `CE mot crie` |
| `paris (cap) est belle` | `Paris est belle` |
| `des    espaces   en   trop` | `des espaces en trop` |

## Le plan

| Séance | Ce que fait `textfix` |
|---|---|
| 3 — aujourd'hui | Lire un fichier, supprimer les espaces en trop, écrire le résultat |
| 4 | Appliquer les marqueurs `(up)`, `(low)` et `(cap)` |
| 5 | Être bien rangé, testé, et documenté |

## Créer le projet

En dehors des dossiers des séances précédentes :

```bash
mkdir textfix
cd textfix
go mod init textfix
```

Ouvrez le dossier `textfix` dans VS Code. Vous y travaillerez jusqu'à la fin du module.

---

# Partie 2 — Quand quelque chose échoue

## Le problème

Certaines opérations peuvent rater. Lire un fichier qui n'existe pas. Transformer le texte `"douze"` en nombre. Écrire dans un dossier interdit.

Un programme sérieux ne plante pas dans ces cas-là. Il **prévient** l'utilisateur avec un message clair.

## La méthode de Go : deux réponses

En Go, une fonction qui peut échouer renvoie **deux** choses : le résultat, et une **erreur**.

Prenons `strconv.Atoi`, qui transforme un texte en nombre :

```go
nombre, err := strconv.Atoi("21")
```

On récupère deux variables d'un coup, séparées par une virgule. `nombre` contient 21, et `err` contient l'erreur éventuelle.

## `nil` : rien à signaler

Quand tout s'est bien passé, `err` vaut **`nil`**. `nil` veut dire « rien », « vide ».

Quand quelque chose a raté, `err` contient une description du problème.

## Toujours vérifier

```go
nombre, err := strconv.Atoi(texte)
if err != nil {
	fmt.Println("Erreur :", texte, "n'est pas un nombre entier.")
	os.Exit(1)
}
fmt.Println("Le double est", nombre*2)
```

`err != nil` se lit : « s'il y a une erreur ». Dans ce cas, on affiche un message et on arrête. Sinon, on continue normalement.

**Vous allez écrire `if err != nil` des dizaines de fois.** C'est normal : c'est la façon dont Go vous oblige à penser aux problèmes au moment où ils peuvent arriver.

## Une image pour retenir

Imaginez un livreur qui vous apporte un colis **et** une fiche. Si la fiche est vide, le colis est bon. Si la fiche porte une remarque — « colis abîmé » —, vous ne l'ouvrez pas sans réfléchir.

Le colis, c'est le résultat. La fiche, c'est `err`.

---

## ✅ Exercice 1 — 25 min

Créez un dossier `seance3` avec son propre `go mod init seance3`, et un sous-dossier `nombre`.

**1.1** Écrivez un programme qui reçoit un nombre en argument et affiche son double :

```bash
go run ./nombre 21
```

```
Le double de 21 est 42
```

**1.2** Utilisez `strconv.Atoi` pour transformer l'argument en nombre. Il faut importer `strconv`.

**1.3** Vérifiez l'erreur. Si l'argument n'est pas un nombre, affichez un message clair et arrêtez le programme.

**1.4** Vérifiez aussi qu'il y a bien un argument, comme en séance 2.

**1.5** Testez avec `21`, avec `douze`, et sans argument.

**Vérification** : les trois cas donnent soit le bon résultat, soit un message clair, et le programme ne plante jamais.

---

# Partie 3 — Lire un fichier

## `os.ReadFile`

```go
contenu, err := os.ReadFile("entree.txt")
```

Même principe que `Atoi` : deux réponses, le contenu et une erreur.

## Une petite conversion

`os.ReadFile` ne renvoie pas directement du texte, mais des **octets** : la forme brute sous laquelle l'ordinateur stocke les fichiers. Pour obtenir du texte, on convertit :

```go
texte := string(contenu)
```

## Une fonction pour lire

On range ce code dans une fonction, qui renvoie elle aussi deux choses : le texte, et une erreur.

```go
func lireFichier(chemin string) (string, error) {
	contenu, err := os.ReadFile(chemin)
	if err != nil {
		return "", fmt.Errorf("impossible de lire le fichier %s", chemin)
	}
	return string(contenu), nil
}
```

**`(string, error)`** — la fonction annonce deux résultats : un texte et une erreur.

**`return "", fmt.Errorf(…)`** — en cas de problème, on renvoie un texte vide et une erreur. `fmt.Errorf` fabrique une erreur avec un message ; `%s` est remplacé par le nom du fichier.

**`return string(contenu), nil`** — tout va bien : on renvoie le texte, et `nil` comme erreur.

## Plusieurs fichiers dans un projet

On écrit cette fonction dans un **nouveau fichier**, `fichiers.go`, à côté de `main.go`. Il commence lui aussi par `package main`.

Tous les fichiers `package main` d'un même dossier forment **un seul programme**. Une fonction écrite dans `fichiers.go` peut être utilisée dans `main.go` directement.

C'est pour cela qu'on lance avec `go run .` : le point inclut tous les fichiers du dossier.

---

## ✅ Exercice 2 — 25 min

Dans votre projet `textfix` :

**2.1** Créez `fichiers.go` avec la fonction `lireFichier`.

**2.2** Dans `main.go`, vérifiez qu'on a bien reçu **deux** arguments : le fichier d'entrée et le fichier de sortie. Sinon, affichez un message et arrêtez.

**2.3** Lisez le fichier d'entrée avec `lireFichier`, et affichez son contenu à l'écran.

**2.4** Créez un fichier `entree.txt` avec quelques lignes de texte. Testez.

**2.5** Testez avec un nom de fichier qui n'existe pas. Que se passe-t-il ?

**Vérification** : le contenu de `entree.txt` s'affiche, et un fichier absent donne un message clair au lieu d'un plantage.

---

# Partie 4 — Écrire un fichier

## `os.WriteFile`

```go
err := os.WriteFile("sortie.txt", []byte(texte), 0644)
```

Trois informations :

`"sortie.txt"` — le nom du fichier à créer.
`[]byte(texte)` — le texte, reconverti en octets.
`0644` — les **permissions** du fichier : qui a le droit de le lire et de le modifier. `0644` est la valeur habituelle pour un fichier texte. Retenez-la telle quelle, sans chercher à la comprendre pour l'instant.

Cette fois, il n'y a qu'une seule réponse : l'erreur. Il n'y a pas de résultat à renvoyer.

## Attention : le fichier est remplacé

Si `sortie.txt` existe déjà, **son contenu est entièrement effacé** et remplacé. Aucune confirmation n'est demandée.

C'est pourquoi notre programme refusera d'écrire dans le fichier qu'il est en train de lire : on perdrait le texte d'origine.

## Une fonction pour écrire

```go
func ecrireFichier(chemin string, texte string) error {
	err := os.WriteFile(chemin, []byte(texte), 0644)
	if err != nil {
		return fmt.Errorf("impossible d'écrire le fichier %s", chemin)
	}
	return nil
}
```

---

## ✅ Exercice 3 — 25 min

**3.1** Ajoutez `ecrireFichier` dans `fichiers.go`.

**3.2** Modifiez `main.go` : au lieu d'afficher le contenu, le programme le **copie** dans le fichier de sortie.

**3.3** Vérifiez dans VS Code que le fichier de sortie a bien été créé, et qu'il est identique à l'entrée.

**3.4** Ajoutez une vérification : si le fichier d'entrée et le fichier de sortie ont le même nom, le programme refuse avec un message.

**3.5** Affichez un message à la fin quand tout s'est bien passé : `Terminé : sortie.txt a été créé.`

**Vérification** : le fichier de sortie est une copie exacte de l'entrée, et le programme refuse d'écrire dans son propre fichier d'entrée.

---

# Partie 5 — Des messages d'erreur clairs

## Deux canaux d'affichage

Un programme lancé dans un terminal a deux canaux pour afficher du texte :

**La sortie normale** — pour les résultats. C'est ce qu'utilise `fmt.Println`.

**La sortie d'erreur** — pour les messages d'erreur. On l'utilise avec `fmt.Fprintln(os.Stderr, …)` :

```go
fmt.Fprintln(os.Stderr, "Erreur : fichier introuvable")
```

À l'écran, vous ne verrez aucune différence. Mais c'est la convention de tous les programmes en ligne de commande, et nous la respecterons dans `textfix`.

## Un bon message d'erreur

Il dit **ce qui ne va pas** et **sur quoi**.

| Mauvais | Bon |
|---|---|
| `Erreur` | `Erreur : impossible de lire le fichier notes.txt` |
| `Ça marche pas` | `Erreur : il faut deux noms de fichiers.` |

Et quand l'utilisateur s'est trompé dans la commande, on lui montre comment l'utiliser :

```
Erreur : il faut deux noms de fichiers.
Utilisation : textfix <fichier_entree> <fichier_sortie>
```

Les chevrons `< >` signalent qu'il faut remplacer le texte par un vrai nom de fichier.

## Supprimer les espaces en trop

Dernière étape de la séance : faire un premier vrai travail sur le texte. On le traite **ligne par ligne**, pour garder la mise en page.

```go
func transformer(texte string) string {
	lignes := strings.Split(texte, "\n")

	for i, ligne := range lignes {
		ligne = strings.TrimRight(ligne, "\r")
		mots := strings.Fields(ligne)
		lignes[i] = strings.Join(mots, " ")
	}

	return strings.Join(lignes, "\n")
}
```

**`strings.Split(texte, "\n")`** — découpe le texte en lignes. `"\n"` représente un retour à la ligne.

**`strings.TrimRight(ligne, "\r")`** — les fichiers créés sous Windows ajoutent un caractère invisible, `\r`, en fin de ligne. On le retire.

**`strings.Fields` puis `strings.Join`** — découper en mots puis recoller avec un seul espace, comme en séance 2.

**`lignes[i] = …`** — on remplace la ligne dans la liste. C'est important : modifier seulement la variable `ligne` ne changerait pas la liste.

---

## ✅ Exercice 4 — 25 min

**4.1** Remplacez tous les `fmt.Println` de messages d'erreur par `fmt.Fprintln(os.Stderr, …)`.

**4.2** Ajoutez un message d'utilisation quand le nombre d'arguments est incorrect.

**4.3** Créez un troisième fichier, `regles.go`, avec la fonction `transformer`.

**4.4** Dans `main.go`, appelez `transformer` entre la lecture et l'écriture : le fichier de sortie contient désormais le texte nettoyé.

**4.5** Testez avec un fichier qui contient des espaces en trop, des lignes vides, et plusieurs lignes.

**Vérification** : les espaces en trop disparaissent, les lignes vides sont conservées, et le projet compte trois fichiers : `main.go`, `fichiers.go` et `regles.go`.

---

## ✅ Exercice 5 — Synthèse — 15 min

Vérifiez que votre programme réagit bien dans toutes les situations.

**5.1** Recopiez et complétez ce tableau en lançant chaque commande :

| Commande | Message obtenu |
|---|---|
| `go run .` | |
| `go run . entree.txt` | |
| `go run . absent.txt sortie.txt` | |
| `go run . entree.txt entree.txt` | |
| `go run . entree.txt dossier_absent/sortie.txt` | |
| `go run . entree.txt sortie.txt` | |

**5.2** Pour chaque message, demandez-vous : quelqu'un qui n'a jamais vu le code comprend-il ce qu'il doit corriger ? Améliorez les messages qui ne passent pas ce test.

**5.3** Demandez à votre voisin de lancer votre programme en faisant exprès de se tromper. Votre programme a-t-il planté une seule fois ?

**Vérification** : seule la dernière commande crée un fichier, et chaque autre cas donne un message clair, sans plantage.

---

# Partie 6 — CampusLink : importer des salles

## Le projet du module

`textfix` vous a servi à apprendre. Le projet sur lequel vous serez évalués s'appelle **CampusLink** : une application qui gère les salles et les équipements du campus.

Votre partie de CampusLink est un programme en ligne de commande qui lit des fichiers de données, les vérifie, et en tire des informations. Tout ce que vous venez de faire avec `textfix` sert directement : lire un fichier, vérifier les erreurs, répartir le code dans plusieurs fichiers.

Il se lance avec une **commande**, puis un nom de fichier :

```bash
go run . import-salles donnees/salles.csv
```

## Un fichier CSV

Les salles sont décrites dans un fichier **CSV**. CSV veut dire « valeurs séparées par des virgules ». C'est un tableau écrit en texte : une ligne par salle, et une virgule entre chaque colonne. Un tableur comme Excel sait l'ouvrir.

```
id,nom,batiment,capacite
S101,Salle Turing,A,30
S102,Salle Lovelace,A,24
S201,Amphi Hopper,B,120
```

La première ligne s'appelle l'**en-tête** : elle donne le nom des colonnes. Ce n'est pas une salle, on ne la traite pas.

Pour séparer les colonnes d'une ligne, on la découpe à chaque virgule :

```go
champs := strings.Split("S101,Salle Turing,A,30", ",")
```

| Index | 0 | 1 | 2 | 3 |
|---|---|---|---|---|
| Valeur | `S101` | `Salle Turing` | `A` | `30` |

Les quatre champs sont du **texte**, y compris `30`. Pour s'en servir comme d'un nombre, il faut le convertir avec `strconv.Atoi`, comme à l'exercice 1.

## Une fiche pour chaque salle : la structure

Une salle, ce n'est pas une seule information, c'est quatre : un identifiant, un nom, un bâtiment, une capacité. Pour les garder ensemble, on crée une **structure**. Imaginez une fiche cartonnée avec quatre cases, chacune avec son nom.

```go
type Salle struct {
	ID       string
	Nom      string
	Batiment string
	Capacite int
}
```

`type Salle struct` crée un nouveau type, nommé `Salle`. Chaque ligne entre les accolades est une **case** de la fiche, avec son nom et son type. La capacité est un `int` : c'est un nombre.

On remplit une fiche en nommant chaque case :

```go
turing := Salle{ID: "S101", Nom: "Salle Turing", Batiment: "A", Capacite: 30}
```

Et on lit une case avec un point :

```go
fmt.Println(turing.Nom)        // Salle Turing
fmt.Println(turing.Capacite)   // 30
```

Une liste de salles s'écrit `[]Salle`, comme une liste de textes s'écrit `[]string`.

## Choisir une commande

CampusLink aura plusieurs commandes. Le premier argument dit laquelle lancer, et les suivants sont les fichiers :

| Index | Contenu |
|---|---|
| `os.Args[1]` | la commande, par exemple `import-salles` |
| `os.Args[2:]` | le reste : les noms de fichiers |

```go
commande := os.Args[1]
var err error

if commande == "aide" {
	fmt.Println(usage)
} else if commande == "import-salles" {
	err = commandeImportSalles(os.Args[2:])
} else {
	err = fmt.Errorf("commande inconnue : %s", commande)
}

if err != nil {
	fmt.Fprintln(os.Stderr, "Erreur :", err)
	os.Exit(1)
}
```

`var err error` crée une variable d'erreur **vide**, `nil`. Chaque commande y range son erreur éventuelle. À la fin, un seul `if` affiche l'erreur et arrête le programme : on ne répète pas ce code dans chaque commande.

## Une ligne invalide n'arrête pas l'import

Un fichier de 200 salles peut contenir une ligne mal écrite. On ne refuse pas tout le fichier pour autant : on **ignore** la ligne, on la **signale** avec son numéro, et on continue.

On sépare le travail en deux fonctions.

`lireSalle(ligne)` transforme **une** ligne en salle, ou renvoie une erreur :

```go
capacite, err := strconv.Atoi(champs[3])
if err != nil {
	return Salle{}, fmt.Errorf("la capacité %q n'est pas un nombre", champs[3])
}
```

`Salle{}` est une fiche vide : en cas d'erreur, il faut quand même renvoyer une salle, mais elle ne sera pas utilisée. `%q` affiche la valeur entre guillemets.

Les vérifications à faire, dans cet ordre :

| Vérification | Message |
|---|---|
| La ligne a 4 champs | `4 champs attendus, 3 trouvés` |
| L'identifiant n'est pas vide | `l'identifiant est vide` |
| La capacité est un nombre | `la capacité "trente" n'est pas un nombre` |
| La capacité est positive | `la capacité doit être positive, 0 trouvé` |

`lireSalles(texte)` parcourt toutes les lignes. Elle renvoie **deux** listes : les salles valides, et les problèmes rencontrés.

```go
salle, err := lireSalle(ligne)
if err != nil {
	problemes = append(problemes, fmt.Sprintf("ligne %d : %v", i+1, err))
	continue
}
salles = append(salles, salle)
```

`fmt.Sprintf` fabrique un texte avec des emplacements, sans l'afficher : `%d` reçoit un nombre, `%v` reçoit l'erreur. Le numéro de ligne est `i+1`, puisque les index commencent à 0 alors que les lignes d'un fichier se comptent à partir de 1.

---

## ✅ Exercice 6 — CampusLink : la commande `import-salles` — 40 min

**6.1** À côté du dossier `textfix`, créez un dossier `campuslink` et initialisez-le avec `go mod init campuslink`. Copiez-y votre fichier `fichiers.go` : il servira tel quel.

**6.2** Créez un dossier `donnees` avec deux fichiers.

`donnees/salles.csv` :

```
id,nom,batiment,capacite
S101,Salle Turing,A,30
S102,Salle Lovelace,A,24
S201,Amphi Hopper,B,120
S202,Salle Hamilton,B,18
```

`donnees/salles-erreurs.csv` :

```
id,nom,batiment,capacite
S101,Salle Turing,A,30
S102,Salle Lovelace,A
S103,Salle Knuth,A,trente
,Salle sans nom,A,20
S104,Salle vide,A,0

S201,Amphi Hopper,B,120
```

**6.3** Créez `salles.go` avec la structure `Salle`.

**6.4** Dans `salles.go`, écrivez `lireSalle(ligne string) (Salle, error)` avec les quatre vérifications du tableau. Retirez les espaces autour de chaque champ avec `strings.TrimSpace`.

**6.5** Dans `salles.go`, écrivez `lireSalles(texte string) ([]Salle, []string)`. Elle ignore l'en-tête et les lignes vides, et signale chaque ligne invalide avec son numéro.

**6.6** Dans `main.go`, écrivez les commandes `aide` et `import-salles`. `import-salles` affiche chaque salle valide sur une ligne, les lignes ignorées sur la sortie d'erreur, puis un résumé.

**6.7** Testez les deux fichiers, puis le programme sans commande, avec une commande inconnue, et avec `import-salles` sans nom de fichier.

**Vérification** : `go run . import-salles donnees/salles-erreurs.csv` affiche :

```
S101 - Salle Turing - bâtiment A - 30 places
S201 - Amphi Hopper - bâtiment B - 120 places
Ligne ignorée : ligne 3 : 4 champs attendus, 3 trouvés
Ligne ignorée : ligne 4 : la capacité "trente" n'est pas un nombre
Ligne ignorée : ligne 5 : l'identifiant est vide
Ligne ignorée : ligne 6 : la capacité doit être positive, 0 trouvé
2 salle(s) importée(s), 4 ligne(s) ignorée(s).
```

et aucune commande mal tapée ne fait planter le programme.

---

# À retenir

## Les erreurs

Une opération qui peut échouer renvoie un résultat **et** une erreur.
`nil` veut dire « pas d'erreur ».
`if err != nil` juste après chaque opération qui peut échouer.

## Les fichiers

`os.ReadFile(chemin)` lit, puis `string(contenu)` donne le texte.
`os.WriteFile(chemin, []byte(texte), 0644)` écrit, et **remplace** le fichier s'il existe.

## L'organisation

Plusieurs fichiers `package main` dans un dossier forment un seul programme.
`main.go` pour le déroulé, `fichiers.go` pour les fichiers, `regles.go` pour le traitement du texte.
On lance avec `go run .` pour inclure tous les fichiers.

## Les messages

Les erreurs sur `os.Stderr`.
Un message dit ce qui ne va pas, et sur quel fichier.

## CampusLink

Un fichier CSV : une ligne par élément, une virgule entre les colonnes, un en-tête en première ligne.
`strings.Split(ligne, ",")` découpe une ligne en champs, qui sont tous du texte.
Une structure regroupe plusieurs informations : `type Salle struct { … }`, puis `salle.Nom`.
`os.Args[1]` est la commande, `os.Args[2:]` ses fichiers.
Une ligne invalide est ignorée et signalée avec son numéro : l'import continue.

## Les erreurs qui coûtent le plus de temps

| Symptôme | Cause |
|---|---|
| `undefined: lireFichier` | Lancé avec `go run main.go` au lieu de `go run .`, ou `fichiers.go` sans `package main` |
| `declared and not used` sur `err` | L'erreur est récupérée mais jamais vérifiée |
| `assignment mismatch: 1 variable but os.ReadFile returns 2 values` | Une seule variable pour une fonction qui renvoie deux choses |
| `cannot use contenu (variable of type []byte) as string value` | Conversion `string(contenu)` oubliée |
| Le fichier d'entrée a été effacé | Même nom pour l'entrée et la sortie, sans vérification |
| Le fichier de sortie est identique à l'entrée | `ligne = …` au lieu de `lignes[i] = …` dans la boucle |
| `index out of range` au lancement | Lecture de `os.Args[1]` avant de vérifier le nombre d'arguments |
| `unknown field Capacité in struct literal of type Salle` | Nom de case mal écrit : un accent, ou une majuscule différente de la structure |
| `s.capacite undefined` | Minuscule au lieu de majuscule : la case s'appelle `Capacite` |
| `cannot use "30" (untyped string constant) as int value in struct literal` | Texte rangé dans une case `int` : il faut d'abord convertir avec `strconv.Atoi` |
| L'en-tête `id,nom,batiment,capacite` est signalé comme ligne invalide | La ligne 0 n'est pas sautée |
| Toutes les lignes ont un numéro décalé d'un cran | `i` affiché au lieu de `i+1` |
| L'import s'arrête à la première ligne invalide | `return` au lieu de `continue` dans la boucle de `lireSalles` |

---

# Prochaine séance

**Le parser (4 h)** : `textfix` reconnaîtra les marqueurs `(up)`, `(low)` et `(cap)`, et transformera le mot qui les précède. CampusLink recevra trois nouvelles commandes : contrôler, compter et exporter les salles.
