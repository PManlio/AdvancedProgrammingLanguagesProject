import { Component, OnInit } from '@angular/core';
import { PazienteServiceService } from 'src/app/services/paziente-service.service';

@Component({
  selector: 'app-logged-in-paziente',
  templateUrl: './logged-in-paziente.component.html',
  styleUrls: ['./logged-in-paziente.component.css']
})
export class LoggedInPazienteComponent implements OnInit {

  constructor(private pazienteService: PazienteServiceService) { }

  public nome: string = localStorage.getItem("nome")

  ngOnInit(): void {
    this.pazienteService.findPsicologoByEmail("psico@logo.t");
    this.pazienteService.getAllPsicologi()
    // domanda: posso avere la benedizione di cambiare le chiamate nel server in POST?
  }

}
