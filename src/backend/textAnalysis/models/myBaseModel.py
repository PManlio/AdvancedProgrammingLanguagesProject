from pydantic import BaseModel

class BodyItem(BaseModel):
    mailPaziente: str
    text: str