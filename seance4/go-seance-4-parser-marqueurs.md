# Séance 4 — Le parser : appliquer les marqueurs

**Application CLI & Parsing en Go — Bachelor 1 — Ycode A2627_4658**
**Séance 4 · 4 h**
**Prérequis : séances 1 à 3**

`textfix` sait lire un fichier, le nettoyer et l'écrire. Aujourd'hui, il fait enfin son vrai travail : repérer les marqueurs `(up)`, `(low)` et `(cap)`, et transformer le mot qui les précède.

Ce travail porte un nom : **parser** un texte. Parser, c'est lire un texte morceau par morceau, et décider pour chaque morceau de ce qu'on en fait. C'est ce que font tous les programmes qui lisent du texte : un navigateur qui lit une page web, un correcteur d'orthographe, un tableur qui lit un fichier.

En dernière partie, CampusLink reçoit trois nouvelles commandes : contrôler les salles, calculer des statistiques, et exporter les données.

---

## Ce que vous saurez faire en fin de séance

Découper un texte en lignes, puis en mots.
Construire une nouvelle liste de mots au fur et à mesure.
Appliquer un marqueur au mot précédent.
Mettre une majuscule à un mot, même accentué.
Empêcher votre programme de planter sur un texte mal écrit.
Compter des éléments par catégorie avec une map, et écrire un fichier d'export.

---

## Déroulé

| Partie | Contenu | Durée |
|---|---|---|
| 1 | Découper en lignes et en mots | 15 min |
| — | Exercice 1 | 20 min |
| 2 | Le principe du parser | 20 min |
| — | Exercice 2 | 35 min |
| 3 | `(low)`, `(cap)` et les lettres accentuées | 15 min |
| — | Exercice 3 | 25 min |
| 4 | Ne jamais planter | 10 min |
| — | Exercice 4 | 25 min |
| 5 | Brancher le parser dans `textfix` | 10 min |
| — | Exercice 5 | 15 min |
| 6 | CampusLink : contrôler, compter, exporter | 15 min |
| — | Exercice 6 — CampusLink | 35 min |

---

# Partie 1 — Découper en lignes et en mots

## Deux découpages

Un texte, c'est des **lignes**. Une ligne, c'est des **mots**. On découpe donc en deux temps.

```go
lignes := strings.Split(texte, "\n")
```

`strings.Split` coupe le texte à chaque retour à la ligne `"\n"`. On obtient une liste de lignes.

```go
mots := strings.Fields(ligne)
```

`strings.Fields` coupe la ligne à chaque espace, en ignorant les espaces en trop. On obtient une liste de mots.

## Une boucle dans une boucle

Pour parcourir tous les mots de tout le texte, on met une boucle à l'intérieur d'une autre :

```go
for _, ligne := range lignes {
	mots := strings.Fields(ligne)
	for _, mot := range mots {
		fmt.Println(mot)
	}
}
```

La boucle extérieure passe d'une ligne à l'autre. Pour chaque ligne, la boucle intérieure passe d'un mot à l'autre.

## `Split` et `Fields` ne sont pas interchangeables

`strings.Split(ligne, " ")` coupe à **chaque** espace : deux espaces à la suite produisent un « mot » vide entre eux. `strings.Fields` ignore les espaces en trop.

**Pour les lignes : `Split` avec `"\n"`. Pour les mots : `Fields`.**

---

## ✅ Exercice 1 — 20 min

Créez un dossier `seance4` avec `go mod init seance4`, et un sous-dossier `decouper`.

**1.1** Créez une variable contenant ce texte de deux lignes :

```go
texte := "le chat dort\nle chien   joue"
```

**1.2** Découpez-le en lignes et affichez le nombre de lignes.

**1.3** Pour chaque ligne, affichez son numéro et son nombre de mots.

**1.4** Affichez chaque mot sur une ligne, précédé d'un tiret :

```
Nombre de lignes : 2
Ligne 1 : 3 mots
  - le
  - chat
  - dort
Ligne 2 : 3 mots
  - le
  - chien
  - joue
```

**Vérification** : la deuxième ligne compte bien 3 mots, malgré les espaces en trop.

---

# Partie 2 — Le principe du parser

## L'idée

On parcourt les mots un par un, et on construit une **nouvelle liste**, appelée `resultat`. Elle est vide au départ.

Pour chaque mot lu, deux possibilités :

**C'est un mot normal** : on l'**ajoute** à la fin de `resultat`.

**C'est un marqueur** : on ne l'ajoute pas. À la place, on **modifie le dernier mot** de `resultat`.

## Pas à pas

Texte : `il est (up) parti`

