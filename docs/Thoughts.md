# Thoughts

Google cloud project: *budgetautomation-414105* er det gamle projekt
nyt projekt: budgetautomation-414505

## Forklaring

- Parent: ignored
  - Det er til udtrœk som fx "johan til fœlles", som ikke har nogen decideret gruppe og bliver håndteret gennem sheets, uden om dette system.
  

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
- Google Sheets
- https://docs.google.com/spreadsheets/d/1Dg3qfLZd3S2ISqYLA7Av-D3njmiWPlcq-tQAodhgeAc/edit?gid=1685114351#gid=1685114351

- Build til et andet system
    - https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs
    - https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures
    - https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags
    - For at byg til windows hvor en terminal bliver aabnet
        - https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
### fyne

- https://docs.fyne.io/explore/widgets
- https://pkg.go.dev/fyne.io/fyne/v2
- https://docs.fyne.io/started/

## Todo

- Der er ingen grund til at upload filen til google sheets for at hente den igen til update totals
- Er det feasible at hente hele sheet'et 1 gang og selv lave alle operationerne
  - Det skulle vœre for at minimere # requests
  - Tror det er nødvendigt at lave så meget som muligt lokalt.
- Som udgangspunkt er det nok nemmere bare at hente 1 kolonne af gangen

- Bestem en sortering for udtrœks grupperne

- Oprettelse af nye felter

- Fjern "faktisk balance", da det er forvirrende
  - Kan bare have et lokalt tjek, der sammenligner konto balance og faktisk balance og kommer med en error / besked hvis der er merkant forskel.
  - Kan faktisk godt lide at have den. Så måske vi skulle beholde den

- Hvad hvis man får penge tilbage og vil fordele dem mellem konti?

- Dette sheet viser kun budget konto'en. Det skal også vœre muligt at se alle konti.

- Lœs PDF udtrœk og ikke kun .csv fil

- håndtering af lån og renter

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

## Angiv måned

- Hvis der er nogen der ikke har en måned, så må det vœre fordi at de er "kommende" og ikke nødvendigvis i denne måned
  - Så dem skal man faktisk ikke tage med.

## Test

- Kan lave insert for Jan Dec
- i udtrœkket kan specified month ligge først, i midten og til sidst
- Udtrœk med blank lines
- Udtrœk med store og små bogstaver
- udtrœk med mellemrum
- Udtrœk med positiv value

## Hvad har jeg lœrt

- Skal huske at overføre det rigtige beløb til Fœlles konto'en
- Håndte kvartal vis / tilfœldig justeringer af fœlles udgifter
    - Hvis der pludselig kommer en "vand" opkrœvning og man så bruger sin egen lønkonto, i stedet for fœlles.

- Håndtering af at kunne bruge forskellige konti (fx at trœkke fra opsparingen)

## Håndtering af kvartalvis kontigenter osv.

- Man Søger bare for at dele beløbet op i 3 og så er det bare det man overfører hver måned

- man kan også lave en kvartalvis overførsel til fœlles konto'en, som bare passer til det beløb 
- Hvis man bare delte overførslerne op så havde vi ikke dette problem.

- Men det er noget der skal kunne håndteres
  - 
- casen: man overførere et fast beløb hver mdr til at dœkke alle forsikringer
  - Så kommer der en surprise opkrœvning (der hører under en eksiterende gruppe), som man betaler fra sin egen konto.
  - Det betyder begge disse udtrœk skal ligge i den samme gruppe.
  - Men fordi den faste allerede er bestemt, så bliver den overskrevet af den eksisterende
  - Det hele kommer fra at man ikke har en individuelt overførsel for her ting
  - Men man kunne jo definere en ny default vœrdi for de faste udgifter, så hver gang der kommer en surpise, så bliver det bare lagt til. 
  - Skal man have en verifier der tjekker om den overføre vœrdi stemmer med hvad der trœkkes.
  

## parse udtrœk fra bank til rigtig format

- Vi skal opdatere udtrœks sheets før vi kan gå videre med det andet
- Så vi skal have en eller anden form for mekanisme, så man ikke kan gå videre før vi har retureneret fra udtrœks update.

## base-line vœrdier til faste udgifter

- Vœlg hvilke Parent's der skal have andre baseline vœrdier
- lœs dem 

## hvilke parents skal have andre baseline vœrdier

- Fœlles udgifter
- mad
- Forsikring

- Alle vi laver en 'total' overførsel for

- Hvad hvis ikke der står noget

- man skal have en måde at fortœlle den at den skal hente disse vœrdier fra sheet

- Hver gruppe skal have et override flag. Som er true per defalt
  - Men hvis ikke, så skal man ind og finde den

- Man definerer hver gruppe
- Efter man har initieret alle excerptGrps, så løver man dem igennem igen og ser hvilke nogen der har override flagget. Som man så skal ind og hente data fra sheets til.
- Det skal ske til sidste i init excrptgrps metoden

