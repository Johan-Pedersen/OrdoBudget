# callStack

- For hver mdr skal man lœse de faste udgifter og sœtte baseline grpTotals baseret på det.
  - Og ved at den lœser dem fra excel arket. Så bliver disse automatisk tilpasset hvad man har til at stå i arket.
- Den skal kun lave udtrœkket for den måned man har gang i.
  - Med mindre man laver det for flere måneder.
- Et excrpt skal have flere mappings
  - Og så får man bare sådan en håndtering som ved andre ukendet
  - Men hvor der kun er de definerede valgmuligheder
- Update totals baseret på excel arket i stedet for sheets
  - der er ingen grund til at upload udtrœk til google sheets for at hente dem ned igen når man skal update Totals
- God håndtering af overførsel mellem konti
  - Hvordan skal man vise at man overfører penge til en opsparings konto og så henter man så pengene ind igen når de skal bruges.
- Dette budget viser kun for budget konto'en, måske skal man også kunne se sine andre konti
- Lœs udtrœk fra PDF og ikke kun csv filer
- håndtering af lån og renter

- Definer protokol for overførsler
  - Hvordan kan man håndtere den her situation med a-kassen. Hvor man tilføjer flere penge end hvad man havde planlagt
    - Denne situation skal ikke overskrive a-kasse beløbet men ligge til det 
    - Overfører samlet set x kr for forsikringer. y af de kr går til a-kasse.
    - Jeg overfører y kr til a-kasse. de bliver kun trukket hver 3 mdr. så skal det gœlde at 3y = hvad jeg betaler til a-kassen.
- Kan man lave log.Fatal mere specifik 
- Håndtering forskellig formattering af udtrœkkene
  - 1.000.000 vs 1000000
  - 1000,00 vs 1000.00
  - og dets permutationer
- Visning af excerptgrps skal ske i sortet rœkkefølge
  - Så det er den samme hver gang
  - Rimlig involved. Ved ikke helt hvordan vi skal gøre det på en god måde når vi gerne vil beholde parents.
- opsœtning af excrptGrpData
  - Skal også vœre muligt at tilføje/fjerne opdelinger
  - Skal indsert i lowercase
- Kunne man automatisere det hele hvis man fik en bank-fuldmagt af brugeren
- Skal man gøre noget specifikt for at håndtere  kvartalvis/iregulœre overførsler til anden konto

