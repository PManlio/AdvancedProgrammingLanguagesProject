import { Component, OnInit } from '@angular/core';
import { JwtService } from 'src/app/services/jwt.service';

@Component({
  selector: 'app-logout-btn',
  templateUrl: './logout-btn.component.html',
  styleUrls: ['./logout-btn.component.css']
})
export class LogoutBtnComponent implements OnInit {

  constructor(private jwt: JwtService) { }

  ngOnInit(): void {
  }

  public logout() {
    this.jwt.removeJWT();
  }

}