| Mot lu | Ce qu'on fait | `resultat` après |
|---|---|---|
| `il` | mot normal : on l'ajoute | `il` |
| `est` | mot normal : on l'ajoute | `il est` |
| `(up)` | marqueur : on modifie le dernier mot | `il EST` |
| `parti` | mot normal : on l'ajoute | `il EST parti` |

Résultat : `il EST parti`. Le marqueur a disparu, et le mot qu'il visait a changé.

**Prenez le temps de bien comprendre ce tableau.** Tout le parser repose sur cette idée.

## Le dernier mot de la liste

Le dernier élément d'une liste est à l'index `len(resultat) - 1`. Une liste de 3 mots a ses mots aux index 0, 1 et 2 : le dernier est bien à l'index 3 − 1 = 2.

```go
dernier := len(resultat) - 1
resultat[dernier] = strings.ToUpper(resultat[dernier])
```

La deuxième ligne remplace le dernier mot par sa version en majuscules.

## Le code

```go
func appliquerMarqueurs(mots []string) []string {
	var resultat []string

	for _, mot := range mots {
		if mot == "(up)" {
			dernier := len(resultat) - 1
			resultat[dernier] = strings.ToUpper(resultat[dernier])
		} else {
			resultat = append(resultat, mot)
		}
	}

	return resultat
}
```

La fonction reçoit une liste de mots, et renvoie une nouvelle liste de mots, transformée.

---

## ✅ Exercice 2 — 35 min

Dans votre projet `textfix` :

**2.1** Sur une feuille, refaites le tableau pas à pas pour le texte `bonjour (up) le monde`. Quel est le résultat ?

**2.2** Dans `regles.go`, écrivez la fonction `appliquerMarqueurs` qui gère le marqueur `(up)`.

**2.3** Pour la tester sans fichier, ajoutez temporairement dans `main`, tout au début :

```go
fmt.Println(appliquerMarqueurs(strings.Fields("il est (up) parti")))
```

Il faut importer `strings` dans `main.go` pendant ce test.

**2.4** Testez avec `bonjour (up) le monde`. Le résultat correspond-il à votre tableau ?

**2.5** Testez avec un marqueur à la fin : `il fait beau (up)`.

**Vérification** : `il est (up) parti` donne `[il EST parti]`, et votre tableau sur papier correspond à ce qu'affiche le programme.

---

# Partie 3 — `(low)`, `(cap)` et les lettres accentuées

## `(low)`

Même principe que `(up)`, avec `strings.ToLower` :

```go
} else if mot == "(low)" {
	dernier := len(resultat) - 1
	resultat[dernier] = strings.ToLower(resultat[dernier])
}
```

`else if` permet d'enchaîner plusieurs cas, comme en séance 2 avec les notes.

## `(cap)` : une majuscule au début

`(cap)` doit mettre la **première lettre** en majuscule, et le reste en minuscules :

| Avant | Après |
|---|---|
| `paris` | `Paris` |
| `PARIS` | `Paris` |
| `élise` | `Élise` |

## Le piège des accents

On pourrait croire qu'il suffit de prendre la première lettre du mot. Mais pour Go, un texte est une suite d'**octets**, et une lettre accentuée comme `é` occupe **deux** octets.

```go
len("élise")    // 6, et non 5 !
```

Si on coupe le mot après le premier octet, on coupe le `é` en deux, et on obtient un caractère illisible.

## La solution : les runes

Une **rune**, c'est une vraie lettre, qu'elle occupe un ou plusieurs octets. On transforme le mot en liste de runes :

```go
lettres := []rune("élise")
len(lettres)    // 5
```

On peut alors modifier la première lettre sans rien abîmer.

## La fonction `capitaliser`

```go
func capitaliser(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}
```

Ligne par ligne :

On met tout le mot en minuscules, puis on le transforme en liste de lettres.
On remplace la première lettre par sa version majuscule. `unicode.ToUpper` agit sur **une seule** lettre ; il faut importer `unicode`.
On recolle les lettres en un texte avec `string(lettres)`.

Le `if mot == ""` protège le cas d'un mot vide, qui n'a pas de première lettre.

---

## ✅ Exercice 3 — 25 min

**3.1** Ajoutez le marqueur `(low)` à `appliquerMarqueurs`.

**3.2** Écrivez la fonction `capitaliser` dans `regles.go`.

**3.3** Ajoutez le marqueur `(cap)`.

**3.4** Testez `paris (cap)`, `PARIS (cap)` et `élise (cap)`.

**3.5** Testez une phrase qui contient les trois marqueurs : `bonjour (up) TOUT (low) le monde (cap)`.

