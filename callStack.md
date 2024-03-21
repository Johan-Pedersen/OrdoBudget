# callStack


- Hvad har jeg lœrt
- Et excrpt skal have flere mappings
  - Og så får man bare sådan en håndtering som ved andre ukendet
  - Men hvor der kun er de definerede valgmuligheder
- Håndtering forskellig formattering af udtrœkkene
  - 1.000.000 vs 1000000
  - 1000,00 vs 1000.00
  - og dets mutationer
- Visning af excerptgrps skal ske i sortet rœkkefølge
  - Så det er den samme hver gang
  - Rimlig involved. Ved ikke helt hvordan vi skal gøre det på en god måde når vi gerne vil beholde parents.
- parse udtrœk fra bank til rigtig format
- opsœtning af excrptGrpData
  - Skal også vœre muligt at tilføje/fjerne opdelinger
  - Skal indsert i lowercase
- Lav rewrite
  - Det virker ikke til at vœre lavet så smart
  - Specilet med den her parent situation
  - Det bliver mange dobblet forloop osv.
  - kan man fold de her json handling sektioner
  - getExcerptGrp er meget ineffektiv
- Kunne man automatisere det hele hvis man fik en bank-fuldmagt af brugeren
- Skal man gøre noget specifikt for at håndtere  kvartalvis/iregulœre overførsler til anden konto

