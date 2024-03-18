# callStack


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