**Vérification** : vous obtenez `Paris`, `Paris`, `Élise`, et la phrase complète donne `BONJOUR tout le Monde`.

---

# Partie 4 — Ne jamais planter

## Le cas dangereux

Que se passe-t-il avec ce texte ?

```
(up) bonjour
```

Le premier mot lu est un marqueur. À ce moment, `resultat` est **vide**. `len(resultat) - 1` vaut `-1`, et `resultat[-1]` n'existe pas. Le programme s'arrête brutalement :

```
panic: runtime error: index out of range [-1]
```

Le mot `panic` signale un **plantage**. Un programme ne devrait jamais planter à cause du contenu d'un fichier : c'est toujours un bug à corriger.

## La protection

Avant de modifier le dernier mot, on vérifie qu'il en existe un :

```go
if len(resultat) == 0 {
	continue
}
```

`continue` veut dire : « passe directement au mot suivant ». Le marqueur est simplement ignoré.

## Les autres situations à prévoir

| Situation | Comportement voulu |
|---|---|
| Marqueur en début de ligne | Ignoré |
| Marqueur seul sur sa ligne | Ignoré, la ligne devient vide |
| Marqueur inconnu, comme `(truc)` | Laissé dans le texte, comme un mot normal |
| Ligne vide | Conservée |
| Espaces en trop | Réduits à un seul |

Pour le marqueur inconnu, il n'y a rien à faire : votre code l'ajoute déjà comme un mot normal, puisqu'il ne correspond à aucun des trois marqueurs.

## Regrouper la vérification

Les trois marqueurs ont besoin de la même vérification. Plutôt que de l'écrire trois fois, on la fait une seule fois pour les trois :

```go
if mot == "(up)" || mot == "(low)" || mot == "(cap)" {
	if len(resultat) == 0 {
		continue
	}
	dernier := len(resultat) - 1
	// puis on choisit la transformation selon le marqueur
}
```

`||` veut dire « ou ». La condition est vraie si le mot est l'un des trois marqueurs.

---

## ✅ Exercice 4 — 25 min

**4.1** Testez `(up) bonjour`. Notez le message d'erreur.

**4.2** Ajoutez la protection et réorganisez votre code : une seule vérification pour les trois marqueurs, grâce à `||`.

**4.3** Testez chaque situation du tableau de la partie 4.

**4.4** Imaginez un autre texte mal écrit qui pourrait faire planter votre programme. Testez-le.

**Vérification** : aucune des situations du tableau ne provoque de `panic`, et `un (truc) inconnu` reste inchangé.

---

# Partie 5 — Brancher le parser dans `textfix`

## Où l'appeler

La fonction `transformer`, écrite en séance 3, découpe déjà chaque ligne en mots. Il suffit d'ajouter **une ligne** pour appliquer les marqueurs à ces mots :

```go
for i, ligne := range lignes {
	ligne = strings.TrimRight(ligne, "\r")
	mots := strings.Fields(ligne)
	mots = appliquerMarqueurs(mots)
	lignes[i] = strings.Join(mots, " ")
}
```

---

## ✅ Exercice 5 — 15 min

**5.1** Retirez la ligne de test ajoutée dans `main` à l'exercice 2, ainsi que l'import de `strings` dans `main.go`.

**5.2** Appelez `appliquerMarqueurs` dans `transformer`.

**5.3** Créez un fichier `essai.txt` :

```
bonjour (up) tout le monde (cap)
CE TEXTE CRIE (low) un peu trop
élise (cap) est arrivée
(up) en début de ligne
un (truc) inconnu

des   espaces   en   trop
```

**5.4** Lancez `go run . essai.txt sortie.txt` et vérifiez le résultat ligne par ligne.

**Vérification** : `sortie.txt` contient exactement :

```
BONJOUR tout le Monde
CE TEXTE crie un peu trop
Élise est arrivée
en début de ligne
un (truc) inconnu

des espaces en trop
```

---

## 🎯 Bonus — pour ceux qui ont terminé

Deux nouveaux marqueurs convertissent un nombre écrit dans une autre base :

| Entrée | Sortie |
|---|---|
| `1E (hex)` | `30` |
| `10 (bin)` | `2` |

`(hex)` lit le mot précédent comme un nombre **hexadécimal** (base 16, avec les chiffres 0 à 9 puis A à F). `(bin)` le lit comme un nombre **binaire** (base 2, avec seulement 0 et 1).

L'outil à utiliser :

```go
valeur, err := strconv.ParseInt("1E", 16, 64)
texte := strconv.FormatInt(valeur, 10)
```

