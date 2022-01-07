from fastapi import FastAPI
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

@app.get("/")
def root():
    print(mongoConn.printConnString())
    return {"msg": "ciao manlio"}

@app.post("/example/post")
def postExample():
    mongoConn.simplePost()

@app.post("/diary")
def postDiary(body: BodyItem):
    diary = DiaryPage(str(body.mailPaziente), str(body.text))
    mongoConn.postDiary(diary.getJSONInfo())

@app.get("/paziente/fulldiary/{email}")
def getAllDiariesOfPatient(email: str):
    return mongoConn.getAllUserDiariesByUserEmail(email)

@app.get("/paziente/metrics/meansentiment/{email}")
def getMeanSentimentOfPatient(email: str):
    return mongoConn.getAnalysisOfUserSentiment(email)

@app.get("/paziente/metrics/gradientsentiment/{email}")
def getGradientSentimentOfPatient(email: str):
    return mongoConn.getGradientOfUserSentiment(email)