import datetime

class Response(object):
    def __init__(self, mailPaziente: str, text: str, sentiment: dict, date: datetime.datetime):
        self.mailPaziente: str = mailPaziente
        self.text: str = text
        self.sentiment: dict = sentiment
        self.date: datetime.datetime = date
