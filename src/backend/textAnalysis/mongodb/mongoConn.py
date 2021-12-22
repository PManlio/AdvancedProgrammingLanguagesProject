# installare pymongo e python-dotenv; fastAPI e "unvicorn[standard]"
import pymongo
from dotenv import load_dotenv
import os

load_dotenv('../.env')

mongousr: str = os.getenv("MONGOUSR")
mongopsw: str = os.getenv("MONGOPSW")
collection: str = "diaries"

client = pymongo.MongoClient(f"mongodb+srv://{mongousr}:{mongopsw}@cluster0.nu9dh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
db = client.test