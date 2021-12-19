import { Component, OnInit } from '@angular/core';
import { Psicologo } from 'src/app/interfaces/psicologo';

@Component({
  selector: 'app-psicologotrovato',
  templateUrl: './psicologotrovato.component.html',
  styleUrls: ['./psicologotrovato.component.css']
})
export class PsicologotrovatoComponent implements OnInit {

  public psicologo: Psicologo = {
    nome: "giulio",
    cognome: "macchevino",
    email: "giulio.m@asd.t",
    cellulare: "123455",
    citta: "ct"
  };

  constructor() { }

  ngOnInit(): void {
  }

}