- Det er trœls at man skal have SpreadSheetId og SpreadSheetsService med i alle funktioner som skal bruge det.
  - Det det skal gives med helt ude fra main funktionen og ind til den inderste funktion der skal lave en form for api kald
  - Kan ikke vœre rigitgt at at det er nødvendig
  - Det burde vœre en form for global variabel der var accessible genenm hele programmet.
    - Men hvordan laver man det og er det overhovedet en god ide?


## Sundhedsforsikring bug

- Vi skal vide hvad values er
- Vi faar et array. Men af hvad?
- Nogen gange er det et array af arrays og andre gange ikk?

- Vi faar en valueRange m. range, MajorDimension, values
    - values
        - Array of Arrays
        - outer array er alt data man har bestilt i sin Range
        - inner arrays er hver en major dimention(?)
    - Major dimention
        - Major dimentions bestemmer om et inner array skal repraesentere en Row eller coloumn

- Men vi faar bare et tomt array
    - Saa det siger jo der ikke er nogen data.
    - Men det er jo ogsaa rigtigt
    - Der er bare forskel paa om vi faar et tomt array eller om vi ikke faar noget array

    - Intet array siger jo at der ikke er nogen major dimention, hvor tomt bare siger der ikke er nogen data

- Det skyldes en eller anden formatering paa sundhedsforsikrings raekken saa der ikke kommer nogen value fra den. Men vand raekken har ikke det problem
- Saa man kan bare kopiere vand cellen og saette ind
- Kan det vare det bare et et whitespace?

## Madkonto udgiften bliver sat ind i overskrifts linjen paa madKonto

- Hvordan fungere ignored
    - Har name i ignored gruppen noget at sige

- Madkonto'en staar baade i config filen til bare at have ingen matches, men ogsaa som ignored
    - Det samme goer faelles udgifter og forsikringer(?)

## Man kan lave et check om de faste overfoerelser passer med hvad der staar i sheets

- Med denne kan man se forskellen paa indtastet og aktuel overfoert

- Skal det vaere et faelles pgm, eller skal det bare vaere til os?
    - Det skal jo bare vare til os. Man kan altid lave eksempler.

- Hvor relevant er det hvad vi overfoere og hvad der bliver trukket paa den anden konto
    - Det handler jo bare om der er + el. - paa faelles konto'en. Det er jo somsaadan ligegyldigt for vores egen bankkonto.
    - Saa kan den saa vaere sket en fejl 

## Et excrpt skal have flere mappings

- Og så får man bare sådan en håndtering som ved andre ukendet
- Men hvor der kun er de definerede valgmuligheder

- Hver gang skal man saa loebe igennem alle matches, det giver saa en liste som man saa kan fremstille i en select correct excrpt Grp.
    - Ligger op til at komme dette i en funktion for sig
    - Funktion til vaelge match

## Lav "hash" funktion. Saa matches altid har den samme index

- Hvor bruger vi unordered lists 
    - excrptGrpTotals

- Hvis vi havde hashfunktioner kunne man sikkert undgaa mange af disse dobbelt loops
    - Det bliver specilt svaert jo flere punkter man har
- Saa skal man nok ikke goere den 2 dyb, saa er det nemmere hvis man ikke har noget parent.
    - Man kan vel hash parent + excrptGrp og saa faar man noget unikt ud. 
    - Hvis man bruger hashing, saa er der jo altid risikoen for collision.
        - ss for collision vil jo altid vaere 1/n

- Hvad vil vi gerne opnaa
    - Man kan paa konstant tid faa en excrptGrp baseret paa et givent udtraek
        - Nej, udtraekkene er jo det fulde udtraek, saa det kan man ikke bruge som key.
    - ExcrptGrps listen skal altid staa i den samme sorterede raekke foelge
        - Hvis dette gaar i mod foerste punkt, kan vi bare lave en sorterings algoritme til vissing af.
        - de skal godt nok have det samme nummer hver gang.
            - De behoever ikke et nummer hvis vi laver det som en gui.
            - Men for nu er det maaske stadig en god ide

### Konstant tids opslag

- Find en excrpt grp for et givent udtraek
    - Kan man nok ikke goere paa konstant tid. Da udtraekkene er lange og vi kun skal match paa et enkelt ord i den liste.
    - Man kunne lave split paa alle ordene, og saa kan man bruge multi threads til at slaa alle ordene op paa samme tid.
    - Og da det bare er en get, behoever de ikke faa en lock paa excrpts listen
    - multithread losningen er dog ikke god hvis man skal have mellemrum i sine matches

- Det er et dobbelt for-loop hver gang man skal finde en excrpt Grp
    - Fx. naar man skal lede efter en med et givent "index". 
    - Hvis alle grupperne skal have et unikt index, saa givere det bedst mening bare at have det i et 1D Array

- Hvilke situationer skal vi bruge det i 
    - Find excrptGrp der matcher et givent udtraek
    - Find excrptGrp der matcher et givent "index"
        - hvis man bruger en 1D liste, kommer denne del automatisk
    - GetExcrptGrp baseret paa name, eller index


### Find excrptGrp der matcher et givent udtraek

- input 
    - streng x, med n ord separeret af ' '
- output
    - liste af ord y1,y2,... yj, der indgaar i x.

