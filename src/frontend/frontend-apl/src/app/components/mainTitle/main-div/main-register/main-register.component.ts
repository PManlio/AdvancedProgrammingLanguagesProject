import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-main-register',
  templateUrl: './main-register.component.html',
  styleUrls: ['./main-register.component.css']
})
export class MainRegisterComponent implements OnInit {

  constructor() { }

  public isSeen: boolean = false;

  ngOnInit(): void {
    this.isSeen = false;
  }

  public showRegistrationModal() {
    this.isSeen = this.isSeen ? false : true;
    console.log("faccio qualcosa");
  }

}
