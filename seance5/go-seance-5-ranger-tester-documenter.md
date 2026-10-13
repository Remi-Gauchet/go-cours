# Séance 5 — Ranger, tester, documenter

**Application CLI & Parsing en Go — Bachelor 1 — Ycode A2627_4658**
**Séance 5 · 4 h**
**Prérequis : séances 1 à 4**

`textfix` fonctionne. Pour en faire un vrai projet, il reste trois choses à faire :

**Ranger** — que chaque fichier ait un rôle clair.
**Tester** — prouver que le programme fait bien ce qu'il doit faire.
**Documenter** — qu'une personne qui ne vous connaît pas puisse s'en servir.

Ce sont ces trois choses qui font la différence entre un exercice et un logiciel qu'on peut confier à quelqu'un.

En dernière partie, vous faites la même chose sur CampusLink, et vous terminez ses commandes avec l'import des équipements.

---

## Ce que vous saurez faire en fin de séance

Expliquer le rôle de chaque fichier de votre projet.
Commenter votre code.
Tester un programme avec des fichiers d'exemple.
Rédiger un README.
Relire le projet d'un camarade et lui faire des remarques utiles.
Livrer CampusLink : code rangé, fichiers de test valides et invalides, README des commandes.

---

## Déroulé

| Partie | Contenu | Durée |
|---|---|---|
| 1 | Un fichier, un rôle | 15 min |
| — | Exercice 1 | 25 min |
| 2 | Tester avec des fichiers | 15 min |
| — | Exercice 2 | 40 min |
| 3 | Le README | 15 min |
| — | Exercice 3 | 30 min |
| 4 | Relire le projet d'un camarade | 10 min |
| — | Exercice 4 | 20 min |
| — | Exercice 5 — finitions et bonus | 15 min |
| 6 | CampusLink : ranger, tester, documenter | 15 min |
| — | Exercice 6 — CampusLink | 40 min |

---

# Partie 1 — Un fichier, un rôle

## Pourquoi ranger

Un programme qu'on vient d'écrire, on le comprend. Le même programme relu trois semaines plus tard, ou par quelqu'un d'autre, c'est une autre histoire.

Un code bien rangé permet de savoir **immédiatement** où chercher. Un bug dans la lecture du fichier ? C'est dans `fichiers.go`. Une règle à ajouter ? C'est dans `regles.go`.

C'est ce qu'on appelle la **séparation des responsabilités** : chaque partie du programme s'occupe d'une seule chose.

## Le rôle de chacun de vos fichiers

| Fichier | Son rôle | Ce qu'il ne fait jamais |
|---|---|---|
| `main.go` | Vérifier les arguments, enchaîner les étapes, afficher les messages | Transformer le texte |
| `fichiers.go` | Lire et écrire les fichiers | Afficher des messages, arrêter le programme |
| `regles.go` | Transformer le texte | Lire ou écrire un fichier, afficher, arrêter le programme |

## Un test simple

Ouvrez `regles.go` et regardez ses imports. On doit y trouver `strings` et `unicode`, mais **jamais** `os` ni `fmt`.

Si `regles.go` a besoin de `os` ou de `fmt`, c'est qu'il fait quelque chose qui n'est pas son rôle : lire un fichier, ou afficher un message.

## Les commentaires

Un **commentaire** est une ligne que Go ignore complètement. Il sert à expliquer le code aux humains. Il commence par `//` :

```go
// capitaliser met une majuscule au début du mot et le reste en minuscules.
func capitaliser(mot string) string {
```

**Un commentaire au-dessus de chaque fonction**, qui dit **ce qu'elle fait**, en une phrase. Par convention, il commence par le nom de la fonction.

Un bon commentaire explique le **pourquoi** ou le **quoi**, pas ce qui se lit déjà dans le code :

| Inutile | Utile |
|---|---|
| `// on ajoute 1 à i` | `// on saute la première ligne, qui contient le titre` |
| `// boucle` | `// transformer applique les règles ligne par ligne.` |

