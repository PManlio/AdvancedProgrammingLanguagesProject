import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';

@Component({
  selector: 'app-main-div',
  templateUrl: './main-div.component.html',
  styleUrls: ['./main-div.component.css']
})
export class MainDivComponent implements OnInit {

  public email: string;
  public password: string;

  constructor() { }

  onSubmit(formInputs: NgForm) {
    this.email = formInputs.value.email;
    this.password = formInputs.value.password;
    console.log(this.email, this.password);
  }

  ngOnInit(): void {
  }

}
