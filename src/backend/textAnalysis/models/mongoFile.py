import json
import datetime

import sys
import os
sys.path.append(os.path.abspath('../sentimentAnalysis'))

from sentimentAnalysis.textBlobAI import sentiment

class DiaryPage:
    def __init__(self, emailPaziente: str, text: str) -> None:
        self.emailPaziente = emailPaziente
        self.text = text
        self.sentiment: dict = sentiment(self.text)
        self.date: datetime.datetime = datetime.datetime.now()

    def sentiment(self):
        return sentiment(self.text)

    def getJSONInfo(self):
        return json.dumps({"emailPaziente": self.emailPaziente, "text": self.text, "sentiment": self.sentiment, "date": str(self.date)})

#example = DiaryPage("m@a.t", "Fin da bambina, sembra assurdo, ma è così. Ricordo che mi isolavo e piangevo pensando sempre la stessa cosa: “Nessuno mi vuole bene”. Ora ho 40 anni e la mia vita è sempre stata un’altalena di malesseri più o meno consapevoli. Eppure gli altri pensano che io sia una donna forte, sempre sorridente, “così solare” mi dicono. L’unico modo che ho per stare bene è quello di non lasciarmi coinvolgere emotivamente. Ma non sempre è possibile e quando capita (nel lavoro, negli affetti, nelle relazioni familiari) e vivo anche una minima delusione, è un disastro. Penso che farei meglio a non esserci, a non vivere, a lasciar perdere tutto. Sono stata da due psicanalisti diversi, uno che lavora in un ospedale pubblico e uno privato (da un euro al minuto) ma, incoraggiamenti a parte, non sono riusciti a indicarmi una via d’uscita. Ho anche seguito una cura a base di farmaci, soprattutto per riuscire a dormire, che mi ha fatto stare meglio, ma quando l’ho sospesa (dopo tre mesi) è tornato tutto esattamente come prima. Sono certa che ne uscirò, com’è accaduto altre volte, usando la ragione, facendomi guidare dal raziocinio e dal buon senso, distinguendo quello che è giusto da quello che è sbagliato. Devo staccare il cuore, però, e vivere freddamente perché, se mi lascio travolgere dai sentimenti e dalle passioni, resto delusa, sto male e piango, non dormo, non voglio vivere. Chissà, forse ha ragione chi ha detto: “Ma gli idioti perché non soffrono di depressione?”. Forse è così e basta, come avere i capelli biondi o il naso storto, forse la depressione non è una malattia, ma è il naso storto. Qualche volta penso che avrei solo bisogno di non essere sola, di vivere come si faceva un tempo in comunità, tutti insieme con i nonni, i cugini, i parenti e i vicini, una vita più semplice e povera, ma forse più ricca di umanità. Nonostante tutto, io continuo a sperare…")

#print(example.getJSONInfo())