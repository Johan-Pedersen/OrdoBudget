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

#### vis current excrpt

- Vi skal have en maade at vise current excrpt paa
    - Dette skal vi bruge databinding
- Vi skal vide hvordan man hugger excrptet op hvis vi skal laese det fra formen.
    - Men behover vi det, vi har jo et current excrpt. Det laver vi bare en to string paa for at vise det i en 1-way databinding til feltet.
    - Vi kan knytte den til vores current excrpt og bindingen tager bare toString metoden.
- Naar vi skal bruge excrptet, saa bruger vi bare vores struct

- Hvordan vil man saa gore det med et current excrpt
    - design patterns?
    - local currentExcrpt i controlleren, som man bare skal husk at update naar man kalder funktioner fra controlleren
    - Men er dette noget der burde ligge i model?
        - Fordi saa faar controllen lige pludselig logik og det skal den jo ikke
        - Men der er ikke noget sted der giver mening at ligge det. Fordi modellen giver bare metoderne der kan kaldes fra controlleren. Men der er ikke noget sted at "persistere" / holde data.
        - Controlleren manipulere jo model. Saa "current excrpt" skal jo ligge i model.
        - Men det behover maaske ikke vaere en decideret varabel.
            - Det er lidt den funktionelle maade at gore det paa
        - Maaske kan man bare hente den man er kommet til?
        - hvad gor vi nu?
            - Vi henter alle values fra sheets og saa lober dem igennem 1 af gangen.
            - Men hvordan ville man gore det 
                - Det er jo lidt af en sideeffekt hvis model skulle update viewet
            

- Skal man give model/view med som argument i controlleren og model/view skal ogsaa have en instance af controlleren.
    - Saa skal man i controlleren definere et model / view interface, som exposer de funktioner der skal vaere tilgaengelig i controlleren
                - Men det kan jo heller ikke rigtig gaa igennem controlleren, ved at moddel kaleder en funktion i controlleren. Som updater viewet.
                    - Det ville ogsaa give cirkulaere dependenciens
                
                - Man henter jo en hel batch og lober igennem dem.
                - Man kan godt update view fra model. Der kommer ingen 

## omdan projekt struktur

- https://github.com/golang-standards/project-layout

- Lav cmd, internal, ui, pakker
- i internal, lav pakker baseret paa hvad denne del af koden giver.
    - Lidt "grop by context" i DDD verdenen


