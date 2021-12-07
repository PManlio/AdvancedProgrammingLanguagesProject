import { Component, Input, OnInit, Output, ViewChild, EventEmitter } from '@angular/core';
import { NgForm } from '@angular/forms';
import { RegistrationInterface } from 'src/app/interfaces/registration-interface';
import { RegisterService } from 'src/app/services/register.service';

@Component({
  selector: 'app-registration-modal',
  templateUrl: './registration-modal.component.html',
  styleUrls: ['./registration-modal.component.css']
})
export class RegistrationModalComponent implements OnInit {

  @ViewChild('registrationModal') thisModal: any;
  private repeatPassword: string;

  @Input() public isSeen: boolean;
  @Output() resetEmitter: EventEmitter<any> = new EventEmitter();

  private registrationInterface: RegistrationInterface;

  constructor(private register: RegisterService) { }

  public chiudiForm() {
    this.isSeen = false;
    this.resetEmitter.emit(this.isSeen)
  }

  ngOnInit(): void {
    this.isSeen = false;
  }

  public registerSubmit(registerForm: NgForm, confirmPassword: any) {

    this.repeatPassword = confirmPassword.value;
    this.registrationInterface = JSON.parse(JSON.stringify(registerForm.value));

    if (this.registrationInterface.password != this.repeatPassword) {
      window.alert("password ripetuta male, riprova");
      return
    }

    // proseguire con il registration
  }
}
