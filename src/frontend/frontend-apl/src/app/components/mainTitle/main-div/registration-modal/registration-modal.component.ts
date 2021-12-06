import { Component, Input, OnInit, Output, ViewChild, EventEmitter } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Paziente } from 'src/app/interfaces/paziente';
import { RegisterService } from 'src/app/services/register.service';
// import { EventEmitter } from 'events';

@Component({
  selector: 'app-registration-modal',
  templateUrl: './registration-modal.component.html',
  styleUrls: ['./registration-modal.component.css']
})
export class RegistrationModalComponent implements OnInit {

  @ViewChild('registrationModal') thisModal: any;

  @Input() public isSeen: boolean;
  @Output() resetEmitter: EventEmitter<any> = new EventEmitter();

  private paziente: Paziente;

  constructor(private register: RegisterService) { }

  public chiudiForm() {
    this.isSeen = false;
    this.resetEmitter.emit(this.isSeen)
  }

  ngOnInit(): void {
    this.isSeen = false;
  }

  public registerSubmit(registerForm: NgForm) {
    if (registerForm.value.password != registerForm.value.repeatPassword) {
      window.alert("password ripetuta male, riprova");
      return
    }
  }
}
