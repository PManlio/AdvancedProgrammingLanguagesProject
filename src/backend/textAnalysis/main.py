# simple test
from typing import Optional
from fastapi import FastAPI

app = FastAPI()

myList = ["manlio", "luca", "helenio"]

@app.get("/")
def root():
    return {"msg":"ciao manlio"}

@app.get("/mylist/{index}")
def getName(index: int, q: Optional[str] = None):
    nome: str = myList[index]
    return {"nome":f"{nome}"}
