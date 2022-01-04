import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { myEnv } from 'src/environments/myEnv';

@Injectable({
  providedIn: 'root'
})
export class DiaryService {

  private diaryServerUrl: string = `${myEnv.textServerUrl}/diary`;
  // private headers = new HttpHeaders(myEnv.headers);

  constructor(private httpClient: HttpClient) { }

  public postDiary(email: string, text: string) {

    let myheaders = {'Content-Type':'application/json'};

    let body = { "mailPaziente": email, "text": text }
    return this.httpClient.post(this.diaryServerUrl, JSON.stringify(body), { headers: myheaders }).toPromise().catch(err => alert(err.status + ' - ' + err.statusText))
  }
}
