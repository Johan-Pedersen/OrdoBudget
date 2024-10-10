# CallStack

- Ryd op efter gammel config metode
    - accouting.createEntry
    - RM data klasse
    - rm local config filer

- Group Ignored har bug
    - for aug, assginer den ikke automatisk til madkonto


- Fix UpdateCommonGrps
    - Skal heller ikke hedde CommonGrp, skal bare vaere en fixedExpense

- **REFACTOR** -> se uber-go/guide og "100 Go mistakes"
- 100 Go mistakes
    - interfaces
        - Er der et sted man ville m. fordel bruge interfaces?
    - error management
    - testing-

- Lav Groups til en liste af pointers
    - (bidirectional assosianion using pointers)
    - Lav Entry til at holde en pointer til en Group
    - Lav Group til at holde en liste af Entry pointers

    - Hvad faar man ud af at bruge pointers
        - Ingen risiko for et infinite deep copy loop
        - Objekterne i sig selv er mindre, hvilket er smart naar go har det med at kopiere objekter hele tiden saa det er nok smart

- pakke til accounting requests

- rename package requests -> request
    - ental

- Der en del funktioner der hedder det samme/naesten det samme
    - accounting.UpdateBudgetReqs gui.updateBudgetReqs
        - Gor naesten 1-1 det samme.
        - Slet gui.updateBudgetReqs

- Ret GetTotal funktionen til GetBalance

- Skal Balances, Groups og Resume vaere exported
    - De burde vaere private med getter og setter metoder


- Update readme
        
- UpdateBudgetReqs skal ikke ligge cli
    - Det er en faelles funktion som bruges af alt UI
- Det skal vaere tydeligt hvilke "sheets" der har hvilket id
- ReadExcrptCsv burde ikke ligge i formatExcrptCsv filen
- har baade getExcrpts og getExcrptsFromSheets, der begge gor det samme
- Refactor
    - Liste af excrptGrps skal vaere et 1D array 
    - Brug multithreading til at match udtraek med excrptGrp
    - LoadExcrptTotal er lidt et maerkeligt navn, for det der goer hovedparten af logikken
    - Aendre tekst paa submit knap
- Kan man lave log.Fatal mere specifik 

- Maaske tilfoj accBalance til updateTotals funktion
    - Skal accBalance fjernes
- Den "eneste" forskel paa ui's er hvordan man finder matches. Ellers så er alt det andet faktis det samme.
    - Man kan saa bare have koden til at stå 1 faelles sted. 
    - Eller som vi har gjort nu, hvor vi har lavet en "backend", begge ui's bruger
    - Man kan samle ui's, saa man faktisk kalder den samme fil hver gang og saa bestemme med flag om man korer cli eller gui. De kan saa begge korer den samme read excrpts og auto-find matches. 
- Build til et andet system
    - https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs
    - https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures
    - https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags
    - For at byg til windows hvor en terminal bliver aabnet
        - https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications

- format haandtering af csv filen
    - Tror den fucker op hvis man har odt format vs text csv
    - Den pt kan den kun klare text csv
- Lav flow for gui delen
- Man skal ogsaa kunne traekke ud for flere aar af gangen.
    - Saa man skal baade angive aar og mdr
- Flyt ui's til at blive kort fra den samme fil
    - Den "eneste" forskel paa ui's er hvordan man finder matches for ukendte excrpts el. excrpts m. flere matches.
    - resten er ens
- handle excrpts i gui
- GUI, kan ikke handle at aabne chid Ignored
- Naar man afbryder et run skal man rulle tilbage
    - Ogsaa update af commongrps
        - Gaar denne igennem for det andet.
- Hvor skal / kan man se logs
- Tilknyt database?
- Naar man laver en error skal ens fields blive repopulated.
- util fil i controller module
- Validation
    - Skal kunne bruges til "debug" budgettet
    - Maaske mest interessent for de faste overfoerlser til faelles kontoen
- Sloring af sheet-id?
- Lav en quick way at opdatere matches paa, saa naar man ser en der burde have et match, saa kan man hurtigt opdatere den og man behoever ikke at huske at goere det bagefter
- Skal kunne huske hvilke mappings du har lavet foer
    - Saa behoever man ikke at lave en config fil. Hvis den selv finder ud af det
- God håndtering af overførsel mellem konti
  - Hvordan skal man vise at man overfører penge til en opsparings konto og så henter man så pengene ind igen når de skal bruges.
- Dette budget viser kun for budget konto'en, måske skal man også kunne se sine andre konti
- Lœs udtrœk fra PDF og ikke kun csv filer
- håndtering af lån og renter
- Kan man bruge concurrency

- Visning af excerptgrps skal ske i sortet rœkkefølge
  - Så det er den samme hver gang
  - Rimlig involved. Ved ikke helt hvordan vi skal gøre det på en god måde når vi gerne vil beholde parents.
  - Det skal ogsaa gerne vises i sammen raekkefolge paa gui'en. Saa det er alligevel noget der skal fixes
- Definer protokol for overførsler
  - Hvordan kan man håndtere den her situation med a-kassen. Hvor man tilføjer flere penge end hvad man havde planlagt
    - Denne situation skal ikke overskrive a-kasse beløbet men ligge til det 
    - Overfører samlet set x kr for forsikringer. y af de kr går til a-kasse.
    - Jeg overfører y kr til a-kasse. de bliver kun trukket hver 3 mdr. så skal det gœlde at 3y = hvad jeg betaler til a-kassen.
- Håndtering forskellig formattering af udtrœkkene
  - 1.000.000 vs 1000000
  - 1000,00 vs 1000.00
  - og dets permutationer
- opsœtning af excrptGrpData
  - Skal også vœre muligt at tilføje/fjerne opdelinger
  - Skal indsert i lowercase
- Kunne man automatisere det hele hvis man fik en bank-fuldmagt af brugeren
- Skal man gøre noget specifikt for at håndtere  kvartalvis/iregulœre overførsler til anden konto
- Kunne laese fra pdf filer
    - Kan blive et problem med formateringen.
