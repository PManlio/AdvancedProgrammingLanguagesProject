# simple test
from typing import Optional
from fastapi import FastAPI, Request
from starlette.middleware.cors import CORSMiddleware
# from fastapi.middleware.cors import CORSMiddleware

import mongodb.mongoConn as mongoConn
from models.mongoFile import DiaryPage
from models.myBaseModel import BodyItem

app = FastAPI()

origins = ["*"]
app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=['*'],
    allow_headers=['*']
)

myList = ["manlio", "luca", "helenio"]

@app.get("/")
def root():
    print(mongoConn.printConnString())
    return {"msg": "ciao manlio"}


@app.get("/mylist/{index}")
def getName(index: int, q: Optional[str] = None):
    nome: str = myList[index]
    return {"nome": f"{nome}"}

@app.post("/example/post")
def postExample():
    mongoConn.simplePost()

@app.post("/diary")
def postDiary(body: BodyItem):
    diary = DiaryPage(str(body.mailPaziente), str(body.text))
    mongoConn.postDiary(diary.getJSONInfo())