- subproblem, indgaar yi i x 
    - contains(yi, x).
    - Lob alle matches igennem.
        - Det gor jeg nu, og er meget langsomt.

- Men hvis man vil have alle matches, saa skal man jo ogsaa lobe alle matches igennem.
- Er der nogen maade at skaere matches fra?

- Multithreading er den bedste losning



## Refactor

### Omdan projektet til en MVC struktur
### Liste af excrptGrps skal vaere et 1D array 
### Update totals baseret på excel arket i stedet for sheets

## GUI

### links

- https://docs.fyne.io/explore/widgets
- https://pkg.go.dev/fyne.io/fyne/v2
- https://docs.fyne.io/started/

### Databinging

- Hvad er databinging
    - Det synkronisere 2 datakilder 
    - Det kan for exemlpel vaere smart at have en databinging mellem UI og Storage

- Kan vi bruge databinging
    - Man kan vel bruge det til form entry, men vil egenligt hellere have det gennem en anden "kanal"
    - Man kan jo bruge det i alle data input fields og saa match dem med en variabel.
        - Kraever det ikke at det er globale variable
            - Kommer jo bare an paa hvordan det er lavet.
        - Vil helst have at alt data kommer ind fra controlleren.

- Skal man bruge data binding til at faa sin form data
    - Hvad gjorde de for der fik databinging
    - Det virker i hvertfald til det er en mulighed

### vis current excrpt

- Det er maaske oplagt at kigge paa noget databinding til dette
- Hvad kan man ellers gore?
    - Man skal jo bruge noget dynamisk opdatering

- Man skal ogsaa kunne knytte en form for action til hver tree element
    - Der har man en OnSelect funktion. Saa man kan laese excrtpTextfeltet og lave en update reg baseret paa det


### Tanker

- Aendre submit btn, til at sige kor

- Hvordan faar man data fra et submit
    - Er det brug af callback
        - Hvad er det

- Alt setup kan vel saadan set gores for man trykker submit.

- Hvad er en callback funktion
    - continuation-passing style
    - Et "callback" er en funktion, der kaldes af en anden funktion
        - Hvor den forste funktion tager funktionen som en parameter
        - Saa det er en hojer ordens funktion
    - Og det er saa et callback fordi vi giver funktionen der skal kore i forskellige events

    - Hvordan er den forskelligt fra en normal funktion
    - Hvad er fordelene ved en callback funktion
    - Hvorfor bruger fyne altid callback funktioner

- NewTree
    - ChildDIDs
    - IsBranch
    - CreatNode
    - UpdateNode

    - Hvorfor skal vi give disse callback funktioner?
        - ChildUID, burde haandteres af strukturen
        - IsBranch, virker ogsaa til at vaere en basal struktur opereation
        - CreatNode, basal operation
        - UpdateNode, Hvis det bare er objekter vi har, saa burde dette heller ikke repraesentere et problem.
    - Hvor kommer tni fra?

- Skal man give disse callBack funktioner hver gang


## omdan projekt struktur

- https://github.com/golang-standards/project-layout

- Lav cmd, internal, ui, pakker
- i internal, lav pakker baseret paa hvad denne del af koden giver.
    - Lidt "grop by context" i DDD verdenen


## Lav en test mode, saa den bare loader en predefinet datasaet

- Det er jo en form for debug mode
- Hvad skal den gore
    - Vi vil gerne undgaa at skulle load en hel masse naar man skal debug
    - Data'en aendre sig jo ikke, saa det er bare et sporgsmaal om at hente data'en fra google sheets i stedet for at skulle upload den og hente den.
    - Man kunne maaske ogsaa bare have noget test data man kunne load ind i vores structs.
    - Saa gaar man helt uden om noget som helst API, saa kan man ogsaa kore det offline.
        - Smart i toget

    - Hvilke steps kan vi springe over
        - Update af excrptSheet
        - Read fra Udtraek
        - Bestem Person
        - Bestem mdr
        - InitExcrptGrps
- Hvordan aktivere man den 
    - Man kan bare saette den som et flag inden man buider
    - eller kan man vel give den som et flag naar man kore programmet
        - kunne vare en sjov maade at gore det paa.

- Hvad skal man bruge for at kalde LoadExcrptGrp
    - Excpts der er hentet fra sheets
        - Det skal man vel bare gemme som en JSON fil, man saa kan unmarshal naar man skal bruge den.
        - Hvad skal vi saa unmarshal/decode fra JSON
            - excrpts fra sheets.ValueRange
            - excrptGrps
            - excrptGrpTotals

    

## Saet debug mode op til at kore gui

- Saa hvordan skal vi gore det  
    - Hver mode kan have en debug mode. Som saa loader predefined data 
        - nok smartest. 
        - Saa kan vi definere noget logik til at load det. (Som vi har gjort)
        - Det er jo bare i init delen at man skal bestemme hvor man loader fra
            - Man kan vel have en shared init del paa tvaers af cli og gui. 
            - Der sker jo det samme, problemet opstaar bare i forhold til maaden at give input
        - Man kan tage imod input, og saa kalde en init funktion. Som bare tager input datasaet og mdr. 
        - Med et debug flag, kan man saa bestemme om det er debug init el. rigtig init man vil kore.

