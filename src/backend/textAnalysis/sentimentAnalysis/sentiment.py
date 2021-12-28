# ALTERNATIVA: https://pypi.org/project/vader-multi/ e https://www.reddit.com/r/ItalyInformatica/comments/kttnj8/librerie_python_per_sentiment_analysis_in_italiano/

import pandas as pd
from pandas.io.parsers import TextFileReader
from sklearn.feature_extraction.text import CountVectorizer # Converte una collezione di documenti di testo in una matrice di contatori di parola per ogni singola parola

# per separare il dataset in due "tronchi", per fare non solo il training ma anche la validation:
from sklearn.model_selection import train_test_split

# importo pure un algoritmo di allenamento (in questo caso va bene un classificatore bayesiano):
from sklearn.naive_bayes import BernoulliNB

# per la metrica di accuratezza etc etc importiamo:
from sklearn.metrics import accuracy_score

# la read_csv funziona pure coi link
dataframe = pd.read_csv('https://raw.githubusercontent.com/pieroit/corso_ml_python_youtube_pollo/master/movie_review.csv')
#dataframe = pd.read_csv('https://gist.githubusercontent.com/vinid/2286d0ae3d0e39153257b7b6607bf189/raw/0073d5c037dd1daf13991a44a148d957b885d9cd/italian_emotion_classification.csv')


# legge i nomi delle colonne del csv
# print(dataframe.head())

# ora, a noi servono i testi, perché dovremo allenare la nostra IA con questi testi
# quindi se nella X avremo i testi...
X = dataframe['text']
#print(X)

# nella Y avremo il sentimento del testo
Y = dataframe['tag']

vect = CountVectorizer(ngram_range=(1,2)) # il countVectorizer lavora col metodo Bag Of Words; puoi specificare anche gli n-grammi
X = vect.fit_transform(X) # Dato il corpo di un testo, impara il dizionario dei vocaboli che legge e returna una matrice documento-termine

#print(X[:2])
#print(type(X))

# una volta vettorizzati i testi, separiamo il dataset in train, test e validation
X_train, X_test, Y_train, Y_test = train_test_split(X, Y)

# costruisco il mio modello e lo alleno sui dataset di training:
model = BernoulliNB() # chiamo la classe che contiene l'algoritmo di allenamento
model.fit(X_train, Y_train)

# ora raccolgo le mie predizioni:
p_train = model.predict(X_train)
p_test = model.predict(X_test)

print(f"le predizioni - train: {p_train}; test: {p_test}")

# misuro l'accuratezza:
acc_train = accuracy_score(Y_train, p_train)
acc_test = accuracy_score(Y_test, p_test)

print(f"Accuracy in training: {acc_train} - Accuracy in test: {acc_test}")
print(f"i parametri del modello sono {model.get_params(deep=True)}")

# TextFileReader()
example_text = "I feel like i'm going to die sooner or later. Someday, i feel, i won't be able to breath anymore. I hear the sound of the bell ringing in my ear: it's my funeral. I can hear people crying, sobbing..."
#example_text = [].append("Quando la festa é finita, la gente ha iniziato a guidare senza essere in condizioni di farlo. Io ho preso la mia macchina con la certezza che ero sobria. Non potevo immaginare, mamma, ció che mi aspettava… qualcosa di inaspettato!")
#example_text = ["Quando la festa é finita, la gente ha iniziato a guidare senza essere in condizioni di farlo. Io ho preso la mia macchina con la certezza che ero sobria. Non potevo immaginare, mamma, ció che mi aspettava… qualcosa di inaspettato!"]
#example_text = ["I feel like i'm going to die sooner or later. Someday, i feel, i won't be able to breath anymore. I hear the sound of the bell ringing in my ear: it's my funeral. I can hear people crying, sobbing..."]

# la BernoulliNB si aspetta in input un dataframe, quindi tramite pandas ne dobbiamo creare uno:
#newDataFrame = pd.DataFrame(columns=['text'])
#newDataFrame.append([example_text], ignore_index=True)
newDataFrame = pd.DataFrame([example_text], columns=["text"])
print(newDataFrame)
newDataFrame = vect.fit_transform(newDataFrame['text'])
print(model.predict(newDataFrame))


#example_text = vect.fit_transform(example_text) # errore di feature
#print(type(example_text))
#print(example_text)
#print(model.predict(example_text[:]))