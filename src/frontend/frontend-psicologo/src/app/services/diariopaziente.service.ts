import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { myEnv } from 'src/environments/myEnv';
import { InfoDiario } from '../interfaces/infodiario';

@Injectable({
  providedIn: 'root'
})
export class DiariopazienteService {

  private diaryServerUrl: string = `${myEnv.textServerUrl}/paziente`;
  private myheaders = {'Content-Type':'application/json'};

  constructor(private httpClient: HttpClient) { }

  public getFullDiaryOfPatient(email: string) {
    return this.httpClient.get<InfoDiario[]>(`${this.diaryServerUrl}/fulldiary/${email}`, { headers: this.myheaders })
  }

  public getSentimentoMedio(email: string) {
    return this.httpClient.get(`${this.diaryServerUrl}/metrics/meansentiment/${email}`, {headers: this.myheaders})
  }
}