- Vil det vaere smart at have en make fil
    - Hvad ville man bruge den til

## Decouple ui og internal

- Man skulle maaske bygge det mere som et API?

- Det handler bare om hvordan vi vil decouple front og backend

- Det er maaske en fin ide at bygge det som et REST API

- Giver det mening at bygge det som et API
    - Det skal i hvertfald afkobles, fordi det er det reneste
        - Det kan jo ogsaa gores med at lave et model interface, som man saa benytter i ui
    - Lav et interface til decoupling, med de funktioner der skal bruges af ui, i en stateless maner.
    - Det bliver svaert at lave den stateless, i forhold til at vide hvilken et excrpt der er det naeste og om vi har haft det excrpt for.
    - Men det styre UI'en jo bare. Fordi ui siger bare giv mig alle excrpts.
    - Dem gemmer de saa og fremviser
    - i deres run tager de saa bare den naeste.
    - backendens primare opgave er jo saadan set bare at hente og gemme data. og sorge for det bliver gjort korrekt

- Men nu har vi ogsaa bare en cli, som skal understottes.
    - Her kan man saa finde ud af om man vil have 1 interface til hver
    - I princip er det jo det samme. pt, er cli'en bare bygget ind i backenden.
        - Det skal den ikke vaere

- Man skal tage alle inteaktionerne ud af furktionerne.
- Saa updateExcrptTotal, skal bare returenere excrptGrpMatches, og saa er det ui's ansvar at vaelge det rigtige match
- Er det saa viewet's ansvar at update excrptGrpTotals og update Resume med det rigtige
    - Det ansvar burde bare ikke ligge i ui.
    - Her kan man vel bruge en callback funktion -> 
        - Det tror jeg ikke ville fungere saa godt med GUI
- Hvordan sikre man sig saa at Resume og excrptGrpTotals bliver opdateret
    - Det skal kaldes fra updateExcrptTotal funktionen, men input til UpdateResume og selve excrptGrpTotals updaten afhaenger af hvilken gruppe man har valgt

- API'et skal bare return alle excrptGrpMatches
    - Det skal vaere en anden funktion
    - UI, kan saa gore hvad de vil med det 

- UpdateExcrptTotal, skal kun update ExcrptGrpTotals array'et og ikke andet
    - pt. gor den for meget
    - Det er maaske fint at den kalder UpdateResume. Ellers glemmer man i UI at updateResume
    
    - Hvordan skal vi gore med state
        - hvordan faar vi fat i excrptGrps
        - Det giver ikke mening at give en pointer med, fordi i teorien behover API og app ikke at vaere paa samme maskine

        - Ellers skal man til at sende excrptGrps og parents hver gang man faar et nyt excrpt
        - og det virker meget upraktisk
        -
        - Det nemmeste virker til at have "Serverside" caching, som nu. Og hvor der ikke findes en enkelt decideret server, saa virker dette fint.


- Haandtering af flere matches for 1 udtraek
    - Det er jo op til ui'en at bestemme hvordan man vil haandtere denne konflikt. Vi skal bare vide hvilken en vi skal update.
    - Men det er maaske det forkerte sted dette ansvar ligger
    - Vi skal nok bare udstille en funktion der henter alle matches, og saa skal ui lave en funktion der haandere konflikten.
    - Eller saa skal man sige tandle total tager en hojre ordens funktion, saa man er tvunget til at bestemme det. 
    - Men det er jo ikke saa tit at gui funktioner returnere en vaerdi, saa det vil ikke passe saa godt.
    - vi maa bare angive dem hver for sig

    - Det traels ved forst at haandtere dobbelt matches el no-matches i ui. Er at man skal finde dem igen.
        - Men maaske kan man gore noget smart.
        - Men man kan jo ikke vide om det er det ene eller det andet

## lav ui/cli del


## graens mellem ui og internal

- Hvor er graensen mellem interal og hvad skal ligge i ui
    - Internals ansvar
        - Det kan ikke kun vaere de metoder der kan ligge i begge ui's
        - Det skal vaere byggeklodserne / frameworket til appliktionen
        - Saa hvis ikke det er en byggeklods, skal den ikke ligge her
    - ui's ansvar
        - Det skal bare vare limen / brugen af internals

## handle excrpts i gui

- Databinding til vise current excrpt
- Hvis et excrpt har flere matches, saa er det kun dem der skal vises i dropdownen.
    - ellers, er det dem alle der skal vises.
- Naar man dobbelt klikker paa en gruppe, skal totals opdateres og der skal hentes et nyt current excrpt.

### vis current Excrpt

- Problemet er at man skal update bindet naar man har assignet et excrpt til en grp.

- Men vi kan have et async loop til at korer alle excrpts i gennem
- Det staar og lytter paa en channel, og mens det gor det er det blocking
    - Saa vi ikke updater bind for den er klar.

- Det er det loop updater vores string binding

- vi skal ogsaa have en function der kaldes onSelect
    - Den skal skrive paa channel'en og tage den selected value, samt excrpt(Maaske fra bind?) og saa update Totals
    - Kan vi faa fat paa excrptet som struct i stedet for string