---

## ✅ Exercice 1 — 25 min

**1.1** Ouvrez chacun de vos trois fichiers et vérifiez qu'il ne fait que ce qui est indiqué dans le tableau de la partie 1.

**1.2** Vérifiez que `regles.go` n'importe ni `os` ni `fmt`.

**1.3** Supprimez tout le code de test qui traîne : lignes `fmt.Println` ajoutées pour vérifier, fonctions qui ne servent plus.

**1.4** Ajoutez un commentaire au-dessus de **chaque** fonction de `fichiers.go` et de `regles.go`.

**1.5** Relancez le programme sur `essai.txt` : le résultat doit être exactement le même qu'en séance 4.

**Vérification** : chaque fonction a un commentaire, `regles.go` n'importe ni `os` ni `fmt`, et le programme fonctionne toujours.

---

# Partie 2 — Tester avec des fichiers

## Pourquoi tester

Quand on modifie un programme, on risque de casser quelque chose qui marchait. Un **test** permet de le vérifier en quelques secondes.

## Le principe

Pour chaque règle, on prépare **deux** fichiers :

**Le fichier d'entrée** — un texte qui utilise la règle.
**Le fichier attendu** — le résultat que le programme **devrait** produire, écrit à la main.

On lance le programme sur le fichier d'entrée, puis on compare sa sortie au fichier attendu. S'ils sont identiques, le test est réussi.

## Où les ranger

Dans un dossier `tests`, avec des noms qui vont par paires :

```
textfix/
├── main.go
├── fichiers.go
├── regles.go
└── tests/
    ├── 01-majuscules.txt
    ├── 01-majuscules.attendu.txt
    ├── 02-minuscules.txt
    ├── 02-minuscules.attendu.txt
    ├── 03-capitales.txt
    ├── 03-capitales.attendu.txt
    ├── 04-limites.txt
    └── 04-limites.attendu.txt
```

**Un fichier par règle.** Quand un test échoue, on sait immédiatement quelle règle est en cause.

## Le fichier attendu s'écrit à la main

C'est la règle la plus importante. Si vous copiez la sortie du programme dans le fichier attendu, le test ne vérifie plus rien : il dit seulement que le programme fait ce qu'il fait.

**Vous écrivez ce qui devrait sortir, en réfléchissant, avant de lancer le programme.**

## Lancer un test

```bash
go run . tests/01-majuscules.txt sortie.txt
```

Le `/` sépare le nom du dossier et le nom du fichier. Sous Windows, `\` fonctionne aussi.

## Comparer deux fichiers

**Dans VS Code** — le plus simple. Dans l'explorateur de gauche, cliquez sur `sortie.txt`, puis maintenez Ctrl (Cmd sur Mac) et cliquez sur le fichier attendu. Faites un clic droit, et choisissez l'option de comparaison (« Compare Selected » si VS Code est en anglais).

Les deux fichiers s'affichent côte à côte, et les différences sont surlignées.

**Dans le terminal** :

```bash
diff sortie.txt tests/01-majuscules.attendu.txt        # Mac, Linux, Git Bash
```

```powershell
fc.exe sortie.txt tests\01-majuscules.attendu.txt      # PowerShell
```

Si les fichiers sont identiques, `diff` n'affiche rien, et `fc.exe` affiche « aucune différence trouvée ».

---

## ✅ Exercice 2 — 40 min

**2.1** Créez le dossier `tests`.

**2.2** Créez les quatre fichiers d'entrée :

`01-majuscules.txt` : des phrases avec `(up)`.
`02-minuscules.txt` : des phrases avec `(low)`.
`03-capitales.txt` : des phrases avec `(cap)`, dont au moins un mot accentué.
`04-limites.txt` : un marqueur en début de ligne, un marqueur inconnu, une ligne vide, des espaces en trop.

**2.3** Pour chacun, écrivez **à la main** le fichier `.attendu.txt` correspondant.

**2.4** Lancez le programme sur chaque fichier d'entrée et comparez la sortie au fichier attendu.

**2.5** Pour chaque différence, cherchez si l'erreur vient du programme ou de votre fichier attendu, et corrigez le bon côté.

**2.6** Cassez volontairement le programme : dans `regles.go`, remplacez `strings.ToUpper` par `strings.ToLower` pour le marqueur `(up)`. Relancez les quatre tests. Lesquels échouent ? Remettez le code d'origine.

**Vérification** : les quatre tests sont identiques à leur fichier attendu, et le code cassé volontairement fait échouer le test des majuscules.

---

# Partie 3 — Le README

## Ce que c'est

Le **README** (« lisez-moi ») est le premier fichier qu'on ouvre en découvrant un projet. Il explique ce que fait le programme, comment l'installer, et comment s'en servir.

Il s'écrit dans un fichier `README.md`, à la racine du projet, en **Markdown** : du texte simple, avec quelques symboles pour la mise en forme.

## Le Markdown en 30 secondes

| On écrit | On obtient |
|---|---|
| `# Titre` | Un grand titre |
| `## Sous-titre` | Un titre plus petit |
| `**gras**` | du texte en **gras** |
| `` `code` `` | du `code` |
| `- élément` | un élément de liste |

