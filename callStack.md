# callStack


- Overskriv alle felter der ikke har fået en vœrdi i totals
- Visning af excerptgrps skal ske i sortet rœkkefølge
  - Så det er den samme hver gang
- håndter vi korrekt positive tal i udgifts sektionen
  - Det er jo rigtig nok at det skal skifte fortegn. selvom det ikke er så pœnt
- Fejl når man selecter gruppe nr uden for mœngden
- håndter blanke linjer i udtrœk
- Test
- parse udtrœk fra bank til rigtig format
- opsœtning af excrptGrpData
  - Skal også vœre muligt at tilføje/fjerne opdelinger
  - Skal indsert i lowercase
- Lav rewrite
  - Det virker ikke til at vœre lavet så smart
  - Specilet med den her parent situation
  - Det bliver mange dobblet forloop osv.
  - kan man fold de her json handling sektioner
- lav assert i forhold til om saldo passer med udtrœk