`ParseInt` lit le texte dans la base indiquée (16 ici). `FormatInt` le réécrit en base 10. Si le mot n'est pas un nombre valide, `err` n'est pas `nil` : dans ce cas, laissez le mot tel quel.

**Vérification** : `il y a 1E (hex) fichiers et 10 (bin) dossiers` donne `il y a 30 fichiers et 2 dossiers`, et `zz (hex)` laisse `zz` inchangé.

---

# Partie 6 — CampusLink : contrôler, compter, exporter

## Trois nouvelles commandes

En séance 3, CampusLink savait importer des salles. Il reçoit aujourd'hui trois commandes de plus :

| Commande | Ce qu'elle fait |
|---|---|
| `controle salles.csv` | Vérifie qu'aucun identifiant de salle n'apparaît deux fois |
| `stats salles.csv` | Affiche le nombre de salles, la capacité totale, et le nombre de salles par bâtiment |
| `export salles.csv propre.csv` | Écrit les salles valides dans un nouveau fichier CSV |

## Ne pas répéter le chargement

Les quatre commandes commencent toutes de la même façon : lire le fichier, le transformer en salles, signaler les lignes ignorées. Plutôt que d'écrire ce code quatre fois, on le range dans une fonction :

```go
func chargerSalles(chemin string) ([]Salle, error) {
	texte, err := lireFichier(chemin)
	if err != nil {
		return nil, err
	}

	salles, problemes := lireSalles(texte)
	for _, p := range problemes {
		fmt.Fprintln(os.Stderr, "Ligne ignorée :", p)
	}

	return salles, nil
}
```

`nil` à la place d'une liste veut dire « pas de liste ». Chaque commande commence alors par une seule ligne : `salles, err := chargerSalles(args[0])`.

C'est la règle à retenir : **quand on écrit deux fois le même code, on en fait une fonction.**

## Compter par catégorie : la map

Pour compter les salles de chaque bâtiment, il faut un compteur par bâtiment. On ne sait pas à l'avance combien il y a de bâtiments, ni leur nom.

Une **map** range des valeurs sous une **clé**. Imaginez un carnet où chaque page porte un nom : on ouvre la page « A », on lit ou on écrit le nombre qui s'y trouve.

```go
parBatiment := map[string]int{}
```

`map[string]int` veut dire : des clés qui sont du texte, des valeurs qui sont des nombres. Les accolades `{}` créent la map, vide au départ.

```go
parBatiment["A"] = parBatiment["A"] + 1
```

Si la clé `"A"` n'existe pas encore, `parBatiment["A"]` vaut **0**. La première salle du bâtiment A donne donc 0 + 1 = 1, la deuxième 1 + 1 = 2. Pas besoin de créer le compteur avant.

| Salle lue | Bâtiment | `parBatiment` après |
|---|---|---|
| S101 | A | A : 1 |
| S102 | A | A : 2 |
| S201 | B | A : 2, B : 1 |
| S202 | B | A : 2, B : 2 |

## L'ordre d'une map n'est pas fixe

On parcourt une map avec `range`, comme une liste. Mais Go donne les clés **dans un ordre différent à chaque lancement**. Pour un affichage stable, on récupère les clés dans une liste, puis on la trie :

```go
var batiments []string
for batiment := range parBatiment {
	batiments = append(batiments, batiment)
}
sort.Strings(batiments)
```

`sort.Strings` trie une liste de textes par ordre alphabétique. Il faut importer `sort`.

## Repérer un doublon

Une autre map, avec des valeurs `true` ou `false`, retient les identifiants déjà rencontrés :

```go
vus := map[string]bool{}

for _, s := range salles {
	if vus[s.ID] {
		// cet identifiant a déjà été vu : c'est un doublon
	}
	vus[s.ID] = true
}
```

Une clé absente vaut `false`. La première fois qu'on lit `S101`, `vus["S101"]` est faux ; on le passe à vrai. La deuxième fois, il est vrai : c'est un doublon.

## Un contrôle qui échoue renvoie une erreur

Si `controle` trouve un problème, il renvoie une erreur. Le programme s'arrête alors avec le code 1, comme pour un fichier introuvable. D'autres programmes pourront ainsi savoir si les données sont propres, sans lire les messages.

## Exporter

Pour écrire une salle dans une ligne CSV, on recolle ses champs avec des virgules. La capacité est un nombre : on la reconvertit en texte avec `strconv.Itoa`, l'inverse de `strconv.Atoi`.

```go
ligne := s.ID + "," + s.Nom + "," + s.Batiment + "," + strconv.Itoa(s.Capacite)
```

