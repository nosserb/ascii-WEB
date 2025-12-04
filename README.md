<a href="https://ibb.co/4whyFz63"><img src="https://i.ibb.co/vC05cp7R/Setting-up-Git-1.png" alt="Setting-up-Git-1" border="0"></a>
<a href="https://ibb.co/jkzX9k2t"><img src="https://i.ibb.co/DfY2Pfyd/Votre-texte-de-paragraphe-43.png" alt="Votre-texte-de-paragraphe-43" border="0"></a>

<a href="https://ibb.co/60fs6gcK"><img src="https://i.ibb.co/ch7TzDKW/Votre-texte-de-paragraphe-45.png" alt="Votre-texte-de-paragraphe-45" border="0"></a>

```
                           Project Overview
                           Module 1 : Banner
                           Module 2 : Render
                           Module 3 : Main
                             Installation
                                 Usage
                             Unit Testing
```

---

<a href="https://ibb.co/H00JNRZ"><img src="https://i.ibb.co/Z33ZHv9/Votre-texte-de-paragraphe-46.png" alt="Votre-texte-de-paragraphe-46" border="0"></a>

**Projet de génération d'art ASCII à partir de fichiers de bannières.**

**Objectif :** Convertir du texte standard en représentations ASCII art avec différentes polices (standard, shadow, thinkertoy).

**Technologies utilisées :**
- Go 1.22.2
- Bibliothèque standard uniquement

---

<a href="https://ibb.co/vx83Pg9Z"><img src="https://i.ibb.co/nsS7cZpz/Votre-texte-de-paragraphe-35.png" alt="Votre-texte-de-paragraphe-35" border="0"></a>

***Responsabilités***

- Lire les fichiers de bannières : `standard`, `shadow`, `thinkertoy`
- Parser la bannière selon le format :
  - 8 lignes + 1 ligne vide par caractère
  - Caractères ASCII de 32 à 126 (95 caractères au total)
- Fonction principale à implémenter :

```go
LoadBanner(path string) (map[rune][]string, error)
```

***Gestion des erreurs***

- Fichier introuvable
- Bannière mal formatée
- Format incorrect (nombre de lignes invalide)

## Tests unitaires requis

***Vérifier le nombre de glyphes***
*Doit retourner 95 caractères*
```
TestBannerGlyphCount()
```

***Vérifier la hauteur de chaque glyph***
*Chaque glyph doit avoir 8 lignes*
```
TestBannerLineHeight()
```

***Livrables***

- `banner.go`
- `banner_test.go`

***Avantages***

- *Développement indépendant des modules 2 et 3*  
- *Travail technique bien défini et équilibré*
- *Peut être testé unitairement sans dépendances*

---

<a href="https://ibb.co/cSrqwTcK"><img src="https://i.ibb.co/G45m9Pvf/Votre-texte-de-paragraphe-36.png" alt="Votre-texte-de-paragraphe-36" border="0"></a>

***Responsabilités***

- Transformer un texte en ASCII-art
- Fonction principale :

```go
PrintASCII(text string, font map[rune][]string) (string, error)
```

***Gestion des cas spéciaux***

- Caractères spéciaux
- Espaces
- Retours à la ligne (`\n`)
- Lignes vides → imprimer 8 lignes vides
- Caractères non supportés → retourner une erreur

***Construction du rendu***

Le rendu se fait **ligne par ligne** :
1. Parcourir chaque ligne de texte
2. Pour chaque ligne de hauteur (0 à 7) :
   - Concaténer la ligne correspondante de chaque caractère
3. Assembler le résultat final

***Tests unitaires requis***

**Rendu d'un seul caractère**
```
TestPrintASCIISingleCharacter()
```
**Rendu d'un mot complet**
```
TestPrintASCIIWord()
```
**Gestion des espaces**           
```
TestPrintASCIIWithSpace()
```
**Gestion des \n**
```
TestPrintASCIINewline()
```
**Gestion des lignes vides**           
```
TestPrintASCIIEmptyLine() 
```
**Caractères non supportés**     
```
TestPrintASCIIUnsupportedCharacter() 
```

***Livrables***

- `render.go` (implémenté)
- `render_test.go` (14 tests)

***Avantages***
```
                Très indépendant → développement immédiat  
                        Logique claire et testable  
                Peu de dépendances avec les autres modules
```

