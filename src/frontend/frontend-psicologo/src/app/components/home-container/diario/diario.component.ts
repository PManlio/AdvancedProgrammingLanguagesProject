import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { InfoDiario } from 'src/app/interfaces/infodiario';
import { DiariopazienteService } from 'src/app/services/diariopaziente.service';

@Component({
  selector: 'app-diario',
  templateUrl: './diario.component.html',
  styleUrls: ['./diario.component.css']
})
export class DiarioComponent implements OnInit {

  public emailPaziente: string;
  private subscriptionToRouteParams: any;

  public infoDiario: InfoDiario[];
  public sentimentoMedio: number;

  constructor(private router: ActivatedRoute, private diario: DiariopazienteService) {
    // this.diario.getFullDiaryOfPatient(this.emailPaziente).subscribe(pagine => {console.log(pagine); this.infoDiario = pagine; })
  }

  ngOnInit(): void {
    this.subscriptionToRouteParams = this.router.params.subscribe(params => {
      this.emailPaziente = params["emailPaziente"];
      this.diario.getFullDiaryOfPatient(this.emailPaziente).subscribe(pagine => {/* console.log(pagine); */ this.infoDiario = pagine; });
      this.diario.getSentimentoMedio(this.emailPaziente).subscribe(val => this.sentimentoMedio = val["averageSentiment"])
    });
  }

}
