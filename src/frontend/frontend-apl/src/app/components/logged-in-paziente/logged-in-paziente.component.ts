import { Component, OnInit } from '@angular/core';
import { PazienteServiceService } from 'src/app/services/paziente-service.service';
import { UserInfoService } from 'src/app/services/user-info.service';

@Component({
  selector: 'app-logged-in-paziente',
  templateUrl: './logged-in-paziente.component.html',
  styleUrls: ['./logged-in-paziente.component.css']
})
export class LoggedInPazienteComponent implements OnInit {

  public nome: string;
  
  constructor(private pazienteService: PazienteServiceService, private userInfo: UserInfoService) {
    this.userInfo.nome.subscribe(() => this.nome = this.userInfo.localNome)
  }

  ngOnInit(): void {  
    this.pazienteService.findPsicologoByEmail("psico@logo.t");
    this.pazienteService.getAllPsicologi()
    // domanda: posso avere la benedizione di cambiare le chiamate nel server in POST?
  }
}
