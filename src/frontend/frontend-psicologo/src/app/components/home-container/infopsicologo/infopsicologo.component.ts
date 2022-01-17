import { Component, OnInit } from '@angular/core';
import { UserInfoService } from 'src/app/services/user-info.service';

@Component({
  selector: 'app-infopsicologo',
  templateUrl: './infopsicologo.component.html',
  styleUrls: ['./infopsicologo.component.css']
})
export class InfopsicologoComponent implements OnInit {

  constructor(private user: UserInfoService) { }

  public nomePsicologo: string;

  ngOnInit(): void {
    this.user.nome.subscribe(val => this.nomePsicologo = val);
  }

}