On rassemble les lignes, en-tête compris, avec `strings.Join`, puis on écrit le fichier avec `ecrireFichier`, déjà présent dans `fichiers.go`.

---

## ✅ Exercice 6 — CampusLink : `controle`, `stats` et `export` — 35 min

Revenez dans votre projet `campuslink`.

**6.1** Écrivez `chargerSalles` dans `main.go`, et utilisez-la dans `commandeImportSalles`. Le résumé final devient `N salle(s) importée(s).`

**6.2** Ajoutez la commande `stats` : nombre de salles, capacité totale, et nombre de salles par bâtiment, triés par ordre alphabétique.

**6.3** Dans `salles.go`, écrivez `trouverDoublons(salles []Salle) []string`, qui renvoie un message pour chaque identifiant déjà rencontré.

**6.4** Ajoutez la commande `controle`. Sans problème, elle affiche `Contrôle terminé : aucun problème trouvé.` Sinon, elle affiche chaque problème et renvoie une erreur.

**6.5** Ajoutez la commande `export`, qui attend deux noms de fichiers et refuse qu'ils soient identiques.

**6.6** Ajoutez les trois commandes au texte d'utilisation.

**6.7** Créez `donnees/salles-doublons.csv` :

```
id,nom,batiment,capacite
S101,Salle Turing,A,30
S102,Salle Lovelace,A,24
S101,Salle Turing bis,A,28
```

**6.8** Testez `stats` et `controle` sur `salles.csv`, `controle` sur `salles-doublons.csv`, et `export` sur `salles-erreurs.csv`.

**Vérification** : `go run . stats donnees/salles.csv` affiche :

```
Nombre de salles : 4
Capacité totale : 192 places
Salles par bâtiment :
  Bâtiment A : 2
  Bâtiment B : 2
```

`controle` signale le doublon `S101` et se termine avec le code 1, et le fichier exporté depuis `salles-erreurs.csv` ne contient que l'en-tête et les deux salles valides.

---

# À retenir

## Découper

`strings.Split(texte, "\n")` pour les lignes, `strings.Fields(ligne)` pour les mots.

## Le parser

On construit une nouvelle liste `resultat`.
Un mot normal y est ajouté. Un marqueur modifie son dernier élément, et n'est pas ajouté.
Le dernier élément est à l'index `len(resultat) - 1`.

## Les lettres

Une lettre accentuée occupe plusieurs octets.
`[]rune(mot)` pour manipuler les lettres une par une.

## Ne jamais planter

On vérifie que `resultat` n'est pas vide avant de modifier son dernier mot.
`continue` passe au mot suivant.
Un `panic` est toujours un bug.

## CampusLink

Quand on écrit deux fois le même code, on en fait une fonction : `chargerSalles`.
`map[string]int{}` crée une map ; une clé absente vaut 0, ou `false` pour une `map[string]bool`.
L'ordre d'une map change à chaque lancement : on trie ses clés avec `sort.Strings`.
`strconv.Itoa` convertit un nombre en texte, l'inverse de `strconv.Atoi`.

## Les erreurs qui coûtent le plus de temps

| Symptôme | Cause |
|---|---|
| `panic: runtime error: index out of range [-1]` | Marqueur en début de ligne, `resultat` encore vide |
| Le marqueur apparaît dans le résultat | Il est ajouté avec `append` au lieu de modifier le dernier mot |
| La liste `resultat` reste vide | `append(resultat, mot)` sans `resultat =` devant |
| `élise (cap)` donne un caractère bizarre | Le mot a été coupé par octets au lieu de runes |
| `undefined: unicode` | `import "unicode"` oublié dans `regles.go` |
| `"strings" imported and not used` dans `main.go` | Import laissé après avoir retiré la ligne de test |
| Toutes les lignes collées en une | Découpage direct en mots, sans passer par les lignes |
| `panic: assignment to entry in nil map` | Map déclarée avec `var parBatiment map[string]int`, sans `{}` : elle n'existe pas encore |
| Les bâtiments s'affichent dans un ordre différent à chaque lancement | Clés de la map parcourues sans être triées |
| `invalid operation: … (mismatched types string and int)` | Texte et nombre collés avec `+` : il faut `strconv.Itoa` |
| Chaque salle est signalée comme doublon | `vus[s.ID] = true` placé avant le test au lieu d'après |
| `controle` trouve un problème mais le code de sortie reste 0 | La commande affiche le problème sans renvoyer d'erreur |

---

# Prochaine séance

**Ranger, tester, documenter (4 h)** : on vérifie que chaque fichier a un rôle clair, on teste le programme avec une série de fichiers, et on rédige sa documentation. CampusLink apprendra aussi à importer les équipements.
