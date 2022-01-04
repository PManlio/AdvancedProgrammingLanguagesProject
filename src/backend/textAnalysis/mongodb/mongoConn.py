# installare pymongo e python-dotenv; fastAPI e "unvicorn[standard]"
import pymongo
from dotenv import load_dotenv
from pathlib import Path
import os
import json

from models.mongoFile import DiaryPage

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