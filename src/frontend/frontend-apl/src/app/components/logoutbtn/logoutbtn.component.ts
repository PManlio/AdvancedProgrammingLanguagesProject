import { Component, OnInit } from '@angular/core';
import { JwtService } from 'src/app/services/jwt.service';

@Component({
  selector: 'app-logoutbtn',
  templateUrl: './logoutbtn.component.html',
  styleUrls: ['./logoutbtn.component.css']
})
export class LogoutbtnComponent implements OnInit {

  constructor(private jwt: JwtService) { }

  ngOnInit(): void { }

  public logout() {
    this.jwt.removeJWT();
  }
}