Pour un bloc de commandes, on l'encadre de trois accents graves :

````
```bash
go run . entree.txt sortie.txt
```
````

Dans VS Code, `Ctrl+Shift+V` affiche l'aperçu du README mis en forme.

## Le modèle à compléter

````markdown
# textfix

(Deux ou trois phrases : à quoi sert ce programme ?)

## Installation

(Ce qu'il faut avoir installé, et la commande pour préparer le projet.)

## Utilisation

```bash
(la commande pour lancer le programme)
```

(Un exemple : un texte d'entrée, et le résultat obtenu.)

## Les règles

| Marqueur | Effet | Exemple | Résultat |
|---|---|---|---|
| `(up)` | | | |
| `(low)` | | | |
| `(cap)` | | | |

## Les messages d'erreur

(Dans quels cas le programme s'arrête, et ce qu'il affiche.)

## Les tests

(Comment lancer un test et comparer le résultat.)

## Organisation du code

(Le rôle de chaque fichier.)

## Limites

(Ce que le programme ne sait pas faire.)

## Auteur

(Votre prénom et votre nom.)
````

## Le bon test

Un README est réussi quand **quelqu'un qui ne vous connaît pas** peut lancer votre programme en le lisant, **sans vous poser une seule question**.

---

## ✅ Exercice 3 — 30 min

**3.1** Créez `README.md` à la racine du projet et copiez-y le modèle.

**3.2** Complétez toutes les sections. Pour les règles, donnez un exemple pour chaque marqueur.

**3.3** Dans la section Limites, notez au moins une chose que votre programme ne gère pas. Un indice : que se passe-t-il avec `bonjour(up)`, sans espace avant le marqueur ?

**3.4** Affichez l'aperçu avec `Ctrl+Shift+V` et vérifiez la mise en forme.

**3.5** Donnez votre README à votre voisin. Il doit réussir à lancer votre programme **en le lisant seulement**. Notez ce qui lui a manqué, et complétez.

**Vérification** : votre voisin a lancé votre programme sans vous poser de question.

---

# Partie 4 — Relire le projet d'un camarade

## Pourquoi

Les développeurs relisent le code de leurs collègues tous les jours. Un regard extérieur repère ce qu'on ne voit plus soi-même.

## Une bonne remarque

Elle est **précise** et **bienveillante**. Elle dit où, et propose une amélioration.

| Pas utile | Utile |
|---|---|
| « C'est pas clair » | « Dans `regles.go`, la fonction `capitaliser` n'a pas de commentaire » |
| « Ton README est nul » | « Le README ne dit pas qu'il faut lancer `go run .` depuis le dossier du projet » |

---

## ✅ Exercice 4 — 20 min

Échangez votre ordinateur avec un camarade, et vérifiez son projet avec cette grille :

| Point vérifié | Oui / Non |
|---|---|
| Le programme se lance sans erreur avec `go run . entree.txt sortie.txt` | |
| Sans argument, il affiche un message clair au lieu de planter | |
| Avec un fichier absent, il affiche un message clair | |
| Il refuse d'écrire dans son propre fichier d'entrée | |
| `regles.go` n'importe ni `os` ni `fmt` | |
| Chaque fonction a un commentaire | |
| Les quatre tests correspondent à leur fichier attendu | |
| Le README permet de lancer le programme sans aide | |
| Au moins une limite est notée dans le README | |

**Vérification** : chaque « Non » est accompagné d'une remarque précise, transmise à votre camarade.

---

## ✅ Exercice 5 — Finitions et bonus — 15 min

**5.1** Corrigez votre projet à partir des remarques reçues.

**5.2** Les bonus ci-dessous sont facultatifs : passez directement à la partie 6 si votre projet n'est pas terminé, et revenez-y à la fin si vous avez le temps.

## 🎯 Bonus A — Les nombres

Si ce n'est pas déjà fait, ajoutez `(hex)` et `(bin)`, présentés en fin de séance 4. Ajoutez une paire de tests `05-nombres`.

## 🎯 Bonus B — La ponctuation

Un signe de ponctuation isolé doit être collé au mot précédent :

| Entrée | Sortie |
|---|---|
| `bonjour , le monde !` | `bonjour, le monde!` |
| `attends ... quoi !?` | `attends... quoi!?` |

Écrivez une fonction `collerPonctuation(mots []string) []string`, sur le même principe que `appliquerMarqueurs` : on construit une nouvelle liste, et quand un mot n'est composé que de ponctuation, on le colle à la fin du dernier mot au lieu de l'ajouter.

Pour savoir si un mot n'est fait que de ponctuation, cette astuce fonctionne :

```go
strings.Trim(mot, ".,!?:;") == ""
```

`strings.Trim` retire les caractères indiqués au début et à la fin. S'il ne reste rien, c'est que le mot n'était que de la ponctuation.

Appelez `collerPonctuation` dans `transformer`, **après** `appliquerMarqueurs`. Ajoutez une paire de tests `06-ponctuation`.

**Vérification** : tous vos tests passent encore, y compris les anciens.

---

# Partie 6 — CampusLink : ranger, tester, documenter

Vous appliquez maintenant à CampusLink les trois gestes de la séance, et vous ajoutez la dernière commande : l'import des équipements.

## Un fichier, un rôle

À la fin de la séance 4, `main.go` contient le choix de la commande **et** le code de chaque commande. Il est devenu long. On le découpe :

| Fichier | Son rôle | Ce qu'il ne fait jamais |
|---|---|---|
| `main.go` | Lire la commande, lancer la bonne fonction, afficher l'erreur éventuelle | Lire un fichier, calculer |
| `commandes.go` | Le code de chaque commande : charger, calculer, afficher | Arrêter le programme |
| `fichiers.go` | Lire et écrire les fichiers | Afficher, arrêter le programme |
| `salles.go` | Transformer du texte CSV en salles, trouver les doublons | Lire un fichier, afficher |
| `equipements.go` | Transformer du texte CSV en équipements | Lire un fichier, afficher |

Le test simple de la partie 1 s'applique toujours : `salles.go` et `equipements.go` n'importent jamais `os`.

## Les équipements

Un équipement est installé dans une salle. Le fichier CSV a lui aussi quatre colonnes :

```
id,nom,salle,quantite
E01,Vidéoprojecteur,S101,1
E02,Ordinateur,S102,24
E03,Micro,S201,2
```

La structure `Equipement` et la fonction `lireEquipement` suivent **exactement** le modèle de `Salle` et de `lireSalle`. Seule une vérification est nouvelle : la salle indiquée doit exister.

Pour le savoir, on range les identifiants des salles dans une map, comme pour les doublons :

```go
sallesConnues := map[string]bool{}
for _, s := range salles {
	sallesConnues[s.ID] = true
}
```

Puis, pour chaque équipement :

```go
if !sallesConnues[equipement.Salle] {
	// la salle n'existe pas : la ligne est ignorée et signalée
}
```

Le `!` veut dire « non » : la condition est vraie quand la salle n'est **pas** connue.

La commande attend donc deux fichiers, les salles puis les équipements :

```bash
go run . import-equipements donnees/salles.csv donnees/equipements.csv
```

## Des fichiers de test valides et invalides

Pour chaque type de données, on prépare deux fichiers :

**Un fichier valide** — toutes les lignes sont correctes. Il vérifie que le programme fait bien son travail.

**Un fichier invalide** — chaque ligne contient une erreur différente, sauf une ou deux lignes correctes. Il vérifie que le programme signale chaque erreur, sans planter, et garde les lignes correctes.

```
tests/
├── salles-valides.csv
├── salles-invalides.csv
├── equipements-valides.csv
├── equipements-invalides.csv
└── RESULTATS-ATTENDUS.md
```

`RESULTATS-ATTENDUS.md` liste, pour chaque commande et chaque fichier, ce que le programme doit afficher. Comme pour `textfix`, ces résultats s'écrivent **à la main, avant** de lancer le programme.

## Comparer un affichage

Pour une commande qui affiche son résultat, on lance la commande et on compare l'écran au résultat attendu.

Sous Mac ou Git Bash, le symbole `>` envoie l'affichage dans un fichier, qu'on compare ensuite avec `diff` :

```bash
go run . stats tests/salles-valides.csv > sortie.txt
```

Seule la sortie normale part dans le fichier : les messages d'erreur restent à l'écran. Sous PowerShell, comparez directement à l'écran.

## Le README des commandes

Le README de CampusLink présente chaque commande avec : ce qu'elle fait, les fichiers qu'elle attend, et un exemple de résultat. C'est la partie que l'utilisateur lira en premier.

---

## ✅ Exercice 6 — CampusLink : équipements, tests et README — 40 min

**6.1** Créez `commandes.go` et déplacez-y toutes les fonctions `commande…`, ainsi que `chargerSalles` et `afficherSalle`. `main.go` ne garde que `usage` et `main`.

**6.2** Créez `equipements.go` avec la structure `Equipement`, et les fonctions `lireEquipement(ligne string) (Equipement, error)` et `lireEquipements(texte string, salles []Salle) ([]Equipement, []string)`. La quantité doit être un nombre positif, et la salle doit exister.

**6.3** Ajoutez la commande `import-equipements <fichier_salles> <fichier_equipements>` et mettez à jour l'utilisation.

**6.4** Créez le dossier `tests` avec les quatre fichiers.

`tests/salles-valides.csv` :

```
id,nom,batiment,capacite
S101,Salle Turing,A,30
S102,Salle Lovelace,A,24
S201,Amphi Hopper,B,120
S202,Salle Hamilton,B,18
```

`tests/salles-invalides.csv` :

```
id,nom,batiment,capacite
S101,Salle Turing,A,30
S102,Salle Lovelace,A
S103,Salle Knuth,A,trente
,Salle sans nom,A,20
S104,Salle vide,A,0
S201,Amphi Hopper,B,120
S201,Amphi Hopper bis,B,100
```

`tests/equipements-valides.csv` :

```
id,nom,salle,quantite
E01,Vidéoprojecteur,S101,1
E02,Ordinateur,S102,24
E03,Micro,S201,2
```

`tests/equipements-invalides.csv` :

```
id,nom,salle,quantite
E01,Vidéoprojecteur,S101,1
E02,Ordinateur,S999,24
E03,Micro,S201,deux
E04,Tableau,S102
```

**6.5** Écrivez à la main `tests/RESULTATS-ATTENDUS.md`, avec le résultat attendu de ces commandes :

| Commande | Fichier(s) |
|---|---|
| `import-salles` | `salles-valides.csv`, puis `salles-invalides.csv` |
| `import-equipements` | `salles-valides.csv` avec `equipements-valides.csv`, puis avec `equipements-invalides.csv` |
| `controle` | `salles-valides.csv`, puis `salles-invalides.csv` |
| `stats` | `salles-valides.csv` |
| `export` | `salles-invalides.csv` |

**6.6** Lancez chaque commande et comparez au résultat attendu. Corrigez le bon côté à chaque différence.

**6.7** Écrivez le `README.md` de CampusLink : description, installation, puis une section **Commandes** avec un exemple pour chacune des six commandes, et les limites connues.

**Vérification** : `salles.go` et `equipements.go` n'importent pas `os`, `go run . import-equipements tests/salles-valides.csv tests/equipements-invalides.csv` signale trois lignes ignorées et importe un seul équipement, et chaque résultat obtenu correspond à `RESULTATS-ATTENDUS.md`.

---

# À retenir

## Ranger

Un fichier, un rôle.
`regles.go` transforme du texte : il n'importe ni `os` ni `fmt`.
Un commentaire `//` au-dessus de chaque fonction, qui dit ce qu'elle fait.

## Tester

Un fichier d'entrée et un fichier attendu par règle, dans un dossier `tests`.
Le fichier attendu s'écrit à la main, avant de lancer le programme.
On relance tous les tests après chaque modification.

## Documenter

Le README explique à quoi sert le programme, comment l'installer, comment l'utiliser.
Un exemple pour chaque règle.
Les limites, honnêtement.
Il est réussi quand quelqu'un s'en sert sans vous poser de question.

## CampusLink

`main.go` choisit la commande, `commandes.go` les exécute, `salles.go` et `equipements.go` transforment les données sans importer `os`.
Un fichier de test valide et un fichier invalide par type de données.
Une map d'identifiants permet de vérifier qu'une salle existe.
Le README présente chaque commande avec un exemple.

## Les erreurs qui coûtent le plus de temps

| Symptôme | Cause |
|---|---|
| `"fmt" imported and not used` après le rangement | Un `Println` de test supprimé, mais l'import est resté |
| Le test échoue alors que les textes semblent identiques | Un espace en trop ou une ligne vide en plus à la fin du fichier attendu |
| Tous les tests échouent sous Windows | Fichiers attendus enregistrés en `CRLF` : cliquez sur `CRLF` en bas à droite de VS Code et choisissez `LF` |
| `Erreur : impossible de lire le fichier tests/01-majuscules.txt` | Terminal lancé depuis un autre dossier que celui du projet |
| Le test passe toujours, même avec le code cassé | Fichier attendu copié depuis la sortie du programme |
| Le README s'affiche sans mise en forme | Fichier nommé `README.txt` au lieu de `README.md` |
| `fc` affiche une erreur incompréhensible | Sous PowerShell, il faut écrire `fc.exe` |
| `undefined: commandeStats` après le découpage | `commandes.go` ne commence pas par `package main` |
| `"strconv" imported and not used` dans `main.go` | Import resté après le déplacement des commandes dans `commandes.go` |
| Tous les équipements sont rejetés : « la salle … n'existe pas » | La map `sallesConnues` n'est jamais remplie, ou le fichier de salles passé est le mauvais |
| `la commande import-equipements attend deux noms de fichiers` | Un seul fichier donné : la commande attend les salles **puis** les équipements |

---

# Et maintenant

Vos deux projets sont complets : ils lisent des fichiers, appliquent des règles, gèrent les erreurs, et sont rangés, testés et documentés. CampusLink contient tout ce que demande le livrable : le code de la CLI, des fichiers de test valides et invalides, et un README avec l'installation, les commandes et des exemples.

Pour chaque nouvelle commande que vous voudriez ajouter, la méthode reste la même : une fonction dans `commandes.go`, un fichier de test valide et un invalide, une section dans le README.
