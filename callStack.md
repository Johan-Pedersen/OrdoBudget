# callStack


- parse udtrœk fra bank til rigtig format
- For hver mdr skal man lœse de faste udgifter og sœtte baseline grpTotals baseret på det.
  - Og ved at den lœser dem fra excel arket. Så bliver disse automatisk tilpasset hvad man har til at stå i arket.
- Definer protokol for overførsler
  - Hvordan kan man håndtere den her situation med a-kassen. Hvor man tilføjer flere penge end hvad man havde planlagt
    - Denne situation skal ikke overskrive a-kasse beløbet men ligge til det 
    - Overfører samlet set x kr for forsikringer. y af de kr går til a-kasse.
    - Jeg overfører y kr til a-kasse. de bliver kun trukket hver 3 mdr. så skal det gœlde at 3y = hvad jeg betaler til a-kassen.
- Kan man lave log.Fatal mere specifik 
- Et excrpt skal have flere mappings
  - Og så får man bare sådan en håndtering som ved andre ukendet
  - Men hvor der kun er de definerede valgmuligheder
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
- Lav rewrite
  - Det virker ikke til at vœre lavet så smart
  - Specilet med den her parent situation
  - Det bliver mange dobblet forloop osv.
  - kan man fold de her json handling sektioner
  - getExcerptGrp er meget ineffektiv
  - Det er meget grimt at alle api requests skal kaldes fra main()
    - Meget grim og cluttered main()
  - Hvis en funktion returnere en error, skal den bare fejle og ikke smide den ud til kalderen
- Kunne man automatisere det hele hvis man fik en bank-fuldmagt af brugeren
- Skal man gøre noget specifikt for at håndtere  kvartalvis/iregulœre overførsler til anden konto

