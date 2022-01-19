# installare pymongo e python-dotenv; fastAPI e "unvicorn[standard]"
import pymongo
from dotenv import load_dotenv
from pathlib import Path
import os
import json

from models.mongoFile import DiaryPage
from models.response import Response
from metrics.metric import gradientSentiment, meanSentiment

pathToEnv = Path("./.env")

load_dotenv(dotenv_path=pathToEnv)

mongousr: str = os.getenv('MONGOUSR')
mongopsw: str = os.getenv("MONGOPSW")
database: str = os.getenv("DBNAME")
collection: str = os.getenv("DBCOLLECTION")

client = pymongo.MongoClient(
    f"mongodb+srv://{mongousr}:{mongopsw}@cluster0.nu9dh.mongodb.net/{database}?retryWrites=true&w=majority")
db = client[database]
dbcollection = db[collection]

def printConnString():
    return (f"questi sono i dati: {mongousr}, {mongopsw}, {database}, {collection}")

def simplePost():
    example = {"nome":"helenio", "cognome":"palmeri"}
    dbcollection.insert_one(example)

def postDiary(JSONdiary: str):
    dbcollection.insert_one(json.loads(JSONdiary))

def getAllUserDiariesByUserEmail(email: str):
    try:
        res = dbcollection.find({"emailPaziente": email})
        return [Response(document["emailPaziente"], document["text"], document["sentiment"], document["date"]) for document in res]
        # return [(document["emailPaziente"], document["sentiment"], document["date"]) for document in res]
    except pymongo.errors.OperationFailure:
        return json.loads({"error": "no document found"})

def getUserDiaryByDate(email: str, date: str):
    try:
        found = dbcollection.find_one({"emailPaziente": email, "date": date})
        return Response(found["emailPaziente"], found["text"], found["sentiment"], found["date"])
    except pymongo.errors.OperationFailure:
        return json.loads({"error": "no document found"})

def getAnalysisOfUserSentiment(email: str):
    try:
        res = dbcollection.find({"emailPaziente": email})
        sentimentArray = [document["sentiment"] for document in res]
        return meanSentiment([pol["polarity"] for pol in sentimentArray])
    except pymongo.errors.OperationFailure:
        return json.loads({"error": "no document found"})

def getGradientOfUserSentiment(email: str):
    res = dbcollection.find({"emailPaziente": email})
    
    results = [{"sentiment":document["sentiment"], "date":document['date']} for document in res]
    sentimentArray = [snt["sentiment"] for snt in results]
    polarityArray = [pol["polarity"] for pol in sentimentArray]
    datesArray = [dt["date"] for dt in results]
    
    return gradientSentiment(polarityArray, datesArray)