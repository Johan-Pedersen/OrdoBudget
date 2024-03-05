# Tanker

Google cloud project: *budgetautomation-414105* er det gamle projekt
nyt projekt: budgetautomation-414505

## nyttige links

- https://developers.google.com/sheets/api/guides/concepts
- https://console.cloud.google.com/iam-admin/iam?orgonly=true&project=budgetautomation-414505&supportedpurview=organizationId,folder,project
- https://developers.google.com/sheets/api/reference/rest?apix=true
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

- Man kan ikke altid sige "Føtex" skal gå til mad. Fordi føtex har mange andre ting
  - Så skal man kunne definere 1 el. flere grupper til føtex og så skal man selv angive hvilken gruppe dette skal i

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

- Hvordan finder vi den rigtige kolonne
  - vi kan hardcode det 
  - Vi har listen af excerpt grps, så man kan bare lave en match og se om elm[0].(string) matcher.
- Lav match på måned 
- nœste step
  - Så har vi rœkken og kolonnen og så er det bare at indsœtte data

## Resume

- Ud for hver udtrœk, skriver vi den udtrœksklasse den er kommet i
- Så har vi et map med den totale sum for hver enkelt udtrœksklasse som så er det vi sender over i regne arket.

- Tilsidst kan vi display hver enkelt udtrœk og dens udtrœksklasse
  - Enten kan vi vœlge at bare lave en sysout, ellers skal vi lave inserts i udtrœkstabellen. Så man også kan se historisk.

## access tokens

- Har jeg et refresh token
- Hvad er er refresh token

- Behøver jeg et refresh token når jeg bruger golangs client library
  - Burde den så ikke stå for det eller er det ikke det de mener?

- Man kan bruge Application Default Credentials(ADC)
  - Gør jeg allerede, men det virker stadig ikke.
  - Nu har jeg default credentials som altid bliver brugt

- Et refresh token bruges til at genopfriske et acccess token, som normalt udløber efter 1 time.
  - Et refresh token udløber efter 7 dage.
  - Der er et limit på 100 refresh tokens pr google account pr clinet ID
    - Men hvis man når sit limit, bliver gamle refresh tokens bare "overskrevet"
    - hvad betyder det

- Jeg har credentials til at få en kode, den kode kan jeg bruge til at få et access token, det access token bruger man til at snakke med api'et.


- Basic steps
  - Obtain OAuth 2.0 credentials from the Google API Console.
    - Done
  - Obtain an access token from the Google Authorization Server
    - Det er dette token der er udløbet
    - Den er nok udløbet fordi mit refresh token er udløbet.
    - Med en desktop app skal man bruge en code varifier der skal bruges til at få authrocation koden man skal bruge til at få access token.
    - Men hvordan kan det have virket før hvis ikke jeg havde det.
  - Examine scopes of access granted by the user
  - Send the access token to an API
    - Det står der ikke rigtig noget om hvordan man gør
  - Refresh the access token, if necessary
    - Når ens access tokens udløber skal man bruge et refresh-token 

- Hvornår skal man bruge et refresh token
  - Der er 7 dags udløbstid på et refresh token
  - Hvornår skal man bruge er nyt access-token
    - Det er hver gang man skal have adgang til api'et
    - Hver access-token varer i 1 time.
  - Men betyder det man kun kan have 100 refresh tokens af gangen
    - Det betyder så at man kun kan have access til 100 API's af gangen.

- 

## Excrpt grps på JSON format

- Vi skal angive excrpt grps og udtrœks mappings

- ideelt set angiver man excrpt grps i en lang liste og så skal man angive man hvilke udtrœk der mapper til hvilke grps
- Det skal jo så bare marshalls til Json form

## Brug regex til excrptGrp matching

- Hvordan vil man bruge regex matching til maps
  - Skal man så selv en map data struktur
  - Du får jo et langt input og så skal du match det bare en lille del at det.
  - Det bliver man jo nødt til at lave en speciel algoritme.

## Opdeling af excerptgroups

- Kan man lave et generisk array af json objekter
