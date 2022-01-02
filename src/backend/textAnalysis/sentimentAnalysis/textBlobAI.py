from textblob import TextBlob
from googletrans import Translator

translator = Translator()

def translateToEnglish(text_it: str)-> str:
    return translator.translate(text_it, dest='en', src='it').text

def sentiment(text_eng: str)-> dict:
    return {"polarity": TextBlob(text_eng).sentiment[0], "subjectivity": TextBlob(text_eng).sentiment[1]}

'''
- Polarity è un float compreso tra [-1,1], dove -1 indica un sentimento negativo mentre +1 indica un sentimento positivo. 
- Subjectivity è un altro float compreso tra [0,1] che indica quanto la frase faccia riferimento ad un'opinione personale.
'''

# frase esempio
#sentence_IT: str = "Fin da bambina, sembra assurdo, ma è così. Ricordo che mi isolavo e piangevo pensando sempre la stessa cosa: “Nessuno mi vuole bene”. Ora ho 40 anni e la mia vita è sempre stata un’altalena di malesseri più o meno consapevoli. Eppure gli altri pensano che io sia una donna forte, sempre sorridente, “così solare” mi dicono. L’unico modo che ho per stare bene è quello di non lasciarmi coinvolgere emotivamente. Ma non sempre è possibile e quando capita (nel lavoro, negli affetti, nelle relazioni familiari) e vivo anche una minima delusione, è un disastro. Penso che farei meglio a non esserci, a non vivere, a lasciar perdere tutto. Sono stata da due psicanalisti diversi, uno che lavora in un ospedale pubblico e uno privato (da un euro al minuto) ma, incoraggiamenti a parte, non sono riusciti a indicarmi una via d’uscita. Ho anche seguito una cura a base di farmaci, soprattutto per riuscire a dormire, che mi ha fatto stare meglio, ma quando l’ho sospesa (dopo tre mesi) è tornato tutto esattamente come prima. Sono certa che ne uscirò, com’è accaduto altre volte, usando la ragione, facendomi guidare dal raziocinio e dal buon senso, distinguendo quello che è giusto da quello che è sbagliato. Devo staccare il cuore, però, e vivere freddamente perché, se mi lascio travolgere dai sentimenti e dalle passioni, resto delusa, sto male e piango, non dormo, non voglio vivere. Chissà, forse ha ragione chi ha detto: “Ma gli idioti perché non soffrono di depressione?”. Forse è così e basta, come avere i capelli biondi o il naso storto, forse la depressione non è una malattia, ma è il naso storto. Qualche volta penso che avrei solo bisogno di non essere sola, di vivere come si faceva un tempo in comunità, tutti insieme con i nonni, i cugini, i parenti e i vicini, una vita più semplice e povera, ma forse più ricca di umanità. Nonostante tutto, io continuo a sperare…"
#print(sentiment(translateToEnglish(sentence_IT)))