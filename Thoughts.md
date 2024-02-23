# Tanker

- go bibliotektet
  - https://pkg.go.dev/google.golang.org/api/sheets/v4
- API doc
  - https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets/
- Sample requests
  - https://developers.google.com/sheets/api/samples

## Todo

- Er det feasible at hente hele sheet'et 1 gang og selv lave alle operationerne
  - Det skulle vœre for at minimere # requests
  - Tror det er nødvendigt at lave så meget som muligt lokalt.
- Som udgangspunkt er det nok nemmere bare at hente 1 kolonne af gangen

- Bestem en sortering
- Hent bank udtrœk
  - Kan create et nyt sheets med denne data
- flyt data til de rigtige kolonner 
  - Det er nok mest rart at vi ikke bare har alle udregninger / data lokalt da man så kan se på dem hvis der er noget der ikke helt stemmer
  - Det skal også vœre muligt at selv justere de resterende udregninger man ikke kan vide på forhånd
- Til sidst skal alle kolonnerne summeres og lœgges ind i måneden der passer med dette udtrœk.
  - Man kan matche det med den dato der står i udtrœkket.
- lœs udtrœks grupperinger fra config og opsœt formattering
  - 
- Brug regex til mapning
  - Brug regexp.MatchString(pattern String, s String)
  - Lav et ignorecase match
- Hvordan skal man håndtere forsikringer og fœlles udgifter
  - De bliver jo bare taget totalt set fra min konto og bliver kun "delt ud" når de bliver trukket fra fœlles kontoen
  - Det samme gœlder "Faste Udgifter"
  - Så skal man selv dele dem op. Men det virker heller ikke helt smart.

  - Det er måske fint bare at holde det sådan her og når man så trœkker data fra sin egen lønkonto. Så er det kun "Hverdag" man opdatere.
- automatisk find den rigtige indsœttelsesrœkke 

## Nordeas's csv fil

- kolonnerne er forskudt med 1 hvis der er brugt 'Nordea Pay'
- Hvilke kolonner skal jeg bruge
  - A (bogførings dato)
  - B (Beløb)
  - G(Tekst)
    - Hvis Teksten er "Nordea Pay køb", så skal man også kigge på H. 
  - H( i normal case, saldo i heltal)
    - ved "Nordea Pay" er det også en del af teksten.
  - I (i normal case, øre for saldo)
    - ved "Nordea Pay", er dette saldo i heltal
  - J
    - ved "Nordea Pay", er dette øre for saldo
  


## Upload CSV fil

- Ark = sheet

- Update udtrœks sheet for hver nye udtrœk
  - Gøres med https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets/request#updatecellsrequest

## Hoved funktionalitet

### Flyt data til de rigtige kolonner (udtrœk)

- Bestem format
  - Man kan samle Tekst i 1 folder indne man uploader det
- Flyt data
  - Kan dette gøres lokalt
- Flyt data over i hoved sheet

### Insert sum data i nyt sheet

- Hvordan angiver man et nyt sheet til insert
- Kan man lave en stor updateReq med alle data i 
  - Det burde man virkelig kunne gøre. 
  - Specielt når data jo bare hœnger sammen.
- Hvorfor skal man i read range både angive sheet navn og i get skal man angive sheet ID, når kun 1 burde vœre nødvendig

### insert data i den rigtige kolonne

- 

## Resume

- Ud for hver udtrœk, skriver vi den udtrœksklasse den er kommet i
- Så har vi et map med den totale sum for hver enkelt udtrœksklasse som så er det vi sender over i regne arket.

- Tilsidst kan vi display hver enkelt udtrœk og dens udtrœksklasse
  - Enten kan vi vœlge at bare lave en sysout, ellers skal vi lave inserts i udtrœkstabellen. Så man også kan se historisk.