- Vil man hente excrpts over i en struct?

## fix loadExcrptTotals

- Vi skal finde en maade hvor metoden baade kan bruges af cli og gui
- Til det kraeves at man har en metode til at bestemme hvilken gruppe man skal tage naar der er flere matches.

- Man kan maaske splitte algoritmen mere ud, saa den i sig selv ikke laver lige saa meget?
- Man kunne ogsaa bare faa den til at returnere alle excrpts der er tvivl om, og saa kna ui delen selv tage stilling til hvad der skal ske med dem.

- Man vil jo gerne have at en metode kun gor 1 ting
- Maaske skulle man lade en anden metode finde alle matches og saa kalde dem i et loop med updateExcrptsTotal.

- Saa adskiller man det med at finde og update, hvor loadExcptsTotals prover at gore det hele.

- tidskomplektiteten er det samme. Det er bare en enkelt faktor i forskel.
    - Det er ogsaa den mest simple losning

- UI'en skal have selMachtGrp funktionen.
- naar de har gjort det kan de kalde UpdateExcrptTotal
- for at de kan kalde den, saa skal man forst decode excrpts.Values, og det er nok der loadExcrpts kommer ind i spil.
    - Her vil det jo hjalpe hvis man havede en struct til at laese excrpts.Values over i. fordi saa kunne man bare have en funktion der gjorde det. 
    - Saa har man en collection af excrpts, som ui derefter kan bruge.

- Hvorfor skal Date i Excrpt struct vaere float
    - Kan det betale sig at lave den til en date type
    - Man kan sige det fungere nu?

- Hvad er dens ansvar
    - UI skal sorge for at finde finde / definere alle matches
    - Derefter skal vi kore dem alle i gennem og UpdateExcrptTotals
    
    - Man kan laese alle excrpts fra CSV'en. De bliver saa laest over i en excrpt struct.
        - Men laeser LoadExcrptTotal filen paa samme maade som ReadExcrptFromCSV
            - Nej, ReadExcrptCsv er ikke lige saa haardfor
            - Den laeser ogsaa hele CSV filen uafh af mdr og ikke ligesaa god error handling
    - Men hvis vi flytter over til at laese fra csv filen, kommer det med at vi ikke skal laese excrpts fra sheets automatisk.
        - Men det er saa lidt en anden opgave ligepludslig
        - Skal have rene linjer, saa tingene ikke flyder sammen og bliver forvirrende.


- Todo:
    - ReadExcrptsCsv
        - laes over i excrpt struct -> J
            - Den skal forfines, saa den er mere haardfor
    - funktion til at finde matches for CLI og GUI 
    - funktion til udregning af account balance
        - update af totals
    - vis excrpts og excrptGrps i GUI
        - stor del
    - Ryd op saa man ikke laengere updater 'temp' sheets med variable


- funktionen ligger det forkerte sted

### Test ReadExcrptCsv

- empty file
- malformed file
- good file
- test cmpMonth functionality.

- maaske skal vi se om vi kan input strings, i stedet for vi skal have saa mange test filer liggende.


### implementer hele flowet for cli

- Den "eneste" forskel paa ui's er hvordan man finder matches. Ellers så er alt det andet faktis det samme.
    - Man kan saa bare have koden til at stå 1 faelles sted. 
    - Eller som vi har gjort nu, hvor vi har lavet en "backend", begge ui's bruger
    - Man kan samle ui's, saa man faktisk kalder den samme fil hver gang og saa bestemme med flag om man korer cli eller gui. De kan saa begge korer den samme read excrpts og auto-find matches. 
    - 

- i alle steps, lober/haandtere man samtlige excrpts.
    - Det er jo ikke eksponentielt bare en faktor, men det er stadig mange gange
    - Den eneste man kan skralde vaek er "update totals"
    - Man kan ogsaa smide update budget loopet vaek
        - Hver gang man har defineret et match, saa kan man lave en update-req
        - i samme omgang kan man update totals

- Hvorfor har vi en totals
    - Den har vaeret til, saa man forst kunne finde alle excrpts og matches. For der efter at upload dem til google sheets.
    - Men hvis vi laver requestet med det samme, saa er der ingen grund til at skulle gemme det i totals.
    - grunden til vi har totals, er fordi for vi updater, saa skal man finde det totale for hver excrptGrp inden man updater sheets.

    - Men det kan man saa komme uden om hvis man kan laegge en vaerdi til det der allered staar i cellen.

- bliver det for meget i forheld til man ogsaa skal vide hvilken celle man skal opdatere.
    - Man kan jo bare hente kolonnen med alle grupperne ud og ligge det over i et map.
    - Saa kan man slaa excrptGrp'en op i mappet for at se hvilken raekke det skal ligge paa 

- Saa skal man ogsaa have logik til at bestemme hvornaar skal man overskrie cellen og hvornaar skal man ligge til.
    - Tror der er mange ting der kan gaa galt i det flow

- Kan det vaere fordi Kirsch golf matcher baade paa "kirsh golf" og "golf"

