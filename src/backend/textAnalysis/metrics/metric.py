import numpy as np
import datetime

# la data arriva come stringa, perciò prima deve essere convertita in datetime
def dateTimeToInteger(strdate: str) -> int:
    date = datetime.datetime.strptime(strdate,'%Y-%m-%d %H:%M:%S.%f')
    return 10000*date.year + 100*date.month + date.day

def meanSentiment(sentiment: list):
    return {"averageSentiment": np.average(sentiment)}


"""
    Tentativo di metrica aggiuntiva, non aggiunto al funzionamento del server TextAnalysis
    I valori prodotti dall'output sono troppo piccoli per essere letti correttamente;
"""
def gradientSentiment(sentiment: list, dates: list):
    # trasformo la lista di date in una lista di interi
    # ne prendo il primo elemento e l'ultimo e faccio: ultimo - primo
    # così ottengo una costante che mi dice quanto tempo è passato dalla
    # prima pagina del diario all'ultima
    intDateList = [dateTimeToInteger(date) for date in dates]
    # totalTime = intDateList[len(intDateList) - 1] - intDateList[0]

    # print(sentiment, intDateList)

    # gradient = np.diff(sentiment) / np.diff(intDateList)
    
    gradient = (sum(sentiment)/len(sentiment)) / (sum(intDateList)/len(intDateList))
    print(gradient)
    return {"gradientSentiment": gradient}
