# CallStack

- hvorfor er den saa langsom efter man har angivet mdr
    - den henter config, men det kan den gore til at starte med i en seperart traad
- naar man opdatere cellen, kan man gore det som "= x1 +x2 ...", saa den er nem at rette i
- haandter mellemrum i input fil navnet
    - skal kunne have punktum i fil navn
- graceful shut down, hvis der sker en fejl
    - lige nu dor den bare instant
- nogen gang kommer der en stor forskel paa faktisk og udregnet konto balance. som bliver rettet op naar man udfylder for naeste mdr

## gui 

- command pattern
    - https://refactoring.guru/design-patterns/command

- Aendre tekst paa submit knap

- handle excrpts i gui
- GUI, kan ikke handle at aabne chid Ignored
- Naar man afbryder et run skal man rulle tilbage
    - Ogsaa update af commongrps
        - Gaar denne igennem for det andet.
- Naar man laver en error skal ens fields blive repopulated.


## udvidelser

- Angiv bank i build script
    - Det fortaeller hvilken parser funktion der skal anvendes, der in-turn reagere paa fil typen
    - Er nodvendigt i for formattering

- Lœs udtrœk fra PDF og ikke kun csv filer
    - behavior pattern

- format haandtering af csv filen
    - Tror den fucker op hvis man har odt format vs text csv
    - Den pt kan den kun klare text csv

- Brug multithreading til at match udtraek med excrptGrp

- Man skal ogsaa kunne traekke ud for flere aar af gangen.
    - Saa man skal baade angive aar og mdr

- Hvor skal / kan man se logs

- Validation
    - Skal kunne bruges til "debug" budgettet
    - Maaske mest interessent for de faste overfoerlser til faelles kontoen
    - Kan man have mere end bare actual account balance

- Håndtering forskellig formattering af udtrœkkene
  - 1.000.000 vs 1000000
  - 1000,00 vs 1000.00
  - og dets permutationer

- Kunne man automatisere det hele hvis man fik en bank-fuldmagt af brugeren

## Refactor


- Hent config paa en smartere maade, saa man ikke henter >6000 raekker
- En bedre maade at tjekke paa om man har en header i config
    - Lige nu tjekker vi bare paa baggrundsfarven
- Der er mange init funktioner man kan kalde fra init metoden i stedet for main
- burde config.GetConfig() returnere en error?
- UpdateBudgetReqs skal ikke ligge cli
    - Det er en faelles funktion som bruges af alt UI
- Default fixed expense til false
- Det er kun cli.go der holder build variable
    - Men disse er ikke synlige for GUI'en fordi alle directories haandteres som en seperart build unit af go. Ligegyldig om det er i samme package