#### laes exclpts fra csv -> J
#### auto find matches

- Er der forskel paa en et match en excrptGrp
    - ExcrptGrp er selve structen
    - et match er et faktisk math mellem ExcrptGrp og excrpt
    - Men det bliver indkaplset i instantieringen af et excrptGrp

- Meget af denne logik ligger allerede i excrptgrps -> skal den det 
    - Man skal bare rette den til saa man hver gang opdater totals.
        - Det gor vi maaske allerede

    - Det var i hvert fald svaert at finde om jeg haved denne logik.
        - Men jeg taenkte heller ikke rigtig over hvad excrptGrps faktis stod for?
        - Foler maaske den skal flyttes.


- Hvordan skal vi opdele pakkerne
    - Hvilke ansvars omraader har vi
    - parser giver god mening, fordi her skal vi parse forskellige filer fra div banker
    - 

    - Laerser fra CSV
        - parseren
        - Til det horer ogsaa excepts.

    - Find auto-matches
        - ved hver match bliver balance updated
        

    - Find resterende matches
        - update balances ved hver match


    - Update google sheets

- Excrpt skal ligge i parser pakken 
    - Hvordan skal man haandtere det med at have defineret typer osv.

- Ansvars omraader i Entry

    - Balance collection
        - haandtering af balance
        - bruger entry
    - Resume collection

    - Entries collection
        - Init collection
        - create
    - Groups collection
        - Create

    - google sheets updates
        - UpdateCommonGrps
        - UpdateExcrptSheet
        
    - Find matches
        - Matcher excrpts med entries + groups

- Entries + Groups

- Balance + Resume
    - Benytter entries + Groups

- Entry pakken kan kaldes accounting
    - Kan holde Entries + Groups + Balance + resume
    - Er det for generelt
        - excrpts kan ogsaa vaere i accouting
            - vi saetter en granse for hvad der horere til input, og hvad der er accunting
    - budget
        - Vi har typerne til at efterligne budget
        - budget kan ogsaa holde alt.

- google Sheets
    - Benytter excrpts
    - ingen grund til at ligge her


- find matches


- InitExcrptGrps / initEtriese skal ligge i parser


- Skal man have en accounting pakke til alt med Balanc, Resume og CommonGrps
    - Skal CommonGrps ikke hedde fixedExpenses?
        - Det er vel ikke nodvendigvis "fixed expenses", det er bare hvad den bliver brugt til

## Fix UpdateCommonGrps

- Det skal gores mere effektivt

- Denne funktionalitet burde ligge hvor man initialisere Balances.

- Det er ikke saa godt design, at ikke rigtig tage hojde for fixed-expenses for til allersidts hvor vi updater sheets. **Fordi Balances viser dermed ikke det sande billed.**
- Hvis Entry.Ind passede til hvilken raekke den havde i sheets. Saa ville man ikke rigtig have et problem. Fordi saa kunne man bare lave en get, naar man instansiterede Balance tabellen
    - Det betyder saa, at vi bliver nodt til at lave en en automatisk getEntries funktion, der henter og indexere alle entries
    - Kan hurtigt blive et problem, fordi hvis der saa er nogen der laver om i kolonnerne saa bliver det hele fucked up
        - Det kan vaere man kan laase sheets, eller lave en disclaimer der fortaeller risikoen ved at makke i den kolonne.
        - Med denne aendring, mindsker det ogsaa det saetup man skal lave for at komme i gang.
    - Hvis Balances er et map, kan man hente alle raekkerne 1 gang. Lobe dem i gennem og saa assign Ind baseret paa deres Row.

### Auto generator for Entries

- Entry.Ind skal vaere entriends paagaeldende rakke i config sheets
    - Den skal nok hedde noget andet. Entry.Row?

- 

### callstack

- Auto generator for config
- Update accounting.UpdateBudgetReqs
    - skal ikke bruge rows parameter laengere.
        - Er erstattet af Entry.Row/Ind
            - fungerede ikke?
        

- Test UpdateBudgetReqs, at fixed expenses bliver taelt med
    - vaere sikker paa at flowet er rigtigt, skal ikke laengere kalde UpdateCommonGrp
- fix UpdateCommonGrp
- Rename UpdateCommonGrps til noget med fixedExpense

## Auto generator for config

### Laes Entry col fra sheets

- For hver reakke danner man en entry/post
- Vi har pre-defined gruppern. og indtal man rammer en ny grp-header, saa bliver det defineret som en ny entry
    - Hvordan kan man bestemme om det skal vaere en header eller ej
        - Man kan hente meta data om Cellen. Saa hvis den ser ud paa en bestemt maade saa taeller det som en grp

- Sheets har namedRanges.
    - https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets#NamedRange
    - Dem kan man benytte til at definere poster og over-poster

- Hvordan skal vi saa haandtere at assign mathes til entries
    - Vi kan have en lokal config fil vi laeser fra
    - Det ville nok vaere mere smart at have det i et sheet.

