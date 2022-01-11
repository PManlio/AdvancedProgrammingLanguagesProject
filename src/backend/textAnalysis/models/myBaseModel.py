from pydantic import BaseModel

class BodyItem(BaseModel):
    mailPaziente: str
    text: str

class BodyRequest(BaseModel):
    mailPaziente: str
    date: str