---

<a href="https://ibb.co/pvMLMkkZ"><img src="https://i.ibb.co/jvQrQjjG/Votre-texte-de-paragraphe-44.png" alt="Votre-texte-de-paragraphe-44" border="0"></a>


***Gestion des arguments de la ligne de commande :***
```
go run . "text" (utilise la police par défaut : standard)
go run . "text" standard
go run . "text" shadow
go run . "text" thinkertoy
```

***Intégration des modules :***

- Appeler LoadBanner() ou LoadFont()
- Appeler PrintASCII()
- Gérer le flux de données entre modules



Validation des entrées

---

<a href="https://ibb.co/MxD0J141"><img src="https://i.ibb.co/wFNm2QXQ/Votre-texte-de-paragraphe-37.png" alt="Votre-texte-de-paragraphe-37" border="0"></a>

***Cloner le projet***

```bash
git clone <repository-url>
cd ascii-art
```

***Vérifier l'installation de Go***

```bash
go version
```
***result :***
```
go version go1.22.2 ou supérieur
```

***Structure des fichiers***
> **! la structure est temporaire !**

```
ascii-art/
├──banner
|   ├── banner_test.go
|   └── banner.go
├── go.mod
├── render.go
├── render_test.go
├── standard           
├── shadow             
└── thinkertoy         
```

---

<a href="https://ibb.co/rKn9sk03"><img src="https://i.ibb.co/dsnZgBMt/Votre-texte-de-paragraphe-38.png" alt="Votre-texte-de-paragraphe-38" border="0"></a>

***Commande de base***

```bash
go run render.go <fontFile> <text>
```

***Exemples***

```bash
go run render.go <police> "Hello"

```

***police disponible***
```
shadow
_| _| _|  

standard
(_) (_) (_) 

tinkertoy
O O O 
```

### Exemples avec cas spéciaux


 ***Texte avec espaces***
```
go run render.go standard "Hello World"
```
***Texte avec retours à la ligne***
```
go run render.go shadow "Line1\nLine2"
```

***Texte multiligne***
```
go run render.go thinkertoy "First\n\nThird"
```

***Sortie attendue***

```bash
go run render.go standard "Hi"
```
**result :** 
```
 _    _  _ 
| |  | |(_)
| |__| | _ 
|  __  || |
| |  | || |
|_|  |_||_|
```

---

<a href="https://ibb.co/WpRTTx5v"><img src="https://i.ibb.co/cXsmmg3h/Votre-texte-de-paragraphe-39.png" alt="Votre-texte-de-paragraphe-39" border="0"></a>

***Lancer tous les tests***

```bash
go test -v
```

***Lancer des tests spécifiques***

```
# Tests du module de rendu uniquement
go test -v -run TestPrintASCII

# Test d'un cas particulier
go test -v -run TestPrintASCIISingleCharacter

# Tests avec coverage
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

***Résultats attendus***

```bash
=== RUN   TestPrintASCIISingleCharacter
--- PASS: TestPrintASCIISingleCharacter (0.00s)
=== RUN   TestPrintASCIIWord
--- PASS: TestPrintASCIIWord (0.00s)
=== RUN   TestPrintASCIIWithSpace
--- PASS: TestPrintASCIIWithSpace (0.00s)
...
PASS
ok      ascii-art    0.123s
```

---

**Concepts clés du projet**

***Architecture modulaire***

>- **Module 1** : Gestion des fichiers de polices
>- **Module 2** : Logique de rendu ASCII
>- **Module 3** : 

***Principes de développement***

>- **Indépendance** : Chaque module peut être développé séparément
>- **Testabilité** : Tests unitaires pour chaque fonctionnalité
>- **Simplicité** : Utilisation de la bibliothèque standard Go uniquement
>- **Clarté** : Code documenté et structure lisible

***Points techniques importants***

>- **Format de police** : 8 lignes + 1 vide par caractère
>- **Plage ASCII** : Caractères 32 à 126 (95 caractères)
>- **Gestion d'erreurs** : Validation stricte des entrées
>- **Rendu ligne par ligne** : Construction progressive du résultat

---

***Guylann Bresson - Kaelig Camesella - Valentin Bosson***  
***18/11/2025***  
***ASCII Art - Zone01 Normandie***