- Saa man kan have et sheet ved siden af, hvor man kan have alle matches.
- Naar man vil update poster og over-poster. definere man bare nye namedRanges, korere init programmet og saa bliver det andet sheet opdateret.
    - Men det skal ogsaa vaere saa man selv kan opdatere det.
    - Det er self ogsaa noget man kunne holde i en database, men her er det nemt for alle at modify
    - Kan ikke bruge database, da de skal have adgang til at skrive nye matches hele tiden.

- Hvordan ser vi forskal paa en GrpName og EntryName
- Fordi vi skal jo bare kunne laese direkte fra Config sheets.
- Man kan laese det direkte fra cellFormat. 
- Saa maa vi definere farver lige som i budgettet som betyder post og over-post
- 

- usecasen hvor man vil tilfoje en ny post, er saa sjaelden at det ikke er saa vaesentling om det er helt optimeret.
    - Vi skal bare have en maade hvor config sheets nemt bliver opdateret naar man saetter budgettet op.
    - Man kan maaske bruge et apps script der holder oje med naar man laver en NamedRange, Der saetter den op i config sheets



### opdater config sheet

- Man kan lave sine egne triggers
    - https://developers.google.com/apps-script/reference/script/spreadsheet-trigger-builder

- Naar man tilfojer en ny raekke, er det en ON_CHANGE event

- lav en knap som henter alle namedRanges over i config sheets
- kan lave en on change trigger, der korere scriptet igen hvis man indsaetter en ny raekke
    - ved ikke helt
    - Denne case sker jo ikke saa tit, saa det ville vaere fint hvis man skulle lave en ny named Range og saa tryk paa knappen

### callStack

- opdater config sheet
- Laes Entry col fra sheets

## August mdr passer ikke med 1000 kr. har 1000 kr mindre end budgettet siger

- Ida ingenor foreningen trakker 2 gang
    - giver - 878 kr
    - Hvor ofte kommer den betaling

- forsikringer giver + 200
- Faelles udgifter giver + 300

- hvis forsikring og faelles passede, saa vil forkellen vaere -1500
- med rettelsen fra ida er der stadig en forskel paa 600 kr

- hverdag stemmer

- bil passer

- indtaegter stemmer

- mad giver stemmer

- opsparing stemmer

## test accounting

### test createGrps

- Vi skal have et sheets.ValueRange som input
- Paa den maade tester vi bare som en black box, og saa tjekker at de rigtige grps kommer ud
- man kan heller ikke rigtig gore noget andet

- Man kan lave en setup funktion, den kan enten lave en mock ValueRange, eller lave en mock getConfig funktion. Hvor man saa har en raekke test configs i et seperaret sheet, som man saa kna kalde i div. test cases.
- Eller skal man lave ValueRange sheets fra bunden, saa er det nok nemmets med test sheets

- Starter med at prove at lave en raa ValueRange variabel

- test mellem rum ved config matches

### call stack

## Lav setup cli argument

- hvad skal man angive
    - budgetSheetId 
    - configSheetId 
    - spreadSheetId 

- Lav guide til hvordan man giver pgm'et adgang til ens speadsheet

- hvor skal man laese input fra
    - Med cli'en burde man bare angive den naar programmet starter
        - Det er jo bare en exe man korere, saa man kan ikke give den input
- hvordan skal vi ogsaa lave setup, naar man ikke kan parse flag 
    - Man kan selv detect om det er forste gang det korere, 
    - i forhold til at saette sheetId og spreadsheetid, burde man kunne angive det naar man builder
        - Jeg tror man kan bruge -ldflags naar man bygger / compiler
            - Man kan bruge -X flaget

- Ogsaa kan man sporge paa hvilken csv fil der skal bruges


- relative path for input file
    - https://forum.golangbridge.org/t/how-to-get-relative-path-from-runtime-caller/15690

- faa abs path til fil, https://stackoverflow.com/questions/17071286/how-can-i-open-files-relative-to-my-gopath 

- Hvad skal man gore med token.json og client_secret
    - De kan jo ikke bare ligge ude i det aabne
    - Og de skal jo ikke vaere i den binaere fil, da de skal vaere lokalt til denne bruger ellers har vi nok it sikkerheds brist
    - i opsaetnings scriptet sorger man faar at lave en mappe
        - der ligger man token og client secret
        - hvordan skal man saa bygge projektet
            - Det kraver jo man har adgang til koden
            - Det skal ikke vaere nodvendigt for mig at bygge selv

- til opsaetning kan man ogsaa angive om det er for 1 eller 2 personer. Hvis man er 2 pers, bliver man spurgt hver gang hvem der skal laves budget for

-

## Lav guide

- naar man laver en kopi, skal man ind og rename funktionen for den virker
    - Hvad sker der naar man trykker paa "implementer" i app script

## build script

- Hvad skal man gore med token.json og client_secret
    - De kan jo ikke bare ligge ude i det aabne
    - Og de skal jo ikke vaere i den binaere fil, da de skal vaere lokalt til denne bruger ellers har vi nok it sikkerheds brist
    - i opsaetnings scriptet sorger man faar at lave en mappe
        - der ligger man token og client secret
        - hvordan skal man saa bygge projektet
            - Det kraver jo man har adgang til koden
            - Det skal ikke vaere nodvendigt for mig at bygge selv

