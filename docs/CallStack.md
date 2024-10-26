# CallStack

## gui 

- Aendre tekst paa submit knap

- handle excrpts i gui
- GUI, kan ikke handle at aabne chid Ignored
- Naar man afbryder et run skal man rulle tilbage
    - Ogsaa update af commongrps
        - Gaar denne igennem for det andet.
- Naar man laver en error skal ens fields blive repopulated.

## udvidelser

- Lœs udtrœk fra PDF og ikke kun csv filer

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

- UpdateBudgetReqs skal ikke ligge cli
    - Det er en faelles funktion som bruges af alt UI