- vi laver et build script for linux og windows, som vi ligger i git repo'et, som man saa kan korere

- Hvordan fungere det med at angive input excrptSheet, naar programmet ligger i usr/local men man kalder det fra desktop.
- Den binaere fil skal vel ogsaa ligge i urs/local/bin el. kan den bare ligge i usr/local/budgetAutomation/bin

- Hvordan skal man lave opsaetning af clinte secret og token
    - syntes den burde komme op af sig selv


## Distribution

- For at jeg kan automatisere noget af det skal der vaere en hjemmeside, som de kan intereagere med i stedet.
- Saa kan jeg bygge en exe fil til dem jeg zipper med client secret og saa er vi good

### callstack

- Hvordan faar man windows til ikke at nakke exe filen med windows defender
- Hvordan kan man automatisk download en client secret
    - Hvis man downloaded det hele fra en hjemmeside, saa kunne man samle det hele i en zip fil man kunne download.
    - Hvad saa med sheetId og spreadSheetId.
        - Dem kunne man saa input i et par felter saa kalder man build scriptet og noget logik til at danne en client secret.
- Hvordan kan man automatisk build systemet

## kor budget for 1 el. 2 personer

- til opsaetning kan man ogsaa angive om det er for 1 eller 2 personer. hvis man er 2 pers, bliver man spurgt hver gang hvem der skal laves budget for

## Brug af pattern i Parser

- Den parser man skal bruge afhaenger af fil type og bank
    - Man bliver nod til at angive bank i opsaetnings scritpet. Da det fortaeller hvilken opsaetning man kan regne med.
    - Naar man har fundet banken, kan programmet selv agere paa baggrund af fil typen
    - bankeg angiver man i opsaetnings scriptet

    - Saa vi skal have et design pattern der understotter dette
        - Det er en form for behaviorial pattern
        - https://refactoring.guru/design-patterns/strategy

- Nordea
    - CSV
    - PDF
- Sparkassen kronjylland
    - PDF

- PDF -> .xlsx
    - Som opforer sig paent

### hvordan skal vi gore 

- Metoden der skal bruges afhaenger baade af bank og fil type

- Skal vi have et interface der hedder parser der implementere ReadInput/ReadExcrpts
    - Her implementere man en funktion per filtype
        - Men saa er der nogen banker der skal implementer funktioner for filtyper de ikke har

- Lave en parser klasse / struct per bank der implementere et interface, med 1 metode parse
    - Der laver man saa et tjek paa fil type der afgore hvilken faktisk parser funktion der skal bruges
    - Hvad opnaar man ved det 
        - At man er garenteret de har en metode parse, som man kan kalde
        - Og saa man kan tjekke paa den som en type

    - Hvordan vaelger man en struct baseret paa banknr
        - Man har vel bare en var af type parser interface. Den saetter man saa bare accordingly

        
- Giver det mening at have en struct, hvis den holder nogen felter 
    - Saa giver det vel bare ekstra kode
    - Giver det noget at have saadan en emytp struct. 
        - Men man skal have den for at kunne bruge et interface.
        - Er interface saa det vaerd?

- Skal man bruge build-tags i forhold til at build UI og cli
    - Det er lidt overkill

- Man kan ligge en init funktion der konvereter -ldflags fra string
    - Hvor skal den init ligge
    - Den er jo generel for alt kode
    - Saa skal de variable man saetter vaere globale, ellers kan man ikke se dem fra cmd
        - Nej, fordi hvis flagene bare er lokalt til cmd pakken saa er det jo fint nok

- at lave (numberVale/stingValue, ser ikke ud til at fungere)
    - Ellers saa er de der bare men sat til nil

- Kan man lave gore saa det kun er vores namedRanges vi faar ind
    - og ikke hente > 6000 rows

### Callstack

## Bedre info og error logger

- lave hjemmestrikket error logger
- hvordan skal vi parse den med til alle funktioner
    - Det nytter ikke noget at have den siddende paa et objekt, som man saa skal lave alle metoderen refere til det 
- Saa skulle man lave loggeren global
    - 

- Man kan bare ligge den i en log pakke

- logging binder sig taet ind til graceful shutdown.
    - De skal jo have en besked at vide naar man fejler
- som binder sig ind i error handling

- start med logging og graceful shutdown

## udgifter staar med negativ fortegn

- Naar man summer sammen, kan man ikke altid *-1 paa. Kun hvis det ikke er "indkomst efter skat"
- saa man skal vide gruppen

- tkan man gore det paa en anden maade
  - 

- Man kan lave en iterator til balance/amounts, med state, saa den ved om disse amounts er indkomst

### lige nu bliver alle celler blanket, hvis ikke det er en entry m en balance

- getAmount giver error, hvis ikke det er en kendt entry
- Men de skal ikke vare blank, dem skal man bare springe over
  - Der er heller ikke nogen grund til at logge dem
- Men hvis entryen findes, men amounts er tom, skal man indsaette blank

- naar man laeser faste udgifter, skal man kun satte ind hvis der staar <> 0